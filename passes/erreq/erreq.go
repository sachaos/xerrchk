package erreq

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "erreq",
	Doc:  "erreq finds binary ops and switch-case statements comparing errors",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var errType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

var unwrapFuncs = []string{
	"golang.org/x/xerrors.Unwrap",
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	checkBinaryOps(pass, inspect)

	checkSwitchStmt(pass, inspect)

	return nil, nil
}

// Check boolean expression using == or !=.
func checkBinaryOps(pass *analysis.Pass, inspect *inspector.Inspector) {
	nodeFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(node ast.Node) {
		binExpr, ok := node.(*ast.BinaryExpr)

		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println(binExpr)

		if !ok {
			return
		}

		if !(binExpr.Op == token.EQL || binExpr.Op == token.NEQ) {
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
}

// Check switch-case statement.
func checkSwitchStmt(pass *analysis.Pass, inspect *inspector.Inspector) {
	nodeFilter := []ast.Node{
		(*ast.SwitchStmt)(nil),
	}
	inspect.Preorder(nodeFilter, func(node ast.Node) {

		switchStmt, ok := node.(*ast.SwitchStmt)

		pass.Reportf(node.Pos(), "do not use wrapped errors as a tag of switch statement.")

		if !ok {
			return
		}

		tagType, ok := pass.TypesInfo.Types[switchStmt.Tag]
		if !ok {
			return
		}

		if !types.AssignableTo(tagType.Type, errType) {
			return
		}

		// if the tag of switch stmt is unwrapped, it should be allowed.
		callExp, ok := switchStmt.Tag.(*ast.CallExpr)
		if ok {
			if expIsUnwrapFunc(callExp.Fun) {
				return
			}
		}

		pass.Reportf(node.Pos(), "do not use wrapped errors as a tag of switch statement.")
	})
}

func expIsUnwrapFunc(exp ast.Expr) bool {
	funIdent, ok := exp.(*ast.Ident)

	fmt.Println(funIdent)

	if !ok {
		return false
	}

	for _, funcName := range unwrapFuncs {
		if funIdent.Name == funcName {
			return true
		}
	}

	return false
}
