# Plan: try the inference-networking-controller proposal, two tracks

**Date:** 2026-09-06 · **Context ticket:** RUN-42428 · **Markers:** ✅ EXISTS upstream or in cluster · 🆕 NEW, we apply it · 🔜 PLANNED, not built or unverified

## TL;DR (the whole thing in 3 lines)

- **Track A, no Dynamo:** new fake-GPU env, install Gateway API + GIE + agentgateway, run CPU vLLM simulators, then apply **by hand** the five objects the page's controller would create. Proves every claim on the page without the controller code, which is not on this machine.
- **Track B, Dynamo:** the honest answer is "mostly needs real GPUs." What we *can* do on fake GPUs is run Dynamo's CPU `mocker` engine on its native Frontend topology, then try one unverified experiment: add an EPP component to it. The documented Gateway-API path with Dynamo's EPP needs GPU workers and a separate env because Dynamo pins much older GIE and agentgateway versions.
- **Order:** A first, all of it. Then B0 and B1 on a second fake-GPU env. B2 only if a Real-GPU env is approved.

## Context

Nael has two documents about the `inference-networking-controller` prototype: the pitch page
(`~/Downloads/architecture-inference-integration.html`) and the landscape report
(`~/dev/z_other/inference-networking-controller-oss-landscape-report.md`). He wants to try the proposal
for real. Findings that shaped this plan:

| Finding | Source | Consequence |
|---|---|---|
| Controller source is **not on this machine**; `~/dev/z_inference-networking-controller` holds only the two docs. Code is on Camiel's machine | local search | Track A applies the controller's outputs as static YAML. User chose this |
| User wants a **new env-in-a-click env with fake GPUs**, not `nasbi-env-01` | user | Env creation is step A0. Fake GPUs are enough for CPU simulators |
| `nasbi-env-01` already has the Dynamo operator (`dynamo-system`) but fake GPUs and no Gateway API | cluster check | Reference only; shows the wizard *can* ship Dynamo, worth checking in the new env |
| GIE v1.6.0 (2026-08-17) moved EPP, `InferenceObjective`, `InferenceModelRewrite`, `EndpointPickerConfig` to llm-d. GIE ships only `InferencePool` now | GIE release notes | The page's API groups are stale. Use GIE **v1.5.0** for the page-faithful run, llm-d for the report-faithful run |
| GIE reference EPP image last tag `v1.5.0`; docs now say use llm-d Router | registry + GIE site | Page-faithful EPP = `registry.k8s.io/gateway-api-inference-extension/epp:v1.5.0` |
| agentgateway charts are at `oci://cr.agentgateway.dev/charts/`, current v1.5.0. The page's ghcr path and v1.0.1 are wrong | agentgateway docs | Use the correct registry |
| llm-d Router v0.10.0: EPP `ghcr.io/llm-d/llm-d-router-endpoint-picker:v0.10.0`, chart `oci://ghcr.io/llm-d/charts/llm-d-router-gateway`, `EndpointPickerConfig` is `llm-d.ai/v1alpha1`, `InferenceObjective` is `llm-d.ai/v1alpha2`, header `x-llm-d-inference-objective`, flow control behind `featureGates: ["flowControl"]` | llm-d repo | Report-faithful variant of Track A |
| Simulator: `ghcr.io/llm-d/llm-d-inference-sim:v0.11.2`, CPU only, port 8000 | GHCR | No GPU needed for Track A |
| Dynamo v1.4.2. Its GAIE install script pins Gateway API v1.5.1, **GIE v1.2.1**, **agentgateway v1.0.0**. Operator creates `${DGD}-pool` + `${DGD}-epp` when a DGD has a `type: epp` component. Workers need GPUs. `dynamo.mocker` is CPU-only but documented only on the Frontend topology. Standalone EPP for vanilla vLLM is unreleased and needs GPUs | ai-dynamo/dynamo main | Track B cannot share a cluster with Track A's versions. B on fake GPUs is partly an experiment |

## Environment

