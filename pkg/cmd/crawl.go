// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/nimbleway-cli/internal/apiquery"
	"github.com/stainless-sdks/nimbleway-cli/internal/requestflag"
	"github.com/stainless-sdks/nimbleway-go"
	"github.com/stainless-sdks/nimbleway-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var crawlList = cli.Command{
	Name:    "list",
	Usage:   "Get crawl data by filters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter crawls by their status.",
			Required:  true,
			QueryPath: "status",
		},
		&requestflag.Flag[any]{
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
	},
	Action:          handleCrawlList,
	HideHelpCommand: true,
}

var crawlRoot = requestflag.WithInnerFlags(cli.Command{
	Name:    "root",
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
	Action:          handleCrawlRoot,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"extract-options": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.debug-options",
			Usage:      "Debug and troubleshooting options for the request",
			InnerField: "debug_options",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.url",
			Usage:      "Target URL to scrape",
			InnerField: "url",
		},
		&requestflag.InnerFlag[any]{
			Name:       "extract-options.browser",
			Usage:      "Browser type to emulate",
			InnerField: "browser",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.city",
			Usage:      "City for geolocation",
			InnerField: "city",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "extract-options.client-timeout",
			Usage:      "Client-side timeout in milliseconds",
			InnerField: "client_timeout",
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
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.disable-ip-check",
			Usage:      "Whether to disable IP address validation",
			InnerField: "disable_ip_check",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.driver",
			Usage:      "Browser driver to use",
			InnerField: "driver",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.dynamic-parser",
			Usage:      "Custom parser configuration as a key-value map",
			InnerField: "dynamic_parser",
		},
		&requestflag.InnerFlag[[]int64]{
			Name:       "extract-options.expected-status-codes",
			Usage:      "Expected HTTP status codes for successful requests",
			InnerField: "expected_status_codes",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.export-userbrowser",
			Usage:      "Whether to export the userbrowser session",
			InnerField: "export_userbrowser",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.format",
			Usage:      "Response format",
			InnerField: "format",
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
			Name:       "extract-options.ip6",
			Usage:      "Whether to use IPv6 for the request",
			InnerField: "ip6",
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
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.markdown",
			Usage:      "Whether to return response in Markdown format",
			InnerField: "markdown",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.metadata",
			Usage:      "Structured metadata about the request execution context",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.method",
			Usage:      "HTTP method for the request",
			InnerField: "method",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.native-mode",
			Usage:      "Native execution mode",
			InnerField: "native_mode",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "extract-options.network-capture",
			Usage:      "Filters for capturing network traffic",
			InnerField: "network_capture",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.no-html",
			Usage:      "Whether to exclude HTML from the response",
			InnerField: "no_html",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.no-userbrowser",
			Usage:      "Whether to disable browser-based rendering",
			InnerField: "no_userbrowser",
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.parse-options",
			Usage:      "Configuration options for parsing behavior",
			InnerField: "parse_options",
		},
		&requestflag.InnerFlag[any]{
			Name:       "extract-options.parser",
			Usage:      "Custom parser configuration as a key-value map",
			InnerField: "parser",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.proxy-provider",
			Usage:      "Proxy provider to use for the request",
			InnerField: "proxy_provider",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.proxy-providers",
			Usage:      "Weighted distribution of proxy providers",
			InnerField: "proxy_providers",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.query-template",
			Usage:      "Query template configuration for structured data extraction",
			InnerField: "query_template",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.raw-headers",
			Usage:      "Whether to return raw HTTP headers in response",
			InnerField: "raw_headers",
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
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "extract-options.render-flow",
			Usage:      "Array of actions to perform during browser rendering",
			InnerField: "render_flow",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.render-options",
			InnerField: "render_options",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "extract-options.request-timeout",
			Usage:      "Request timeout in milliseconds",
			InnerField: "request_timeout",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.return-response-headers-as-header",
			Usage:      "Whether to return response headers in HTTP headers",
			InnerField: "return_response_headers_as_header",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.save-userbrowser",
			Usage:      "Whether to save the userbrowser session for reuse",
			InnerField: "save_userbrowser",
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
		&requestflag.InnerFlag[bool]{
			Name:       "extract-options.skip-ubct",
			Usage:      "Whether to skip userbrowser creation template processing",
			InnerField: "skip_ubct",
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.template",
			Usage:      "Userbrowser creation template configuration",
			InnerField: "template",
		},
		&requestflag.InnerFlag[string]{
			Name:       "extract-options.type",
			Usage:      "Type of query or scraping template",
			InnerField: "type",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "extract-options.userbrowser-creation-template-rendered",
			Usage:      "Pre-rendered userbrowser creation template configuration",
			InnerField: "userbrowser_creation_template_rendered",
		},
	},
})

var crawlStatus = cli.Command{
	Name:    "status",
	Usage:   "Get crawl data",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The unique identifier of the crawl task.",
			Required: true,
		},
	},
	Action:          handleCrawlStatus,
	HideHelpCommand: true,
}

var crawlTerminate = cli.Command{
	Name:    "terminate",
	Usage:   "Cancel crawl task",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "The unique identifier of the crawl task.",
			Required: true,
		},
	},
	Action:          handleCrawlTerminate,
	HideHelpCommand: true,
}

func handleCrawlList(ctx context.Context, cmd *cli.Command) error {
	client := nimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := nimblego.CrawlListParams{}

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
	_, err = client.Crawl.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "crawl list", obj, format, transform)
}

func handleCrawlRoot(ctx context.Context, cmd *cli.Command) error {
	client := nimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := nimblego.CrawlRootParams{}

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
	_, err = client.Crawl.Root(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "crawl root", obj, format, transform)
}

func handleCrawlStatus(ctx context.Context, cmd *cli.Command) error {
	client := nimblego.NewClient(getDefaultRequestOptions(cmd)...)
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "crawl status", obj, format, transform)
}

func handleCrawlTerminate(ctx context.Context, cmd *cli.Command) error {
	client := nimblego.NewClient(getDefaultRequestOptions(cmd)...)
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "crawl terminate", obj, format, transform)
}
