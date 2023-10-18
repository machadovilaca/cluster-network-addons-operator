/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2023 Red Hat, Inc.
 *
 */

package main

import (
	"github.com/kubevirt/cluster-network-addons-operator/pkg/monitoring/metrics"
	"github.com/kubevirt/cluster-network-addons-operator/pkg/monitoring/rules"
	"github.com/machadovilaca/operator-observability/pkg/operatormetrics"
	dto "github.com/prometheus/client_model/go"
)

// This should be used only for very rare cases where the naming conventions that are explained in the best practices:
// https://sdk.operatorframework.io/docs/best-practices/observability-best-practices/#metrics-guidelines
// should be ignored.
var excludedMetrics = map[string]struct{}{}

// ReadMetrics reads the metrics and parse them to a MetricFamily
func ReadMetrics() []*dto.MetricFamily {
	err := metrics.SetupMetrics()
	if err != nil {
		panic(err)
	}

	err = rules.SetupRules()
	if err != nil {
		panic(err)
	}

	var metricFamily []*dto.MetricFamily

	for _, metric := range metrics.ListMetrics() {
		if _, ok := excludedMetrics[metric.GetOpts().Name]; ok {
			continue
		}
		metricFamily = append(metricFamily, toMetricFamily(metric.GetOpts(), metric.GetType()))
	}

	for _, rule := range rules.ListRecordingRules() {
		if _, ok := excludedMetrics[rule.GetOpts().Name]; ok {
			continue
		}
		metricFamily = append(metricFamily, toMetricFamily(rule.GetOpts(), rule.GetType()))
	}

	return metricFamily
}

func toMetricFamily(opts operatormetrics.MetricOpts, mType operatormetrics.MetricType) *dto.MetricFamily {
	metricType := dto.MetricType_UNTYPED

	switch {
	case mType == operatormetrics.CounterType || mType == operatormetrics.CounterVecType:
		metricType = dto.MetricType_COUNTER
	case mType == operatormetrics.GaugeType || mType == operatormetrics.GaugeVecType:
		metricType = dto.MetricType_GAUGE
	case mType == operatormetrics.HistogramType || mType == operatormetrics.HistogramVecType:
		metricType = dto.MetricType_HISTOGRAM
	case mType == operatormetrics.SummaryType || mType == operatormetrics.SummaryVecType:
		metricType = dto.MetricType_SUMMARY
	}

	return &dto.MetricFamily{
		Name: &opts.Name,
		Help: &opts.Help,
		Type: &metricType,
	}
}
