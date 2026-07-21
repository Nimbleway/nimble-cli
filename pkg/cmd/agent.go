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

var agentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a Web Search Agent. Either pass `template` to materialize a pre-built\ntemplate (its fields, goals, sources, and suggested questions are copied), or\ndefine the agent from scratch with `display_name`, `goals`, `sources`, and an\noptional `output_schema` for structured results.",
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
		&requestflag.Flag[*string]{
			Name:     "skill",
			Usage:    "Skill or operating context for the agent.",
			BodyPath: "skill",
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
	},
	Action:          handleAgentsCreate,
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

var agentsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an agent with a\n[JSON Patch](https://datatracker.ietf.org/doc/html/rfc6902) document — an array\nof `{op, path, value}` operations applied to the agent, e.g.\n`[{\"op\": \"replace\", \"path\": \"/display_name\", \"value\": \"My agent\"}]`. Returns the\nupdated agent.",
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
	Action:          handleAgentsUpdate,
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

var agentsList = cli.Command{
	Name:    "list",
	Usage:   "List the active Web Search Agents in your account. Results are scoped to the\nworkspace resolved from your token (or the optional `workspace_id` query\nparameter) and paginated with `offset`/`limit`.",
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
	Action:          handleAgentsList,
	HideHelpCommand: true,
}

var agentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deactivate an agent. This is a soft delete: the agent can no longer start new\nruns, but its existing runs and their results remain retrievable.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
	},
	Action:          handleAgentsDelete,
	HideHelpCommand: true,
}

var agentsGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve a single Web Search Agent by ID.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
	},
	Action:          handleAgentsGet,
	HideHelpCommand: true,
}

var agentsRun = requestflag.WithInnerFlags(cli.Command{
	Name:    "run",
	Usage:   "Creates a minimal persistent Web Search Agent and starts a run for it. The\nresponse includes `web_search_agent_id` for later agent and run queries.",
	Suggest: true,
	Flags: []cli.Flag{
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
	Action:          handleAgentsRun,
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

func handleAgentsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.AgentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.New(ctx, params, options...)
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
		Title:          "agents create",
		Transform:      transform,
	})
}

func handleAgentsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.AgentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Update(
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
		Title:          "agents update",
		Transform:      transform,
	})
}

func handleAgentsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agents.List(ctx, params, options...)
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
		Title:          "agents list",
		Transform:      transform,
	})
}

func handleAgentsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Agents.Delete(ctx, cmd.Value("agent-id").(string), options...)
}

func handleAgentsGet(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agents.Get(ctx, cmd.Value("agent-id").(string), options...)
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
		Title:          "agents get",
		Transform:      transform,
	})
}

func handleAgentsRun(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agents.Run(ctx, params, options...)
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
		Title:          "agents run",
		Transform:      transform,
	})
}
