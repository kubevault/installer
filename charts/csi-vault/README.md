# CSI Vault

[CSI Vault by AppsCode](https://github.com/kubevault/csi-driver) - Kubernetes CSI Driver for HashiCorp Vault

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install csi-vault appscode/csi-vault -n kubevault
```

## Introduction

This chart deploys a Vault CSI Driver on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes v1.16+
- `--allow-privileged` flag must be set to true for both the API server and the kubelet
- (If you use Docker) The Docker daemon of the cluster nodes must allow shared mounts
- Pre-installed HashiCorp Vault server.

## Installing the Chart

To install the chart with the release name `csi-vault`:

```console
$ helm install csi-vault appscode/csi-vault -n kubevault
```

The command deploys a Vault CSI Driver on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `csi-vault`:

```console
$ helm delete csi-vault -n kubevault
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `csi-vault` chart and their default values.

|              Parameter               |                                                      Description                                                       |                                          Default                                           |
|--------------------------------------|------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------|
| nameOverride                         | Overrides name template                                                                                                | `""`                                                                                       |
| fullnameOverride                     | Overrides fullname template                                                                                            | `""`                                                                                       |
| attacher.name                        | Name of the attacher container                                                                                         | `attacher`                                                                                 |
| attacher.registry                    | Docker registry used to pull CSI attacher image                                                                        | `k8s.gcr.io/sig-storage`                                                                   |
| attacher.repository                  | CSI attacher container image                                                                                           | `csi-attacher`                                                                             |
| attacher.tag                         | CSI attacher container image tag                                                                                       | `v3.1.0`                                                                                   |
| attacher.pullPolicy                  | CSI attacher container image pull policy                                                                               | `IfNotPresent`                                                                             |
| plugin.name                          | Name of the CSI driver container                                                                                       | `plugin`                                                                                   |
| plugin.registry                      | Docker registry used to pull CSI driver image                                                                          | `kubevault`                                                                                |
| plugin.repository                    | CSI driver container image                                                                                             | `csi-vault`                                                                                |
| plugin.tag                           | CSI driver container image tag                                                                                         | `v0.4.0-rc.0`                                                                              |
| plugin.pullPolicy                    | CSI driver container image pull policy                                                                                 | `IfNotPresent`                                                                             |
| plugin.securityContext               | Security options the CSI driver container should run with                                                              | `{"allowPrivilegeEscalation":true,"capabilities":{"add":["SYS_ADMIN"]},"privileged":true}` |
| provisioner.name                     | Name of the provisioner container                                                                                      | `provisioner`                                                                              |
| provisioner.registry                 | Docker registry used to pull CSI provisioner image                                                                     | `k8s.gcr.io/sig-storage`                                                                   |
| provisioner.repository               | CSI provisioner container image                                                                                        | `csi-provisioner`                                                                          |
| provisioner.tag                      | CSI provisioner container image tag                                                                                    | `v2.2.0`                                                                                   |
| provisioner.pullPolicy               | CSI provisioner container image pull policy                                                                            | `IfNotPresent`                                                                             |
| nodeRegistrar.name                   | Name of the CSI driver node registrar container                                                                        | `node-registrar`                                                                           |
| nodeRegistrar.registry               | Docker registry used to pull CSI driver node registrar image                                                           | `k8s.gcr.io/sig-storage`                                                                   |
| nodeRegistrar.repository             | CSI driver node registrar container image                                                                              | `csi-node-driver-registrar`                                                                |
| nodeRegistrar.tag                    | CSI driver node registrar container image tag                                                                          | `v2.1.0`                                                                                   |
| nodeRegistrar.pullPolicy             | CSI driver node registrar container image pull policy                                                                  | `IfNotPresent`                                                                             |
| driverName                           | Vault CSI driver name                                                                                                  | `secrets.csi.kubevault.com`                                                                |
| pluginAddress                        | Vault CSI driver endpoint address                                                                                      | `/csi/csi.sock`                                                                            |
| pluginDir                            | Vault CSI driver plugin directory                                                                                      | `/csi`                                                                                     |
| components.controller                | Name of controller driver component                                                                                    | `controller`                                                                               |
| components.node                      | Name of node driver component                                                                                          | `node`                                                                                     |
| criticalAddon                        | If true, installs CSI driver as critical addon                                                                         | `false`                                                                                    |
| logLevel                             | Log level for operator                                                                                                 | `3`                                                                                        |
| nodeSelector                         | Node labels for pod assignment                                                                                         | `{"beta.kubernetes.io/arch":"amd64","beta.kubernetes.io/os":"linux"}`                      |
| rbac.create                          | Specifies whether RBAC resources should be created                                                                     | `true`                                                                                     |
| apiserver.useKubeapiserverFqdnForAks | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true) | `true`                                                                                     |
| apiserver.healthcheck.enabled        | healthcheck configures the readiness and liveliness probes for the CSI driver pod.                                     | `true`                                                                                     |
| enableAnalytics                      | If true, sends usage analytics                                                                                         | `true`                                                                                     |
| monitoring.agent                     | Name of monitoring agent (either "prometheus.io/operator" or "prometheus.io/builtin")                                  | `"none"`                                                                                   |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install csi-vault appscode/csi-vault -n kubevault --set attacher.name=attacher
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install csi-vault appscode/csi-vault -n kubevault --values values.yaml
```
