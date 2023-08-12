//nolint:testpackage
package noyoda

import (
	"os"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	t.Parallel()

	cur, _ := os.Getwd()
	testDir := cur + "/testdata"

	analysistest.Run(t, testDir, NewAnalyzer())
}
