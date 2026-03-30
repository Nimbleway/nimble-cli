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

var agentsGenerationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Create Agent Generation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-name",
			Required: true,
			BodyPath: "agent_name",
		},
		&requestflag.Flag[string]{
			Name:     "prompt",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[any]{
			Name:     "input-schema",
			BodyPath: "input_schema",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "output-schema",
			BodyPath: "output_schema",
		},
		&requestflag.Flag[string]{
			Name:     "from-agent",
			Required: true,
			BodyPath: "from_agent",
		},
	},
	Action:          handleAgentsGenerationsCreate,
	HideHelpCommand: true,
}

var agentsGenerationsGet = cli.Command{
	Name:    "get",
	Usage:   "Get Agent Generation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "generation-id",
			Required: true,
		},
	},
	Action:          handleAgentsGenerationsGet,
	HideHelpCommand: true,
}

func handleAgentsGenerationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.AgentGenerationNewParams{}

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
	_, err = client.Agents.Generations.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents:generations create", obj, format, transform)
}

func handleAgentsGenerationsGet(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Agents.Generations.Get(ctx, cmd.Value("generation-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents:generations get", obj, format, transform)
}
