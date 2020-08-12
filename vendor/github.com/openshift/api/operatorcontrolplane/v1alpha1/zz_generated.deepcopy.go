// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogEntry) DeepCopyInto(out *LogEntry) {
	*out = *in
	in.Start.DeepCopyInto(&out.Start)
	out.Latency = in.Latency
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogEntry.
func (in *LogEntry) DeepCopy() *LogEntry {
	if in == nil {
		return nil
	}
	out := new(LogEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutageEntry) DeepCopyInto(out *OutageEntry) {
	*out = *in
	in.Start.DeepCopyInto(&out.Start)
	in.End.DeepCopyInto(&out.End)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutageEntry.
func (in *OutageEntry) DeepCopy() *OutageEntry {
	if in == nil {
		return nil
	}
	out := new(OutageEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkConnectivityCheck) DeepCopyInto(out *PodNetworkConnectivityCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkConnectivityCheck.
func (in *PodNetworkConnectivityCheck) DeepCopy() *PodNetworkConnectivityCheck {
	if in == nil {
		return nil
	}
	out := new(PodNetworkConnectivityCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodNetworkConnectivityCheck) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkConnectivityCheckCondition) DeepCopyInto(out *PodNetworkConnectivityCheckCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkConnectivityCheckCondition.
func (in *PodNetworkConnectivityCheckCondition) DeepCopy() *PodNetworkConnectivityCheckCondition {
	if in == nil {
		return nil
	}
	out := new(PodNetworkConnectivityCheckCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkConnectivityCheckList) DeepCopyInto(out *PodNetworkConnectivityCheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodNetworkConnectivityCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkConnectivityCheckList.
func (in *PodNetworkConnectivityCheckList) DeepCopy() *PodNetworkConnectivityCheckList {
	if in == nil {
		return nil
	}
	out := new(PodNetworkConnectivityCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodNetworkConnectivityCheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkConnectivityCheckSpec) DeepCopyInto(out *PodNetworkConnectivityCheckSpec) {
	*out = *in
	out.TLSClientCert = in.TLSClientCert
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkConnectivityCheckSpec.
func (in *PodNetworkConnectivityCheckSpec) DeepCopy() *PodNetworkConnectivityCheckSpec {
	if in == nil {
		return nil
	}
	out := new(PodNetworkConnectivityCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodNetworkConnectivityCheckStatus) DeepCopyInto(out *PodNetworkConnectivityCheckStatus) {
	*out = *in
	if in.Successes != nil {
		in, out := &in.Successes, &out.Successes
		*out = make([]LogEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Failures != nil {
		in, out := &in.Failures, &out.Failures
		*out = make([]LogEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Outages != nil {
		in, out := &in.Outages, &out.Outages
		*out = make([]OutageEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PodNetworkConnectivityCheckCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodNetworkConnectivityCheckStatus.
func (in *PodNetworkConnectivityCheckStatus) DeepCopy() *PodNetworkConnectivityCheckStatus {
	if in == nil {
		return nil
	}
	out := new(PodNetworkConnectivityCheckStatus)
	in.DeepCopyInto(out)
	return out
}