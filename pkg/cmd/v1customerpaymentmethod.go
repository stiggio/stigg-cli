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

var v1CustomersPaymentMethodAttach = cli.Command{
	Name:  "attach",
	Usage: "Perform payment-method attachment on a Customer",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "id",
		},
		&requestflag.Flag[string]{
			Name:     "integration-id",
			Usage:    "Integration details",
			BodyPath: "integrationId",
		},
		&requestflag.Flag[string]{
			Name:     "payment-method-id",
			Usage:    "Billing provider payment method id",
			BodyPath: "paymentMethodId",
		},
		&requestflag.Flag[string]{
			Name:     "vendor-identifier",
			Usage:    "The vendor identifier of integration",
			BodyPath: "vendorIdentifier",
		},
		&requestflag.Flag[string]{
			Name:     "billing-currency",
			BodyPath: "billingCurrency",
		},
	},
	Action:          handleV1CustomersPaymentMethodAttach,
	HideHelpCommand: true,
}

var v1CustomersPaymentMethodDetach = cli.Command{
	Name:  "detach",
	Usage: "Perform payment-method detachment on a Customer",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "id",
		},
	},
	Action:          handleV1CustomersPaymentMethodDetach,
	HideHelpCommand: true,
}

func handleV1CustomersPaymentMethodAttach(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1CustomerPaymentMethodAttachParams{}

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
	_, err = client.V1.Customers.PaymentMethod.Attach(
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
	return ShowJSON(os.Stdout, "v1:customers:payment-method attach", obj, format, transform)
}

func handleV1CustomersPaymentMethodDetach(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.V1.Customers.PaymentMethod.Detach(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:customers:payment-method detach", obj, format, transform)
}
