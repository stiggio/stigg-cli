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

var v1PlansCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new plan in draft status.",
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
			Usage:    "The product ID to associate the plan with",
			Required: true,
			BodyPath: "productId",
		},
		&requestflag.Flag[*string]{
			Name:     "billing-id",
			Usage:    "The unique identifier for the entity in the billing provider",
			BodyPath: "billingId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-trial-config",
			Usage:    "Default trial configuration for the plan",
			BodyPath: "defaultTrialConfig",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "The description of the package",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Metadata associated with the entity",
			BodyPath: "metadata",
		},
		&requestflag.Flag[*string]{
			Name:     "parent-plan-id",
			Usage:    "The ID of the parent plan, if applicable",
			BodyPath: "parentPlanId",
		},
		&requestflag.Flag[*string]{
			Name:     "pricing-type",
			Usage:    "The pricing type of the package",
			BodyPath: "pricingType",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			Usage:    "The status of the package",
			BodyPath: "status",
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
	Action:          handleV1PlansCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"default-trial-config": {
		&requestflag.InnerFlag[float64]{
			Name:       "default-trial-config.duration",
			Usage:      "The duration of the trial in the specified units",
			InnerField: "duration",
		},
		&requestflag.InnerFlag[string]{
			Name:       "default-trial-config.units",
			Usage:      "The time unit for the trial duration (DAY or MONTH)",
			InnerField: "units",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "default-trial-config.budget",
			Usage:      "Budget configuration for the trial",
			InnerField: "budget",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "default-trial-config.trial-end-behavior",
			Usage:      "Behavior when the trial ends (CONVERT_TO_PAID or CANCEL_SUBSCRIPTION)",
			InnerField: "trialEndBehavior",
		},
	},
})

var v1PlansRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a plan by its unique identifier, including entitlements and pricing\ndetails.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
	Action:          handleV1PlansRetrieve,
	HideHelpCommand: true,
}

var v1PlansUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing plan's properties such as display name, description, and\nmetadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[*string]{
			Name:     "billing-id",
			Usage:    "The unique identifier for the entity in the billing provider",
			BodyPath: "billingId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "charges",
			Usage:    "Pricing configuration to set on the plan draft",
			BodyPath: "charges",
		},
		&requestflag.Flag[any]{
			Name:     "compatible-addon-id",
			BodyPath: "compatibleAddonIds",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-trial-config",
			Usage:    "Default trial configuration for the plan",
			BodyPath: "defaultTrialConfig",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "The description of the package",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "The display name of the package",
			BodyPath: "displayName",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Metadata associated with the entity",
			BodyPath: "metadata",
		},
		&requestflag.Flag[*string]{
			Name:     "parent-plan-id",
			Usage:    "The ID of the parent plan, if applicable",
			BodyPath: "parentPlanId",
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
	Action:          handleV1PlansUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"charges": {
		&requestflag.InnerFlag[string]{
			Name:       "charges.pricing-type",
			Usage:      "The pricing type (FREE, PAID, or CUSTOM)",
			InnerField: "pricingType",
		},
		&requestflag.InnerFlag[string]{
			Name:       "charges.billing-id",
			Usage:      "Deprecated: billing integration ID",
			InnerField: "billingId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "charges.minimum-spend",
			Usage:      "Minimum spend configuration per billing period",
			InnerField: "minimumSpend",
		},
		&requestflag.InnerFlag[string]{
			Name:       "charges.overage-billing-period",
			Usage:      "When overage charges are billed",
			InnerField: "overageBillingPeriod",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "charges.overage-pricing-models",
			Usage:      "Array of overage pricing model configurations",
			InnerField: "overagePricingModels",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "charges.pricing-models",
			Usage:      "Array of pricing model configurations",
			InnerField: "pricingModels",
		},
	},
	"default-trial-config": {
		&requestflag.InnerFlag[float64]{
			Name:       "default-trial-config.duration",
			Usage:      "The duration of the trial in the specified units",
			InnerField: "duration",
		},
		&requestflag.InnerFlag[string]{
			Name:       "default-trial-config.units",
			Usage:      "The time unit for the trial duration (DAY or MONTH)",
			InnerField: "units",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "default-trial-config.budget",
			Usage:      "Budget configuration for the trial",
			InnerField: "budget",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "default-trial-config.trial-end-behavior",
			Usage:      "Behavior when the trial ends (CONVERT_TO_PAID or CANCEL_SUBSCRIPTION)",
			InnerField: "trialEndBehavior",
		},
	},
})

