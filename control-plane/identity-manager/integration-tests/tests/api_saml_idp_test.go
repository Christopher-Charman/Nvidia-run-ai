/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package tests

import (
	"context"
	"encoding/base64"
	"net/http"
	"testing"

	openapi "github.com/dsx-ai-factory/run-ai/control-plane/identity-manager/pkg/client/generated"
	"github.com/stretchr/testify/require"
)

var samlMetadataUrl string

const samlEntityId = "urn:example:idp"

// samlDescriptorMissingSsoBinding is a descriptor as a real identity provider
// publishes one, except that its third SingleSignOnService carries no Binding
// attribute. Keycloak requires that attribute on every endpoint, so the
// document passes structural validation here and is then rejected by
// Keycloak's stricter parser.
const samlDescriptorMissingSsoBinding = `<md:EntityDescriptor xmlns="urn:oasis:names:tc:SAML:2.0:metadata" xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" xmlns:saml="urn:oasis:names:tc:SAML:2.0:assertion" xmlns:ds="http://www.w3.org/2000/09/xmldsig#" entityID="urn:example:idp">
  <md:IDPSSODescriptor WantAuthnRequestsSigned="true" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
    <md:KeyDescriptor use="signing">
      <ds:KeyInfo>
        <ds:KeyName>saml-metadata-test</ds:KeyName>
        <ds:X509Data>
          <ds:X509Certificate>MIICszCCAZugAwIBAgIBATANBgkqhkiG9w0BAQsFADAdMRswGQYDVQQDExJzYW1sLW1ldGFkYXRhLXRlc3QwHhcNMjYwMTAxMDAwMDAwWhcNMzYwMTAxMDAwMDAwWjAdMRswGQYDVQQDExJzYW1sLW1ldGFkYXRhLXRlc3QwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDYFtON1NuNMhHxEJxa49gEpzkgIqpbSbJZ4R6MhZIuG8XkN/wwcRHElJ1Y7wCloDsYMAjZ1/1A6pPqb+/D12oECPVugshUgdtDFduQYrOv7G6aaylMWRFiHyd/y2hj+N4KUqDrPymRUo3ZZJOwpQxk4DLjjhTe9Qp0dWnpHuLIbuSZ3VvfJ9FEOsZqmcgDhq7/0OYkKw9HpA04UaCOzgGxYPxC7NMpSrYT4kq7PIWFSkxd4sOspqlnAwoR5ZVlYckPlZ+4HNihXkOTDbMbT3XIuvjJb/wpFiNyRiY5XnkIMfPYJv8lU3Vq2JsQMZdp+DHUwzHazktfyYjoMIBpBaB3AgMBAAEwDQYJKoZIhvcNAQELBQADggEBAFSr8FgQ56kPSa9e1XiN0pD6KNxG4zE+MD22L2gPtcP5ckZG7g1Yaofzil8Djbh1gbj/VnrfK2WEYehNbeDglQ4K5TlmFxnpsNl4GiIJk9CVFVTaubSSUfOKuJplnXdpBcqdYhCFMXTUBaNyRyHtKxCFperhpyqBXpwEjatMfvhA/QTbdWI5q3V/FwYO1Ixb24C1GWt0146/f3Hpnt8GAlX1qDW7gyZVODTTgKf9X7HIoLIdUtqeNJJTLGhk3AWHSO2hikc3SpzPHwUjR28omLyKHtUD03fHThT5R5nsCgzD3JyT5MqWajkKcC6fB3eVCRtxhQh5IkjFsHuSs/e0wGM=</ds:X509Certificate>
        </ds:X509Data>
      </ds:KeyInfo>
    </md:KeyDescriptor>
    <md:ArtifactResolutionService Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP" Location="https://idp.example.com/auth/realms/saml/protocol/saml/resolve" index="0"></md:ArtifactResolutionService>
    <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://idp.example.com/auth/realms/saml/protocol/saml"></md:SingleLogoutService>
    <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://idp.example.com/auth/realms/saml/protocol/saml"></md:SingleLogoutService>
    <md:SingleLogoutService Binding="urn:oasis:names:tc:SAML:2.0:bindings:SOAP" Location="https://idp.example.com/auth/realms/saml/protocol/saml"></md:SingleLogoutService>
    <md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:persistent</md:NameIDFormat>
    <md:NameIDFormat>urn:oasis:names:tc:SAML:2.0:nameid-format:transient</md:NameIDFormat>
    <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
    <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://idp.example.com/auth/realms/saml/protocol/saml"></md:SingleSignOnService>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://idp.example.com/auth/realms/saml/protocol/saml"></md:SingleSignOnService>
    <md:SingleSignOnService Location="https://idp.example.com/auth/realms/saml/protocol/saml"></md:SingleSignOnService>
    <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Artifact" Location="https://idp.example.com/auth/realms/saml/protocol/saml"></md:SingleSignOnService>
  </md:IDPSSODescriptor>
</md:EntityDescriptor>`

