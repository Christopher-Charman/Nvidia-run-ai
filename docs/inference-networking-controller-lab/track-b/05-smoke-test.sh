#!/usr/bin/env bash
# Track B smoke test.
#   ./05-smoke-test.sh b1     B.1  request through Dynamo's own Frontend (no Gateway)
#   ./05-smoke-test.sh b2     B.2  request through agentgateway → our HTTPRoute → operator's pool → Dynamo EPP → mocker
set -uo pipefail
NS="${NS:-dynamo-lab}"
DGD="${DGD:-mocker-agg}"
MODEL="${MODEL:-Qwen/Qwen3-0.6B}"
AGW_NS="${AGW_NS:-agentgateway-system}"
MODE="${1:-b1}"
PASS=0; FAIL=0
ok()   { printf '  \033[1;32m✓ %s\033[0m\n' "$*"; PASS=$((PASS+1)); }
bad()  { printf '  \033[1;31m✗ %s\033[0m\n' "$*"; FAIL=$((FAIL+1)); }
step() { printf '\n\033[1m[%s] %s\033[0m\n' "$1" "$2"; }
pf_pids=(); cleanup() { for p in "${pf_pids[@]:-}"; do [ -n "$p" ] && kill "$p" 2>/dev/null; done; }; trap cleanup EXIT
body='{"model":"'"$MODEL"'","messages":[{"role":"user","content":"Explain KV-aware routing in one sentence."}],"max_tokens":32}'

step 0 "DGD and pods"
kubectl get dynamographdeployment "$DGD" -n "$NS" -o jsonpath='{.status.state}{"\n"}' 2>/dev/null
kubectl get pods -n "$NS" -o wide

if [ "$MODE" = "b1" ]; then
  step 1 "Frontend Service exists"
  kubectl get svc -n "$NS" "${DGD}-frontend" && ok "svc ${DGD}-frontend" || bad "no ${DGD}-frontend Service"
  step 2 "Request through Dynamo Frontend (documented path)"
  kubectl port-forward -n "$NS" "svc/${DGD}-frontend" 18000:8000 >/dev/null 2>&1 & pf_pids+=($!); sleep 3
  r=$(curl -s -m 60 localhost:18000/v1/chat/completions -H 'Content-Type: application/json' -d "$body")
  echo "$r" | head -c 400; echo
  echo "$r" | grep -q '"choices"' && ok "mocker answered through the Frontend" || bad "no choices in response"
  step 3 "Models list"
  curl -s -m 20 localhost:18000/v1/models | head -c 300; echo
fi

if [ "$MODE" = "b2" ]; then
  step 1 "Operator-generated objects: ${DGD}-pool and ${DGD}-epp"
  kubectl get inferencepool "${DGD}-pool" -n "$NS" -o jsonpath='selector={.spec.selector} epp={.spec.endpointPickerRef.name}:{.spec.endpointPickerRef.port.number}{"\n"}' 2>/dev/null \
    && ok "InferencePool ${DGD}-pool exists (✅ Dynamo created it)" || bad "no InferencePool ${DGD}-pool"
  kubectl get svc,deploy -n "$NS" | grep -E "${DGD}-epp" && ok "EPP Service/Deployment exist" || bad "no ${DGD}-epp"
  step 2 "Our HTTPRoute (🆕) accepted and resolved"
  conds=$(kubectl get httproute "$DGD" -n "$NS" -o jsonpath='{range .status.parents[*].conditions[*]}{.type}={.status} {end}')
  echo "  $conds"
  echo "$conds" | grep -q 'Accepted=True' && ok "Accepted=True" || bad "not Accepted"
  echo "$conds" | grep -q 'ResolvedRefs=True' && ok "ResolvedRefs=True" || bad "refs not resolved"
  step 3 "Request through agentgateway → our route → Dynamo pool → Dynamo EPP → mocker"
  GWSVC=$(kubectl get svc -n "$AGW_NS" -l gateway.networking.k8s.io/gateway-name=inference-gateway -o jsonpath='{.items[0].metadata.name}')
  echo "  gateway svc: ${GWSVC:-<none>}"
  kubectl port-forward -n "$AGW_NS" "svc/$GWSVC" 18080:80 >/dev/null 2>&1 & pf_pids+=($!); sleep 3
  r=$(curl -s -m 120 localhost:18080/v1/chat/completions -H "X-Gateway-Model-Name: $MODEL" -H 'Content-Type: application/json' -d "$body")
  echo "$r" | head -c 400; echo
  echo "$r" | grep -q '"choices"' && ok "mocker answered through the Gateway + Dynamo EPP  ← ExistingPool mode works on CPU" || bad "no choices; see EPP logs below"
  step 4 "Dynamo EPP logs (last 15 lines)"
  kubectl logs -n "$NS" "deploy/${DGD}-epp" --tail=15 2>/dev/null | cut -c1-220
fi

if [ "$MODE" = "b2prime" ]; then
  step 1 "Our objects (🆕): pool mocker-frontend-pool selects Dynamo Frontend pods, EPP mocker-frontend-epp"
  kubectl get inferencepool mocker-frontend-pool -n "$NS" -o jsonpath='selector={.spec.selector.matchLabels} epp={.spec.endpointPickerRef.name}{"\n"}' && ok "pool exists" || bad "no pool"
  [ "$(kubectl get deploy mocker-frontend-epp -n "$NS" -o jsonpath='{.status.availableReplicas}')" = "1" ] && ok "EPP available" || bad "EPP not available"
  step 2 "Our HTTPRoute mocker-frontend accepted and resolved"
  conds=$(kubectl get httproute mocker-frontend -n "$NS" -o jsonpath='{range .status.parents[*].conditions[*]}{.type}={.status} {end}'); echo "  $conds"
  echo "$conds" | grep -q 'Accepted=True' && ok "Accepted=True" || bad "not Accepted"
  echo "$conds" | grep -q 'ResolvedRefs=True' && ok "ResolvedRefs=True" || bad "refs not resolved"
  step 3 "Request: agentgateway → our route → our pool → our GIE EPP → Dynamo Frontend pod → mocker"
  GWSVC=$(kubectl get svc -n "$AGW_NS" -l gateway.networking.k8s.io/gateway-name=inference-gateway -o jsonpath='{.items[0].metadata.name}')
  kubectl port-forward -n "$AGW_NS" "svc/$GWSVC" 18080:80 >/dev/null 2>&1 & pf_pids+=($!); sleep 3
  got=0; for i in 1 2 3; do
    r=$(curl -s -m 90 localhost:18080/v1/chat/completions -H 'Host: mocker.test' -H 'Content-Type: application/json' -d "$body")
    echo "$r" | grep -q '"choices"' && { got=1; break; }; sleep 4
  done
  echo "$r" | head -c 300; echo
  [ "$got" = 1 ] && ok "completion via our EPP over Dynamo-owned pods  ← Dynamo as a pod owner works" || bad "no completion; check EPP logs"
fi

printf '\n\033[1mResult: %d passed, %d failed\033[0m\n' "$PASS" "$FAIL"
[ "$FAIL" -eq 0 ]
