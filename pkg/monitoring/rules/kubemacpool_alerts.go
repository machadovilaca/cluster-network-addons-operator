package rules

import (
	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var kubemacpoolAlerts = []promv1.Rule{
	{
		Alert: "KubeMacPoolDuplicateMacsFound",
		Expr:  intstr.FromString("kubevirt_cnao_kubemacpool_duplicate_macs != 0"),
		For:   "5m",
		Annotations: map[string]string{
			"summary":     "Duplicate macs found.",
			"description": "There are {{ $value }} duplicate KubeMacPool MAC addresses.",
		},
		Labels: map[string]string{
			"severity":               "warning",
			"operator_health_impact": "warning",
		},
	},
	{
		Alert: "KubemacpoolDown",
		Expr:  intstr.FromString("kubevirt_cnao_cr_kubemacpool_aggregated == 1 and kubevirt_cnao_kubemacpool_manager_up == 0"),
		For:   "5m",
		Annotations: map[string]string{
			"summary":     "KubeMacpool is deployed by CNAO CR but KubeMacpool pod is down.",
			"description": "The total count of running KubeMacPool manager pods is 0 (zero) for more than 5 minutes.",
		},
		Labels: map[string]string{
			"severity":               "critical",
			"operator_health_impact": "critical",
		},
	},
}