func init() {
	samlMetadataUrl = "http://" + KeycloakDomain + ":8090/auth/realms/master/protocol/saml/descriptor"
}

func TestSAMLIdp_CreateAndDelete(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-create")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	idpRequest := makeSamlIdpRequest(alias, entityId, metadataUrl, mappers)
	resp, httpResp, err := OpenapiClient.IdpsAPI.CreateIdp(ctx).IdpCreationRequest(idpRequest).Execute()
	defer deleteIDP(t, ctx, alias)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, httpResp.StatusCode)
	require.Equal(t, alias, resp.Alias)
}

func TestSAMLIdp_Get(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-get")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	defer deleteIDP(t, ctx, alias)
	resp, err := OpenapiClient.IdpsAPI.GetIdp(ctx, alias).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestSAMLIdp_Update(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-update")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	defer deleteIDP(t, ctx, alias)
	newEntityId := entityId + "-updated"
	updateRequest := makeSamlIdpRequest(alias, newEntityId, metadataUrl, mappers)

	_, err := OpenapiClient.IdpsAPI.UpdateIdp(ctx, alias).IdpCreationRequest(updateRequest).Execute()
	require.NoError(t, err)
}

func TestSAMLIdp_List(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-list")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	defer deleteIDP(t, ctx, alias)

	listResp, httpResp, err := OpenapiClient.IdpsAPI.GetIdps(ctx).Execute()
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, httpResp.StatusCode)
	found := false
	for _, idp := range listResp {
		if idp.Alias == alias {
			found = true
			break
		}
	}
	require.True(t, found, "Created SAML IdP should be in the list")
}

func TestSAMLIdp_Get_NotFound(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-get-notfound")
	alias := "non-existent-saml-idp-alias"
	httpResp, err := OpenapiClient.IdpsAPI.GetIdp(ctx, alias).Execute()
	OpenapiClient.IdpsAPI.GetIdp(ctx, alias)
	requireApiError(t, err, httpResp, http.StatusNotFound)
}

func TestSAMLIdp_Create_Conflict(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-conflict")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	defer deleteIDP(t, ctx, alias)

	idpRequest := makeSamlIdpRequest(alias, entityId, metadataUrl, mappers)
	_, httpResp, err := OpenapiClient.IdpsAPI.CreateIdp(ctx).IdpCreationRequest(idpRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusConflict)
}

func TestSAMLIdp_Create_InvalidBase64MetadataFile(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-create-invalid-base64")
	alias := GenerateRandomString("saml-idp-invalid-base64")
	idpRequest := makeSamlIdpRawFileRequest(alias, samlEntityId, "metadata.xml", "not!valid!base64", GetDefaultIDPMappers())

	_, httpResp, err := OpenapiClient.IdpsAPI.CreateIdp(ctx).IdpCreationRequest(idpRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusBadRequest)
}

func TestSAMLIdp_Create_UnparseableMetadataFile(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-create-unparseable")
	alias := GenerateRandomString("saml-idp-unparseable")
	idpRequest := makeSamlIdpFileRequest(alias, samlEntityId, "metadata.xml", "this is not saml metadata", GetDefaultIDPMappers())

	_, httpResp, err := OpenapiClient.IdpsAPI.CreateIdp(ctx).IdpCreationRequest(idpRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusBadRequest)
}

func TestSAMLIdp_Create_MetadataRejectedByKeycloak(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-create-keycloak-rejects")
	alias := GenerateRandomString("saml-idp-keycloak-rejects")
	idpRequest := makeSamlIdpFileRequest(alias, samlEntityId, "metadata.xml", samlDescriptorMissingSsoBinding, GetDefaultIDPMappers())

	_, httpResp, err := OpenapiClient.IdpsAPI.CreateIdp(ctx).IdpCreationRequest(idpRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusBadRequest)
}

func TestSAMLIdp_Create_UnresolvableMetadataUrl(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-create-unresolvable-url")
	alias := GenerateRandomString("saml-idp-unresolvable-url")
	idpRequest := makeSamlIdpRequest(alias, samlEntityId, "http://no-such-host.invalid/descriptor", GetDefaultIDPMappers())

	_, httpResp, err := OpenapiClient.IdpsAPI.CreateIdp(ctx).IdpCreationRequest(idpRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusBadRequest)
}

