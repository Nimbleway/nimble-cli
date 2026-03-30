// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/Nimbleway/nimble-cli/internal/apiquery"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
	"github.com/Nimbleway/nimble-go"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var agentsList = cli.Command{
	Name:    "list",
	Usage:   "List Agent Templates",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Number of results per page",
			Default:   100,
			QueryPath: "limit",
		},
		&requestflag.Flag[any]{
			Name:      "managed-by",
			Usage:     "Filter templates by attribution",
			QueryPath: "managed_by",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "Pagination offset",
			Default:   0,
			QueryPath: "offset",
		},
		&requestflag.Flag[any]{
			Name:      "privacy",
			Usage:     "Filter by privacy level",
			QueryPath: "privacy",
		},
		&requestflag.Flag[any]{
			Name:      "search",
			Usage:     "Search templates by name, domain, or vertical",
			QueryPath: "search",
		},
	},
	Action:          handleAgentsList,
	HideHelpCommand: true,
}

var agentsGet = cli.Command{
	Name:    "get",
	Usage:   "Get Agent Template",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "template-name",
			Required: true,
		},
	},
	Action:          handleAgentsGet,
	HideHelpCommand: true,
}

var agentsPublish = cli.Command{
	Name:    "publish",
	Usage:   "Publish Agent Version",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-name",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "version-id",
			Required: true,
			BodyPath: "version_id",
		},
	},
	Action:          handleAgentsPublish,
	HideHelpCommand: true,
}

var agentsRun = cli.Command{
	Name:    "run",
	Usage:   "Execute WSA Realtime Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent",
			Required: true,
			BodyPath: "agent",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Required: true,
			BodyPath: "params",
		},
		&requestflag.Flag[[]string]{
			Name:     "format",
			Usage:    "Response formats to include. All disabled by default.",
			BodyPath: "formats",
		},
		&requestflag.Flag[bool]{
			Name:     "localization",
			Default:  false,
			BodyPath: "localization",
		},
	},
	Action:          handleAgentsRun,
	HideHelpCommand: true,
}

var agentsRunAsync = cli.Command{
	Name:    "run-async",
	Usage:   "Execute WSA Async Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent",
			Required: true,
			BodyPath: "agent",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Required: true,
			BodyPath: "params",
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "URL to call back when async operation completes",
			BodyPath: "callback_url",
		},
		&requestflag.Flag[[]string]{
			Name:     "format",
			Usage:    "Response formats to include. All disabled by default.",
			BodyPath: "formats",
		},
		&requestflag.Flag[bool]{
			Name:     "localization",
			Default:  false,
			BodyPath: "localization",
		},
		&requestflag.Flag[bool]{
			Name:     "storage-compress",
			Usage:    "Whether to compress stored data",
			BodyPath: "storage_compress",
		},
		&requestflag.Flag[string]{
			Name:     "storage-object-name",
			Usage:    "Custom name for the stored object",
			BodyPath: "storage_object_name",
		},
		&requestflag.Flag[string]{
			Name:     "storage-type",
			Usage:    "Type of storage to use for results",
			BodyPath: "storage_type",
		},
		&requestflag.Flag[string]{
			Name:     "storage-url",
			Usage:    "URL for storage location",
			BodyPath: "storage_url",
		},
	},
	Action:          handleAgentsRunAsync,
	HideHelpCommand: true,
}

var agentsRunBatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "run-batch",
	Usage:   "Execute WSA Batch Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "input",
			Required: true,
			BodyPath: "inputs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "shared-inputs",
			Required: true,
			BodyPath: "shared_inputs",
		},
	},
	Action:          handleAgentsRunBatch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"input": {
		&requestflag.InnerFlag[[]string]{
			Name:       "input.formats",
			Usage:      "Response formats to include. All disabled by default.",
			InnerField: "formats",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "input.localization",
			InnerField: "localization",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "input.params",
			InnerField: "params",
		},
	},
	"shared-inputs": {
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.agent",
			InnerField: "agent",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "shared-inputs.formats",
			Usage:      "Response formats to include. All disabled by default.",
			InnerField: "formats",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "shared-inputs.localization",
			InnerField: "localization",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "shared-inputs.params",
			InnerField: "params",
		},
	},
})

func handleAgentsList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.AgentListParams{}

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
	_, err = client.Agents.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents list", obj, format, transform)
}

func handleAgentsGet(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("template-name") && len(unusedArgs) > 0 {
		cmd.Set("template-name", unusedArgs[0])
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
	_, err = client.Agents.Get(ctx, cmd.Value("template-name").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents get", obj, format, transform)
}

func handleAgentsPublish(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-name") && len(unusedArgs) > 0 {
		cmd.Set("agent-name", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.AgentPublishParams{}

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
	_, err = client.Agents.Publish(
		ctx,
		cmd.Value("agent-name").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents publish", obj, format, transform)
}

func handleAgentsRun(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.AgentRunParams{}

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
	_, err = client.Agents.Run(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents run", obj, format, transform)
}

func handleAgentsRunAsync(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.AgentRunAsyncParams{}

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
	_, err = client.Agents.RunAsync(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents run-async", obj, format, transform)
}

func handleAgentsRunBatch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.AgentRunBatchParams{}

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
	_, err = client.Agents.RunBatch(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents run-batch", obj, format, transform)
}
