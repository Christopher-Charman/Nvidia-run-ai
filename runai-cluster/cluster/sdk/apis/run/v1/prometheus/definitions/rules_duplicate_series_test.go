//go:build prometheus_rules
// +build prometheus_rules

/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package definitions

import (
	"fmt"
	"strings"
	"testing"
)

var monitoringStacks = []string{"monitoring-a", "monitoring-b"}

func twice(series, values string) []Input {
	inputs := make([]Input, 0, len(monitoringStacks))
	for _, stack := range monitoringStacks {
		scraped := fmt.Sprintf(`%s{endpoint="%s"}`, series, stack)
		if open := strings.Index(series, "{"); open >= 0 {
			scraped = fmt.Sprintf(`%s{endpoint="%s", %s`, series[:open], stack, series[open+1:])
		}
		inputs = append(inputs, Input{Series: scraped, Values: values})
	}
	return inputs
}

func single(series, values string) []Input {
	return []Input{{Series: series, Values: values}}
}

func inputs(groups ...[]Input) []Input {
	var all []Input
	for _, g := range groups {
		all = append(all, g...)
	}
	return all
}

func TestAllRulesSurviveDuplicateThirdPartySeries(t *testing.T) {
	thirdParty := []struct{ series, values string }{
		{`kube_node_info{node="n1"}`, "1x20"},
		{`kube_node_status_capacity{node="n1", resource="cpu", unit="core"}`, "16x20"},
		{`kube_node_status_capacity{node="n1", resource="memory", unit="byte"}`, "68719476736x20"},
		{`kube_node_status_capacity{node="n1", resource="nvidia_com_gpu", unit="integer"}`, "4x20"},
		{`kube_node_status_condition{node="n1", condition="Ready", status="true"}`, "1x20"},
		{`kube_node_spec_unschedulable{node="n1"}`, "0x20"},
		{`kube_pod_info{pod="wl-pod", namespace="team-a", node="n1", uid="uid-1", pod_ip="10.0.0.5"}`, "1x20"},
		{`kube_pod_info{pod="dcgm-exporter-x", namespace="runai", node="n1", uid="uid-2", pod_ip="10.0.0.6", created_by_name="dcgm-exporter"}`, "1x20"},
		{`kube_pod_info{pod="node-exporter-x", namespace="monitoring", node="n1", uid="uid-3", pod_ip="10.0.0.7"}`, "1x20"},
		{`kube_pod_status_phase{pod="wl-pod", namespace="team-a", phase="Running"}`, "1x20"},
		{`kube_pod_completion_time{pod="wl-pod", namespace="team-a", uid="uid-1"}`, "1700000000x20"},
		{`kube_pod_container_resource_requests{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="cpu", unit="core"}`, "2x20"},
		{`kube_pod_container_resource_requests{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="memory", unit="byte"}`, "1073741824x20"},
		{`kube_pod_container_resource_limits{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="cpu", unit="core"}`, "4x20"},
		{`kube_pod_container_resource_limits{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="memory", unit="byte"}`, "2147483648x20"},
		{`kube_pod_container_resource_limits{pod="cp-pod", namespace="runai", uid="uid-4", container="cp", node="n1", resource="cpu", unit="core"}`, "1x20"},
		{`kube_pod_container_resource_limits{pod="cp-pod", namespace="runai", uid="uid-4", container="cp", node="n1", resource="memory", unit="byte"}`, "1000x20"},
		{`kube_pod_container_status_restarts_total{pod="cp-pod", namespace="runai", container="cp"}`, "0+1x20"},
		{`kube_deployment_status_replicas_available{namespace="runai", deployment="d1"}`, "1x20"},
		{`kube_deployment_spec_replicas{namespace="runai", deployment="d1"}`, "2x20"},
		{`kube_deployment_status_replicas_unavailable{namespace="runai", deployment="d1"}`, "1x20"},
		{`kube_statefulset_status_replicas_ready{namespace="runai", statefulset="s1"}`, "1x20"},
		{`kube_statefulset_status_replicas{namespace="runai", statefulset="s1"}`, "2x20"},
		{`kube_statefulset_status_replicas_available{namespace="runai", statefulset="s1"}`, "0x20"},
		{`kube_statefulset_replicas{namespace="runai", statefulset="s1"}`, "2x20"},
		{`kube_daemonset_status_current_number_scheduled{namespace="runai", daemonset="ds1"}`, "1x20"},
		{`kube_daemonset_status_desired_number_scheduled{namespace="runai", daemonset="ds1"}`, "2x20"},
		{`kube_daemonset_status_number_misscheduled{namespace="runai", daemonset="ds1"}`, "0x20"},
		{`kube_daemonset_status_updated_number_scheduled{namespace="runai", daemonset="ds1"}`, "1x20"},
		{`kube_daemonset_status_number_available{namespace="runai", daemonset="ds1"}`, "1x20"},
		{`kube_daemonset_status_number_unavailable{namespace="runai", daemonset="ds1"}`, "1x20"},
		{`container_cpu_usage_seconds_total{pod="wl-pod", namespace="team-a", container="main"}`, "0+30x20"},
		{`container_cpu_usage_seconds_total{pod="wl-pod", namespace="team-a", container="", id="/pod.slice"}`, "0+30x20"},
		{`container_cpu_usage_seconds_total{pod="cp-pod", namespace="runai", container="cp"}`, "0+27x20"},
		{`container_memory_usage_bytes{pod="wl-pod", namespace="team-a", container="main"}`, "2000x20"},
		{`container_memory_usage_bytes{pod="wl-pod", namespace="team-a", container="", id="/pod.slice"}`, "2000x20"},
		{`container_memory_working_set_bytes{pod="cp-pod", namespace="runai", container="cp"}`, "950x20"},
		{`node_cpu_seconds_total{job="node-exporter", pod="node-exporter-x", namespace="monitoring", cpu="0", mode="idle"}`, "0+360x20"},
		{`node_memory_MemTotal_bytes{job="node-exporter", pod="node-exporter-x", namespace="monitoring"}`, "68719476736x20"},
		{`node_memory_MemAvailable_bytes{job="node-exporter", pod="node-exporter-x", namespace="monitoring"}`, "34359738368x20"},
		{`node_memory_Cached_bytes{job="node-exporter", pod="node-exporter-x", namespace="monitoring"}`, "1000x20"},
		{`node_memory_Buffers_bytes{job="node-exporter", pod="node-exporter-x", namespace="monitoring"}`, "1000x20"},
		{`node_memory_MemFree_bytes{job="node-exporter", pod="node-exporter-x", namespace="monitoring"}`, "1000x20"},
		{`node_memory_Slab_bytes{job="node-exporter", pod="node-exporter-x", namespace="monitoring"}`, "1000x20"},
		{`DCGM_FI_DEV_GPU_UTIL{instance="10.0.0.6:9400", UUID="GPU-1", gpu="0", modelName="A100", exported_pod="wl-pod", exported_namespace="team-a"}`, "75x20"},
		{`DCGM_FI_DEV_FB_USED{instance="10.0.0.6:9400", UUID="GPU-1", gpu="0", modelName="A100", exported_pod="wl-pod", exported_namespace="team-a"}`, "1024x20"},
		{`DCGM_FI_DEV_FB_FREE{instance="10.0.0.6:9400", UUID="GPU-1", gpu="0", modelName="A100", exported_pod="wl-pod", exported_namespace="team-a"}`, "1024x20"},
		{`DCGM_FI_PROF_GR_ENGINE_ACTIVE{instance="10.0.0.6:9400", UUID="GPU-1", gpu="0", modelName="A100", exported_pod="wl-pod", exported_namespace="team-a"}`, "0.7x20"},
		{`DCGM_FI_DEV_NVLINK_BANDWIDTH_TOTAL{instance="10.0.0.6:9400", UUID="GPU-1", gpu="0", modelName="A100", exported_pod="wl-pod", exported_namespace="team-a"}`, "0+100x20"},
	}

	var all []Input
	for _, m := range thirdParty {
		all = append(all, twice(m.series, m.values)...)
	}
	all = append(all, single(`runai_node_nodepool{node="n1", nodepool="np1"}`, "1x20")...)

	newTester(t).Test(TestCase{
		Interval: "30s",
		Inputs:   all,
		Tests: []Test{
			{Expr: `count(runai_node_info_enriched)`, EvalTime: "8m", Samples: []Sample{{Value: 1}}},
			{Expr: `sum(runai_cpu_cores_count_per_node)`, EvalTime: "8m", Samples: []Sample{{Value: 16}}},
			{Expr: `sum(runai_total_memory_bytes_per_node)`, EvalTime: "8m", Samples: []Sample{{Value: 68719476736}}},
			{Expr: `sum(runai_gpu_count_per_node)`, EvalTime: "8m", Samples: []Sample{{Value: 1}}},
			{Expr: `sum(runai_gpu_memory_total_mebibytes_per_gpu)`, EvalTime: "8m", Samples: []Sample{{Value: 2048}}},
			{Expr: `sum(runai_gpu_memory_used_mebibytes_per_gpu)`, EvalTime: "8m", Samples: []Sample{{Value: 1024}}},
			{Expr: `sum(runai_gpu_utilization_per_gpu)`, EvalTime: "8m", Samples: []Sample{{Value: 75}}},
			{Expr: `sum(runai_allocated_cpu_cores_per_node)`, EvalTime: "8m", Samples: []Sample{{Value: 2}}},
			{Expr: `sum(runai_allocated_memory_bytes_per_node)`, EvalTime: "8m", Samples: []Sample{{Value: 1073741824}}},
			{Expr: `sum(runai_used_cpu_cores_per_node)`, EvalTime: "8m", Samples: []Sample{{Value: 4}}},
			{Expr: `sum(runai_used_memory_bytes_per_node)`, EvalTime: "8m", Samples: []Sample{{Value: 34359738368}}},
		},
	})
}

