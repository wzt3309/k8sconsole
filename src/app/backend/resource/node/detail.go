package node

import (
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/resource/common"
	"github.com/wzt3309/k8sconsole/src/app/backend/resource/pod"
	"k8s.io/api/core/v1"
)

// NodeAllocatedResources describes node allocated resources.
type NodeAllocatedResources struct {
	/**
		Kubernetes has a new metric called Millicores that is used to measure CPU usage.
		It is a CPU core split into 1000 units (milli = 1000).
	   - 1. 1 cpu with 1 core has 1000m
	   - 2. 1 cpu with 2 core has 2*1000m = 2000m
	 */
	// CPURequests is number of allocated millicores.
	CPURequests int64 `json:"cpuRequests"`

	// CPURequestsFraction is a fraction of CPU, that is allocated.
	CPURequestsFraction float64 `json:"cpuRequestFraction"`

	// CPULimits is defined CPU limit.
	CPULimits int64 `json:"cpuLimits"`

	// CPULimitsFraction is a fraction of defined CPU limit.
	// Note. can be over 100%, i.e. overcommitted.
	CPULimitsFraction float64 `json:"cpuLimitsFraction"`

	// CPUCapacity is specified node CPU capacity in millicores.
	CPUCapacity int64 `json:"cpuCapacity"`

	// MemoryRequests is a fraction of memory, that is allocated.
	MemoryRequests int64 `json:"memoryRequests"`

	// MemoryRequestsFraction is a fraction of memory, that is allocated.
	MemoryRequestsFraction float64 `json:"memoryRequestsFraction"`

	// MemoryLimits is defined memory limit.
	MemoryLimits int64 `json:"memoryLimits"`

	// MemoryLimitsFraction is a fraction of defined memory limit, can be over 100%, i.e.
	// overcommitted.
	MemoryLimitsFraction float64 `json:"memoryLimitsFraction"`

	// MemoryCapacity is specified node memory capacity in bytes.
	MemoryCapacity int64 `json:"memoryCapacity"`

	// AllocatedPods in number of currently allocated pods on the node.
	AllocatedPods int `json:"allocatedPods"`

	// PodCapacity is maximum number of pods, that can be allocated on the node.
	PodCapacity int64 `json:"podCapacity"`

	// PodFraction is a fraction of pods, that can be allocated on given node.
	PodFraction float64 `json:"podFraction"`
}

// NodeDetail is a presentation layer view of Kubernetes Node resource. This means it is Node plus
// additional augmented data we can get from other sources.
type NodeDetail struct {
	ObjectMeta api.ObjectMeta `json:"objectMeta"`
	TypeMeta   api.TypeMeta   `json:"typeMeta"`

	// NodePhase is the current lifecycle phase of the node.
	Phase v1.NodePhase `json:"phase"`

	// Resources allocated by node.
	AllocatedResources NodeAllocatedResources `json:"allocatedResources"`

	// PodCIDR represents the pod IP range assigned to the node.
	PodCIDR string `json:"podCIDR"`

	// ID of the node assigned by the cloud provider.
	ProviderID string `json:"providerID"`

	// Unschedulable controls node schedulability of new pods. By default node is schedulable.
	Unschedulable bool `json:"unschedulable"`

	// Set of ids/uuids to uniquely identify the node.
	NodeInfo v1.NodeSystemInfo `json:"nodeInfo"`

	// Conditions is an array of current node conditions.
	Conditions []common.Condition `json:"conditions"`

	// Container images of the node.
	ContainerImages []string `json:"containerImages"`

	// PodList contains information about pods belonging to this node.
	PodList pod.PodList `json:"podList"`

	// Events is list of events associated to the node.
	EventList common.EventList `json:"eventList"`

	// Taints
	Taints []v1.Taint `json:"taints,omitempty"`

	// Addresses is a list of addresses reachable to the node. Queried from cloud provider, if available.
	Addresses []v1.NodeAddress `json:"addresses,omitempty"`

	// List of non-critical errors, that occurred during resource retrieval.
	Errors []error `json:"errors"`
}
