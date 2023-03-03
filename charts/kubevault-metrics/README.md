# KubeVault Metrics

[KubeVault Metrics](https://github.com/kubevault) - KubeVault State Metrics

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/kubevault-metrics --version=v2022.12.28
$ helm upgrade -i kubevault-metrics appscode/kubevault-metrics -n kubevault --create-namespace --version=v2022.12.28
```

## Introduction

This chart deploys KubeVault metrics configurations on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `kubevault-metrics`:

```bash
$ helm upgrade -i kubevault-metrics appscode/kubevault-metrics -n kubevault --create-namespace --version=v2022.12.28
```

The command deploys KubeVault metrics configurations on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `kubevault-metrics`:

```bash
$ helm uninstall kubevault-metrics -n kubevault
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