func TestAccountingResourcesPerNodepoolWithDuplicateSeries(t *testing.T) {
	newAccountingTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`kube_node_status_capacity{node="n1", resource="cpu", unit="core"}`, "16x5"),
			twice(`kube_node_status_capacity{node="n1", resource="memory", unit="byte"}`, "68719476736x5"),
			twice(`kube_node_status_capacity{node="n1", resource="nvidia_com_gpu", unit="integer"}`, "4x5"),
			twice(`runai_node_nodepool{node="n1", nodepool="np1"}`, "1x5"),
		),
		Tests: []Test{
			{
				Expr:     `resources_per_nodepool`,
				EvalTime: "3m",
				Samples: []Sample{
					{Labels: `resources_per_nodepool{nodepool="np1", resource="cpu"}`, Value: 16},
					{Labels: `resources_per_nodepool{nodepool="np1", resource="memory"}`, Value: 68719476736},
					{Labels: `resources_per_nodepool{nodepool="np1", resource="nvidia_com_gpu"}`, Value: 4},
				},
			},
		},
	})
}

func TestContainerMemoryUsageAlertsWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`container_memory_working_set_bytes{namespace="runai", pod="cp-pod", container="cp"}`, "950x10"),
			twice(`kube_pod_container_resource_limits{namespace="runai", pod="cp-pod", container="cp", resource="memory", unit="byte", node="n1"}`, "1000x10"),
		),
		Tests: []Test{
			{
				Expr:     `((max by(container, pod, namespace) (container_memory_working_set_bytes{namespace=~"runai|runai-backend"})>0) / on(container, pod, namespace) (max by(container, pod, namespace) (kube_pod_container_resource_limits{namespace=~"runai|runai-backend", unit="byte"})>0)) > 0.8`,
				EvalTime: "5m",
				Samples:  []Sample{{Labels: `{container="cp", namespace="runai", pod="cp-pod"}`, Value: 0.95}},
			},
		},
		Alerts: []AlertTest{
			{
				AlertName: "RunaiContainerMemoryUsageWarning",
				EvalTime:  "5m",
				Expected: []ExpectAlert{{
					Labels: map[string]string{
						"alertname": "RunaiContainerMemoryUsageWarning",
						"severity":  "warning",
						"namespace": "runai",
						"pod":       "cp-pod",
						"container": "cp",
					},
					Annotations: map[string]string{
						"summary":     "Runai container is using more than 80% of its memory limit.",
						"description": "Runai container runai/cp-pod/cp is using more than 80% of its memory limit.",
					},
				}},
			},
			{
				AlertName: "RunaiContainerMemoryUsageCritical",
				EvalTime:  "5m",
				Expected: []ExpectAlert{{
					Labels: map[string]string{
						"alertname": "RunaiContainerMemoryUsageCritical",
						"severity":  "critical",
						"namespace": "runai",
						"pod":       "cp-pod",
						"container": "cp",
					},
					Annotations: map[string]string{
						"summary":     "Runai container is using more than 90% of its memory limit.",
						"description": "Runai container runai/cp-pod/cp is using more than 90% of its memory limit.",
					},
				}},
			},
		},
	})
}

func TestCpuUsageAlertWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`container_cpu_usage_seconds_total{namespace="runai", pod="cp-pod", container="cp"}`, "0+54x10"),
			twice(`kube_pod_container_resource_limits{namespace="runai", pod="cp-pod", container="cp", resource="cpu", unit="core", node="n1"}`, "1x10"),
		),
		Tests: []Test{
			{
				Expr:     `((max by(container, pod, namespace) (rate(container_cpu_usage_seconds_total{namespace=~"runai|runai-backend"}[2m]))>0) / on(container, pod, namespace) (max by(container, pod, namespace) (kube_pod_container_resource_limits{namespace=~"runai|runai-backend", unit="core"})>0)) > 0.8`,
				EvalTime: "5m",
				Samples:  []Sample{{Labels: `{container="cp", namespace="runai", pod="cp-pod"}`, Value: 0.9}},
			},
		},
		Alerts: []AlertTest{
			{
				AlertName: "RunaiCpuUsageWarning",
				EvalTime:  "5m",
				Expected: []ExpectAlert{{
					Labels: map[string]string{
						"alertname": "RunaiCpuUsageWarning",
						"severity":  "warning",
						"namespace": "runai",
						"pod":       "cp-pod",
						"container": "cp",
					},
					Annotations: map[string]string{
						"summary":     "Runai container is using more than 80% of its CPU limit.",
						"description": "Runai container runai/cp-pod/cp is using more than 80% of its CPU limit.",
					},
				}},
			},
		},
	})
}

func TestNodeInfoEnrichedWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`kube_node_info{node="n1"}`, "1x5"),
			twice(`kube_node_status_condition{node="n1", condition="Ready", status="true"}`, "1x5"),
			twice(`kube_node_spec_unschedulable{node="n1"}`, "0x5"),
			single(`runai_node_nodepool{node="n1", nodepool="np1"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `count(runai_node_info_enriched)`, EvalTime: "3m", Samples: []Sample{{Value: 1}}},
			{
				Expr:     `runai_node_info_enriched`,
				EvalTime: "3m",
				Samples:  []Sample{{Labels: `runai_node_info_enriched{node="n1", nodepool="np1", ready="true", schedulable="true"}`, Value: 1}},
			},
		},
	})
}

func TestNodeCapacityWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`kube_node_status_capacity{node="n1", resource="cpu", unit="core"}`, "16x5"),
			twice(`kube_node_status_capacity{node="n1", resource="memory", unit="byte"}`, "68719476736x5"),
			single(`runai_node_info_enriched{node="n1", nodepool="np1", modelName="A100", ready="true", schedulable="true"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `sum(runai_cpu_cores_count_per_node{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 16}}},
			{Expr: `sum(runai_total_memory_bytes_per_node{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 68719476736}}},
		},
	})
}

func TestGpuCountPerNodeFallbackWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`kube_node_info{node="n1"}`, "1x5"),
			single(`runai_node_info_enriched{node="n1", nodepool="np1", ready="true", schedulable="true"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `sum(runai_gpu_count_per_node{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 0}}},
		},
	})
}

func TestPodResourcesWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`kube_pod_status_phase{pod="wl-pod", namespace="team-a", phase="Running"}`, "1x5"),
			twice(`kube_pod_status_phase{pod="wl-pod", namespace="team-a", phase="Succeeded"}`, "0x5"),
			twice(`kube_pod_container_resource_requests{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="cpu", unit="core"}`, "2x5"),
			twice(`kube_pod_container_resource_requests{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="memory", unit="byte"}`, "1073741824x5"),
			twice(`kube_pod_container_resource_limits{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="cpu", unit="core"}`, "4x5"),
			twice(`kube_pod_container_resource_limits{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="memory", unit="byte"}`, "2147483648x5"),
			single(`runai_node_info_enriched{node="n1", nodepool="np1", ready="true", schedulable="true"}`, "1x5"),
			single(`runai_pod_info_enriched_nodepool_unique{pod_name="wl-pod", pod_namespace="team-a", pod_uuid="uid-1", node="n1", nodepool="np1", workload_id="wl-1", workload_name="wl"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `sum(runai_allocated_cpu_cores_per_node{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 2}}},
			{Expr: `sum(runai_allocated_memory_bytes_per_node{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 1073741824}}},
			{Expr: `sum(runai_cpu_requested_cores_per_pod{pod_name="wl-pod"})`, EvalTime: "3m", Samples: []Sample{{Value: 2}}},
			{Expr: `sum(runai_cpu_limit_cores_per_pod{pod_name="wl-pod"})`, EvalTime: "3m", Samples: []Sample{{Value: 4}}},
			{Expr: `sum(runai_memory_requested_bytes_per_pod{pod_name="wl-pod"})`, EvalTime: "3m", Samples: []Sample{{Value: 1073741824}}},
			{Expr: `sum(runai_memory_limit_bytes_per_pod{pod_name="wl-pod"})`, EvalTime: "3m", Samples: []Sample{{Value: 2147483648}}},
			{Expr: `sum(runai_active_job_cpu_requested_cores{workload_id="wl-1"})`, EvalTime: "3m", Samples: []Sample{{Value: 2}}},
			{Expr: `sum(runai_active_job_memory_requested_bytes{workload_id="wl-1"})`, EvalTime: "3m", Samples: []Sample{{Value: 1073741824}}},
			{Expr: `sum(runai_cpu_limits_per_active_workload{workload_id="wl-1"})`, EvalTime: "3m", Samples: []Sample{{Value: 4}}},
			{Expr: `sum(runai_memory_limits_per_active_workload{workload_id="wl-1"})`, EvalTime: "3m", Samples: []Sample{{Value: 2147483648}}},
		},
	})
}

func TestPodResourcesWithConflictingNodeLabel(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			single(`kube_pod_container_resource_requests{endpoint="monitoring-a", pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="cpu", unit="core"}`, "2x5"),
			single(`kube_pod_container_resource_requests{endpoint="monitoring-b", pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="", resource="cpu", unit="core"}`, "2x5"),
			single(`kube_pod_container_resource_requests{endpoint="monitoring-a", pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="memory", unit="byte"}`, "1073741824x5"),
			single(`kube_pod_container_resource_requests{endpoint="monitoring-b", pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="", resource="memory", unit="byte"}`, "1073741824x5"),
			twice(`kube_pod_status_phase{pod="wl-pod", namespace="team-a", phase="Running"}`, "1x5"),
			single(`runai_node_info_enriched{node="n1", nodepool="np1", ready="true", schedulable="true"}`, "1x5"),
			single(`runai_pod_info_enriched_nodepool_unique{pod_name="wl-pod", pod_namespace="team-a", pod_uuid="uid-1", node="n1", nodepool="np1", workload_id="wl-1", workload_name="wl"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `sum(runai_cpu_requested_cores_per_pod{pod_name="wl-pod"})`, EvalTime: "3m", Samples: []Sample{{Value: 2}}},
			{Expr: `sum(runai_memory_requested_bytes_per_pod{pod_name="wl-pod"})`, EvalTime: "3m", Samples: []Sample{{Value: 1073741824}}},
			{Expr: `sum(runai_active_job_cpu_requested_cores{workload_id="wl-1"})`, EvalTime: "3m", Samples: []Sample{{Value: 2}}},
			{Expr: `sum(runai_allocated_cpu_cores_per_node{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 2}}},
		},
	})
}

func TestPendingRequestsWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`kube_pod_container_resource_requests{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="cpu", unit="core"}`, "2x5"),
			twice(`kube_pod_container_resource_requests{pod="wl-pod", namespace="team-a", uid="uid-1", container="main", node="n1", resource="memory", unit="byte"}`, "1073741824x5"),
			single(`runai_pod_info_enriched_nodepool_unique{pod_name="wl-pod", pod_namespace="team-a", pod_uuid="uid-1", node="n1", nodepool="np1", workload_id="wl-1", workload_name="wl"}`, "1x5"),
			single(`runai_workload_info_enriched{workload_id="wl-1", workload_name="wl", nodepool="np1", detailed_status="Pending"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `sum(runai_pending_requested_cpu_cores_per_workload{workload_id="wl-1"})`, EvalTime: "3m", Samples: []Sample{{Value: 2}}},
			{Expr: `sum(runai_pending_requested_memory_bytes_per_workload{workload_id="wl-1"})`, EvalTime: "3m", Samples: []Sample{{Value: 1073741824}}},
		},
	})
}

func TestPodCompletionTimeWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`kube_pod_completion_time{pod="wl-pod", namespace="team-a", uid="uid-1"}`, "1700000000x5"),
			single(`runai_pod_info_enriched_nodepool_unique{pod_name="wl-pod", pod_namespace="team-a", pod_uuid="uid-1", node="n1", nodepool="np1", workload_id="wl-1"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `count(runai_pod_completion_time)`, EvalTime: "3m", Samples: []Sample{{Value: 1}}},
			{Expr: `max(runai_pod_completion_time)`, EvalTime: "3m", Samples: []Sample{{Value: 1700000000}}},
		},
	})
}

func TestPodUsageWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`container_cpu_usage_seconds_total{container="main", pod="wl-pod", namespace="team-a"}`, "0+60x10"),
			twice(`container_cpu_usage_seconds_total{container="sidecar", pod="wl-pod", namespace="team-a"}`, "0+30x10"),
			twice(`container_memory_usage_bytes{container="main", pod="wl-pod", namespace="team-a"}`, "2000x10"),
			twice(`container_memory_usage_bytes{container="sidecar", pod="wl-pod", namespace="team-a"}`, "1000x10"),
			single(`runai_pod_info_enriched_nodepool_unique{pod_name="wl-pod", pod_namespace="team-a", pod_uuid="uid-1", node="n1", nodepool="np1", workload_id="wl-1"}`, "1x10"),
		),
		Tests: []Test{
			{Expr: `sum(runai_used_cpu_cores_per_pod{pod_name="wl-pod"})`, EvalTime: "5m", Samples: []Sample{{Value: 1.5}}},
			{Expr: `sum(runai_used_memory_bytes_per_pod{pod_name="wl-pod"})`, EvalTime: "5m", Samples: []Sample{{Value: 3000}}},
		},
	})
}

func TestNodeUsageWithDuplicateNodeExporters(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "30s",
		Inputs: inputs(
			single(`node_cpu_seconds_total{job="node-exporter", mode="idle", cpu="0", pod="node-exporter-a", namespace="monitoring"}`, "0+360x20"),
			single(`node_cpu_seconds_total{job="node-exporter", mode="idle", cpu="0", pod="node-exporter-b", namespace="monitoring"}`, "0+360x20"),
			single(`node_memory_MemTotal_bytes{job="node-exporter", pod="node-exporter-a", namespace="monitoring"}`, "68719476736x20"),
			single(`node_memory_MemTotal_bytes{job="node-exporter", pod="node-exporter-b", namespace="monitoring"}`, "68719476736x20"),
			single(`node_memory_MemAvailable_bytes{job="node-exporter", pod="node-exporter-a", namespace="monitoring"}`, "34359738368x20"),
			single(`node_memory_MemAvailable_bytes{job="node-exporter", pod="node-exporter-b", namespace="monitoring"}`, "34359738368x20"),
			twice(`kube_pod_info{pod="node-exporter-a", namespace="monitoring", node="n1", uid="ne-a"}`, "1x20"),
			twice(`kube_pod_info{pod="node-exporter-b", namespace="monitoring", node="n1", uid="ne-b"}`, "1x20"),
			single(`runai_cpu_cores_count_per_node{node="n1", nodepool="np1", ready="true", schedulable="true"}`, "16x20"),
			single(`runai_node_info_enriched{node="n1", nodepool="np1", ready="true", schedulable="true"}`, "1x20"),
			single(`runai_node_nodepool{node="n1", nodepool="np1"}`, "1x20"),
			twice(`kube_node_info{node="n1"}`, "1x20"),
		),
		Tests: []Test{
			{Expr: `sum(runai_used_cpu_cores_per_node{node="n1"})`, EvalTime: "5m", Samples: []Sample{{Value: 4}}},
			{Expr: `sum(runai_used_memory_bytes_per_node{node="n1"})`, EvalTime: "5m", Samples: []Sample{{Value: 34359738368}}},
		},
	})
}

func TestRawNodeExporterRepublishWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			single(`node_cpu_seconds_total{mode="idle", cpu="0", pod="node-exporter-a", namespace="monitoring"}`, "0+60x5"),
			single(`node_memory_MemTotal_bytes{pod="node-exporter-a", namespace="monitoring"}`, "68719476736x5"),
			single(`node_memory_MemAvailable_bytes{pod="node-exporter-a", namespace="monitoring"}`, "34359738368x5"),
			twice(`kube_pod_info{pod="node-exporter-a", namespace="monitoring", node="n1", uid="ne-a"}`, "1x5"),
			twice(`kube_node_info{node="n1"}`, "1x5"),
			single(`runai_node_nodepool{node="n1", nodepool="np1"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `count(runai_node_cpu_seconds_total{node="n1", nodepool="np1"})`, EvalTime: "3m", Samples: []Sample{{Value: 1}}},
			{Expr: `max(runai_node_memory_mem_total_bytes{node="n1", nodepool="np1"})`, EvalTime: "3m", Samples: []Sample{{Value: 68719476736}}},
			{Expr: `max(runai_node_memory_mem_available_bytes{node="n1", nodepool="np1"})`, EvalTime: "3m", Samples: []Sample{{Value: 34359738368}}},
		},
	})
}

