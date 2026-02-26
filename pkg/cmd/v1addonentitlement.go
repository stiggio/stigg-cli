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

var v1AddonsEntitlementsCreate = requestflag.WithInnerFlags(cli.Command{
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
}, map[string][]requestflag.HasOuterFlag{
	"entitlement": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "entitlement.credit",
			Usage:      "Credit entitlement to create",
			InnerField: "credit",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "entitlement.feature",
			Usage:      "Feature entitlement to create",
			InnerField: "feature",
		},
	},
})

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
		&requestflag.Flag[map[string]any]{
			Name:     "credit",
			Usage:    "Credit entitlement fields to update",
			BodyPath: "credit",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "feature",
			Usage:    "Feature entitlement fields to update",
			BodyPath: "feature",
		},
	},
	Action:          handleV1AddonsEntitlementsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"credit": {
		&requestflag.InnerFlag[float64]{
			Name:       "credit.amount",
			Usage:      "Credit grant amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "credit.behavior",
			Usage:      "Entitlement behavior (Increment or Override)",
			InnerField: "behavior",
		},
		&requestflag.InnerFlag[string]{
			Name:       "credit.cadence",
			Usage:      "Credit grant cadence (MONTH or YEAR)",
			InnerField: "cadence",
		},
		&requestflag.InnerFlag[string]{
			Name:       "credit.description",
			Usage:      "Description of the entitlement",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "credit.display-name-override",
			Usage:      "Override display name for the entitlement",
			InnerField: "displayNameOverride",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "credit.hidden-from-widgets",
			Usage:      "Widget types where this entitlement is hidden",
			InnerField: "hiddenFromWidgets",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "credit.is-custom",
			Usage:      "Whether this is a custom entitlement",
			InnerField: "isCustom",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "credit.is-granted",
			Usage:      "Whether the entitlement is granted",
			InnerField: "isGranted",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "credit.order",
			Usage:      "Display order of the entitlement",
			InnerField: "order",
		},
	},
	"feature": {
		&requestflag.InnerFlag[string]{
			Name:       "feature.behavior",
			Usage:      "Entitlement behavior (Increment or Override)",
			InnerField: "behavior",
		},
		&requestflag.InnerFlag[string]{
			Name:       "feature.description",
			Usage:      "Description of the entitlement",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "feature.display-name-override",
			Usage:      "Override display name for the entitlement",
			InnerField: "displayNameOverride",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "feature.enum-values",
			Usage:      "Allowed enum values for the feature entitlement",
			InnerField: "enumValues",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature.has-soft-limit",
			Usage:      "Whether the usage limit is a soft limit",
			InnerField: "hasSoftLimit",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature.has-unlimited-usage",
			Usage:      "Whether usage is unlimited",
			InnerField: "hasUnlimitedUsage",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "feature.hidden-from-widgets",
			Usage:      "Widget types where this entitlement is hidden",
			InnerField: "hiddenFromWidgets",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature.is-custom",
			Usage:      "Whether this is a custom entitlement",
			InnerField: "isCustom",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature.is-granted",
			Usage:      "Whether the entitlement is granted",
			InnerField: "isGranted",
		},
		&requestflag.InnerFlag[any]{
			Name:       "feature.monthly-reset-period-configuration",
			Usage:      "Configuration for monthly reset period",
			InnerField: "monthlyResetPeriodConfiguration",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "feature.order",
			Usage:      "Display order of the entitlement",
			InnerField: "order",
		},
		&requestflag.InnerFlag[string]{
			Name:       "feature.reset-period",
			Usage:      "Period at which usage resets",
			InnerField: "resetPeriod",
		},
		&requestflag.InnerFlag[any]{
			Name:       "feature.usage-limit",
			Usage:      "Maximum allowed usage for the feature",
			InnerField: "usageLimit",
		},
		&requestflag.InnerFlag[any]{
			Name:       "feature.weekly-reset-period-configuration",
			Usage:      "Configuration for weekly reset period",
			InnerField: "weeklyResetPeriodConfiguration",
		},
		&requestflag.InnerFlag[any]{
			Name:       "feature.yearly-reset-period-configuration",
			Usage:      "Configuration for yearly reset period",
			InnerField: "yearlyResetPeriodConfiguration",
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
