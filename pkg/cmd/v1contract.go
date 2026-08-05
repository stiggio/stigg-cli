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

var v1ContractsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a contract for a customer together with all of its (custom)\nsubscriptions in a single atomic operation. Every new subscription is created\ninside one transaction — any validation or creation failure rolls the whole\ncontract back. Each subscription entry is either a new subscription to create or\na reference to an existing custom subscription. Returns the created contract.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "customer-id",
			Usage:    "The customer ref ID the contract belongs to",
			Required: true,
			BodyPath: "customerId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "subscription",
			Usage:    "The subscriptions to attach to the contract (must be non-empty). Each entry is either a new subscription to create or a reference to an existing custom subscription.",
			Required: true,
			BodyPath: "subscriptions",
		},
		&requestflag.Flag[any]{
			Name:     "activation-end-date",
			Usage:    "Optional contract activation end date",
			BodyPath: "activationEndDate",
		},
		&requestflag.Flag[any]{
			Name:     "activation-start-date",
			Usage:    "Optional contract activation start date",
			BodyPath: "activationStartDate",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Optional contract name",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "po-number",
			Usage:    "Optional purchase-order number",
			BodyPath: "poNumber",
		},
		&requestflag.Flag[bool]{
			Name:     "setup-billing",
			Usage:    "Whether to set up billing for the contract by creating a billing contract in the connected billing provider. When false, the contract only provisions access (grants entitlements) and no billing contract is created. Defaults to true.",
			Default:  true,
			BodyPath: "setupBilling",
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
	Action:          handleV1ContractsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"subscription": {
		&requestflag.InnerFlag[string]{
			Name:       "subscription.existing-subscription-id",
			Usage:      "The subscription ref ID of an already-created custom subscription to link",
			InnerField: "existingSubscriptionId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "subscription.new-subscription",
			Usage:      "A new subscription to create, using the same body the provision-subscription endpoint accepts",
			InnerField: "newSubscription",
		},
	},
})

var v1ContractsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single contract by its ID, enriched with a preview of its upcoming\n(next) invoice when one is available. Returns 404 when no contract with that ID\nexists in the environment.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
	Action:          handleV1ContractsRetrieve,
	HideHelpCommand: true,
}

var v1ContractsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Updates a contract's metadata (name, PO number, activation dates) and optionally\nre-links its subscriptions. Best-effort re-syncs the change to the connected\nbilling provider.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[any]{
			Name:     "activation-end-date",
			Usage:    "New activation end date",
			BodyPath: "activationEndDate",
		},
		&requestflag.Flag[any]{
			Name:     "activation-start-date",
			Usage:    "New activation start date",
			BodyPath: "activationStartDate",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "New contract name",
			BodyPath: "name",
		},
		&requestflag.Flag[*string]{
			Name:     "po-number",
			Usage:    "New purchase-order number",
			BodyPath: "poNumber",
		},
		&requestflag.Flag[bool]{
			Name:     "setup-billing",
			Usage:    "Enable billing on a provision-access-only contract by creating a billing contract in the connected billing provider. Only takes effect when true and the contract has no billing yet; omitting it leaves billing unchanged. Billing is never removed by an update.",
			BodyPath: "setupBilling",
		},
		&requestflag.Flag[[]string]{
			Name:     "subscription-id",
			Usage:    "When provided, replaces the set of subscriptions linked to the contract (subscription ref IDs)",
			BodyPath: "subscriptionIds",
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
	Action:          handleV1ContractsUpdate,
	HideHelpCommand: true,
}

var v1ContractsList = cli.Command{
	Name:    "list",
	Usage:   "Retrieves a cursor-paginated list of contracts in the environment, fetched live\nfrom the connected billing provider. Each contract is enriched with a preview of\nits upcoming (next) invoice when one is available. Returns an empty list when no\nbilling provider is connected. Supports filtering by customer external ID,\nstate, and name.",
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
		&requestflag.Flag[string]{
			Name:      "customer-external-id",
			Usage:     "Filter by the exact external ID of the customer the contract belongs to",
			QueryPath: "customerExternalId",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Filter by exact contract name",
			QueryPath: "name",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Filter by contract state. Supports comma-separated values for multiple states",
			QueryPath: "state",
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
	Action:          handleV1ContractsList,
	HideHelpCommand: true,
}

var v1ContractsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a contract: cancels the contract in the connected billing provider and\ncancels every subscription linked to it.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
	Action:          handleV1ContractsDelete,
	HideHelpCommand: true,
}

func handleV1ContractsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1ContractNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Contracts.New(ctx, params, options...)
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
		Title:          "v1:contracts create",
		Transform:      transform,
	})
}

func handleV1ContractsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1ContractGetParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Contracts.Get(
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
		Title:          "v1:contracts retrieve",
		Transform:      transform,
	})
}

func handleV1ContractsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1ContractUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Contracts.Update(
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
		Title:          "v1:contracts update",
		Transform:      transform,
	})
}

func handleV1ContractsList(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1ContractListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.V1.Contracts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:contracts list",
			Transform:      transform,
		})
	} else {
		iter := client.V1.Contracts.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "v1:contracts list",
			Transform:      transform,
		})
	}
}

func handleV1ContractsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := stigg.V1ContractDeleteParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.V1.Contracts.Delete(
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
		Title:          "v1:contracts delete",
		Transform:      transform,
	})
}
