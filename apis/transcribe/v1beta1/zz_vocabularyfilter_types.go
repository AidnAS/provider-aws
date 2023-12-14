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

type VocabularyFilterInitParameters struct {

	// The language code you selected for your vocabulary filter. Refer to the supported languages page for accepted codes.
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with words argument.
	VocabularyFilterFileURI *string `json:"vocabularyFilterFileUri,omitempty" tf:"vocabulary_filter_file_uri,omitempty"`

	// - A list of terms to include in the vocabulary. Conflicts with vocabulary_filter_file_uri argument.
	Words []*string `json:"words,omitempty" tf:"words,omitempty"`
}

type VocabularyFilterObservation struct {

	// ARN of the VocabularyFilter.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Generated download URI.
	DownloadURI *string `json:"downloadUri,omitempty" tf:"download_uri,omitempty"`

	// VocabularyFilter name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The language code you selected for your vocabulary filter. Refer to the supported languages page for accepted codes.
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with words argument.
	VocabularyFilterFileURI *string `json:"vocabularyFilterFileUri,omitempty" tf:"vocabulary_filter_file_uri,omitempty"`

	// - A list of terms to include in the vocabulary. Conflicts with vocabulary_filter_file_uri argument.
	Words []*string `json:"words,omitempty" tf:"words,omitempty"`
}

type VocabularyFilterParameters struct {

	// The language code you selected for your vocabulary filter. Refer to the supported languages page for accepted codes.
	// +kubebuilder:validation:Optional
	LanguageCode *string `json:"languageCode,omitempty" tf:"language_code,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The Amazon S3 location (URI) of the text file that contains your custom VocabularyFilter. Conflicts with words argument.
	// +kubebuilder:validation:Optional
	VocabularyFilterFileURI *string `json:"vocabularyFilterFileUri,omitempty" tf:"vocabulary_filter_file_uri,omitempty"`

	// - A list of terms to include in the vocabulary. Conflicts with vocabulary_filter_file_uri argument.
	// +kubebuilder:validation:Optional
	Words []*string `json:"words,omitempty" tf:"words,omitempty"`
}

// VocabularyFilterSpec defines the desired state of VocabularyFilter
type VocabularyFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VocabularyFilterParameters `json:"forProvider"`
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
	InitProvider VocabularyFilterInitParameters `json:"initProvider,omitempty"`
}

// VocabularyFilterStatus defines the observed state of VocabularyFilter.
type VocabularyFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VocabularyFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VocabularyFilter is the Schema for the VocabularyFilters API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VocabularyFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.languageCode) || (has(self.initProvider) && has(self.initProvider.languageCode))",message="spec.forProvider.languageCode is a required parameter"
	Spec   VocabularyFilterSpec   `json:"spec"`
	Status VocabularyFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VocabularyFilterList contains a list of VocabularyFilters
type VocabularyFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VocabularyFilter `json:"items"`
}

// Repository type metadata.
var (
	VocabularyFilter_Kind             = "VocabularyFilter"
	VocabularyFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VocabularyFilter_Kind}.String()
	VocabularyFilter_KindAPIVersion   = VocabularyFilter_Kind + "." + CRDGroupVersion.String()
	VocabularyFilter_GroupVersionKind = CRDGroupVersion.WithKind(VocabularyFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&VocabularyFilter{}, &VocabularyFilterList{})
}
