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

var v1CustomersRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a customer by their unique identifier, including billing information\nand subscription status.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1CustomersRetrieve,
	HideHelpCommand: true,
}

var v1CustomersUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an existing customer's properties such as name, email, and billing\ninformation.",
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
			Name:     "coupon-id",
			Usage:    "Customer level coupon",
			BodyPath: "couponId",
		},
		&requestflag.Flag[any]{
			Name:     "email",
			Usage:    "The email of the customer",
			BodyPath: "email",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "integration",
			Usage:    "List of integrations",
			BodyPath: "integrations",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "The name of the customer",
			BodyPath: "name",
		},
	},
	Action:          handleV1CustomersUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"integration": {
		&requestflag.InnerFlag[string]{
			Name:       "integration.id",
			Usage:      "Integration details",
			InnerField: "id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "integration.synced-entity-id",
			Usage:      "Synced entity id",
			InnerField: "syncedEntityId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "integration.vendor-identifier",
			Usage:      "The vendor identifier of integration",
			InnerField: "vendorIdentifier",
		},
	},
})

var v1CustomersList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of customers in the environment.",
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
		&requestflag.Flag[string]{
			Name:      "email",
			Usage:     "Filter by exact customer email address",
			QueryPath: "email",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Filter by exact customer name",
			QueryPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1CustomersList,
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

var v1CustomersArchive = cli.Command{
	Name:    "archive",
	Usage:   "Archives a customer, preventing new subscriptions. Optionally cancels existing\nsubscriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1CustomersArchive,
	HideHelpCommand: true,
}

var v1CustomersImport = requestflag.WithInnerFlags(cli.Command{
	Name:    "import",
	Usage:   "Imports multiple customers in bulk. Used for migrating customer data from\nexternal systems.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "customer",
			Usage:    "List of customer objects to import",
			Required: true,
			BodyPath: "customers",
		},
	},
	Action:          handleV1CustomersImport,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"customer": {
		&requestflag.InnerFlag[string]{
			Name:       "customer.id",
			Usage:      "Customer slug",
			InnerField: "id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "customer.email",
			Usage:      "The email of the customer",
			InnerField: "email",
		},
		&requestflag.InnerFlag[any]{
			Name:       "customer.name",
			Usage:      "The name of the customer",
			InnerField: "name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "customer.metadata",
			Usage:      "Additional metadata",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "customer.payment-method-id",
			Usage:      "Billing provider payment method id",
			InnerField: "paymentMethodId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "customer.updated-at",
			Usage:      "Timestamp of when the record was last updated",
			InnerField: "updatedAt",
		},
	},
})

var v1CustomersListResources = cli.Command{
	Name:    "list-resources",
	Usage:   "Retrieves a paginated list of resources within the same customer.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1CustomersListResources,
	HideHelpCommand: true,
}

var v1CustomersProvision = requestflag.WithInnerFlags(cli.Command{
	Name:    "provision",
	Usage:   "Creates a new customer and optionally provisions an initial subscription in a\nsingle operation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "Customer slug",
			Required: true,
			BodyPath: "id",
		},
		&requestflag.Flag[any]{
			Name:     "billing-id",
			Usage:    "The unique identifier for the entity in the billing provider",
			BodyPath: "billingId",
		},
		&requestflag.Flag[any]{
			Name:     "coupon-id",
			Usage:    "Customer level coupon",
			BodyPath: "couponId",
		},
		&requestflag.Flag[any]{
			Name:     "default-payment-method",
			Usage:    "The default payment method details",
			BodyPath: "defaultPaymentMethod",
		},
		&requestflag.Flag[any]{
			Name:     "email",
			Usage:    "The email of the customer",
			BodyPath: "email",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "integration",
			Usage:    "List of integrations",
			BodyPath: "integrations",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "name",
			Usage:    "The name of the customer",
			BodyPath: "name",
		},
	},
	Action:          handleV1CustomersProvision,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"default-payment-method": {
		&requestflag.InnerFlag[any]{
			Name:       "default-payment-method.billing-id",
			Usage:      "The default payment method id",
			InnerField: "billingId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "default-payment-method.card-expiry-month",
			Usage:      "The expiration month of the default payment method",
			InnerField: "cardExpiryMonth",
		},
		&requestflag.InnerFlag[any]{
			Name:       "default-payment-method.card-expiry-year",
			Usage:      "The expiration year of the default payment method",
			InnerField: "cardExpiryYear",
		},
		&requestflag.InnerFlag[any]{
			Name:       "default-payment-method.card-last4-digits",
			Usage:      "The last 4 digits of the default payment method",
			InnerField: "cardLast4Digits",
		},
		&requestflag.InnerFlag[string]{
			Name:       "default-payment-method.type",
			Usage:      "The default payment method type",
			InnerField: "type",
		},
	},
	"integration": {
		&requestflag.InnerFlag[string]{
			Name:       "integration.id",
			Usage:      "Integration details",
			InnerField: "id",
		},
		&requestflag.InnerFlag[any]{
			Name:       "integration.synced-entity-id",
			Usage:      "Synced entity id",
			InnerField: "syncedEntityId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "integration.vendor-identifier",
			Usage:      "The vendor identifier of integration",
			InnerField: "vendorIdentifier",
		},
	},
})

var v1CustomersUnarchive = cli.Command{
	Name:    "unarchive",
	Usage:   "Restores an archived customer, allowing them to create new subscriptions again.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1CustomersUnarchive,
	HideHelpCommand: true,
}

func handleV1CustomersRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Customers.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:customers retrieve", obj, format, transform)
}

func handleV1CustomersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerUpdateParams{}

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
	_, err = client.V1.Customers.Update(
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
	return ShowJSON(os.Stdout, "v1:customers update", obj, format, transform)
}

func handleV1CustomersList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerListParams{}

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
		_, err = client.V1.Customers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:customers list", obj, format, transform)
	} else {
		iter := client.V1.Customers.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:customers list", iter, format, transform, maxItems)
	}
}

func handleV1CustomersArchive(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Customers.Archive(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:customers archive", obj, format, transform)
}

func handleV1CustomersImport(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerImportParams{}

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
	_, err = client.V1.Customers.Import(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:customers import", obj, format, transform)
}

func handleV1CustomersListResources(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerListResourcesParams{}

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
		_, err = client.V1.Customers.ListResources(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:customers list-resources", obj, format, transform)
	} else {
		iter := client.V1.Customers.ListResourcesAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:customers list-resources", iter, format, transform, maxItems)
	}
}

func handleV1CustomersProvision(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerProvisionParams{}

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
	_, err = client.V1.Customers.Provision(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:customers provision", obj, format, transform)
}

func handleV1CustomersUnarchive(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Customers.Unarchive(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:customers unarchive", obj, format, transform)
}
