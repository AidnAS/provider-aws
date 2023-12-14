// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ServerCertificateInitParameters struct {

	// encoded format.
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// encoded public key certificates
	// of the chain.
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format /cloudfront/your_path_here.
	// See IAM Identifiers for more details on IAM Paths.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ServerCertificateObservation struct {

	// The Amazon Resource Name (ARN) specifying the server certificate.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// encoded format.
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// encoded public key certificates
	// of the chain.
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// Date and time in RFC3339 format on which the certificate is set to expire.
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The unique Server Certificate name
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format /cloudfront/your_path_here.
	// See IAM Identifiers for more details on IAM Paths.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Date and time in RFC3339 format when the server certificate was uploaded.
	UploadDate *string `json:"uploadDate,omitempty" tf:"upload_date,omitempty"`
}

type ServerCertificateParameters struct {

	// encoded format.
	// +kubebuilder:validation:Optional
	CertificateBody *string `json:"certificateBody,omitempty" tf:"certificate_body,omitempty"`

	// encoded public key certificates
	// of the chain.
	// +kubebuilder:validation:Optional
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`

	// The IAM path for the server certificate.  If it is not
	// included, it defaults to a slash (/). If this certificate is for use with
	// AWS CloudFront, the path must be in format /cloudfront/your_path_here.
	// See IAM Identifiers for more details on IAM Paths.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// encoded format.
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ServerCertificateSpec defines the desired state of ServerCertificate
type ServerCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerCertificateParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ServerCertificateInitParameters `json:"initProvider,omitempty"`
}

// ServerCertificateStatus defines the observed state of ServerCertificate.
type ServerCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServerCertificate is the Schema for the ServerCertificates API. Provides an IAM Server Certificate
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ServerCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificateBody) || (has(self.initProvider) && has(self.initProvider.certificateBody))",message="spec.forProvider.certificateBody is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateKeySecretRef)",message="spec.forProvider.privateKeySecretRef is a required parameter"
	Spec   ServerCertificateSpec   `json:"spec"`
	Status ServerCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerCertificateList contains a list of ServerCertificates
type ServerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerCertificate `json:"items"`
}

// Repository type metadata.
var (
	ServerCertificate_Kind             = "ServerCertificate"
	ServerCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerCertificate_Kind}.String()
	ServerCertificate_KindAPIVersion   = ServerCertificate_Kind + "." + CRDGroupVersion.String()
	ServerCertificate_GroupVersionKind = CRDGroupVersion.WithKind(ServerCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerCertificate{}, &ServerCertificateList{})
}
