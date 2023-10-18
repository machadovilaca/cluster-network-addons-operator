package rules

import (
	"fmt"

	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"
	"github.com/machadovilaca/operator-observability/pkg/operatorrules"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var operatorRecordingRules = []operatorrules.RecordingRule{
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "kubevirt_cnao_operator_up",
			Help: "Total count of running CNAO operators",
		},
		MetricType: operatormetrics.GaugeType,
		Expr:       intstr.FromString(fmt.Sprintf("sum(up{namespace='%s', pod=~'cluster-network-addons-operator-.*'} or vector(0))", GetOperandNamespace())),
	},
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "kubevirt_cnao_cr_kubemacpool_aggregated",
			Help: "Total count of KubeMacPool manager pods deployed by CNAO CR",
		},
		MetricType: operatormetrics.GaugeType,
		Expr:       intstr.FromString(fmt.Sprintf("sum(kubevirt_cnao_cr_kubemacpool_deployed{namespace='%s'} or vector(0))", GetOperandNamespace())),
	},
}
