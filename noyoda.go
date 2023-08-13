package noyoda

import (
	"flag"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = `remove yoda conditions

Yoda condition is a expression/statement style to prevent accidental assignments like if x = 3 instead if x == 3.
Go does not needs this check.`

//nolint:gochecknoglobals
var (
	includeConst bool
	flagset      flag.FlagSet
)

func init() {
	flagset.BoolVar(&includeConst, "include-const", false, "should include const (default is false)")
}

func NewAnalyzer() *analysis.Analyzer {
	//nolint:exhaustruct,exhaustivestruct
	return &analysis.Analyzer{
		Name:     "noyoda",
		Doc:      doc,
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
		Flags:    flagset,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	//nolint:forcetypeassert
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
		(*ast.CaseClause)(nil),
	}

	ins.Preorder(nodeFilter, func(node ast.Node) {
		bexprs := parseBinaryExpressions(node)

		for _, bexpr := range bexprs {
			lval, ok := bexpr.X.(*ast.BasicLit)

			if !ok && !includeConst {
				continue
			}

			var lvalStr string
			if ok {
				lvalStr = lval.Value
			} else {
				n, ok := bexpr.X.(*ast.Ident)

				if !ok {
					continue
				}

				if n.Obj.Kind != ast.Con {
					continue
				}

				lvalStr = n.Name
			}

			rval, ok := bexpr.Y.(*ast.Ident)

			if !ok {
				continue
			}

			pass.Reportf(node.Pos(), "yoda condition: %s %s %s should be %s %s %s",
				lvalStr, bexpr.Op.String(), rval.Name,
				rval.Name, bexpr.Op.String(), lvalStr,
			)
		}
	})

	//nolint:nilnil
	return nil, nil
}

func parseBinaryExpressions(n ast.Node) []*ast.BinaryExpr {
	var bexprs []*ast.BinaryExpr

	switch node := n.(type) {
	case *ast.IfStmt:
		bexpr, ok := node.Cond.(*ast.BinaryExpr)

		if !ok {
			return bexprs
		}

		bexprs = append(bexprs, bexpr)

	case *ast.CaseClause:
		for _, v := range node.List {
			bexpr, ok := v.(*ast.BinaryExpr)

			if !ok {
				continue
			}

			bexprs = append(bexprs, bexpr)
		}
	default:
		panic("should never reach here, node is neither IfStmt nor CaseClause")
	}

	var ret []*ast.BinaryExpr

	for _, expr := range bexprs {
		exprs := recurseBinaryExpressions(expr)
		if exprs != nil {
			ret = append(ret, exprs...)
		}
	}

	return ret
}

func recurseBinaryExpressions(expr *ast.BinaryExpr) []*ast.BinaryExpr {
	xexpr, xok := expr.X.(*ast.BinaryExpr)
	yexpr, yok := expr.Y.(*ast.BinaryExpr)

	switch {
	case !xok && !yok: // expr does not contain another binary expression
		return []*ast.BinaryExpr{expr}
	case xok && yok: // both have binary expression
		xs := recurseBinaryExpressions(xexpr)
		ys := recurseBinaryExpressions(yexpr)

		return append(xs, ys...)

	case xok: // only x have binary expression
		return recurseBinaryExpressions(xexpr)
	case yok: // only y have binary expression
		return recurseBinaryExpressions(yexpr)

	default:
		panic("Unknown state")
	}
}