```
   ┌──────────────────────────────────────────────────────────────────────┐
   │ env-in-a-click wizard  ✅   https://argocd-infra-main-02.runailabs.com │
   └───────────────┬──────────────────────────────────┬───────────────────┘
                   │ creates                          │ creates
                   ▼                                  ▼
   ┌───────────────────────────────┐   ┌───────────────────────────────┐
   │ env A  nasbi-run-42428    🆕  │   │ env B  nasbi-run-42428-dyn 🆕 │
   │  Kind · Fake GPU · Self-Hosted│   │  Kind · Fake GPU · Self-Hosted│
   │  Gateway API 1.6 · GIE 1.5    │   │  Dynamo 1.4.2 · its own pins  │
   │  agentgateway 1.5 · llm-d 0.10│   │  GIE 1.2.1 · agentgateway 1.0 │
   └───────────────────────────────┘   └───────────────────────────────┘
          Track A                             Track B0 / B1
                                                       │ only if approved
                                                       ▼
                                       ┌───────────────────────────────┐
                                       │ env C  EKS · Real GPU     🔜  │
                                       │  Track B2, documented path    │
                                       └───────────────────────────────┘
```

**Caption:** two cheap fake-GPU envs keep the incompatible version sets apart. A third, real-GPU env is optional.

Wizard choices for A and B: distribution **Kind**, Run:ai **Self-Hosted** (default, keeps Run:ai available as an optional pod owner), GPU **Fake** (default 16 × Tesla-V100), third-party: nginx (default), Knative off, others off. Env names are assumptions; change if the ticket differs.

```
Pseudocode: get a cluster
1. Run the env-in-a-click skill with the choices above.                       🆕
2. Wait for ArgoCD to report the env healthy.
3. kc add --cluster nasbi-run-42428 --session                                 ✅ scopes KUBECONFIG to this terminal
4. Ask: kubectl get nodes shows worker-0-0 with nvidia.com/gpu 16?  → if no, wait or kc-debug.  ✗
5. Ask: kubectl get crd | grep -i dynamo returns anything?  → if yes, note it, the wizard now ships Dynamo.
6. Confirm no gatewayclass resource exists.                                  ✓ clean slate for Gateway API
```

---

## Track A — the page's stack, no Dynamo, no controller code

### A.1 What we install vs what we apply

```
 LAYER            OBJECT                                   MARKER   WHO
 ─────────────    ───────────────────────────────────────  ──────   ──────────────────────
 prerequisites    Gateway API CRDs v1.6.0 standard          ✅       kubectl apply, admin step
                  GIE CRDs v1.5.0 (InferencePool +          ✅       kubectl apply, admin step
                    x-k8s.io Objective, ModelRewrite)
                  agentgateway-crds + agentgateway v1.5.0   ✅       helm, admin step
                  Gateway  inference-agw                    ✅       kubectl apply, admin step
 pods             llm-d-inference-sim ×3, label             ✅       kubectl apply, "pod owner" step
                    app: vllm-llama3
 the controller's InferencePool  llama3-pool                🆕       we apply by hand
 output, by hand  EPP SA + ConfigMap + Deployment + Service 🆕       we apply by hand
                  HTTPRoute  llama3-route                   🆕       we apply by hand
                  InferenceObjective ×3 tiers               🔜       we apply by hand, Phase 1
                  InferenceModelRewrite alias + canary      🔜       we apply by hand, Phase 1
```

**Caption:** the 🆕 rows are exactly the objects the page says the controller creates. Applying them by hand is the controller with a human as the reconcile loop.

### A.2 Install order

