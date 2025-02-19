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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReplicatedJobStatusApplyConfiguration represents a declarative configuration of the ReplicatedJobStatus type for use
// with apply.
type ReplicatedJobStatusApplyConfiguration struct {
	Name               *string  `json:"name,omitempty"`
	Ready              *int32   `json:"ready,omitempty"`
	Succeeded          *int32   `json:"succeeded,omitempty"`
	Failed             *int32   `json:"failed,omitempty"`
	Active             *int32   `json:"active,omitempty"`
	Pending            *int32   `json:"pending,omitempty"`
	LastTransitionTime *v1.Time `json:"lastTransitionTime,omitempty"`
}

// ReplicatedJobStatusApplyConfiguration constructs a declarative configuration of the ReplicatedJobStatus type for use with
// apply.
func ReplicatedJobStatus() *ReplicatedJobStatusApplyConfiguration {
	return &ReplicatedJobStatusApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *ReplicatedJobStatusApplyConfiguration) WithName(value string) *ReplicatedJobStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithReady sets the Ready field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ready field is set to the value of the last call.
func (b *ReplicatedJobStatusApplyConfiguration) WithReady(value int32) *ReplicatedJobStatusApplyConfiguration {
	b.Ready = &value
	return b
}

// WithSucceeded sets the Succeeded field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Succeeded field is set to the value of the last call.
func (b *ReplicatedJobStatusApplyConfiguration) WithSucceeded(value int32) *ReplicatedJobStatusApplyConfiguration {
	b.Succeeded = &value
	return b
}

// WithFailed sets the Failed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Failed field is set to the value of the last call.
func (b *ReplicatedJobStatusApplyConfiguration) WithFailed(value int32) *ReplicatedJobStatusApplyConfiguration {
	b.Failed = &value
	return b
}

// WithActive sets the Active field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Active field is set to the value of the last call.
func (b *ReplicatedJobStatusApplyConfiguration) WithActive(value int32) *ReplicatedJobStatusApplyConfiguration {
	b.Active = &value
	return b
}

// WithPending sets the Pending field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pending field is set to the value of the last call.
func (b *ReplicatedJobStatusApplyConfiguration) WithPending(value int32) *ReplicatedJobStatusApplyConfiguration {
	b.Pending = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *ReplicatedJobStatusApplyConfiguration) WithLastTransitionTime(value v1.Time) *ReplicatedJobStatusApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}
