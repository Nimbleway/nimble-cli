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

var fastSerpRun = cli.Command{
	Name:    "run",
	Usage:   "Fast SERP",
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
	Action:          handleFastSerpRun,
	HideHelpCommand: true,
}

func handleFastSerpRun(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.FastSerpRunParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.FastSerp.Run(ctx, params, options...)
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
		Title:          "fast-serp run",
		Transform:      transform,
	})
}
