#!/bin/bash

usage() { 
  cat << EndOfMessage
Prepares the @grafig/lib project running the following stages:
  - yarn lint
  - yarn coverage --badges
  - yarn docs
  - yarn build

Usage:
  yarn prepare [flags]

Flags:
  -B, --build     Include/exclude the build stage.
  -D, --docs      Include/exclude the docs stage.
  -h, --help      Help for prepare.
  -i, --include   Run the stages specified with flags only.
  -L, --lint      Include/exclude the lint stage.
  -T, --test      Include/exclude the test stage.
  -x, --exclude   Run the stages NOT specified with flags only.
EndOfMessage
}

if [ $# -eq 0 ]; then
  yarn lint
  yarn coverage --badges
  yarn docs
  yarn build
  exit 0
fi

for arg in "$@"; do
  shift
  case "$arg" in
    '--build')    set -- "$@" '-B'   ;;
    '--docs')     set -- "$@" '-D'   ;;
    '--help')     set -- "$@" '-h'   ;;
    '--include')  set -- "$@" '-i'   ;;
    '--lint')     set -- "$@" '-L'   ;;
    '--test')     set -- "$@" '-T'   ;;
    '--exclude')  set -- "$@" '-x'   ;;
    *)            set -- "$@" "$arg" ;;
  esac
done

build=false
docs=false
lint=false
test=false
include=false
exclude=false

OPTIND=1
while getopts "BDhiLTx" opt
do
  case "$opt" in
    'B') build=true ;;
    'D') docs=true ;;
    'h') usage; exit 0 ;;
    'i') include=true ;;
    'L') lint=true 0 ;;
    'T') test=true ;;
    'x') exclude=true ;;
    '?') usage >&2; exit 1 ;;
  esac
done
shift $(( OPTIND - 1))

if [ $include = $exclude ]; then
  usage >&2; exit 1
fi

if [ $include = "true" ]; then
  if [ $lint = "true" ] && ! yarn lint; then exit 1;fi
  if [ $test = "true" ] && ! yarn coverage --badges; then exit 1;fi
  if [ $docs = "true" ] && ! yarn docs; then exit 1;fi
  if [ $build = "true" ] && ! yarn build; then exit 1;fi
elif [ $exclude = "true" ]; then
  if [ $lint = "false" ] && ! yarn lint; then exit 1; fi
  if [ $test = "false" ] && ! yarn coverage --badges; then exit 1; fi
  if [ $docs = "false" ] && ! yarn docs; then exit 1; fi
  if [ $build = "false" ] && ! yarn build; then exit 1; fi
fi