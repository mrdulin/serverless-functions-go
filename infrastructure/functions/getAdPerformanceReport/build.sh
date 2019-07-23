#!/usr/bin/env bash

cd "$(dirname "$0")"
export GO111MODULE=on

go build -v -o ${GOPATH}/bin ./getAdPerformanceReport.go