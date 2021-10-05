# KubeVault

[KubeVault by AppsCode](https://github.com/kubevault) - HashiCorp Vault operator for Kubernetes

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubevault appscode/kubevault -n kubevault
```

## Introduction

This chart deploys a KubeVault operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubevault`:

```console
$ helm install kubevault appscode/kubevault -n kubevault
```

The command deploys a KubeVault operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubevault`:

```console
$ helm delete kubevault -n kubevault
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubevault` chart and their default values.

|                Parameter                |                                                                                                                                                                                  Description                                                                                                                                                                                   | Default |
|-----------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| global.license                          | License for the product. Get a license by following the steps from [here](https://kubevault.com/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubevault \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubevault \` <br> `--set global.license=<license file content>` | `""`    |
| global.registry                         | Docker registry used to pull KubeVault related images                                                                                                                                                                                                                                                                                                                          | `""`    |
| global.registryFQDN                     | Docker registry fqdn used to pull KubeVault related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                     | `""`    |
| global.imagePullSecrets                 | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubevault \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                           | `[]`    |
| global.skipCleaner                      | Skip generating cleaner job YAML                                                                                                                                                                                                                                                                                                                                               | `false` |
| global.monitoring.enabled               | If true, enables monitoring KubeVault operator                                                                                                                                                                                                                                                                                                                                 | `false` |
| global.monitoring.agent                 | Name of monitoring agent ("prometheus.io" or "prometheus.io/operator" or "prometheus.io/builtin")                                                                                                                                                                                                                                                                              | `""`    |
| global.monitoring.serviceMonitor.labels | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/operator`.                                                                                                                                                                                                            | `{}`    |
| kubevault-crds.enabled                  | If enabled, installs the kubevault-crds chart                                                                                                                                                                                                                                                                                                                                  | `false` |
| kubevault-catalog.enabled               | If enabled, installs the kubevault-catalog chart                                                                                                                                                                                                                                                                                                                               | `true`  |
| kubevault-operator.enabled              | If enabled, installs the kubevault-operator chart                                                                                                                                                                                                                                                                                                                              | `true`  |
| kubevault-metrics.enabled               | If enabled, installs the kubevault-metrics chart                                                                                                                                                                                                                                                                                                                               | `false` |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubevault appscode/kubevault -n kubevault --set global.registry=kubevault
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubevault appscode/kubevault -n kubevault --values values.yaml
```
