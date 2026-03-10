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

var v1EventsCreditsGetAutoRecharge = cli.Command{
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
	Action:          handleV1EventsCreditsGetAutoRecharge,
	HideHelpCommand: true,
}

var v1EventsCreditsGetUsage = cli.Command{
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
			Name:      "currency-id",
			Usage:     "Filter by currency ID",
			QueryPath: "currencyId",
		},
		&requestflag.Flag[string]{
			Name:      "resource-id",
			Usage:     "Filter by resource ID",
			QueryPath: "resourceId",
		},
		&requestflag.Flag[string]{
			Name:      "time-range",
			Usage:     "Time range for usage data (LAST_DAY, LAST_WEEK, LAST_MONTH, LAST_YEAR). Defaults to LAST_MONTH",
			QueryPath: "timeRange",
		},
	},
	Action:          handleV1EventsCreditsGetUsage,
	HideHelpCommand: true,
}

var v1EventsCreditsListLedger = cli.Command{
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
	Action:          handleV1EventsCreditsListLedger,
	HideHelpCommand: true,
}

func handleV1EventsCreditsGetAutoRecharge(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1EventCreditGetAutoRechargeParams{}

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
	_, err = client.V1.Events.Credits.GetAutoRecharge(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:events:credits get-auto-recharge", obj, format, transform)
}

func handleV1EventsCreditsGetUsage(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1EventCreditGetUsageParams{}

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
	_, err = client.V1.Events.Credits.GetUsage(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:events:credits get-usage", obj, format, transform)
}

func handleV1EventsCreditsListLedger(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1EventCreditListLedgerParams{}

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
		_, err = client.V1.Events.Credits.ListLedger(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:events:credits list-ledger", obj, format, transform)
	} else {
		iter := client.V1.Events.Credits.ListLedgerAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:events:credits list-ledger", iter, format, transform, maxItems)
	}
}
