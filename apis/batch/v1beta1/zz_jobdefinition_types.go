// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ContainersInitParameters struct {

	// An array of arguments to the entrypoint. If this isn't specified, the CMD of the container image is used. This corresponds to the args member in the Entrypoint portion of the Pod in Kubernetes. Environment variable references are expanded using the container's environment.
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// The entrypoint for the container. This isn't run within a shell. If this isn't specified, the ENTRYPOINT of the container image is used. Environment variable references are expanded using the container's environment.
	Command []*string `json:"command,omitempty" tf:"command,omitempty"`

	// The environment variables to pass to a container. See EKS Environment below.
	Env []EnvInitParameters `json:"env,omitempty" tf:"env,omitempty"`

	// The Docker image used to start the container.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// The image pull policy for the container. Supported values are Always, IfNotPresent, and Never.
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty" tf:"image_pull_policy,omitempty"`

	// The name of the container. If the name isn't specified, the default name "Default" is used. Each container in a pod must have a unique name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type and amount of resources to assign to a container. The supported resources include memory, cpu, and nvidia.com/gpu.
	Resources []ResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// The security context for a job.
	SecurityContext []SecurityContextInitParameters `json:"securityContext,omitempty" tf:"security_context,omitempty"`

	// The volume mounts for the container.
	VolumeMounts []VolumeMountsInitParameters `json:"volumeMounts,omitempty" tf:"volume_mounts,omitempty"`
}

type ContainersObservation struct {

	// An array of arguments to the entrypoint. If this isn't specified, the CMD of the container image is used. This corresponds to the args member in the Entrypoint portion of the Pod in Kubernetes. Environment variable references are expanded using the container's environment.
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// The entrypoint for the container. This isn't run within a shell. If this isn't specified, the ENTRYPOINT of the container image is used. Environment variable references are expanded using the container's environment.
	Command []*string `json:"command,omitempty" tf:"command,omitempty"`

	// The environment variables to pass to a container. See EKS Environment below.
	Env []EnvObservation `json:"env,omitempty" tf:"env,omitempty"`

	// The Docker image used to start the container.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// The image pull policy for the container. Supported values are Always, IfNotPresent, and Never.
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty" tf:"image_pull_policy,omitempty"`

	// The name of the container. If the name isn't specified, the default name "Default" is used. Each container in a pod must have a unique name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type and amount of resources to assign to a container. The supported resources include memory, cpu, and nvidia.com/gpu.
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// The security context for a job.
	SecurityContext []SecurityContextObservation `json:"securityContext,omitempty" tf:"security_context,omitempty"`

	// The volume mounts for the container.
	VolumeMounts []VolumeMountsObservation `json:"volumeMounts,omitempty" tf:"volume_mounts,omitempty"`
}

type ContainersParameters struct {

	// An array of arguments to the entrypoint. If this isn't specified, the CMD of the container image is used. This corresponds to the args member in the Entrypoint portion of the Pod in Kubernetes. Environment variable references are expanded using the container's environment.
	// +kubebuilder:validation:Optional
	Args []*string `json:"args,omitempty" tf:"args,omitempty"`

	// The entrypoint for the container. This isn't run within a shell. If this isn't specified, the ENTRYPOINT of the container image is used. Environment variable references are expanded using the container's environment.
	// +kubebuilder:validation:Optional
	Command []*string `json:"command,omitempty" tf:"command,omitempty"`

	// The environment variables to pass to a container. See EKS Environment below.
	// +kubebuilder:validation:Optional
	Env []EnvParameters `json:"env,omitempty" tf:"env,omitempty"`

	// The Docker image used to start the container.
	// +kubebuilder:validation:Optional
	Image *string `json:"image" tf:"image,omitempty"`

	// The image pull policy for the container. Supported values are Always, IfNotPresent, and Never.
	// +kubebuilder:validation:Optional
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty" tf:"image_pull_policy,omitempty"`

	// The name of the container. If the name isn't specified, the default name "Default" is used. Each container in a pod must have a unique name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The type and amount of resources to assign to a container. The supported resources include memory, cpu, and nvidia.com/gpu.
	// +kubebuilder:validation:Optional
	Resources []ResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// The security context for a job.
	// +kubebuilder:validation:Optional
	SecurityContext []SecurityContextParameters `json:"securityContext,omitempty" tf:"security_context,omitempty"`

	// The volume mounts for the container.
	// +kubebuilder:validation:Optional
	VolumeMounts []VolumeMountsParameters `json:"volumeMounts,omitempty" tf:"volume_mounts,omitempty"`
}

