# Track B · Dynamo owns the middle

Dynamo's operator creates the InferencePool, the EPP, and the worker pods. We create **only the HTTPRoute**. This is the report's `ExistingPool` mode. Markers: ✅ Dynamo or admin · 🆕 what we create · 🔜 experiment, undocumented upstream.

Plan and status log: [../../inference-networking-controller-lab-plan.md](../../inference-networking-controller-lab-plan.md).

## Reality check first

| Fact | Consequence |
|---|---|
| Dynamo workers need real GPUs | On a fake-GPU env only `dynamo.mocker` (CPU) can run |
| `dynamo.mocker` is documented only on the Frontend topology | B.1 is documented, B.2 (mocker behind the EPP) is an experiment |
| Dynamo's own install script pins Gateway API v1.5.1, GIE **v1.2.1**, agentgateway **v1.0.0** | Different cluster from Track A, by design |
| The env shipped Dynamo operator **1.0.1**: `nvidia.com/v1alpha1`, `spec.services` map, `componentType: frontend/worker/epp`, `eppConfig` supported | Manifests use the 1.0.1 schema |
| NGC repos `mocker-runtime` and `epp-image` are gated (403 anonymous); `dynamo-frontend`, `vllm-runtime`, `kubernetes-operator` are public | B.1 uses public images. B.2 needs an NGC pull secret |

## Files, in apply order

| # | File | Marker | What it is |
|---|---|---|---|
| 0 | `00-check-dynamo.sh` | ✅ | Confirms the operator, its version, NATS, and that no Gateway API is installed yet |
| 1 | `01-mocker-agg.yaml` | ✅ | B.1: Dynamo-native topology. Frontend + one CPU mocker decode worker. Copied from `examples/backends/mocker/deploy/agg.yaml` |
| 2 | `02-gaie-prereqs.sh` | ✅ | B.2: Dynamo's own `install_gaie_crd_agentgateway.sh` pins, Gateway `inference-gateway` in `agentgateway-system` |
| 3 | `03-mocker-agg-epp.yaml` | ⛔ | B.2: same DGD plus Dynamo's `epp` component. **Blocked**: `epp-image` is gated on NGC. Still written against the 1.4.2 schema, needs the 1.0.1 rewrite once a pull secret exists |
| 4 | `04-httproute.yaml` | 🆕 | B.2: the ONLY object we own. Header match `X-Gateway-Model-Name`, backend `mocker-agg-pool` |
| 5 | `05-smoke-test.sh` | ✓ | `b1` through the Frontend, `b2` through the Gateway with Dynamo's EPP, `b2prime` through the Gateway with our EPP |
| 6 | `06-b2prime-our-epp-over-dynamo.yaml` | 🆕 | B.2′: our GIE EPP v1.2.1 + our InferencePool selecting the Dynamo Frontend pods. Rendered from the GIE chart, names and resources adapted |
| 7 | `07-b2prime-httproute.yaml` | 🆕 | B.2′: our route, host `mocker.test`, backend `mocker-frontend-pool` |
| | `apply.sh` | | Substitutes `${DYNAMO_VERSION}` into the DGD manifests |

## Names used everywhere

| Thing | Name |
|---|---|
| Namespace for workloads | `dynamo-lab` |
| DGD | `mocker-agg` |
| Model | `Qwen/Qwen3-0.6B` |
| Pool created by the operator | `mocker-agg-pool` |
| EPP Service created by the operator | `mocker-agg-epp` |
| Gateway | `inference-gateway` in `agentgateway-system`, class `agentgateway` |
| Our HTTPRoute | `mocker-agg`, in `dynamo-lab` |

## What Track B proves

| Question | Result 2026-09-06 |
|---|---|
| Dynamo runs on CPU with mocker | ✅ B.1: completion through `svc/mocker-agg-frontend` |
| Operator creates `${DGD}-pool` + `${DGD}-epp` from an `epp` service | ⛔ B.2 not run: `epp-image` gated on NGC |
| Our HTTPRoute alone routes into a Dynamo-owned pool, `ExistingPool` | ⛔ same blocker |
| Dynamo pods work as a plain pod owner under **our** EPP + pool, `ManagedEPP` | ✅ B.2′: completion through agentgateway → our route → our pool → GIE EPP → Dynamo Frontend → mocker |

## How to unblock B.2

Only you can do this step, it involves your NGC API key. With `KUBECONFIG=~/.kube/config.nasbi-run-42428-dyn`, create a `docker-registry` secret named `ngc-pull` in namespace `dynamo-lab`, server `nvcr.io`, username `$oauthtoken`, password your NGC key. Then say "b2" and I rewrite `03-mocker-agg-epp.yaml` to the 1.0.1 schema with `imagePullSecrets`, apply it, and run `05-smoke-test.sh b2`.
