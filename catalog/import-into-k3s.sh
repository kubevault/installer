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

set -x

if [ -z "${IMAGE_REGISTRY}" ]; then
    echo "IMAGE_REGISTRY is not set"
    exit 1
fi

TARBALL=${1:-}
tar -zxvf $TARBALL

k3s ctr images import images/appscode-images-vault-1.10.3.tar
k3s ctr images import images/appscode-images-vault-1.11.5.tar
k3s ctr images import images/appscode-images-vault-1.12.1.tar
k3s ctr images import images/appscode-images-vault-1.13.3.tar
k3s ctr images import images/appscode-images-vault-1.14.10.tar
k3s ctr images import images/appscode-images-vault-1.15.6.tar
k3s ctr images import images/appscode-images-vault-1.16.3.tar
k3s ctr images import images/appscode-images-vault-1.17.6.tar
k3s ctr images import images/appscode-images-vault-1.18.4.tar
k3s ctr images import images/appscode-images-vault-1.2.0.tar
k3s ctr images import images/appscode-images-vault-1.2.2.tar
k3s ctr images import images/appscode-images-vault-1.2.3.tar
k3s ctr images import images/appscode-images-vault-1.5.9.tar
k3s ctr images import images/appscode-images-vault-1.6.5.tar
k3s ctr images import images/appscode-images-vault-1.7.2.tar
k3s ctr images import images/appscode-images-vault-1.7.3.tar
k3s ctr images import images/appscode-images-vault-1.8.2.tar
k3s ctr images import images/appscode-images-vault-1.9.2.tar
k3s ctr images import images/appscode-kubectl-nonroot-1.31.tar
k3s ctr images import images/kubevault-vault-exporter-v0.1.1.tar
k3s ctr images import images/kubevault-vault-operator-v0.24.0-rc.0.tar
k3s ctr images import images/kubevault-vault-unsealer-v0.24.0-rc.0.tar
k3s ctr images import images/openbao-openbao-2.4.3.tar
k3s ctr images import images/library-vault-0.11.5.tar
