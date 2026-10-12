# 001 · Setup: real GPUs, real model, EPP from the adaptive branch

## TL;DR (the whole thing in 3 lines)

- One Kubernetes cluster with **4 GPUs** (one node is fine), **4 vLLM pods** serving one real model, an **Inference Gateway** (kgateway) in front, and the **EPP built from Nir's `adaptive-with-pressure` branch** behind it.
- Prometheus scrapes the EPP and the vLLM pods. That is where the numbers come from.
- Everything below copies manifests that already exist in the GAIE repo; only the GPU count, the model, and the EPP image are changed.

Status: setup guide, written 2026-10-12 from the pinned branch revision `6c0f8d8`. Not yet executed end to end on a cluster. Versions marked **check** must be confirmed at run time.

## What you end up with

```text
 ┌──────────────────────────────┐
 │ 📈 Load generator (doc 002)  │   GuideLLM pod inside the cluster
 └──────────────┬───────────────┘
                │ HTTP POST /v1/completions
 ╔══════════════▼═══════════════════════════════════════════════════════════╗
 ║ Kubernetes cluster · namespace epp-lab                                    ║
 ║                                                                           ║
 ║  ┌──────────────────┐  ext-proc gRPC   ┌───────────────────────────────┐  ║
 ║  │ ✅ Gateway        │────────────────►│ 🆕 EPP                         │  ║
 ║  │ kgateway          │◄────────────────│  image: adaptive-with-pressure │  ║
 ║  │ + HTTPRoute       │  chosen pod IP  │  env: ENABLE_EXPERIMENTAL_...  │  ║
 ║  └────────┬─────────┘                  │  plugins config (profiles)     │  ║
 ║           │ forward request            └───────────────┬───────────────┘  ║
 ║           ▼                                            │ scrape /metrics   ║
 ║  ┌─────────────────────────────────────────────────────▼────────────────┐ ║
 ║  │ ✅ InferencePool "vllm-model" · 4 × vLLM pods · 1 GPU each            │ ║
 ║  │   real model, --enable-prefix-caching, --max-num-seqs 256            │ ║
 ║  │   exposes vllm:num_requests_waiting / running / gpu_cache_usage_perc │ ║
 ║  └──────────────────────────────────────────────────────────────────────┘ ║
 ║  ┌───────────────────────┐                                                ║
 ║  │ 📊 Prometheus         │  scrapes EPP (ServiceMonitor) + vLLM (PodMonitor)║
 ║  └───────────────────────┘                                                ║
 ╚═══════════════════════════════════════════════════════════════════════════╝
```

✅ = exists in GAIE / the cluster today · 🆕 = the only custom piece, the EPP image from the adaptive branch.

## Prerequisites

| Need | Minimum | Notes |
|---|---|---|
| Kubernetes cluster | 1.30+, one node with **4 GPUs** | 8 GPUs lets you reproduce Nir's 8-pod setup later. GPU Operator or device plugin installed, `nvidia.com/gpu` allocatable. |
| GPU type | A100 40 GB / H100 / L40S | Fits an 8B model in bf16 with room for KV cache. L4 (24 GB) works with a 7–8B model in fp8 or a 3B model. |
| Model | `Qwen/Qwen3-8B` (no gate) or `meta-llama/Llama-3.1-8B-Instruct` (HF token) | Nir's branch manifest uses `Qwen/Qwen3-32B` with TP=2 on 16 GPUs; that is the "later" step, not the minimum. |
| Container registry | reachable from the cluster | For the EPP image you build. |
| Local tools | `kubectl`, `helm`, `docker` (buildx), `git` | Present on this machine. |
| Cluster access | `kc add --session` | Per the personal rules. |

## Sizing

| GPUs available | Pods | Model | TP | Why |
|---|---|---|---|---|
| **4** (minimum) | 4 | 8B bf16 | 1 | Enough pods for a hotspot to be visible (1 hot, 3 cold) |
| 8 | 8 | 8B bf16 | 1 | Nir's reported benchmark shape |
| 16 | 8 | 32B | 2 | Nir's current branch manifest |

