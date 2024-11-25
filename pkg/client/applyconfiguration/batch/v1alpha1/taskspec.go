/*
Copyright The Volcano Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	batchv1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

// TaskSpecApplyConfiguration represents a declarative configuration of the TaskSpec type for use
// with apply.
type TaskSpecApplyConfiguration struct {
	Name           *string                             `json:"name,omitempty"`
	Replicas       *int32                              `json:"replicas,omitempty"`
	MinAvailable   *int32                              `json:"minAvailable,omitempty"`
	Template       *v1.PodTemplateSpec                 `json:"template,omitempty"`
	Policies       []LifecyclePolicyApplyConfiguration `json:"policies,omitempty"`
	TopologyPolicy *batchv1alpha1.NumaPolicy           `json:"topologyPolicy,omitempty"`
	MaxRetry       *int32                              `json:"maxRetry,omitempty"`
	DependsOn      *DependsOnApplyConfiguration        `json:"dependsOn,omitempty"`
}

// TaskSpecApplyConfiguration constructs a declarative configuration of the TaskSpec type for use with
// apply.
func TaskSpec() *TaskSpecApplyConfiguration {
	return &TaskSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *TaskSpecApplyConfiguration) WithName(value string) *TaskSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *TaskSpecApplyConfiguration) WithReplicas(value int32) *TaskSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithMinAvailable sets the MinAvailable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinAvailable field is set to the value of the last call.
func (b *TaskSpecApplyConfiguration) WithMinAvailable(value int32) *TaskSpecApplyConfiguration {
	b.MinAvailable = &value
	return b
}

// WithTemplate sets the Template field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Template field is set to the value of the last call.
func (b *TaskSpecApplyConfiguration) WithTemplate(value v1.PodTemplateSpec) *TaskSpecApplyConfiguration {
	b.Template = &value
	return b
}

// WithPolicies adds the given value to the Policies field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Policies field.
func (b *TaskSpecApplyConfiguration) WithPolicies(values ...*LifecyclePolicyApplyConfiguration) *TaskSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPolicies")
		}
		b.Policies = append(b.Policies, *values[i])
	}
	return b
}

// WithTopologyPolicy sets the TopologyPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TopologyPolicy field is set to the value of the last call.
func (b *TaskSpecApplyConfiguration) WithTopologyPolicy(value batchv1alpha1.NumaPolicy) *TaskSpecApplyConfiguration {
	b.TopologyPolicy = &value
	return b
}

// WithMaxRetry sets the MaxRetry field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MaxRetry field is set to the value of the last call.
func (b *TaskSpecApplyConfiguration) WithMaxRetry(value int32) *TaskSpecApplyConfiguration {
	b.MaxRetry = &value
	return b
}

// WithDependsOn sets the DependsOn field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DependsOn field is set to the value of the last call.
func (b *TaskSpecApplyConfiguration) WithDependsOn(value *DependsOnApplyConfiguration) *TaskSpecApplyConfiguration {
	b.DependsOn = value
	return b
}