func TestSAMLIdp_Update_NotFound(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-update-notfound")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	idpRequest := makeSamlIdpRequest(alias, entityId, metadataUrl, mappers)
	httpResp, err := OpenapiClient.IdpsAPI.UpdateIdp(ctx, alias).IdpCreationRequest(idpRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusNotFound)
}

func TestSAMLIdp_Update_InvalidBase64MetadataFile(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-update-invalid-base64")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	defer deleteIDP(t, ctx, alias)
	invalidRequest := makeSamlIdpRawFileRequest(alias, entityId, "metadata.xml", "not!valid!base64", mappers)

	httpResp, err := OpenapiClient.IdpsAPI.UpdateIdp(ctx, alias).IdpCreationRequest(invalidRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusBadRequest)
}

func TestSAMLIdp_Update_UnparseableMetadataFile(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-update-unparseable")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	defer deleteIDP(t, ctx, alias)
	invalidRequest := makeSamlIdpFileRequest(alias, entityId, "metadata.xml", "this is not saml metadata", mappers)

	httpResp, err := OpenapiClient.IdpsAPI.UpdateIdp(ctx, alias).IdpCreationRequest(invalidRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusBadRequest)
}

func TestSAMLIdp_Update_MetadataRejectedByKeycloak(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-update-keycloak-rejects")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	defer deleteIDP(t, ctx, alias)
	invalidRequest := makeSamlIdpFileRequest(alias, entityId, "metadata.xml", samlDescriptorMissingSsoBinding, mappers)

	httpResp, err := OpenapiClient.IdpsAPI.UpdateIdp(ctx, alias).IdpCreationRequest(invalidRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusBadRequest)
}

func TestSAMLIdp_Update_Invalid(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-invalid")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	defer deleteIDP(t, ctx, alias)
	invalidRequest := makeSamlIdpRequest(alias, entityId, "", mappers) // missing entityId

	httpResp, err := OpenapiClient.IdpsAPI.UpdateIdp(ctx, alias).IdpCreationRequest(invalidRequest).Execute()
	requireApiError(t, err, httpResp, http.StatusBadRequest)
}

func TestSAMLIdp_Delete_NotFound(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-delete-notfound")
	alias := "non-existent-saml-idp-alias"
	httpResp, err := OpenapiClient.IdpsAPI.DeleteIdp(ctx, alias).Execute()
	requireApiError(t, err, httpResp, http.StatusNotFound)
}

func TestSAMLIdp_Create_MissingFields(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-missingfields")
	_, httpResp, err := OpenapiClient.IdpsAPI.CreateIdp(ctx).IdpCreationRequest(openapi.IdpCreationRequest{
		Name: StrPtr("missing-metadataxmltype"),
		Type: openapi.IDPCREATIONREQUESTTYPEENUM_SAML,
		SamlData: *openapi.NewNullableSamlCreationData(&openapi.SamlCreationData{
			MetadataXmlType: nil,
			MetadataXmlUrl:  StrPtr(samlMetadataUrl),
			EntityId:        StrPtr(samlEntityId),
		}),
		Mappers: GetDefaultIDPMappers(),
	}).Execute()
	requireApiError(t, err, httpResp, http.StatusBadRequest)
}

func TestSAMLIdp_List_Empty(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-listempty")
	listResp, _, err := OpenapiClient.IdpsAPI.GetIdps(ctx).Execute()
	require.NoError(t, err)
	require.Empty(t, listResp)
}

func TestSAMLIdp_Get_AfterDelete(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-get-afterdelete")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	deleteIDP(t, ctx, alias)

	_, err := OpenapiClient.IdpsAPI.GetIdp(ctx, alias).Execute()
	require.Error(t, err)
}

func TestSAMLIdp_GetIdpMappers(t *testing.T) {
	t.Parallel()
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-get-idp-mappers")
	alias, entityId, metadataUrl, mappers := newRandomSamlIdpTestDataWithRealm()
	createSamlIdp(t, ctx, alias, entityId, metadataUrl, mappers)
	defer deleteIDP(t, ctx, alias)

	result, _, err := OpenapiClient.IdpsAPI.GetIdpMappers(ctx, alias).Execute()
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, mappers.Gid, result.Gid)
	require.Equal(t, mappers.Uid, result.Uid)
	require.Equal(t, mappers.Groups, result.Groups)
	require.Equal(t, mappers.SupplementaryGroups, result.SupplementaryGroups)
	require.Equal(t, mappers.Email, result.Email)
	require.Equal(t, mappers.FirstName, result.FirstName)
	require.Equal(t, mappers.LastName, result.LastName)
}

