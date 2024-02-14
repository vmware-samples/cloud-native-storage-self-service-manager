module cns.vmware.com/cns-manager

go 1.19

require (
	github.com/antihax/optional v1.0.0
	github.com/go-logr/zapr v1.2.3
	go.uber.org/zap v1.24.0
	golang.org/x/oauth2 v0.5.0
	sigs.k8s.io/controller-runtime v0.13.0
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/onsi/gomega v1.27.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	k8s.io/apimachinery v0.25.2 // indirect
)

replace (
	github.com/go-logr/logr => github.com/go-logr/logr v1.2.0
	github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.4.1
	github.com/kubernetes-csi/csi-lib-utils => github.com/kubernetes-csi/csi-lib-utils v0.11.0
	k8s.io/api => k8s.io/api v0.25.2
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.25.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.25.2
	k8s.io/apiserver => k8s.io/apiserver v0.25.2
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.25.2
	k8s.io/client-go => k8s.io/client-go v0.25.2
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.25.2
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.25.2
	k8s.io/code-generator => k8s.io/code-generator v0.25.2
	k8s.io/component-base => k8s.io/component-base v0.25.2
	k8s.io/component-helpers => k8s.io/component-helpers v0.25.2
	k8s.io/controller-manager => k8s.io/controller-manager v0.25.2
	k8s.io/cri-api => k8s.io/cri-api v0.25.2
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.25.2
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.25.2
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.25.2
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.25.2
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.25.2
	k8s.io/kubectl => k8s.io/kubectl v0.25.2
	k8s.io/kubelet => k8s.io/kubelet v0.25.2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.25.2
	k8s.io/metrics => k8s.io/metrics v0.25.2
	k8s.io/mount-utils => k8s.io/mount-utils v0.24.6
	k8s.io/node-api => k8s.io/node-api v0.25.2
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.25.2
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.25.2
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.25.2
	k8s.io/sample-controller => k8s.io/sample-controller v0.25.2
)
