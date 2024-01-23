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

	common "github.com/upbound/provider-aws/config/common"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this APICache.
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

func (mg *APICache) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"GraphQLAPI",

			"GraphQLAPIList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.APIIDRef,
			Selector:     mg.Spec.ForProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"GraphQLAPI",

			"GraphQLAPIList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.APIIDRef,
			Selector:     mg.Spec.InitProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this APIKey.
func (mg *APIKey) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"GraphQLAPI",

			"GraphQLAPIList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.APIIDRef,
			Selector:     mg.Spec.ForProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Datasource.
func (mg *Datasource) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"GraphQLAPI",

			"GraphQLAPIList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.APIIDRef,
			Selector:     mg.Spec.ForProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DynamodbConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("dynamodb.aws.upbound.io",

				"v1beta1",
				"Table", "TableList",
			)
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DynamodbConfig[i3].TableName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.DynamodbConfig[i3].TableNameRef,
				Selector:     mg.Spec.ForProvider.DynamodbConfig[i3].TableNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DynamodbConfig[i3].TableName")
		}
		mg.Spec.ForProvider.DynamodbConfig[i3].TableName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DynamodbConfig[i3].TableNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1",
			"Role", "RoleList",
		)
		if err !=

			nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.ServiceRoleArnRef,
			Selector:     mg.Spec.ForProvider.ServiceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceRoleArn")
	}
	mg.Spec.ForProvider.ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceRoleArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.DynamodbConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("dynamodb.aws.upbound.io",

				"v1beta1",
				"Table", "TableList",
			)
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DynamodbConfig[i3].TableName),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.DynamodbConfig[i3].TableNameRef,
				Selector:     mg.Spec.InitProvider.DynamodbConfig[i3].TableNameSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DynamodbConfig[i3].TableName")
		}
		mg.Spec.InitProvider.DynamodbConfig[i3].TableName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DynamodbConfig[i3].TableNameRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

			"v1beta1",
			"Role", "RoleList",
		)
		if err !=

			nil {
			return errors.
				Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceRoleArn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.InitProvider.ServiceRoleArnRef,
			Selector:     mg.Spec.InitProvider.ServiceRoleArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceRoleArn")
	}
	mg.Spec.InitProvider.ServiceRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceRoleArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Function.
func (mg *Function) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"GraphQLAPI",

			"GraphQLAPIList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.APIIDRef,
			Selector:     mg.Spec.ForProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"Datasource",

			"DatasourceList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataSource),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DataSourceRef,
			Selector:     mg.Spec.ForProvider.DataSourceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DataSource")
	}
	mg.Spec.ForProvider.DataSource = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DataSourceRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"GraphQLAPI",

			"GraphQLAPIList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.APIIDRef,
			Selector:     mg.Spec.InitProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.APIID")
	}
	mg.Spec.InitProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"Datasource",

			"DatasourceList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataSource),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DataSourceRef,
			Selector:     mg.Spec.InitProvider.DataSourceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DataSource")
	}
	mg.Spec.InitProvider.DataSource = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DataSourceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this GraphQLAPI.
func (mg *GraphQLAPI) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.LogConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

				"v1beta1",
				"Role", "RoleList",
			)
			if err !=

				nil {
				return errors.
					Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LogConfig[i3].CloudwatchLogsRoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.LogConfig[i3].CloudwatchLogsRoleArnRef,
				Selector:     mg.Spec.ForProvider.LogConfig[i3].CloudwatchLogsRoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.LogConfig[i3].CloudwatchLogsRoleArn")
		}
		mg.Spec.ForProvider.LogConfig[i3].CloudwatchLogsRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.LogConfig[i3].CloudwatchLogsRoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.UserPoolConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io",

				"v1beta1",
				"UserPool",

				"UserPoolList",
			)
			if err !=
				nil {

				return errors.Wrap(err,

					"failed to get the reference target managed resource and its list for reference resolution",
				)
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.UserPoolConfig[i3].UserPoolID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.UserPoolConfig[i3].UserPoolIDRef,
				Selector:     mg.Spec.ForProvider.UserPoolConfig[i3].UserPoolIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.UserPoolConfig[i3].UserPoolID")
		}
		mg.Spec.ForProvider.UserPoolConfig[i3].UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.UserPoolConfig[i3].UserPoolIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.LogConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("iam.aws.upbound.io",

				"v1beta1",
				"Role", "RoleList",
			)
			if err !=

				nil {
				return errors.
					Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LogConfig[i3].CloudwatchLogsRoleArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.LogConfig[i3].CloudwatchLogsRoleArnRef,
				Selector:     mg.Spec.InitProvider.LogConfig[i3].CloudwatchLogsRoleArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.LogConfig[i3].CloudwatchLogsRoleArn")
		}
		mg.Spec.InitProvider.LogConfig[i3].CloudwatchLogsRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.LogConfig[i3].CloudwatchLogsRoleArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.UserPoolConfig); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io",

				"v1beta1",
				"UserPool",

				"UserPoolList",
			)
			if err !=
				nil {

				return errors.Wrap(err,

					"failed to get the reference target managed resource and its list for reference resolution",
				)
			}

			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.UserPoolConfig[i3].UserPoolID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.UserPoolConfig[i3].UserPoolIDRef,
				Selector:     mg.Spec.InitProvider.UserPoolConfig[i3].UserPoolIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.UserPoolConfig[i3].UserPoolID")
		}
		mg.Spec.InitProvider.UserPoolConfig[i3].UserPoolID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.UserPoolConfig[i3].UserPoolIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this Resolver.
func (mg *Resolver) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList

	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"GraphQLAPI",

			"GraphQLAPIList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.APIIDRef,
			Selector:     mg.Spec.ForProvider.APIIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.APIID")
	}
	mg.Spec.ForProvider.APIID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.APIIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"Datasource",

			"DatasourceList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DataSource),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DataSourceRef,
			Selector:     mg.Spec.ForProvider.DataSourceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DataSource")
	}
	mg.Spec.ForProvider.DataSource = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DataSourceRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appsync.aws.upbound.io",

			"v1beta1",
			"Datasource",

			"DatasourceList",
		)
		if err != nil {

			return errors.Wrap(err,

				"failed to get the reference target managed resource and its list for reference resolution",
			)
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DataSource),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.DataSourceRef,
			Selector:     mg.Spec.InitProvider.DataSourceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DataSource")
	}
	mg.Spec.InitProvider.DataSource = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DataSourceRef = rsp.ResolvedReference

	return nil
}