```
Pseudocode: Track A
 1. kubectl apply --server-side -f gateway-api v1.6.0 standard-install.yaml                    ✅
 2. kubectl apply -f GIE v1.5.0 manifests.yaml                                                 ✅
    → expect CRDs: inferencepools.inference.networking.k8s.io,
      inferenceobjectives + inferencemodelrewrites .inference.networking.x-k8s.io
 3. helm upgrade -i agentgateway-crds oci://cr.agentgateway.dev/charts/agentgateway-crds --version v1.5.0 -n agentgateway-system --create-namespace   ✅
 4. helm upgrade -i agentgateway      oci://cr.agentgateway.dev/charts/agentgateway      --version v1.5.0 -n agentgateway-system --set inferenceExtension.enabled=true   ✅
    → kubectl get gatewayclass agentgateway  shows ACCEPTED True
 5. kubectl create ns model-serving; apply Gateway inference-agw, class agentgateway, HTTP :80  ✅
    → kubectl wait --for=condition=Programmed gateway/inference-agw
    → no LoadBalancer in Kind: the Gateway Service stays Pending for an external IP. Use port-forward. That is fine.
 6. Apply sim Deployment: 3 replicas, image llm-d-inference-sim:v0.11.2, args --model llama3 --port 8000,
    labels app: vllm-llama3 and llm-d.ai/engine-type: vllm, port 8000 named http                ✅ pod owner step
    → optional 6b: add resources nvidia.com/gpu: 1 + schedulerName runai-scheduler + project label
      to exercise Run:ai as the pod owner on fake GPUs. Sim never touches CUDA, so it still runs.
 7. Apply InferencePool llama3-pool: selector app: vllm-llama3, targetPorts 8000,
    endpointPickerRef llama3-epp port 9002, failureMode FailOpen                                 🆕
 8. Apply EPP: ServiceAccount + RBAC, ConfigMap kind EndpointPickerConfig with
    queue-scorer 30 / kv-cache 20 / prefix-cache 10, Deployment image
    registry.k8s.io/gateway-api-inference-extension/epp:v1.5.0 with --pool-name llama3-pool,
    secure-serving left at default true, Service port 9002 + 9090                                🆕
    → kubectl logs deploy/llama3-epp | grep secure-serving  shows true
 9. Apply HTTPRoute llama3-route: parentRef inference-agw, hostname llama3.test,
    path /v1, backendRef group inference.networking.k8s.io kind InferencePool name llama3-pool,
    timeouts.request 300s                                                                       🆕
    → status.parents[0].conditions Accepted True, ResolvedRefs True
10. Smoke test, see A.3.                                                                        ✓
11. Phase 1: apply 3 InferenceObjectives (prod-gold 100, internal 10, batch-jobs −1)
    and 1 InferenceModelRewrite (llama3 → 90/10 split), x-k8s.io/v1alpha2 groups              🔜
12. Swap test, the page's "pluggable EPP" claim: change only the EPP Deployment image to
    ghcr.io/llm-d/llm-d-router-endpoint-picker:v0.10.0 and the ConfigMap apiVersion to
    llm-d.ai/v1alpha1. Ask: does routing still work?  → if yes, the claim holds for this pair.  ✓
    → if the x-k8s.io Objectives are ignored by llm-d, that is the report's stale-API finding, confirmed live.  ✗ expected
```

### A.3 Verification, straight from the page's smoke-test guide

| # | Check | Command | Expect |
|---|---|---|---|
| 1 | Stack health | `kubectl get pods -n model-serving` | EPP 1/1, 3 sim pods Running |
| 2 | Route backend kind | `kubectl get httproute llama3-route -n model-serving -o jsonpath='{.spec.rules[0].backendRefs}'` | `kind=InferencePool` |
| 3 | Route accepted | `... -o jsonpath='{.status.parents[*].conditions[*]}'` | `Accepted=True`, `ResolvedRefs=True` |
| 4 | EPP sees pods | port-forward EPP 9090, `curl localhost:9090/metrics \| grep inference_pool_ready_pods` | `3` |
| 5 | Sim works alone | port-forward a sim pod 8000, POST `/v1/chat/completions` model `llama3` | JSON with `choices[0]` |
| 6 | End to end | `kubectl port-forward -n agentgateway-system svc/<gateway-svc> 8080:80`, then `curl -H 'Host: llama3.test' localhost:8080/v1/chat/completions ...` | JSON completion, EPP log `Scheduling succeeded` |
| 7 | Load awareness | fire 30 parallel requests, then `kubectl logs deploy/llama3-epp` | picks are not round-robin, busy pods skipped |
| 8 | Phase 1 header | repeat 6 with `-H 'x-gateway-inference-objective: prod-gold'` | request served, EPP log shows objective resolved. **Not** shed under load, flow control is off, as the page admits |
| 9 | Canary | 20 requests, inspect sim logs for the `model` field received | roughly 18 stable / 2 canary |
| 10 | Two models on one Gateway | apply a second pool + route for `mistral` with only `/v1` path | collision unless hostnames differ, confirms report §12 |

### A.3b Results, run on `nasbi-run-42428` on 2026-09-06

