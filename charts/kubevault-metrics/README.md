# KubeVault Metrics

[KubeVault Metrics](https://github.com/kubevault) - KubeVault Prometheus Metrics

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubevault-metrics appscode/kubevault-metrics -n kube-system
```

## Introduction

This chart deploys KubeVault metrics on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubevault-metrics`:

```console
$ helm install kubevault-metrics appscode/kubevault-metrics -n kube-system
```

The command deploys KubeVault metrics on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubevault-metrics`:

```console
$ helm delete kubevault-metrics -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


