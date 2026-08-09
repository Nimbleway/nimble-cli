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

## Authentication

Run `nimble login` once and the CLI stores your credentials, so later commands
need no `--api-key` flag:

```sh
nimble login
```

Login offers two methods:

- **Browser** (default): opens your browser to approve access, then stores the
  API key it retrieves for your account.
- **Paste an API key**: enter a key directly. It is validated before it is saved.

Two related commands round this out:

```sh
nimble whoami   # show the active credential and where it came from
nimble logout   # remove the stored credential
```

Credentials are written to `~/.nimble/credentials.json` with `0600` permissions.
Set `NIMBLE_CONFIG_DIR` to store them elsewhere.

### Credential priority

When a command needs an API key, the CLI uses the first source available:

1. `--api-key` flag (explicit per-command override)
2. Stored credential (from `nimble login`)
3. `NIMBLE_API_KEY` environment variable

Note that a stored credential takes precedence over `NIMBLE_API_KEY`. If a
command uses an unexpected key, run `nimble whoami` to see which source is
active, and `nimble logout` to fall back to the environment variable.

## Usage

The CLI follows a resource-based command structure:

```sh
nimble [resource] <command> [flags...]
```

```sh
nimble extract run --url https://example.com
```

To override the stored credential for a single command, pass `--api-key`:

```sh
nimble extract run \
  --api-key 'My API Key' \
  --url https://example.com
```

For details about specific commands, use the `--help` flag.

### Environment variables

| Environment variable | Required | Default value  |
| -------------------- | -------- | -------------- |
| `NIMBLE_API_KEY`     | no       | `null`         |
| `NIMBLE_CONFIG_DIR`  | no       | `~/.nimble`    |
| `CLIENT_SOURCE`      | no       | `"sdk"`        |

### Global flags

- `--api-key` (can also be set with `NIMBLE_API_KEY` env var, or stored via `nimble login`)
- `--client-source` (can also be set with `CLIENT_SOURCE` env var)
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

## Linking different Go SDK versions

You can link the CLI against a different version of the Nimble Go SDK
for development purposes using the `./scripts/link` script.

To link to a specific version from a repository (version can be a branch,
git tag, or commit hash):

```bash
./scripts/link github.com/org/repo@version
```

To link to a local copy of the SDK:

```bash
./scripts/link ../path/to/githubcomnimblewaynimblego-go
```

If you run the link script without any arguments, it will default to `../githubcomnimblewaynimblego-go`.
