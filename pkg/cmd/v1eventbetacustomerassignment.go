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

var v1EventsBetaCustomersAssignmentsList = cli.Command{
	Name:    "list",
	Usage:   "Returns a cursor-paginated list of capability assignments for the given\ncustomer. An assignment ties an entity to a capability with a usage limit and\nreset cadence.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
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
		&requestflag.Flag[string]{
			Name:      "capability-id",
			Usage:     "Filter assignments to a specific capability refId",
			QueryPath: "capabilityId",
		},
		&requestflag.Flag[string]{
			Name:      "entity-id",
			Usage:     "Filter assignments to a specific entity refId",
			QueryPath: "entityId",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:       "x-account-id",
			HeaderPath: "X-ACCOUNT-ID",
		},
		&requestflag.Flag[string]{
			Name:       "x-environment-id",
			HeaderPath: "X-ENVIRONMENT-ID",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleV1EventsBetaCustomersAssignmentsList,
	HideHelpCommand: true,
}

var v1EventsBetaCustomersAssignmentsUpsert = requestflag.WithInnerFlags(cli.Command{
	Name:    "upsert",
	Usage:   "Batched create-or-update of capability assignments. Existing assignments matched\nby (entityId, capabilityId) are updated; new pairs are created. On update,\nomitted fields (usageLimit, cadence) are preserved; on create both are required\nby the governance service.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "assignment",
			Usage:    "Assignments to upsert (1–100 per request)",
			Required: true,
			BodyPath: "assignments",
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
	Action:          handleV1EventsBetaCustomersAssignmentsUpsert,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"assignment": {
		&requestflag.InnerFlag[string]{
			Name:       "assignment.capability-id",
			Usage:      "The capability refId this assignment grants",
			InnerField: "capabilityId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assignment.entity-id",
			Usage:      "The entity refId this assignment is attached to",
			InnerField: "entityId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "assignment.cadence",
			Usage:      "Usage-reset cadence (required on create). Currently only `MONTH` is supported",
			InnerField: "cadence",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "assignment.usage-limit",
			Usage:      "Maximum usage allowed within one cadence window (required on create)",
			InnerField: "usageLimit",
		},
	},
})

func handleV1EventsBetaCustomersAssignmentsList(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1EventBetaCustomerAssignmentListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Events.Beta.Customers.Assignments.List(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:events:beta:customers:assignments list",
			Transform:      transform,
		})
	} else {
		iter := client.V1.Events.Beta.Customers.Assignments.ListAutoPaging(
			ctx,
			cmd.Value("id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:events:beta:customers:assignments list",
			Transform:      transform,
		})
	}
}

func handleV1EventsBetaCustomersAssignmentsUpsert(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := stigg.V1EventBetaCustomerAssignmentUpsertParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Events.Beta.Customers.Assignments.Upsert(
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
		Title:          "v1:events:beta:customers:assignments upsert",
		Transform:      transform,
	})
}
