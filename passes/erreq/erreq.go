package erreq

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "erreq",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

const Doc = "erreq finds binary expression which is comparing error"

var errType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.BinaryExpr)(nil)}

	inspect.Preorder(nodeFilter, func(node ast.Node) {
		binExpr, ok := node.(*ast.BinaryExpr)
		if !ok {
			return
		}

		xType, ok := pass.TypesInfo.Types[binExpr.X]
		if !ok {
			return
		}

		if !types.AssignableTo(xType.Type, errType) {
			return
		}

		yType, ok := pass.TypesInfo.Types[binExpr.Y]
		if !ok {
			return
		}

		if yType.IsNil() {
			return
		}

		pass.Reportf(node.Pos(), "do not compare error with \"==\" or \"!=\"")
	})
	return nil, nil
}
