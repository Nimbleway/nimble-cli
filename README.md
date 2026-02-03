# Nimbleway CLI

The official CLI for the [Nimbleway REST API](docs.nimbleway.com).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

To test or install the CLI locally, you need [Go](https://go.dev/doc/install) version 1.22 or later installed.

```sh
go install 'github.com/stainless-sdks/nimbleway-cli/cmd/nimble@latest'
```

Once you have run `go install`, the binary is placed in your Go bin directory:

- **Default location**: `$HOME/go/bin` (or `$GOPATH/bin` if GOPATH is set)
- **Check your path**: Run `go env GOPATH` to see the base directory

If commands aren't found after installation, add the Go bin directory to your PATH:

```sh
# Add to your shell profile (.zshrc, .bashrc, etc.)
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Running Locally

After cloning the git repository for this project, you can use the
`scripts/run` script to run the tool locally:

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
nimble [resource] <command> [flags...]
```

```sh
nimble extract \
  --debug-options '{collect_har: true, no_retry_mode: true, record_screen: true, redact: true, show_cursor: true, solve_captcha: true, trace: true, upload_engine_logs: true, verbose: true, with_proxy_usage: true}' \
  --url https://example.com \
  --browser chrome \
  --city 'Los Angeles' \
  --client-timeout 25000 \
  --consent-header \
  --cookies '{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}' \
  --country US \
  --device desktop \
  --disable-ip-check \
  --driver vx8 \
  --dynamic-parser '{myParser: bar}' \
  --expected-status-code 200 \
  --expected-status-code 201 \
  --export-userbrowser \
  --format json \
  --headers '{User-Agent: CustomBot/1.0, Accept-Language: en-US}' \
  --http2 \
  --ip6 \
  --is-xhr \
  --locale en-US \
  --markdown \
  --metadata '{account_name: acme-corp, definition_id: 456, definition_name: product-scraper, endpoint: /api/v2/scrape, execution_id: exec-abc123, flowit_task_id: task-xyz789, input_id: input-123, pipeline_execution_id: 12345, query_template_id: template-qry-001, source: web-app, template_id: 789, template_name: e-commerce-template}' \
  --method GET \
  --native-mode requester \
  --network-capture '{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}' \
  --no-html \
  --no-userbrowser \
  --os windows \
  --parse \
  --parse-options '{merge_dynamic: true}' \
  --parser '{myParser: bar}' \
  --proxy-provider brightdata \
  --proxy-providers '{brightdata: 70, oxylabs: 30}' \
  --query-template '{id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, api_type: WEB, pagination: {next_page_params: {foo: bar}}, params: {foo: bar}}' \
  --raw-headers \
  --referrer-type random \
  --render \
  --render-flow '{wait: bar}' \
  --render-flow '{click: bar}' \
  --render-options '{adblock: true, blocked_domains: [ads.example.com, tracker.com], browser_engine: chrome, cache: false, connector_type: webit-cdp, disabled_resources: [image, stylesheet], enable_2captcha: true, extensions: [extension-id-1, extension-id-2], fingerprint_id: fp-abc123, hackium_configuration: {collect_logs: true, do_not_fix_math_salt: true, enable_document_element_spoof: true, enable_document_has_focus: true, enable_fake_navigation_history: true, enable_key_ordering: true, enable_sniffer: true, enable_verbose_logs: true}, headless: true, include_iframes: true, load_local_storage: true, local_storage_keys_to_load: [authToken, userId], mouse_strategy: linear, no_accept_encoding: true, override_permissions: true, random_header_order: true, render_type: load, store_local_storage: true, timeout: 30000, typing_interval: 100, typing_strategy: simple, userbrowser: true, wait_until: networkidle2, with_performance_metrics: true}' \
  --request-timeout 30000 \
  --return-response-headers-as-header \
  --save-userbrowser \
  --session '{id: id, prefetch_userbrowser: true, retry: true, timeout: 1}' \
  --skill dynamic-content \
  --skip-ubct \
  --state CA \
  --tag campaign-2024-q1 \
  --template '{name: x, params: {foo: bar}}' \
  --type generic \
  --userbrowser-creation-template-rendered '{id: id, allowed_parameter_names: [x], render_flow_rendered: [{foo: bar}]}'
```

For details about specific commands, use the `--help` flag.

### Global Flags

- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)

### Passing files as arguments

To pass files to your API, you can use the `@myfile.ext` syntax:

```bash
nimble <command> --arg @abe.jpg
```

Files can also be passed inside JSON or YAML blobs:

```bash
nimble <command> --arg '{image: "@abe.jpg"}'
# Equivalent:
nimble <command> <<YAML
arg:
  image: "@abe.jpg"
YAML
```

If you need to pass a string literal that begins with an `@` sign, you can
escape the `@` sign to avoid accidentally passing a file.

```bash
nimble <command> --username '\@abe'
```

#### Explicit encoding

For JSON endpoints, the CLI tool does filetype sniffing to determine whether the
file contents should be sent as a string literal (for plain text files) or as a
base64-encoded string literal (for binary files). If you need to explicitly send
the file as either plain text or base64-encoded data, you can use
`@file://myfile.txt` (for string encoding) or `@data://myfile.dat` (for
base64-encoding). Note that absolute paths will begin with `@file://` or
`@data://`, followed by a third `/` (for example, `@file:///tmp/file.txt`).

```bash
nimble <command> --arg @data://file.txt
```