| # | Check | Result | Evidence |
|---|---|---|---|
| 1 | Stack health | ✅ | EPP 1/1, 3 sims Running |
| 2 | Route backend kind | ✅ | `kind=InferencePool` |
| 3 | Route accepted | ✅ | `Accepted=True ResolvedRefs=True` |
| 3b | EPP TLS | ✅ | flags log `"secure-serving":true` |
| 4 | EPP sees pods | ✅ after RBAC fix | `inference_pool_ready_pods{name="llama3-pool"} 3`. Needed `system:auth-delegator` + `/metrics` reader for the token |
| 5 | Sim works alone | ✅ | direct completion |
| 6 | End to end | ✅ | completion via agentgateway → EPP → sim |
| 7 | Load awareness | ✅ | 30-request burst spread 11 / 7 / 4, not round-robin |
| 8 | Phase 1 header | ✅ tracked, not enforced | `objectiveKey=batch-jobs priority=-1` in EPP log, request still served |
| 9 | Canary | ✅ | 4 of 20 rewritten to `llama3-v2` (rule 90/10) |
| 10 | Two `/v1` routes on one Gateway | ✅ collision reproduced | both Accepted=True, older route silently took 10/10 requests |
| 12 | EPP swap GIE → llm-d v0.10.0 | ✅ works | routing intact; llm-d still honors `x-k8s.io` objectives and rewrite, both header names |

Fixes made along the way: smoke script `wait` blocked on port-forwards; EPP metrics need TokenReview rights; sim needs `--lora-modules` so the canary target exists.

### A.4 What Track A proves and does not prove

| Page claim | Provable here | Note |
|---|---|---|
| One CR creates the stack | No, no controller | We prove the *stack* works, not the *automation* |
| Smart routing beats round-robin | Yes | Check 7 |
| Only agentgateway supports InferencePool | No, and the report says it is false | Optional A.5: repeat with Istio |
| Pluggable EPP by image swap | Partly | Check 12. Works for GIE→llm-d. Report says Dynamo is different, see Track B |
| Lanes | Tracked, not enforced | Check 8. Enforcing needs `featureGates: ["flowControl"]` in an llm-d config, that is the report's Phase 0 fix |
| Canary via rewrite | Yes | Check 9 |

### A.5 Optional, report-faithful variant

Replace steps 7 to 9 with the llm-d chart, which is what the report says the controller should generate:

```
helm upgrade -i llama3 oci://ghcr.io/llm-d/charts/llm-d-router-gateway --version v0.10.0 -n model-serving \
  --set router.modelServers.matchLabels.app=vllm-llama3 \
  --set router.epp.resources.requests.cpu=100m --set router.epp.resources.requests.memory=128Mi \
  --set router.epp.resources.limits.memory=512Mi \
  --set provider.name=none --set httpRoute.create=true --set httpRoute.inferenceGatewayName=inference-agw
```
Then compare the rendered objects with our hand-written ones. The diff is the report's §12 gap list, made concrete. Chart defaults request 8 CPU / 8 Gi, so the resource overrides are mandatory on Kind.

### A.6 Later, when the controller source arrives from Camiel

```
1. Clone next to the docs, e.g. ~/dev/inference-networking-controller.
2. make install (CRDs), then KUBECONFIG=<session> go run ./cmd/main.go --metrics-bind-address=:8114 --health-probe-bind-address=:8115
3. Delete the hand-applied 🆕 objects, apply one ModelEndpoint with the same names.
4. Ask: does the controller recreate the identical objects?  → diff against A.2 output.  ✓
```

---

## Track B — Dynamo

### B.0 Reality check before touching anything

```
 What the page implies                      What upstream says today
 ─────────────────────────────────────────  ────────────────────────────────────────────────────────
 "Dynamo EPP can replace the EPP box"       True, but only for pods that Dynamo's operator created
                                            (a DynamoGraphDeployment with a type: epp component)
 "needs its operator"                       Confirmed. Operator auto-creates ${DGD}-pool and ${DGD}-epp
 fake GPUs are fine                         ✗ Dynamo workers need real GPUs
                                            ✓ except dynamo.mocker, CPU only, but documented
                                              only on the Frontend topology, never behind an EPP
 same cluster as Track A                    ✗ Dynamo's script pins GIE v1.2.1 + agentgateway v1.0.0
                                              vs Track A's GIE v1.5.0 + agentgateway v1.5.0
 standalone Dynamo EPP for vanilla vLLM     unreleased, build-your-own image, and still needs GPUs
```

