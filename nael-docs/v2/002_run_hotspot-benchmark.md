# 002 · Run: the hotspot benchmark, three EPP modes, one set of numbers

## TL;DR (the whole thing in 3 lines)

- One workload that **creates a hotspot on purpose**: every request shares a long prefix, so the prefix scorer pulls them all to one pod.
- Run it three times with only the EPP mode changed: **A static**, **B continuous (Nir)**, **C discrete (ours)**. Same pods, same model, same load, same seed.
- Collect TTFT p50/p99, throughput, and the per-pod queue spread. If C is not at least on par with B on this workload, there is nothing to submit.

Status: run plan, written 2026-10-12. Run C needs the prototype from the v2 proposal and cannot be executed yet. Runs A and B can.

## The runs

```text
                 ┌────────────────────────────────────────────────────────────┐
  same setup     │  4 × vLLM pods · 1 model · prefix caching on (doc 001)      │
  (doc 001)      └────────────────────────────┬───────────────────────────────┘
                                              │
            ┌─────────────────────────────────┼─────────────────────────────────┐
            ▼                                 ▼                                 ▼
  ┌──────────────────────┐        ┌──────────────────────┐        ┌──────────────────────┐
  │ A · static           │        │ B · continuous       │        │ C · discrete         │
  │ fixed weights        │        │ configurator ON      │        │ profiles + rules     │
  │ prefix 2 · KV 1 · q 1│        │ weights follow I, P  │        │ cv / P → one profile │
  │ ✅ runs today        │        │ ✅ runs today        │        │ 🔜 needs prototype   │
  └──────────┬───────────┘        └──────────┬───────────┘        └──────────┬───────────┘
             │                               │                               │
             └───────────────────────────────┼───────────────────────────────┘
                                             ▼
                   same hotspot workload · same concurrency sweep · 3 repetitions each
                                             ▼
                   TTFT p50 / p99 · throughput · queue spread · profile switches (C)
```

A is the baseline everyone has. B is what Nir reported. C is the question.

## 1. The workload

**Goal:** make the prefix scorer concentrate traffic. All requests carry the same long system prompt, so under cache-oriented weights the EPP sends them to the one pod that already holds that prefix. That pod's queue grows while three sit idle: the hotspot.

GuideLLM is the tool Nir used (`config/manifests/guidellm.yaml`, image `ghcr.io/vllm-project/guidellm:v0.5.3`). Deploy it as a long-running pod and `exec` into it.

```bash
kubectl apply -f config/manifests/pvc.yaml            # guidellm-results-pvc, from the branch
kubectl apply -f config/manifests/guidellm.yaml       # edit namespace: nrozenba → epp-lab
kubectl exec -it guidellm -- bash
```

Inside the pod, one sweep:

```bash
GATEWAY=http://inference-gateway.epp-lab.svc:80       # or the IP from doc 001 step 7
MODEL=Qwen/Qwen3-8B

guidellm benchmark \
  --target "$GATEWAY" \
  --model "$MODEL" \
  --processor "$MODEL" \
  --rate-type concurrent \
  --rate 8,16,32,64,96 \
  --max-seconds 120 \
  --data "prefix_tokens=1024,prompt_tokens=128,output_tokens=128" \
  --output-path /data/run-A-rep1.json
```

| Argument | Why this value |
|---|---|
| `--rate-type concurrent --rate 8,…,96` | Concurrency sweep. With 4 pods × `max-num-seqs 256` queues form well before 96 only if the hotspot does; that is the point. |
| `prefix_tokens=1024` | The shared prefix. Long enough that a cache hit is worth a lot to the prefix scorer. **check**: confirm the key name with `guidellm benchmark --help` for v0.5.3; if absent, use a dataset file where every row repeats the same 1024-token system prompt. |
| `prompt_tokens=128, output_tokens=128` | Short unique tail and short generation, so the experiment is about routing, not decode time. |
| `--max-seconds 120` | Long enough for the background configurator (B) to react several times. |

