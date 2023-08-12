package main

import (
	"github.com/eugercek/noyoda"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(noyoda.NewAnalyzer())
}
