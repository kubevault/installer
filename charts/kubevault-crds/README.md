# KubeVault CRDs

[KubeVault CRDs](https://github.com/kubevault) - KubeVault Custom Resource Definitions

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubevault-crds --version=v2025.2.10
$ helm upgrade -i kubevault-crds appscode/kubevault-crds -n kubevault --create-namespace --version=v2025.2.10
```

## Introduction

This chart deploys KubeVault crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubevault-crds`:

```bash
$ helm upgrade -i kubevault-crds appscode/kubevault-crds -n kubevault --create-namespace --version=v2025.2.10
```

The command deploys KubeVault crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubevault-crds`:

```bash
$ helm uninstall kubevault-crds -n kubevault
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


