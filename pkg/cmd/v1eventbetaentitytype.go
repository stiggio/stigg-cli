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

var v1EventsBetaEntityTypesList = cli.Command{
	Name:    "list",
	Usage:   "Returns a cursor-paginated list of entity types defined in the environment.\nEntity types are vendor-defined categories of resource that can be governed\n(e.g. Org, Team, User).",
	Suggest: true,
	Flags: []cli.Flag{
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1EventsBetaEntityTypesList,
	HideHelpCommand: true,
}

var v1EventsBetaEntityTypesUpsert = requestflag.WithInnerFlags(cli.Command{
	Name:    "upsert",
	Usage:   "Batched create-or-update of entity types. Existing types matched by id are\nupdated; new ids are created. Idempotent — re-submitting the same payload\nconverges to the same state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "type",
			Usage:    "Entity types to upsert (1–100 per request)",
			Required: true,
			BodyPath: "types",
		},
	},
	Action:          handleV1EventsBetaEntityTypesUpsert,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"type": {
		&requestflag.InnerFlag[string]{
			Name:       "type.id",
			Usage:      "The unique identifier for the entity",
			InnerField: "id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "type.attribution-keys",
			Usage:      `Dimension keys used to attribute usage events to instances of this type (e.g. ["orgId"]). Empty array means no attribution.`,
			InnerField: "attributionKeys",
		},
		&requestflag.InnerFlag[string]{
			Name:       "type.display-name",
			Usage:      "The display name for the entity type",
			InnerField: "displayName",
		},
	},
})

func handleV1EventsBetaEntityTypesList(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := stigg.V1EventBetaEntityTypeListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Events.Beta.EntityTypes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:events:beta:entity-types list",
			Transform:      transform,
		})
	} else {
		iter := client.V1.Events.Beta.EntityTypes.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:events:beta:entity-types list",
			Transform:      transform,
		})
	}
}

func handleV1EventsBetaEntityTypesUpsert(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1EventBetaEntityTypeUpsertParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Events.Beta.EntityTypes.Upsert(ctx, params, options...)
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
		Title:          "v1:events:beta:entity-types upsert",
		Transform:      transform,
	})
}
