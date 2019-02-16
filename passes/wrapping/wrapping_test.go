package wrapping_test

import (
	"testing"

	"github.com/sachaos/errchk/passes/wrapping"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, wrapping.Analyzer, "a")
}