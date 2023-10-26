module kubevault.dev/installer

go 1.18

require (
	github.com/Masterminds/semver/v3 v3.2.1
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/gogo/protobuf v1.3.2
	github.com/google/gofuzz v1.2.0
	github.com/spf13/pflag v1.0.5
	github.com/yudai/gojsondiff v1.0.0
	gomodules.xyz/go-sh v0.1.0
	gomodules.xyz/semvers v0.0.0-20220316103017-cfbe8c37b63d
	k8s.io/api v0.25.3
	k8s.io/apimachinery v0.25.3
	kmodules.xyz/client-go v0.25.39
	kmodules.xyz/schema-checker v0.4.1
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1-0.20220316001817-d5090ed65664 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gobeam/stringy v0.0.5 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/gomega v1.24.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/stretchr/testify v1.8.3 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	gomodules.xyz/encoding v0.0.7 // indirect
	gomodules.xyz/jsonpath v0.0.2 // indirect
	gomodules.xyz/mergo v0.3.13 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/utils v0.0.0-20221012122500-cfd413dd9e85 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.3.0

replace github.com/Masterminds/sprig/v3 => github.com/gomodules/sprig/v3 v3.2.3-0.20220405051441-0a8a99bac1b8

replace sigs.k8s.io/controller-runtime => github.com/kmodules/controller-runtime v0.13.1-0.20230725210630-f8273f9c4412

replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.6

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.25.2-0.20230817060727-aa2afd31b825

replace k8s.io/kubernetes => github.com/kmodules/kubernetes v1.26.0-alpha.0.0.20220917022409-54ac9f3f1315
