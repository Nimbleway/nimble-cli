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

var taskAgentCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Web Search Agent instance.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "agent-name",
			Usage:    "Stable agent name.",
			BodyPath: "agent_name",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			Usage:    "Agent description shown to users.",
			BodyPath: "description",
		},
		&requestflag.Flag[*string]{
			Name:     "display-name",
			Usage:    "Human-friendly agent name shown to users.",
			BodyPath: "display_name",
		},
		&requestflag.Flag[*string]{
			Name:     "domain-expertise",
			Usage:    "Domain expertise or operating context for the agent.",
			BodyPath: "domain_expertise",
		},
		&requestflag.Flag[string]{
			Name:     "effort",
			Usage:    "Default effort level for this agent's runs.",
			Default:  "high",
			BodyPath: "effort",
		},
		&requestflag.Flag[[]string]{
			Name:     "goal",
			Usage:    "Ordered goals for the agent to follow.",
			BodyPath: "goals",
		},
		&requestflag.Flag[*string]{
			Name:     "icon",
			Usage:    "Icon identifier used when presenting the agent.",
			BodyPath: "icon",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			Usage:    "Whether the agent can be used to start new runs.",
			Default:  true,
			BodyPath: "is_active",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "output-schema",
			Usage:    "JSON schema describing the structured output the agent should produce.",
			BodyPath: "output_schema",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sources",
			Usage:    "Source guidance for the agent.",
			BodyPath: "sources",
		},
		&requestflag.Flag[[]string]{
			Name:     "suggested-question",
			Usage:    "Suggested prompts users can run with this agent.",
			BodyPath: "suggested_questions",
		},
		&requestflag.Flag[*string]{
			Name:     "template",
			Usage:    "Template name to materialize this instance from. When set, the scalar fields and child rows are copied from the template.",
			BodyPath: "template",
		},
		&requestflag.Flag[*string]{
			Name:     "use-case",
			Usage:    "Primary use case supported by the agent.",
			BodyPath: "use_case",
		},
		&requestflag.Flag[*string]{
			Name:     "workspace-id",
			Usage:    "Workspace identifier to associate with the agent.",
			BodyPath: "workspace_id",
		},
	},
	Action:          handleTaskAgentCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"sources": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sources.allow",
			Usage:      "Source groups the agent is allowed to use.",
			InnerField: "allow",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sources.avoid",
			Usage:      "Free-text guidance describing sources or domains to avoid.",
			InnerField: "avoid",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sources.block",
			Usage:      "Source groups the agent should not use.",
			InnerField: "block",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sources.prioritize",
			Usage:      "Free-text guidance describing sources or domains to prioritize.",
			InnerField: "prioritize",
		},
	},
})

var taskAgentUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update Agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "body",
			Usage:    "A JSON Patch document per RFC 6902 — a JSON array of patch operations.",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleTaskAgentUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"body": {
		&requestflag.InnerFlag[string]{
			Name:       "body.op",
			Usage:      `Allowed values: "add", "remove", "replace", "move", "copy", "test".`,
			InnerField: "op",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "body.from",
			InnerField: "from",
		},
		&requestflag.InnerFlag[any]{
			Name:       "body.value",
			InnerField: "value",
		},
	},
})

var taskAgentList = cli.Command{
	Name:    "list",
	Usage:   "List Web Search Agent instances.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   100,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Default:   0,
			QueryPath: "offset",
		},
		&requestflag.Flag[*string]{
			Name:      "workspace-id",
			QueryPath: "workspace_id",
		},
	},
	Action:          handleTaskAgentList,
	HideHelpCommand: true,
}

var taskAgentDeactivate = cli.Command{
	Name:    "deactivate",
	Usage:   "Deactivate Agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
	},
	Action:          handleTaskAgentDeactivate,
	HideHelpCommand: true,
}

var taskAgentGet = cli.Command{
	Name:    "get",
	Usage:   "Get Agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
	},
	Action:          handleTaskAgentGet,
	HideHelpCommand: true,
}

var taskAgentRun = requestflag.WithInnerFlags(cli.Command{
	Name:    "run",
	Usage:   "Create a research run for a Web Search Agent instance.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:     "input",
			Usage:    "User prompt or task instructions for the run.",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[*string]{
			Name:     "effort",
			Usage:    "Canonical effort tier names for the research graph.",
			BodyPath: "effort",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-events",
			Usage:    "Whether to stream run events when supported.",
			Default:  false,
			BodyPath: "enable_events",
		},
		&requestflag.Flag[any]{
			Name:     "input-data",
			Usage:    "Existing records to ENRICH: a list of partial rows, or a single object, mirroring output_schema's shape.",
			BodyPath: "input_data",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "output-schema",
			Usage:    "JSON schema overriding the agent's default structured output for this run.",
			BodyPath: "output_schema",
		},
		&requestflag.Flag[*string]{
			Name:     "previous-interaction-id",
			Usage:    "Previous interaction identifier used to continue a conversation.",
			BodyPath: "previous_interaction_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sources",
			Usage:    "Source guidance overriding the agent default.",
			BodyPath: "sources",
		},
	},
	Action:          handleTaskAgentRun,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"sources": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sources.allow",
			Usage:      "Source groups the agent is allowed to use.",
			InnerField: "allow",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sources.avoid",
			Usage:      "Free-text guidance describing sources or domains to avoid.",
			InnerField: "avoid",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sources.block",
			Usage:      "Source groups the agent should not use.",
			InnerField: "block",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sources.prioritize",
			Usage:      "Free-text guidance describing sources or domains to prioritize.",
			InnerField: "prioritize",
		},
	},
})

func handleTaskAgentCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.TaskAgentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TaskAgent.New(ctx, params, options...)
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
		Title:          "task-agent create",
		Transform:      transform,
	})
}

func handleTaskAgentUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := githubcomnimblewaynimblego.TaskAgentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TaskAgent.Update(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "task-agent update",
		Transform:      transform,
	})
}

func handleTaskAgentList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.TaskAgentListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TaskAgent.List(ctx, params, options...)
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
		Title:          "task-agent list",
		Transform:      transform,
	})
}

func handleTaskAgentDeactivate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	return client.TaskAgent.Deactivate(ctx, cmd.Value("agent-id").(string), options...)
}

func handleTaskAgentGet(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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
	_, err = client.TaskAgent.Get(ctx, cmd.Value("agent-id").(string), options...)
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
		Title:          "task-agent get",
		Transform:      transform,
	})
}

func handleTaskAgentRun(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := githubcomnimblewaynimblego.TaskAgentRunParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.TaskAgent.Run(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "task-agent run",
		Transform:      transform,
	})
}