**Caption:** on fake GPUs, Track B is one confirmed step and one experiment. The full documented path is Track B2 and needs real GPUs.

### B.1 On env B, fake GPUs: Dynamo-native topology with mocker  ✅ documented

```
Pseudocode
1. kc add --cluster nasbi-run-42428-dyn --session
2. Ask: is the Dynamo operator already there (kubectl get crd | grep dynamo)?
   → if yes, skip 3.   → if no, install:
3. helm upgrade --install dynamo-platform https://helm.ngc.nvidia.com/nvidia/ai-dynamo/charts/dynamo-platform-1.4.2.tgz \
     -n dynamo-system --create-namespace --wait --timeout=10m                                  ✅
   → skip gpu-operator, the env has the fake one and mocker requests no GPU
4. Apply examples/backends/mocker/deploy/agg.yaml: Frontend + one decode worker running
   python3 -m dynamo.mocker --model-path Qwen/Qwen3-0.6B                                        ✅
5. kubectl port-forward svc/mocker-agg-frontend 8000:8000; POST /v1/chat/completions             ✓
6. Result: the page's "Dynamo-native frontend routing" box, running on CPU.
```

### B.2 On env B, fake GPUs: mocker behind Dynamo's EPP  🔜 unverified experiment

```
Pseudocode
1. Run Dynamo's install_gaie_crd_agentgateway.sh, which applies its own pins:
   Gateway API v1.5.1, GIE v1.2.1, agentgateway v1.0.0, plus AgentgatewayParameters
   disabling istio injection, Gateway in agentgateway-system.                                  ✅ Dynamo's versions
2. Add a component  type: epp  to the mocker DGD, no eppConfig, image dynamo-frontend:1.4.2
   (1.5.0 Rust EPP is not released; the doc's 1.5.0 references are forward-looking).
   Add frontendSidecar to the decode worker as the doc requires.                               🔜
3. Ask: did the operator create mocker-agg-pool and mocker-agg-epp?  → if no, stop, record why.  ✗
4. Apply HTTPRoute: backendRef InferencePool mocker-agg-pool port 8000,
   header match X-Gateway-Model-Name: Qwen/Qwen3-0.6B, timeout 300s.                           🆕 the ExistingPool story
5. curl through the Gateway. Ask: does the mocker answer?
   → if yes: Dynamo EPP + InferencePool + our HTTPRoute proven on CPU. Big win.  ✓
   → if no: mocker does not publish what the EPP needs. Documented gap, move to B3.  ✗ acceptable
```

This is exactly the report's `ExistingPool` mode: Dynamo owns pool + EPP + pods, we own only the route.

### B.3 On env C, real GPUs: the documented path  🔜 needs approval and budget

Only if B.2 fails or a real workload is required. New env-in-a-click env, EKS, GPU type Real. Then the Dynamo GAIE doc as written: DGD with vLLM workers publishing KV events, `type: epp` component, our HTTPRoute. Not planned in detail until approved.

---

### B.4 Results, run on `nasbi-run-42428-dyn` on 2026-09-06

| Step | Result | Evidence |
|---|---|---|
| B.0 env + Dynamo | ✅ | Wizard shipped Dynamo operator **1.0.1** (`nvidia.com/v1alpha1`, `services` map), NATS. Not 1.4.2 as planned |
| Gateway stack, Dynamo pins | ✅ | Gateway API 1.5.1, GIE 1.2.1, agentgateway 1.0.0, Gateway `inference-gateway` Programmed |
| B.1 mocker via Dynamo Frontend | ✅ | completion for `Qwen/Qwen3-0.6B`; mocker publishes KV events over NATS, block size 64 |
| B.2 Dynamo EPP behind Gateway | ⛔ blocked | `nvcr.io/nvidia/ai-dynamo/epp-image` is gated (403 anonymous). Needs an NGC pull secret only the user can create |
| B.2′ our EPP over Dynamo pods | ✅ | GIE EPP v1.2.1 + our pool selecting Dynamo Frontend pods + our route. Completion returned. Dynamo as a pod owner works |

