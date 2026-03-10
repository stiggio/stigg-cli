// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stiggio/stigg-cli/internal/apiquery"
	"github.com/stiggio/stigg-cli/internal/requestflag"
	"github.com/stiggio/stigg-go"
	"github.com/stiggio/stigg-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var v1AddonsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new addon in draft status, associated with a specific product.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The unique identifier for the entity",
			Required: true,
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "The display name of the package",
			Required: true,
			BodyPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:     "product-id",
			Usage:    "The product id of the package",
			Required: true,
			BodyPath: "productId",
		},
		&requestflag.Flag[any]{
			Name:     "billing-id",
			Usage:    "The unique identifier for the entity in the billing provider",
			BodyPath: "billingId",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "The description of the package",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "max-quantity",
			Usage:    "The maximum quantity of this addon that can be added to a subscription",
			BodyPath: "maxQuantity",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Metadata associated with the entity",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "pricing-type",
			Usage:    "The pricing type of the package",
			BodyPath: "pricingType",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The status of the package",
			BodyPath: "status",
		},
	},
	Action:          handleV1AddonsCreate,
	HideHelpCommand: true,
}

var v1AddonsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves an addon by its unique identifier, including entitlements and pricing\ndetails.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1AddonsRetrieve,
	HideHelpCommand: true,
}

var v1AddonsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates an existing addon's properties such as display name, description, and\nmetadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "billing-id",
			Usage:    "The unique identifier for the entity in the billing provider",
			BodyPath: "billingId",
		},
		&requestflag.Flag[any]{
			Name:     "dependency",
			Usage:    "List of addons the addon is dependant on",
			BodyPath: "dependencies",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "The description of the package",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "The display name of the package",
			BodyPath: "displayName",
		},
		&requestflag.Flag[any]{
			Name:     "max-quantity",
			Usage:    "The maximum quantity of this addon that can be added to a subscription",
			BodyPath: "maxQuantity",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Metadata associated with the entity",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The status of the package",
			BodyPath: "status",
		},
	},
	Action:          handleV1AddonsUpdate,
	HideHelpCommand: true,
}

var v1AddonsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of addons in the environment.",
	Suggest: true,
	Flags: []cli.Flag{
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
			Name:      "product-id",
			Usage:     "Filter by product ID",
			QueryPath: "productId",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by status. Supports comma-separated values for multiple statuses",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1AddonsList,
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

var v1AddonsArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archives an addon, preventing it from being used in new subscriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1AddonsArchive,
	HideHelpCommand: true,
}

var v1AddonsCreateDraft = cli.Command{
	Name:    "create-draft",
	Usage:   "Creates a draft version of an existing addon for modification before publishing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1AddonsCreateDraft,
	HideHelpCommand: true,
}

var v1AddonsPublish = cli.Command{
	Name:    "publish",
	Usage:   "Publishes a draft addon, making it available for use in subscriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "migration-type",
			Usage:    "The migration type of the package",
			Required: true,
			BodyPath: "migrationType",
		},
	},
	Action:          handleV1AddonsPublish,
	HideHelpCommand: true,
}

var v1AddonsRemoveDraft = cli.Command{
	Name:    "remove-draft",
	Usage:   "Removes a draft version of an addon.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1AddonsRemoveDraft,
	HideHelpCommand: true,
}

