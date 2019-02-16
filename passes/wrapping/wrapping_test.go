package wrapping_test

import (
	"testing"

	"github.com/sachaos/errchk/passes/wrapping"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestPublicOption(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, wrapping.Analyzer, "a/option/public")
}

func TestAllOption(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, wrapping.Analyzer, "a/option/all")
}
