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

	// ResolveReferences of this PrivateDNSNamespace.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *PrivateDNSNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1",
			"VPC", "VPCList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPC),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.VPCRef,
			Selector:     mg.Spec.ForProvider.VPCSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPC")
	}
	mg.Spec.ForProvider.VPC = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io",

			"v1beta1",
			"VPC", "VPCList",
		)
		if err != nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")

		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPC),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.VPCRef,
			Selector:     mg.Spec.InitProvider.VPCSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPC")
	}
	mg.Spec.InitProvider.VPC = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DNSConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("servicediscovery.aws.upbound.io",

				"v1beta1", "PrivateDNSNamespace",

				"PrivateDNSNamespaceList",
			)
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DNSConfig[i3].NamespaceID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.DNSConfig[i3].NamespaceIDRef,
				Selector:     mg.Spec.ForProvider.DNSConfig[i3].NamespaceIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DNSConfig[i3].NamespaceID")
		}
		mg.Spec.ForProvider.DNSConfig[i3].NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DNSConfig[i3].NamespaceIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.DNSConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("servicediscovery.aws.upbound.io",

				"v1beta1", "PrivateDNSNamespace",

				"PrivateDNSNamespaceList",
			)
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DNSConfig[i3].NamespaceID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.DNSConfig[i3].NamespaceIDRef,
				Selector:     mg.Spec.InitProvider.DNSConfig[i3].NamespaceIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DNSConfig[i3].NamespaceID")
		}
		mg.Spec.InitProvider.DNSConfig[i3].NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DNSConfig[i3].NamespaceIDRef = rsp.ResolvedReference

	}

	return nil
}
