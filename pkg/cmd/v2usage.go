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

var v2UsageReport = requestflag.WithInnerFlags(cli.Command{
	Name:    "report",
	Usage:   "Reports usage measurements for metered features. The reported usage is used to\ntrack, limit, and bill customer consumption. Credit balances are returned once\nper wallet in `credits`, not on each measurement.",
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
	Action:          handleV2UsageReport,
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
		&requestflag.InnerFlag[float64]{
			Name:       "usage.value",
			Usage:      "The value to report for usage. Fractional values are accepted and rounded half away from zero to 5 decimal places. A value with a fractional part must have an absolute value below 34,359,738,368 (2^35); whole numbers are accepted up to ±9,007,199,254,740,991.",
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
		&requestflag.InnerFlag[string]{
			Name:       "usage.idempotency-key",
			Usage:      "A key you provide to safely retry the same usage report without double-counting it. Reports with a previously-seen idempotency key are deduplicated for 7 days; after that window a retry is treated as new usage.",
			InnerField: "idempotencyKey",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "usage.resource-id",
			Usage:      "The customer resource this usage applies to. Optional — only required if the customer has multiple resources (for example, one subscription per workspace or site) and usage needs to be tracked separately per resource; omit it to report usage at the customer level.",
			InnerField: "resourceId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "usage.update-behavior",
			Usage:      "How the reported value is applied: DELTA (default) adds it to the feature's current usage; SET treats it as the new absolute usage total, and Stigg computes the delta internally.",
			InnerField: "updateBehavior",
		},
	},
})

func handleV2UsageReport(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V2UsageReportParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V2.Usage.Report(ctx, params, options...)
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
		Title:          "v2:usage report",
		Transform:      transform,
	})
}
