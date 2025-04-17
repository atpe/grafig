#!/bin/bash

if
  [ -e internal/parser/JavaScriptLexer.g4 ] &&
  [ -e internal/parser/JavaScriptParser.g4 ] &&
  [ -e internal/parser/javascript_lexer_base.go ] &&
  [ -e internal/parser/javascript_parser_base.go ]
then
  exit 0
else
  yarn prepare:download
  yarn prepare:transform
  yarn prepare:generate
fi
