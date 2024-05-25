#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

source "${CODEGEN_PKG}/kube_codegen.sh"

THIS_PKG="k8s-controller-sample"

#kube::codegen::gen_helpers \
#    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
#    "${SCRIPT_ROOT}/api"

kube::codegen::gen_client \
    --with-watch \
    --output-dir "${SCRIPT_ROOT}/api/generated" \
    --output-pkg "${THIS_PKG}/api/generated" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}/api"