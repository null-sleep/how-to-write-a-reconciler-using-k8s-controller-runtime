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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SpeakerSpec defines the desired state of Speaker
type SpeakerSpec struct {
	// Name of the Speaker.
	// +kubebuilder:validation:Type=string
	// +required
	Name string `json:"name"`

	// +kubebuilder:validation:Type=string
	Bio string `json:"bio,omitempty"`

	// Email of the Speaker
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="^[\w-\.]+@([\w-]+\.)+[\w-]{2,15}$"
	Email string `json:"email,omitempty"`
}

// SpeakerStatus defines the observed state of Speaker
type SpeakerStatus struct {
	// ObservedGeneration is the last observed generation of the Speaker object.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	ID int64 `json:"id,omitempty"`

	Conditions []metav1.Condition `json:"conditions:omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Speaker",type=string,JSONPath=`.spec.name`
// +kubebuilder:printcolumn:name="Email",type=string,JSONPath=`.spec.email`
// Speaker is the Schema for the speakers API
type Speaker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpeakerSpec   `json:"spec,omitempty"`
	Status SpeakerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SpeakerList contains a list of Speaker
type SpeakerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Speaker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Speaker{}, &SpeakerList{})
}
