#!/bin/bash

base_url=https://raw.githubusercontent.com/antlr/grammars-v4/refs/heads/master/javascript/javascript

printf "transforming grammar files...\n"

cd internal || exit

if hash python 2>/dev/null; then
  python -c "$(curl -s ${base_url}/Go/transformGrammar.py)" >/dev/null
else
  printf "grammar files not transformed: python does not exist in PATH\n"
  printf "either manually correct errors or install python and run 'yarn transform'"
  exit 1
fi