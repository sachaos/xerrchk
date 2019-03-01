package erreq_test

import (
	"testing"

	"github.com/sachaos/xerrchk/passes/erreq"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, erreq.Analyzer, "a")
}
