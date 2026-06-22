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

var crawlList = cli.Command{
	Name:    "list",
	Usage:   "Crawl by Filter",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*string]{
			Name:      "cursor",
			Usage:     "Cursor for pagination.",
			Default:   nil,
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Number of crawls to return per page.",
			Default:   100,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter crawls by their status.",
			Default:   "all",
			QueryPath: "status",
		},
	},
	Action:          handleCrawlList,
	HideHelpCommand: true,
}

var crawlRun = requestflag.WithInnerFlags(cli.Command{
	Name:    "run",
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
	Action:          handleCrawlRun,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"extract-options": {
		&requestflag.InnerFlag[any]{
			Name:       "extract-options.body",
			Usage:      "Request body for POST, PUT, PATCH methods",
			InnerField: "body",
		},
		&requestflag.InnerFlag[any]{
			Name:       "extract-options.browser",
			Usage:      "Browser type to emulate",
			InnerField: "browser",
		},
		&requestflag.InnerFlag[[]map[string]any]{
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
			Name:       "extract-options.markdown-backend",
			Usage:      `Selects which markdown conversion strategy to use. "full_page" converts the entire HTML page. "main_content" uses Mozilla Readability to extract the main article content before converting.`,
			InnerField: "markdown_backend",
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
		&requestflag.InnerFlag[any]{
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

var crawlStatus = cli.Command{
	Name:    "status",
	Usage:   "Get crawl data",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The unique identifier of the crawl task.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleCrawlStatus,
	HideHelpCommand: true,
}

var crawlTerminate = cli.Command{
	Name:    "terminate",
	Usage:   "Cancel Crawl",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "The unique identifier of the crawl task.",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleCrawlTerminate,
	HideHelpCommand: true,
}

func handleCrawlList(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.CrawlListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Crawl.List(ctx, params, options...)
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
		Title:          "crawl list",
		Transform:      transform,
	})
}

func handleCrawlRun(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.CrawlRunParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Crawl.Run(ctx, params, options...)
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
		Title:          "crawl run",
		Transform:      transform,
	})
}

func handleCrawlStatus(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Crawl.Status(ctx, cmd.Value("id").(string), options...)
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
		Title:          "crawl status",
		Transform:      transform,
	})
}

func handleCrawlTerminate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Crawl.Terminate(ctx, cmd.Value("id").(string), options...)
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
		Title:          "crawl terminate",
		Transform:      transform,
	})
}