Fewer than 4 pods makes "imbalance" almost meaningless: with 2 pods the CV of two numbers is a coin flip.

## Steps

### 0. Namespace, variables, HF token

```bash
export NS=epp-lab
export MODEL=Qwen/Qwen3-8B                 # or meta-llama/Llama-3.1-8B-Instruct
export POOL=vllm-model                     # InferencePool name and pod label
export REGISTRY=<your-registry>/epp-lab    # where the EPP image goes
export EPP_TAG=adaptive-6c0f8d8

kubectl create namespace $NS
kubectl config set-context --current --namespace=$NS
# only if the model is gated:
kubectl create secret generic hf-token --from-literal=token=<your-hf-token>
```

### 1. Gateway API + Inference Extension CRDs

```bash
export IGW_LATEST_RELEASE=$(curl -s https://api.github.com/repos/kubernetes-sigs/gateway-api-inference-extension/releases \
  | jq -r '[.[] | select(.prerelease == false)][0].tag_name')      # check: v1.x
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api-inference-extension/releases/download/${IGW_LATEST_RELEASE}/manifests.yaml
```

This is step "Install the Inference Extension CRDs" of the GAIE getting-started guide.

### 2. Gateway provider: kgateway

```bash
KGTW_VERSION=v2.2.0-main      # check: the version named in site-src/guides/index.md at the pinned revision
helm upgrade -i --create-namespace --namespace kgateway-system --version $KGTW_VERSION \
  kgateway-crds oci://cr.kgateway.dev/kgateway-dev/charts/kgateway-crds
helm upgrade -i --namespace kgateway-system --version $KGTW_VERSION \
  kgateway oci://cr.kgateway.dev/kgateway-dev/charts/kgateway --set inferenceExtension.enabled=true
kubectl apply -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api-inference-extension/refs/tags/${IGW_LATEST_RELEASE}/config/manifests/gateway/kgateway/gateway.yaml
```

Istio or NGINX Gateway Fabric work too; the guide has tabs for each. kgateway is the lightest.

### 3. vLLM pods (real model)

Adapted from `config/manifests/vllm/gpu-deployment.yaml` at the pinned revision: 4 replicas, 1 GPU each, no LoRA, prefix caching on, no node pin.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: vllm-model
spec:
  replicas: 4
  selector:
    matchLabels:
      app: vllm-model
  template:
    metadata:
      labels:
        app: vllm-model
    spec:
      containers:
      - name: vllm
        image: vllm/vllm-openai:v0.13.0        # check: version used in the branch manifest
        command: ["python3", "-m", "vllm.entrypoints.openai.api_server"]
        args:
        - --model
        - Qwen/Qwen3-8B                        # $MODEL
        - --port
        - "8000"
        - --max-num-seqs
        - "256"
        - --enable-prefix-caching
        - --gpu-memory-utilization
        - "0.90"
        env:
        - name: VLLM_USE_V1
          value: "1"
        # only if the model is gated:
        # - name: HUGGING_FACE_HUB_TOKEN
        #   valueFrom: { secretKeyRef: { name: hf-token, key: token } }
        ports:
        - { containerPort: 8000, name: http }
        readinessProbe:
          httpGet: { path: /health, port: http }
          periodSeconds: 2
        startupProbe:
          httpGet: { path: /health, port: http }
          failureThreshold: 600
          periodSeconds: 2
        resources:
          limits: { nvidia.com/gpu: 1 }
          requests: { nvidia.com/gpu: 1 }
        volumeMounts:
        - { mountPath: /dev/shm, name: shm }
        - { mountPath: /root/.cache/huggingface, name: hf-cache }
      volumes:
      - name: shm
        emptyDir: { medium: Memory }
      - name: hf-cache
        emptyDir: {}                            # use a PVC if you will redeploy often
