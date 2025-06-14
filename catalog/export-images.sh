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

mkdir -p images

OS=$(uname -o)
if [ "${OS}" = "GNU/Linux" ]; then
    OS=Linux
fi
ARCH=$(uname -m)
if [ "${ARCH}" = "aarch64" ]; then
    ARCH=arm64
fi
curl -sL "https://github.com/google/go-containerregistry/releases/latest/download/go-containerregistry_${OS}_${ARCH}.tar.gz" >/tmp/go-containerregistry.tar.gz
tar -zxvf /tmp/go-containerregistry.tar.gz -C /tmp/
mv /tmp/crane images

CMD="./images/crane"

$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.10.3 images/appscode-images-vault-1.10.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.11.5 images/appscode-images-vault-1.11.5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.12.1 images/appscode-images-vault-1.12.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.13.3 images/appscode-images-vault-1.13.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.14.10 images/appscode-images-vault-1.14.10.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.15.6 images/appscode-images-vault-1.15.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.16.3 images/appscode-images-vault-1.16.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.17.6 images/appscode-images-vault-1.17.6.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.18.4 images/appscode-images-vault-1.18.4.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.2.0 images/appscode-images-vault-1.2.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.2.2 images/appscode-images-vault-1.2.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.2.3 images/appscode-images-vault-1.2.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.5.9 images/appscode-images-vault-1.5.9.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.6.5 images/appscode-images-vault-1.6.5.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.7.2 images/appscode-images-vault-1.7.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.7.3 images/appscode-images-vault-1.7.3.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.8.2 images/appscode-images-vault-1.8.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/appscode-images/vault:1.9.2 images/appscode-images-vault-1.9.2.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubevault/vault-exporter:v0.1.1 images/kubevault-vault-exporter-v0.1.1.tar
$CMD pull --allow-nondistributable-artifacts --insecure ghcr.io/kubevault/vault-operator:v0.22.0 images/kubevault-vault-operator-v0.22.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure kubevault/vault-unsealer:v0.22.0 images/kubevault-vault-unsealer-v0.22.0.tar
$CMD pull --allow-nondistributable-artifacts --insecure vault:0.11.5 images/library-vault-0.11.5.tar

tar -czvf images.tar.gz images
