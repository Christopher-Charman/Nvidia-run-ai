/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Workload Client with slim list", func() {
	var (
		server        *httptest.Server
		client        *BackendWorkloadServiceClient
		capturedQuery url.Values
	)

	BeforeEach(func() {
		// Return a slim response: only the identity fields
		server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			capturedQuery = r.URL.Query()

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"workloads": [
					{
						"id": "slim-1",
						"name": "slim-workload-1",
						"type": "Training",
						"projectId": "1",
						"projectName": "test-project",
						"clusterId": "cluster-1",
						"externalWorkloadName": "ew-training-slim-1"
					}
				]
			}`))
		}))

		var err error
		client, err = NewWorkloadServiceClient(server.URL, "test-token")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		server.Close()
	})

	It("should send verbosity=slim and parse the slim workloads", func() {
		response, err := client.WithSlimList().List(context.Background(), 10, 0, []string{"name==slim-workload-1"}, false)

		Expect(err).ToNot(HaveOccurred())
		Expect(capturedQuery.Get("verbosity")).To(Equal("slim"))
		Expect(response.Workloads).To(HaveLen(1))
		Expect(response.Workloads[0].Id).To(Equal("slim-1"))
		Expect(response.Workloads[0].Name).To(Equal("slim-workload-1"))
		Expect(response.Workloads[0].Type).To(Equal("Training"))
		Expect(response.Workloads[0].ProjectName).To(Equal("test-project"))
		Expect(response.Workloads[0].GetExternalWorkloadName()).To(Equal("ew-training-slim-1"))
	})

	It("should not send verbosity by default", func() {
		_, err := client.List(context.Background(), 10, 0, []string{}, false)

		Expect(err).ToNot(HaveOccurred())
		_, hasVerbosity := capturedQuery["verbosity"]
		Expect(hasVerbosity).To(BeFalse())
	})
})
