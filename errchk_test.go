package errchk_test

import (
	"testing"

	"github.com/sachaos/errchk"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, errchk.Analyzer, "a")
}