package main

import (
	"fmt"
	"os"

	"github.com/machadovilaca/operator-observability/pkg/testutil"

	"github.com/kubevirt/cluster-network-addons-operator/pkg/monitoring/rules"
)

func main() {
	numErrors := 0

	err := rules.SetupRules()
	if err != nil {
		panic(err)
	}

	linter := testutil.New()

	problems := linter.LintAlerts(rules.ListAlerts())
	for _, problem := range problems {
		numErrors++
		fmt.Printf("Alert '%s' - %s\n", problem.ResourceName, problem.Description)
	}

	problems = linter.LintRecordingRules(rules.ListRecordingRules())
	for _, problem := range problems {
		numErrors++
		fmt.Printf("Recording Rule '%s' - %s\n", problem.ResourceName, problem.Description)
	}

	if numErrors > 0 {
		fmt.Printf("%d validations failed!\n", numErrors)
		os.Exit(1)
	}

	fmt.Println("All validations passed for alerts and recording rules")
}
