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

var internalBetaEventQueuesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Get event queue by queue name",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
	},
	Action:          handleInternalBetaEventQueuesRetrieve,
	HideHelpCommand: true,
}

var internalBetaEventQueuesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update event queue configuration",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-assume-role-arn",
			BodyPath: "allowedAssumeRoleArns",
		},
		&requestflag.Flag[bool]{
			Name:     "create-low-priority-queues",
			Usage:    "Whether to create separate low-priority queues for standard topic events",
			BodyPath: "createLowPriorityQueues",
		},
		&requestflag.Flag[[]string]{
			Name:     "event-type",
			BodyPath: "eventTypes",
		},
	},
	Action:          handleInternalBetaEventQueuesUpdate,
	HideHelpCommand: true,
}

var internalBetaEventQueuesList = cli.Command{
	Name:            "list",
	Usage:           "List all event queues for the current environment",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleInternalBetaEventQueuesList,
	HideHelpCommand: true,
}

var internalBetaEventQueuesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an event queue and tear down its infrastructure",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "queue-name",
			Required: true,
		},
	},
	Action:          handleInternalBetaEventQueuesDelete,
	HideHelpCommand: true,
}

var internalBetaEventQueuesProvision = cli.Command{
	Name:    "provision",
	Usage:   "Provision SQS queue, SNS subscriptions, and IAM role for the current environment",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "AWS region for the SQS queue (e.g., us-east-1, eu-west-1)",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-assume-role-arn",
			Usage:    "Additional IAM role ARNs allowed to assume the external role for queue access",
			BodyPath: "allowedAssumeRoleArns",
		},
		&requestflag.Flag[bool]{
			Name:     "create-low-priority-queues",
			Usage:    "Whether to create separate low-priority queues for standard topic events",
			Default:  false,
			BodyPath: "createLowPriorityQueues",
		},
		&requestflag.Flag[[]string]{
			Name:     "event-type",
			Usage:    "Event types to subscribe to. Defaults to entitlements, measurements, and migrations.",
			Default:  []string{"ENTITLEMENTS_UPDATED", "MEASUREMENT_REPORTED", "SUBSCRIPTIONS_MIGRATED"},
			BodyPath: "eventTypes",
		},
		&requestflag.Flag[string]{
			Name:     "suffix",
			Usage:    "Optional suffix to allow multiple queues for the same environment and region",
			BodyPath: "suffix",
		},
	},
	Action:          handleInternalBetaEventQueuesProvision,
	HideHelpCommand: true,
}

func handleInternalBetaEventQueuesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-name") && len(unusedArgs) > 0 {
		cmd.Set("queue-name", unusedArgs[0])
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
	_, err = client.Internal.Beta.EventQueues.Get(ctx, cmd.Value("queue-name").(string), options...)
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
		Title:          "internal:beta:event-queues retrieve",
		Transform:      transform,
	})
}

func handleInternalBetaEventQueuesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-name") && len(unusedArgs) > 0 {
		cmd.Set("queue-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.InternalBetaEventQueueUpdateParams{}

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
	_, err = client.Internal.Beta.EventQueues.Update(
		ctx,
		cmd.Value("queue-name").(string),
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
		Title:          "internal:beta:event-queues update",
		Transform:      transform,
	})
}

func handleInternalBetaEventQueuesList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Internal.Beta.EventQueues.List(ctx, options...)
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
		Title:          "internal:beta:event-queues list",
		Transform:      transform,
	})
}

func handleInternalBetaEventQueuesDelete(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("queue-name") && len(unusedArgs) > 0 {
		cmd.Set("queue-name", unusedArgs[0])
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
	_, err = client.Internal.Beta.EventQueues.Delete(ctx, cmd.Value("queue-name").(string), options...)
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
		Title:          "internal:beta:event-queues delete",
		Transform:      transform,
	})
}

func handleInternalBetaEventQueuesProvision(ctx context.Context, cmd *cli.Command) error {
	client := stigg.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := stigg.InternalBetaEventQueueProvisionParams{}

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
	_, err = client.Internal.Beta.EventQueues.Provision(ctx, params, options...)
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
		Title:          "internal:beta:event-queues provision",
		Transform:      transform,
	})
}
