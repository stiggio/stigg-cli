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

var v1CreditsGetAutoRecharge = cli.Command{
	Name:    "get-auto-recharge",
	Usage:   "Retrieves the automatic recharge configuration for a customer and currency.\nReturns default settings if no configuration exists.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "currency-id",
			Usage:     "Filter by currency ID (required)",
			Required:  true,
			QueryPath: "currencyId",
		},
		&requestflag.Flag[string]{
			Name:      "customer-id",
			Usage:     "Filter by customer ID (required)",
			Required:  true,
			QueryPath: "customerId",
		},
	},
	Action:          handleV1CreditsGetAutoRecharge,
	HideHelpCommand: true,
}

var v1CreditsGetUsage = cli.Command{
	Name:    "get-usage",
	Usage:   "Retrieves credit usage time-series data for a customer, grouped by feature, over\na specified time range.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "customer-id",
			Usage:     "Filter by customer ID (required)",
			Required:  true,
			QueryPath: "customerId",
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
			Usage:     "Filter by currency ID",
			QueryPath: "currencyId",
		},
		&requestflag.Flag[any]{
			Name:      "end-date",
			Usage:     "End date for the credit usage time range (ISO 8601). Defaults to now when startDate is provided",
			QueryPath: "endDate",
		},
		&requestflag.Flag[string]{
			Name:      "group-by",
			Usage:     "Comma-separated list of feature dimension keys to group usage series by (up to 3). Each key matches /^[a-zA-Z0-9_$-]+$/",
			QueryPath: "groupBy",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "resource-id",
			Usage:     "Filter by resource ID",
			QueryPath: "resourceId",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			Usage:     "Start date for the credit usage time range (ISO 8601). Takes precedence over timeRange when provided",
			QueryPath: "startDate",
		},
		&requestflag.Flag[string]{
			Name:      "time-range",
			Usage:     "Time range for usage data (LAST_DAY, LAST_WEEK, LAST_MONTH, LAST_YEAR). Defaults to LAST_MONTH",
			QueryPath: "timeRange",
		},
	},
	Action:          handleV1CreditsGetUsage,
	HideHelpCommand: true,
}

var v1CreditsListLedger = cli.Command{
	Name:    "list-ledger",
	Usage:   "Retrieves a paginated list of credit ledger events for a customer.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "customer-id",
			Usage:     "Filter by customer ID (required)",
			Required:  true,
			QueryPath: "customerId",
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
			Usage:     "Filter by currency ID",
			QueryPath: "currencyId",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "resource-id",
			Usage:     "Filter by resource ID",
			QueryPath: "resourceId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1CreditsListLedger,
	HideHelpCommand: true,
}

func handleV1CreditsGetAutoRecharge(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1CreditGetAutoRechargeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.GetAutoRecharge(ctx, params, options...)
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
		Title:          "v1:credits get-auto-recharge",
		Transform:      transform,
	})
}

func handleV1CreditsGetUsage(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1CreditGetUsageParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.GetUsage(ctx, params, options...)
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
		Title:          "v1:credits get-usage",
		Transform:      transform,
	})
}

func handleV1CreditsListLedger(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1CreditListLedgerParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Credits.ListLedger(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:credits list-ledger",
			Transform:      transform,
		})
	} else {
		iter := client.V1.Credits.ListLedgerAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:credits list-ledger",
			Transform:      transform,
		})
	}
}
