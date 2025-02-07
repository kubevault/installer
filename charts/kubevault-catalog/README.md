# KubeVault Catalog

[KubeVault Catalog by AppsCode](https://github.com/kubevault/operator) - Catalog for KubeVault supported versions supported by KubeVault

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubevault-catalog --version=v2025.2.10
$ helm upgrade -i kubevault-catalog appscode/kubevault-catalog -n kubevault --create-namespace --version=v2025.2.10
```

## Introduction

This chart deploys HashiCorp KubeVault Catalog on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubevault-catalog`:

```bash
$ helm upgrade -i kubevault-catalog appscode/kubevault-catalog -n kubevault --create-namespace --version=v2025.2.10
```

The command deploys HashiCorp KubeVault Catalog on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubevault-catalog`:

```bash
$ helm uninstall kubevault-catalog -n kubevault
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubevault-catalog` chart and their default values.

|       Parameter       |                   Description                   |           Default            |
|-----------------------|-------------------------------------------------|------------------------------|
| nameOverride          | Overrides name template                         | <code>""</code>              |
| fullnameOverride      | Overrides fullname template                     | <code>""</code>              |
| proxies.dockerHub     |                                                 | <code>""</code>              |
| proxies.dockerLibrary |                                                 | <code>""</code>              |
| proxies.ghcr          |                                                 | <code>ghcr.io</code>         |
| proxies.kubernetes    |                                                 | <code>registry.k8s.io</code> |
| proxies.appscode      |                                                 | <code>r.appscode.com</code>  |
| skipDeprecated        | Set true to avoid deploying deprecated versions | <code>true</code>            |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i kubevault-catalog appscode/kubevault-catalog -n kubevault --create-namespace --version=v2025.2.10 --set proxies.ghcr=ghcr.io
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i kubevault-catalog appscode/kubevault-catalog -n kubevault --create-namespace --version=v2025.2.10 --values values.yaml
```