```

```bash
kubectl apply -f vllm-model.yaml
kubectl rollout status deploy/vllm-model --timeout=20m     # first pull + model download
```

Why these flags matter for the experiment:

| Flag | Effect on the signals |
|---|---|
| `--enable-prefix-caching` | Without it the prefix scorer never sees a hit and no hotspot forms. |
| `--max-num-seqs 256` | Sets how many requests run before they queue. Lower = queues form sooner = pressure visible at lower load. |
| 1 GPU per pod, 4 pods | 4 independent queues, so imbalance across pods is a real quantity. |

### 4. Build the EPP image from the adaptive branch

```bash
git clone https://github.com/nirrozenbaum/gateway-api-inference-extension.git epp-adaptive
cd epp-adaptive
git checkout 6c0f8d8496b56d68c30369520f3c4477adf8cd42      # adaptive-with-pressure, pinned
IMAGE_REGISTRY=$REGISTRY IMAGE_TAG=$REGISTRY/epp:$EPP_TAG make image-push
```

`make image-push` is the Makefile target `image-build` with `--push`. If the cluster can pull from your local Docker daemon (kind, or a single-node cluster you control), `make image-load` instead.

### 5. Install the InferencePool + EPP with the chart from the same checkout

Using the local chart guarantees the chart and the image come from the same revision.

```bash
helm install $POOL ./config/charts/inferencepool \
  --set inferencePool.modelServers.matchLabels.app=$POOL \
  --set inferencePool.targetPorts[0].number=8000 \
  --set provider.name=none \
  --set experimentalHttpRoute.enabled=true \
  --set experimentalHttpRoute.inferenceGatewayName=inference-gateway \
  --set inferenceExtension.image.hub=$REGISTRY \
  --set inferenceExtension.image.name=epp \
  --set inferenceExtension.image.tag=$EPP_TAG \
  --set inferenceExtension.monitoring.prometheus.enabled=true \
  --set inferenceExtension.flags.v=2 \
  --set-json 'inferenceExtension.env=[{"name":"ENABLE_EXPERIMENTAL_ADAPTIVE_CONFIGURATOR","value":"false"}]'
```

`provider.name=none` is what the guide uses for kgateway and NGINX; `experimentalHttpRoute` creates the HTTPRoute to the Gateway named `inference-gateway` (created in step 2).

The adaptive configurator is **off** by default here. Doc 002 switches it per run.

### 6. Prometheus

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm upgrade -i kube-prometheus-stack prometheus-community/kube-prometheus-stack \
  --namespace monitoring --create-namespace \
  --set prometheus.prometheusSpec.serviceMonitorSelectorNilUsesHelmValues=false \
  --set prometheus.prometheusSpec.podMonitorSelectorNilUsesHelmValues=false

# EPP metrics need the metrics-reader secret the chart's ServiceMonitor refers to
kubectl apply -f config/manifests/metrics-reader.yaml
```

Scrape the vLLM pods too:

```yaml
apiVersion: monitoring.coreos.com/v1
kind: PodMonitor
metadata:
  name: vllm-model
  namespace: epp-lab
spec:
  selector:
    matchLabels: { app: vllm-model }
  podMetricsEndpoints:
  - port: http
    path: /metrics
    interval: 5s
```

### 7. Smoke test

```bash
IP=$(kubectl get gateway/inference-gateway -o jsonpath='{.status.addresses[0].value}')
PORT=80
curl -s ${IP}:${PORT}/v1/completions -H 'Content-Type: application/json' -d "{
  \"model\": \"$MODEL\",
  \"prompt\": \"Write as if you were a critic: San Francisco\",
  \"max_tokens\": 32
}" | jq .

# EPP picked a pod?
kubectl logs deploy/$POOL-epp --tail=20 | grep -i -E 'selected|profile|pod'
```

## Verify before moving to doc 002

