//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasExternalDNSRule) DeepCopyInto(out *MaasExternalDNSRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasExternalDNSRule.
func (in *MaasExternalDNSRule) DeepCopy() *MaasExternalDNSRule {
	if in == nil {
		return nil
	}
	out := new(MaasExternalDNSRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasInstance) DeepCopyInto(out *MaasInstance) {
	*out = *in
	out.Auth = in.Auth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasInstance.
func (in *MaasInstance) DeepCopy() *MaasInstance {
	if in == nil {
		return nil
	}
	out := new(MaasInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasInstanceRules) DeepCopyInto(out *MaasInstanceRules) {
	*out = *in
	if in.OSImages != nil {
		in, out := &in.OSImages, &out.OSImages
		*out = make([]OSImage, len(*in))
		copy(*out, *in)
	}
	if in.Nameservers != nil {
		in, out := &in.Nameservers, &out.Nameservers
		*out = make([]Nameserver, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasInstanceRules.
func (in *MaasInstanceRules) DeepCopy() *MaasInstanceRules {
	if in == nil {
		return nil
	}
	out := new(MaasInstanceRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasValidator) DeepCopyInto(out *MaasValidator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasValidator.
func (in *MaasValidator) DeepCopy() *MaasValidator {
	if in == nil {
		return nil
	}
	out := new(MaasValidator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaasValidator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasValidatorList) DeepCopyInto(out *MaasValidatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MaasValidator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasValidatorList.
func (in *MaasValidatorList) DeepCopy() *MaasValidatorList {
	if in == nil {
		return nil
	}
	out := new(MaasValidatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaasValidatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasValidatorSpec) DeepCopyInto(out *MaasValidatorSpec) {
	*out = *in
	out.MaasInstance = in.MaasInstance
	out.MaasExternalDNSRule = in.MaasExternalDNSRule
	in.MaasInstanceRules.DeepCopyInto(&out.MaasInstanceRules)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasValidatorSpec.
func (in *MaasValidatorSpec) DeepCopy() *MaasValidatorSpec {
	if in == nil {
		return nil
	}
	out := new(MaasValidatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaasValidatorStatus) DeepCopyInto(out *MaasValidatorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaasValidatorStatus.
func (in *MaasValidatorStatus) DeepCopy() *MaasValidatorStatus {
	if in == nil {
		return nil
	}
	out := new(MaasValidatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSImage) DeepCopyInto(out *OSImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSImage.
func (in *OSImage) DeepCopy() *OSImage {
	if in == nil {
		return nil
	}
	out := new(OSImage)
	in.DeepCopyInto(out)
	return out
}
