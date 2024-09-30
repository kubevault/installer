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

OS=$(uname -o)
if [ "${OS}" = "GNU/Linix" ]; then
    OS=Linux
fi
ARCH=$(uname -m)
if [ "${ARCH}" = "aarch64" ]; then
    ARCH=arm64
fi
curl -sL "https://github.com/google/go-containerregistry/releases/latest/download/go-containerregistry_${OS}_${ARCH}.tar.gz" >/tmp/go-containerregistry.tar.gz
tar -zxvf /tmp/go-containerregistry.tar.gz -C /tmp/
mv /tmp/crane .

CMD="./crane"

$CMD cp ghcr.io/kubevault/secrets-store-reader:v2024.9.30 $IMAGE_REGISTRY/kubevault/secrets-store-reader:v2024.9.30
$CMD cp ghcr.io/kubevault/vault-exporter:v0.1.1 $IMAGE_REGISTRY/kubevault/vault-exporter:v0.1.1
$CMD cp ghcr.io/kubevault/vault-operator:v0.19.0 $IMAGE_REGISTRY/kubevault/vault-operator:v0.19.0
$CMD cp kubevault/vault-unsealer:v0.19.0 $IMAGE_REGISTRY/kubevault/vault-unsealer:v0.19.0
$CMD cp vault:0.11.5 $IMAGE_REGISTRY/vault:0.11.5
$CMD cp vault:1.10.3 $IMAGE_REGISTRY/vault:1.10.3
$CMD cp vault:1.11.5 $IMAGE_REGISTRY/vault:1.11.5
$CMD cp vault:1.12.1 $IMAGE_REGISTRY/vault:1.12.1
$CMD cp vault:1.13.3 $IMAGE_REGISTRY/vault:1.13.3
$CMD cp vault:1.2.0 $IMAGE_REGISTRY/vault:1.2.0
$CMD cp vault:1.2.2 $IMAGE_REGISTRY/vault:1.2.2
$CMD cp vault:1.2.3 $IMAGE_REGISTRY/vault:1.2.3
$CMD cp vault:1.5.9 $IMAGE_REGISTRY/vault:1.5.9
$CMD cp vault:1.6.5 $IMAGE_REGISTRY/vault:1.6.5
$CMD cp vault:1.7.2 $IMAGE_REGISTRY/vault:1.7.2
$CMD cp vault:1.7.3 $IMAGE_REGISTRY/vault:1.7.3
$CMD cp vault:1.8.2 $IMAGE_REGISTRY/vault:1.8.2
$CMD cp vault:1.9.2 $IMAGE_REGISTRY/vault:1.9.2
