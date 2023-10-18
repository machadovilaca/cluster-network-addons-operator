package rules

import (
	"fmt"
	"os"

	"github.com/machadovilaca/operator-observability/pkg/operatorrules"
	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

const (
	defaultRunbookURLTemplate = "https://kubevirt.io/monitoring/runbooks/"
	runbookURLTemplateEnv     = "RUNBOOK_URL_TEMPLATE"
)

var (
	recordingRules = [][]operatorrules.RecordingRule{
		operatorRecordingRules,
		kubemacpoolRecordingRules,
	}

	alerts = [][]promv1.Rule{
		operatorAlerts,
		kubemacpoolAlerts,
	}
)

func SetupRules() error {
	err := operatorrules.RegisterRecordingRules(recordingRules...)
	if err != nil {
		return err
	}

	runbookUrl := GetRunbookURLTemplate()

	for _, alertList := range alerts {
		for _, alert := range alertList {
			// add component labels
			alert.Labels["kubernetes_operator_part_of"] = "kubevirt"
			alert.Labels["kubernetes_operator_component"] = "cluster-network-addons-operator"

			// add runbook url
			alert.Annotations["runbook_url"] = fmt.Sprintf("%s%s", runbookUrl, alert.Alert)
		}
	}

	return operatorrules.RegisterAlerts(alerts...)
}

func BuildPrometheusRule() (*promv1.PrometheusRule, error) {
	rules, err := operatorrules.BuildPrometheusRule(
		"prometheus-rules-cluster-network-addons-operator",
		GetOperandNamespace(),
		map[string]string{"prometheus.cnao.io": "true"},
	)
	if err != nil {
		return nil, err
	}

	return rules, nil
}

func ListRecordingRules() []operatorrules.RecordingRule {
	return operatorrules.ListRecordingRules()
}

func ListAlerts() []promv1.Rule {
	return operatorrules.ListAlerts()
}

func GetRunbookURLTemplate() string {
	runbookURLTemplate, exists := os.LookupEnv(runbookURLTemplateEnv)
	if !exists {
		runbookURLTemplate = defaultRunbookURLTemplate
	}

	return runbookURLTemplate
}

func GetOperandNamespace() string {
	return os.Getenv("OPERAND_NAMESPACE")
}
