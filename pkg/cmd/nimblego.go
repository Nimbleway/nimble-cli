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

var extract = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract",
	Usage:   "Webit v2 Realtime extract Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "debug-options",
			Usage:    "Debug and troubleshooting options for the request",
			Required: true,
			BodyPath: "debug_options",
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "Target URL to scrape",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[any]{
			Name:     "browser",
			Usage:    "Browser type to emulate",
			BodyPath: "browser",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "City for geolocation",
			BodyPath: "city",
		},
		&requestflag.Flag[float64]{
			Name:     "client-timeout",
			Usage:    "Client-side timeout in milliseconds",
			BodyPath: "client_timeout",
		},
		&requestflag.Flag[bool]{
			Name:     "consent-header",
			Usage:    "Whether to automatically handle cookie consent headers",
			BodyPath: "consent_header",
		},
		&requestflag.Flag[any]{
			Name:     "cookies",
			Usage:    "Browser cookies as array of cookie objects",
			BodyPath: "cookies",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Country code for geolocation and proxy selection",
			Default:  "ALL",
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "device",
			Usage:    "Device type for browser emulation",
			Default:  "desktop",
			BodyPath: "device",
		},
		&requestflag.Flag[bool]{
			Name:     "disable-ip-check",
			Usage:    "Whether to disable IP address validation",
			BodyPath: "disable_ip_check",
		},
		&requestflag.Flag[string]{
			Name:     "driver",
			Usage:    "Browser driver to use",
			BodyPath: "driver",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "dynamic-parser",
			Usage:    "Custom parser configuration as a key-value map",
			BodyPath: "dynamic_parser",
		},
		&requestflag.Flag[[]int64]{
			Name:     "expected-status-code",
			Usage:    "Expected HTTP status codes for successful requests",
			BodyPath: "expected_status_codes",
		},
		&requestflag.Flag[bool]{
			Name:     "export-userbrowser",
			Usage:    "Whether to export the userbrowser session",
			Default:  false,
			BodyPath: "export_userbrowser",
		},
		&requestflag.Flag[string]{
			Name:     "format",
			Usage:    "Response format",
			Default:  "json",
			BodyPath: "format",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "headers",
			Usage:    "Custom HTTP headers to include in the request",
			Default:  map[string]any{},
			BodyPath: "headers",
		},
		&requestflag.Flag[bool]{
			Name:     "http2",
			Usage:    "Whether to use HTTP/2 protocol",
			Default:  false,
			BodyPath: "http2",
		},
		&requestflag.Flag[bool]{
			Name:     "ip6",
			Usage:    "Whether to use IPv6 for the request",
			BodyPath: "ip6",
		},
		&requestflag.Flag[bool]{
			Name:     "is-xhr",
			Usage:    "Whether to emulate XMLHttpRequest behavior",
			Default:  false,
			BodyPath: "is_xhr",
		},
		&requestflag.Flag[string]{
			Name:     "locale",
			Usage:    "Locale for browser language and region settings",
			Default:  "en",
			BodyPath: "locale",
		},
		&requestflag.Flag[bool]{
			Name:     "markdown",
			Usage:    "Whether to return response in Markdown format",
			BodyPath: "markdown",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Structured metadata about the request execution context",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "method",
			Usage:    "HTTP method for the request",
			Default:  "GET",
			BodyPath: "method",
		},
		&requestflag.Flag[string]{
			Name:     "native-mode",
			Usage:    "Native execution mode",
			BodyPath: "native_mode",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "network-capture",
			Usage:    "Filters for capturing network traffic",
			BodyPath: "network_capture",
		},
		&requestflag.Flag[bool]{
			Name:     "no-html",
			Usage:    "Whether to exclude HTML from the response",
			Default:  false,
			BodyPath: "no_html",
		},
		&requestflag.Flag[bool]{
			Name:     "no-userbrowser",
			Usage:    "Whether to disable browser-based rendering",
			BodyPath: "no_userbrowser",
		},
		&requestflag.Flag[string]{
			Name:     "os",
			Usage:    "Operating system to emulate",
			BodyPath: "os",
		},
		&requestflag.Flag[bool]{
			Name:     "parse",
			Usage:    "Whether to parse the response content",
			Default:  false,
			BodyPath: "parse",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "parse-options",
			Usage:    "Configuration options for parsing behavior",
			BodyPath: "parse_options",
		},
		&requestflag.Flag[any]{
			Name:     "parser",
			Usage:    "Custom parser configuration as a key-value map",
			BodyPath: "parser",
		},
		&requestflag.Flag[string]{
			Name:     "proxy-provider",
			Usage:    "Proxy provider to use for the request",
			Default:  "proxit",
			BodyPath: "proxy_provider",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "proxy-providers",
			Usage:    "Weighted distribution of proxy providers",
			BodyPath: "proxy_providers",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "query-template",
			Usage:    "Query template configuration for structured data extraction",
			BodyPath: "query_template",
		},
		&requestflag.Flag[bool]{
			Name:     "raw-headers",
			Usage:    "Whether to return raw HTTP headers in response",
			Default:  false,
			BodyPath: "raw_headers",
		},
		&requestflag.Flag[string]{
			Name:     "referrer-type",
			Usage:    "Referrer policy for the request",
			BodyPath: "referrer_type",
		},
		&requestflag.Flag[bool]{
			Name:     "render",
			Usage:    "Whether to render JavaScript content using a browser",
			BodyPath: "render",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "render-flow",
			Usage:    "Array of actions to perform during browser rendering",
			BodyPath: "render_flow",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "render-options",
			BodyPath: "render_options",
		},
		&requestflag.Flag[float64]{
			Name:     "request-timeout",
			Usage:    "Request timeout in milliseconds",
			BodyPath: "request_timeout",
		},
		&requestflag.Flag[bool]{
			Name:     "return-response-headers-as-header",
			Usage:    "Whether to return response headers in HTTP headers",
			BodyPath: "return_response_headers_as_header",
		},
		&requestflag.Flag[bool]{
			Name:     "save-userbrowser",
			Usage:    "Whether to save the userbrowser session for reuse",
			Default:  false,
			BodyPath: "save_userbrowser",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "session",
			BodyPath: "session",
		},
		&requestflag.Flag[any]{
			Name:     "skill",
			Usage:    "Skills or capabilities required for the request",
			BodyPath: "skill",
		},
		&requestflag.Flag[bool]{
			Name:     "skip-ubct",
			Usage:    "Whether to skip userbrowser creation template processing",
			BodyPath: "skip_ubct",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "US state for geolocation (only valid when country is US)",
			BodyPath: "state",
		},
		&requestflag.Flag[string]{
			Name:     "tag",
			Usage:    "User-defined tag for request identification",
			BodyPath: "tag",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "template",
			Usage:    "Userbrowser creation template configuration",
			BodyPath: "template",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of query or scraping template",
			Default:  "generic",
			BodyPath: "type",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "userbrowser-creation-template-rendered",
			Usage:    "Pre-rendered userbrowser creation template configuration",
			BodyPath: "userbrowser_creation_template_rendered",
		},
	},
	Action:          handleExtract,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"debug-options": {
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.collect-har",
			InnerField: "collect_har",
		},
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.no-retry-mode",
			InnerField: "no_retry_mode",
		},
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.record-screen",
			InnerField: "record_screen",
		},
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.redact",
			InnerField: "redact",
		},
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.show-cursor",
			InnerField: "show_cursor",
		},
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.solve-captcha",
			InnerField: "solve_captcha",
		},
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.trace",
			InnerField: "trace",
		},
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.upload-engine-logs",
			InnerField: "upload_engine_logs",
		},
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.verbose",
			InnerField: "verbose",
		},
		&requestflag.InnerFlag[any]{
			Name:       "debug-options.with-proxy-usage",
			InnerField: "with_proxy_usage",
		},
	},
	"metadata": {
		&requestflag.InnerFlag[string]{
			Name:       "metadata.account-name",
			Usage:      "Account name associated with the request",
			InnerField: "account_name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "metadata.definition-id",
			Usage:      "Definition identifier",
			InnerField: "definition_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.definition-name",
			Usage:      "Name of the definition",
			InnerField: "definition_name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.endpoint",
			Usage:      "API endpoint being called",
			InnerField: "endpoint",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.execution-id",
			Usage:      "Unique identifier for this execution",
			InnerField: "execution_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.flowit-task-id",
			Usage:      "FlowIt task identifier",
			InnerField: "flowit_task_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.input-id",
			Usage:      "Input data identifier",
			InnerField: "input_id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "metadata.pipeline-execution-id",
			Usage:      "Identifier for the pipeline execution",
			InnerField: "pipeline_execution_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.query-template-id",
			Usage:      "Query template identifier",
			InnerField: "query_template_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.source",
			Usage:      "Source system or application making the request",
			InnerField: "source",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "metadata.template-id",
			Usage:      "Template identifier",
			InnerField: "template_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.template-name",
			Usage:      "Name of the template",
			InnerField: "template_name",
		},
	},
	"network-capture": {
		&requestflag.InnerFlag[string]{
			Name:       "network-capture.method",
			InnerField: "method",
		},
		&requestflag.InnerFlag[any]{
			Name:       "network-capture.resource-type",
			Usage:      "Resource type for network capture filtering",
			InnerField: "resource_type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "network-capture.status-code",
			InnerField: "status_code",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "network-capture.url",
			InnerField: "url",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "network-capture.validation",
			InnerField: "validation",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "network-capture.wait-for-requests-count",
			InnerField: "wait_for_requests_count",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "network-capture.wait-for-requests-count-timeout",
			InnerField: "wait_for_requests_count_timeout",
		},
	},
	"parse-options": {
		&requestflag.InnerFlag[bool]{
			Name:       "parse-options.merge-dynamic",
			Usage:      "Whether to merge dynamic parsing results with static results",
			InnerField: "merge_dynamic",
		},
	},
	"query-template": {
		&requestflag.InnerFlag[string]{
			Name:       "query-template.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "query-template.api-type",
			InnerField: "api_type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "query-template.pagination",
			InnerField: "pagination",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "query-template.params",
			InnerField: "params",
		},
	},
	"render-options": {
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.adblock",
			Usage:      "Whether to enable ad blocking",
			InnerField: "adblock",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "render-options.blocked-domains",
			Usage:      "Domains to block from loading",
			InnerField: "blocked_domains",
		},
		&requestflag.InnerFlag[any]{
			Name:       "render-options.browser-engine",
			Usage:      "Browser engine to use, or weighted distribution of engines",
			InnerField: "browser_engine",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.cache",
			Usage:      "Whether to enable browser caching",
			InnerField: "cache",
		},
		&requestflag.InnerFlag[string]{
			Name:       "render-options.connector-type",
			Usage:      "Type of browser connector to use",
			InnerField: "connector_type",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "render-options.disabled-resources",
			Usage:      "Types of resources to block from loading",
			InnerField: "disabled_resources",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.enable-2captcha",
			Usage:      "Whether to use 2Captcha service for solving captchas",
			InnerField: "enable_2captcha",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "render-options.extensions",
			Usage:      "Browser extensions to load",
			InnerField: "extensions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "render-options.fingerprint-id",
			Usage:      "Fingerprint identifier for browser customization",
			InnerField: "fingerprint_id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "render-options.hackium-configuration",
			Usage:      "Configuration for Hackium browser modifications",
			InnerField: "hackium_configuration",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.headless",
			Usage:      "Whether to run browser in headless mode",
			InnerField: "headless",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.include-iframes",
			Usage:      "Whether to include iframe content in the result",
			InnerField: "include_iframes",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.load-local-storage",
			Usage:      "Whether to load previously stored localStorage data",
			InnerField: "load_local_storage",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "render-options.local-storage-keys-to-load",
			Usage:      "Specific localStorage keys to load",
			InnerField: "local_storage_keys_to_load",
		},
		&requestflag.InnerFlag[string]{
			Name:       "render-options.mouse-strategy",
			Usage:      "Strategy for simulating mouse movements",
			InnerField: "mouse_strategy",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.no-accept-encoding",
			Usage:      "Disable content encoding to avoid cached responses",
			InnerField: "no_accept_encoding",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.override-permissions",
			Usage:      "Whether to override default browser permissions",
			InnerField: "override_permissions",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.random-header-order",
			Usage:      "Whether to randomize HTTP header order",
			InnerField: "random_header_order",
		},
		&requestflag.InnerFlag[string]{
			Name:       "render-options.render-type",
			Usage:      "Type of render completion to wait for",
			InnerField: "render_type",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.store-local-storage",
			Usage:      "Whether to store localStorage data for future sessions",
			InnerField: "store_local_storage",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "render-options.timeout",
			Usage:      "Maximum time in milliseconds to wait for page render",
			InnerField: "timeout",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "render-options.typing-interval",
			Usage:      "Interval in milliseconds between key presses",
			InnerField: "typing_interval",
		},
		&requestflag.InnerFlag[string]{
			Name:       "render-options.typing-strategy",
			Usage:      "Strategy for simulating keyboard typing",
			InnerField: "typing_strategy",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.userbrowser",
			Usage:      "Whether to use a persistent browser session",
			InnerField: "userbrowser",
		},
		&requestflag.InnerFlag[string]{
			Name:       "render-options.wait-until",
			Usage:      "Browser event to wait for before considering page loaded",
			InnerField: "wait_until",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "render-options.with-performance-metrics",
			Usage:      "Whether to collect performance metrics during rendering",
			InnerField: "with_performance_metrics",
		},
	},
	"session": {
		&requestflag.InnerFlag[string]{
			Name:       "session.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "session.prefetch-userbrowser",
			InnerField: "prefetch_userbrowser",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "session.retry",
			InnerField: "retry",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "session.timeout",
			InnerField: "timeout",
		},
	},
	"template": {
		&requestflag.InnerFlag[string]{
			Name:       "template.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "template.params",
			InnerField: "params",
		},
	},
	"userbrowser-creation-template-rendered": {
		&requestflag.InnerFlag[string]{
			Name:       "userbrowser-creation-template-rendered.id",
			InnerField: "id",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "userbrowser-creation-template-rendered.allowed-parameter-names",
			InnerField: "allowed_parameter_names",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "userbrowser-creation-template-rendered.render-flow-rendered",
			InnerField: "render_flow_rendered",
		},
	},
})

