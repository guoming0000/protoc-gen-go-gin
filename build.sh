#!/bin/bash
#
# lxy@20230907
#
set -ve
cd cmd/protoc-gen-openapi/ && go build && cd ../../
cd cmd/protoc-gen-go-errors/ && go build && cd ../../
cd cmd/protoc-gen-go-gin/ && go build && cd ../../

