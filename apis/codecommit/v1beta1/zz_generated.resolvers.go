/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this ApprovalRuleTemplateAssociation.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *ApprovalRuleTemplateAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("codecommit.aws.upbound.io",

			"v1beta1",
			"ApprovalRuleTemplate",

			"ApprovalRuleTemplateList",
		)
		if err != nil {

			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ApprovalRuleTemplateName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.ApprovalRuleTemplateNameRef,
			Selector:     mg.Spec.ForProvider.ApprovalRuleTemplateNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ApprovalRuleTemplateName")
	}
	mg.Spec.ForProvider.ApprovalRuleTemplateName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ApprovalRuleTemplateNameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("codecommit.aws.upbound.io",

			"v1beta1",
			"Repository",

			"RepositoryList",
		)
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RepositoryName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RepositoryNameRef,
			Selector:     mg.Spec.ForProvider.RepositoryNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RepositoryName")
	}
	mg.Spec.ForProvider.RepositoryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Trigger.
func (mg *Trigger) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("codecommit.aws.upbound.io",

			"v1beta1",
			"Repository",

			"RepositoryList",
		)
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RepositoryName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RepositoryNameRef,
			Selector:     mg.Spec.ForProvider.RepositoryNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RepositoryName")
	}
	mg.Spec.ForProvider.RepositoryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RepositoryNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Trigger); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io",

				"v1beta1",
				"Topic", "TopicList",
			)
			if err !=
				nil {
				return errors.
					Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Trigger[i3].DestinationArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Trigger[i3].DestinationArnRef,
				Selector:     mg.Spec.ForProvider.Trigger[i3].DestinationArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Trigger[i3].DestinationArn")
		}
		mg.Spec.ForProvider.Trigger[i3].DestinationArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Trigger[i3].DestinationArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("codecommit.aws.upbound.io",

			"v1beta1",
			"Repository",

			"RepositoryList",
		)
		if err !=
			nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.RepositoryName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.RepositoryNameRef,
			Selector:     mg.Spec.InitProvider.RepositoryNameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.RepositoryName")
	}
	mg.Spec.InitProvider.RepositoryName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.RepositoryNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Trigger); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("sns.aws.upbound.io",

				"v1beta1",
				"Topic", "TopicList",
			)
			if err !=
				nil {
				return errors.
					Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Trigger[i3].DestinationArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.Trigger[i3].DestinationArnRef,
				Selector:     mg.Spec.InitProvider.Trigger[i3].DestinationArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Trigger[i3].DestinationArn")
		}
		mg.Spec.InitProvider.Trigger[i3].DestinationArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Trigger[i3].DestinationArnRef = rsp.ResolvedReference

	}

	return nil
}
