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

var v1BetaCustomersEntitlementsCheck = cli.Command{
	Name:    "check",
	Usage:   "Experimental — request and response shapes may change without notice. Same\nsemantics as `Check entitlement`, plus an optional `dimensions` query param that\nresolves to per-entity governance limits surfaced as `chains` on the response.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[string]{
			Name:      "currency-id",
			Usage:     "Currency ID (refId) to check for credit entitlements. Mutually exclusive with `featureId`.",
			QueryPath: "currencyId",
		},
		&requestflag.Flag[map[string]any]{
			Name:      "dimensions",
			Usage:     "Optional attribution map (e.g. `dimensions[userId]=u1`). When provided, the response includes a `chains` array with per-entity governance limits.",
			QueryPath: "dimensions",
		},
		&requestflag.Flag[string]{
			Name:      "feature-id",
			Usage:     "Feature ID (refId) to check. Mutually exclusive with `currencyId`.",
			QueryPath: "featureId",
		},
		&requestflag.Flag[int64]{
			Name:      "requested-usage",
			Usage:     "Requested usage amount to evaluate against the entitlement limit (numeric features only)",
			QueryPath: "requestedUsage",
		},
		&requestflag.Flag[[]string]{
			Name:      "requested-value",
			Usage:     "Requested values to evaluate against allowed values (enum features only)",
			QueryPath: "requestedValues",
		},
		&requestflag.Flag[string]{
			Name:      "resource-id",
			Usage:     "Resource ID to scope the entitlement check to a specific resource",
			QueryPath: "resourceId",
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
	Action:          handleV1BetaCustomersEntitlementsCheck,
	HideHelpCommand: true,
}

func handleV1BetaCustomersEntitlementsCheck(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1BetaCustomerEntitlementCheckParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1Beta.Customers.Entitlements.Check(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "v1-beta:customers:entitlements check",
		Transform:      transform,
	})
}
