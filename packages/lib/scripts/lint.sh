#!/bin/bash

usage() { 
  cat << EndOfMessage
Lints the @grafig/lib project.

Usage:
  yarn lint [flags]

Flags:
  -h, --help            Help for build.
EndOfMessage
}

for arg in "$@"; do
  shift
  case "$arg" in
    '--help')     set -- "$@" '-h'   ;;
    *)            set -- "$@" "$arg" ;;
  esac
done

OPTIND=1
while getopts "h" opt
do
  case "$opt" in
    'h') usage; exit 0 ;;
    '?') usage >&2; exit 1 ;;
  esac
done
shift $(( OPTIND - 1))

eslint .