func TestDcgmMetricsWithDuplicateSeries(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			single(`DCGM_FI_DEV_GPU_UTIL{instance="10.0.0.5:9400", UUID="GPU-1", gpu="0", modelName="A100", exported_pod="wl-pod", exported_namespace="team-a"}`, "75x5"),
			twice(`kube_pod_info{pod="dcgm-exporter-xyz", namespace="runai", node="n1", pod_ip="10.0.0.5", created_by_name="dcgm-exporter"}`, "1x5"),
			twice(`kube_node_info{node="n1"}`, "1x5"),
			single(`runai_node_nodepool{node="n1", nodepool="np1"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `count(runai_dcgm_gpu_utilization{node="n1", nodepool="np1"})`, EvalTime: "3m", Samples: []Sample{{Value: 1}}},
			{Expr: `max(runai_dcgm_gpu_utilization{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 75}}},
		},
	})
}

func TestDcgmMetricsWithDuplicateExporterScrape(t *testing.T) {
	newTester(t).Test(TestCase{
		Interval: "1m",
		Inputs: inputs(
			twice(`DCGM_FI_DEV_FB_USED{instance="10.0.0.5:9400", UUID="GPU-1", gpu="0", modelName="A100"}`, "512x5"),
			twice(`DCGM_FI_DEV_FB_FREE{instance="10.0.0.5:9400", UUID="GPU-1", gpu="0", modelName="A100"}`, "512x5"),
			twice(`DCGM_FI_DEV_GPU_UTIL{instance="10.0.0.5:9400", UUID="GPU-1", gpu="0", modelName="A100"}`, "75x5"),
			twice(`kube_pod_info{pod="dcgm-exporter-xyz", namespace="runai", node="n1", pod_ip="10.0.0.5", created_by_name="dcgm-exporter"}`, "1x5"),
			twice(`kube_node_info{node="n1"}`, "1x5"),
			single(`runai_node_nodepool{node="n1", nodepool="np1"}`, "1x5"),
		),
		Tests: []Test{
			{Expr: `count(runai_dcgm_gpu_total_mebibytes{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 1}}},
			{Expr: `sum(runai_gpu_memory_total_mebibytes_per_gpu{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 1024}}},
			{Expr: `sum(runai_gpu_memory_used_mebibytes_per_gpu{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 512}}},
			{Expr: `sum(runai_gpu_utilization_per_gpu{node="n1"})`, EvalTime: "3m", Samples: []Sample{{Value: 75}}},
		},
	})
}
