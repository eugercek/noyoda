package noyoda

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = `remove yoda conditions

Yoda condition is a expression/statement style to prevent accidental assignments like if x = 3 instead if x == 3.
Go does not needs this check.
`

func NewAnalyzer() *analysis.Analyzer {
	//nolint:exhaustruct,exhaustivestruct
	return &analysis.Analyzer{
		Name:     "noyoda",
		Doc:      doc,
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	//nolint:forcetypeassert
	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.IfStmt)(nil),
	}

	ins.Preorder(nodeFilter, func(node ast.Node) {
		//nolint:varnamelen
		stmt, ok := node.(*ast.IfStmt)

		if !ok {
			return
		}

		bexpr, ok := stmt.Cond.(*ast.BinaryExpr)

		if !ok {
			return
		}

		lval, ok := bexpr.X.(*ast.BasicLit)

		if !ok {
			return
		}

		rval, ok := bexpr.Y.(*ast.Ident)

		if !ok {
			return
		}

		pass.Reportf(node.Pos(), "yoda condition: %s %s %s should be %s %s %s",
			lval.Value, bexpr.Op.String(), rval.Name,
			rval.Name, bexpr.Op.String(), lval.Value,
		)
	})

	//nolint:nilnil
	return nil, nil
}
