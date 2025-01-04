package render

import (
	"bytes"
	"go/printer"
	"go/token"
)

func Node(fset *token.FileSet, x any) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}

	return buf.String()
}
