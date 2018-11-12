package api

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

// CsrfToken is used to secure requests from CSRF attacks
type CsrfToken struct {
	// Token generated on request for validation
	Token string `json:"token"`
}

// ObjectMeta is metadata about an instance of resource
type ObjectMeta struct {
	// Object name and the name is unique within a namespace.
	Name 								string 							`json:"name,omitempty"`

	// Any empty namespace equivalent to the 'default' namespace.
	// Not all objects are required to be scoped to a namespace - the value of this field
	// for those objects will be empty
	Namespace 					string 							`json:"namespace,omitempty"`

	// Labels are k-v pairs that may be scope and select individual resources.
	Labels 							map[string]string		`json:"labels,omitempty"`

	// Annotations are unstructured key value data stored with a resource that be set by external tooling.
	Annotations 				map[string]string 	`json:"annotations,omitempty"`

	// CreationTimestamp is a timestamp representing the apiserver time when this object
	// was created.
	CreationTimestamp		metaV1.Time					`json:"creationTimestamp,omitempty"`
}

// TypeMeta describes the type of an object in response and request
type TypeMeta struct {
	// kind of an object
	Kind	ResourceKind	`json:"kind,omitempty"`
}

// ListMeta describes list of objects.
type ListMeta struct {
	TotalItems int `json:"totalItems"`
}

// NewObjectMeta creates a new instance of ObjectMate struct based on k8s object meta.
func NewObjectMeta(k8sObjectMeta metaV1.ObjectMeta) ObjectMeta {
	return ObjectMeta{
		Name: k8sObjectMeta.Name,
		Namespace: k8sObjectMeta.Namespace,
		Labels: k8sObjectMeta.Labels,
		Annotations: k8sObjectMeta.Annotations,
		CreationTimestamp: k8sObjectMeta.CreationTimestamp,
	}
}

// NewTypeMeta creates new type meta for the resource kind
func NewTypeMeta(kind ResourceKind) TypeMeta {
	return TypeMeta{
		Kind: kind,
	}
}

type ResourceKind string

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

// ListEverything is a list options used to list all resource without any filtering.
var ListEverything = metaV1.ListOptions{
	LabelSelector: labels.Everything().String(),
	FieldSelector: fields.Everything().String(),
}
