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
	Usage:   "List Templates",
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
			Usage:     "Filter public templates by attribution",
			QueryPath: "managed_by",
		},
		&requestflag.Flag[int64]{
			Name:      "offset",
			Usage:     "Pagination offset",
			Default:   0,
			QueryPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:      "privacy",
			Usage:     "Filter by privacy level",
			Default:   "public",
			QueryPath: "privacy",
		},
	},
	Action:          handleAgentsList,
	HideHelpCommand: true,
}

var agentsAsync = cli.Command{
	Name:    "async",
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
	Action:          handleAgentsAsync,
	HideHelpCommand: true,
}

var agentsGet = cli.Command{
	Name:    "get",
	Usage:   "Get Template",
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
		&requestflag.Flag[bool]{
			Name:     "localization",
			Default:  false,
			BodyPath: "localization",
		},
	},
	Action:          handleAgentsRun,
	HideHelpCommand: true,
}

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

func handleAgentsAsync(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.AgentAsyncParams{}

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
	_, err = client.Agents.Async(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents async", obj, format, transform)
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
