#!/usr/bin/env python3
# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
"""Rewrite absolute run-ai-docs.nvidia.com / Redoc-anchor links in tag md files to root-relative Fern paths.
Mapping follows the spec-source link conversion (control-plane specs). usage: relativize-links.py <dir> <product-slug>"""
import re, sys, pathlib
d, product = sys.argv[1], sys.argv[2]
SH = "/run-ai/self-hosted/documentation"
M = [  # (regex on old URL path after domain, new relative path); anchors dropped where headings no longer exist
 (r"(self-hosted/(latest|2\.25)/)infrastructure-setup/authentication/overview/?(#.*)?", SH+"/platform-management/authentication/overview"),
 (r"self-hosted/(latest|2\.25)/infrastructure-setup/procedures/event-history", SH+"/installation-and-configuration/operate/observability/event-history"),
 (r"self-hosted/latest/platform-management/policies/?(#.*)?", SH+"/platform-management/policies/policies-and-rules"),
 (r"self-hosted/latest/platform-management/aiinitiatives/(resources/nodes|resources/node-pools|organization/projects|organization/departments)/?(#.*)?", None),
 (r"self-hosted/latest/workloads-in-nvidia-run-ai/using-training/distributed-training/distributed-training-models", SH+"/workloads-in-nvidia-run-ai/using-training/distributed-training-models"),
 (r"self-hosted/latest/(workloads-in-nvidia-run-ai/.+)", None),
 (r"saas/(platform-management/.+)", None),
]
def conv(url):
    p = re.sub(r"^https://run-ai-docs\.nvidia\.com/", "", url)
    for rx, new in M:
        m = re.fullmatch(rx, p)
        if not m: continue
        if new: return new
        if "aiinitiatives" in rx: return f"{SH}/platform-management/manage-ai-initiatives/{m.group(1)}"
        return f"{SH}/{m.group(1)}"
    return None
def tag(m):  # Redoc "#tag/Foo-Bar" -> Fern tag page in the same product
    slug = m.group(1).lower()
    return f"](/run-ai/{product}/reference/rest-api/endpoints/authentication-and-authorization/{slug})"
bad = []
for f in pathlib.Path(d).glob("*.md"):
    s = f.read_text()
    def link(m):
        n = conv(m.group(1)); 
        if n is None: bad.append(m.group(1)); return m.group(0)
        return f"]({n})"
    s = re.sub(r"\]\((https://run-ai-docs\.nvidia\.com/[^)]*)\)", link, s)
    s = re.sub(r"\]\((?:/2\.25/)?#tag/([A-Za-z-]+)\)", tag, s)
    f.write_text(s)
print("unmapped:", bad)
