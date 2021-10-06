#!/bin/bash
go clean -testcache
go test ./interfaces/... $1