type EksPropertiesInitParameters struct {

	// The properties for the Kubernetes pod resources of a job. See pod_properties below.
	PodProperties []PodPropertiesInitParameters `json:"podProperties,omitempty" tf:"pod_properties,omitempty"`
}

type EksPropertiesObservation struct {

	// The properties for the Kubernetes pod resources of a job. See pod_properties below.
	PodProperties []PodPropertiesObservation `json:"podProperties,omitempty" tf:"pod_properties,omitempty"`
}

type EksPropertiesParameters struct {

	// The properties for the Kubernetes pod resources of a job. See pod_properties below.
	// +kubebuilder:validation:Optional
	PodProperties []PodPropertiesParameters `json:"podProperties" tf:"pod_properties,omitempty"`
}

type EmptyDirInitParameters struct {

	// The medium to store the volume. The default value is an empty string, which uses the storage of the node.
	Medium *string `json:"medium,omitempty" tf:"medium,omitempty"`

	// The maximum size of the volume. By default, there's no maximum size defined.
	SizeLimit *string `json:"sizeLimit,omitempty" tf:"size_limit,omitempty"`
}

type EmptyDirObservation struct {

	// The medium to store the volume. The default value is an empty string, which uses the storage of the node.
	Medium *string `json:"medium,omitempty" tf:"medium,omitempty"`

	// The maximum size of the volume. By default, there's no maximum size defined.
	SizeLimit *string `json:"sizeLimit,omitempty" tf:"size_limit,omitempty"`
}

type EmptyDirParameters struct {

	// The medium to store the volume. The default value is an empty string, which uses the storage of the node.
	// +kubebuilder:validation:Optional
	Medium *string `json:"medium,omitempty" tf:"medium,omitempty"`

	// The maximum size of the volume. By default, there's no maximum size defined.
	// +kubebuilder:validation:Optional
	SizeLimit *string `json:"sizeLimit" tf:"size_limit,omitempty"`
}

