//nolint:testpackage
package noyoda

import (
	"os"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

//nolint:gochecknoglobals
var testDir string

func init() {
	cur, _ := os.Getwd()
	testDir = cur + "/testdata/"
}

func TestIf(t *testing.T) {
	t.Parallel()

	analysistest.Run(t, testDir+"ifcond", NewAnalyzer())
}

func TestSwitch(t *testing.T) {
	t.Parallel()

	analysistest.Run(t, testDir+"switchcond", NewAnalyzer())
}

func TestConst(t *testing.T) {
	t.Parallel()

	a := NewAnalyzer()

	if err := a.Flags.Set("include-const", "true"); err != nil {
		panic(err)
	}

	analysistest.Run(t, testDir+"constcond", a)
}
