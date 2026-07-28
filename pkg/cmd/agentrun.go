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

var agentsRunsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Start an agent run. The run executes asynchronously: the response returns\nimmediately with status `queued`, then poll `GET .../runs/{run_id}` until\n`completed` and fetch the output from `GET .../runs/{run_id}/result` — or set\n`enable_events: true` and follow `GET .../runs/{run_id}/events` for live\nprogress.",
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
			Name:     "agent-name",
			Usage:    "Stable agent name. On this no-agent-id route, an unseen name creates a new agent; an existing name reuses it. Ignored on the /{agent_id}/runs route.",
			BodyPath: "agent_name",
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
		&requestflag.Flag[string]{
			Name:     "origin",
			Usage:    "Origin of public API runs. Public requests are always API-originated.",
			Default:  "api",
			BodyPath: "origin",
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
		&requestflag.Flag[*string]{
			Name:     "skill",
			Usage:    "Skill override for this run. One-time only, except when this run creates a new agent via agent_name, in which case it becomes the new agent's stored skill.",
			BodyPath: "skill",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sources",
			Usage:    "Source guidance overriding the agent default.",
			BodyPath: "sources",
		},
		&requestflag.Flag[*string]{
			Name:     "use-case",
			Usage:    "Only settable when this run creates a new agent (via agent_name, or when no agent is resolved), in which case it becomes the new agent's stored use_case. For a run against an existing agent, this must match the agent's own use_case — passing the same value is accepted as a no-op, a different value is rejected.",
			BodyPath: "use_case",
		},
	},
	Action:          handleAgentsRunsCreate,
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

var agentsRunsList = cli.Command{
	Name:    "list",
	Usage:   "List the runs of an agent, newest first, paginated with `offset`/`limit`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
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
	},
	Action:          handleAgentsRunsList,
	HideHelpCommand: true,
}

var agentsRunsGet = cli.Command{
	Name:    "get",
	Usage:   "Retrieve a run's current state. Poll this endpoint after creating a run: the run\nis finished once `status` is `completed`, `failed`, or `cancelled`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
	},
	Action:          handleAgentsRunsGet,
	HideHelpCommand: true,
}

var agentsRunsResult = cli.Command{
	Name:    "result",
	Usage:   "Fetch the output of a completed run. The `output` is `type: \"text\"` (a prose\nanswer) or `type: \"json\"` (structured data matching the output schema), plus\n`trust` metadata with per-claim citations for the answer.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
	},
	Action:          handleAgentsRunsResult,
	HideHelpCommand: true,
}

var agentsRunsStreamEvents = cli.Command{
	Name:    "stream-events",
	Usage:   "Stream a run's progress as\n[server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)\n(`text/event-stream`). Create the run with `enable_events: true` to have events\npublished. A keep-alive comment is sent every 15 seconds.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			Required:  true,
			PathParam: "run_id",
		},
	},
	Action:          handleAgentsRunsStreamEvents,
	HideHelpCommand: true,
}

func handleAgentsRunsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.AgentRunNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Runs.New(
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
		Title:          "agents:runs create",
		Transform:      transform,
	})
}

func handleAgentsRunsList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.AgentRunListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Runs.List(
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
		Title:          "agents:runs list",
		Transform:      transform,
	})
}

func handleAgentsRunsGet(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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

	params := githubcomnimblewaynimblego.AgentRunGetParams{
		AgentID: cmd.Value("agent-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Runs.Get(
		ctx,
		cmd.Value("run-id").(string),
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
		Title:          "agents:runs get",
		Transform:      transform,
	})
}

func handleAgentsRunsResult(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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

	params := githubcomnimblewaynimblego.AgentRunResultParams{
		AgentID: cmd.Value("agent-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Runs.Result(
		ctx,
		cmd.Value("run-id").(string),
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
		Title:          "agents:runs result",
		Transform:      transform,
	})
}

func handleAgentsRunsStreamEvents(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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

	params := githubcomnimblewaynimblego.AgentRunStreamEventsParams{
		AgentID: cmd.Value("agent-id").(string),
	}

	return client.Agents.Runs.StreamEvents(
		ctx,
		cmd.Value("run-id").(string),
		params,
		options...,
	)
}
