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

var extractTemplatesGenerationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create Extract Template Generation Public V2",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "Instructions for generating the extract template.",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "Example URL used to generate the extract template.",
			BodyPath: "url",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "input-schema",
			Usage:    "Optional JSON schema describing expected input parameters.",
			BodyPath: "input_schema",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Metadata to attach to the generated extract template.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[*string]{
			Name:     "name",
			Usage:    "Optional stable name for the generated extract template.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "output-schema",
			Usage:    "Optional JSON schema describing desired extracted output.",
			BodyPath: "output_schema",
		},
		&requestflag.Flag[string]{
			Name:     "from-extract-template",
			Usage:    "Name of the source extract template to refine.",
			BodyPath: "from_extract_template",
		},
	},
	Action:          handleExtractTemplatesGenerationsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"metadata": {
		&requestflag.InnerFlag[*string]{
			Name:       "metadata.description",
			Usage:      "Description for the generated template.",
			InnerField: "description",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "metadata.display-name",
			Usage:      "Human-friendly display name for the generated template.",
			InnerField: "display_name",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "metadata.tags",
			Usage:      "Tags to associate with the generated template.",
			InnerField: "tags",
		},
	},
})

var extractTemplatesGenerationsGet = cli.Command{
	Name:    "get",
	Usage:   "Get Extract Template Generation Public V2",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "generation-id",
			Required:  true,
			PathParam: "generation_id",
		},
	},
	Action:          handleExtractTemplatesGenerationsGet,
	HideHelpCommand: true,
}

func handleExtractTemplatesGenerationsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.ExtractTemplateGenerationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.Templates.Generations.New(ctx, params, options...)
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
		Title:          "extract:templates:generations create",
		Transform:      transform,
	})
}

func handleExtractTemplatesGenerationsGet(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Extract.Templates.Generations.Get(ctx, cmd.Value("generation-id").(string), options...)
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
		Title:          "extract:templates:generations get",
		Transform:      transform,
	})
}
