/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// StateMachineParameters defines the desired state of StateMachine
type StateMachineParameters struct {
	// Region is which region the StateMachine will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The Amazon States Language definition of the state machine. See Amazon States
	// Language (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).
	// +kubebuilder:validation:Required
	Definition *string `json:"definition"`
	// Defines what execution history events are logged and where they are logged.
	//
	// By default, the level is set to OFF. For more information see Log Levels
	// (https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html)
	// in the Step Functions User Guide.
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty"`
	// The name of the state machine.
	//
	// A name must not contain:
	//
	//    * white space
	//
	//    * brackets < > { } [ ]
	//
	//    * wildcard characters ? *
	//
	//    * special characters " # % \ ^ | ~ ` $ & , ; : /
	//
	//    * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable logging with CloudWatch Logs, the name should only contain 0-9,
	// A-Z, a-z, - and _.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// Set to true to publish the first version of the state machine during creation.
	// The default is false.
	Publish *bool `json:"publish,omitempty"`
	// Tags to be added when creating a state machine.
	//
	// An array of key-value pairs. For more information, see Using Cost Allocation
	// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
	// in the Amazon Web Services Billing and Cost Management User Guide, and Controlling
	// Access Using IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html).
	//
	// Tags may only contain Unicode letters, digits, white space, or these symbols:
	// _ . : / = + - @.
	Tags []*Tag `json:"tags,omitempty"`
	// Selects whether X-Ray tracing is enabled.
	TracingConfiguration *TracingConfiguration `json:"tracingConfiguration,omitempty"`
	// Sets description about the state machine version. You can only set the description
	// if the publish parameter is set to true. Otherwise, if you set versionDescription,
	// but publish to false, this API action throws ValidationException.
	VersionDescription           *string `json:"versionDescription,omitempty"`
	CustomStateMachineParameters `json:",inline"`
}

// StateMachineSpec defines the desired state of StateMachine
type StateMachineSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       StateMachineParameters `json:"forProvider"`
}

// StateMachineObservation defines the observed state of StateMachine
type StateMachineObservation struct {
	// The date the state machine is created.
	CreationDate *metav1.Time `json:"creationDate,omitempty"`
	// The Amazon Resource Name (ARN) that identifies the created state machine.
	StateMachineARN *string `json:"stateMachineARN,omitempty"`
	// The Amazon Resource Name (ARN) that identifies the created state machine
	// version. If you do not set the publish parameter to true, this field returns
	// null value.
	StateMachineVersionARN *string `json:"stateMachineVersionARN,omitempty"`

	CustomStateMachineObservation `json:",inline"`
}

// StateMachineStatus defines the observed state of StateMachine.
type StateMachineStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          StateMachineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StateMachine is the Schema for the StateMachines API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StateMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StateMachineSpec   `json:"spec"`
	Status            StateMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StateMachineList contains a list of StateMachines
type StateMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StateMachine `json:"items"`
}

// Repository type metadata.
var (
	StateMachineKind             = "StateMachine"
	StateMachineGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StateMachineKind}.String()
	StateMachineKindAPIVersion   = StateMachineKind + "." + GroupVersion.String()
	StateMachineGroupVersionKind = GroupVersion.WithKind(StateMachineKind)
)

func init() {
	SchemeBuilder.Register(&StateMachine{}, &StateMachineList{})
}
