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

// SecretParameters defines the desired state of Secret
type SecretParameters struct {
	// Region is which region the Secret will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// (Optional) Specifies a user-provided description of the secret.
	Description *string `json:"description,omitempty"`
	// (Optional) Specifies the ARN, Key ID, or alias of the AWS KMS customer master
	// key (CMK) to be used to encrypt the SecretString or SecretBinary values in
	// the versions stored in this secret.
	//
	// You can specify any of the supported ways to identify a AWS KMS key ID. If
	// you need to reference a CMK in a different account, you can use only the
	// key ARN or the alias ARN.
	//
	// If you don't specify this value, then Secrets Manager defaults to using the
	// AWS account's default CMK (the one named aws/secretsmanager). If a AWS KMS
	// CMK with that name doesn't yet exist, then Secrets Manager creates it for
	// you automatically the first time it needs to encrypt a version's SecretString
	// or SecretBinary fields.
	//
	// You can use the account default CMK to encrypt and decrypt only if you call
	// this operation using credentials from the same account that owns the secret.
	// If the secret resides in a different account, then you must create a custom
	// CMK and specify the ARN in this field.
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	// (Optional) Specifies a list of user-defined tags that are attached to the
	// secret. Each tag is a "Key" and "Value" pair of strings. This operation only
	// appends tags to the existing list of tags. To remove tags, you must use UntagResource.
	//
	//    * Secrets Manager tag key names are case sensitive. A tag with the key
	//    "ABC" is a different tag from one with key "abc".
	//
	//    * If you check tags in IAM policy Condition elements as part of your security
	//    strategy, then adding or removing a tag can change permissions. If the
	//    successful completion of this operation would result in you losing your
	//    permissions for this secret, then this operation is blocked and returns
	//    an Access Denied error.
	//
	// This parameter requires a JSON text string argument. For information on how
	// to format a JSON parameter for the various command line tool environments,
	// see Using JSON for Parameters (https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json)
	// in the AWS CLI User Guide. For example:
	//
	// [{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]
	//
	// If your command-line tool or SDK requires quotation marks around the parameter,
	// you should use single quotes to avoid confusion with the double quotes required
	// in the JSON text.
	//
	// The following basic restrictions apply to tags:
	//
	//    * Maximum number of tags per secret—50
	//
	//    * Maximum key length—127 Unicode characters in UTF-8
	//
	//    * Maximum value length—255 Unicode characters in UTF-8
	//
	//    * Tag keys and values are case sensitive.
	//
	//    * Do not use the aws: prefix in your tag names or values because AWS reserves
	//    it for AWS use. You can't edit or delete tag names or values with this
	//    prefix. Tags with this prefix do not count against your tags per secret
	//    limit.
	//
	//    * If you use your tagging schema across multiple services and resources,
	//    remember other services might have restrictions on allowed characters.
	//    Generally allowed characters: letters, spaces, and numbers representable
	//    in UTF-8, plus the following special characters: + - = . _ : / @.
	Tags                   []*Tag `json:"tags,omitempty"`
	CustomSecretParameters `json:",inline"`
}

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SecretParameters `json:"forProvider"`
}

// SecretObservation defines the observed state of Secret
type SecretObservation struct {
	// The Amazon Resource Name (ARN) of the secret that you just created.
	//
	// Secrets Manager automatically adds several random characters to the name
	// at the end of the ARN when you initially create a secret. This affects only
	// the ARN and not the actual friendly name. This ensures that if you create
	// a new secret with the same name as an old secret that you previously deleted,
	// then users with access to the old secret don't automatically get access to
	// the new secret because the ARNs are different.
	ARN *string `json:"arn,omitempty"`
}

// SecretStatus defines the observed state of Secret.
type SecretStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Secret is the Schema for the Secrets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:deprecatedversion:warning="Please use v1beta1 version of this resource."
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretSpec   `json:"spec"`
	Status            SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

// Repository type metadata.
var (
	SecretKind             = "Secret"
	SecretGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretKind}.String()
	SecretKindAPIVersion   = SecretKind + "." + GroupVersion.String()
	SecretGroupVersionKind = GroupVersion.WithKind(SecretKind)
)

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}
