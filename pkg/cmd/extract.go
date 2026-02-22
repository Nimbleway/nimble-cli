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

var extractAsync = requestflag.WithInnerFlags(cli.Command{
	Name:    "async",
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
			Name:     "browser",
			Usage:    "Browser type to emulate",
			BodyPath: "browser",
		},
		&requestflag.Flag[[]any]{
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
			Default:  []string{"html"},
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
		&requestflag.Flag[bool]{
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

func handleExtractAsync(ctx context.Context, cmd *cli.Command) error {
	client := githubcomnimblewaynimblego.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := githubcomnimblewaynimblego.ExtractAsyncParams{}

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
	_, err = client.Extract.Async(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "extract async", obj, format, transform)
}
