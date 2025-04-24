# @grafig/syntax

![Golang](https://img.shields.io/badge/golang-grey?logo=go) ![Version](https://img.shields.io/github/package-json/v/atpe/grafig?filename=packages/syntax/package.json)

A tool to transpile _FigScript_ into vanilla JavaScript.

## Installation

>This project requires the [Golang distributable](https://go.dev/doc/install) and [ANTLR4](https://www.antlr.org/download.html) to be installed, with their binaries included in the `PATH` environment variable. Additionally, to use the provided [transformGrammar.py](https://github.com/antlr/grammars-v4/blob/master/javascript/javascript/Go/transformGrammar.py) script, `python` must also be installed and included in `PATH`.

If `yarn install` was run in the workspace root, the installation process should already be completed. If not, run `yarn install` from within this package. This will additionally run the `prepare` script to download and transform the required grammar and parser files into the `internal/parser` directory.

Once the installation has been completed successfully, the binary can be built by running:

```bash
yarn build
```

## Usage

Once the binary has been sucessfully built, all operations are aliased through `yarn` - e.g.:

```bash
yarn syntax transpile example/file.fs
# is the same as running
bin/figsyntax transpile example/file.fs
```

The available commands are:

- `parse`
- `analyse`
- `transpile`

To see the usage of any command run:

```bash
yarn syntax [command] --help
```
