package testgo

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "mylinter is ..."

var (
	s        string
	Analyzer = &analysis.Analyzer{
		Name: "mylinter",
		Doc:  doc,
		Run:  run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}
)

func init() {
	Analyzer.Flags.StringVar(&s, "mylinter.s", "", "s default false")
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ImportSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ImportSpec:
			if n.Name == nil {
				return
			}
			if !strings.Contains(n.Path.Value, s) {
				return
			}
			pass.Reportf(n.Pos(), "%s is imported as %s", n.Path.Value, n.Name.Name)
		}
	})

	return nil, nil
}
