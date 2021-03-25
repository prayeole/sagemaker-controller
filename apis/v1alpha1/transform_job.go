// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TransformJobSpec defines the desired state of TransformJob
type TransformJobSpec struct {
	BatchStrategy           *string            `json:"batchStrategy,omitempty"`
	DataProcessing          *DataProcessing    `json:"dataProcessing,omitempty"`
	Environment             map[string]*string `json:"environment,omitempty"`
	ExperimentConfig        *ExperimentConfig  `json:"experimentConfig,omitempty"`
	MaxConcurrentTransforms *int64             `json:"maxConcurrentTransforms,omitempty"`
	MaxPayloadInMB          *int64             `json:"maxPayloadInMB,omitempty"`
	ModelClientConfig       *ModelClientConfig `json:"modelClientConfig,omitempty"`
	// +kubebuilder:validation:Required
	ModelName *string `json:"modelName"`
	Tags      []*Tag  `json:"tags,omitempty"`
	// +kubebuilder:validation:Required
	TransformInput *TransformInput `json:"transformInput"`
	// +kubebuilder:validation:Required
	TransformJobName *string `json:"transformJobName"`
	// +kubebuilder:validation:Required
	TransformOutput *TransformOutput `json:"transformOutput"`
	// +kubebuilder:validation:Required
	TransformResources *TransformResources `json:"transformResources"`
}

// TransformJobStatus defines the observed state of TransformJob
type TransformJobStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions         []*ackv1alpha1.Condition `json:"conditions"`
	FailureReason      *string                  `json:"failureReason,omitempty"`
	TransformJobStatus *string                  `json:"transformJobStatus,omitempty"`
}

// TransformJob is the Schema for the TransformJobs API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="FailureReason",type=string,JSONPath=`.status.failureReason`
// +kubebuilder:printcolumn:name="TransformJobStatus",type=string,JSONPath=`.status.transformJobStatus`
type TransformJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransformJobSpec   `json:"spec,omitempty"`
	Status            TransformJobStatus `json:"status,omitempty"`
}

// TransformJobList contains a list of TransformJob
// +kubebuilder:object:root=true
type TransformJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransformJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TransformJob{}, &TransformJobList{})
}