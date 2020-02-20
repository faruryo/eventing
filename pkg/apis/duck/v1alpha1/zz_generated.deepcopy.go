// +build !ignore_autogenerated

/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1beta1 "knative.dev/eventing/pkg/apis/duck/v1beta1"
	apis "knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelTemplateSpec) DeepCopyInto(out *ChannelTemplateSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelTemplateSpec.
func (in *ChannelTemplateSpec) DeepCopy() *ChannelTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ChannelTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChannelTemplateSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelTemplateSpecInternal) DeepCopyInto(out *ChannelTemplateSpecInternal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelTemplateSpecInternal.
func (in *ChannelTemplateSpecInternal) DeepCopy() *ChannelTemplateSpecInternal {
	if in == nil {
		return nil
	}
	out := new(ChannelTemplateSpecInternal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChannelTemplateSpecInternal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Channelable) DeepCopyInto(out *Channelable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Channelable.
func (in *Channelable) DeepCopy() *Channelable {
	if in == nil {
		return nil
	}
	out := new(Channelable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Channelable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelableList) DeepCopyInto(out *ChannelableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Channelable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelableList.
func (in *ChannelableList) DeepCopy() *ChannelableList {
	if in == nil {
		return nil
	}
	out := new(ChannelableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChannelableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelableSpec) DeepCopyInto(out *ChannelableSpec) {
	*out = *in
	in.SubscribableTypeSpec.DeepCopyInto(&out.SubscribableTypeSpec)
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelableSpec.
func (in *ChannelableSpec) DeepCopy() *ChannelableSpec {
	if in == nil {
		return nil
	}
	out := new(ChannelableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelableStatus) DeepCopyInto(out *ChannelableStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	in.AddressStatus.DeepCopyInto(&out.AddressStatus)
	in.SubscribableTypeStatus.DeepCopyInto(&out.SubscribableTypeStatus)
	if in.ErrorChannel != nil {
		in, out := &in.ErrorChannel, &out.ErrorChannel
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelableStatus.
func (in *ChannelableStatus) DeepCopy() *ChannelableStatus {
	if in == nil {
		return nil
	}
	out := new(ChannelableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeliverySpec) DeepCopyInto(out *DeliverySpec) {
	*out = *in
	if in.DeadLetterSink != nil {
		in, out := &in.DeadLetterSink, &out.DeadLetterSink
		*out = new(duckv1.Destination)
		(*in).DeepCopyInto(*out)
	}
	if in.Retry != nil {
		in, out := &in.Retry, &out.Retry
		*out = new(int32)
		**out = **in
	}
	if in.BackoffPolicy != nil {
		in, out := &in.BackoffPolicy, &out.BackoffPolicy
		*out = new(BackoffPolicyType)
		**out = **in
	}
	if in.BackoffDelay != nil {
		in, out := &in.BackoffDelay, &out.BackoffDelay
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeliverySpec.
func (in *DeliverySpec) DeepCopy() *DeliverySpec {
	if in == nil {
		return nil
	}
	out := new(DeliverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeliveryStatus) DeepCopyInto(out *DeliveryStatus) {
	*out = *in
	if in.DeadLetterChannel != nil {
		in, out := &in.DeadLetterChannel, &out.DeadLetterChannel
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeliveryStatus.
func (in *DeliveryStatus) DeepCopy() *DeliveryStatus {
	if in == nil {
		return nil
	}
	out := new(DeliveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Resource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceList) DeepCopyInto(out *ResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceList.
func (in *ResourceList) DeepCopy() *ResourceList {
	if in == nil {
		return nil
	}
	out := new(ResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscribable) DeepCopyInto(out *Subscribable) {
	*out = *in
	if in.Subscribers != nil {
		in, out := &in.Subscribers, &out.Subscribers
		*out = make([]SubscriberSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscribable.
func (in *Subscribable) DeepCopy() *Subscribable {
	if in == nil {
		return nil
	}
	out := new(Subscribable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscribableStatus) DeepCopyInto(out *SubscribableStatus) {
	*out = *in
	if in.Subscribers != nil {
		in, out := &in.Subscribers, &out.Subscribers
		*out = make([]SubscriberStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscribableStatus.
func (in *SubscribableStatus) DeepCopy() *SubscribableStatus {
	if in == nil {
		return nil
	}
	out := new(SubscribableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscribableType) DeepCopyInto(out *SubscribableType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscribableType.
func (in *SubscribableType) DeepCopy() *SubscribableType {
	if in == nil {
		return nil
	}
	out := new(SubscribableType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscribableType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscribableTypeList) DeepCopyInto(out *SubscribableTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SubscribableType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscribableTypeList.
func (in *SubscribableTypeList) DeepCopy() *SubscribableTypeList {
	if in == nil {
		return nil
	}
	out := new(SubscribableTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubscribableTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscribableTypeSpec) DeepCopyInto(out *SubscribableTypeSpec) {
	*out = *in
	if in.Subscribable != nil {
		in, out := &in.Subscribable, &out.Subscribable
		*out = new(Subscribable)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscribableTypeSpec.
func (in *SubscribableTypeSpec) DeepCopy() *SubscribableTypeSpec {
	if in == nil {
		return nil
	}
	out := new(SubscribableTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscribableTypeStatus) DeepCopyInto(out *SubscribableTypeStatus) {
	*out = *in
	if in.SubscribableStatus != nil {
		in, out := &in.SubscribableStatus, &out.SubscribableStatus
		*out = new(SubscribableStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscribableTypeStatus.
func (in *SubscribableTypeStatus) DeepCopy() *SubscribableTypeStatus {
	if in == nil {
		return nil
	}
	out := new(SubscribableTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberSpec) DeepCopyInto(out *SubscriberSpec) {
	*out = *in
	if in.SubscriberURI != nil {
		in, out := &in.SubscriberURI, &out.SubscriberURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplyURI != nil {
		in, out := &in.ReplyURI, &out.ReplyURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.DeadLetterSinkURI != nil {
		in, out := &in.DeadLetterSinkURI, &out.DeadLetterSinkURI
		*out = new(apis.URL)
		(*in).DeepCopyInto(*out)
	}
	if in.Delivery != nil {
		in, out := &in.Delivery, &out.Delivery
		*out = new(v1beta1.DeliverySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberSpec.
func (in *SubscriberSpec) DeepCopy() *SubscriberSpec {
	if in == nil {
		return nil
	}
	out := new(SubscriberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriberStatus) DeepCopyInto(out *SubscriberStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriberStatus.
func (in *SubscriberStatus) DeepCopy() *SubscriberStatus {
	if in == nil {
		return nil
	}
	out := new(SubscriberStatus)
	in.DeepCopyInto(out)
	return out
}
