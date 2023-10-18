package rules

import (
	"fmt"

	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"
	"github.com/machadovilaca/operator-observability/pkg/operatorrules"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var kubemacpoolRecordingRules = []operatorrules.RecordingRule{
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "kubevirt_cnao_kubemacpool_duplicate_macs",
			Help: "Total count of duplicate KubeMacPool MAC addresses",
		},
		MetricType: operatormetrics.GaugeType,
		Expr:       intstr.FromString(fmt.Sprintf("sum(kubevirt_kmp_duplicate_macs{namespace='%s'} or vector(0))", GetOperandNamespace())),
	},
	{
		MetricsOpts: operatormetrics.MetricOpts{
			Name: "kubevirt_cnao_kubemacpool_manager_up",
			Help: "Total count of running KubeMacPool manager pods",
		},
		MetricType: operatormetrics.GaugeType,
		Expr:       intstr.FromString(fmt.Sprintf("sum(up{namespace='%s', pod=~'kubemacpool-mac-controller-manager-.*'} or vector(0))", GetOperandNamespace())),
	},
}