func TestSAMLIdp_Create_FileUpload_FromDisk(t *testing.T) {
	ctx, _, _ := TestStepsClient.NewTestContextWithRealm("saml-idp-file-disk")
	alias := GenerateRandomString("saml-idp-file-disk")
	entityId := "urn:test:idp"
	fileName := "test-metadata.xml"
	// Use a hardcoded XML string instead of reading from disk
	xmlContent := `<EntityDescriptor entityID="urn:test:idp" xmlns="urn:oasis:names:tc:SAML:2.0:metadata">
  <IDPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol"/>
</EntityDescriptor>`
	mappers := GetDefaultIDPMappers()
	idpRequest := makeSamlIdpFileRequest(alias, entityId, fileName, xmlContent, mappers)
	defer deleteIDP(t, ctx, alias)

	resp, httpResp, err := OpenapiClient.IdpsAPI.CreateIdp(ctx).IdpCreationRequest(idpRequest).Execute()
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, http.StatusOK, httpResp.StatusCode)
	require.Equal(t, alias, resp.Alias)
}

func makeSamlIdpRequest(alias, entityId, metadataUrl string, mappers *openapi.Mappers) openapi.IdpCreationRequest {
	metadataType := openapi.SAMLCREATIONDATAMETADATAXMLTYPEENUM_URL
	return openapi.IdpCreationRequest{
		Name: StrPtr(alias),
		Type: openapi.IDPCREATIONREQUESTTYPEENUM_SAML,
		SamlData: *openapi.NewNullableSamlCreationData(&openapi.SamlCreationData{
			MetadataXmlType: &metadataType,
			MetadataXmlUrl:  StrPtr(metadataUrl),
			EntityId:        StrPtr(entityId),
		}),
		Mappers: mappers,
	}
}

func makeSamlIdpFileRequest(alias, entityId, fileName, xmlContent string, mappers *openapi.Mappers) openapi.IdpCreationRequest {
	metadataType := openapi.SAMLCREATIONDATAMETADATAXMLTYPEENUM_FILE
	xmlBase64 := base64.StdEncoding.EncodeToString([]byte(xmlContent))
	return openapi.IdpCreationRequest{
		Name: StrPtr(alias),
		Type: openapi.IDPCREATIONREQUESTTYPEENUM_SAML,
		SamlData: *openapi.NewNullableSamlCreationData(&openapi.SamlCreationData{
			MetadataXmlType: &metadataType,
			MetadataXmlFile: StrPtr(xmlBase64),
			FileName:        StrPtr(fileName),
			EntityId:        StrPtr(entityId),
		}),
		Mappers: mappers,
	}
}

// makeSamlIdpRawFileRequest sends metadataXmlFile verbatim, unlike
// makeSamlIdpFileRequest, so the base64 decoding itself can be exercised.
func makeSamlIdpRawFileRequest(alias, entityId, fileName, metadataXmlFile string, mappers *openapi.Mappers) openapi.IdpCreationRequest {
	metadataType := openapi.SAMLCREATIONDATAMETADATAXMLTYPEENUM_FILE
	return openapi.IdpCreationRequest{
		Name: StrPtr(alias),
		Type: openapi.IDPCREATIONREQUESTTYPEENUM_SAML,
		SamlData: *openapi.NewNullableSamlCreationData(&openapi.SamlCreationData{
			MetadataXmlType: &metadataType,
			MetadataXmlFile: StrPtr(metadataXmlFile),
			FileName:        StrPtr(fileName),
			EntityId:        StrPtr(entityId),
		}),
		Mappers: mappers,
	}
}

func newRandomSamlIdpTestDataWithRealm() (string, string, string, *openapi.Mappers) {
	alias := GenerateRandomString("saml")
	entityId := samlEntityId
	metadataUrl := samlMetadataUrl
	mappers := GetDefaultIDPMappers()
	return alias, entityId, metadataUrl, mappers
}

func createSamlIdp(t *testing.T, ctx context.Context, alias, entityId, metadataUrl string, mappers *openapi.Mappers) {
	idpRequest := makeSamlIdpRequest(alias, entityId, metadataUrl, mappers)
	createIDP(t, ctx, idpRequest)
}

func createIDP(t *testing.T, ctx context.Context, idpCreationRequest openapi.IdpCreationRequest) {
	err := TestStepsClient.CreateIDP(ctx, idpCreationRequest)
	require.NoError(t, err)
}

func deleteIDP(t *testing.T, ctx context.Context, alias string) {
	err := TestStepsClient.DeleteIDP(ctx, alias)
	require.NoError(t, err)
}
