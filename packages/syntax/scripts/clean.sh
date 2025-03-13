#!/bin/bash

printf "cleaning package..."

git clean -Xfq 'bin/*' 'internal/*'