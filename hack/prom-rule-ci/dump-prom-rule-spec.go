package main

import (
	"encoding/json"
	"fmt"
	"github.com/kubevirt/cluster-network-addons-operator/pkg/monitoring/rules"
)

func main() {
	err := rules.SetupRules()
	if err != nil {
		panic(err)
	}

	promRule, err := rules.BuildPrometheusRule()
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(promRule.Spec)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
