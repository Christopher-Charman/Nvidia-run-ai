# Static OpenAPI schemas for asset spec validation

Assets-service validates the free-form `objectSpec` payload of k8s-object-backed
assets (currently: resource claim templates) against a static copy of the
Kubernetes OpenAPI v3 schema for the corresponding resource, checked in to this
folder and loaded via `go:embed` at build time.

Static (not runtime-fetched) so:

- the control plane can validate without a live cluster reachable;
- validation is deterministic across environments;
- no bootstrap ordering problem (assets-service does not depend on any specific
  cluster being registered before it starts serving).

## Files

- `rct_object_spec.schema.json` — the schema bundle for
  `io.k8s.api.resource.v1.ResourceClaimSpec` plus every schema it transitively
  `$ref`s (`DeviceClaim`, `DeviceRequest`, `DeviceSelector`, ...). Loaded by
  `pkg/services/api_rct_service.go` and consumed by `validateRctObjectSpec`.

## How to refresh

When Kubernetes ships a change to `resource.k8s.io/v1` (rare — the API group
went GA at 1.34 and is expected to be stable), regenerate the bundle from a
live cluster running the minimum supported K8s version. Either shape works —
the loader (`extractComponentSchemas` in `api_rct_service.go`) accepts the
full OpenAPI v3 document or just the schemas dictionary:

```bash
# Simplest — the whole document (a few hundred KB):
kubectl get --raw '/openapi/v3/apis/resource.k8s.io/v1' \
  > control-plane/assets-service/pkg/services/schemas/rct_object_spec.schema.json

# Or the schemas dictionary only (a bit smaller):
kubectl get --raw '/openapi/v3/apis/resource.k8s.io/v1' \
  | jq '.components.schemas' \
  > control-plane/assets-service/pkg/services/schemas/rct_object_spec.schema.json
```

Do NOT slice down to only `ResourceClaimSpec` — its `$ref`s
(`DeviceClaim`, `DeviceRequest`, `DeviceSelector`, ...) need to resolve
locally, so the whole `components.schemas` dictionary must stay together.

Commit the file with a note in the commit message identifying the K8s server
version it came from (`kubectl version -o json | jq .serverVersion`).
