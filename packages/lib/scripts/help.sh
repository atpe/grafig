#!/bin/bash

usage() {
  cat << EndOfMessage
Use a subcommand with --help to see more usage information.

Usage:
  yarn run [command]

Available Commands:
  build       Build the @grafig/lib distributable.
  coverage    Generate project coverage.
  docs        Generate project documentation.
  help        Help for @grafig/lib.
  lint        Lint project code.
  prepare     Prepare the @grafig/lib project.
  test        Test project code.
Flags:
  -h, --help   help for yarn

Use "yarn run [command] --help" for more information about a command.
EndOfMessage
}

case "$1" in
  'build')    yarn build --help ;;
  'lint')     yarn lint --help ;;
  'docs')     yarn docs --help ;;
  'test')     yarn test --help ;;
  *) usage ;;
esac


