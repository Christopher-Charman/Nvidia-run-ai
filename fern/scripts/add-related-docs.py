#!/usr/bin/env python3
# SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
# SPDX-License-Identifier: Apache-2.0
"""Append a 'Related documentation' section to tag summary md files.
Links that used to live in OpenAPI descriptions are maintained here, in md only. Idempotent."""
import pathlib
SH = "/run-ai/self-hosted/documentation"
def api(product, tag): return f"/run-ai/{product}/reference/rest-api/endpoints/{tag}"
UID = ("User identity in containers", SH+"/installation-and-configuration/configuration/security/user-identity-in-containers")
INF = ("Inference requirements", SH+"/installation-and-configuration/system-requirements/networking-setup#standard-inference-wildcard-certificate")
PROJ = ("Projects (GPU over quota)", SH+"/platform-management/manage-ai-initiatives/organization/projects")
K = lambda t, p: (t, "https://kubernetes.io/docs/"+p)
REL = {
 "access-rules": [("Create access rules", SH+"/platform-management/authentication/access-rules")],
 "roles": [("Roles", SH+"/platform-management/authentication/roles")],
 "tokens": [("Request an API token", "/run-ai/self-hosted/reference/rest-api/getting-started/how-to-authenticate-to-the-api#request-an-api-token")],
 "projects": [PROJ], "departments": [PROJ],
 "nodes": [K("Kubernetes topology manager policies", "tasks/administer-cluster/topology-manager/#topology-manager-policies")],
 "node-pools": [("Node pools", SH+"/platform-management/manage-ai-initiatives/resources/node-pools"), ("Projects", SH+"/platform-management/manage-ai-initiatives/organization/projects")],
 "environment": [K("Images", "concepts/containers/images"), UID],
 "nfs": [K("NFS volumes", "concepts/storage/volumes#nfs")],
 "pvc": [K("Storage classes", "concepts/storage/storage-classes")],
 "s3": [("AWS credentials access keys", "https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html")],
 "nvidia-nim": [INF], "distributed-inferences": [INF],
 "policy": [("jq manual", "https://jqlang.org/manual/")],
 "workloads": [("Get workload endpoints", "@get-workload-endpoints")],
}
for f in pathlib.Path("docs/rest-api").glob("*/tags/*.md"):
    product = f.parent.parent.name
    items = REL.get(f.stem)
    if not items: continue
    s = f.read_text()
    if "Related documentation" in s: continue
    rows = []
    for t, u in items:
        if u.startswith("@"): continue
        if u.startswith("/run-ai/self-hosted") and product == "multi-tenant" and "reference" not in u: continue
        rows.append(f"- [{t}]({u})")
    if rows: f.write_text(s.rstrip("\n") + "\n\n## Related documentation\n\n" + "\n".join(rows) + "\n")
