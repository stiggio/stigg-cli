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

var v1ProductsArchiveProduct = cli.Command{
	Name:    "archive-product",
	Usage:   "Archives a product, preventing new subscriptions. All plans and addons are\narchived.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1ProductsArchiveProduct,
	HideHelpCommand: true,
}

var v1ProductsCreateProduct = cli.Command{
	Name:    "create-product",
	Usage:   "Creates a new product.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The unique identifier for the entity",
			Required: true,
			BodyPath: "id",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Description of the product",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name of the product",
			BodyPath: "displayName",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Additional metadata for the product",
			BodyPath: "metadata",
		},
		&requestflag.Flag[bool]{
			Name:     "multiple-subscriptions",
			Usage:    "Indicates if multiple subscriptions to this product are allowed",
			Default:  false,
			BodyPath: "multipleSubscriptions",
		},
	},
	Action:          handleV1ProductsCreateProduct,
	HideHelpCommand: true,
}

var v1ProductsDuplicateProduct = cli.Command{
	Name:    "duplicate-product",
	Usage:   "Duplicates an existing product, including its plans, addons, and configuration.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The unique identifier for the entity",
			Required: true,
			BodyPath: "id",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Description of the product",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name of the product",
			BodyPath: "displayName",
		},
	},
	Action:          handleV1ProductsDuplicateProduct,
	HideHelpCommand: true,
}

var v1ProductsListProducts = requestflag.WithInnerFlags(cli.Command{
	Name:    "list-products",
	Usage:   "Retrieves a paginated list of products in the environment.",
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by product status. Supports comma-separated values for multiple statuses",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1ProductsListProducts,
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

var v1ProductsUnarchiveProduct = cli.Command{
	Name:    "unarchive-product",
	Usage:   "Restores an archived product, allowing new subscriptions to be created.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1ProductsUnarchiveProduct,
	HideHelpCommand: true,
}

var v1ProductsUpdateProduct = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-product",
	Usage:   "Updates an existing product's properties such as display name, description, and\nmetadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Description of the product",
			BodyPath: "description",
		},
		&requestflag.Flag[string]{
			Name:     "display-name",
			Usage:    "Display name of the product",
			BodyPath: "displayName",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Additional metadata for the product",
			BodyPath: "metadata",
		},
		&requestflag.Flag[bool]{
			Name:     "multiple-subscriptions",
			Usage:    "Indicates if multiple subscriptions to this product are allowed",
			BodyPath: "multipleSubscriptions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "product-settings",
			BodyPath: "productSettings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "usage-reset-cutoff-rule",
			Usage:    "Rule defining when usage resets upon subscription update.",
			BodyPath: "usageResetCutoffRule",
		},
	},
	Action:          handleV1ProductsUpdateProduct,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"product-settings": {
		&requestflag.InnerFlag[string]{
			Name:       "product-settings.subscription-cancellation-time",
			Usage:      "Time when the subscription will be cancelled",
			InnerField: "subscriptionCancellationTime",
		},
		&requestflag.InnerFlag[string]{
			Name:       "product-settings.subscription-end-setup",
			Usage:      "Setup for the end of the subscription",
			InnerField: "subscriptionEndSetup",
		},
		&requestflag.InnerFlag[string]{
			Name:       "product-settings.subscription-start-setup",
			Usage:      "Setup for the start of the subscription",
			InnerField: "subscriptionStartSetup",
		},
		&requestflag.InnerFlag[any]{
			Name:       "product-settings.downgrade-plan-id",
			Usage:      "ID of the plan to downgrade to at the end of the billing period",
			InnerField: "downgradePlanId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "product-settings.prorate-at-end-of-billing-period",
			InnerField: "prorateAtEndOfBillingPeriod",
		},
		&requestflag.InnerFlag[any]{
			Name:       "product-settings.subscription-start-plan-id",
			Usage:      "ID of the plan to start the subscription with",
			InnerField: "subscriptionStartPlanId",
		},
	},
	"usage-reset-cutoff-rule": {
		&requestflag.InnerFlag[string]{
			Name:       "usage-reset-cutoff-rule.behavior",
			Usage:      "Behavior of the usage reset cutoff rule",
			InnerField: "behavior",
		},
	},
})

func handleV1ProductsArchiveProduct(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Products.ArchiveProduct(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:products archive-product", obj, format, transform)
}

func handleV1ProductsCreateProduct(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1ProductNewProductParams{}

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
	_, err = client.V1.Products.NewProduct(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:products create-product", obj, format, transform)
}

func handleV1ProductsDuplicateProduct(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1ProductDuplicateProductParams{}

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
	_, err = client.V1.Products.DuplicateProduct(
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
	return ShowJSON(os.Stdout, "v1:products duplicate-product", obj, format, transform)
}

func handleV1ProductsListProducts(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1ProductListProductsParams{}

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
		_, err = client.V1.Products.ListProducts(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:products list-products", obj, format, transform)
	} else {
		iter := client.V1.Products.ListProductsAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:products list-products", iter, format, transform, maxItems)
	}
}

func handleV1ProductsUnarchiveProduct(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Products.UnarchiveProduct(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:products unarchive-product", obj, format, transform)
}

func handleV1ProductsUpdateProduct(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1ProductUpdateProductParams{}

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
	_, err = client.V1.Products.UpdateProduct(
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
	return ShowJSON(os.Stdout, "v1:products update-product", obj, format, transform)
}