var extractTemplate = cli.Command{
	Name:    "extract-template",
	Usage:   "Execute WSA Template Realtime Endpoint",
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
	},
	Action:          handleExtractTemplate,
	HideHelpCommand: true,
}

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
			Usage:    "Maximum number of subagents to execute in parallel for WSA focus modes (shopping, social, geo). Ignored for traditional SERP focus modes. Default: 3, Range: 1-5.",
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
		&requestflag.Flag[string]{
			Name:     "topic",
			Usage:    "Search focus/specialization (general, news, or location)",
			Default:  "general",
			BodyPath: "topic",
		},
	},
	Action:          handleSearch,
	HideHelpCommand: true,
}

func handleExtract(ctx context.Context, cmd *cli.Command) error {
	client := nimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := nimblego.ExtractParams{}

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
	_, err = client.Extract(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "extract", obj, format, transform)
}

func handleExtractTemplate(ctx context.Context, cmd *cli.Command) error {
	client := nimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := nimblego.ExtractTemplateParams{}

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
	_, err = client.ExtractTemplate(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "extract-template", obj, format, transform)
}

func handleMap(ctx context.Context, cmd *cli.Command) error {
	client := nimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := nimblego.MapParams{}

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
	client := nimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := nimblego.SearchParams{}

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