var v1PlansList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of plans in the environment.",
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
		&requestflag.Flag[[]string]{
			Name:      "status",
			Usage:     "Filter by status. Supports comma-separated values for multiple statuses",
			QueryPath: "status",
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
	Action:          handleV1PlansList,
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

var v1PlansArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archives a plan, preventing it from being used in new subscriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
	Action:          handleV1PlansArchive,
	HideHelpCommand: true,
}

var v1PlansCreateDraft = cli.Command{
	Name:    "create-draft",
	Usage:   "Creates a draft version of an existing plan for modification before publishing.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
	Action:          handleV1PlansCreateDraft,
	HideHelpCommand: true,
}

var v1PlansListCharges = cli.Command{
	Name:    "list-charges",
	Usage:   "Retrieves the list of charges configured on a plan.",
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
	Action:          handleV1PlansListCharges,
	HideHelpCommand: true,
}

var v1PlansListOverageCharges = cli.Command{
	Name:    "list-overage-charges",
	Usage:   "Retrieves the list of overage charges configured on a plan.",
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
	Action:          handleV1PlansListOverageCharges,
	HideHelpCommand: true,
}

var v1PlansPublish = cli.Command{
	Name:    "publish",
	Usage:   "Publishes a draft plan, making it available for use in subscriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:     "migration-type",
			Usage:    "The migration type of the package",
			Required: true,
			BodyPath: "migrationType",
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
	Action:          handleV1PlansPublish,
	HideHelpCommand: true,
}

var v1PlansRemoveDraft = cli.Command{
	Name:    "remove-draft",
	Usage:   "Removes a draft version of a plan.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
	Action:          handleV1PlansRemoveDraft,
	HideHelpCommand: true,
}

func handleV1PlansCreate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := stigg.V1PlanNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Plans.New(ctx, params, options...)
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
		Title:          "v1:plans create",
		Transform:      transform,
	})
}

func handleV1PlansRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1PlanGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Plans.Get(
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
		Title:          "v1:plans retrieve",
		Transform:      transform,
	})
}

func handleV1PlansUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1PlanUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Plans.Update(
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
		Title:          "v1:plans update",
		Transform:      transform,
	})
}

func handleV1PlansList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := stigg.V1PlanListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Plans.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:plans list",
			Transform:      transform,
		})
	} else {
		iter := client.V1.Plans.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:plans list",
			Transform:      transform,
		})
	}
}

func handleV1PlansArchive(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1PlanArchiveParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Plans.Archive(
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
		Title:          "v1:plans archive",
		Transform:      transform,
	})
}

func handleV1PlansCreateDraft(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1PlanNewDraftParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Plans.NewDraft(
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
		Title:          "v1:plans create-draft",
		Transform:      transform,
	})
}

func handleV1PlansListCharges(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1PlanListChargesParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Plans.ListCharges(
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
			Title:          "v1:plans list-charges",
			Transform:      transform,
		})
	} else {
		iter := client.V1.Plans.ListChargesAutoPaging(
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
			Title:          "v1:plans list-charges",
			Transform:      transform,
		})
	}
}

func handleV1PlansListOverageCharges(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1PlanListOverageChargesParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Plans.ListOverageCharges(
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
			Title:          "v1:plans list-overage-charges",
			Transform:      transform,
		})
	} else {
		iter := client.V1.Plans.ListOverageChargesAutoPaging(
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
			Title:          "v1:plans list-overage-charges",
			Transform:      transform,
		})
	}
}

func handleV1PlansPublish(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1PlanPublishParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Plans.Publish(
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
		Title:          "v1:plans publish",
		Transform:      transform,
	})
}

func handleV1PlansRemoveDraft(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1PlanRemoveDraftParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Plans.RemoveDraft(
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
		Title:          "v1:plans remove-draft",
		Transform:      transform,
	})
}
