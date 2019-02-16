package main

import (
	"github.com/sachaos/xerrchk/passes/isas"
	"github.com/sachaos/xerrchk/passes/wrapping"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		wrapping.Analyzer,
		isas.Analyzer,
	)
}