var v1AddonsSetPricing = requestflag.WithInnerFlags(cli.Command{
	Name:    "set-pricing",
	Usage:   "Sets the pricing configuration for an addon.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "pricing-type",
			Usage:    "The pricing type (FREE, PAID, or CUSTOM)",
			Required: true,
			BodyPath: "pricingType",
		},
		&requestflag.Flag[string]{
			Name:     "billing-id",
			Usage:    "Deprecated: billing integration ID",
			BodyPath: "billingId",
		},
		&requestflag.Flag[any]{
			Name:     "minimum-spend",
			Usage:    "Minimum spend configuration per billing period",
			BodyPath: "minimumSpend",
		},
		&requestflag.Flag[string]{
			Name:     "overage-billing-period",
			Usage:    "When overage charges are billed",
			BodyPath: "overageBillingPeriod",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "overage-pricing-model",
			Usage:    "Array of overage pricing model configurations",
			BodyPath: "overagePricingModels",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "pricing-model",
			Usage:    "Array of pricing model configurations",
			BodyPath: "pricingModels",
		},
	},
	Action:          handleV1AddonsSetPricing,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"minimum-spend": {
		&requestflag.InnerFlag[string]{
			Name:       "minimum-spend.billing-period",
			Usage:      "The billing period",
			InnerField: "billingPeriod",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "minimum-spend.minimum",
			Usage:      "The minimum spend amount",
			InnerField: "minimum",
		},
	},
	"overage-pricing-model": {
		&requestflag.InnerFlag[string]{
			Name:       "overage-pricing-model.billing-model",
			Usage:      "The billing model for overages",
			InnerField: "billingModel",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "overage-pricing-model.price-periods",
			Usage:      "Price periods for overage pricing",
			InnerField: "pricePeriods",
		},
		&requestflag.InnerFlag[string]{
			Name:       "overage-pricing-model.billing-cadence",
			Usage:      "The billing cadence for overages",
			InnerField: "billingCadence",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "overage-pricing-model.entitlement",
			Usage:      "Entitlement configuration for the overage feature",
			InnerField: "entitlement",
		},
		&requestflag.InnerFlag[string]{
			Name:       "overage-pricing-model.feature-id",
			Usage:      "The feature ID for overage pricing",
			InnerField: "featureId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "overage-pricing-model.top-up-custom-currency-id",
			Usage:      "Custom currency ID for overage top-up",
			InnerField: "topUpCustomCurrencyId",
		},
	},
	"pricing-model": {
		&requestflag.InnerFlag[string]{
			Name:       "pricing-model.billing-model",
			Usage:      "The billing model (FLAT_FEE, PER_UNIT, USAGE_BASED, CREDIT_BASED)",
			InnerField: "billingModel",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "pricing-model.price-periods",
			Usage:      "Array of price period configurations (at least one required)",
			InnerField: "pricePeriods",
		},
		&requestflag.InnerFlag[string]{
			Name:       "pricing-model.billing-cadence",
			Usage:      "The billing cadence (RECURRING or ONE_OFF)",
			InnerField: "billingCadence",
		},
		&requestflag.InnerFlag[string]{
			Name:       "pricing-model.feature-id",
			Usage:      "The feature ID this pricing model is associated with",
			InnerField: "featureId",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pricing-model.max-unit-quantity",
			Usage:      "Maximum number of units (max 999999)",
			InnerField: "maxUnitQuantity",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "pricing-model.min-unit-quantity",
			Usage:      "Minimum number of units",
			InnerField: "minUnitQuantity",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "pricing-model.monthly-reset-period-configuration",
			Usage:      "Monthly reset period configuration",
			InnerField: "monthlyResetPeriodConfiguration",
		},
		&requestflag.InnerFlag[string]{
			Name:       "pricing-model.reset-period",
			Usage:      "The usage reset period",
			InnerField: "resetPeriod",
		},
		&requestflag.InnerFlag[string]{
			Name:       "pricing-model.tiers-mode",
			Usage:      "The tiered pricing mode (VOLUME or GRADUATED)",
			InnerField: "tiersMode",
		},
		&requestflag.InnerFlag[string]{
			Name:       "pricing-model.top-up-custom-currency-id",
			Usage:      "The custom currency ID for top-up pricing",
			InnerField: "topUpCustomCurrencyId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "pricing-model.weekly-reset-period-configuration",
			Usage:      "Weekly reset period configuration",
			InnerField: "weeklyResetPeriodConfiguration",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "pricing-model.yearly-reset-period-configuration",
			Usage:      "Yearly reset period configuration",
			InnerField: "yearlyResetPeriodConfiguration",
		},
	},
})

func handleV1AddonsCreate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1AddonNewParams{}

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
	_, err = client.V1.Addons.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:addons create", obj, format, transform)
}

func handleV1AddonsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Addons.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:addons retrieve", obj, format, transform)
}

func handleV1AddonsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1AddonUpdateParams{}

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
	_, err = client.V1.Addons.Update(
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
	return ShowJSON(os.Stdout, "v1:addons update", obj, format, transform)
}

func handleV1AddonsList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1AddonListParams{}

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
		_, err = client.V1.Addons.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:addons list", obj, format, transform)
	} else {
		iter := client.V1.Addons.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:addons list", iter, format, transform, maxItems)
	}
}

func handleV1AddonsArchive(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Addons.Archive(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:addons archive", obj, format, transform)
}

func handleV1AddonsCreateDraft(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Addons.NewDraft(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:addons create-draft", obj, format, transform)
}

func handleV1AddonsPublish(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1AddonPublishParams{}

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
	_, err = client.V1.Addons.Publish(
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
	return ShowJSON(os.Stdout, "v1:addons publish", obj, format, transform)
}

func handleV1AddonsRemoveDraft(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Addons.RemoveDraft(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:addons remove-draft", obj, format, transform)
}

func handleV1AddonsSetPricing(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1AddonSetPricingParams{}

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
	_, err = client.V1.Addons.SetPricing(
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
	return ShowJSON(os.Stdout, "v1:addons set-pricing", obj, format, transform)
}
