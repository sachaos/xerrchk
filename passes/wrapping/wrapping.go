package wrapping

import (
	"fmt"
	"go/token"
	"go/types"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"

	"golang.org/x/tools/go/analysis"
)

var scope string
var errType = types.Universe.Lookup("error").Type().Underlying().(*types.Interface)

var Analyzer = &analysis.Analyzer{
	Name: "wrapping",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

const Doc = "wrapping detect unwrapped error"
const xerrorsPath = "golang.org/x/xerrors"
const scopeAll = "all"
const scopePublic = "public"

func init() {
	Analyzer.Flags.StringVar(&scope, "scope", scopeAll, "scope of checking (public, all)")
}

func run(pass *analysis.Pass) (interface{}, error) {
	err := checkFlags()
	if err != nil {
		return nil, err
	}

	srcFuncs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, srcFunc := range srcFuncs {
		positions := wrappingErrPositions(srcFunc)
		for _, pos := range positions {
			pass.Reportf(pos, "wrap with xerrros.Errorf or xerrors.Opaque")
		}
	}

	return nil, nil
}

func checkFlags() error {
	if scope == scopeAll || scope == scopePublic {
		return nil
	}
	return fmt.Errorf("unknown scope '%s'", scope)
}

func wrappingErrPositions(srcFunc *ssa.Function) []token.Pos {
	if scope == scopePublic && isPrivate(srcFunc) {
		return nil
	}

	if !isReturningErr(srcFunc) {
		return nil
	}

	var positions []token.Pos
	for _, block := range srcFunc.Blocks {
		for _, instr := range block.Instrs {
			val, ok := instr.(ssa.Value)
			if !ok {
				continue
			}

			if isErr(val.Type()) && !isCallingXerrors(val) && isReachToReturn(val) {
				positions = append(positions, convertToOriginVal(val).Pos())
			}
		}
	}

	return positions
}

func isPrivate(function *ssa.Function) bool {
	if function.Parent() != nil {
		return isPrivate(function.Parent())
	}

	name := function.Name()
	return unicode.IsLower(rune(name[0]))
}

func isCallingXerrors(val ssa.Value) bool {
	call, ok := val.(*ssa.Call)
	if !ok {
		return false
	}

	f := call.Common().StaticCallee()
	if f == nil {
		return false
	}

	return removeVendor(f.Pkg.Pkg.Path()) == xerrorsPath
}

func removeVendor(path string) string {
	s := strings.Split(path, "/")
	for i := range s {
		if s[i] == "vendor" {
			return strings.Join(s[i+1:], "/")
		}
	}
	return path
}

func isReachToReturn(val ssa.Value) bool {
	for _, ref := range *val.Referrers() {
		_, ok := ref.(*ssa.Return)
		if ok {
			return true
		}
	}
	return false
}

func isReturningErr(srcFunc *ssa.Function) bool {
	results := srcFunc.Signature.Results()
	for i := 0; i < results.Len(); i++ {
		v := results.At(i)
		if isErr(v.Type()) {
			return true
		}
	}
	return false
}

func convertToOriginVal(val ssa.Value) ssa.Value {
	switch v := val.(type) {
	case *ssa.Extract:
		return v.Tuple
	case *ssa.MakeInterface:
		if val.Pos() != 0 {
			return val
		}
		return v.X
	}
	return val
}

func isErr(v types.Type) bool {
	return types.AssignableTo(v, errType)
}
