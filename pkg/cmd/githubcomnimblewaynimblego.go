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

var map_ = cli.Command{
	Name:    "map",
	Usage:   "Create map task",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "Url to map.",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Country code for geolocation and proxy selection",
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "domain-filter",
			Usage:    "Includes subdomains of the main domain in the mapping process.",
			Default:  "all",
			BodyPath: "domain_filter",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Maximum number of links to return.",
			Default:  5000,
			BodyPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			Usage:    "Locale for browser language and region settings",
			BodyPath: "locale",
		},
		&requestflag.Flag[string]{
			Name:     "sitemap",
			Usage:    "Sitemap and other methods will be used together to find URLs.",
			Default:  "include",
			BodyPath: "sitemap",
		},
	},
	Action:          handleMap,
	HideHelpCommand: true,
}

var search = cli.Command{
	Name:    "search",
	Usage:   "Search",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "query",
			Usage:    "Search query string",
			Required: true,
			BodyPath: "query",
		},
		&requestflag.Flag[any]{
			Name:     "content-type",
			Usage:    "Filter by content type (only supported with focus=general). Supports semantic groups ('documents', 'spreadsheets', 'presentations') and specific formats ('pdf', 'docx', 'xlsx', etc.)",
			BodyPath: "content_type",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Default:  "US",
			BodyPath: "country",
		},
		&requestflag.Flag[bool]{
			Name:     "deep-search",
			Usage:    "If True, fetches and extracts full page content for each search result. If False, returns only metadata (title, snippet, URL)",
			Default:  true,
			BodyPath: "deep_search",
		},
		&requestflag.Flag[any]{
			Name:     "end-date",
			Usage:    "Filter results before this date (format: YYYY-MM-DD or YYYY)",
			BodyPath: "end_date",
		},
		&requestflag.Flag[any]{
			Name:     "exclude-domain",
			Usage:    "List of domains to exclude from search results. Maximum 50 domains.",
			BodyPath: "exclude_domains",
		},
		&requestflag.Flag[bool]{
			Name:     "include-answer",
			Usage:    "Generate LLM answer summary based on search result snippets (works with both deep_search=True and False)",
			Default:  false,
			BodyPath: "include_answer",
		},
		&requestflag.Flag[any]{
			Name:     "include-domain",
			Usage:    "List of domains to include in search results. Maximum 50 domains.",
			BodyPath: "include_domains",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			Default:  "en",
			BodyPath: "locale",
		},
		&requestflag.Flag[int64]{
			Name:     "max-subagents",
			Usage:    "Maximum number of subagents to execute in parallel for WSA focus modes (shopping, social, geo). Ignored for traditional SERP focus modes. Default: 3, Range: 1-10.",
			Default:  3,
			BodyPath: "max_subagents",
		},
		&requestflag.Flag[int64]{
			Name:     "num-results",
			Usage:    "Maximum number of results to return (actual count may be less)",
			Default:  3,
			BodyPath: "num_results",
		},
		&requestflag.Flag[string]{
			Name:     "parsing-type",
			Usage:    "Output format: plain_text, markdown, or simplified_html",
			Default:  "markdown",
			BodyPath: "parsing_type",
		},
		&requestflag.Flag[any]{
			Name:     "search-engine",
			Usage:    "Enum representing the search engines supported by Nimble\n⚠️ DEPRECATED: This parameter is ignored. Use 'focus' parameter instead.",
			BodyPath: "search_engine",
		},
		&requestflag.Flag[any]{
			Name:     "start-date",
			Usage:    "Filter results after this date (format: YYYY-MM-DD or YYYY)",
			BodyPath: "start_date",
		},
		&requestflag.Flag[any]{
			Name:     "time-range",
			Usage:    "Time range filters passed to Webit SERP API as 'time' parameter.",
			BodyPath: "time_range",
		},
		&requestflag.Flag[any]{
			Name:     "topic",
			Usage:    "Search focus/specialization. Can be a single focus mode (e.g., 'shopping', 'social') or a list of explicit subagent names (e.g., ['amazon_serp', 'target_serp'])",
			Default:  "general",
			BodyPath: "topic",
		},
	},
	Action:          handleSearch,
	HideHelpCommand: true,
}

func handleMap(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.MapParams{}

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
	_, err = client.Map(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "map", obj, format, transform)
}

func handleSearch(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.SearchParams{}

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
	_, err = client.Search(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search", obj, format, transform)
}
