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

var v1CreditsConsumptionConsume = cli.Command{
	Name:    "consume",
	Usage:   "Consumes a specified amount of credits directly from a customer wallet, with no\nfeature mapping. Returns the optimistic balance.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    "The amount of credits to consume",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "currency-id",
			Usage:    "The credit currency to consume from (required)",
			Required: true,
			BodyPath: "currencyId",
		},
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "The customer to consume credits from (required)",
			Required: true,
			BodyPath: "customerId",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "A unique key used to deduplicate the consumption (required)",
			Required: true,
			BodyPath: "idempotencyKey",
		},
		&requestflag.Flag[any]{
			Name:     "created-at",
			Usage:    "Optional timestamp the consumption is attributed to",
			BodyPath: "createdAt",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dimensions",
			Usage:    "Optional dimensions describing the consumption",
			BodyPath: "dimensions",
		},
		&requestflag.Flag[string]{
			Name:     "resource-id",
			Usage:    "Optional resource the consumption is attributed to",
			BodyPath: "resourceId",
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
	Action:          handleV1CreditsConsumptionConsume,
	HideHelpCommand: true,
}

var v1CreditsConsumptionConsumeAsync = requestflag.WithInnerFlags(cli.Command{
	Name:    "consume-async",
	Usage:   "Consumes credits directly from customer wallets asynchronously. Consumptions are\nreconciled asynchronously into the credit balances.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "consumption",
			Usage:    "The credit consumptions to report (up to 1000)",
			Required: true,
			BodyPath: "consumptions",
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
	Action:          handleV1CreditsConsumptionConsumeAsync,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"consumption": {
		&requestflag.InnerFlag[float64]{
			Name:       "consumption.amount",
			Usage:      "The amount of credits to consume",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "consumption.currency-id",
			Usage:      "The credit currency to consume from (required)",
			InnerField: "currencyId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "consumption.customer-id",
			Usage:      "The customer to consume credits from (required)",
			InnerField: "customerId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "consumption.idempotency-key",
			Usage:      "A unique key used to deduplicate the consumption (required)",
			InnerField: "idempotencyKey",
		},
		&requestflag.InnerFlag[any]{
			Name:       "consumption.created-at",
			Usage:      "Optional timestamp the consumption is attributed to",
			InnerField: "createdAt",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "consumption.dimensions",
			Usage:      "Optional dimensions describing the consumption",
			InnerField: "dimensions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "consumption.resource-id",
			Usage:      "Optional resource the consumption is attributed to",
			InnerField: "resourceId",
		},
	},
})

func handleV1CreditsConsumptionConsume(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1CreditConsumptionConsumeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.Consumption.Consume(ctx, params, options...)
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
		Title:          "v1:credits:consumption consume",
		Transform:      transform,
	})
}

func handleV1CreditsConsumptionConsumeAsync(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1CreditConsumptionConsumeAsyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.Consumption.ConsumeAsync(ctx, params, options...)
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
		Title:          "v1:credits:consumption consume-async",
		Transform:      transform,
	})
}
