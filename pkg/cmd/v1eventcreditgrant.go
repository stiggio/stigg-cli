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

var v1EventsCreditsGrantsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new credit grant for a customer with specified amount, type, and\noptional billing configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "amount",
			Usage:    "The credit amount to grant",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "currency-id",
			Usage:    "The credit currency ID (required)",
			Required: true,
			BodyPath: "currencyId",
		},
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "The customer ID to grant credits to (required)",
			Required: true,
			BodyPath: "customerId",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "The display name for the credit grant",
			Required: true,
			BodyPath: "displayName",
		},
		&requestflag.Flag[string]{
			Name:     "grant-type",
			Usage:    "The type of credit grant (PAID, PROMOTIONAL, RECURRING)",
			Required: true,
			BodyPath: "grantType",
		},
		&requestflag.Flag[bool]{
			Name:     "await-payment-confirmation",
			Usage:    "Whether to wait for payment confirmation before returning (default: true)",
			BodyPath: "awaitPaymentConfirmation",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "billing-information",
			Usage:    "Billing information for the credit grant",
			BodyPath: "billingInformation",
		},
		&requestflag.Flag[string]{
			Name:     "comment",
			Usage:    "An optional comment on the credit grant",
			BodyPath: "comment",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "cost",
			Usage:    "The monetary cost of the credit grant",
			BodyPath: "cost",
		},
		&requestflag.Flag[any]{
			Name:     "effective-at",
			Usage:    "The date when the credit grant becomes effective",
			BodyPath: "effectiveAt",
		},
		&requestflag.Flag[any]{
			Name:     "expire-at",
			Usage:    "The date when the credit grant expires",
			BodyPath: "expireAt",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata for the credit grant",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "payment-collection-method",
			Usage:    "The payment collection method (CHARGE, INVOICE, NONE)",
			BodyPath: "paymentCollectionMethod",
		},
		&requestflag.Flag[int64]{
			Name:     "priority",
			Usage:    "The priority of the credit grant (lower number = higher priority)",
			BodyPath: "priority",
		},
		&requestflag.Flag[string]{
			Name:     "resource-id",
			Usage:    "The resource ID to scope the grant to",
			BodyPath: "resourceId",
		},
	},
	Action:          handleV1EventsCreditsGrantsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"billing-information": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "billing-information.billing-address",
			Usage:      "The billing address",
			InnerField: "billingAddress",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "billing-information.invoice-days-until-due",
			Usage:      "Days until the invoice is due",
			InnerField: "invoiceDaysUntilDue",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "billing-information.is-invoice-paid",
			Usage:      "Whether the invoice is already paid",
			InnerField: "isInvoicePaid",
		},
	},
	"cost": {
		&requestflag.InnerFlag[float64]{
			Name:       "cost.amount",
			Usage:      "The price amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "cost.currency",
			Usage:      "ISO 4217 currency code",
			InnerField: "currency",
		},
	},
})

var v1EventsCreditsGrantsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of credit grants for a customer.",
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
		&requestflag.Flag[map[string]any]{
			Name:      "created-at",
			Usage:     "Filter by creation date using range operators: gt, gte, lt, lte",
			QueryPath: "createdAt",
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
			Usage:     "Filter by resource ID. When omitted, only grants without a resource are returned",
			QueryPath: "resourceId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1EventsCreditsGrantsList,
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

var v1EventsCreditsGrantsVoid = cli.Command{
	Name:    "void",
	Usage:   "Voids an existing credit grant, preventing further consumption of the remaining\ncredits.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1EventsCreditsGrantsVoid,
	HideHelpCommand: true,
}

func handleV1EventsCreditsGrantsCreate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1EventCreditGrantNewParams{}

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
	_, err = client.V1.Events.Credits.Grants.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:events:credits:grants create", obj, format, transform)
}

func handleV1EventsCreditsGrantsList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1EventCreditGrantListParams{}

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
		_, err = client.V1.Events.Credits.Grants.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:events:credits:grants list", obj, format, transform)
	} else {
		iter := client.V1.Events.Credits.Grants.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:events:credits:grants list", iter, format, transform, maxItems)
	}
}

func handleV1EventsCreditsGrantsVoid(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Events.Credits.Grants.Void(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:events:credits:grants void", obj, format, transform)
}
