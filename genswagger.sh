#!/bin/bash
#
# lxy@20230612
#
set -ve

which yq || {
   # https://github.com/mikefarah/yq
   # go install github.com/mikefarah/yq/v4@latest
   # brew install yq
   wget http://qiniu.brightguo.com/sunmi/yq -O /usr/local/bin/yq &&  chmod +x /usr/bin/yq
}

cd cmd/protoc-gen-openapi
go build
cd ../../

go run swagger.go api/article/auth.proto


protoc --proto_path=. \
        --proto_path=./third_party \
        --openapi_out=fq_schema_naming=true,default_response=false,output_mode=source_relative:. \
        api/article/auth.swagger.proto

yq -Poj api/article/auth.swagger.yaml > api/article/auth.swagger.json
# rm api/article/auth.swagger.yaml