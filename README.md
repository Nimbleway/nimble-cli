# Nimble CLI

The official CLI for the [Nimble REST API](docs.nimbleway.com).

It is generated with [Stainless](https://www.stainless.com/).

<!-- x-release-please-start-version -->

## Installation

### Installing with Go

To test or install the CLI locally, you need [Go](https://go.dev/doc/install) version 1.22 or later installed.

```sh
go install 'github.com/Nimbleway/nimble-cli/cmd/nimble@latest'
```

Once you have run `go install`, the binary is placed in your Go bin directory:

- **Default location**: `$HOME/go/bin` (or `$GOPATH/bin` if GOPATH is set)
- **Check your path**: Run `go env GOPATH` to see the base directory

If commands aren't found after installation, add the Go bin directory to your PATH:

```sh
# Add to your shell profile (.zshrc, .bashrc, etc.)
export PATH="$PATH:$(go env GOPATH)/bin"
```

<!-- x-release-please-end -->

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
nimble crawl \
  --url url \
  --allow-external-links \
  --allow-subdomains \
  --callback '{url: https://example.com, events: [started], headers: {foo: string}, metadata: {foo: bar}}' \
  --crawl-entire-domain \
  --exclude-path /exclude-this-path \
  --exclude-path /and-this-path \
  --extract-options "{browser: chrome, browser_actions: [{goto: https://example.com/login}, {wait_for_element: '#login-form'}, {fill: {selector: '#username', value: user@example.com, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {fill: {selector: '#password', value: password123, click_on_element: true, delay: 1000, mode: type, mouse_movement_strategy: linear, required: 'true', scroll: true, skip: 'true', timeout: 0, typing_interval: 1000, typing_strategy: simple, visible: true}}, {click: '#submit'}, {screenshot: {format: png, full_page: true, quality: 0, required: 'true', skip: 'true'}}], city: Los Angeles, consent_header: true, cookies: [{creation: creation, domain: domain, expires: expires, extensions: [string], hostOnly: true, httpOnly: true, lastAccessed: lastAccessed, maxAge: Infinity, name: name, path: path, pathIsDefault: true, sameSite: strict, secure: true, value: value}], country: US, device: desktop, driver: vx8, expected_status_codes: [200, 201], formats: [html], headers: {User-Agent: CustomBot/1.0, Accept-Language: en-US}, http2: true, is_xhr: true, locale: en-US, method: GET, network_capture: [{method: GET, resource_type: document, status_code: 100, url: {value: value, type: exact}, validation: true, wait_for_requests_count: 0, wait_for_requests_count_timeout: 1}], os: windows, parse: true, parser: {myParser: bar}, referrer_type: random, render: true, request_timeout: 30000, session: {id: id, prefetch_userbrowser: true, retry: true, timeout: 1}, skill: dynamic-content, state: CA, tag: campaign-2024-q1, url: url}" \
  --ignore-query-parameters \
  --include-path /include-this-path \
  --include-path /and-this-path \
  --limit 100 \
  --max-discovery-depth 3 \
  --name 'The best crawl ever' \
  --sitemap include
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
