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

var crawl = requestflag.WithInnerFlags(cli.Command{
	Name:    "crawl",
	Usage:   "Create crawl task",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "Url to crawl.",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[bool]{
			Name:     "allow-external-links",
			Usage:    "Allows the crawler to follow links to external websites.",
			Default:  false,
			BodyPath: "allow_external_links",
		},
		&requestflag.Flag[bool]{
			Name:     "allow-subdomains",
			Usage:    "Allows the crawler to follow links to subdomains of the main domain.",
			Default:  false,
			BodyPath: "allow_subdomains",
		},
		&requestflag.Flag[any]{
			Name:     "callback",
			Usage:    "Webhook configuration for receiving crawl results.",
			BodyPath: "callback",
		},
		&requestflag.Flag[bool]{
			Name:     "crawl-entire-domain",
			Usage:    "Allows the crawler to follow internal links to sibling or parent URLs, not just child paths.",
			Default:  false,
			BodyPath: "crawl_entire_domain",
		},
		&requestflag.Flag[[]string]{
			Name:     "exclude-path",
			Usage:    "URL pathname regex patterns that exclude matching URLs from the crawl.",
			BodyPath: "exclude_paths",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "extract-options",
			BodyPath: "extract_options",
		},
		&requestflag.Flag[bool]{
			Name:     "ignore-query-parameters",
			Usage:    "Do not re-scrape the same path with different (or none) query parameters.",
			Default:  false,
			BodyPath: "ignore_query_parameters",
		},
		&requestflag.Flag[[]string]{
			Name:     "include-path",
			Usage:    "URL pathname regex patterns that include matching URLs in the crawl.",
			BodyPath: "include_paths",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Maximum number of pages to crawl.",
			Default:  5000,
			BodyPath: "limit",
		},
		&requestflag.Flag[int64]{
			Name:     "max-discovery-depth",
			Usage:    "Maximum depth to crawl based on discovery order.",
			Default:  5,
			BodyPath: "max_discovery_depth",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the crawl.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "sitemap",
			Usage:    "Sitemap and other methods will be used together to find URLs.",
			Default:  "include",
			BodyPath: "sitemap",
		},
	},
	Action:          handleCrawl,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"extract-options": {
		&requestflag.InnerFlag[any]{
			Name:       "extract-options.browser",
			Usage:      "Browser type to emulate",
			InnerField: "browser",
		},
		&requestflag.InnerFlag[[]any]{
			Name:       "extract-options.browser-actions",
			Usage:      "Array of browser automation actions to execute sequentially",
			InnerField: "browser_actions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.city",
			Usage:      "City for geolocation",
			InnerField: "city",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.consent-header",
			Usage:      "Whether to automatically handle cookie consent headers",
			InnerField: "consent_header",
		},
		&requestflag.InnerFlag[any]{
			Name:       "extract-options.cookies",
			Usage:      "Browser cookies as array of cookie objects",
			InnerField: "cookies",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.country",
			Usage:      "Country code for geolocation and proxy selection",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.device",
			Usage:      "Device type for browser emulation",
			InnerField: "device",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.driver",
			Usage:      "Browser driver to use",
			InnerField: "driver",
		},
		&requestflag.InnerFlag[[]int64]{
			Name:       "extract-options.expected-status-codes",
			Usage:      "Expected HTTP status codes for successful requests",
			InnerField: "expected_status_codes",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "extract-options.formats",
			Usage:      "List of acceptable response formats in order of preference",
			InnerField: "formats",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.headers",
			Usage:      "Custom HTTP headers to include in the request",
			InnerField: "headers",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.http2",
			Usage:      "Whether to use HTTP/2 protocol",
			InnerField: "http2",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.is-xhr",
			Usage:      "Whether to emulate XMLHttpRequest behavior",
			InnerField: "is_xhr",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.locale",
			Usage:      "Locale for browser language and region settings",
			InnerField: "locale",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.method",
			Usage:      "HTTP method for the request",
			InnerField: "method",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "extract-options.network-capture",
			Usage:      "Filters for capturing network traffic",
			InnerField: "network_capture",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.os",
			Usage:      "Operating system to emulate",
			InnerField: "os",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.parse",
			Usage:      "Whether to parse the response content",
			InnerField: "parse",
		},
		&requestflag.InnerFlag[any]{
			Name:       "extract-options.parser",
			Usage:      "Custom parser configuration as a key-value map",
			InnerField: "parser",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.referrer-type",
			Usage:      "Referrer policy for the request",
			InnerField: "referrer_type",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.render",
			Usage:      "Whether to render JavaScript content using a browser",
			InnerField: "render",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "extract-options.request-timeout",
			Usage:      "Request timeout in milliseconds",
			InnerField: "request_timeout",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.session",
			InnerField: "session",
		},
		&requestflag.InnerFlag[any]{
			Name:       "extract-options.skill",
			Usage:      "Skills or capabilities required for the request",
			InnerField: "skill",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.state",
			Usage:      "US state for geolocation (only valid when country is US)",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.tag",
			Usage:      "User-defined tag for request identification",
			InnerField: "tag",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.url",
			Usage:      "Target URL to scrape",
			InnerField: "url",
		},
	},
})

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

func handleCrawl(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.CrawlParams{}

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
	_, err = client.Crawl(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "crawl", obj, format, transform)
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
