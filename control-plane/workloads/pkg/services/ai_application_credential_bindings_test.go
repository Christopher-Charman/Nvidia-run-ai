/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package services

import (
	"encoding/base64"
	"encoding/json"

	"github.com/dsx-ai-factory/run-ai/control-plane/workloads/internal/generated/pkg/api"
	assetsClient "github.com/dsx-ai-factory/run-ai/control-plane/workloads/pkg/clients/assets"
	testing "github.com/dsx-ai-factory/run-ai/control-plane/workloads/pkg/test"
	"github.com/dsx-ai-factory/run-ai/runai-application-packages/convutils"
	assetsApi "github.com/dsx-ai-factory/run-ai/runai-clients/assets-service-client/generated"
	"github.com/dsx-ai-factory/run-ai/runai-cluster/common/cluster-capabilities/checker"
	clusterClient "github.com/dsx-ai-factory/run-ai/runai-cluster/common/cluster-capabilities/clients/cluster"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AI Application runtime credential bindings", func() {
	const assetID = "12345678-abcd-4000-8000-000000000001"
	var request api.AiApplicationCreationRequest
	var credential assetsApi.CredentialsListResponseEntry
	ctx := testing.SetContext(testing.TestTenant, testing.TestUser, nil)
	BeforeEach(func() {
		clusterClient.SetMock(clusterClient.NewSimpleMockClient().WithCluster(testing.ClusterUuid1, testing.BuildDisplayedCluster(testing.ClusterUuid1, "3.1.0", nil)))
		checker.InitClusterCapabilitiesChecker("", testing.GetDefaultClustersCacheConfig(), "")
		credential = assetsApi.CredentialsListResponseEntry{Meta: assetsApi.AssetMeta{
			Id: assetID, Name: "selected", Kind: assetsApi.ASSETKIND_NGC_API_KEY, Scope: assetsApi.SCOPE_TENANT, CreatedBy: testing.TestUser,
		}}
		request = api.AiApplicationCreationRequest{
			Name: "my-app", ClusterId: testing.ClusterUuid1, ProjectId: "1",
			Spec: api.AiApplicationSpec{CredentialBindings: []api.AiApplicationCredentialBinding{
				{Variable: "global.ngcApiSecret.name", AssetId: assetID, CredentialType: "ngc-api-key"},
			}},
		}
	})
	process := func() (bool, *api.Error) {
		assetsClient.SetMock(assetsClient.NewSimpleMock().WithCredentials([]assetsApi.CredentialsListResponseEntry{credential}))
		return processAIApplicationCredentialBindings(ctx, &request, "department")
	}

	It("keeps IDs separate from values and derives ownership on the server", func() {
		request.Spec.Values = map[string]interface{}{"unrelated": "preserved"}
		request.Spec.CredentialBindings[0].UserCredential = true
		request.Spec.CredentialBindings[0].ExistingSecretName = "untrusted-client-name"
		personal, err := process()
		Expect(err).To(BeNil())
		Expect(personal).To(BeFalse())
		Expect(request.Spec.CredentialBindings[0].UserCredential).To(BeFalse())
		Expect(request.Spec.CredentialBindings[0].ExistingSecretName).To(BeEmpty())
		Expect(request.Spec.Values).To(Equal(map[string]interface{}{"unrelated": "preserved"}))
	})

	It("accepts two uses of the same personal NGC asset and redacts its IDs for another viewer", func() {
		credential.Meta.Private.Set(convutils.Ref(true))
		request.Spec.CredentialBindings = append(request.Spec.CredentialBindings, api.AiApplicationCredentialBinding{
			Variable: "global.imagePullSecret.name", AssetId: assetID, CredentialType: "image-pull-secret",
		})
		personal, err := process()
		Expect(err).To(BeNil())
		Expect(personal).To(BeTrue())
		Expect(request.Spec.CredentialBindings[0].UserCredential).To(BeTrue())
		Expect(request.Spec.CredentialBindings[1].UserCredential).To(BeTrue())
		app := &api.AiApplication{CreatedBy: testing.TestUser, Spec: request.Spec}
		sanitizeAIApplicationResponse(testing.SetContext(testing.TestTenant, testing.AnotherUser, nil), app)
		Expect(app.Spec.CredentialBindings[0].AssetId).To(BeEmpty())
		Expect(app.Spec.CredentialBindings[1].AssetId).To(BeEmpty())
	})

	It("rejects deleted and inaccessible selections with a reselection message", func() {
		assetsClient.SetMock(assetsClient.NewSimpleMock())
		_, err := processAIApplicationCredentialBindings(ctx, &request, "department")
		Expect(err).NotTo(BeNil())
		Expect(err.Message).To(ContainSubstring("no longer exists or is not accessible. Please reselect"))
	})

	DescribeTable("interprets the credential API deletion timestamp",
		func(metadataJSON string, deleted bool) {
			var responseMetadata assetsApi.AssetMeta
			Expect(json.Unmarshal([]byte(metadataJSON), &responseMetadata)).To(Succeed())
			credential.Meta.DeletedAt = responseMetadata.DeletedAt
			_, err := process()
			if deleted {
				Expect(err).NotTo(BeNil())
				Expect(err.Message).To(ContainSubstring("no longer exists or is not accessible. Please reselect"))
			} else {
				Expect(err).To(BeNil())
			}
		},
		Entry("omitted timestamp is active", `{}`, false),
		Entry("null timestamp is active", `{"deletedAt":null}`, false),
		Entry("zero timestamp returned for active assets is active", `{"deletedAt":"0001-01-01T00:00:00Z"}`, false),
		Entry("nonzero timestamp is deleted", `{"deletedAt":"2026-09-28T10:00:00Z"}`, true),
	)

	It("drops overridden literals and personal placeholders before persistence and legacy resolution", func() {
		request.Spec.Values = map[string]interface{}{"global": map[string]interface{}{"ngcApiSecret": map[string]interface{}{
			"name": "USERCREDENTIAL:deleted-personal-credential", "create": "test-only-sensitive-literal", "keep": true,
		}}}
		_, err := process()
		Expect(err).To(BeNil())
		used, err := processAIApplicationCredentials(ctx, request.Name, &request.Spec)
		Expect(err).To(BeNil())
		Expect(used).To(BeFalse())
		raw, marshalErr := json.Marshal(request.Spec.Values)
		Expect(marshalErr).NotTo(HaveOccurred())
		Expect(string(raw)).To(MatchJSON(`{"global":{"ngcApiSecret":{"name":null,"create":null,"keep":true}}}`))
	})

	It("rejects another user's personal credential and credentials from another project", func() {
		credential.Meta.Private.Set(convutils.Ref(true))
		credential.Meta.CreatedBy = testing.AnotherUser
		_, err := process()
		Expect(err).NotTo(BeNil())
		credential.Meta.Private.Set(convutils.Ref(false))
		credential.Meta.Scope = assetsApi.SCOPE_PROJECT
		credential.Meta.ClusterId.Set(convutils.Ref(request.ClusterId))
		credential.Meta.ProjectId.Set(convutils.Ref(int32(2)))
		_, err = process()
		Expect(err).NotTo(BeNil())
		Expect(err.Message).To(ContainSubstring("target project"))
	})

	DescribeTable("enforces credential scope against the target project",
		func(scope assetsApi.Scope, clusterID, departmentID string, projectID int32, valid bool) {
			credential.Meta.Scope = scope
			credential.Meta.ClusterId.Set(convutils.Ref(clusterID))
			credential.Meta.DepartmentId.Set(convutils.Ref(departmentID))
			credential.Meta.ProjectId.Set(convutils.Ref(projectID))
			_, err := process()
			Expect(err == nil).To(Equal(valid))
		},
		Entry("tenant", assetsApi.SCOPE_TENANT, "", "", int32(0), true),
		Entry("matching cluster", assetsApi.SCOPE_CLUSTER, testing.ClusterUuid1, "", int32(0), true),
		Entry("different cluster", assetsApi.SCOPE_CLUSTER, testing.ClusterUuid2, "", int32(0), false),
		Entry("matching department", assetsApi.SCOPE_DEPARTMENT, testing.ClusterUuid1, "department", int32(0), true),
		Entry("different department", assetsApi.SCOPE_DEPARTMENT, testing.ClusterUuid1, "other", int32(0), false),
		Entry("same department ID in another cluster", assetsApi.SCOPE_DEPARTMENT, testing.ClusterUuid2, "department", int32(0), false),
		Entry("matching project", assetsApi.SCOPE_PROJECT, testing.ClusterUuid1, "department", int32(1), true),
		Entry("same project ID in another cluster", assetsApi.SCOPE_PROJECT, testing.ClusterUuid2, "department", int32(1), false),
		Entry("unknown scope", assetsApi.Scope("unknown"), testing.ClusterUuid1, "department", int32(1), false),
	)

	DescribeTable("enforces compatible Secret representations",
		func(kind assetsApi.AssetKind, fieldType string, valid bool) {
			credential.Meta.Kind = kind
			request.Spec.CredentialBindings[0].CredentialType = fieldType
			_, err := process()
			Expect(err == nil).To(Equal(valid))
		},
		Entry("NGC API", assetsApi.ASSETKIND_NGC_API_KEY, "ngc-api-key", true),
		Entry("NGC image pull", assetsApi.ASSETKIND_NGC_API_KEY, "image-pull-secret", true),
		Entry("NGC generic", assetsApi.ASSETKIND_NGC_API_KEY, "generic", true),
		Entry("registry image pull", assetsApi.ASSETKIND_DOCKER_REGISTRY, "image-pull-secret", true),
		Entry("registry is not generic", assetsApi.ASSETKIND_DOCKER_REGISTRY, "generic", false),
		Entry("generic is not NGC", assetsApi.ASSETKIND_GENERIC_SECRET, "ngc-api-key", false),
		Entry("password unsupported", assetsApi.ASSETKIND_PASSWORD, "generic", false),
		Entry("old enum rejected", assetsApi.ASSETKIND_DOCKER_REGISTRY, "docker-registry", false),
	)

	It("normalizes base64 values while preserving unrelated overrides", func() {
		request.Spec.ValuesBase64 = base64.StdEncoding.EncodeToString([]byte("pulls:\n  - name: existing\n    extra: preserve\n"))
		_, err := process()
		Expect(err).To(BeNil())
		Expect(request.Spec.ValuesBase64).To(BeEmpty())
		Expect(request.Spec.Values["pulls"]).To(Equal([]interface{}{map[string]interface{}{"name": "existing", "extra": "preserve"}}))
	})

	It("rejects conflicting destinations and mutually exclusive value inputs", func() {
		request.Spec.CredentialBindings = append(request.Spec.CredentialBindings, api.AiApplicationCredentialBinding{
			Variable: "global.ngcApiSecret.create", AssetId: assetID, CredentialType: "generic",
		})
		_, err := process()
		Expect(err).NotTo(BeNil())
		request.Spec.CredentialBindings = request.Spec.CredentialBindings[:1]
		request.Spec.Values = map[string]interface{}{}
		request.Spec.ValuesBase64 = "e30="
		_, err = process()
		Expect(err).NotTo(BeNil())
	})

	It("gates older clusters before their agents could ignore bindings", func() {
		clusterClient.SetMock(clusterClient.NewSimpleMockClient().WithCluster(testing.ClusterUuid1, testing.BuildDisplayedCluster(testing.ClusterUuid1, "3.0.0", nil)))
		checker.InitClusterCapabilitiesChecker("", testing.GetDefaultClustersCacheConfig(), "")
		_, err := process()
		Expect(err).NotTo(BeNil())
	})

	It("takes existing Secret names only from authorized asset metadata and never copies data", func() {
		credential.Meta.Kind = assetsApi.ASSETKIND_GENERIC_SECRET
		request.Spec.CredentialBindings[0].CredentialType = "generic"
		credential.Spec.GenericSecret.Set(&assetsApi.GenericSecretSpec{ExistingSecretName: *assetsApi.NewNullableString(convutils.Ref("preexisting-reference"))})
		_, err := process()
		Expect(err).To(BeNil())
		Expect(request.Spec.CredentialBindings[0].ExistingSecretName).To(Equal("preexisting-reference"))
		raw, marshalErr := json.Marshal(request.Spec)
		Expect(marshalErr).NotTo(HaveOccurred())
		Expect(string(raw)).NotTo(ContainSubstring("keyValuePairs"))
	})
})
