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

var v1CreditsCustomCurrenciesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new custom currency in the environment.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The unique identifier for the new custom currency",
			Required: true,
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "The display name of the custom currency",
			Required: true,
			BodyPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the currency",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata to attach to the custom currency",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "symbol",
			Usage:    "The symbol used to represent the custom currency",
			BodyPath: "symbol",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "units",
			Usage:    "Singular and plural unit labels for a custom currency. Both fields are required when supplied.",
			BodyPath: "units",
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
	Action:          handleV1CreditsCustomCurrenciesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"units": {
		&requestflag.InnerFlag[string]{
			Name:       "units.plural",
			Usage:      "Plural form of the unit label",
			InnerField: "plural",
		},
		&requestflag.InnerFlag[string]{
			Name:       "units.singular",
			Usage:      "Singular form of the unit label",
			InnerField: "singular",
		},
	},
})

var v1CreditsCustomCurrenciesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing custom currency. Only the supplied fields are modified.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "currency-id",
			Required:  true,
			PathParam: "currencyId",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "A human-readable description of the custom currency. Send an empty string to clear.",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "The display name of the custom currency",
			BodyPath: "displayName",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata to attach to the custom currency",
			BodyPath: "metadata",
		},
		&requestflag.Flag[*string]{
			Name:     "symbol",
			Usage:    "The symbol used to represent the custom currency. Send an empty string to clear.",
			BodyPath: "symbol",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "units",
			Usage:    "Singular and plural unit labels for a custom currency. Both fields are required when supplied.",
			BodyPath: "units",
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
	Action:          handleV1CreditsCustomCurrenciesUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"units": {
		&requestflag.InnerFlag[string]{
			Name:       "units.plural",
			Usage:      "Plural form of the unit label",
			InnerField: "plural",
		},
		&requestflag.InnerFlag[string]{
			Name:       "units.singular",
			Usage:      "Singular form of the unit label",
			InnerField: "singular",
		},
	},
})

var v1CreditsCustomCurrenciesList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of custom currencies in the environment. Archived\ncurrencies are excluded by default; pass `status=ARCHIVED` (or\n`status=ACTIVE,ARCHIVED`) to include them.",
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "status",
			Usage:     "Filter by custom currency status. Supports comma-separated values (e.g., `ACTIVE,ARCHIVED`). Defaults to `ACTIVE`.",
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
	Action:          handleV1CreditsCustomCurrenciesList,
	HideHelpCommand: true,
}

var v1CreditsCustomCurrenciesArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archives a custom currency. Fails if the currency is still associated with any\nactive plan or addon — use the associated-entities endpoint first to inspect\ndependencies.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "currency-id",
			Required:  true,
			PathParam: "currencyId",
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
	Action:          handleV1CreditsCustomCurrenciesArchive,
	HideHelpCommand: true,
}

var v1CreditsCustomCurrenciesListAssociatedEntities = cli.Command{
	Name:    "list-associated-entities",
	Usage:   "Lists the active plans and addons that reference a custom currency. Useful\nbefore archiving to inspect dependencies.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "currency-id",
			Required:  true,
			PathParam: "currencyId",
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
	Action:          handleV1CreditsCustomCurrenciesListAssociatedEntities,
	HideHelpCommand: true,
}

var v1CreditsCustomCurrenciesUnarchive = cli.Command{
	Name:    "unarchive",
	Usage:   "Restores a previously archived custom currency. Fails if another active currency\nwith the same ID already exists.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "currency-id",
			Required:  true,
			PathParam: "currencyId",
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
	Action:          handleV1CreditsCustomCurrenciesUnarchive,
	HideHelpCommand: true,
}

func handleV1CreditsCustomCurrenciesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1CreditCustomCurrencyNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.CustomCurrencies.New(ctx, params, options...)
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
		Title:          "v1:credits:custom-currencies create",
		Transform:      transform,
	})
}

func handleV1CreditsCustomCurrenciesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("currency-id") && len(unusedArgs) > 0 {
		cmd.Set("currency-id", unusedArgs[0])
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

	params := stigg.V1CreditCustomCurrencyUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.CustomCurrencies.Update(
		ctx,
		cmd.Value("currency-id").(string),
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
		Title:          "v1:credits:custom-currencies update",
		Transform:      transform,
	})
}

func handleV1CreditsCustomCurrenciesList(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1CreditCustomCurrencyListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Credits.CustomCurrencies.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:credits:custom-currencies list",
			Transform:      transform,
		})
	} else {
		iter := client.V1.Credits.CustomCurrencies.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:credits:custom-currencies list",
			Transform:      transform,
		})
	}
}

func handleV1CreditsCustomCurrenciesArchive(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("currency-id") && len(unusedArgs) > 0 {
		cmd.Set("currency-id", unusedArgs[0])
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

	params := stigg.V1CreditCustomCurrencyArchiveParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.CustomCurrencies.Archive(
		ctx,
		cmd.Value("currency-id").(string),
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
		Title:          "v1:credits:custom-currencies archive",
		Transform:      transform,
	})
}

func handleV1CreditsCustomCurrenciesListAssociatedEntities(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("currency-id") && len(unusedArgs) > 0 {
		cmd.Set("currency-id", unusedArgs[0])
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

	params := stigg.V1CreditCustomCurrencyListAssociatedEntitiesParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.CustomCurrencies.ListAssociatedEntities(
		ctx,
		cmd.Value("currency-id").(string),
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
		Title:          "v1:credits:custom-currencies list-associated-entities",
		Transform:      transform,
	})
}

func handleV1CreditsCustomCurrenciesUnarchive(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("currency-id") && len(unusedArgs) > 0 {
		cmd.Set("currency-id", unusedArgs[0])
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

	params := stigg.V1CreditCustomCurrencyUnarchiveParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.CustomCurrencies.Unarchive(
		ctx,
		cmd.Value("currency-id").(string),
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
		Title:          "v1:credits:custom-currencies unarchive",
		Transform:      transform,
	})
}
