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

var v1CouponsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new discount coupon with percentage or fixed amount off, applicable to\ncustomer subscriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The unique identifier for the entity",
			Required: true,
			BodyPath: "id",
		},
		&requestflag.Flag[any]{
			Name:     "amounts-off",
			Usage:    "Fixed amount discounts in different currencies",
			Required: true,
			BodyPath: "amountsOff",
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Description of the coupon",
			Required: true,
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "duration-in-months",
			Usage:    "Duration of the coupon validity in months",
			Required: true,
			BodyPath: "durationInMonths",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Metadata associated with the entity",
			Required: true,
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the coupon",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "percent-off",
			Usage:    "Percentage discount off the original price",
			Required: true,
			BodyPath: "percentOff",
		},
	},
	Action:          handleV1CouponsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"amounts-off": {
		&requestflag.InnerFlag[float64]{
			Name:       "amounts-off.amount",
			Usage:      "The price amount",
			InnerField: "amount",
		},
		&requestflag.InnerFlag[string]{
			Name:       "amounts-off.currency",
			Usage:      "The price currency",
			InnerField: "currency",
		},
	},
})

var v1CouponsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a coupon by its unique identifier.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1CouponsRetrieve,
	HideHelpCommand: true,
}

var v1CouponsList = requestflag.WithInnerFlags(cli.Command{
	Name:    "list",
	Usage:   "Retrieves a paginated list of coupons in the environment.",
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
			Usage:     "Filter by coupon status. Supports comma-separated values for multiple statuses",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Filter by coupon type (FIXED or PERCENTAGE)",
			QueryPath: "type",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1CouponsList,
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

var v1CouponsArchiveCoupon = cli.Command{
	Name:    "archive-coupon",
	Usage:   "Archives a coupon, preventing it from being applied to new subscriptions.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleV1CouponsArchiveCoupon,
	HideHelpCommand: true,
}

var v1CouponsUpdateCoupon = cli.Command{
	Name:    "update-coupon",
	Usage:   "Updates an existing coupon's properties such as name, description, and metadata.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "description",
			Usage:    "Description of the coupon",
			BodyPath: "description",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Metadata associated with the entity",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the coupon",
			BodyPath: "name",
		},
	},
	Action:          handleV1CouponsUpdateCoupon,
	HideHelpCommand: true,
}

func handleV1CouponsCreate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CouponNewParams{}

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
	_, err = client.V1.Coupons.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:coupons create", obj, format, transform)
}

func handleV1CouponsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Coupons.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:coupons retrieve", obj, format, transform)
}

func handleV1CouponsList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CouponListParams{}

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
		_, err = client.V1.Coupons.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "v1:coupons list", obj, format, transform)
	} else {
		iter := client.V1.Coupons.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "v1:coupons list", iter, format, transform, maxItems)
	}
}

func handleV1CouponsArchiveCoupon(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Coupons.ArchiveCoupon(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:coupons archive-coupon", obj, format, transform)
}

func handleV1CouponsUpdateCoupon(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CouponUpdateCouponParams{}

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
	_, err = client.V1.Coupons.UpdateCoupon(
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
	return ShowJSON(os.Stdout, "v1:coupons update-coupon", obj, format, transform)
}
