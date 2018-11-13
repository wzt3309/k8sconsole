package node

import (
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"k8s.io/api/core/v1"
)

type NodeList struct {
	ListMeta 	api.ListMeta 	`json:"listMeta"`
	Nodes 		[]Node 				`json:"nodes"`
	// List of non-critical errors, that occurred during resource retrieval.
	Errors 		[]error 			`json:"errors"`
}

type Node struct {
	ObjectMeta 	api.ObjectMeta `json:"objectMeta"`
	TypeMeta 		api.TypeMeta	 `json:"typeMeta"`
	Ready 			v1.ConditionStatus `json:"ready"`
}
