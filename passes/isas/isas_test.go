package isas_test

import (
	"testing"

	"github.com/sachaos/errchk/passes/isas"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, isas.Analyzer, "a")
}