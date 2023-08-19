package noyoda

import (
	"flag"
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = `remove yoda conditions

Yoda condition is a expression/statement style to prevent accidental assignments like if x = 3 instead if x == 3.
Go does not needs this check.`

var (
	includeConst bool
	skipRange    bool
	flagset      flag.FlagSet
)

func init() {
	flagset.BoolVar(&includeConst, "include-const", false, "should include const (default is false)")
	flagset.BoolVar(&skipRange, "skip-range", true,
		"should skip (10 < a && a < 20) like range conditions from yoda conditions (default is false)")
}

func NewAnalyzer() *analysis.Analyzer {
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
			lval, rval, ok := isYodaCondition(bexpr)

			if !ok {
				continue
			}

			newText := fmt.Sprintf("%s %s %s", rval, bexpr.Op.String(), lval)
			errorMsg := fmt.Sprintf("yoda condition: %s %s %s should be %s",
				lval, bexpr.Op.String(), rval,
				newText,
			)

			pass.Report(analysis.Diagnostic{
				Pos:      bexpr.Pos(),
				End:      bexpr.End(),
				Category: "noyoda",
				Message:  errorMsg,
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: errorMsg,
						TextEdits: []analysis.TextEdit{
							{
								Pos:     bexpr.Pos(),
								End:     bexpr.End(),
								NewText: []byte(newText),
							},
						},
					},
				},
			})
		}
	})

	//nolint:nilnil
	return nil, nil
}

// isYodaCondition checks if an expression written in yoda condition style
// many expression may seem like "yoda condition" style for example: (2 * time.Minute)
// which is very natural form of saying "2 minutes", since it's not in an if or switch-case
// "statement" this is not yoda condition.
//
// # Usage of isYodaCondition should be on ast.IfStmt.Cond and ast.CaseClause
//
// This function would be very good use case of naked return :(.
func isYodaCondition(expr *ast.BinaryExpr) (lval, rval string, ok bool) {
	lval, ok = parseYodaConstant(expr.X)

	if !ok {
		return "", "", false
	}

	r, ok := expr.Y.(*ast.Ident)

	if !ok {
		return "", "", false
	}

	rval = r.Name

	return lval, rval, true
}

// parseBinaryExpressions extracts binary expressions from a node.
// Recursively looks for all binary expressions.
//
// For functionality of this package, only ast.IfStmt and ast.caseClause is enough
// Any other node as argument will panic.
func parseBinaryExpressions(n ast.Node) []*ast.BinaryExpr {
	var bexprs []*ast.BinaryExpr

	switch node := n.(type) {
	case *ast.IfStmt:
		bexpr, ok := node.Cond.(*ast.BinaryExpr)

		if !ok {
			return nil
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

// recurseBinaryExpressions traverses and accumulates all binary expressions under given expr argument.
func recurseBinaryExpressions(expr *ast.BinaryExpr) []*ast.BinaryExpr {
	xexpr, xok := expr.X.(*ast.BinaryExpr)
	yexpr, yok := expr.Y.(*ast.BinaryExpr)

	switch {
	case !xok && !yok: // expr does not contain another binary expression
		return []*ast.BinaryExpr{expr}
	case xok && yok: // both have binary expression
		// Check 10 > a && a > 5 like conditions
		// This must be above from recursion because we need a root bexpr
		// and look at its children (if any)
		if skipRange && isNumberRangExpression(expr) {
			return []*ast.BinaryExpr{}
		}

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

// parseYodaConstant checks if the expression is valid yoda expression's left hand
// Which is only valid when left is basic literal or `includeConst` enabled and left is constant.
func parseYodaConstant(e ast.Expr) (val string, ok bool) {
	switch expr := e.(type) {
	case *ast.BasicLit:
		return expr.Value, true
	case *ast.Ident:
		if !includeConst {
			return "", false
		}

		if expr.Obj.Kind != ast.Con {
			return "", false
		}

		return expr.Name, true
	default:
		return "", false
	}
}

// isNumberRangExpression checks if a binary expression is a "number range expression"
// a number range expression is something like: 10 > a && a > 5 or 1 < a && a < 20
// since this is very natural way of saying: 10 > a > 5 and 1 < a 20, noyoda
// ignores this type of yoda condition (10 > a)
//
// This function is limit on the cyclomatic complexity, carefully edit.
func isNumberRangExpression(top *ast.BinaryExpr) bool {
	left, ok := top.X.(*ast.BinaryExpr)

	if !ok {
		return false
	}

	right, ok := top.Y.(*ast.BinaryExpr)

	if !ok {
		return false
	}

	if top.Op.String() != "&&" {
		return false
	}

	ok = numberRangeOperatorMatch(left.Op.String(), right.Op.String())

	if !ok {
		return false
	}

	// Here we know that expression is like a < b  && c < d
	// Now we need to check if `a` and `d `are yoda constant
	// And b and c are same variable

	_, llok := parseYodaConstant(left.X)
	lr, lrok := left.Y.(*ast.Ident)

	rl, rlok := right.X.(*ast.Ident)
	_, rrok := parseYodaConstant(right.Y)

	if !(llok && lrok && rlok && rrok) {
		return false
	}

	if lr.Name != rl.Name {
		return false
	}

	return true
}

// numberRangeOperatorMatch checks can l and r operators can be part of
// "number range expression". look at isNumberRangExpression's docstring for explanation.
func numberRangeOperatorMatch(l, r string) bool {
	switch l {
	case ">", ">=":
		return r == ">" || r == ">="
	case "<", "<=":
		return r == "<" || r == "<="
	}

	return false
}
