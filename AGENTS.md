# AGENTS.md

This file provides guidance to coding agents (e.g. Claude Code, claude.ai/code) when working with code in this repository.

## Repository purpose

Go module `kubevault.dev/installer` — Helm charts and supporting tooling for installing [KubeVault](https://kubevault.com/). It also exposes a Go API package (`apis/installer/v1alpha1/`) describing chart values so other KubeVault components can consume strongly typed installation parameters.

Charts shipped:
- `charts/kubevault` — meta chart (operator + webhook + catalog + metrics).
- `charts/kubevault-operator` — the KubeVault operator (the main controller manager).
- `charts/kubevault-webhook-server` — admission webhook server.
- `charts/kubevault-catalog` — `VaultServerVersion` catalog.
- `charts/kubevault-crds` — CRDs only.
- `charts/kubevault-metrics` — Prometheus metrics exporter.
- `charts/kubevault-opscenter` — OpsCenter UI integration.
- `charts/kubevault-grafana-dashboards` — packaged Grafana dashboards.
- `charts/kubevault-certified`, `charts/kubevault-certified-crds` — Red Hat-certified packaging.

## Architecture

- `charts/` — one subdirectory per Helm chart. Each has `Chart.yaml`, `values.yaml`, `templates/`, plus generated artifacts `doc.yaml`, `README.md`, and (where applicable) `values.openapiv3_schema.yaml` and `crds/`.
- `apis/installer/v1alpha1/` — Go types backing chart values. Used for OpenAPI/values-schema generation and as a typed surface for downstream programs. Single API group: `installer:v1alpha1`.
- `crds/` — top-level CRDs (chart-specific CRDs live under each chart's `crds/`).
- `catalog/` — image catalog driven by `kmodules.xyz/image-packer`. `imagelist.yaml` is the source of truth.
- `hack/scripts/` — codegen/release helpers (`update-catalog.sh`, `update-chart-dependencies.sh`, `import-crds.sh`, `ct.sh`, etc.).
- `tests/` — chart-testing (`ct`) configuration.
- `lintconf.yaml` — YAML lint config.
- `third_party/` — vendored third-party assets.
- `vendor/` — Go vendored deps.
- `DEVELOPMENT.md` — developer guide.

## Common commands

All Make targets run inside `ghcr.io/appscode/golang-dev` — Docker must be running.

- `make gen` — regenerate everything: clientset + manifests + values schemas + chart docs. Run after any change to `apis/installer/v1alpha1/*_types.go`.
- `make clientset` — regenerate Go client only.
- `make manifests` — `gen-crds gen-values-schema gen-chart-doc`.
- `make gen-values-schema` — regenerate `values.openapiv3_schema.yaml` from the Go types.
- `make gen-chart-doc` — regenerate per-chart `README.md`.
- `make update-charts` — refresh chart-level metadata.
- `make refresh` — `gen update-catalog fmt`. **ALWAYS run this before opening a PR** so generated files are current.
- `make fmt`, `make lint`, `make unit-tests` / `make test` — standard.
- `make ct` — chart-testing lint+test.
- `make verify` — `verify-modules`; `go mod tidy && go mod vendor` must leave the tree clean.
- `make add-license` / `make check-license` — manage license headers.

Auxiliary helpers (invoked outside Make):

- `./hack/scripts/update-catalog.sh` — regenerate `catalog/` from `imagelist.yaml`.
- `./hack/scripts/import-crds.sh` — pull CRDs from `kubevault.dev/apimachinery` into each chart's `crds/`.

Run a single Go test (requires a local Go toolchain):

```
go test ./apis/installer/v1alpha1/... -run TestName -v
```

## Conventions

- Module path is `kubevault.dev/installer` (vanity URL). Imports must use that.
- Edit `apis/installer/v1alpha1/*_types.go` to change a chart's values surface, then run `make gen` so the generated clientset, `values.openapiv3_schema.yaml`, and per-chart `README.md` stay in sync. Do not hand-edit `zz_generated.*.go`, generated chart `README.md` files, `values.openapiv3_schema.yaml`, or anything under `catalog/` except `imagelist.yaml`.
- License: `LICENSE.md` (AppsCode); use `make add-license` to apply headers to new Go files.
- Sign off commits (`git commit -s`); contributions follow the DCO.
- Vendor directory is checked in — `go mod tidy && go mod vendor` must leave the tree clean (enforced by `verify-modules`).
- CRDs for charts are **imported** from `kubevault.dev/apimachinery` via `import-crds.sh`; do not hand-author `charts/*/crds/*.yaml`.
- The `-certified` charts mirror the standard charts for Red Hat certification — bumps must apply to both pairs (`kubevault-operator` ↔ `kubevault-certified`, `kubevault-crds` ↔ `kubevault-certified-crds`).
