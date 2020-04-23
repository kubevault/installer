{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "csi-vault.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "csi-vault.fullname" -}}
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
{{- define "csi-vault.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "csi-vault.labels" -}}
helm.sh/chart: {{ include "csi-vault.chart" . }}
{{ include "csi-vault.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "csi-vault.selectorLabels" -}}
app.kubernetes.io/name: {{ include "csi-vault.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "csi-vault.attacher.fullname" -}}
{{- printf "%s-%s" (include "csi-vault.fullname" .) .Values.attacher.name | trunc 63 | trimSuffix "-" -}}
{{ end }}

{{- define "csi-vault.node-registrar.fullname" -}}
{{- printf "%s-%s" (include "csi-vault.fullname" .) .Values.nodeRegistrar.name | trunc 63 | trimSuffix "-" -}}
{{ end }}


{{- define "csi-vault.controller.fullname" -}}
{{- printf "%s-%s" (include "csi-vault.fullname" .) "controller" | trunc 63 | trimSuffix "-" -}}
{{ end }}

{{- define "csi-vault.controller.name" -}}
{{- printf "%s-%s" .Release.Name "controller" | trunc 63 | trimSuffix "-" -}}
{{ end }}

{{- define "csi-vault.node.fullname" -}}
{{- printf "%s-%s" (include "csi-vault.fullname" .) "node" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "csi-vault.node.name" -}}
{{- printf "%s-%s" .Release.Name "node" | trunc 63 | trimSuffix "-" -}}
{{ end }}

{{- define "csi-vault.node" -}}
valueFrom:
  fieldRef:
    fieldPath: spec.nodeName
{{- end -}}