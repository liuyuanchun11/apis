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

// HyperJobSpecApplyConfiguration represents a declarative configuration of the HyperJobSpec type for use
// with apply.
type HyperJobSpecApplyConfiguration struct {
	ReplicatedJobs []ReplicatedJobApplyConfiguration `json:"replicatedJobs,omitempty"`
	MinAvailable   *int32                            `json:"minAvailable,omitempty"`
	Plugins        map[string][]string               `json:"plugins,omitempty"`
	SuccessPolicy  *SuccessPolicyApplyConfiguration  `json:"successPolicy,omitempty"`
	FailurePolicy  *FailurePolicyApplyConfiguration  `json:"failurePolicy,omitempty"`
	StartupPolicy  *StartupPolicyApplyConfiguration  `json:"startupPolicy,omitempty"`
}

// HyperJobSpecApplyConfiguration constructs a declarative configuration of the HyperJobSpec type for use with
// apply.
func HyperJobSpec() *HyperJobSpecApplyConfiguration {
	return &HyperJobSpecApplyConfiguration{}
}

// WithReplicatedJobs adds the given value to the ReplicatedJobs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ReplicatedJobs field.
func (b *HyperJobSpecApplyConfiguration) WithReplicatedJobs(values ...*ReplicatedJobApplyConfiguration) *HyperJobSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithReplicatedJobs")
		}
		b.ReplicatedJobs = append(b.ReplicatedJobs, *values[i])
	}
	return b
}

// WithMinAvailable sets the MinAvailable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MinAvailable field is set to the value of the last call.
func (b *HyperJobSpecApplyConfiguration) WithMinAvailable(value int32) *HyperJobSpecApplyConfiguration {
	b.MinAvailable = &value
	return b
}

// WithPlugins puts the entries into the Plugins field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Plugins field,
// overwriting an existing map entries in Plugins field with the same key.
func (b *HyperJobSpecApplyConfiguration) WithPlugins(entries map[string][]string) *HyperJobSpecApplyConfiguration {
	if b.Plugins == nil && len(entries) > 0 {
		b.Plugins = make(map[string][]string, len(entries))
	}
	for k, v := range entries {
		b.Plugins[k] = v
	}
	return b
}

// WithSuccessPolicy sets the SuccessPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SuccessPolicy field is set to the value of the last call.
func (b *HyperJobSpecApplyConfiguration) WithSuccessPolicy(value *SuccessPolicyApplyConfiguration) *HyperJobSpecApplyConfiguration {
	b.SuccessPolicy = value
	return b
}

// WithFailurePolicy sets the FailurePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailurePolicy field is set to the value of the last call.
func (b *HyperJobSpecApplyConfiguration) WithFailurePolicy(value *FailurePolicyApplyConfiguration) *HyperJobSpecApplyConfiguration {
	b.FailurePolicy = value
	return b
}

// WithStartupPolicy sets the StartupPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartupPolicy field is set to the value of the last call.
func (b *HyperJobSpecApplyConfiguration) WithStartupPolicy(value *StartupPolicyApplyConfiguration) *HyperJobSpecApplyConfiguration {
	b.StartupPolicy = value
	return b
}
