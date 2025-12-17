// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/stigg-cli/internal/apiquery"
	"github.com/stainless-sdks/stigg-cli/internal/requestflag"
	"github.com/stainless-sdks/stigg-go"
	"github.com/stainless-sdks/stigg-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var v1CustomersCreate = cli.Command{
	Name:  "create",
	Usage: "Create a new Customer",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "The email of the customer",
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "external-id",
			Usage:    "Customer slug",
			BodyPath: "externalId",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the customer",
			BodyPath: "name",
		},
		&requestflag.Flag[any]{
			Name:     "default-payment-method",
			Usage:    "The default payment method details",
			BodyPath: "defaultPaymentMethod",
		},
		&requestflag.Flag[[]any]{
			Name:     "integration",
			Usage:    "List of integrations",
			BodyPath: "integrations",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Additional metadata",
			BodyPath: "metadata",
		},
	},
	Action:          handleV1CustomersCreate,
	HideHelpCommand: true,
}

var v1CustomersRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get a single Customer by id",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "id",
		},
	},
	Action:          handleV1CustomersRetrieve,
	HideHelpCommand: true,
}

var v1CustomersUpdate = cli.Command{
	Name:  "update",
	Usage: "Update an existing Customer",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "id",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "The email of the customer",
			BodyPath: "email",
		},
		&requestflag.Flag[[]any]{
			Name:     "integration",
			Usage:    "List of integrations",
			BodyPath: "integrations",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Additional metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "The name of the customer",
			BodyPath: "name",
		},
	},
	Action:          handleV1CustomersUpdate,
	HideHelpCommand: true,
}

var v1CustomersList = cli.Command{
	Name:  "list",
	Usage: "Get a list of Customers",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "ending-before",
			Usage:     "Ending before this UUID for pagination",
			QueryPath: "ending_before",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Items per page",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "starting-after",
			Usage:     "Starting after this UUID for pagination",
			QueryPath: "starting_after",
		},
	},
	Action:          handleV1CustomersList,
	HideHelpCommand: true,
}

var v1CustomersArchive = cli.Command{
	Name:  "archive",
	Usage: "Perform archive on a Customer",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "id",
		},
	},
	Action:          handleV1CustomersArchive,
	HideHelpCommand: true,
}

var v1CustomersUnarchive = cli.Command{
	Name:  "unarchive",
	Usage: "Perform unarchive on a Customer",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "id",
		},
	},
	Action:          handleV1CustomersUnarchive,
	HideHelpCommand: true,
}

func handleV1CustomersCreate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerNewParams{}

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
	_, err = client.V1.Customers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:customers create", obj, format, transform)
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
		return streamOutput("v1:customers list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "v1:customers list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
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
