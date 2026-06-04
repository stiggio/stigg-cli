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

var v1EventsDataExportDestinationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Register a destination on the environment's DATA_EXPORT integration.\nLazy-creates the integration row + provider recipient on first call. Idempotent\non destinationId.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "destination-id",
			Usage:    "The provider destination ID returned by the embedded SDK on connect",
			Required: true,
			BodyPath: "destinationId",
		},
		&requestflag.Flag[string]{
			Name:     "destination-type",
			Usage:    "The destination type (e.g. snowflake, bigquery)",
			Required: true,
			BodyPath: "destinationType",
		},
	},
	Action:          handleV1EventsDataExportDestinationsCreate,
	HideHelpCommand: true,
}

var v1EventsDataExportDestinationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Remove a destination from the DATA_EXPORT integration metadata. Idempotent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "destination-id",
			Required:  true,
			PathParam: "destinationId",
		},
	},
	Action:          handleV1EventsDataExportDestinationsDelete,
	HideHelpCommand: true,
}

func handleV1EventsDataExportDestinationsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1EventDataExportDestinationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Events.DataExport.Destinations.New(ctx, params, options...)
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
		Title:          "v1:events:data-export:destinations create",
		Transform:      transform,
	})
}

func handleV1EventsDataExportDestinationsDelete(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("destination-id") && len(unusedArgs) > 0 {
		cmd.Set("destination-id", unusedArgs[0])
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
	_, err = client.V1.Events.DataExport.Destinations.Delete(ctx, cmd.Value("destination-id").(string), options...)
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
		Title:          "v1:events:data-export:destinations delete",
		Transform:      transform,
	})
}
