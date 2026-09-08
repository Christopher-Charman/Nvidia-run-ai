#!/usr/bin/env bash
# Track B step 2 · B.2 prerequisites, ✅ upstream. Adapted from Dynamo's own
# deploy/inference-gateway/scripts/install_gaie_crd_agentgateway.sh (ai-dynamo/dynamo main), same pins:
#   Gateway API v1.5.1 · GIE v1.2.1 · agentgateway v1.0.0
# These are four minor versions behind Track A on purpose: Dynamo's operator and EPP were tested against them.
set -euo pipefail
NAMESPACE="${NAMESPACE:-dynamo-lab}"
AGW_NAMESPACE="${AGW_NAMESPACE:-agentgateway-system}"
GATEWAY_API_VERSION="${GATEWAY_API_VERSION:-v1.5.1}"
IGW_RELEASE="${IGW_RELEASE:-v1.2.1}"
AGW_VERSION="${AGW_VERSION:-v1.0.0}"
step() { printf '\n\033[1;32m== %s\033[0m\n' "$*"; }

step "0/5 workload namespace $NAMESPACE"
kubectl create namespace "$NAMESPACE" --dry-run=client -o yaml | kubectl apply -f -

step "1/5 Gateway API $GATEWAY_API_VERSION standard CRDs"
kubectl apply --server-side --force-conflicts -f \
  "https://github.com/kubernetes-sigs/gateway-api/releases/download/${GATEWAY_API_VERSION}/standard-install.yaml"

step "2/5 GIE $IGW_RELEASE CRDs (InferencePool group inference.networking.k8s.io + x-k8s.io extras)"
kubectl apply -f \
  "https://github.com/kubernetes-sigs/gateway-api-inference-extension/releases/download/${IGW_RELEASE}/manifests.yaml"
kubectl get crd | grep -E 'inference\.networking'

step "3/5 agentgateway $AGW_VERSION CRDs + controller, inference extension on"
helm upgrade -i --create-namespace --namespace "$AGW_NAMESPACE" --version "$AGW_VERSION" \
  agentgateway-crds oci://cr.agentgateway.dev/charts/agentgateway-crds
helm upgrade -i --namespace "$AGW_NAMESPACE" --version "$AGW_VERSION" agentgateway \
  oci://cr.agentgateway.dev/charts/agentgateway \
  --set inferenceExtension.enabled=true --wait --timeout 5m

step "4/5 AgentgatewayParameters (disable istio sidecar injection on the proxy, no-op without istio)"
kubectl apply --server-side -n "$AGW_NAMESPACE" -f - <<'EOF'
apiVersion: agentgateway.dev/v1alpha1
kind: AgentgatewayParameters
metadata:
  name: inference-gateway-params
spec:
  deployment:
    spec:
      template:
        metadata:
          annotations:
            sidecar.istio.io/inject: "false"
EOF

step "5/5 Gateway inference-gateway in $AGW_NAMESPACE, routes allowed from any namespace"
kubectl apply -n "$AGW_NAMESPACE" -f - <<'EOF'
apiVersion: gateway.networking.k8s.io/v1
kind: Gateway
metadata:
  name: inference-gateway
spec:
  gatewayClassName: agentgateway
  infrastructure:
    parametersRef:
      group: agentgateway.dev
      kind: AgentgatewayParameters
      name: inference-gateway-params
  listeners:
    - name: http
      port: 80
      protocol: HTTP
      allowedRoutes:
        namespaces:
          from: All
EOF
kubectl wait gateway/inference-gateway -n "$AGW_NAMESPACE" --for=condition=Programmed --timeout=180s
kubectl get gateway -n "$AGW_NAMESPACE"
echo "OK: Dynamo-pinned Gateway stack installed. Kind has no LoadBalancer: use port-forward."
