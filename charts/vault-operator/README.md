# Vault Operator

[Vault Operator by AppsCode](https://github.com/kubevault/operator) - HashiCorp Vault Operator for Kubernetes

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install vault-operator appscode/vault-operator -n kube-system
```

## Introduction

This chart bootstraps a [HashiCorp Vault controller](https://github.com/kubevault/operator) deployment on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the chart with the release name `vault-operator`:

```console
$ helm install vault-operator appscode/vault-operator -n kube-system
```

The command deploys Vault operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `vault-operator`:

```console
$ helm uninstall vault-operator -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the Vault chart and their default values.

| Parameter                               | Description                                                                                                                                                                                                                                                                                                                                                 | Default                                                   |
| --------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| `replicaCount`                          | Number of Vault operator replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                           | `1`                                                       |
| `operator.registry`                     | Docker registry used to pull Vault operator image                                                                                                                                                                                                                                                                                                           | `kubevault`                                               |
| `operator.repository`                   | Vault operator container image                                                                                                                                                                                                                                                                                                                              | `vault-operator`                                          |
| `operator.tag`                          | Vault operator container image tag                                                                                                                                                                                                                                                                                                                          | `v0.3.0`                                                   |
| `cleaner.registry`                      | Docker registry used to pull Webhook cleaner image                                                                                                                                                                                                                                                                                                          | `appscode`                                                |
| `cleaner.repository`                    | Webhook cleaner container image                                                                                                                                                                                                                                                                                                                             | `kubectl`                                                 |
| `cleaner.tag`                           | Webhook cleaner container image tag                                                                                                                                                                                                                                                                                                                         | `v1.11`                                                   |
| `imagePullSecrets`                      | Specify image pull secrets                                                                                                                                                                                                                                                                                                                                  | `[]`                                                      |
| `imagePullPolicy`                       | Image pull policy                                                                                                                                                                                                                                                                                                                                           | `IfNotPresent`                                            |
| `criticalAddon`                         | If true, installs Vault operator as critical addon                                                                                                                                                                                                                                                                                                          | `false`                                                   |
| `logLevel`                              | Log level for operator                                                                                                                                                                                                                                                                                                                                      | `3`                                                       |
| `affinity`                              | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                           | `{}`                                                      |
| `nodeSelector`                          | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                              | `{}`                                                      |
| `tolerations`                           | Tolerations used pod assignment                                                                                                                                                                                                                                                                                                                             | `[]`                                                      |
| `serviceAccount.create`                 | If `true`, create a new service account                                                                                                                                                                                                                                                                                                                     | `true`                                                    |
| `serviceAccount.name`                   | Service account to be used. If not set and `serviceAccount.create` is `true`, a name is generated using the fullname template                                                                                                                                                                                                                               | ``                                                        |
| `apiserver.groupPriorityMinimum`        | The minimum priority the group should have.                                                                                                                                                                                                                                                                                                                 | 10000                                                     |
| `apiserver.versionPriority`             | The ordering of this API inside of the group.                                                                                                                                                                                                                                                                                                               | 15                                                        |
| `apiserver.enableValidatingWebhook`     | Enable validating webhooks for Vault CRDs                                                                                                                                                                                                                                                                                                                   | true                                                      |
| `apiserver.enableMutatingWebhook`       | Enable mutating webhooks for Vault CRDs                                                                                                                                                                                                                                                                                                                     | true                                                      |
| `apiserver.ca`                          | CA certificate used by main Kubernetes api server                                                                                                                                                                                                                                                                                                           | `not-ca-cert`                                             |
| `apiserver.bypassValidatingWebhookXray` | If true, bypasses validating webhook xray checks                                                                                                                                                                                                                                                                                                            | `false`                                                   |
| `apiserver.useKubeapiserverFqdnForAks`  | If true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522                                                                                                                                                                                                                                                     | `true`                                                    |
| `apiserver.healthcheck.enabled`         | Enable readiness and liveliness probes                                                                                                                                                                                                                                                                                                                      | `false`                                                   |
| `apiserver.servingCerts.generate`       | If true, generate on install/upgrade the certs that allow the kube-apiserver (and potentially ServiceMonitor) to authenticate vault-operator pods. Otherwise specify in `apiserver.servingCerts.{caCrt, serverCrt, serverKey}`. See also: [example terraform](https://github.com/kubevault/installer/blob/master/charts/vault-operator/example-terraform.tf) | `true`                                                    |
| `enableAnalytics`                       | Send usage events to Google Analytics                                                                                                                                                                                                                                                                                                                       | `true`                                                    |
| `monitoring.agent`                      | Specify which monitoring agent to use for monitoring Vault. It accepts either `prometheus.io/builtin` or `prometheus.io/coreos-operator`.                                                                                                                                                                                                                   | `none`                                                    |
| `monitoring.operator`                   | Specify whether to monitor Vault operator.                                                                                                                                                                                                                                                                                                                  | `false`                                                   |
| `monitoring.prometheus.namespace`       | Specify the namespace where Prometheus server is running or will be deployed.                                                                                                                                                                                                                                                                               | Release namespace                                         |
| `monitoring.serviceMonitor.labels`      | Specify the labels for ServiceMonitor. Prometheus crd will select ServiceMonitor using these labels. Only usable when monitoring agent is `prometheus.io/coreos-operator`.                                                                                                                                                                                  | `app: <generated app name>` and `release: <release name>` |
| `clusterName`                           | Specify the name of cluster used in a multi-cluster setup                                                                                                                                                                                                                                                                                                   |                                                           |

Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install vault-operator appscode/vault-operator -n kube-system --set image.tag=v0.3.0
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install vault-operator appscode/vault-operator -n kube-system --values values.yaml
```
