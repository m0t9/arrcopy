// Package arrcopy provides analyzer prohibiting for-range loops with arrays copying.
package arrcopy

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/m0t9/arrcopy/internal/render"
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

		// Optimization is not needed for loops ignoring array's items (value).
		if id, ok := rng.Value.(*ast.Ident); rng.Value == nil || (ok && id.Name == "_") {
			return
		}

		vt := pass.TypesInfo.TypeOf(rng.X)

		// Optimization is not applicable to call and cast expressions.
		switch rng.X.(type) {
		case *ast.CallExpr, *ast.TypeAssertExpr:
			return
		}

		if _, isArray := vt.(*types.Array); isArray {
			arr := render.Node(pass.Fset, rng.X)
			pass.Report(analysis.Diagnostic{
				Message: fmt.Sprintf(`for-range loop over array '%s' found. Use for-range over '&%s' instead`,
					arr, arr),
				Pos: rng.Pos(),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: fmt.Sprintf(`replace '%s' with '&%s'`, arr, arr),
						TextEdits: []analysis.TextEdit{
							{
								Pos:     rng.X.Pos(),
								End:     rng.X.End(),
								NewText: []byte(fmt.Sprintf(`&%s`, arr)),
							},
						},
					},
				},
			})
		}
	})

	return nil, nil
}
