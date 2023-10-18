package rules

import (
	"fmt"

	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var operatorAlerts = []promv1.Rule{
	{
		Alert: "CnaoDown",
		Expr:  intstr.FromString("kubevirt_cnao_operator_up == 0"),
		For:   "5m",
		Annotations: map[string]string{
			"summary":     "CNAO pod is down.",
			"description": "The total count of running CNAO operators is 0 (zero) for more than 5 minutes.",
		},
		Labels: map[string]string{
			"severity":               "warning",
			"operator_health_impact": "warning",
		},
	},
	{
		Alert: "NetworkAddonsConfigNotReady",
		Expr:  intstr.FromString(fmt.Sprintf("sum(kubevirt_cnao_cr_ready{namespace='%s'} or vector(0)) == 0", GetOperandNamespace())),
		For:   "5m",
		Annotations: map[string]string{
			"summary":     "CNAO CR NetworkAddonsConfig is not ready.",
			"description": "There is no CNAO CR ready for more than 5 minutes.",
		},
		Labels: map[string]string{
			"severity":               "warning",
			"operator_health_impact": "warning",
		},
	},
}
