#!/usr/bin/env bash

set -xe

source hack/components/yaml-utils.sh

TMP_PATH=$(mktemp -d -p /tmp -t alerts-test-XXXXXX)
PROMETHERUS_IMAGE=quay.io/prometheus/prometheus:v2.15.2
OPERAND_NAMESPACE=test-ns

teardown() {
  rm -rf "${TMP_PATH}"
}

# get_prometheus_rule_spec extracts the spec from the prometheus rule instance
get_prometheus_rule_spec() {
  local output_prom_spec=$1
  prom_rule_spec=$(OPERAND_NAMESPACE=${OPERAND_NAMESPACE} go run ./hack/prom-rule-ci)
  echo "${prom_rule_spec}" >${output_prom_spec}
}

# test_setup prepares all the needed files in the tmp folder
test_setup() {
  local tests_yaml=$1
  local tmp_prom_spec=$2
  local tmp_tests_file=$3

  get_prometheus_rule_spec ${tmp_prom_spec}

  cp ${tests_yaml} ${tmp_tests_file}
  sed -i 's|{{ .Namespace }}|'${OPERAND_NAMESPACE}'|g' ${tmp_tests_file}
}

verify_prom_spec() {
  local tmp_prom_spec=$1

  ${OCI_BIN} run --rm --entrypoint=/bin/promtool \
  -v ${tmp_prom_spec}:/tmp/rules.verify.yaml:ro \
  ${PROMETHERUS_IMAGE} \
  check rules /tmp/rules.verify.yaml
}

check_rules() {
  local tmp_prom_spec=$1
  local tmp_tests_file=$2

  ${OCI_BIN} run --rm --entrypoint=/bin/promtool \
  -v ${tmp_prom_spec}:/tmp/rules.verify.yaml:ro \
  -v ${tmp_tests_file}:/tmp/rules.test.yaml:ro \
  ${PROMETHERUS_IMAGE} \
  test rules /tmp/rules.test.yaml
}

main() {
  local tests_file="${1:?}"
  local tmp_prom_spec=${TMP_PATH}/prom-spec
  local tmp_tests_file=${TMP_PATH}/prom-rules-tests.yaml

  trap teardown EXIT SIGINT

  test_setup ${tests_file} ${tmp_prom_spec} ${tmp_tests_file}

  verify_prom_spec ${tmp_prom_spec}
  check_rules ${tmp_prom_spec} ${tmp_tests_file}
}

main "$@"
