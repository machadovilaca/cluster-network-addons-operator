package main

import (
	"fmt"
	"github.com/kubevirt/cluster-network-addons-operator/pkg/monitoring/rules"

	"github.com/machadovilaca/operator-observability/pkg/docs"

	"github.com/kubevirt/cluster-network-addons-operator/pkg/monitoring/metrics"
)

const tpl = `# Cluster Network Addons Operator Metrics

{{- range . }}

{{ $deprecatedVersion := "" -}}
{{- with index .ExtraFields "DeprecatedVersion" -}}
    {{- $deprecatedVersion = printf " in %s" . -}}
{{- end -}}

{{- $stabilityLevel := "" -}}
{{- if and (.ExtraFields.StabilityLevel) (ne .ExtraFields.StabilityLevel "STABLE") -}}
	{{- $stabilityLevel = printf "[%s%s] " .ExtraFields.StabilityLevel $deprecatedVersion -}}
{{- end -}}

### {{ .Name }}
{{ print $stabilityLevel }}{{ .Help }}. Type: {{ .Type -}}.

{{- end }}

## Developing new metrics

All metrics documented here are auto-generated and reflect exactly what is being
exposed. After developing new metrics or changing old ones please regenerate
this document.
`

func main() {
	mustSucceed(metrics.SetupMetrics())
	mustSucceed(rules.SetupRules())

	docsString := docs.BuildMetricsDocsWithCustomTemplate(metrics.ListMetrics(), rules.ListRecordingRules(), tpl)
	fmt.Print(docsString)
}

func mustSucceed(err error) {
	if err != nil {
		panic(err)
	}
}
