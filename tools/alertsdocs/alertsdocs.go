package main

import (
	"fmt"

	"github.com/machadovilaca/operator-observability/pkg/docs"

	"github.com/kubevirt/cluster-network-addons-operator/pkg/monitoring/rules"
)

const tpl = `# Cluster Network Addons Operator Alerts

{{- range . }}

### {{.Name}}
**Summary:** {{ index .Annotations "summary" }}.

**Description:** {{ index .Annotations "description" }}.

**Severity:** {{ index .Labels "severity" }}.

**Operator health impact:** {{ index .Labels "operator_health_impact" }}.
{{- if .For }}

**For:** {{ .For }}.
{{- end -}}
{{- end }}

## Developing new alerts

All alerts documented here are auto-generated and reflect exactly what is being
exposed. After developing new alerts or changing old ones please regenerate
this document.
`

func main() {
	err := rules.SetupRules()
	if err != nil {
		panic(err)
	}

	docsString := docs.BuildAlertsDocsWithCustomTemplate(rules.ListAlerts(), tpl)
	fmt.Print(docsString)
}
