package wrapping_test

import (
	"testing"

	"github.com/sachaos/xerrchk/passes/wrapping"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestWrapping(t *testing.T) {
	testdata := analysistest.TestData()
	wrapping.Analyzer.Flags.Set("scope", "public")
	analysistest.Run(t, testdata, wrapping.Analyzer, "a")
}

func TestPublicOption(t *testing.T) {
	testdata := analysistest.TestData()
	wrapping.Analyzer.Flags.Set("scope", "public")
	analysistest.Run(t, testdata, wrapping.Analyzer, "a/option/public")
}

func TestAllOption(t *testing.T) {
	testdata := analysistest.TestData()
	wrapping.Analyzer.Flags.Set("scope", "all")
	analysistest.Run(t, testdata, wrapping.Analyzer, "a/option/all")
}
