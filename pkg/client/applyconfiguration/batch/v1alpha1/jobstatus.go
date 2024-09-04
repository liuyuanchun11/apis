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

// JobStatusApplyConfiguration represents an declarative configuration of the JobStatus type for use
// with apply.
type JobStatusApplyConfiguration struct {
	State               *JobStateApplyConfiguration            `json:"state,omitempty"`
	MinAvailable        *int32                                 `json:"minAvailable,omitempty"`
	TaskStatusCount     map[string]TaskStateApplyConfiguration `json:"taskStatusCount,omitempty"`
	Pending             *int32                                 `json:"pending,omitempty"`
	Running             *int32                                 `json:"running,omitempty"`
	Succeeded           *int32                                 `json:"succeeded,omitempty"`
	Failed              *int32                                 `json:"failed,omitempty"`
	Terminating         *int32                                 `json:"terminating,omitempty"`
	Unknown             *int32                                 `json:"unknown,omitempty"`
	Version             *int32                                 `json:"version,omitempty"`
	RetryCount          *int32                                 `json:"retryCount,omitempty"`
	RunningDuration     *v1.Duration                           `json:"runningDuration,omitempty"`
	ControlledResources map[string]string                      `json:"controlledResources,omitempty"`
	Conditions          []JobConditionApplyConfiguration       `json:"conditions,omitempty"`
}

// JobStatusApplyConfiguration constructs an declarative configuration of the JobStatus type for use with
// apply.
func JobStatus() *JobStatusApplyConfiguration {
	return &JobStatusApplyConfiguration{}
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithState(value *JobStateApplyConfiguration) *JobStatusApplyConfiguration {
	b.State = value
	return b
}

// WithMinAvailable sets the MinAvailable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinAvailable field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithMinAvailable(value int32) *JobStatusApplyConfiguration {
	b.MinAvailable = &value
	return b
}

// WithTaskStatusCount puts the entries into the TaskStatusCount field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the TaskStatusCount field,
// overwriting an existing map entries in TaskStatusCount field with the same key.
func (b *JobStatusApplyConfiguration) WithTaskStatusCount(entries map[string]TaskStateApplyConfiguration) *JobStatusApplyConfiguration {
	if b.TaskStatusCount == nil && len(entries) > 0 {
		b.TaskStatusCount = make(map[string]TaskStateApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.TaskStatusCount[k] = v
	}
	return b
}

// WithPending sets the Pending field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Pending field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithPending(value int32) *JobStatusApplyConfiguration {
	b.Pending = &value
	return b
}

// WithRunning sets the Running field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Running field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithRunning(value int32) *JobStatusApplyConfiguration {
	b.Running = &value
	return b
}

// WithSucceeded sets the Succeeded field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Succeeded field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithSucceeded(value int32) *JobStatusApplyConfiguration {
	b.Succeeded = &value
	return b
}

// WithFailed sets the Failed field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Failed field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithFailed(value int32) *JobStatusApplyConfiguration {
	b.Failed = &value
	return b
}

// WithTerminating sets the Terminating field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Terminating field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithTerminating(value int32) *JobStatusApplyConfiguration {
	b.Terminating = &value
	return b
}

// WithUnknown sets the Unknown field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Unknown field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithUnknown(value int32) *JobStatusApplyConfiguration {
	b.Unknown = &value
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithVersion(value int32) *JobStatusApplyConfiguration {
	b.Version = &value
	return b
}

// WithRetryCount sets the RetryCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RetryCount field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithRetryCount(value int32) *JobStatusApplyConfiguration {
	b.RetryCount = &value
	return b
}

// WithRunningDuration sets the RunningDuration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunningDuration field is set to the value of the last call.
func (b *JobStatusApplyConfiguration) WithRunningDuration(value v1.Duration) *JobStatusApplyConfiguration {
	b.RunningDuration = &value
	return b
}

// WithControlledResources puts the entries into the ControlledResources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ControlledResources field,
// overwriting an existing map entries in ControlledResources field with the same key.
func (b *JobStatusApplyConfiguration) WithControlledResources(entries map[string]string) *JobStatusApplyConfiguration {
	if b.ControlledResources == nil && len(entries) > 0 {
		b.ControlledResources = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ControlledResources[k] = v
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *JobStatusApplyConfiguration) WithConditions(values ...*JobConditionApplyConfiguration) *JobStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
