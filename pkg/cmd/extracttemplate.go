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

var extractTemplatesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Patch Extract Template Public V2",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "extract-template-name",
			Required:  true,
			PathParam: "extract_template_name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "body",
			Usage:    "A JSON Patch document per RFC 6902 — a JSON array of patch operations.",
			Required: true,
			BodyRoot: true,
		},
	},
	Action:          handleExtractTemplatesUpdate,
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

var extractTemplatesList = cli.Command{
	Name:    "list",
	Usage:   "List Extract Templates Public V2",
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
	},
	Action:          handleExtractTemplatesList,
	HideHelpCommand: true,
}

var extractTemplatesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete Extract Template Public V2",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "extract-template-name",
			Required:  true,
			PathParam: "extract_template_name",
		},
	},
	Action:          handleExtractTemplatesDelete,
	HideHelpCommand: true,
}

var extractTemplatesAsync = cli.Command{
	Name:    "async",
	Usage:   "Execute Extraction Template Async Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Required: true,
			BodyPath: "params",
		},
		&requestflag.Flag[string]{
			Name:     "template",
			Required: true,
			BodyPath: "template",
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
	Action:          handleExtractTemplatesAsync,
	HideHelpCommand: true,
}

var extractTemplatesBatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "batch",
	Usage:   "Execute Extraction Template Batch Endpoint",
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
	Action:          handleExtractTemplatesBatch,
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
			Name:       "shared-inputs.template",
			InnerField: "template",
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

var extractTemplatesGet = cli.Command{
	Name:    "get",
	Usage:   "Get Extract Template Public V2",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "extract-template-name",
			Required:  true,
			PathParam: "extract_template_name",
		},
	},
	Action:          handleExtractTemplatesGet,
	HideHelpCommand: true,
}

var extractTemplatesRun = cli.Command{
	Name:    "run",
	Usage:   "Execute Extraction Template Realtime Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Required: true,
			BodyPath: "params",
		},
		&requestflag.Flag[string]{
			Name:     "template",
			Required: true,
			BodyPath: "template",
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
	Action:          handleExtractTemplatesRun,
	HideHelpCommand: true,
}

func handleExtractTemplatesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("extract-template-name") && len(unusedArgs) > 0 {
		cmd.Set("extract-template-name", unusedArgs[0])
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

	params := githubcomnimblewaynimblego.ExtractTemplateUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.Templates.Update(
		ctx,
		cmd.Value("extract-template-name").(string),
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
		Title:          "extract:templates update",
		Transform:      transform,
	})
}

func handleExtractTemplatesList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.ExtractTemplateListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.Templates.List(ctx, params, options...)
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
		Title:          "extract:templates list",
		Transform:      transform,
	})
}

func handleExtractTemplatesDelete(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("extract-template-name") && len(unusedArgs) > 0 {
		cmd.Set("extract-template-name", unusedArgs[0])
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

	return client.Extract.Templates.Delete(ctx, cmd.Value("extract-template-name").(string), options...)
}

func handleExtractTemplatesAsync(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.ExtractTemplateAsyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.Templates.Async(ctx, params, options...)
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
		Title:          "extract:templates async",
		Transform:      transform,
	})
}

func handleExtractTemplatesBatch(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.ExtractTemplateBatchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.Templates.Batch(ctx, params, options...)
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
		Title:          "extract:templates batch",
		Transform:      transform,
	})
}

func handleExtractTemplatesGet(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("extract-template-name") && len(unusedArgs) > 0 {
		cmd.Set("extract-template-name", unusedArgs[0])
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
	_, err = client.Extract.Templates.Get(ctx, cmd.Value("extract-template-name").(string), options...)
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
		Title:          "extract:templates get",
		Transform:      transform,
	})
}

func handleExtractTemplatesRun(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.ExtractTemplateRunParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract.Templates.Run(ctx, params, options...)
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
		Title:          "extract:templates run",
		Transform:      transform,
	})
}
