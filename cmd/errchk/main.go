package main

import (
	"github.com/sachaos/errchk/passes/wrapping"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(wrapping.Analyzer) }
