package render

import (
	"bytes"
	"go/printer"
	"go/token"
)

// Node renders an AST node (x) using provided token.FileSet.
func Node(fset *token.FileSet, x any) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}

	return buf.String()
}
