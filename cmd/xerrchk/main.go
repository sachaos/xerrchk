package main

import (
	"github.com/sachaos/xerrchk/passes/erreq"
	"github.com/sachaos/xerrchk/passes/wrapping"
	"github.com/tenntenn/gosa/passes/nilerr"
	"github.com/tenntenn/gosa/passes/wraperrfmt"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		wrapping.Analyzer,
		erreq.Analyzer,
		wraperrfmt.Analyzer,
		nilerr.Analyzer,
	)
}
