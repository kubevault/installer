module kubevault.dev/installer

go 1.21.5

require (
	github.com/Masterminds/semver/v3 v3.2.1
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/gogo/protobuf v1.3.2
	github.com/google/go-containerregistry v0.17.0
	github.com/google/gofuzz v1.2.0
	github.com/spf13/pflag v1.0.5
	github.com/yudai/gojsondiff v1.0.0
	gomodules.xyz/go-sh v0.1.0
	gomodules.xyz/semvers v0.0.0-20220316103017-cfbe8c37b63d
	k8s.io/api v0.29.0
	k8s.io/apimachinery v0.29.0
	k8s.io/klog/v2 v2.110.1
	kmodules.xyz/client-go v0.29.6
	kmodules.xyz/go-containerregistry v0.0.12
	kmodules.xyz/schema-checker v0.4.1
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.14.3 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/docker/cli v24.0.7+incompatible // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/docker v24.0.7+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.7.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/gobeam/stringy v0.0.5 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc3 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/sirupsen/logrus v1.9.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/vbatts/tar-split v0.11.3 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	gomodules.xyz/encoding v0.0.7 // indirect
	gomodules.xyz/jsonpath v0.0.2 // indirect
	gomodules.xyz/mergo v0.3.13 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gotest.tools/v3 v3.1.0 // indirect
	k8s.io/utils v0.0.0-20231127182322-b307cd553661 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)

replace (
	github.com/Masterminds/sprig/v3 => github.com/gomodules/sprig/v3 v3.2.3-0.20220405051441-0a8a99bac1b8
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.6
	sigs.k8s.io/yaml => github.com/kmodules/yaml v1.4.1-0.20231224133800-a4e3f1abb174
)
