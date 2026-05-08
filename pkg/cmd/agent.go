// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/Nimbleway/nimble-cli/internal/apiquery"
	"github.com/Nimbleway/nimble-cli/internal/requestflag"
	"github.com/Nimbleway/nimble-go"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var agentList = cli.Command{
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
		&requestflag.Flag[*string]{
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
		&requestflag.Flag[*string]{
			Name:      "privacy",
			Usage:     "Filter by privacy level",
			QueryPath: "privacy",
		},
		&requestflag.Flag[*string]{
			Name:      "search",
			Usage:     "Search templates by name, domain, or vertical",
			QueryPath: "search",
		},
	},
	Action:          handleAgentList,
	HideHelpCommand: true,
}

var agentGenerate = requestflag.WithInnerFlags(cli.Command{
	Name:    "generate",
	Usage:   "Create Agent Generation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			BodyPath: "url",
		},
		&requestflag.Flag[*string]{
			Name:     "agent-name",
			BodyPath: "agent_name",
		},
		&requestflag.Flag[any]{
			Name:     "input-schema",
			BodyPath: "input_schema",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "output-schema",
			BodyPath: "output_schema",
		},
		&requestflag.Flag[string]{
			Name:     "from-agent",
			BodyPath: "from_agent",
		},
	},
	Action:          handleAgentGenerate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"metadata": {
		&requestflag.InnerFlag[*string]{
			Name:       "metadata.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "metadata.display-name",
			InnerField: "display_name",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "metadata.tags",
			InnerField: "tags",
		},
	},
})

var agentGet = cli.Command{
	Name:    "get",
	Usage:   "Get Agent Template",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "template-name",
			Required:  true,
			PathParam: "template_name",
		},
	},
	Action:          handleAgentGet,
	HideHelpCommand: true,
}

var agentGetGeneration = cli.Command{
	Name:    "get-generation",
	Usage:   "Get Agent Generation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "generation-id",
			Required:  true,
			PathParam: "generation_id",
		},
	},
	Action:          handleAgentGetGeneration,
	HideHelpCommand: true,
}

var agentPublish = cli.Command{
	Name:    "publish",
	Usage:   "Publish Agent Version",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-name",
			Required:  true,
			PathParam: "agent_name",
		},
		&requestflag.Flag[string]{
			Name:     "version-id",
			Required: true,
			BodyPath: "version_id",
		},
	},
	Action:          handleAgentPublish,
	HideHelpCommand: true,
}

var agentRun = cli.Command{
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
	Action:          handleAgentRun,
	HideHelpCommand: true,
}

var agentRunAsync = cli.Command{
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
	Action:          handleAgentRunAsync,
	HideHelpCommand: true,
}

var agentRunBatch = requestflag.WithInnerFlags(cli.Command{
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
	Action:          handleAgentRunBatch,
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

func handleAgentList(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := githubcomnimblewaynimblego.AgentListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agent.List(ctx, params, options...)
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
		Title:          "agent list",
		Transform:      transform,
	})
}

func handleAgentGenerate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := githubcomnimblewaynimblego.AgentGenerateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agent.Generate(ctx, params, options...)
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
		Title:          "agent generate",
		Transform:      transform,
	})
}

func handleAgentGet(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agent.Get(ctx, cmd.Value("template-name").(string), options...)
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
		Title:          "agent get",
		Transform:      transform,
	})
}

func handleAgentGetGeneration(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("generation-id") && len(unusedArgs) > 0 {
		cmd.Set("generation-id", unusedArgs[0])
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
	_, err = client.Agent.GetGeneration(ctx, cmd.Value("generation-id").(string), options...)
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
		Title:          "agent get-generation",
		Transform:      transform,
	})
}

func handleAgentPublish(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-name") && len(unusedArgs) > 0 {
		cmd.Set("agent-name", unusedArgs[0])
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

	params := githubcomnimblewaynimblego.AgentPublishParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agent.Publish(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "agent publish",
		Transform:      transform,
	})
}

func handleAgentRun(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := githubcomnimblewaynimblego.AgentRunParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agent.Run(ctx, params, options...)
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
		Title:          "agent run",
		Transform:      transform,
	})
}

func handleAgentRunAsync(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := githubcomnimblewaynimblego.AgentRunAsyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agent.RunAsync(ctx, params, options...)
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
		Title:          "agent run-async",
		Transform:      transform,
	})
}

func handleAgentRunBatch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := githubcomnimblewaynimblego.AgentRunBatchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agent.RunBatch(ctx, params, options...)
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
		Title:          "agent run-batch",
		Transform:      transform,
	})
}