Run each sweep **3 times** per mode and report the median. Name files `run-<A|B|C>-rep<n>.json`.

## 2. Switching modes between runs

| Mode | Command (from the `epp-adaptive` checkout, doc 001 step 5) |
|---|---|
| A static | `helm upgrade $POOL ./config/charts/inferencepool --reuse-values --set-json 'inferenceExtension.env=[{"name":"ENABLE_EXPERIMENTAL_ADAPTIVE_CONFIGURATOR","value":"false"}]'` |
| B continuous | same with `"value":"true"`. Record the five `ADAPTIVE_CONFIGURATOR_*` values in effect; with defaults the pressure cap is inactive. Optional B2: add `{"name":"ADAPTIVE_CONFIGURATOR_CONSERVATIVE_MAX_WEIGHT","value":"0.6"}`. |
| C discrete | image with the handler + `--set inferenceExtension.pluginsConfigFile=adaptive-profiles.yaml` (the YAML from the v2 proposal). 🔜 |

After every switch: `kubectl rollout status deploy/$POOL-epp`, then **wait 60 s** so the configurator's smoothed signals and the vLLM prefix caches start from a comparable state, then start the sweep.

Between repetitions of the same mode, restart the vLLM pods **only if** you want cold caches every time. Decide once and write it in the results log. Recommended: do not restart; warm caches are the realistic case.

## 3. What to record

### From GuideLLM (per run file)

| Metric | Where in the report | Shows |
|---|---|---|
| TTFT p50 / p99 per concurrency | `benchmarks[].metrics.time_to_first_token_ms` | The user-facing cost of a hotspot (queueing before prefill) |
| Output tokens/s per concurrency | `benchmarks[].metrics.output_tokens_per_second` | That spreading did not cost throughput |
| Request latency p99 | `benchmarks[].metrics.request_latency` | Tail behavior |
| Error / timeout count | `benchmarks[].metrics.requests_per_second` and error totals | Must be 0 for a run to count |

### From Prometheus (PromQL, over each run's time window)

| Name | Query | Shows |
|---|---|---|
| Queue per pod | `vllm:num_requests_waiting{namespace="epp-lab"}` | The hotspot itself: one line up, three flat |
| **Queue spread** (the imbalance you can see) | `max(vllm:num_requests_waiting{namespace="epp-lab"}) - min(vllm:num_requests_waiting{namespace="epp-lab"})` | Should be large in A, small in B and C |
| KV usage per pod | `vllm:gpu_cache_usage_perc{namespace="epp-lab"}` | Whether the hot pod is also memory-hot |
| Prefix hit rate per pod | `rate(vllm:prefix_cache_hits[30s]) / rate(vllm:prefix_cache_queries[30s])` | The price of distribution: hits drop when traffic spreads. **check** metric names against the vLLM version. |
| Pool averages (EPP) | `inference_pool_average_queue_size`, `inference_pool_average_kv_cache_utilization` | Sanity, note integer truncation of the queue average |
| Nir's signals (B) | `adaptive_*` series from the EPP, see `hack/adaptive-metrics.sh` | When the configurator reacted, and how far it moved D |
| Our decisions (C) | `adaptive_profile_selections_total{profile,rule,reason}` 🔜 | How often each profile was chosen and how often the regime switched |

Export each query for the run window with the Prometheus HTTP API, for example:

```bash
PROM=http://localhost:9090          # kubectl port-forward -n monitoring svc/kube-prometheus-stack-prometheus 9090
START=<unix-seconds>; END=<unix-seconds>
curl -sG "$PROM/api/v1/query_range" \
  --data-urlencode 'query=max(vllm:num_requests_waiting{namespace="epp-lab"}) - min(vllm:num_requests_waiting{namespace="epp-lab"})' \
  --data-urlencode "start=$START" --data-urlencode "end=$END" --data-urlencode "step=5" \
  > spread-run-A-rep1.json
```