Images: public on NGC `kubernetes-operator`, `dynamo-frontend`, `vllm-runtime`, `sglang-runtime`, `dynamo-planner` (no 1.0.x tag). Gated: `mocker-runtime`, `epp-image`. B.1 runs on `dynamo-frontend:1.0.1` + `vllm-runtime:1.0.1`.

**To unblock B.2 yourself** (I must not handle the key): in a terminal with `KUBECONFIG=~/.kube/config.nasbi-run-42428-dyn`, create a docker-registry secret named `ngc-pull` in namespace `dynamo-lab` for server `nvcr.io`, user `$oauthtoken`, password = your NGC API key, then say "b2". The DGD will get `imagePullSecrets` and the `componentType: epp` service from `03-mocker-agg-epp.yaml`, rewritten to the 1.0.1 schema.

## What could go wrong, and the answer

| Risk | Mitigation |
|---|---|
| Gateway Service stuck Pending, no LoadBalancer in Kind | Expected. Use `port-forward` for every test |
| agentgateway `endpointPickerRef.port must be specified` | Always set `port: 9002` in the pool |
| `ext_proc failed` in gateway logs | EPP not on TLS. Confirm `secure-serving: true` in EPP logs |
| GIE v1.5.0 and llm-d v0.10 both define `InferenceObjective` under different groups | Install GIE first, then llm-d CRDs. Both can coexist, they are different groups. Objectives in the wrong group are silently ignored, that is the point of check 12 |
| llm-d chart OOM-kills on Kind | Set the resource overrides in A.5 |
| Two tracks on one cluster fight over GIE versions | Separate envs, by design |
| `kc list` crashes | 0-byte `~/.kube/config.local`. Cosmetic, not fixed by this plan |
| Env creation fails or hangs | `/kc-debug` skill against the ArgoCD app |
| Leaving envs running | `kc hibernate --cluster <name>` after each session |

## Deliverables

1. Env A and env B created, kubeconfigs via `kc add --session`.
2. A manifests folder for Track A, one file per object, applied in order. Location to be agreed, suggested `docs/inference-networking-controller-lab/track-a/`.
3. A results table filled in from A.3 checks 1 to 12 and B.1 / B.2 outcomes.
4. An addendum to `docs/inference-networking-controller-oss-view.md`: "What we proved live" with the same ✅ / 🆕 / 🔜 markers.

## Code references

| Ref | Location | Status |
|---|---|---|
| agentgateway EPP TLS behaviour | `controller/pkg/agentgateway/plugins/inference_plugin.go` in agentgateway main | ✅ upstream, read by research |
| llm-d chart pool gating | `routerlib/templates/_inferencepool.yaml`, `router.inferencePool.create` | ✅ upstream |
| llm-d headers | `pkg/epp/metadata/consts.go` | ✅ upstream |
| Dynamo GAIE install script | `deploy/inference-gateway/scripts/install_gaie_crd_agentgateway.sh` | ✅ upstream |
| Dynamo mocker example | `examples/backends/mocker/deploy/agg.yaml` | ✅ upstream |
| Controller source | Camiel's machine | ❌ not available |

## 🔧 If you only do one thing

Run Track A steps 1 to 10 on env A. Ninety minutes, no GPUs, and it settles whether the page's core promise, smarter-than-round-robin routing through an InferencePool, actually works with today's upstream versions.

## 🧠 Remember it as

Track A is the page with a human as the controller. Track B is a neighbor with its own stack, and on fake GPUs we can only knock on its door.

---

## Status log

Updated as work progresses. Newest entry first.

