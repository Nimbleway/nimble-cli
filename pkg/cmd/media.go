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

var mediaRun = requestflag.WithInnerFlags(cli.Command{
	Name:    "run",
	Usage:   "Download media from a URL. Waits for the result before responding.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			BodyPath: "country",
		},
		&requestflag.Flag[[]string]{
			Name:     "expected-mime-type",
			BodyPath: "expected_mime_types",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			BodyPath: "locale",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "storage",
			BodyPath: "storage",
		},
	},
	Action:          handleMediaRun,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"storage": {
		&requestflag.InnerFlag[string]{
			Name:       "storage.url",
			InnerField: "url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "storage.object-name",
			InnerField: "object_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "storage.type",
			Usage:      `Allowed values: "s3", "gcs", "do".`,
			InnerField: "type",
		},
	},
})

var mediaRunAsync = requestflag.WithInnerFlags(cli.Command{
	Name:    "run-async",
	Usage:   "Download media from a URL asynchronously. Returns a task ID immediately.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "URL to call back when async operation completes",
			BodyPath: "callback_url",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			BodyPath: "country",
		},
		&requestflag.Flag[[]string]{
			Name:     "expected-mime-type",
			BodyPath: "expected_mime_types",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			BodyPath: "locale",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "storage",
			BodyPath: "storage",
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
	Action:          handleMediaRunAsync,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"storage": {
		&requestflag.InnerFlag[string]{
			Name:       "storage.url",
			InnerField: "url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "storage.object-name",
			InnerField: "object_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "storage.type",
			Usage:      `Allowed values: "s3", "gcs", "do".`,
			InnerField: "type",
		},
	},
})

func handleMediaRun(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.MediaRunParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Media.Run(ctx, params, options...)
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
		Title:          "media run",
		Transform:      transform,
	})
}

func handleMediaRunAsync(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.MediaRunAsyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Media.RunAsync(ctx, params, options...)
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
		Title:          "media run-async",
		Transform:      transform,
	})
}
