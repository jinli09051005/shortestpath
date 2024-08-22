//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (C) 2024 JinLi Co.,Ltd. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Display) DeepCopyInto(out *Display) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Display.
func (in *Display) DeepCopy() *Display {
	if in == nil {
		return nil
	}
	out := new(Display)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Display) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisplayList) DeepCopyInto(out *DisplayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Display, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisplayList.
func (in *DisplayList) DeepCopy() *DisplayList {
	if in == nil {
		return nil
	}
	out := new(DisplayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DisplayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisplaySpec) DeepCopyInto(out *DisplaySpec) {
	*out = *in
	out.StartNode = in.StartNode
	if in.TargetNodes != nil {
		in, out := &in.TargetNodes, &out.TargetNodes
		*out = make([]TargetNode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisplaySpec.
func (in *DisplaySpec) DeepCopy() *DisplaySpec {
	if in == nil {
		return nil
	}
	out := new(DisplaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DisplayStatus) DeepCopyInto(out *DisplayStatus) {
	*out = *in
	in.LastUpdate.DeepCopyInto(&out.LastUpdate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DisplayStatus.
func (in *DisplayStatus) DeepCopy() *DisplayStatus {
	if in == nil {
		return nil
	}
	out := new(DisplayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Edge) DeepCopyInto(out *Edge) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Edge.
func (in *Edge) DeepCopy() *Edge {
	if in == nil {
		return nil
	}
	out := new(Edge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnownNodes) DeepCopyInto(out *KnownNodes) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnownNodes.
func (in *KnownNodes) DeepCopy() *KnownNodes {
	if in == nil {
		return nil
	}
	out := new(KnownNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KnownNodes) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnownNodesList) DeepCopyInto(out *KnownNodesList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KnownNodes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnownNodesList.
func (in *KnownNodesList) DeepCopy() *KnownNodesList {
	if in == nil {
		return nil
	}
	out := new(KnownNodesList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KnownNodesList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnownNodesSpec) DeepCopyInto(out *KnownNodesSpec) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]Node, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnownNodesSpec.
func (in *KnownNodesSpec) DeepCopy() *KnownNodesSpec {
	if in == nil {
		return nil
	}
	out := new(KnownNodesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnownNodesStatus) DeepCopyInto(out *KnownNodesStatus) {
	*out = *in
	in.LastUpdate.DeepCopyInto(&out.LastUpdate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnownNodesStatus.
func (in *KnownNodesStatus) DeepCopy() *KnownNodesStatus {
	if in == nil {
		return nil
	}
	out := new(KnownNodesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	if in.PreNode != nil {
		in, out := &in.PreNode, &out.PreNode
		*out = new(Node)
		(*in).DeepCopyInto(*out)
	}
	if in.Edges != nil {
		in, out := &in.Edges, &out.Edges
		*out = make([]Edge, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartNode) DeepCopyInto(out *StartNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartNode.
func (in *StartNode) DeepCopy() *StartNode {
	if in == nil {
		return nil
	}
	out := new(StartNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetNode) DeepCopyInto(out *TargetNode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetNode.
func (in *TargetNode) DeepCopy() *TargetNode {
	if in == nil {
		return nil
	}
	out := new(TargetNode)
	in.DeepCopyInto(out)
	return out
}
