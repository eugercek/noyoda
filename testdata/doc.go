/*
	Package testdata have test go files which is used by golang.org/x/tools/go/analysis/analysistest package

Generally first yoda conditions tested afterwards non yoda conditions tested (to not create wrong diagnostics)

10, 20, 30 ... -> yoda condition
100, 200, 300 ... -> no yoda condition

In recursive tests we check 4 condition (n = noyoda y = yoda)

	n y
	y n
	n n
	y y

In case test we check

	y
	y,y
	n
	n,n
*/
package testdata
