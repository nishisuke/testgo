package main

import (
	mylinter "github.com/nishisuke/testgo"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(mylinter.Analyzer) }
