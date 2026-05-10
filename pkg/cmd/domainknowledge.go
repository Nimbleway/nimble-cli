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

var domainKnowledgeGetDriver = cli.Command{
	Name:    "get-driver",
	Usage:   "Resolves the suggested driver for a given URL or agent name. Exactly one of\n`url` or `agent` must be provided.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent",
			Usage:     "Agent name to resolve driver for (e.g. nimble-ecommerce).",
			QueryPath: "agent",
		},
		&requestflag.Flag[string]{
			Name:      "url",
			Usage:     "Target domain to resolve driver for (e.g. amazon.com).",
			QueryPath: "url",
		},
	},
	Action:          handleDomainKnowledgeGetDriver,
	HideHelpCommand: true,
}

func handleDomainKnowledgeGetDriver(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.DomainKnowledgeGetDriverParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.DomainKnowledge.GetDriver(ctx, params, options...)
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
		Title:          "domain-knowledge get-driver",
		Transform:      transform,
	})
}