## 4. How to present "better"

One table, medians of 3 repetitions, one row per concurrency level, three column groups:

| Concurrency | A static TTFT p99 | B continuous TTFT p99 | C discrete TTFT p99 | A tok/s | B tok/s | C tok/s | A spread (mean) | B spread | C spread | C profile switches / min |
|---|---|---|---|---|---|---|---|---|---|---|
| 8 | | | | | | | | | | |
| 16 | | | | | | | | | | |
| 32 | | | | | | | | | | |
| 64 | | | | | | | | | | |
| 96 | | | | | | | | | | |

Plus three charts, each with the three modes overlaid:

1. **Queue per pod over time** at concurrency 64. This is the picture: A shows one line climbing, B and C show four lines close together.
2. **TTFT p99 vs concurrency.**
3. **Prefix hit rate vs concurrency.** Shows what distribution costs; the honest counterweight.

The claim to make, and nothing more: *on a shared-prefix hotspot workload, discrete profile selection reaches the same tail latency as continuous weight adaptation, with N profile switches per minute and a human-readable reason for each.* If the numbers do not support it, say that instead.

## 5. Pitfalls that invalidate a run

| Pitfall | Symptom | Fix |
|---|---|---|
| Prefix caching off or prefix too short | No hotspot in A; all three modes look identical | Check `vllm:prefix_cache_hits` increases; use `prefix_tokens ≥ 1024` |
| `max-num-seqs` too high for the load | Queues stay 0, pressure stays 0, nothing to adapt to | Lower to 64, or raise concurrency |
| Changing two things at once | Cannot attribute a difference | Only the EPP env / config differs between A, B, C |
| One repetition | Jitter dominates | 3 repetitions, median, show spread |
| Comparing B with unknown pressure settings | Unfair to Nir's design | Record the five env values; run B and B2 (`Cmax 0.6`) |
| Configurator still warming up | B's first 30 s look like A | 60 s settle time after every switch |
| Model download during a run | TTFT spikes unrelated to routing | `kubectl rollout status` before starting; keep the HF cache on a PVC |

## 6. Results log

Keep one file per campaign under `nael-docs/v2/results/<YYYY-MM-DD>-hotspot/`: the GuideLLM JSONs, the exported PromQL JSONs, and a `README.md` with the table from section 4 and these header lines:

```text
Date · cluster · GPU type × count · model · vLLM version · EPP image tag (6c0f8d8 / ours)
Mode A env · Mode B env (all five ADAPTIVE_* values) · Mode C plugins config
Workload line (the exact guidellm command) · repetitions · cache policy (warm / cold)
```

## Code references

| Ref | Location | Used for |
|---|---|---|
| 1 | `config/manifests/guidellm.yaml`, `config/manifests/pvc.yaml` (branch `6c0f8d8`) | Load generator pod and its results PVC |
| 2 | `hack/adaptive-metrics.sh` (branch `6c0f8d8`) | Reading Nir's `adaptive*` metrics from the EPP |
| 3 | `cmd/epp/runner/runner.go:98-103` | The env vars that define modes A and B |
| 4 | `pkg/epp/backend/metrics/logger.go:77-104` | How `inference_pool_average_*` are computed (integer division on the queue) |
| 5 | `nael-docs/v2/epp-upstream_issue-template.md` | Mode C plugins config and the `adaptive_profile_selections_total` metric |
| 6 | `config/manifests/benchmark/benchmark.yaml` (branch `6c0f8d8`) | Alternative load tool (`latency_throughput_curve.sh`, ShareGPT); not used here because it has no shared-prefix knob |

## 🔧 If you only do one thing

Run A once at concurrency 64 and look at `vllm:num_requests_waiting` per pod. If you do not see one pod's queue climbing while the others stay flat, fix the workload before running anything else.

## 🧠 Remember it as

Make the hotspot on purpose, then change one switch at a time and watch four queues.
