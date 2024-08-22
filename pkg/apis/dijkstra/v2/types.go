package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KnownNodes
// KnownNodes资源对象结构定义
type KnownNodes struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// +optional
	Spec KnownNodesSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	// +optional
	Status KnownNodesStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// KnownNodesList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KnownNodesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []KnownNodes `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// KnownNodesSpec defines the desired state of KnownNodes
type KnownNodesSpec struct {
	// A type of node identity
	// +optional
	NodeIdentity string `json:"nodeIdentity" protobuf:"bytes,1,rep,name=nodeIdentity"`
	// To node cost unit
	// +optional
	CostUnit string `json:"costUnit,omitempty" protobuf:"bytes,2,rep,name=costUnit"`
	// Known nodes information
	// +optional
	// +patchMergeKey=id
	// +patchStrategy=merge
	Nodes []Node `json:"nodes,omitempty" patchStrategy:"merge" patchMergeKey:"id" protobuf:"bytes,3,rep,name=nodes"`
}

type Node struct {
	// Node id
	// +optional
	ID int32 `json:"id,omitempty" protobuf:"bytes,1,rep,name=id"`
	// Node name
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,2,rep,name=name"`
	// Previous node
	// +optional
	PreNode *Node `json:"preNode,omitempty" protobuf:"bytes,3,rep,name=preNode"`
	// Node edges
	// +optional
	Edges []Edge `json:"edges,omitempty" protobuf:"bytes,4,rep,name=edges"`
}

type Edge struct {
	// To node id
	// +optional
	To int32 `json:"id,omitempty" protobuf:"bytes,1,rep,name=id"`
	// To node cost
	// +optional
	Cost int32 `json:"cost,omitempty" protobuf:"bytes,2,rep,name=cost"`
}

// KnownNodesStatus defines the observed state of KnownNodes
type KnownNodesStatus struct {
	// Last Update Time
	// +optional
	LastUpdate metav1.Time `json:"lastUpdate,omitempty" protobuf:"bytes,1,opt,name=lastUpdate"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Display
type Display struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// +optional
	Spec DisplaySpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	// +optional
	Status DisplayStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// DisplayList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DisplayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Display `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// DisplaySpec defines the desired state of Display
type DisplaySpec struct {
	// A type of node identity
	// +optional
	NodeIdentity string `json:"nodeIdentity,omitempty" protobuf:"bytes,1,rep,name=nodeIdentity"`
	// Start node information
	// +optional
	StartNode StartNode `json:"startNode,omitempty" protobuf:"bytes,2,rep,name=startNode"`
	// Target nodes information
	// +optional
	// +patchMergeKey=id
	// +patchStrategy=merge
	TargetNodes []TargetNode `json:"targetNodes,omitempty" patchStrategy:"merge" patchMergeKey:"id" protobuf:"bytes,3,rep,name=targetNodes"`
	// Algorithms used to calculate the shortest path, including dijkstra and floyd algorithms
	// +optional
	Algorithm string `json:"algorithm,omitempty" protobuf:"bytes,4,rep,name=algorithm"`
}

type StartNode struct {
	// Node id
	// +optional
	ID int32 `json:"id,omitempty" protobuf:"bytes,1,rep,name=id"`
	// Node name
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,2,rep,name=name"`
}

type TargetNode struct {
	// Target node id
	// +optional
	ID int32 `json:"id,omitempty" protobuf:"bytes,1,rep,name=id"`
	// Target node name
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,2,rep,name=name"`
	// Start node to target node distance
	// +optional
	Distance int32 `json:"distance,omitempty" protobuf:"bytes,3,rep,name=distance"`
	// Start node to target node path
	// +optional
	Path string `json:"path,omitempty" protobuf:"bytes,4,rep,name=path"`
}

// DisplayStatus defines the observed state of Display
type DisplayStatus struct {
	// Last Update Time
	// +optional
	LastUpdate metav1.Time `json:"lastUpdate,omitempty" protobuf:"bytes,1,opt,name=lastUpdate"`
	// Dispaly  ShortestPath Compute Status
	// +optional
	ComputeStatus string `json:"computeStatus,omitempty" protobuf:"bytes,2,opt,name=computeStatus"`
}
