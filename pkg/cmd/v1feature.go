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

var v1FeaturesArchiveFeature = cli.Command{
	Name:    "archive-feature",
	Usage:   "Archives a feature, preventing it from being used in new entitlements.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1FeaturesArchiveFeature,
	HideHelpCommand: true,
}

var v1FeaturesCreateFeature = requestflag.WithInnerFlags(cli.Command{
	Name:    "create-feature",
	Usage:   "Creates a new feature with the specified type, metering, and configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The unique identifier for the feature",
			Required: true,
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "The display name for the feature",
			Required: true,
			BodyPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:     "feature-type",
			Usage:    "The type of the feature",
			Required: true,
			BodyPath: "featureType",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description for the feature",
			BodyPath: "description",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "enum-configuration",
			Usage:    "The configuration data for the feature",
			BodyPath: "enumConfiguration",
		},
		&requestflag.Flag[string]{
			Name:     "feature-status",
			Usage:    "The status of the feature",
			Default:  "ACTIVE",
			BodyPath: "featureStatus",
		},
		&requestflag.Flag[string]{
			Name:     "feature-units",
			Usage:    "The units for the feature",
			BodyPath: "featureUnits",
		},
		&requestflag.Flag[string]{
			Name:     "feature-units-plural",
			Usage:    "The plural units for the feature",
			BodyPath: "featureUnitsPlural",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "The additional metadata for the feature",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "meter-type",
			Usage:    "The meter type for the feature",
			BodyPath: "meterType",
		},
		&requestflag.Flag[any]{
			Name:     "unit-transformation",
			Usage:    "Unit transformation to be applied to the reported usage",
			BodyPath: "unitTransformation",
		},
	},
	Action:          handleV1FeaturesCreateFeature,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"enum-configuration": {
		&requestflag.InnerFlag[string]{
			Name:       "enum-configuration.display-name",
			Usage:      "The display name for the enum configuration entity",
			InnerField: "displayName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "enum-configuration.value",
			Usage:      "The unique value identifier for the enum configuration entity",
			InnerField: "value",
		},
	},
	"unit-transformation": {
		&requestflag.InnerFlag[int64]{
			Name:       "unit-transformation.divide",
			Usage:      "Divide usage by this number",
			InnerField: "divide",
		},
		&requestflag.InnerFlag[string]{
			Name:       "unit-transformation.feature-units",
			Usage:      "Singular feature units after the transformation",
			InnerField: "featureUnits",
		},
		&requestflag.InnerFlag[string]{
			Name:       "unit-transformation.feature-units-plural",
			Usage:      "Plural feature units after the transformation",
			InnerField: "featureUnitsPlural",
		},
		&requestflag.InnerFlag[string]{
			Name:       "unit-transformation.round",
			Usage:      "After division, either round the result up or down",
			InnerField: "round",
		},
	},
})

var v1FeaturesListFeatures = requestflag.WithInnerFlags(cli.Command{
	Name:    "list-features",
	Usage:   "Retrieves a paginated list of features in the environment.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "Filter by entity ID",
			QueryPath: "id",
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
		&requestflag.Flag[string]{
			Name:      "feature-type",
			Usage:     "Filter by feature type. Supports comma-separated values for multiple types",
			QueryPath: "featureType",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "meter-type",
			Usage:     "Filter by meter type. Supports comma-separated values for multiple types",
			QueryPath: "meterType",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by feature status. Supports comma-separated values for multiple statuses",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1FeaturesListFeatures,
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

var v1FeaturesRetrieveFeature = cli.Command{
	Name:    "retrieve-feature",
	Usage:   "Retrieves a feature by its unique identifier.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1FeaturesRetrieveFeature,
	HideHelpCommand: true,
}

var v1FeaturesUnarchiveFeature = cli.Command{
	Name:    "unarchive-feature",
	Usage:   "Restores an archived feature, allowing it to be used in entitlements again.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1FeaturesUnarchiveFeature,
	HideHelpCommand: true,
}

var v1FeaturesUpdateFeature = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-feature",
	Usage:   "Updates an existing feature's properties such as display name, description, and\nconfiguration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "The description for the feature",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "The display name for the feature",
			BodyPath: "displayName",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "enum-configuration",
			Usage:    "The configuration data for the feature",
			BodyPath: "enumConfiguration",
		},
		&requestflag.Flag[string]{
			Name:     "feature-units",
			Usage:    "The units for the feature",
			BodyPath: "featureUnits",
		},
		&requestflag.Flag[string]{
			Name:     "feature-units-plural",
			Usage:    "The plural units for the feature",
			BodyPath: "featureUnitsPlural",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "The additional metadata for the feature",
			BodyPath: "metadata",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "meter",
			BodyPath: "meter",
		},
		&requestflag.Flag[any]{
			Name:     "unit-transformation",
			Usage:    "Unit transformation to be applied to the reported usage",
			BodyPath: "unitTransformation",
		},
	},
	Action:          handleV1FeaturesUpdateFeature,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"enum-configuration": {
		&requestflag.InnerFlag[string]{
			Name:       "enum-configuration.display-name",
			Usage:      "The display name for the enum configuration entity",
			InnerField: "displayName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "enum-configuration.value",
			Usage:      "The unique value identifier for the enum configuration entity",
			InnerField: "value",
		},
	},
	"meter": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "meter.aggregation",
			InnerField: "aggregation",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "meter.filters",
			InnerField: "filters",
		},
	},
	"unit-transformation": {
		&requestflag.InnerFlag[int64]{
			Name:       "unit-transformation.divide",
			Usage:      "Divide usage by this number",
			InnerField: "divide",
		},
		&requestflag.InnerFlag[string]{
			Name:       "unit-transformation.feature-units",
			Usage:      "Singular feature units after the transformation",
			InnerField: "featureUnits",
		},
		&requestflag.InnerFlag[string]{
			Name:       "unit-transformation.feature-units-plural",
			Usage:      "Plural feature units after the transformation",
			InnerField: "featureUnitsPlural",
		},
		&requestflag.InnerFlag[string]{
			Name:       "unit-transformation.round",
			Usage:      "After division, either round the result up or down",
			InnerField: "round",
		},
	},
})

func handleV1FeaturesArchiveFeature(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Features.ArchiveFeature(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:features archive-feature", obj, format, transform)
}

func handleV1FeaturesCreateFeature(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1FeatureNewFeatureParams{}

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
	_, err = client.V1.Features.NewFeature(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:features create-feature", obj, format, transform)
}

func handleV1FeaturesListFeatures(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1FeatureListFeaturesParams{}

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
		_, err = client.V1.Features.ListFeatures(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:features list-features", obj, format, transform)
	} else {
		iter := client.V1.Features.ListFeaturesAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:features list-features", iter, format, transform, maxItems)
	}
}

func handleV1FeaturesRetrieveFeature(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Features.GetFeature(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:features retrieve-feature", obj, format, transform)
}

func handleV1FeaturesUnarchiveFeature(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Features.UnarchiveFeature(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:features unarchive-feature", obj, format, transform)
}

func handleV1FeaturesUpdateFeature(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1FeatureUpdateFeatureParams{}

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
	_, err = client.V1.Features.UpdateFeature(
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
	return ShowJSON(os.Stdout, "v1:features update-feature", obj, format, transform)
}
