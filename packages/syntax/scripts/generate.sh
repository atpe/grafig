#!/bin/bash

printf "generating parsers..."

go generate ./...
go mod tidy