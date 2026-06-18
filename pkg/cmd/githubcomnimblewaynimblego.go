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

var extract = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract",
	Usage:   "Extract",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "Target URL to scrape",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[any]{
			Name:     "body",
			Usage:    "Request body for POST, PUT, PATCH methods",
			BodyPath: "body",
		},
		&requestflag.Flag[any]{
			Name:     "browser",
			Usage:    "Browser type to emulate",
			BodyPath: "browser",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "browser-action",
			Usage:    "Array of browser automation actions to execute sequentially",
			BodyPath: "browser_actions",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "City for geolocation",
			BodyPath: "city",
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
		&requestflag.Flag[string]{
			Name:     "driver",
			Usage:    "Browser driver to use",
			BodyPath: "driver",
		},
		&requestflag.Flag[[]int64]{
			Name:     "expected-status-code",
			Usage:    "Expected HTTP status codes for successful requests",
			BodyPath: "expected_status_codes",
		},
		&requestflag.Flag[[]string]{
			Name:     "format",
			Usage:    "List of acceptable response formats in order of preference",
			BodyPath: "formats",
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
		&requestflag.Flag[string]{
			Name:     "markdown-backend",
			Usage:    `Selects which markdown conversion strategy to use. "full_page" converts the entire HTML page. "main_content" uses Mozilla Readability to extract the main article content before converting.`,
			Default:  "full_page",
			BodyPath: "markdown_backend",
		},
		&requestflag.Flag[string]{
			Name:     "method",
			Usage:    "HTTP method for the request",
			Default:  "GET",
			BodyPath: "method",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "network-capture",
			Usage:    "Filters for capturing network traffic",
			BodyPath: "network_capture",
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
		&requestflag.Flag[any]{
			Name:     "parser",
			Usage:    "Custom parser configuration as a key-value map",
			BodyPath: "parser",
		},
		&requestflag.Flag[string]{
			Name:     "referrer-type",
			Usage:    "Referrer policy for the request",
			BodyPath: "referrer_type",
		},
		&requestflag.Flag[any]{
			Name:     "render",
			Usage:    "Whether to render JavaScript content using a browser",
			BodyPath: "render",
		},
		&requestflag.Flag[float64]{
			Name:     "request-timeout",
			Usage:    "Request timeout in milliseconds",
			BodyPath: "request_timeout",
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
	},
	Action:          handleExtract,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"network-capture": {
		&requestflag.InnerFlag[string]{
			Name:       "network-capture.method",
			Usage:      `Allowed values: "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH".`,
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
})

var extractAsync = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract-async",
	Usage:   "Extract Async Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "url",
			Usage:    "Target URL to scrape",
			Required: true,
			BodyPath: "url",
		},
		&requestflag.Flag[any]{
			Name:     "body",
			Usage:    "Request body for POST, PUT, PATCH methods",
			BodyPath: "body",
		},
		&requestflag.Flag[any]{
			Name:     "browser",
			Usage:    "Browser type to emulate",
			BodyPath: "browser",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "browser-action",
			Usage:    "Array of browser automation actions to execute sequentially",
			BodyPath: "browser_actions",
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "URL to call back when async operation completes",
			BodyPath: "callback_url",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "City for geolocation",
			BodyPath: "city",
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
		&requestflag.Flag[string]{
			Name:     "driver",
			Usage:    "Browser driver to use",
			BodyPath: "driver",
		},
		&requestflag.Flag[[]int64]{
			Name:     "expected-status-code",
			Usage:    "Expected HTTP status codes for successful requests",
			BodyPath: "expected_status_codes",
		},
		&requestflag.Flag[[]string]{
			Name:     "format",
			Usage:    "List of acceptable response formats in order of preference",
			BodyPath: "formats",
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
		&requestflag.Flag[string]{
			Name:     "markdown-backend",
			Usage:    `Selects which markdown conversion strategy to use. "full_page" converts the entire HTML page. "main_content" uses Mozilla Readability to extract the main article content before converting.`,
			Default:  "full_page",
			BodyPath: "markdown_backend",
		},
		&requestflag.Flag[string]{
			Name:     "method",
			Usage:    "HTTP method for the request",
			Default:  "GET",
			BodyPath: "method",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "network-capture",
			Usage:    "Filters for capturing network traffic",
			BodyPath: "network_capture",
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
		&requestflag.Flag[any]{
			Name:     "parser",
			Usage:    "Custom parser configuration as a key-value map",
			BodyPath: "parser",
		},
		&requestflag.Flag[string]{
			Name:     "referrer-type",
			Usage:    "Referrer policy for the request",
			BodyPath: "referrer_type",
		},
		&requestflag.Flag[any]{
			Name:     "render",
			Usage:    "Whether to render JavaScript content using a browser",
			BodyPath: "render",
		},
		&requestflag.Flag[float64]{
			Name:     "request-timeout",
			Usage:    "Request timeout in milliseconds",
			BodyPath: "request_timeout",
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
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "US state for geolocation (only valid when country is US)",
			BodyPath: "state",
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
		&requestflag.Flag[string]{
			Name:     "tag",
			Usage:    "User-defined tag for request identification",
			BodyPath: "tag",
		},
	},
	Action:          handleExtractAsync,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"network-capture": {
		&requestflag.InnerFlag[string]{
			Name:       "network-capture.method",
			Usage:      `Allowed values: "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH".`,
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
})

var extractBatch = requestflag.WithInnerFlags(cli.Command{
	Name:    "extract-batch",
	Usage:   "Extract Batch Endpoint",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "input",
			Usage:    "Array of extraction requests. Each object can include extraction parameters and async/storage settings.",
			Required: true,
			BodyPath: "inputs",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "shared-inputs",
			Usage:    "Shared parameters applied to the entire batch. Can include extraction parameters and async/storage settings.",
			BodyPath: "shared_inputs",
		},
	},
	Action:          handleExtractBatch,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"input": {
		&requestflag.InnerFlag[any]{
			Name:       "input.body",
			Usage:      "Request body for POST, PUT, PATCH methods",
			InnerField: "body",
		},
		&requestflag.InnerFlag[any]{
			Name:       "input.browser",
			Usage:      "Browser type to emulate",
			InnerField: "browser",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "input.browser-actions",
			Usage:      "Array of browser automation actions to execute sequentially",
			InnerField: "browser_actions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.callback-url",
			Usage:      "URL to call back when async operation completes",
			InnerField: "callback_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.city",
			Usage:      "City for geolocation",
			InnerField: "city",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "input.consent-header",
			Usage:      "Whether to automatically handle cookie consent headers",
			InnerField: "consent_header",
		},
		&requestflag.InnerFlag[any]{
			Name:       "input.cookies",
			Usage:      "Browser cookies as array of cookie objects",
			InnerField: "cookies",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.country",
			Usage:      "Country code for geolocation and proxy selection",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.device",
			Usage:      "Device type for browser emulation",
			InnerField: "device",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.driver",
			Usage:      "Browser driver to use",
			InnerField: "driver",
		},
		&requestflag.InnerFlag[[]int64]{
			Name:       "input.expected-status-codes",
			Usage:      "Expected HTTP status codes for successful requests",
			InnerField: "expected_status_codes",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "input.formats",
			Usage:      "List of acceptable response formats in order of preference",
			InnerField: "formats",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "input.headers",
			Usage:      "Custom HTTP headers to include in the request",
			InnerField: "headers",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "input.http2",
			Usage:      "Whether to use HTTP/2 protocol",
			InnerField: "http2",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "input.is-xhr",
			Usage:      "Whether to emulate XMLHttpRequest behavior",
			InnerField: "is_xhr",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.locale",
			Usage:      "Locale for browser language and region settings",
			InnerField: "locale",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.markdown-backend",
			Usage:      `Selects which markdown conversion strategy to use. "full_page" converts the entire HTML page. "main_content" uses Mozilla Readability to extract the main article content before converting.`,
			InnerField: "markdown_backend",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.method",
			Usage:      "HTTP method for the request",
			InnerField: "method",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "input.network-capture",
			Usage:      "Filters for capturing network traffic",
			InnerField: "network_capture",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.os",
			Usage:      "Operating system to emulate",
			InnerField: "os",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "input.parse",
			Usage:      "Whether to parse the response content",
			InnerField: "parse",
		},
		&requestflag.InnerFlag[any]{
			Name:       "input.parser",
			Usage:      "Custom parser configuration as a key-value map",
			InnerField: "parser",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.referrer-type",
			Usage:      "Referrer policy for the request",
			InnerField: "referrer_type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "input.render",
			Usage:      "Whether to render JavaScript content using a browser. Use 'auto' to let the engine select the candidate config per domain.",
			InnerField: "render",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "input.request-timeout",
			Usage:      "Request timeout in milliseconds",
			InnerField: "request_timeout",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "input.session",
			InnerField: "session",
		},
		&requestflag.InnerFlag[any]{
			Name:       "input.skill",
			Usage:      "Skills or capabilities required for the request",
			InnerField: "skill",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.state",
			Usage:      "US state for geolocation (only valid when country is US)",
			InnerField: "state",
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
		&requestflag.InnerFlag[string]{
			Name:       "input.tag",
			Usage:      "User-defined tag for request identification",
			InnerField: "tag",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.url",
			Usage:      "Target URL to scrape",
			InnerField: "url",
		},
	},
	"shared-inputs": {
		&requestflag.InnerFlag[any]{
			Name:       "shared-inputs.body",
			Usage:      "Request body for POST, PUT, PATCH methods",
			InnerField: "body",
		},
		&requestflag.InnerFlag[any]{
			Name:       "shared-inputs.browser",
			Usage:      "Browser type to emulate",
			InnerField: "browser",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "shared-inputs.browser-actions",
			Usage:      "Array of browser automation actions to execute sequentially",
			InnerField: "browser_actions",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.callback-url",
			Usage:      "URL to call back when async operation completes",
			InnerField: "callback_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.city",
			Usage:      "City for geolocation",
			InnerField: "city",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "shared-inputs.consent-header",
			Usage:      "Whether to automatically handle cookie consent headers",
			InnerField: "consent_header",
		},
		&requestflag.InnerFlag[any]{
			Name:       "shared-inputs.cookies",
			Usage:      "Browser cookies as array of cookie objects",
			InnerField: "cookies",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.country",
			Usage:      "Country code for geolocation and proxy selection",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.device",
			Usage:      "Device type for browser emulation",
			InnerField: "device",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.driver",
			Usage:      "Browser driver to use",
			InnerField: "driver",
		},
		&requestflag.InnerFlag[[]int64]{
			Name:       "shared-inputs.expected-status-codes",
			Usage:      "Expected HTTP status codes for successful requests",
			InnerField: "expected_status_codes",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "shared-inputs.formats",
			Usage:      "List of acceptable response formats in order of preference",
			InnerField: "formats",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "shared-inputs.headers",
			Usage:      "Custom HTTP headers to include in the request",
			InnerField: "headers",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "shared-inputs.http2",
			Usage:      "Whether to use HTTP/2 protocol",
			InnerField: "http2",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "shared-inputs.is-xhr",
			Usage:      "Whether to emulate XMLHttpRequest behavior",
			InnerField: "is_xhr",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.locale",
			Usage:      "Locale for browser language and region settings",
			InnerField: "locale",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.markdown-backend",
			Usage:      `Selects which markdown conversion strategy to use. "full_page" converts the entire HTML page. "main_content" uses Mozilla Readability to extract the main article content before converting.`,
			InnerField: "markdown_backend",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.method",
			Usage:      "HTTP method for the request",
			InnerField: "method",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "shared-inputs.network-capture",
			Usage:      "Filters for capturing network traffic",
			InnerField: "network_capture",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.os",
			Usage:      "Operating system to emulate",
			InnerField: "os",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "shared-inputs.parse",
			Usage:      "Whether to parse the response content",
			InnerField: "parse",
		},
		&requestflag.InnerFlag[any]{
			Name:       "shared-inputs.parser",
			Usage:      "Custom parser configuration as a key-value map",
			InnerField: "parser",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.referrer-type",
			Usage:      "Referrer policy for the request",
			InnerField: "referrer_type",
		},
		&requestflag.InnerFlag[any]{
			Name:       "shared-inputs.render",
			Usage:      "Whether to render JavaScript content using a browser. Use 'auto' to let the engine select the candidate config per domain.",
			InnerField: "render",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "shared-inputs.request-timeout",
			Usage:      "Request timeout in milliseconds",
			InnerField: "request_timeout",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "shared-inputs.session",
			InnerField: "session",
		},
		&requestflag.InnerFlag[any]{
			Name:       "shared-inputs.skill",
			Usage:      "Skills or capabilities required for the request",
			InnerField: "skill",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.state",
			Usage:      "US state for geolocation (only valid when country is US)",
			InnerField: "state",
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
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.tag",
			Usage:      "User-defined tag for request identification",
			InnerField: "tag",
		},
		&requestflag.InnerFlag[string]{
			Name:       "shared-inputs.url",
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
			Usage:    "Country code for geo-targeted results (e.g., 'US', 'GB', 'IL')",
			Default:  "US",
			BodyPath: "country",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "debug-params",
			Usage:    "Internal-only. Gated to allowlisted accounts; ignored otherwise.",
			BodyPath: "debug_params",
		},
		&requestflag.Flag[*bool]{
			Name:     "deep-search",
			Usage:    "Deprecated. Use search_depth instead. true maps to 'deep', false maps to 'lite'.",
			BodyPath: "deep_search",
		},
		&requestflag.Flag[*string]{
			Name:     "end-date",
			Usage:    "Filter results before this date (format: YYYY-MM-DD or YYYY)",
			BodyPath: "end_date",
		},
		&requestflag.Flag[any]{
			Name:     "exclude-domain",
			Usage:    "List of domains to exclude from search results. Maximum 50 domains.",
			BodyPath: "exclude_domains",
		},
		&requestflag.Flag[any]{
			Name:     "focus",
			Usage:    "Search focus mode (e.g., 'general', 'news', 'shopping') or a list of explicit subagent names (e.g., ['amazon_serp', 'target_serp'])",
			Default:  "general",
			BodyPath: "focus",
		},
		&requestflag.Flag[bool]{
			Name:     "include-answer",
			Usage:    "Generate an LLM-powered answer summary based on search result snippets.",
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
			Usage:    "Language/locale code (e.g., 'en', 'fr', 'de')",
			Default:  "en",
			BodyPath: "locale",
		},
		&requestflag.Flag[int64]{
			Name:     "max-results",
			Usage:    "Maximum number of results to return. Actual count may be lower depending on availability.",
			Default:  3,
			BodyPath: "max_results",
		},
		&requestflag.Flag[int64]{
			Name:     "max-subagents",
			Usage:    "Maximum number of subagents to execute in parallel for WSA focus modes (shopping, social, geo). Ignored for SERP focus modes.",
			Default:  3,
			BodyPath: "max_subagents",
		},
		&requestflag.Flag[string]{
			Name:     "output-format",
			Usage:    "Output format: plain_text, markdown, or simplified_html",
			Default:  "markdown",
			BodyPath: "output_format",
		},
		&requestflag.Flag[*string]{
			Name:     "search-depth",
			Usage:    "Controls content richness and latency of search results.\n\n- lite: Token-efficient metadata for high-volume pipelines (title, URL, description only)\n- fast: Rich content (~2K chars) optimized for AI agents\n- deep: Full page content via Webit scraping for comprehensive analysis",
			BodyPath: "search_depth",
		},
		&requestflag.Flag[*string]{
			Name:     "start-date",
			Usage:    "Filter results after this date (format: YYYY-MM-DD or YYYY)",
			BodyPath: "start_date",
		},
		&requestflag.Flag[*string]{
			Name:     "time-range",
			Usage:    "Time range filters passed to Webit SERP API as 'time' parameter.",
			BodyPath: "time_range",
		},
	},
	Action:          handleSearch,
	HideHelpCommand: true,
}

func handleExtract(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.ExtractParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Extract(ctx, params, options...)
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
		Title:          "extract",
		Transform:      transform,
	})
}

func handleExtractAsync(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.ExtractAsyncParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExtractAsync(ctx, params, options...)
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
		Title:          "extract-async",
		Transform:      transform,
	})
}

func handleExtractBatch(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.ExtractBatchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ExtractBatch(ctx, params, options...)
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
		Title:          "extract-batch",
		Transform:      transform,
	})
}

func handleMap(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.MapParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Map(ctx, params, options...)
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
		Title:          "map",
		Transform:      transform,
	})
}

func handleSearch(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomnimblewaynimblego.SearchParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Search(ctx, params, options...)
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
		Title:          "search",
		Transform:      transform,
	})
}
