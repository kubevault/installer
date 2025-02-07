# KubeVault

[KubeVault by AppsCode](https://github.com/kubevault) - HashiCorp Vault operator for Kubernetes

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubevault --version=v2025.2.10
$ helm upgrade -i kubevault appscode/kubevault -n kubevault --create-namespace --version=v2025.2.10
```

## Introduction

This chart deploys a KubeVault operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubevault`:

```bash
$ helm upgrade -i kubevault appscode/kubevault -n kubevault --create-namespace --version=v2025.2.10
```

The command deploys a KubeVault operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubevault`:

```bash
$ helm uninstall kubevault -n kubevault
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubevault` chart and their default values.

|                       Parameter                       |                                                                                                                                                                                  Description                                                                                                                                                                                   |      Default       |
|-------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------|
| global.license                                        | License for the product. Get a license by following the steps from [here](https://kubevault.com/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubevault \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubevault \` <br> `--set global.license=<license file content>` | <code>""</code>    |
| global.licenseSecretName                              | Name of Secret with the license as key.txt key                                                                                                                                                                                                                                                                                                                                 | <code>""</code>    |
| global.registry                                       | Docker registry used to pull KubeVault related images                                                                                                                                                                                                                                                                                                                          | <code>""</code>    |
| global.registryFQDN                                   | Docker registry fqdn used to pull KubeVault related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                     | <code>""</code>    |
| global.imagePullSecrets                               | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubevault \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                           | <code>[]</code>    |
| global.monitoring.agent                               | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                                                                                                                                                           | <code>""</code>    |
| global.monitoring.serviceMonitor.labels               | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                            | <code>{}</code>    |
| kubevault-catalog.enabled                             | If enabled, installs the kubevault-catalog chart                                                                                                                                                                                                                                                                                                                               | <code>true</code>  |
| kubevault-operator.enabled                            | If enabled, installs the kubevault-operator chart                                                                                                                                                                                                                                                                                                                              | <code>true</code>  |
| kubevault-webhook-server.enabled                      | If enabled, installs the kubevault-webhook-server chart                                                                                                                                                                                                                                                                                                                        | <code>true</code>  |
| ace-user-roles.enabled                                | If enabled, installs the ace-user-roles chart                                                                                                                                                                                                                                                                                                                                  | <code>true</code>  |
| ace-user-roles.enableClusterRoles.ace                 |                                                                                                                                                                                                                                                                                                                                                                                | <code>false</code> |
| ace-user-roles.enableClusterRoles.appcatalog          |                                                                                                                                                                                                                                                                                                                                                                                | <code>true</code>  |
| ace-user-roles.enableClusterRoles.catalog             |                                                                                                                                                                                                                                                                                                                                                                                | <code>false</code> |
| ace-user-roles.enableClusterRoles.cert-manager        |                                                                                                                                                                                                                                                                                                                                                                                | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubedb              |                                                                                                                                                                                                                                                                                                                                                                                | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubedb-ui           |                                                                                                                                                                                                                                                                                                                                                                                | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubestash           |                                                                                                                                                                                                                                                                                                                                                                                | <code>false</code> |
| ace-user-roles.enableClusterRoles.kubevault           |                                                                                                                                                                                                                                                                                                                                                                                | <code>true</code>  |
| ace-user-roles.enableClusterRoles.license-proxyserver |                                                                                                                                                                                                                                                                                                                                                                                | <code>true</code>  |
| ace-user-roles.enableClusterRoles.metrics             |                                                                                                                                                                                                                                                                                                                                                                                | <code>false</code> |
| ace-user-roles.enableClusterRoles.prometheus          |                                                                                                                                                                                                                                                                                                                                                                                | <code>false</code> |
| ace-user-roles.enableClusterRoles.stash               |                                                                                                                                                                                                                                                                                                                                                                                | <code>false</code> |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubevault appscode/kubevault -n kubevault --create-namespace --version=v2025.2.10 --set global.registry=kubevault
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubevault appscode/kubevault -n kubevault --create-namespace --version=v2025.2.10 --values values.yaml
```
