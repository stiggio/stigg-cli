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

var v1CreditsReservationsRelease = cli.Command{
	Name:    "release",
	Usage:   "Releases a credit reservation, consuming the given amount and returning the rest\nof the held credits to the balance. Retry a failed release with the same\nreservation ID; it is applied once.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "reservation-id",
			Required:  true,
			PathParam: "reservationId",
		},
		&requestflag.Flag[float64]{
			Name:     "consumed-amount",
			Usage:    "The amount of credits to consume at release. Up to the reserved amount is charged to the credits that backed the reservation; the rest returns to the balance, and any amount above the reservation is consumed normally.",
			Required: true,
			BodyPath: "consumedAmount",
		},
		&requestflag.Flag[string]{
			Name:     "currency-id",
			Usage:    "The credit currency of the reservation",
			Required: true,
			BodyPath: "currencyId",
		},
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "The customer the reservation was made for",
			Required: true,
			BodyPath: "customerId",
		},
		&requestflag.Flag[string]{
			Name:     "resource-id",
			Usage:    "The resource the reservation was made for, if any",
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
	Action:          handleV1CreditsReservationsRelease,
	HideHelpCommand: true,
}

func handleV1CreditsReservationsRelease(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("reservation-id") && len(unusedArgs) > 0 {
		cmd.Set("reservation-id", unusedArgs[0])
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

	params := stigg.V1CreditReservationReleaseParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Credits.Reservations.Release(
		ctx,
		cmd.Value("reservation-id").(string),
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
		Title:          "v1:credits:reservations release",
		Transform:      transform,
	})
}
