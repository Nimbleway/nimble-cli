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
	Usage:   "Create a new workspace-scoped Web Search Agent. Pass `template` to clone from a\nnamed template.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:     "agent-name",
			BodyPath: "agent_name",
		},
		&requestflag.Flag[*string]{
			Name:     "description",
			BodyPath: "description",
		},
		&requestflag.Flag[*string]{
			Name:     "display-name",
			BodyPath: "display_name",
		},
		&requestflag.Flag[*string]{
			Name:     "domain-expertise",
			BodyPath: "domain_expertise",
		},
		&requestflag.Flag[string]{
			Name:     "effort",
			Default:  "research",
			BodyPath: "effort",
		},
		&requestflag.Flag[[]string]{
			Name:     "goal",
			BodyPath: "goals",
		},
		&requestflag.Flag[*string]{
			Name:     "icon",
			BodyPath: "icon",
		},
		&requestflag.Flag[bool]{
			Name:     "is-active",
			Default:  true,
			BodyPath: "is_active",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "output-schema",
			BodyPath: "output_schema",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sources",
			BodyPath: "sources",
		},
		&requestflag.Flag[[]string]{
			Name:     "suggested-question",
			BodyPath: "suggested_questions",
		},
		&requestflag.Flag[*string]{
			Name:     "template",
			Usage:    "Template name to materialise this instance from. When set, scalar fields and child rows are copied from the template.",
			BodyPath: "template",
		},
		&requestflag.Flag[*string]{
			Name:     "use-case",
			Usage:    `Allowed values: "research", "enrichment", "dataset_building".`,
			BodyPath: "use_case",
		},
		&requestflag.Flag[*string]{
			Name:     "workspace-id",
			BodyPath: "workspace_id",
		},
	},
	Action:          handleTaskAgentCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"sources": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sources.allow",
			InnerField: "allow",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sources.avoid",
			InnerField: "avoid",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sources.block",
			InnerField: "block",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sources.prioritize",
			InnerField: "prioritize",
		},
	},
})

var taskAgentUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Apply a JSON Patch document (`application/json-patch+json`) to an agent you own.\nEach operation must be a `replace` with path `/field_name`.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "body",
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
			Usage:      `Allowed values: "replace".`,
			InnerField: "op",
		},
		&requestflag.InnerFlag[string]{
			Name:       "body.path",
			InnerField: "path",
		},
		&requestflag.InnerFlag[any]{
			Name:       "body.value",
			InnerField: "value",
		},
	},
})

var taskAgentList = cli.Command{
	Name:    "list",
	Usage:   "List active Web Search Agents visible to the caller. Includes agents scoped to\nthe caller's workspace.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "effort",
			QueryPath: "effort",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Default:   20,
			QueryPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Default:   0,
			QueryPath: "offset",
		},
		&requestflag.Flag[*string]{
			Name:      "use-case",
			QueryPath: "use_case",
		},
	},
	Action:          handleTaskAgentList,
	HideHelpCommand: true,
}

var taskAgentDeactivate = cli.Command{
	Name:    "deactivate",
	Usage:   "Deactivate an agent you own. The agent is marked inactive but not deleted.",
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
	Usage:   "Fetch a single Web Search Agent by id.",
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
	Usage:   "Create and enqueue a research run for a Web Search Agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agent_id",
		},
		&requestflag.Flag[string]{
			Name:     "input",
			Required: true,
			BodyPath: "input",
		},
		&requestflag.Flag[bool]{
			Name:     "enable-events",
			Default:  false,
			BodyPath: "enable_events",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "output-schema",
			BodyPath: "output_schema",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "sources",
			BodyPath: "sources",
		},
	},
	Action:          handleTaskAgentRun,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"sources": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sources.allow",
			InnerField: "allow",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sources.avoid",
			InnerField: "avoid",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "sources.block",
			InnerField: "block",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "sources.prioritize",
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
