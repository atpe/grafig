# figsyntax

![image](https://badgen.net/badge/version/v0.0.2/grey)

A tool for parsing _FigScript_ and transpiling into native JavaScript.

## Installation

This project requires the [Golang distributable](https://go.dev/doc/install) and [ANTLR4](https://www.antlr.org/download.html) to be installed, with their binaries included in the `PATH` environment variable. Additionally, to use the provided [transformGrammar.py](https://github.com/antlr/grammars-v4/blob/master/javascript/javascript/Go/transformGrammar.py) script, `python` must also be installed and included in `PATH`.

If `yarn install` was run in the root package, the installation process should already be completed, otherwise run `yarn install` in this directory.

The `preinstall` script, which runs before `install`, is equivalent to running:

```bash
yarn clean      # delete untracked files
yarn download   # download ANTLR files
yarn transform  # transform grammars
yarn generate   # generate parsers
yarn build      # build package binary
```

Once the installation has been completed successfully, build the binary:

```bash
yarn build
```

## Usage

Once the binary has been sucessfully built, all operations are aliased through `yarn` - e.g.:

```bash
yarn parse example/simple.js
# OR
bin/figsyntax parse example/simple.js
```
