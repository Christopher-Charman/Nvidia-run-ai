/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package composer

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Automatic exposed URLs", func() {
	DescribeTable("preserves the URL used by networking and authentication",
		func(domain, workload, suffix string, subdomain bool, expected string) {
			actual, err := ComposeAutoExposedURL(domain, "team", workload, suffix, subdomain)
			Expect(err).NotTo(HaveOccurred())
			Expect(actual.String()).To(Equal(expected))
		},
		Entry("first subdomain URL", "https://example.com", "job", "", true, "https://team-job.example.com/"),
		Entry("subdomain suffix", "https://example.com", "job", "metrics", true, "https://team-job-metrics.example.com/"),
		Entry("path suffix", "https://example.com", "job", "metrics", false, "https://example.com/team/job/metrics"),
		Entry("configured path, query and fragment", "https://example.com/base?key=value#section", "job", "metrics", false,
			"https://example.com/team/job/metrics//base?key=value#section"),
		Entry("subdomain preserves port and path", "https://example.com:8443/base", "job", "", true,
			"https://team-job.example.com:8443/base"),
		Entry("long hostname uses the existing stable hash", "https://example.com", strings.Repeat("x", 60), "metrics", true,
			"https://team-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-9e3eddb17a.example.com/"),
	)

	DescribeTable("rejects an unusable cluster domain",
		func(domain string) {
			actual, err := ComposeAutoExposedURL(domain, "team", "job", "", true)
			Expect(err).To(HaveOccurred())
			Expect(actual).To(BeNil())
		},
		Entry("empty", ""),
		Entry("malformed escape", "https://example.com/%zz"),
	)
})
