package main

import (
	"github.com/sachaos/errchk/passes/isas"
	"github.com/sachaos/errchk/passes/wrapping"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		wrapping.Analyzer,
		isas.Analyzer,
	)
}
