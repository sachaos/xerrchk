package isas_test

import (
	"testing"

	"github.com/sachaos/xerrchk/passes/isas"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, isas.Analyzer, "a")
}
