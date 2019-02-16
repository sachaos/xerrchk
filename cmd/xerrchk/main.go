package main

import (
	"github.com/sachaos/xerrchk/passes/isas"
	"github.com/sachaos/xerrchk/passes/wrapping"
	"github.com/tenntenn/gosa/passes/nilerr"
	"github.com/tenntenn/gosa/passes/wraperrfmt"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		wrapping.Analyzer,
		isas.Analyzer,
		wraperrfmt.Analyzer,
		nilerr.Analyzer,
	)
}
