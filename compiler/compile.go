package compiler

import (
	"endito/document"
)

// Piler struct maintains a reference to the root
// of its document tree to compile
type Piler struct {
	Root *document.Node
}

// NewPiler constructor returns a new Piler
func NewPiler(r *document.Node) *Piler {
	return &Piler{
		Root: r,
	}
}

// Compile is responsible for tokenizing the existing document, parsing the
// tokens, and executing the requested updates
//
// The compiler struct contains references to a document tree that it is
// responsible for compiling
func (p *Piler) Compile() {
}