| When | Step | Status | Notes |
|---|---|---|---|
| 2026-09-07 11:30 | Envs A + B woken | ✅ done | Both auto-hibernated overnight (label `hibernating`, auto-delete after 1 week ≈ 13 Sep). Woke via `env-in-a-click/status: wake-up`. Stacks intact on both: routes Accepted/ResolvedRefs, pods restarted and passing probes |
| 2026-09-06 17:10 | B.2′ · our GIE EPP + pool over Dynamo Frontend pods | ✅ passed | agentgateway 1.0.0 → our HTTPRoute (host `mocker.test`) → our InferencePool (selector `nvidia.com/dynamo-component-type: frontend`) → GIE EPP v1.2.1 → Dynamo Frontend pod → mocker. Completion returned. Proves Dynamo as a **pod owner** under ManagedEPP. Two rename bugs fixed on the way (SA name, `--pool-name`). Not the same as B.2, which needs Dynamo's own gated EPP image |
| 2026-09-06 16:58 | B.1 · mocker on Dynamo Frontend, CPU | ✅ passed | DGD `mocker-agg` Ready. `POST /v1/chat/completions` via `svc/mocker-agg-frontend` returned a completion for `Qwen/Qwen3-0.6B`. Mocker log: KV events published over NATS, `block_size 64`, discovery via Kubernetes. Images: `dynamo-frontend:1.0.1` + `vllm-runtime:1.0.1` |
| 2026-09-06 16:58 | B.2 · Dynamo EPP behind the Gateway | ⛔ blocked on credentials | Needs `nvcr.io/nvidia/ai-dynamo/epp-image` (403 anonymous). Requires an NGC pull secret in `dynamo-lab` that only the user can create. Fallback B.2′ in progress: our public GIE EPP + InferencePool in front of the Dynamo Frontend pods, i.e. Dynamo as a pod owner under ManagedEPP |
| 2026-09-06 16:52 | NGC image access map | ✅ found | Anonymous pull probe: **public** `kubernetes-operator`, `dynamo-frontend`, `vllm-runtime`, `sglang-runtime`, `dynamo-planner` (no 1.0.x tag). **Gated (403)** `mocker-runtime`, `epp-image`. B.1 switched to `dynamo-frontend:1.0.1` + `vllm-runtime:1.0.1` (wheel includes `dynamo.mocker`). B.2 needs `epp-image` → requires an NGC pull secret the user must create |
| 2026-09-06 16:45 | B.1 · mocker DGD on Dynamo 1.0.1 | ⛔ blocked on image pull | DGD accepted (schema `nvidia.com/v1alpha1`, `services` map). Operator created Frontend + decode deployments and Services. But `nvcr.io/nvidia/ai-dynamo/mocker-runtime` returns **403 on anonymous pull** for both `1.0.1` and `1.0.0`, while `kubernetes-operator:1.0.1` pulled fine → gated NGC repo, cluster has no NGC pull secret. Probing alternatives |
| 2026-09-06 16:34 | B0 · env `nasbi-run-42428-dyn` up, Dynamo present | ✅ done, ⚠️ version | Cluster up in ~5 min. Wizard shipped **Dynamo operator 1.0.1** (`kubernetes-operator:1.0.1`) + NATS, not 1.4.2. B.2 manifest assumed the 1.4.2 legacy `type: epp` contract; must re-check against the 1.0.1 CRD. Gateway stack installed with Dynamo's pins, Gateway `inference-gateway` Programmed |
| 2026-09-06 16:30 | Track B manifests written | ✅ done | `track-b/`: check script, B.1 mocker DGD, Dynamo-pinned GAIE prereqs (Gateway API 1.5.1 · GIE 1.2.1 · agentgateway 1.0.0), B.2 DGD with legacy `type: epp` (`epp-image`, `eppConfig` from Dynamo v1.4.2 example), our HTTPRoute, smoke test. Three assumptions listed in `03-mocker-agg-epp.yaml` |
| 2026-09-06 16:20 | B0 · env `nasbi-run-42428-dyn` provisioning | 🔄 in progress | ArgoCD app created via CLI, same chart as env A plus `dynamo.enabled=true`, `dynamo.grove.enabled=true`. URL: https://argocd-infra-main-02.runailabs.com/applications/nasbi-run-42428-dyn |
| 2026-09-06 16:18 | Track A wiring diagram | ✅ done | `track-a/wiring.png` + `wiring.mmd`, link table in `track-a/README.md`. Eleven links, three kinds: name ref, label selector, mount |
| 2026-09-06 16:08 | Check 12 · pluggable EPP, GIE → llm-d Router v0.10.0 | ✅ works, with a surprise | Image swap only, same flags. Rollout clean, route stayed resolved, e2e completion returned. **Surprise:** llm-d v0.10.0 still watches `inference.networking.x-k8s.io` InferenceObjective and honored our objectives (`objectiveKey=batch-jobs priority=-1`) and rewrite (`llama3-v2` seen), with both header names. No `llm-d.ai` CRDs installed. So the page's 'pluggable EPP' claim holds for this pair, and the report's stale-group gap is not yet breaking. EPP reverted to GIE baseline afterwards |
| 2026-09-06 16:02 | Check 4 · EPP metrics | ✅ 9/9 Phase 0 | Root cause: EPP metrics endpoint authenticates bearer tokens via TokenReview, and the EPP SA lacked `system:auth-delegator`. Added ClusterRoleBinding + a `/metrics` reader ClusterRole to `04-epp.yaml`. `inference_pool_ready_pods{name="llama3-pool"} 3` |
| 2026-09-06 15:58 | Check 10 · `/v1` route collision (report §12) | ✅ reproduced | Two hostless `/v1` HTTPRoutes on the same Gateway, one → InferencePool, one → Service. **Both** show Accepted=True ResolvedRefs=True. 10 requests with an unmatched Host all went to the older route (EPP saw 10/10). Silent collision confirmed; hostnames or body-based model routing are required. Test objects cleaned up |
| 2026-09-06 15:52 | A.2 step 11 · Phase 1 objectives + rewrite, checks 8-9 | ✅ done | 3 InferenceObjectives + 1 InferenceModelRewrite applied (x-k8s.io/v1alpha2). Header `x-gateway-inference-objective: batch-jobs` → EPP log `objectiveKey=batch-jobs priority=-1`. Priority **tracked, not enforced**, as the page says. Canary: 4 of last 20 requests rewritten to `llama3-v2` (rule 90/10). EPP `secure-serving: true` confirmed from its flags log. Metrics endpoint returns 401 without a token; adding metrics-reader RBAC |
| 2026-09-06 15:45 | A.2 steps 5-10 · stack applied, Phase 0 smoke test | ✅ done, 8/9 | Gateway Programmed, 3 sims + EPP Running, HTTPRoute Accepted=True ResolvedRefs=True. **End-to-end request through agentgateway → EPP → sim returned a completion.** 30-request burst spread 11/7/4 across pods, i.e. load-aware, not round-robin. Check 4 (EPP metric `inference_pool_ready_pods`) failed on metric name, being investigated. Smoke script bug fixed: `wait` was blocking on port-forwards |
| 2026-09-06 15:37 | A.2 steps 1-4 · prerequisites installed | ✅ done | Gateway API v1.6.0, GIE v1.5.0 CRDs, agentgateway v1.5.0 with inferenceExtension. GatewayClass `agentgateway` Accepted=True. Ran while Run:ai install was still in progress, no conflict |
| 2026-09-06 15:36 | A0 · env `nasbi-run-42428` cluster up | ✅ done | `create-infra` + `kubeconfig-exposer` healthy after ~5 min. `create-third-party` (nginx, fake GPU) still running, `create-runai` pending. Kubeconfig saved to `~/.kube/config.nasbi-run-42428` (default context untouched) |
| 2026-09-06 15:33 | Track A manifests written | ✅ done | `docs/inference-networking-controller-lab/track-a/`: prereqs script, 7 manifests, smoke-test script, llm-d EPP swap. YAML and shell syntax validated. Schemas copied from `helm template` of the GIE v1.5.0 inferencepool chart |
| 2026-09-06 15:28 | A0 · env `nasbi-run-42428` provisioning | 🔄 in progress | ArgoCD app created via `argocd app create` (Chrome extension was offline, so CLI instead of wizard). Same chart/versions as today's viper env `dima-riga`. Kind · Fake GPU · Self-Hosted · nginx only · Dynamo off. URL: https://argocd-infra-main-02.runailabs.com/applications/nasbi-run-42428 |
| 2026-09-06 15:20 | Plan written, decisions taken | ✅ done | Manual stack first · new fake-GPU env |

**Done:** Track A complete. Track B: B.1 passed, B.2′ passed, B.2 proper blocked on an NGC pull secret. Envs A and B both still running.
**Next up:** waiting on the user for (a) NGC pull secret `ngc-pull` in `dynamo-lab` on env B → "b2", (b) controller source from Camiel → "a6", (c) "commit" for the docs. Envs auto-delete ≈ 13 Sep.
