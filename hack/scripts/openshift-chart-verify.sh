#!/bin/bash

# Copyright AppsCode Inc. and Contributors
#
# Licensed under the AppsCode Community License 1.0.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Packages charts/kubevault-certified, drops it into a local checkout of
# appscode/openshift-helm-charts, and runs redhat-certification/chart-verifier
# against it to produce a certification report.
#
# Usage: openshift-chart-verify.sh <kubevault-license-path> [version]
#   kubevault-license-path  path to a KubeVault Enterprise license file;
#                           passed to chart-verifier as the global.license
#                           chart value.
#   version                 chart version to package, with or without a "v"
#                           prefix (default: version field in
#                           charts/kubevault-certified/Chart.yaml).
#
# Requires a reachable Kubernetes cluster on the current kubeconfig context
# (e.g. `kind create cluster --image kindest/node:v1.34.0`); this script does
# not create one.
#
# Env overrides:
#   OPENSHIFT_HELM_CHARTS_DIR  checkout of appscode/openshift-helm-charts
#                              (default: ~/go/src/github.com/appscode/openshift-helm-charts)

set -eou pipefail

if [ $# -lt 1 ]; then
    echo "Usage: $0 <kubevault-license-path> [version]" >&2
    exit 1
fi

LICENSE_PATH="$1"
if [ ! -f "$LICENSE_PATH" ]; then
    echo "license file not found: $LICENSE_PATH" >&2
    exit 1
fi

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
OPENSHIFT_HELM_CHARTS_DIR="${OPENSHIFT_HELM_CHARTS_DIR:-$HOME/go/src/github.com/appscode/openshift-helm-charts}"

CHART_VERSION="${2:-$(grep -E '^version:' "$REPO_ROOT/charts/kubevault-certified/Chart.yaml" | awk '{print $2}')}"
CHART_VERSION="${CHART_VERSION#v}"

echo "==> Installing chart-verifier 1.14.1"
go install github.com/redhat-certification/chart-verifier@1.14.1
CHART_VERIFIER="$(go env GOPATH)/bin/chart-verifier"

echo "==> Checking for a reachable Kubernetes cluster"
if ! kubectl cluster-info >/dev/null 2>&1; then
    echo "no reachable Kubernetes cluster on the current kubeconfig context." >&2
    echo "create one first, e.g.: kind create cluster --image kindest/node:v1.34.0" >&2
    exit 1
fi

PACKAGE_DIR="$(mktemp -d)"
trap 'rm -rf "$PACKAGE_DIR"' EXIT
helm package "$REPO_ROOT/charts/kubevault-certified" --destination "$PACKAGE_DIR"

CHART_DEST_DIR="$OPENSHIFT_HELM_CHARTS_DIR/charts/partners/appscode/kubevault-certified/$CHART_VERSION"
mkdir -p "$CHART_DEST_DIR"

CHART_PACKAGE="$(find "$PACKAGE_DIR" -maxdepth 1 -name 'kubevault-certified-*.tgz')"
CHART_TGZ="$CHART_DEST_DIR/kubevault-certified-$CHART_VERSION.tgz"
cp "$CHART_PACKAGE" "$CHART_TGZ"

echo "==> Installing kubevault-certified-crds CRDs into the cluster"
helm upgrade --install kubevault-certified-crds "$REPO_ROOT/charts/kubevault-certified-crds"

echo "==> Running chart-verifier"
"$CHART_VERIFIER" verify \
    -S global.distro.ubi=operator \
    -G global.license="$LICENSE_PATH" \
    "$CHART_TGZ" \
    >"$CHART_DEST_DIR/report.yaml"

echo "==> Done: $CHART_DEST_DIR"
