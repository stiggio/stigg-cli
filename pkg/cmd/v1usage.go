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

var v1UsageHistory = cli.Command{
	Name:    "history",
	Usage:   "Retrieves historical usage data for a customer's metered feature over time.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "customer-id",
			Required:  true,
			PathParam: "customerId",
		},
		&requestflag.Flag[string]{
			Name:      "feature-id",
			Required:  true,
			PathParam: "featureId",
		},
		&requestflag.Flag[any]{
			Name:      "start-date",
			Usage:     "The start date of the range",
			Required:  true,
			QueryPath: "startDate",
		},
		&requestflag.Flag[any]{
			Name:      "end-date",
			Usage:     "The end date of the range",
			QueryPath: "endDate",
		},
		&requestflag.Flag[string]{
			Name:      "group-by",
			Usage:     "Criteria by which to group the usage history",
			QueryPath: "groupBy",
		},
		&requestflag.Flag[*string]{
			Name:      "resource-id",
			Usage:     "Resource id",
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
	Action:          handleV1UsageHistory,
	HideHelpCommand: true,
}

var v1UsageReport = requestflag.WithInnerFlags(cli.Command{
	Name:    "report",
	Usage:   "Reports usage measurements for metered features. The reported usage is used to\ntrack, limit, and bill customer consumption.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "usage",
			Usage:    "A list of usage reports to be submitted in bulk",
			Required: true,
			BodyPath: "usages",
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
	Action:          handleV1UsageReport,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"usage": {
		&requestflag.InnerFlag[string]{
			Name:       "usage.customer-id",
			Usage:      "Customer id",
			InnerField: "customerId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "usage.feature-id",
			Usage:      "Feature id",
			InnerField: "featureId",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "usage.value",
			Usage:      "The value to report for usage",
			InnerField: "value",
		},
		&requestflag.InnerFlag[any]{
			Name:       "usage.created-at",
			Usage:      "Timestamp of when the record was created",
			InnerField: "createdAt",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "usage.dimensions",
			Usage:      "Additional dimensions for the usage report",
			InnerField: "dimensions",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "usage.resource-id",
			Usage:      "Resource id",
			InnerField: "resourceId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "usage.update-behavior",
			Usage:      "The method by which the usage value should be updated",
			InnerField: "updateBehavior",
		},
	},
})

func handleV1UsageHistory(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("feature-id") && len(unusedArgs) > 0 {
		cmd.Set("feature-id", unusedArgs[0])
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

	params := stigg.V1UsageHistoryParams{
		CustomerID: cmd.Value("customer-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Usage.History(
		ctx,
		cmd.Value("feature-id").(string),
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
		Title:          "v1:usage history",
		Transform:      transform,
	})
}

func handleV1UsageReport(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1UsageReportParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Usage.Report(ctx, params, options...)
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
		Title:          "v1:usage report",
		Transform:      transform,
	})
}
