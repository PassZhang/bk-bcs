// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by main. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in BackendList) DeepCopyInto(out *BackendList) {
	{
		in := &in
		*out = make(BackendList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Backend)
				**out = **in
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendList.
func (in BackendList) DeepCopy() BackendList {
	if in == nil {
		return nil
	}
	out := new(BackendList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudListener) DeepCopyInto(out *CloudListener) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudListener.
func (in *CloudListener) DeepCopy() *CloudListener {
	if in == nil {
		return nil
	}
	out := new(CloudListener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudListener) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudListenerBackendHealthStatus) DeepCopyInto(out *CloudListenerBackendHealthStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudListenerBackendHealthStatus.
func (in *CloudListenerBackendHealthStatus) DeepCopy() *CloudListenerBackendHealthStatus {
	if in == nil {
		return nil
	}
	out := new(CloudListenerBackendHealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudListenerHealthStatus) DeepCopyInto(out *CloudListenerHealthStatus) {
	*out = *in
	if in.RulesHealth != nil {
		in, out := &in.RulesHealth, &out.RulesHealth
		*out = make([]*CloudListenerRuleHealthStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CloudListenerRuleHealthStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudListenerHealthStatus.
func (in *CloudListenerHealthStatus) DeepCopy() *CloudListenerHealthStatus {
	if in == nil {
		return nil
	}
	out := new(CloudListenerHealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudListenerList) DeepCopyInto(out *CloudListenerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudListener, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudListenerList.
func (in *CloudListenerList) DeepCopy() *CloudListenerList {
	if in == nil {
		return nil
	}
	out := new(CloudListenerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudListenerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudListenerRuleHealthStatus) DeepCopyInto(out *CloudListenerRuleHealthStatus) {
	*out = *in
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]*CloudListenerBackendHealthStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CloudListenerBackendHealthStatus)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudListenerRuleHealthStatus.
func (in *CloudListenerRuleHealthStatus) DeepCopy() *CloudListenerRuleHealthStatus {
	if in == nil {
		return nil
	}
	out := new(CloudListenerRuleHealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudListenerSpec) DeepCopyInto(out *CloudListenerSpec) {
	*out = *in
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(CloudListenerTls)
		**out = **in
	}
	if in.TargetGroup != nil {
		in, out := &in.TargetGroup, &out.TargetGroup
		*out = new(TargetGroup)
		(*in).DeepCopyInto(*out)
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make(RuleList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Rule)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudListenerSpec.
func (in *CloudListenerSpec) DeepCopy() *CloudListenerSpec {
	if in == nil {
		return nil
	}
	out := new(CloudListenerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudListenerStatus) DeepCopyInto(out *CloudListenerStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	if in.HealthStatus != nil {
		in, out := &in.HealthStatus, &out.HealthStatus
		*out = new(CloudListenerHealthStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudListenerStatus.
func (in *CloudListenerStatus) DeepCopy() *CloudListenerStatus {
	if in == nil {
		return nil
	}
	out := new(CloudListenerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudListenerTls) DeepCopyInto(out *CloudListenerTls) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudListenerTls.
func (in *CloudListenerTls) DeepCopy() *CloudListenerTls {
	if in == nil {
		return nil
	}
	out := new(CloudListenerTls)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudLoadBalancer) DeepCopyInto(out *CloudLoadBalancer) {
	*out = *in
	if in.PublicIPs != nil {
		in, out := &in.PublicIPs, &out.PublicIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VIPS != nil {
		in, out := &in.VIPS, &out.VIPS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudLoadBalancer.
func (in *CloudLoadBalancer) DeepCopy() *CloudLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(CloudLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	if in.TargetGroup != nil {
		in, out := &in.TargetGroup, &out.TargetGroup
		*out = new(TargetGroup)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RuleList) DeepCopyInto(out *RuleList) {
	{
		in := &in
		*out = make(RuleList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Rule)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleList.
func (in RuleList) DeepCopy() RuleList {
	if in == nil {
		return nil
	}
	out := new(RuleList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroup) DeepCopyInto(out *TargetGroup) {
	*out = *in
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(TargetGroupHealthCheck)
		**out = **in
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make(BackendList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Backend)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroup.
func (in *TargetGroup) DeepCopy() *TargetGroup {
	if in == nil {
		return nil
	}
	out := new(TargetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupHealthCheck) DeepCopyInto(out *TargetGroupHealthCheck) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupHealthCheck.
func (in *TargetGroupHealthCheck) DeepCopy() *TargetGroupHealthCheck {
	if in == nil {
		return nil
	}
	out := new(TargetGroupHealthCheck)
	in.DeepCopyInto(out)
	return out
}