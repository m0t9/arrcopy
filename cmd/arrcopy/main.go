package main

import (
	"github.com/m0t9/goperflint/internal/arrcopy"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(arrcopy.NewAnalyzer())
}
