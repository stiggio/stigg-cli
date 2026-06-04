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

var v1EventsDataExportMintScopedToken = cli.Command{
	Name:    "mint-scoped-token",
	Usage:   "Mint a scoped JWT for the FE embedded SDK. Lazy-creates the DATA_EXPORT\nintegration if needed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "application-origin",
			Usage:    "FE origin the resulting JWT is bound to (provider-side anti-fraud)",
			Required: true,
			BodyPath: "applicationOrigin",
		},
		&requestflag.Flag[string]{
			Name:     "destination-type",
			Usage:    "Pin the token to a specific warehouse connect flow",
			BodyPath: "destinationType",
		},
	},
	Action:          handleV1EventsDataExportMintScopedToken,
	HideHelpCommand: true,
}

var v1EventsDataExportTriggerSync = cli.Command{
	Name:    "trigger-sync",
	Usage:   "Trigger a sync for one destination or all destinations under the provider\nentity.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "destination-id",
			Usage:    "Provider destination ID to sync. Omit to sync all destinations.",
			BodyPath: "destinationId",
		},
	},
	Action:          handleV1EventsDataExportTriggerSync,
	HideHelpCommand: true,
}

func handleV1EventsDataExportMintScopedToken(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1EventDataExportMintScopedTokenParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Events.DataExport.MintScopedToken(ctx, params, options...)
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
		Title:          "v1:events:data-export mint-scoped-token",
		Transform:      transform,
	})
}

func handleV1EventsDataExportTriggerSync(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1EventDataExportTriggerSyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Events.DataExport.TriggerSync(ctx, params, options...)
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
		Title:          "v1:events:data-export trigger-sync",
		Transform:      transform,
	})
}
