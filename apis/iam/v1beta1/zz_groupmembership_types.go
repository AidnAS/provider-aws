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

type GroupMembershipInitParameters struct {

	// The name to identify the Group Membership
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GroupMembershipObservation struct {

	// –  The IAM Group name to attach the list of users to
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name to identify the Group Membership
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of IAM User names to associate with the Group
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type GroupMembershipParameters struct {

	// –  The IAM Group name to attach the list of users to
	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Reference to a Group to populate group.
	// +kubebuilder:validation:Optional
	GroupRef *v1.Reference `json:"groupRef,omitempty" tf:"-"`

	// Selector for a Group to populate group.
	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// The name to identify the Group Membership
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// References to User to populate users.
	// +kubebuilder:validation:Optional
	UserRefs []v1.Reference `json:"userRefs,omitempty" tf:"-"`

	// Selector for a list of User to populate users.
	// +kubebuilder:validation:Optional
	UserSelector *v1.Selector `json:"userSelector,omitempty" tf:"-"`

	// A list of IAM User names to associate with the Group
	// +crossplane:generate:reference:type=User
	// +crossplane:generate:reference:refFieldName=UserRefs
	// +crossplane:generate:reference:selectorFieldName=UserSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

// GroupMembershipSpec defines the desired state of GroupMembership
type GroupMembershipSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupMembershipParameters `json:"forProvider"`
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
	InitProvider GroupMembershipInitParameters `json:"initProvider,omitempty"`
}

// GroupMembershipStatus defines the observed state of GroupMembership.
type GroupMembershipStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupMembershipObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembership is the Schema for the GroupMemberships API. Provides a top level resource to manage IAM Group membership for IAM Users.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type GroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   GroupMembershipSpec   `json:"spec"`
	Status GroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupMembershipList contains a list of GroupMemberships
type GroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupMembership `json:"items"`
}

// Repository type metadata.
var (
	GroupMembership_Kind             = "GroupMembership"
	GroupMembership_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupMembership_Kind}.String()
	GroupMembership_KindAPIVersion   = GroupMembership_Kind + "." + CRDGroupVersion.String()
	GroupMembership_GroupVersionKind = CRDGroupVersion.WithKind(GroupMembership_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupMembership{}, &GroupMembershipList{})
}
