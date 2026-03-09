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

var v1AddonsEntitlementsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates one or more entitlements (feature or credit) on a draft addon.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "addon-id",
			Required: true,
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "entitlement",
			Usage:    "Entitlements to create",
			Required: true,
			BodyPath: "entitlements",
		},
	},
	Action:          handleV1AddonsEntitlementsCreate,
	HideHelpCommand: true,
}

var v1AddonsEntitlementsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing entitlement on a draft addon.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "addon-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "UpdateFeatureEntitlementRequest",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "behavior",
			Usage:    "Entitlement behavior (Increment or Override)",
			BodyPath: "behavior",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the entitlement",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-name-override",
			Usage:    "Override display name for the entitlement",
			BodyPath: "displayNameOverride",
		},
		&requestflag.Flag[[]string]{
			Name:     "enum-value",
			Usage:    "Allowed enum values for the feature entitlement",
			BodyPath: "enumValues",
		},
		&requestflag.Flag[bool]{
			Name:     "has-soft-limit",
			Usage:    "Whether the usage limit is a soft limit",
			BodyPath: "hasSoftLimit",
		},
		&requestflag.Flag[bool]{
			Name:     "has-unlimited-usage",
			Usage:    "Whether usage is unlimited",
			BodyPath: "hasUnlimitedUsage",
		},
		&requestflag.Flag[[]string]{
			Name:     "hidden-from-widget",
			Usage:    "Widget types where this entitlement is hidden",
			BodyPath: "hiddenFromWidgets",
		},
		&requestflag.Flag[bool]{
			Name:     "is-custom",
			Usage:    "Whether this is a custom entitlement",
			BodyPath: "isCustom",
		},
		&requestflag.Flag[bool]{
			Name:     "is-granted",
			Usage:    "Whether the entitlement is granted",
			BodyPath: "isGranted",
		},
		&requestflag.Flag[any]{
			Name:     "monthly-reset-period-configuration",
			Usage:    "Configuration for monthly reset period",
			BodyPath: "monthlyResetPeriodConfiguration",
		},
		&requestflag.Flag[float64]{
			Name:     "order",
			Usage:    "Display order of the entitlement",
			BodyPath: "order",
		},
		&requestflag.Flag[string]{
			Name:     "reset-period",
			Usage:    "Period at which usage resets",
			BodyPath: "resetPeriod",
		},
		&requestflag.Flag[any]{
			Name:     "usage-limit",
			Usage:    "Maximum allowed usage for the feature",
			BodyPath: "usageLimit",
		},
		&requestflag.Flag[any]{
			Name:     "weekly-reset-period-configuration",
			Usage:    "Configuration for weekly reset period",
			BodyPath: "weeklyResetPeriodConfiguration",
		},
		&requestflag.Flag[any]{
			Name:     "yearly-reset-period-configuration",
			Usage:    "Configuration for yearly reset period",
			BodyPath: "yearlyResetPeriodConfiguration",
		},
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    "Credit grant amount",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "cadence",
			Usage:    "Credit grant cadence (MONTH or YEAR)",
			BodyPath: "cadence",
		},
	},
	Action:          handleV1AddonsEntitlementsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"monthly-reset-period-configuration": {
		&requestflag.InnerFlag[string]{
			Name:       "monthly-reset-period-configuration.according-to",
			Usage:      "Reset anchor (SubscriptionStart or StartOfTheMonth)",
			InnerField: "accordingTo",
		},
	},
	"weekly-reset-period-configuration": {
		&requestflag.InnerFlag[string]{
			Name:       "weekly-reset-period-configuration.according-to",
			Usage:      "Reset anchor (SubscriptionStart or specific day)",
			InnerField: "accordingTo",
		},
	},
	"yearly-reset-period-configuration": {
		&requestflag.InnerFlag[string]{
			Name:       "yearly-reset-period-configuration.according-to",
			Usage:      "Reset anchor (SubscriptionStart)",
			InnerField: "accordingTo",
		},
	},
})

var v1AddonsEntitlementsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves a list of entitlements for an addon.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "addon-id",
			Required: true,
		},
	},
	Action:          handleV1AddonsEntitlementsList,
	HideHelpCommand: true,
}

var v1AddonsEntitlementsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an entitlement from a draft addon.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "addon-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1AddonsEntitlementsDelete,
	HideHelpCommand: true,
}

func handleV1AddonsEntitlementsCreate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("addon-id") && len(unusedArgs) > 0 {
		cmd.Set("addon-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1AddonEntitlementNewParams{}

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
	_, err = client.V1.Addons.Entitlements.New(
		ctx,
		cmd.Value("addon-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:addons:entitlements create", obj, format, transform)
}

func handleV1AddonsEntitlementsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1AddonEntitlementUpdateParams{
		AddonID: cmd.Value("addon-id").(string),
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Addons.Entitlements.Update(
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
	return ShowJSON(os.Stdout, "v1:addons:entitlements update", obj, format, transform)
}

func handleV1AddonsEntitlementsList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("addon-id") && len(unusedArgs) > 0 {
		cmd.Set("addon-id", unusedArgs[0])
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
	_, err = client.V1.Addons.Entitlements.List(ctx, cmd.Value("addon-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:addons:entitlements list", obj, format, transform)
}

func handleV1AddonsEntitlementsDelete(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1AddonEntitlementDeleteParams{
		AddonID: cmd.Value("addon-id").(string),
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
	_, err = client.V1.Addons.Entitlements.Delete(
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
	return ShowJSON(os.Stdout, "v1:addons:entitlements delete", obj, format, transform)
}
