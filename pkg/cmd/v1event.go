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

var v1EventsReport = requestflag.WithInnerFlags(cli.Command{
	Name:    "report",
	Usage:   "Reports raw usage events for event-based metering. Events are ingested\nasynchronously and aggregated into usage totals.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "event",
			Usage:    "A list of usage events to report",
			Required: true,
			BodyPath: "events",
		},
	},
	Action:          handleV1EventsReport,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"event": {
		&requestflag.InnerFlag[string]{
			Name:       "event.customer-id",
			Usage:      "Customer id",
			InnerField: "customerId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.event-name",
			Usage:      "The name of the usage event",
			InnerField: "eventName",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.idempotency-key",
			Usage:      "Idempotency key",
			InnerField: "idempotencyKey",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "event.dimensions",
			Usage:      "Dimensions associated with the usage event",
			InnerField: "dimensions",
		},
		&requestflag.InnerFlag[any]{
			Name:       "event.resource-id",
			Usage:      "Resource id",
			InnerField: "resourceId",
		},
		&requestflag.InnerFlag[any]{
			Name:       "event.timestamp",
			Usage:      "Timestamp",
			InnerField: "timestamp",
		},
	},
})

func handleV1EventsReport(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.V1EventReportParams{}

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
	_, err = client.V1.Events.Report(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "v1:events report", obj, format, transform)
}
