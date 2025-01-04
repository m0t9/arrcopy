// Package arrcopy provides analyzer prohibiting for-range loops with arrays copying.
package arrcopy

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// NewAnalyzer returns analyzer instance for arrcopy linter.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "arrcopy",
		Doc:      "reports for-range looping over arrays with copying",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	filter := []ast.Node{(*ast.RangeStmt)(nil)}

	inspector.Preorder(filter, func(n ast.Node) {
		rng := n.(*ast.RangeStmt)
		vt := pass.TypesInfo.TypeOf(rng.X)
		if _, isArray := vt.(*types.Array); isArray {
			pass.Reportf(rng.Pos(),
				`for-range loop over array '%s' found. Use for-range over '&%s' instead`,
				rng.X, rng.X)
		}
	})
	return nil, nil
}
