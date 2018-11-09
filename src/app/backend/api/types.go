package api

// List of all resource kinds
// TODO(wzt3309) Need to add more resource kinds
const (
	ResourceKindConfigMap  = "configmap"
	ResourceKindDaemonSet  = "daemonset"
	ResourceKindDeployment = "deployment"
	ResourceKindNamespace  = "namespace"
	ResourceKindNode       = "node"
	ResourceKindPod        = "pod"
	ResourceKindReplicaSet = "replicaset"
	ResourceKindSecret     = "secret"
	ResourceKindService    = "service"
)

// ClientType represents type of client that is used to perform generic operations on resources.
// Different resources belong to different client, i.e. Deployments belongs to extension client
// and StatefulSets to apps client.
type ClientType string

// List of client types.
// TODO(wzt3309) Need to add more client types
const (
	ClientTypeDefault         = "restclient"
	ClientTypeExtensionClient = "extensionclient"
)

// Mapping from resource kind to K8s apiserver API path.
// TODO(wzt3309) Need to add more mappings
var KindToAPIMapping = map[string]struct {
	// k8s resource name
	Resource string
	// Client type used by given resource, i.e. deployments using extension client
	ClientType ClientType
	// Is this object global scoped (not below a namespace), i.e. 'kubectl get node'
	Namespaced bool
}{
	ResourceKindConfigMap:  {"configmaps", ClientTypeDefault, true},
	ResourceKindDaemonSet:  {"daemonsets", ClientTypeExtensionClient, true},
	ResourceKindDeployment: {"deployments", ClientTypeExtensionClient, true},
	ResourceKindNamespace:  {"namespaces", ClientTypeDefault, false},
	ResourceKindNode:       {"nodes", ClientTypeDefault, false},
	ResourceKindPod:        {"pods", ClientTypeDefault, true},
	ResourceKindReplicaSet: {"replicasets", ClientTypeExtensionClient, true},
	ResourceKindSecret:     {"secrets", ClientTypeDefault, true},
	ResourceKindService:    {"services", ClientTypeDefault, true},
}
