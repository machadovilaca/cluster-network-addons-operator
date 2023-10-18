package monitoring

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/kubevirt/cluster-network-addons-operator/pkg/monitoring/rules"
	"github.com/kubevirt/cluster-network-addons-operator/pkg/render"
)

const (
	defaultServiceAccountName  = "prometheus-k8s"
	defaultMonitoringNamespace = "monitoring"
)

func RenderMonitoring(manifestDir string, monitoringAvailable bool) ([]*unstructured.Unstructured, error) {
	if !monitoringAvailable {
		return nil, nil
	}

	// render the manifests on disk
	data := render.MakeRenderData()
	data.Data["Namespace"] = rules.GetOperandNamespace()
	data.Data["MonitoringNamespace"] = GetMonitoringNamespace()
	data.Data["MonitoringServiceAccount"] = getServiceAccount()

	objs, err := render.RenderDir(filepath.Join(manifestDir, "monitoring"), &data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to render monitoring manifests")
	}

	promRule, err := rules.BuildPrometheusRule()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build PrometheusRule")
	}

	unstructuredObj, err := convertToUnstructured(promRule)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert PrometheusRule to Unstructured")
	}

	objs = append(objs, unstructuredObj)

	return objs, nil
}

func GetMonitoringNamespace() string {
	monitoringNamespaceFromEnv := os.Getenv("MONITORING_NAMESPACE")

	if monitoringNamespaceFromEnv != "" {
		return monitoringNamespaceFromEnv
	}
	return defaultMonitoringNamespace
}

func getServiceAccount() string {
	monitoringServiceAccountFromEnv := os.Getenv("MONITORING_SERVICE_ACCOUNT")

	if monitoringServiceAccountFromEnv != "" {
		return monitoringServiceAccountFromEnv
	}
	return defaultServiceAccountName
}

func convertToUnstructured(obj interface{}) (*unstructured.Unstructured, error) {
	runtimeObject, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	unstrObj := &unstructured.Unstructured{}
	unstrObj.SetUnstructuredContent(runtimeObject)

	return unstrObj, nil
}
