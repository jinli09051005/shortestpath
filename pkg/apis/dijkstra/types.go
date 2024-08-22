package dijkstra

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KnownNodes
type KnownNodes struct {
	metav1.TypeMeta
	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta
	// +optional
	Spec KnownNodesSpec
	// +optional
	Status KnownNodesStatus
}

// KnownNodesList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KnownNodesList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []KnownNodes
}

// KnownNodesSpec defines the desired state of KnownNodes
type KnownNodesSpec struct {
	// A type of node identity
	// +optional
	NodeIdentity string
	// To node cost unit
	// +optional
	CostUnit string
	// Known nodes information
	// +optional
	Nodes []Node
}

type Node struct {
	// Node id
	// +optional
	ID int32
	// Node name
	// +optional
	Name string
	// Previous node
	// +optional
	PreNode *Node
	// Node edges
	// +optional
	Edges []Edge
}

type Edge struct {
	// To node id
	// +optional
	To int32
	// To node cost
	// +optional
	Cost int32
}

// KnownNodesStatus defines the observed state of KnownNodes
type KnownNodesStatus struct {
	// Last Update Time
	// +optional
	LastUpdate metav1.Time
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Display
type Display struct {
	metav1.TypeMeta
	// Standard object's metadata.
	// +optional
	metav1.ObjectMeta
	// +optional
	Spec DisplaySpec
	// +optional
	Status DisplayStatus
}

// DisplayList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DisplayList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Display
}

// DisplaySpec defines the desired state of Display
type DisplaySpec struct {
	// A type of node identity
	// +optional
	NodeIdentity string
	// Start node information
	// +optional
	StartNode StartNode
	// Target nodes information
	// +optional
	TargetNodes []TargetNode
	// Algorithms used to calculate the shortest path, including dijkstra and floyd algorithms
	// +optional
	Algorithm string
}

type StartNode struct {
	// Node id
	// +optional
	ID int32
	// Node name
	// +optional
	Name string
}

type TargetNode struct {
	// Target node id
	// +optional
	ID int32
	// Target node name
	// +optional
	Name string
	// Start node to target node distance
	// +optional
	Distance int32
	// Start node to target node path
	// +optional
	Path string
}

// DisplayStatus defines the observed state of Display
type DisplayStatus struct {
	// Last Update Time
	// +optional
	LastUpdate metav1.Time
	// Dispaly  ShortestPath Compute Status
	// +optional
	ComputeStatus string
}
