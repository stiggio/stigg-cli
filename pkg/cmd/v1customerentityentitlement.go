// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stiggio/stigg-cli/internal/apiquery"
	"github.com/stiggio/stigg-cli/internal/requestflag"
	"github.com/stiggio/stigg-go"
	"github.com/stiggio/stigg-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var v1CustomersEntitiesEntitlementsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a cursor-paginated list of entity entitlements for the given customer.\nAn entity entitlement ties an entity to a feature or currency with a usage limit\nand reset cadence.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "after",
			Usage:     "Return items that come after this cursor",
			QueryPath: "after",
		},
		&requestflag.Flag[string]{
			Name:      "before",
			Usage:     "Return items that come before this cursor",
			QueryPath: "before",
		},
		&requestflag.Flag[string]{
			Name:      "currency-id",
			Usage:     "Filter entity entitlements to a specific currency, by its ID. Mutually exclusive with `featureId`.",
			QueryPath: "currencyId",
		},
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "Filter entity entitlements to a specific entity ID",
			QueryPath: "entityId",
		},
		&requestflag.Flag[string]{
			Name:      "feature-id",
			Usage:     "Filter entity entitlements to a specific feature, by its ID. Mutually exclusive with `currencyId`.",
			QueryPath: "featureId",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:       "x-account-id",
			HeaderPath: "X-ACCOUNT-ID",
		},
		&requestflag.Flag[string]{
			Name:       "x-environment-id",
			HeaderPath: "X-ENVIRONMENT-ID",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1CustomersEntitiesEntitlementsList,
	HideHelpCommand: true,
}

var v1CustomersEntitiesEntitlementsUpsert = requestflag.WithInnerFlags(cli.Command{
	Name:    "upsert",
	Usage:   "Batched create-or-update of entity entitlements. Existing entity entitlements\nmatched by (entityId, featureId or currencyId, scopeEntityIds) are updated; new\nones are created. On update, omitted fields (usageLimit, cadence) are preserved;\non create both are required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "entity-entitlement",
			Usage:    "Entity entitlements to upsert (1–100 per request)",
			Required: true,
			BodyPath: "entityEntitlements",
		},
		&requestflag.Flag[string]{
			Name:       "x-account-id",
			HeaderPath: "X-ACCOUNT-ID",
		},
		&requestflag.Flag[string]{
			Name:       "x-environment-id",
			HeaderPath: "X-ENVIRONMENT-ID",
		},
	},
	Action:          handleV1CustomersEntitiesEntitlementsUpsert,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"entity-entitlement": {
		&requestflag.InnerFlag[string]{
			Name:       "entity-entitlement.entity-id",
			Usage:      "The entity ID this entity entitlement is attached to",
			InnerField: "entityId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "entity-entitlement.cadence",
			Usage:      "Usage-reset cadence (required on create) as an ISO-8601 single-unit duration, e.g. `P1M`, `P30D`, `PT1M`.",
			InnerField: "cadence",
		},
		&requestflag.InnerFlag[string]{
			Name:       "entity-entitlement.currency-id",
			Usage:      "Currency ID this entity entitlement grants (credit limits). Mutually exclusive with `featureId`.",
			InnerField: "currencyId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "entity-entitlement.feature-id",
			Usage:      "Feature ID this entity entitlement grants. Mutually exclusive with `currencyId`.",
			InnerField: "featureId",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "entity-entitlement.parent-id",
			Usage:      "Parent entity ID in the hierarchy. Omit to leave the current parent untouched (a new node defaults to a root); `null` detaches to a root; an ID sets or changes the parent. Reparenting or detaching an existing node is leaf-only; place a parent before its children, or in the same batch.",
			InnerField: "parentId",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "entity-entitlement.scope-entity-ids",
			Usage:      "Dimension-scoped sub-limit: the set of entity IDs this limit applies to (order-insensitive, unique, at most 5). Omit or leave empty for the node-wide limit.",
			InnerField: "scopeEntityIds",
		},
		&requestflag.InnerFlag[*float64]{
			Name:       "entity-entitlement.usage-limit",
			Usage:      "Maximum usage allowed within one cadence window (required on create)",
			InnerField: "usageLimit",
		},
	},
})

func handleV1CustomersEntitiesEntitlementsList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := stigg.V1CustomerEntityEntitlementListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Customers.Entities.Entitlements.List(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:customers:entities:entitlements list",
			Transform:      transform,
		})
	} else {
		iter := client.V1.Customers.Entities.Entitlements.ListAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:customers:entities:entitlements list",
			Transform:      transform,
		})
	}
}

func handleV1CustomersEntitiesEntitlementsUpsert(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := stigg.V1CustomerEntityEntitlementUpsertParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Customers.Entities.Entitlements.Upsert(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1:customers:entities:entitlements upsert",
		Transform:      transform,
	})
}
