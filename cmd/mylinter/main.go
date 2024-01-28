package main

import (
	"mylinter"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(mylinter.Analyzer) }
