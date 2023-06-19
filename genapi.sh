#!/bin/bash
#
# lxy@20230612
#
set -ve
go build

go generate generate.go

