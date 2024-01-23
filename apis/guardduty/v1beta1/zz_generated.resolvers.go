/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Filter.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *Filter) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("guardduty.aws.upbound.io",

			"v1beta1",
			"Detector",

			"DetectorList",
		)
		if err !=
			nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DetectorID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DetectorIDRef,
			Selector:     mg.Spec.ForProvider.DetectorIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DetectorID")
	}
	mg.Spec.ForProvider.DetectorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DetectorIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Member.
func (mg *Member) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("guardduty.aws.upbound.io",

			"v1beta1",
			"Detector",

			"DetectorList",
		)
		if err !=
			nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccountID),
			Extract:      resource.ExtractParamPath("account_id", true),
			Reference:    mg.Spec.ForProvider.AccountIDRef,
			Selector:     mg.Spec.ForProvider.AccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccountID")
	}
	mg.Spec.ForProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccountIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("guardduty.aws.upbound.io",

			"v1beta1",
			"Detector",

			"DetectorList",
		)
		if err !=
			nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DetectorID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DetectorIDRef,
			Selector:     mg.Spec.ForProvider.DetectorIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DetectorID")
	}
	mg.Spec.ForProvider.DetectorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DetectorIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("guardduty.aws.upbound.io",

			"v1beta1",
			"Detector",

			"DetectorList",
		)
		if err !=
			nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AccountID),
			Extract:      resource.ExtractParamPath("account_id", true),
			Reference:    mg.Spec.InitProvider.AccountIDRef,
			Selector:     mg.Spec.InitProvider.AccountIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AccountID")
	}
	mg.Spec.InitProvider.AccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AccountIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("guardduty.aws.upbound.io",

			"v1beta1",
			"Detector",

			"DetectorList",
		)
		if err !=
			nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DetectorID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DetectorIDRef,
			Selector:     mg.Spec.InitProvider.DetectorIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DetectorID")
	}
	mg.Spec.InitProvider.DetectorID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DetectorIDRef = rsp.ResolvedReference

	return nil
}
