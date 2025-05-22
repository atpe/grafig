#!/bin/bash

usage() { 
  cat << EndOfMessage
Builds, and optionally serves, the project coverage.

Usage:
  yarn coverage [flags]

Flags:
  -b, --badges      Generate coverage badges.
  -h, --help        Help for coverage.
  -s, --serve       Serve the generated documentation.
EndOfMessage
}

for arg in "$@"; do
  shift
  case "$arg" in
    '--badges')   set -- "$@" '-b'   ;;
    '--help')     set -- "$@" '-h'   ;;
    '--serve')    set -- "$@" '-s'   ;;
    *)            set -- "$@" "$arg" ;;
  esac
done

badges="false"
serve="false"

OPTIND=1
while getopts "bhs" opt
do
  case "$opt" in
    'b') badges="true" ;;
    'h') usage; exit 0 ;;
    's') serve=true ;;
    '?') usage >&2; exit 1 ;;
  esac
done
shift $(( OPTIND - 1))

jest --coverage --silent --verbose false 

if [ $badges = "true" ]; then jest-coverage-badges; fi
if [ $serve = "true" ]; then parcel coverage/lcov-report/index.html --port 3002; fi
