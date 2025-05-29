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

set -eou pipefail

crd_dir=${1:-}

api_repo_url=https://github.com/kubevault/apimachinery.git
api_repo_tag=${KUBEVAULT_APIMACHINERY_TAG:-master}

if [ "$#" -ne 1 ]; then
    if [ "${api_repo_tag}" == "master" ]; then
        echo "Error: missing path_to_input_crds_directory"
        echo "Usage: import-crds.sh <path_to_input_crds_directory>"
        exit 1
    fi

    tmp_dir=$(mktemp -d -t api-XXXXXXXXXX)
    # always cleanup temp dir
    # ref: https://opensource.com/article/20/6/bash-trap
    trap \
        "{ rm -rf "${tmp_dir}"; }" \
        SIGINT SIGTERM ERR EXIT

    mkdir -p ${tmp_dir}
    pushd $tmp_dir
    git clone $api_repo_url
    repo_dir=$(ls -b1)
    cd $repo_dir
    git checkout $api_repo_tag
    popd
    crd_dir=${tmp_dir}/${repo_dir}/crds
fi

KMODULES_CUSTOM_RESOURCES_TAG=${KMODULES_CUSTOM_RESOURCES_TAG:-v0.32.0}
KUBEOPS_SUPERVISOR_TAG=${KUBEOPS_SUPERVISOR_TAG:-v0.0.9}
OPEN_VIZ_APIMACHINERY_TAG=${OPEN_VIZ_APIMACHINERY_TAG:-v0.0.8}

crd-importer \
    --no-description \
    --input=${crd_dir} \
    --input=https://github.com/kmodules/custom-resources/raw/${KMODULES_CUSTOM_RESOURCES_TAG}/crds/metrics.appscode.com_metricsconfigurations.v1.yaml \
    --out=./charts/kubevault-crds/crds

crd-importer --v=v1 \
    --no-description \
    --annotations 'config.kubernetes.io/local-config=true' \
    --input=${crd_dir} \
    --out=./charts/kubevault-crds/crds \
    --group=catalog.kubevault.com

crd-importer --v=v1 \
    --no-description \
    --annotations 'config.kubernetes.io/local-config=true' \
    --input=${crd_dir} \
    --out=./charts/kubevault-crds/crds \
    --group=kubevault.com

crd-importer \
    --no-description \
    --input=${crd_dir} \
    --out=. --output-yaml=./crds/kubevault-crds.yaml

crd-importer --v=v1 \
    --no-description \
    --input=${crd_dir} \
    --out=./charts/kubevault-catalog/crds \
    --group=catalog.kubevault.com

crd-importer --v=v1 \
    --no-description \
    --input=${crd_dir} \
    --out=. --output-yaml=./crds/kubevault-catalog-crds.yaml \
    --group=catalog.kubevault.com

crd-importer --v=v1 \
    --no-description \
    --input=${crd_dir} \
    --out=./charts/kubevault-operator/crds \
    --group=kubevault.com

crd-importer \
    --no-description \
    --input=https://github.com/kmodules/custom-resources/raw/${KMODULES_CUSTOM_RESOURCES_TAG}/crds/metrics.appscode.com_metricsconfigurations.yaml \
    --out=./charts/kubevault-metrics/crds

crd-importer \
    --no-description \
    --input=https://github.com/kubeops/supervisor/raw/${KUBEOPS_SUPERVISOR_TAG}/crds/supervisor.appscode.com_recommendations.yaml \
    --out=./charts/kubevault-operator/crds

crd-importer \
    --no-description \
    --input=https://github.com/open-viz/apimachinery/raw/${OPEN_VIZ_APIMACHINERY_TAG}/crds/openviz.dev_grafanadashboards.yaml \
    --out=./charts/kubevault-grafana-dashboards/crds
