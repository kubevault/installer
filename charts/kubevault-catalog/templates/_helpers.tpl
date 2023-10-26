{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "kubevault-catalog.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "kubevault-catalog.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "kubevault-catalog.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "kubevault-catalog.labels" -}}
helm.sh/chart: {{ include "kubevault-catalog.chart" . }}
{{ include "kubevault-catalog.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "kubevault-catalog.selectorLabels" -}}
app.kubernetes.io/name: {{ include "kubevault-catalog.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "kubevault-catalog.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "kubevault-catalog.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "image.dockerHub" -}}
{{ prepend (list ._repo) (list .Values.proxies.dockerHub .Values.registryFQDN | compact | first) | compact | join "/" }}
{{- end }}

{{- define "image.dockerLibrary" -}}
{{ prepend (list ._repo) (list .Values.proxies.dockerLibrary .Values.proxies.dockerHub .Values.registryFQDN | compact | first) | compact | join "/" }}
{{- end }}

{{- define "image.ghcr" -}}
{{ prepend (list ._repo) (list .Values.proxies.ghcr .Values.registryFQDN | compact | first) | compact | join "/" }}
{{- end }}

{{- define "image.kubernetes" -}}
{{ prepend (list ._repo) (list .Values.proxies.kubernetes .Values.registryFQDN | compact | first) | compact | join "/" }}
{{- end }}

{{- define "image.appscode" -}}
{{ prepend (list ._repo) (list .Values.proxies.appscode .Values.registryFQDN | compact | first) | compact | join "/" }}
{{- end }}
