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

var v1BetaCustomersRetrieveGovernance = cli.Command{
	Name:    "retrieve-governance",
	Usage:   "Queries the customer's governance hierarchy tree, returning a cursor-paginated\nlist of nodes with their usage configuration (limit, cadence, scope) and current\nusage, sortable and filterable by usage. Each node carries `parentId` so the\ntree can be rebuilt client-side. Usage is read from a periodically-refreshed\nread model and never gates access.",
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
		&requestflag.Flag[[]string]{
			Name:      "currency-id",
			Usage:     "Currency ids to include, repeated per value (e.g. `?currencyIds=credits`). Omit both featureIds and currencyIds for tree mode.",
			QueryPath: "currencyIds",
		},
		&requestflag.Flag[string]{
			Name:      "entity-id-search",
			Usage:     "Case-insensitive substring match on the entity id (`%`/`_` matched literally).",
			QueryPath: "entityIdSearch",
		},
		&requestflag.Flag[[]string]{
			Name:      "entity-type-id",
			Usage:     "Filter to one or more entity types, repeated per value (e.g. `?entityTypeIds=team&entityTypeIds=user`).",
			QueryPath: "entityTypeIds",
		},
		&requestflag.Flag[[]string]{
			Name:      "feature-id",
			Usage:     "Feature ids to include, repeated per value (e.g. `?featureIds=ai-tokens&featureIds=seats`). Omit both featureIds and currencyIds for tree mode — every node in the hierarchy with no usage configuration attached.",
			QueryPath: "featureIds",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[float64]{
			Name:      "min-utilization",
			Usage:     "Only nodes with utilization ≥ this value (e.g. 0.8 for ≥80%, 1 for at/over limit).",
			QueryPath: "minUtilization",
		},
		&requestflag.Flag[string]{
			Name:      "order",
			Usage:     "Sort direction: `asc` or `desc` (default `desc`).",
			Default:   "desc",
			QueryPath: "order",
		},
		&requestflag.Flag[string]{
			Name:      "scope",
			Usage:     "Filter by configuration scope: `all` (default), `nodeWide` (`[]` only), or `scoped` (non-empty only).",
			Default:   "all",
			QueryPath: "scope",
		},
		&requestflag.Flag[string]{
			Name:      "sort-by",
			Usage:     "Sort key: `utilization` (default, cross-capability-safe), `currentUsage`, `usageLimit`, `scopeSize`, `id`, or `createdAt`.",
			Default:   "utilization",
			QueryPath: "sortBy",
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
	Action:          handleV1BetaCustomersRetrieveGovernance,
	HideHelpCommand: true,
}

func handleV1BetaCustomersRetrieveGovernance(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1BetaCustomerGetGovernanceParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1Beta.Customers.GetGovernance(
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
		Title:          "v1-beta:customers retrieve-governance",
		Transform:      transform,
	})
}
