# KubeVault Opscenter

[KubeVault Opscenter by AppsCode](https://github.com/kubevault) - KubeVault Opscenter

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubevault-opscenter --version=v2023.03.03
$ helm upgrade -i kubevault-opscenter appscode/kubevault-opscenter -n kubevault --create-namespace --version=v2023.03.03
```

## Introduction

This chart deploys a KubeVault Opscenter on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubevault-opscenter`:

```bash
$ helm upgrade -i kubevault-opscenter appscode/kubevault-opscenter -n kubevault --create-namespace --version=v2023.03.03
```

The command deploys a KubeVault Opscenter on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubevault-opscenter`:

```bash
$ helm uninstall kubevault-opscenter -n kubevault
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubevault-opscenter` chart and their default values.

|                Parameter                |                                                                                                                                                                                  Description                                                                                                                                                                                   |       Default        |
|-----------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| global.license                          | License for the product. Get a license by following the steps from [here](https://kubevault.com/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubevault \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubevault \` <br> `--set global.license=<license file content>` | <code>""</code>      |
| global.registry                         | Docker registry used to pull KubeVault related images                                                                                                                                                                                                                                                                                                                          | <code>""</code>      |
| global.registryFQDN                     | Docker registry fqdn used to pull KubeVault related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                     | <code>ghcr.io</code> |
| global.imagePullSecrets                 | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubevault \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                           | <code>[]</code>      |
| global.monitoring.agent                 | Name of monitoring agent (one of "prometheus.io", "prometheus.io/operator", "prometheus.io/builtin")                                                                                                                                                                                                                                                                           | <code>""</code>      |
| global.monitoring.serviceMonitor.labels | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                            | <code>{}</code>      |
| kubevault-metrics.enabled               | If enabled, installs the kubevault-metrics chart                                                                                                                                                                                                                                                                                                                               | <code>true</code>    |
| kubevault-grafana-dashboards.enabled    | If enabled, installs the kubevault-grafana-dashboards chart                                                                                                                                                                                                                                                                                                                    | <code>true</code>    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubevault-opscenter appscode/kubevault-opscenter -n kubevault --create-namespace --version=v2023.03.03 --set global.registryFQDN=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubevault-opscenter appscode/kubevault-opscenter -n kubevault --create-namespace --version=v2023.03.03 --values values.yaml
```
