#!/bin/bash

base_url=https://raw.githubusercontent.com/antlr/grammars-v4/refs/heads/master/javascript/javascript

printf "downloading grammar files...\n"

curl -s --output internal/parser/JavaScriptLexer.g4 "${base_url}/JavaScriptLexer.g4"
curl -s --output internal/parser/JavaScriptParser.g4 "${base_url}/JavaScriptParser.g4"

printf "downloading parser files...\n"

curl -s --output internal/parser/javascript_lexer_base.go "${base_url}/Go/javascript_lexer_base.go"
curl -s --output internal/parser/javascript_parser_base.go "${base_url}/Go/javascript_parser_base.go"