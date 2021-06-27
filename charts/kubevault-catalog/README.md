# KubeVault Catalog

[KubeVault Catalog by AppsCode](https://github.com/kubevault/operator) - Catalog for KubeVault supported versions supported by KubeVault

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubevault-catalog appscode/kubevault-catalog -n kubevault
```

## Introduction

This chart deploys HashiCorp KubeVault Catalog on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubevault-catalog`:

```console
$ helm install kubevault-catalog appscode/kubevault-catalog -n kubevault
```

The command deploys HashiCorp KubeVault Catalog on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubevault-catalog`:

```console
$ helm delete kubevault-catalog -n kubevault
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubevault-catalog` chart and their default values.

|    Parameter     |                   Description                   |   Default   |
|------------------|-------------------------------------------------|-------------|
| nameOverride     | Overrides name template                         | `""`        |
| fullnameOverride | Overrides fullname template                     | `""`        |
| image.registry   | Docker registry used to pull Vault server image | `kubevault` |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubevault-catalog appscode/kubevault-catalog -n kubevault --set image.registry=kubevault
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubevault-catalog appscode/kubevault-catalog -n kubevault --values values.yaml
```
