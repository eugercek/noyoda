//nolint:testpackage
package noyoda

import (
	"fmt"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestIf(t *testing.T) {
	t.Parallel()

	analysistest.Run(t, analysistest.TestData(), NewAnalyzer(), "ifcond")
}

func TestSwitch(t *testing.T) {
	t.Parallel()

	analysistest.Run(t, analysistest.TestData(), NewAnalyzer(), "switchcond")
}

func TestConst(t *testing.T) {
	t.Parallel()

	a := NewAnalyzer()

	if err := a.Flags.Set("include-const", "true"); err != nil {
		panic(err)
	}

	analysistest.Run(t, analysistest.TestData(), a, "constcond")
}

func TestRange(t *testing.T) {
	t.Parallel()

	a := NewAnalyzer()

	if err := a.Flags.Set("skip-range", "true"); err != nil {
		panic(err)
	}

	analysistest.Run(t, analysistest.TestData(), a, "rangecond")
}

func Test_rangeOperatorMatch(t *testing.T) {
	t.Parallel()

	tests := []struct {
		l, r string
		ok   bool
	}{
		{">", ">", true},
		{">", ">=", true},
		{">=", ">", true},
		{">=", ">=", true},
		//
		{"<", "<", true},
		{"<", "<=", true},
		{"<=", "<", true},
		{"<=", "<=", true},
		//
		{">", "<", false},
		{">", "<=", false},
		{">=", "<", false},
		{">=", "<=", false},
		//
		{"<", ">", false},
		{"<", ">=", false},
		{"<=", ">", false},
		{"<=", ">=", false},
	}

	for _, ts := range tests {
		ts := ts
		t.Run(fmt.Sprintf("%s and %s", ts.l, ts.r), func(t *testing.T) {
			t.Parallel()

			ok := rangeOperatorMatch(ts.l, ts.r)

			if ok != ts.ok {
				t.Errorf("want %t got %t", ts.ok, ok)
			}
		})
	}
}
