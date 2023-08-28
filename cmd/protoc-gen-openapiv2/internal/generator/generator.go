// Package generator provides an abstract interface to code generators.
package generator

import (
	"github.com/guoming0000/protoc-gen-go-gin/cmd/protoc-gen-openapiv2/internal/descriptor"
)

// Generator is an abstraction of code generators.
type Generator interface {
	// Generate generates output files from input .proto files.
	Generate(targets []*descriptor.File) ([]*descriptor.ResponseFile, error)
}
