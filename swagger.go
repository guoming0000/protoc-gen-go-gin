package main

import (
	"os"

	"github.com/guoming0000/protoc-gen-go-gin/cmd/protoc-gen-openapi/generator"
)

func main() {
	generator.RecreateProto(os.Args[1], os.Args[2])
}
