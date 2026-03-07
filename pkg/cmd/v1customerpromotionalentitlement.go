// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/stigg-cli/internal/apiquery"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
	"github.com/stiggio/stigg-go"
	"github.com/stiggio/stigg-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var v1CustomersPromotionalEntitlementsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Grants promotional entitlements to a customer, providing feature access outside\ntheir subscription. Entitlements can be time-limited or permanent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "promotional-entitlement",
			Usage:    "Promotional entitlements to grant",
			Required: true,
			BodyPath: "promotionalEntitlements",
		},
	},
	Action:          handleV1CustomersPromotionalEntitlementsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"promotional-entitlement": {
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.custom-end-date",
			Usage:      "The custom end date of the promotional entitlement",
			InnerField: "customEndDate",
		},
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.enum-values",
			Usage:      "The enum values of the entitlement",
			InnerField: "enumValues",
		},
		&requestflag.InnerFlag[string]{
			Name:       "promotional-entitlement.feature-id",
			Usage:      "The unique identifier of the entitlement feature",
			InnerField: "featureId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.has-soft-limit",
			Usage:      "Whether the entitlement has a soft limit",
			InnerField: "hasSoftLimit",
		},
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.has-unlimited-usage",
			Usage:      "Whether the entitlement has an unlimited usage",
			InnerField: "hasUnlimitedUsage",
		},
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.is-visible",
			Usage:      "Whether the entitlement is visible",
			InnerField: "isVisible",
		},
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.monthly-reset-period-configuration",
			Usage:      "The monthly reset period configuration of the entitlement, defined when reset period is monthly",
			InnerField: "monthlyResetPeriodConfiguration",
		},
		&requestflag.InnerFlag[string]{
			Name:       "promotional-entitlement.period",
			Usage:      "The grant period of the promotional entitlement",
			InnerField: "period",
		},
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.reset-period",
			Usage:      "The reset period of the entitlement",
			InnerField: "resetPeriod",
		},
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.usage-limit",
			Usage:      "The usage limit of the entitlement",
			InnerField: "usageLimit",
		},
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.weekly-reset-period-configuration",
			Usage:      "The weekly reset period configuration of the entitlement, defined when reset period is weekly",
			InnerField: "weeklyResetPeriodConfiguration",
		},
		&requestflag.InnerFlag[any]{
			Name:       "promotional-entitlement.yearly-reset-period-configuration",
			Usage:      "The yearly reset period configuration of the entitlement, defined when reset period is yearly",
			InnerField: "yearlyResetPeriodConfiguration",
		},
	},
})

var v1CustomersPromotionalEntitlementsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of a customer's promotional entitlements.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
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
		&requestflag.Flag[map[string]any]{
			Name:      "created-at",
			Usage:     "Filter by creation date using range operators: gt, gte, lt, lte",
			QueryPath: "createdAt",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by promotional entitlement status. Supports comma-separated values for multiple statuses",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1CustomersPromotionalEntitlementsList,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"created-at": {
		&requestflag.InnerFlag[any]{
			Name:       "created-at.gt",
			Usage:      "Greater than the specified createdAt value",
			InnerField: "gt",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.gte",
			Usage:      "Greater than or equal to the specified createdAt value",
			InnerField: "gte",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.lt",
			Usage:      "Less than the specified createdAt value",
			InnerField: "lt",
		},
		&requestflag.InnerFlag[any]{
			Name:       "created-at.lte",
			Usage:      "Less than or equal to the specified createdAt value",
			InnerField: "lte",
		},
	},
})

var v1CustomersPromotionalEntitlementsRevoke = cli.Command{
	Name:    "revoke",
	Usage:   "Revokes a previously granted promotional entitlement from a customer for a\nspecific feature.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "feature-id",
			Required: true,
		},
	},
	Action:          handleV1CustomersPromotionalEntitlementsRevoke,
	HideHelpCommand: true,
}

func handleV1CustomersPromotionalEntitlementsCreate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerPromotionalEntitlementNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Customers.PromotionalEntitlements.New(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:customers:promotional-entitlements create", obj, format, transform)
}

func handleV1CustomersPromotionalEntitlementsList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerPromotionalEntitlementListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Customers.PromotionalEntitlements.List(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:customers:promotional-entitlements list", obj, format, transform)
	} else {
		iter := client.V1.Customers.PromotionalEntitlements.ListAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:customers:promotional-entitlements list", iter, format, transform, maxItems)
	}
}

func handleV1CustomersPromotionalEntitlementsRevoke(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("feature-id") && len(unusedArgs) > 0 {
		cmd.Set("feature-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerPromotionalEntitlementRevokeParams{
		ID: cmd.Value("id").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Customers.PromotionalEntitlements.Revoke(
		ctx,
		cmd.Value("feature-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:customers:promotional-entitlements revoke", obj, format, transform)
}
