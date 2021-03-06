# KubeVault CRDs

[KubeVault CRDs](https://github.com/kubevault) - KubeVault Custom Resource Definitions

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubevault-crds appscode/kubevault-crds -n kubevault
```

## Introduction

This chart deploys KubeVault crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubevault-crds`:

```console
$ helm install kubevault-crds appscode/kubevault-crds -n kubevault
```

The command deploys KubeVault crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubevault-crds`:

```console
$ helm delete kubevault-crds -n kubevault
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


