package main

import (
	"github.com/sachaos/errchk"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(errchk.Analyzer) }