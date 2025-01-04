// Package arrcopy provides analyzer prohibiting for-range loops with arrays copying.
package arrcopy

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

// NewAnalyzer returns analyzer instance for arrcopy linter.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "arrcopy",
		Doc:  "reports for-range looping over arrays with copying",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			rng, ok := n.(*ast.RangeStmt)
			if !ok {
				return true
			}
			vt := pass.TypesInfo.TypeOf(rng.X)
			if _, isArray := vt.(*types.Array); isArray {
				pass.Reportf(rng.Pos(),
					`for-range loop over array '%s' found. Use for-range over '&%s' instead`,
					rng.X, rng.X)
			}
			return true
		})
	}
	return nil, nil
}
