/*
Copyright 2022.

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

// The [KubeBuilder documentation]
// (https://book.kubebuilder.io/reference/markers/crd-validation.html?highlight=required#crd-validation)
// states that if you include the kubebuilder:validation:Required marker at the
// packate level as shown below, all fields in the package are required unless
// explicitly marked optional.  However, it doesn't seem to make a difference.
// Without it, all fields that aren't explicitly marked as optional are treated
// as required anyway.
// +kubebuilder:validation:Required
package v1alpha1

// ClTracepointProgramInfo contains the Tracepoint program details
type ClTracepointProgramInfo struct {
	// links is the list of points to which the program is attached.  The list items
	// are optional and may be updated after the bpf program has been loaded
	// +optional
	Links []ClTracepointAttachInfo `json:"links,omitempty"`
}

type ClTracepointAttachInfo struct {
	// name refers to the name of a kernel tracepoint to attach the
	// bpf program to.
	// +kubebuilder:validation:Pattern="^[a-zA-Z][a-zA-Z0-9_]+."
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=64
	Name string `json:"name"`

	// This field is optional since we included the +optional and omitempty
	// markers.  It is not be included in the serialized output if empty because
	// we inluded omitembpty.  Best practice is to use both markers together.
	// When omitted, the default value is zero (the default for the int type).
	// +optional
	OptionalField1 int `json:"optionalField1,omitempty"`

	// This field is optional since we included the +optional marker.  However,
	// It is included in the serialized output even if empty. When omitted, the
	// default value is zero (the default for the int type).
	// +optional
	OptionalField2 int `json:"optionalField2"`

	// This field is optional since we included the omitempty marker.  It is
	// also not be included in the serialized output if empty. When omitted, the
	// default value is zero (the default for the int type).
	OptionalField3 int `json:"optionalField3,omitempty"`

	// This field is optional since we included the +optional and omitempty
	// markers.  It is not be included in the serialized output if empty because
	// we inluded omitempty.  Best practice is to use both markers together.
	// When omitted, the default value is nil (the default for pointers). It is
	// useful to use a pointer type when you want to distinguish between a field
	// that is not set and a field that is set to the default value.
	// +optional
	OptionalField4 *int `json:"optionalField4,omitempty"`

	// This field is optional since we included the +optional marker.
	// Apparently becuase it's a pointer field, it's omitted from the serialized
	// output even though we don't have omitempty. When omitted, the default
	// value is nil (the default for pointers). It is useful to use a pointer
	// type when you want to distinguish between a field that is not set and a
	// field that is set to the default value.
	// +optional
	OptionalField5 *int `json:"optionalField5"`

	// This field is optional since we included the omitempty marker.  It is
	// also not be included in the serialized output if empty. When omitted, the
	// default value is nil (the default for pointers). It is useful to use a
	// pointer type when you want to distinguish between a field that is not set
	// and a field that is set to the default value.
	OptionalField6 *int `json:"optionalField6,omitempty"`

	// This field is optional since we included the +optional marker.  It will
	// be included in the serialized output even if empty since we gave it a
	// default value.
	// +optional
	// +kubebuilder:default=42
	DefaultField int `json:"defaultField,omitempty"`

	// This field is required since we didn't make it optional.
	RequiredField1 int `json:"requiredField1"`

	// This field is required since we included the +required marker.
	// +required
	RequiredField2 int `json:"requiredField2"`

	// This field is required since we included the "validation:Required"
	// marker.
	// +kubebuilder:validation:Required
	RequiredField3 int `json:"requiredField3"`

	// You might think this field is required since we included the
	// "validation:Required" marker.  However, omitempty takes precedence,
	// making this field optional.
	// +kubebuilder:validation:Required
	NotReallyRequiredField1 int `json:"notReallyRequiredField1,omitempty"`
}

type ClTracepointProgramInfoState struct {
	// links is the list of attach points for the BPF program on the given node. Each entry
	// in *AttachInfoState represents a specific, unique attach point that is
	// derived from *AttachInfo by fully expanding any selectors.  Each entry
	// also contains information about the attach point required by the
	// reconciler
	// +optional
	Links []ClTracepointAttachInfoState `json:"links,omitempty"`
}

type ClTracepointAttachInfoState struct {
	AttachInfoStateCommon `json:",inline"`

	// The name of a kernel tracepoint to attach the bpf program to.
	Name string `json:"name"`
}
