#!/bin/bash

yarn syntax:analyse -s -o "test/test.json" "test/*.js"
cr -itQ -f json -o "test/control.json" "test/"