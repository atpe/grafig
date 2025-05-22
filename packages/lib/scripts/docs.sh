#!/bin/bash

usage() { 
  cat << EndOfMessage
Builds the project documention.

Usage:
  yarn docs [flags]

Flags:
  -h, --help        Help for docs.
  -s, --serve       Serve the generated documentation.
EndOfMessage
}

for arg in "$@"; do
  shift
  case "$arg" in
    '--help')     set -- "$@" '-h'   ;;
    '--serve')    set -- "$@" '-s'   ;;
    *)            set -- "$@" "$arg" ;;
  esac
done

serve="false"

OPTIND=1
while getopts "hs" opt
do
  case "$opt" in
    'h') usage; exit 0 ;;
    's') serve=true ;;
    '?') usage >&2; exit 1 ;;
  esac
done
shift $(( OPTIND - 1))

typedoc "src/index.ts"

if [ $serve = "true" ]; then parcel docs/index.html --port 3001; fi