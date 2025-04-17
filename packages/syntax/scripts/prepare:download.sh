#!/bin/bash

base_url=https://raw.githubusercontent.com/antlr/grammars-v4/refs/heads/master/javascript/javascript

printf "downloading required grammar files...\n"

curl -s --output internal/parser/JavaScriptLexer.g4 "${base_url}/JavaScriptLexer.g4"
curl -s --output internal/parser/JavaScriptParser.g4 "${base_url}/JavaScriptParser.g4"

printf "downloading required parser files...\n"

curl -s --output internal/parser/javascript_lexer_base.go "${base_url}/Go/javascript_lexer_base.go"
curl -s --output internal/parser/javascript_parser_base.go "${base_url}/Go/javascript_parser_base.go"

printf "transforming grammar files...\n"

cd internal || exit

if hash python 2>/dev/null; then
  python -c "$(curl -s ${base_url}/Go/transformGrammar.py)" >/dev/null
else
  printf "grammar files not transformed: python does not exist in PATH\n"
  printf "either manually correct errors or install python and run 'yarn transform'"
  exit 1
fi