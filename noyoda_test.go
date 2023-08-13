//nolint:testpackage
package noyoda

import (
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