type EnvInitParameters struct {

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the environment variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnvObservation struct {

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of the environment variable.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type EnvParameters struct {

	// Specifies the name of the job definition.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value of the environment variable.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type EvaluateOnExitInitParameters struct {

	// Specifies the action to take if all of the specified conditions are met. The values are not case sensitive. Valid values: retry, exit.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// A glob pattern to match against the decimal representation of the exit code returned for a job.
	OnExitCode *string `json:"onExitCode,omitempty" tf:"on_exit_code,omitempty"`

	// A glob pattern to match against the reason returned for a job.
	OnReason *string `json:"onReason,omitempty" tf:"on_reason,omitempty"`

	// A glob pattern to match against the status reason returned for a job.
	OnStatusReason *string `json:"onStatusReason,omitempty" tf:"on_status_reason,omitempty"`
}

type EvaluateOnExitObservation struct {

	// Specifies the action to take if all of the specified conditions are met. The values are not case sensitive. Valid values: retry, exit.
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// A glob pattern to match against the decimal representation of the exit code returned for a job.
	OnExitCode *string `json:"onExitCode,omitempty" tf:"on_exit_code,omitempty"`

	// A glob pattern to match against the reason returned for a job.
	OnReason *string `json:"onReason,omitempty" tf:"on_reason,omitempty"`

	// A glob pattern to match against the status reason returned for a job.
	OnStatusReason *string `json:"onStatusReason,omitempty" tf:"on_status_reason,omitempty"`
}

type EvaluateOnExitParameters struct {

	// Specifies the action to take if all of the specified conditions are met. The values are not case sensitive. Valid values: retry, exit.
	// +kubebuilder:validation:Optional
	Action *string `json:"action" tf:"action,omitempty"`

	// A glob pattern to match against the decimal representation of the exit code returned for a job.
	// +kubebuilder:validation:Optional
	OnExitCode *string `json:"onExitCode,omitempty" tf:"on_exit_code,omitempty"`

	// A glob pattern to match against the reason returned for a job.
	// +kubebuilder:validation:Optional
	OnReason *string `json:"onReason,omitempty" tf:"on_reason,omitempty"`

	// A glob pattern to match against the status reason returned for a job.
	// +kubebuilder:validation:Optional
	OnStatusReason *string `json:"onStatusReason,omitempty" tf:"on_status_reason,omitempty"`
}

type HostPathInitParameters struct {

	// The path of the file or directory on the host to mount into containers on the pod.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type HostPathObservation struct {

	// The path of the file or directory on the host to mount into containers on the pod.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type HostPathParameters struct {

	// The path of the file or directory on the host to mount into containers on the pod.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`
}

type JobDefinitionInitParameters struct {

	// A valid container properties provided as a single valid JSON document. This parameter is only valid if the type parameter is container.
	ContainerProperties *string `json:"containerProperties,omitempty" tf:"container_properties,omitempty"`

	// When updating a job definition a new revision is created. This parameter determines if the previous version is deregistered (INACTIVE) or left  ACTIVE. Defaults to true.
	DeregisterOnNewRevision *bool `json:"deregisterOnNewRevision,omitempty" tf:"deregister_on_new_revision,omitempty"`

	// A valid eks properties. This parameter is only valid if the type parameter is container.
	EksProperties []EksPropertiesInitParameters `json:"eksProperties,omitempty" tf:"eks_properties,omitempty"`

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A valid node properties provided as a single valid JSON document. This parameter is required if the type parameter is multinode.
	NodeProperties *string `json:"nodeProperties,omitempty" tf:"node_properties,omitempty"`

	// Specifies the parameter substitution placeholders to set in the job definition.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The platform capabilities required by the job definition. If no value is specified, it defaults to EC2. To run the job on Fargate resources, specify FARGATE.
	// +listType=set
	PlatformCapabilities []*string `json:"platformCapabilities,omitempty" tf:"platform_capabilities,omitempty"`

	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is false.
	PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition. Maximum number of retry_strategy is 1.  Defined below.
	RetryStrategy []RetryStrategyInitParameters `json:"retryStrategy,omitempty" tf:"retry_strategy,omitempty"`

	// The scheduling priority of the job definition. This only affects jobs in job queues with a fair share policy. Jobs with a higher scheduling priority are scheduled before jobs with a lower scheduling priority. Allowed values 0 through 9999.
	SchedulingPriority *float64 `json:"schedulingPriority,omitempty" tf:"scheduling_priority,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of timeout is 1. Defined below.
	Timeout []TimeoutInitParameters `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of job definition. Must be container or multinode.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type JobDefinitionObservation struct {

	// The Amazon Resource Name of the job definition, includes revision (:#).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ARN without the revision number.
	ArnPrefix *string `json:"arnPrefix,omitempty" tf:"arn_prefix,omitempty"`

	// A valid container properties provided as a single valid JSON document. This parameter is only valid if the type parameter is container.
	ContainerProperties *string `json:"containerProperties,omitempty" tf:"container_properties,omitempty"`

	// When updating a job definition a new revision is created. This parameter determines if the previous version is deregistered (INACTIVE) or left  ACTIVE. Defaults to true.
	DeregisterOnNewRevision *bool `json:"deregisterOnNewRevision,omitempty" tf:"deregister_on_new_revision,omitempty"`

	// A valid eks properties. This parameter is only valid if the type parameter is container.
	EksProperties []EksPropertiesObservation `json:"eksProperties,omitempty" tf:"eks_properties,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A valid node properties provided as a single valid JSON document. This parameter is required if the type parameter is multinode.
	NodeProperties *string `json:"nodeProperties,omitempty" tf:"node_properties,omitempty"`

	// Specifies the parameter substitution placeholders to set in the job definition.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The platform capabilities required by the job definition. If no value is specified, it defaults to EC2. To run the job on Fargate resources, specify FARGATE.
	// +listType=set
	PlatformCapabilities []*string `json:"platformCapabilities,omitempty" tf:"platform_capabilities,omitempty"`

	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is false.
	PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition. Maximum number of retry_strategy is 1.  Defined below.
	RetryStrategy []RetryStrategyObservation `json:"retryStrategy,omitempty" tf:"retry_strategy,omitempty"`

	// The revision of the job definition.
	Revision *float64 `json:"revision,omitempty" tf:"revision,omitempty"`

	// The scheduling priority of the job definition. This only affects jobs in job queues with a fair share policy. Jobs with a higher scheduling priority are scheduled before jobs with a lower scheduling priority. Allowed values 0 through 9999.
	SchedulingPriority *float64 `json:"schedulingPriority,omitempty" tf:"scheduling_priority,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of timeout is 1. Defined below.
	Timeout []TimeoutObservation `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of job definition. Must be container or multinode.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type JobDefinitionParameters struct {

	// A valid container properties provided as a single valid JSON document. This parameter is only valid if the type parameter is container.
	// +kubebuilder:validation:Optional
	ContainerProperties *string `json:"containerProperties,omitempty" tf:"container_properties,omitempty"`

	// When updating a job definition a new revision is created. This parameter determines if the previous version is deregistered (INACTIVE) or left  ACTIVE. Defaults to true.
	// +kubebuilder:validation:Optional
	DeregisterOnNewRevision *bool `json:"deregisterOnNewRevision,omitempty" tf:"deregister_on_new_revision,omitempty"`

	// A valid eks properties. This parameter is only valid if the type parameter is container.
	// +kubebuilder:validation:Optional
	EksProperties []EksPropertiesParameters `json:"eksProperties,omitempty" tf:"eks_properties,omitempty"`

	// Specifies the name of the job definition.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A valid node properties provided as a single valid JSON document. This parameter is required if the type parameter is multinode.
	// +kubebuilder:validation:Optional
	NodeProperties *string `json:"nodeProperties,omitempty" tf:"node_properties,omitempty"`

	// Specifies the parameter substitution placeholders to set in the job definition.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The platform capabilities required by the job definition. If no value is specified, it defaults to EC2. To run the job on Fargate resources, specify FARGATE.
	// +kubebuilder:validation:Optional
	// +listType=set
	PlatformCapabilities []*string `json:"platformCapabilities,omitempty" tf:"platform_capabilities,omitempty"`

	// Specifies whether to propagate the tags from the job definition to the corresponding Amazon ECS task. Default is false.
	// +kubebuilder:validation:Optional
	PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition. Maximum number of retry_strategy is 1.  Defined below.
	// +kubebuilder:validation:Optional
	RetryStrategy []RetryStrategyParameters `json:"retryStrategy,omitempty" tf:"retry_strategy,omitempty"`

	// The scheduling priority of the job definition. This only affects jobs in job queues with a fair share policy. Jobs with a higher scheduling priority are scheduled before jobs with a lower scheduling priority. Allowed values 0 through 9999.
	// +kubebuilder:validation:Optional
	SchedulingPriority *float64 `json:"schedulingPriority,omitempty" tf:"scheduling_priority,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of timeout is 1. Defined below.
	// +kubebuilder:validation:Optional
	Timeout []TimeoutParameters `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// The type of job definition. Must be container or multinode.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MetadataInitParameters struct {

	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type MetadataObservation struct {

	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type MetadataParameters struct {

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`
}

type PodPropertiesInitParameters struct {

	// The properties of the container that's used on the Amazon EKS pod. See containers below.
	Containers []ContainersInitParameters `json:"containers,omitempty" tf:"containers,omitempty"`

	// The DNS policy for the pod. The default value is ClusterFirst. If the host_network argument is not specified, the default is ClusterFirstWithHostNet. ClusterFirst indicates that any DNS query that does not match the configured cluster domain suffix is forwarded to the upstream nameserver inherited from the node. For more information, see Pod's DNS policy in the Kubernetes documentation.
	DNSPolicy *string `json:"dnsPolicy,omitempty" tf:"dns_policy,omitempty"`

	// Indicates if the pod uses the hosts' network IP address. The default value is true. Setting this to false enables the Kubernetes pod networking model. Most AWS Batch workloads are egress-only and don't require the overhead of IP allocation for each pod for incoming connections.
	HostNetwork *bool `json:"hostNetwork,omitempty" tf:"host_network,omitempty"`

	// Metadata about the Kubernetes pod.
	Metadata []MetadataInitParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the service account that's used to run the pod.
	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name,omitempty"`

	// Specifies the volumes for a job definition that uses Amazon EKS resources. AWS Batch supports emptyDir, hostPath, and secret volume types.
	Volumes []VolumesInitParameters `json:"volumes,omitempty" tf:"volumes,omitempty"`
}

type PodPropertiesObservation struct {

	// The properties of the container that's used on the Amazon EKS pod. See containers below.
	Containers []ContainersObservation `json:"containers,omitempty" tf:"containers,omitempty"`

	// The DNS policy for the pod. The default value is ClusterFirst. If the host_network argument is not specified, the default is ClusterFirstWithHostNet. ClusterFirst indicates that any DNS query that does not match the configured cluster domain suffix is forwarded to the upstream nameserver inherited from the node. For more information, see Pod's DNS policy in the Kubernetes documentation.
	DNSPolicy *string `json:"dnsPolicy,omitempty" tf:"dns_policy,omitempty"`

	// Indicates if the pod uses the hosts' network IP address. The default value is true. Setting this to false enables the Kubernetes pod networking model. Most AWS Batch workloads are egress-only and don't require the overhead of IP allocation for each pod for incoming connections.
	HostNetwork *bool `json:"hostNetwork,omitempty" tf:"host_network,omitempty"`

	// Metadata about the Kubernetes pod.
	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the service account that's used to run the pod.
	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name,omitempty"`

	// Specifies the volumes for a job definition that uses Amazon EKS resources. AWS Batch supports emptyDir, hostPath, and secret volume types.
	Volumes []VolumesObservation `json:"volumes,omitempty" tf:"volumes,omitempty"`
}

type PodPropertiesParameters struct {

	// The properties of the container that's used on the Amazon EKS pod. See containers below.
	// +kubebuilder:validation:Optional
	Containers []ContainersParameters `json:"containers" tf:"containers,omitempty"`

	// The DNS policy for the pod. The default value is ClusterFirst. If the host_network argument is not specified, the default is ClusterFirstWithHostNet. ClusterFirst indicates that any DNS query that does not match the configured cluster domain suffix is forwarded to the upstream nameserver inherited from the node. For more information, see Pod's DNS policy in the Kubernetes documentation.
	// +kubebuilder:validation:Optional
	DNSPolicy *string `json:"dnsPolicy,omitempty" tf:"dns_policy,omitempty"`

	// Indicates if the pod uses the hosts' network IP address. The default value is true. Setting this to false enables the Kubernetes pod networking model. Most AWS Batch workloads are egress-only and don't require the overhead of IP allocation for each pod for incoming connections.
	// +kubebuilder:validation:Optional
	HostNetwork *bool `json:"hostNetwork,omitempty" tf:"host_network,omitempty"`

	// Metadata about the Kubernetes pod.
	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// The name of the service account that's used to run the pod.
	// +kubebuilder:validation:Optional
	ServiceAccountName *string `json:"serviceAccountName,omitempty" tf:"service_account_name,omitempty"`

	// Specifies the volumes for a job definition that uses Amazon EKS resources. AWS Batch supports emptyDir, hostPath, and secret volume types.
	// +kubebuilder:validation:Optional
	Volumes []VolumesParameters `json:"volumes,omitempty" tf:"volumes,omitempty"`
}

type ResourcesInitParameters struct {

	// +mapType=granular
	Limits map[string]*string `json:"limits,omitempty" tf:"limits,omitempty"`

	// +mapType=granular
	Requests map[string]*string `json:"requests,omitempty" tf:"requests,omitempty"`
}

type ResourcesObservation struct {

	// +mapType=granular
	Limits map[string]*string `json:"limits,omitempty" tf:"limits,omitempty"`

	// +mapType=granular
	Requests map[string]*string `json:"requests,omitempty" tf:"requests,omitempty"`
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Limits map[string]*string `json:"limits,omitempty" tf:"limits,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Requests map[string]*string `json:"requests,omitempty" tf:"requests,omitempty"`
}

type RetryStrategyInitParameters struct {

	// The number of times to move a job to the RUNNABLE status. You may specify between 1 and 10 attempts.
	Attempts *float64 `json:"attempts,omitempty" tf:"attempts,omitempty"`

	// The evaluate on exit conditions under which the job should be retried or failed. If this parameter is specified, then the attempts parameter must also be specified. You may specify up to 5 configuration blocks.
	EvaluateOnExit []EvaluateOnExitInitParameters `json:"evaluateOnExit,omitempty" tf:"evaluate_on_exit,omitempty"`
}

type RetryStrategyObservation struct {

	// The number of times to move a job to the RUNNABLE status. You may specify between 1 and 10 attempts.
	Attempts *float64 `json:"attempts,omitempty" tf:"attempts,omitempty"`

	// The evaluate on exit conditions under which the job should be retried or failed. If this parameter is specified, then the attempts parameter must also be specified. You may specify up to 5 configuration blocks.
	EvaluateOnExit []EvaluateOnExitObservation `json:"evaluateOnExit,omitempty" tf:"evaluate_on_exit,omitempty"`
}

type RetryStrategyParameters struct {

	// The number of times to move a job to the RUNNABLE status. You may specify between 1 and 10 attempts.
	// +kubebuilder:validation:Optional
	Attempts *float64 `json:"attempts,omitempty" tf:"attempts,omitempty"`

	// The evaluate on exit conditions under which the job should be retried or failed. If this parameter is specified, then the attempts parameter must also be specified. You may specify up to 5 configuration blocks.
	// +kubebuilder:validation:Optional
	EvaluateOnExit []EvaluateOnExitParameters `json:"evaluateOnExit,omitempty" tf:"evaluate_on_exit,omitempty"`
}

type SecretInitParameters struct {

	// Specifies whether the secret or the secret's keys must be defined.
	Optional *bool `json:"optional,omitempty" tf:"optional,omitempty"`

	// The name of the secret. The name must be allowed as a DNS subdomain name.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type SecretObservation struct {

	// Specifies whether the secret or the secret's keys must be defined.
	Optional *bool `json:"optional,omitempty" tf:"optional,omitempty"`

	// The name of the secret. The name must be allowed as a DNS subdomain name.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type SecretParameters struct {

	// Specifies whether the secret or the secret's keys must be defined.
	// +kubebuilder:validation:Optional
	Optional *bool `json:"optional,omitempty" tf:"optional,omitempty"`

	// The name of the secret. The name must be allowed as a DNS subdomain name.
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type SecurityContextInitParameters struct {
	Privileged *bool `json:"privileged,omitempty" tf:"privileged,omitempty"`

	ReadOnlyRootFileSystem *bool `json:"readOnlyRootFileSystem,omitempty" tf:"read_only_root_file_system,omitempty"`

	RunAsGroup *float64 `json:"runAsGroup,omitempty" tf:"run_as_group,omitempty"`

	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty" tf:"run_as_non_root,omitempty"`

	RunAsUser *float64 `json:"runAsUser,omitempty" tf:"run_as_user,omitempty"`
}

type SecurityContextObservation struct {
	Privileged *bool `json:"privileged,omitempty" tf:"privileged,omitempty"`

	ReadOnlyRootFileSystem *bool `json:"readOnlyRootFileSystem,omitempty" tf:"read_only_root_file_system,omitempty"`

	RunAsGroup *float64 `json:"runAsGroup,omitempty" tf:"run_as_group,omitempty"`

	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty" tf:"run_as_non_root,omitempty"`

	RunAsUser *float64 `json:"runAsUser,omitempty" tf:"run_as_user,omitempty"`
}

type SecurityContextParameters struct {

	// +kubebuilder:validation:Optional
	Privileged *bool `json:"privileged,omitempty" tf:"privileged,omitempty"`

	// +kubebuilder:validation:Optional
	ReadOnlyRootFileSystem *bool `json:"readOnlyRootFileSystem,omitempty" tf:"read_only_root_file_system,omitempty"`

	// +kubebuilder:validation:Optional
	RunAsGroup *float64 `json:"runAsGroup,omitempty" tf:"run_as_group,omitempty"`

	// +kubebuilder:validation:Optional
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty" tf:"run_as_non_root,omitempty"`

	// +kubebuilder:validation:Optional
	RunAsUser *float64 `json:"runAsUser,omitempty" tf:"run_as_user,omitempty"`
}

type TimeoutInitParameters struct {

	// The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is 60 seconds.
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds,omitempty" tf:"attempt_duration_seconds,omitempty"`
}

type TimeoutObservation struct {

	// The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is 60 seconds.
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds,omitempty" tf:"attempt_duration_seconds,omitempty"`
}

type TimeoutParameters struct {

	// The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is 60 seconds.
	// +kubebuilder:validation:Optional
	AttemptDurationSeconds *float64 `json:"attemptDurationSeconds,omitempty" tf:"attempt_duration_seconds,omitempty"`
}

type VolumeMountsInitParameters struct {

	// The path of the file or directory on the host to mount into containers on the pod.
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
}

type VolumeMountsObservation struct {

	// The path of the file or directory on the host to mount into containers on the pod.
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path,omitempty"`

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
}

type VolumeMountsParameters struct {

	// The path of the file or directory on the host to mount into containers on the pod.
	// +kubebuilder:validation:Optional
	MountPath *string `json:"mountPath" tf:"mount_path,omitempty"`

	// Specifies the name of the job definition.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`
}

type VolumesInitParameters struct {
	EmptyDir []EmptyDirInitParameters `json:"emptyDir,omitempty" tf:"empty_dir,omitempty"`

	// The path of the file or directory on the host to mount into containers on the pod.
	HostPath []HostPathInitParameters `json:"hostPath,omitempty" tf:"host_path,omitempty"`

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Secret []SecretInitParameters `json:"secret,omitempty" tf:"secret,omitempty"`
}

type VolumesObservation struct {
	EmptyDir []EmptyDirObservation `json:"emptyDir,omitempty" tf:"empty_dir,omitempty"`

	// The path of the file or directory on the host to mount into containers on the pod.
	HostPath []HostPathObservation `json:"hostPath,omitempty" tf:"host_path,omitempty"`

	// Specifies the name of the job definition.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Secret []SecretObservation `json:"secret,omitempty" tf:"secret,omitempty"`
}

type VolumesParameters struct {

	// +kubebuilder:validation:Optional
	EmptyDir []EmptyDirParameters `json:"emptyDir,omitempty" tf:"empty_dir,omitempty"`

	// The path of the file or directory on the host to mount into containers on the pod.
	// +kubebuilder:validation:Optional
	HostPath []HostPathParameters `json:"hostPath,omitempty" tf:"host_path,omitempty"`

	// Specifies the name of the job definition.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Secret []SecretParameters `json:"secret,omitempty" tf:"secret,omitempty"`
}

// JobDefinitionSpec defines the desired state of JobDefinition
type JobDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     JobDefinitionParameters `json:"forProvider"`
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
	InitProvider JobDefinitionInitParameters `json:"initProvider,omitempty"`
}

// JobDefinitionStatus defines the observed state of JobDefinition.
type JobDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        JobDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// JobDefinition is the Schema for the JobDefinitions API. Provides a Batch Job Definition resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type JobDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   JobDefinitionSpec   `json:"spec"`
	Status JobDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobDefinitionList contains a list of JobDefinitions
type JobDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobDefinition `json:"items"`
}

// Repository type metadata.
var (
	JobDefinition_Kind             = "JobDefinition"
	JobDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobDefinition_Kind}.String()
	JobDefinition_KindAPIVersion   = JobDefinition_Kind + "." + CRDGroupVersion.String()
	JobDefinition_GroupVersionKind = CRDGroupVersion.WithKind(JobDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&JobDefinition{}, &JobDefinitionList{})
}
