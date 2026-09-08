#!/usr/bin/env bash
# Track B step 0 · what did the env ship? Read-only.
set -uo pipefail
NS_OP="${NS_OP:-dynamo-system}"

printf '\n\033[1m== Dynamo CRDs\033[0m\n'
kubectl get crd | grep -E 'nvidia.com' || echo "  none: Dynamo operator NOT installed"

printf '\n\033[1m== Operator + platform pods in %s\033[0m\n' "$NS_OP"
kubectl get pods -n "$NS_OP" -o wide 2>/dev/null || echo "  namespace missing"

printf '\n\033[1m== Operator image tag (this decides DYNAMO_VERSION for the manifests)\033[0m\n'
kubectl get deploy -n "$NS_OP" -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.spec.template.spec.containers[*].image}{"\n"}{end}' 2>/dev/null

printf '\n\033[1m== Event plane (NATS) and etcd\033[0m\n'
kubectl get svc -n "$NS_OP" 2>/dev/null | grep -iE 'nats|etcd' || echo "  no nats/etcd Service found"

printf '\n\033[1m== Gateway API already present? (expect: no, Dynamo pins its own versions)\033[0m\n'
kubectl api-resources 2>/dev/null | grep -iE 'gateway.networking|inference.networking' || echo "  none, clean slate"

printf '\n\033[1m== Fake GPU node\033[0m\n'
kubectl get nodes -o json | python3 -c "import sys,json; [print(' ',n['metadata']['name'], 'nvidia.com/gpu =', n['status']['allocatable'].get('nvidia.com/gpu','-')) for n in json.load(sys.stdin)['items']]"

printf '\n\033[1m== Existing DGDs (expect none)\033[0m\n'
kubectl get dynamographdeployments -A 2>/dev/null || true
