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

var serpRun = cli.Command{
	Name:    "run",
	Usage:   "SERP",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "search-engine",
			Usage:    "The search engine to query.",
			Required: true,
			BodyPath: "search_engine",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "ISO Alpha-2 country code used to access the target search engine (e.g. US, DE, GB).",
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "device",
			Usage:    "Device type used for the search request.",
			BodyPath: "device",
		},
		&requestflag.Flag[string]{
			Name:     "domain",
			Usage:    `Top-level domain for the search engine (e.g. "com", "co.uk", "de").`,
			Default:  "com",
			BodyPath: "domain",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			Usage:    "Locale used for the search request.",
			BodyPath: "locale",
		},
		&requestflag.Flag[string]{
			Name:     "location",
			Usage:    "Geo-location for the search (canonical Google location name).",
			BodyPath: "location",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of results to return (1–100).",
			BodyPath: "num_results",
		},
		&requestflag.Flag[int64]{
			Name:     "page",
			Usage:    "The result page number for pagination.",
			BodyPath: "page",
		},
		&requestflag.Flag[bool]{
			Name:     "parse",
			Usage:    "When true, the SERP response is parsed into structured JSON.",
			Default:  true,
			BodyPath: "parse",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "The search keyword or phrase to query.",
			BodyPath: "query",
		},
		&requestflag.Flag[bool]{
			Name:     "render",
			Usage:    "Whether to render the page in a browser before extracting.",
			Default:  false,
			BodyPath: "render",
		},
		&requestflag.Flag[bool]{
			Name:     "show-hidden-results",
			Usage:    "When true, disables Google result filtering (filter=0) so omitted/duplicate and highly similar pages are also returned. Applies to Google search engines.",
			BodyPath: "show_hidden_results",
		},
	},
	Action:          handleSerpRun,
	HideHelpCommand: true,
}

var serpRunAsync = cli.Command{
	Name:    "run-async",
	Usage:   "SERP Async Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "search-engine",
			Usage:    "The search engine to query.",
			Required: true,
			BodyPath: "search_engine",
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "URL to call back when async operation completes",
			BodyPath: "callback_url",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "ISO Alpha-2 country code used to access the target search engine (e.g. US, DE, GB).",
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "device",
			Usage:    "Device type used for the search request.",
			BodyPath: "device",
		},
		&requestflag.Flag[string]{
			Name:     "domain",
			Usage:    `Top-level domain for the search engine (e.g. "com", "co.uk", "de").`,
			Default:  "com",
			BodyPath: "domain",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			Usage:    "Locale used for the search request.",
			BodyPath: "locale",
		},
		&requestflag.Flag[string]{
			Name:     "location",
			Usage:    "Geo-location for the search (canonical Google location name).",
			BodyPath: "location",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Number of results to return (1–100).",
			BodyPath: "num_results",
		},
		&requestflag.Flag[int64]{
			Name:     "page",
			Usage:    "The result page number for pagination.",
			BodyPath: "page",
		},
		&requestflag.Flag[bool]{
			Name:     "parse",
			Usage:    "When true, the SERP response is parsed into structured JSON.",
			Default:  true,
			BodyPath: "parse",
		},
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "The search keyword or phrase to query.",
			BodyPath: "query",
		},
		&requestflag.Flag[bool]{
			Name:     "render",
			Usage:    "Whether to render the page in a browser before extracting.",
			Default:  false,
			BodyPath: "render",
		},
		&requestflag.Flag[bool]{
			Name:     "show-hidden-results",
			Usage:    "When true, disables Google result filtering (filter=0) so omitted/duplicate and highly similar pages are also returned. Applies to Google search engines.",
			BodyPath: "show_hidden_results",
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
	Action:          handleSerpRunAsync,
	HideHelpCommand: true,
}

var serpRunBatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "run-batch",
	Usage:   "SERP Batch Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "input",
			Usage:    "Array of SERP requests. Each object can include search parameters and async/storage settings.",
			Required: true,
			BodyPath: "inputs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "shared-inputs",
			Usage:    "Shared parameters applied to the entire batch. Can include search parameters and async/storage settings.",
			BodyPath: "shared_inputs",
		},
	},
	Action:          handleSerpRunBatch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"input": {
		&requestflag.InnerFlag[string]{
			Name:       "input.callback-url",
			Usage:      "URL to call back when async operation completes",
			InnerField: "callback_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.country",
			Usage:      "ISO Alpha-2 country code used to access the target search engine (e.g. US, DE, GB).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.device",
			Usage:      "Device type used for the search request.",
			InnerField: "device",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.domain",
			Usage:      `Top-level domain for the search engine (e.g. "com", "co.uk", "de").`,
			InnerField: "domain",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.locale",
			Usage:      "Locale used for the search request.",
			InnerField: "locale",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.location",
			Usage:      "Geo-location for the search (canonical Google location name).",
			InnerField: "location",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "input.num-results",
			Usage:      "Number of results to return (1–100).",
			InnerField: "num_results",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "input.page",
			Usage:      "The result page number for pagination.",
			InnerField: "page",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "input.parse",
			Usage:      "When true, the SERP response is parsed into structured JSON.",
			InnerField: "parse",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.query",
			Usage:      "The search keyword or phrase to query.",
			InnerField: "query",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "input.render",
			Usage:      "Whether to render the page in a browser before extracting.",
			InnerField: "render",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.search-engine",
			Usage:      "The search engine to query.",
			InnerField: "search_engine",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "input.show-hidden-results",
			Usage:      "When true, disables Google result filtering (filter=0) so omitted/duplicate and highly similar pages are also returned. Applies to Google search engines.",
			InnerField: "show_hidden_results",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "input.storage-compress",
			Usage:      "Whether to compress stored data",
			InnerField: "storage_compress",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.storage-object-name",
			Usage:      "Custom name for the stored object",
			InnerField: "storage_object_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.storage-type",
			Usage:      "Type of storage to use for results",
			InnerField: "storage_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.storage-url",
			Usage:      "URL for storage location",
			InnerField: "storage_url",
		},
	},
	"shared-inputs": {
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.callback-url",
			Usage:      "URL to call back when async operation completes",
			InnerField: "callback_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.country",
			Usage:      "ISO Alpha-2 country code used to access the target search engine (e.g. US, DE, GB).",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.device",
			Usage:      "Device type used for the search request.",
			InnerField: "device",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.domain",
			Usage:      `Top-level domain for the search engine (e.g. "com", "co.uk", "de").`,
			InnerField: "domain",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.locale",
			Usage:      "Locale used for the search request.",
			InnerField: "locale",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.location",
			Usage:      "Geo-location for the search (canonical Google location name).",
			InnerField: "location",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "shared-inputs.num-results",
			Usage:      "Number of results to return (1–100).",
			InnerField: "num_results",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "shared-inputs.page",
			Usage:      "The result page number for pagination.",
			InnerField: "page",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "shared-inputs.parse",
			Usage:      "When true, the SERP response is parsed into structured JSON.",
			InnerField: "parse",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.query",
			Usage:      "The search keyword or phrase to query.",
			InnerField: "query",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "shared-inputs.render",
			Usage:      "Whether to render the page in a browser before extracting.",
			InnerField: "render",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.search-engine",
			Usage:      "The search engine to query.",
			InnerField: "search_engine",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "shared-inputs.show-hidden-results",
			Usage:      "When true, disables Google result filtering (filter=0) so omitted/duplicate and highly similar pages are also returned. Applies to Google search engines.",
			InnerField: "show_hidden_results",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "shared-inputs.storage-compress",
			Usage:      "Whether to compress stored data",
			InnerField: "storage_compress",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.storage-object-name",
			Usage:      "Custom name for the stored object",
			InnerField: "storage_object_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.storage-type",
			Usage:      "Type of storage to use for results",
			InnerField: "storage_type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.storage-url",
			Usage:      "URL for storage location",
			InnerField: "storage_url",
		},
	},
})

func handleSerpRun(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.SerpRunParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Serp.Run(ctx, params, options...)
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
		Title:          "serp run",
		Transform:      transform,
	})
}

func handleSerpRunAsync(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.SerpRunAsyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Serp.RunAsync(ctx, params, options...)
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
		Title:          "serp run-async",
		Transform:      transform,
	})
}

func handleSerpRunBatch(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.SerpRunBatchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Serp.RunBatch(ctx, params, options...)
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
		Title:          "serp run-batch",
		Transform:      transform,
	})
}