| Check | Command | Expect |
|---|---|---|
| 4 vLLM pods ready, 1 GPU each | `kubectl get pods -l app=vllm-model -o wide` | 4 × `1/1 Running`, spread over GPUs |
| InferencePool accepted | `kubectl get inferencepool $POOL -o yaml \| yq .status` | `Accepted=True`, `ResolvedRefs=True` |
| HTTPRoute attached | `kubectl get httproute -o yaml \| yq '.items[].status'` | parent `inference-gateway`, `Accepted=True` |
| EPP sees 4 endpoints | `kubectl logs deploy/$POOL-epp \| grep -c 'pod added'` or EPP metric `inference_pool_ready_pods` | 4 |
| Prometheus scrapes both | Prometheus UI → Targets | EPP ServiceMonitor and vLLM PodMonitor `UP` |
| Prefix caching active | `curl <pod-ip>:8000/metrics \| grep prefix_cache` | `vllm:prefix_cache_hits` / `queries` counters increase after two identical prompts |

## Switching EPP behavior (used by doc 002)

| Run | What changes | How |
|---|---|---|
| **A · static** | fixed weights, default plugins config | `ENABLE_EXPERIMENTAL_ADAPTIVE_CONFIGURATOR=false` |
| **B · continuous (Nir)** | configurator on, defaults | `ENABLE_EXPERIMENTAL_ADAPTIVE_CONFIGURATOR=true`; note defaults make the pressure cap inactive (`CONSERVATIVE_MAX_WEIGHT=1`); set `ADAPTIVE_CONFIGURATOR_CONSERVATIVE_MAX_WEIGHT=0.6` for a second variant |
| **C · discrete (ours)** | predefined profiles + rules | requires the prototype; delivered as a plugins config file via `inferenceExtension.pluginsConfigFile` plus the image with the handler |

```bash
# flip a run without reinstalling
helm upgrade $POOL ./config/charts/inferencepool --reuse-values \
  --set-json 'inferenceExtension.env=[{"name":"ENABLE_EXPERIMENTAL_ADAPTIVE_CONFIGURATOR","value":"true"}]'
kubectl rollout status deploy/$POOL-epp
```

## Teardown

```bash
helm uninstall $POOL
kubectl delete -f vllm-model.yaml
kubectl delete -f https://raw.githubusercontent.com/kubernetes-sigs/gateway-api-inference-extension/refs/tags/${IGW_LATEST_RELEASE}/config/manifests/gateway/kgateway/gateway.yaml
helm uninstall kgateway kgateway-crds -n kgateway-system
helm uninstall kube-prometheus-stack -n monitoring
kubectl delete namespace $NS
```

## Code references

All GAIE paths are at revision `6c0f8d8496b56d68c30369520f3c4477adf8cd42` of `nirrozenbaum/gateway-api-inference-extension`.

| Ref | Location | Used for |
|---|---|---|
| 1 | `site-src/guides/index.md` | CRD install, kgateway install, gateway manifest, smoke test |
| 2 | `config/manifests/vllm/gpu-deployment.yaml` | Base for the vLLM deployment (Nir's uses Qwen3-32B, TP=2, 8 replicas, node-pinned) |
| 3 | `config/charts/inferencepool/values.yaml` | Image, env, `pluginsConfigFile`, monitoring knobs |
| 4 | `site-src/_includes/epp.md` | `helm install` flags per provider |
| 5 | `config/manifests/metrics-reader.yaml` | Secret the EPP ServiceMonitor authenticates with |
| 6 | `cmd/epp/runner/runner.go:98-103` | The six `ADAPTIVE_CONFIGURATOR_*` env vars |
| 7 | `Makefile:198-233` | `image-build`, `image-push`, `image-load`, `image-kind` |
| 8 | `hack/adaptive-metrics.sh` | Nir's loop that curls the EPP `/metrics` for `adaptive*` series |

## 🔧 If you only do one thing

Get step 7's curl to return a completion through the gateway with the adaptive-branch EPP image. Everything else is tuning.

## 🧠 Remember it as

Four real GPUs, four real queues, one EPP you built, one Prometheus watching all five.
