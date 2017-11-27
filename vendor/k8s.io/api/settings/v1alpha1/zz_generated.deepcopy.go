// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PodPreset, InType: reflect.TypeOf(&PodPreset{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PodPresetList, InType: reflect.TypeOf(&PodPresetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_v1alpha1_PodPresetSpec, InType: reflect.TypeOf(&PodPresetSpec{})},
	)
}

func DeepCopy_v1alpha1_PodPreset(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodPreset)
		out := out.(*PodPreset)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_v1alpha1_PodPresetSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

func DeepCopy_v1alpha1_PodPresetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodPresetList)
		out := out.(*PodPresetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]PodPreset, len(*in))
			for i := range *in {
				if err := DeepCopy_v1alpha1_PodPreset(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func DeepCopy_v1alpha1_PodPresetSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*PodPresetSpec)
		out := out.(*PodPresetSpec)
		*out = *in
		if newVal, err := c.DeepCopy(&in.Selector); err != nil {
			return err
		} else {
			out.Selector = *newVal.(*v1.LabelSelector)
		}
		if in.Env != nil {
			in, out := &in.Env, &out.Env
			*out = make([]core_v1.EnvVar, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*core_v1.EnvVar)
				}
			}
		}
		if in.EnvFrom != nil {
			in, out := &in.EnvFrom, &out.EnvFrom
			*out = make([]core_v1.EnvFromSource, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*core_v1.EnvFromSource)
				}
			}
		}
		if in.Volumes != nil {
			in, out := &in.Volumes, &out.Volumes
			*out = make([]core_v1.Volume, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*core_v1.Volume)
				}
			}
		}
		if in.VolumeMounts != nil {
			in, out := &in.VolumeMounts, &out.VolumeMounts
			*out = make([]core_v1.VolumeMount, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}
