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

CMD="./crane"

$CMD push --allow-nondistributable-artifacts --insecure images/kubevault-vault-exporter-v0.1.1.tar $IMAGE_REGISTRY/kubevault/vault-exporter:v0.1.1
$CMD push --allow-nondistributable-artifacts --insecure images/kubevault-vault-operator-v0.20.0-rc.0.tar $IMAGE_REGISTRY/kubevault/vault-operator:v0.20.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/kubevault-vault-unsealer-v0.20.0-rc.0.tar $IMAGE_REGISTRY/kubevault/vault-unsealer:v0.20.0-rc.0
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-0.11.5.tar $IMAGE_REGISTRY/vault:0.11.5
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.10.3.tar $IMAGE_REGISTRY/vault:1.10.3
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.11.5.tar $IMAGE_REGISTRY/vault:1.11.5
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.12.1.tar $IMAGE_REGISTRY/vault:1.12.1
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.13.3.tar $IMAGE_REGISTRY/vault:1.13.3
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.2.0.tar $IMAGE_REGISTRY/vault:1.2.0
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.2.2.tar $IMAGE_REGISTRY/vault:1.2.2
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.2.3.tar $IMAGE_REGISTRY/vault:1.2.3
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.5.9.tar $IMAGE_REGISTRY/vault:1.5.9
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.6.5.tar $IMAGE_REGISTRY/vault:1.6.5
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.7.2.tar $IMAGE_REGISTRY/vault:1.7.2
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.7.3.tar $IMAGE_REGISTRY/vault:1.7.3
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.8.2.tar $IMAGE_REGISTRY/vault:1.8.2
$CMD push --allow-nondistributable-artifacts --insecure images/library-vault-1.9.2.tar $IMAGE_REGISTRY/vault:1.9.2
