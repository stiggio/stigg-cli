# Stigg CLI

The official CLI for the Stigg REST API.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/stigg-cli/cmd/stigg@latest'
```

### Running Locally

```sh
go run cmd/stigg/main.go
```

## Usage

The CLI follows a resource-based command structure:

```sh
stigg [resource] [command] [flags]
```

```sh
stigg v1:customers retrieve \
  --id REPLACE_ME
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
