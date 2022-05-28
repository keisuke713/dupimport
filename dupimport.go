package main

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "dupimport",
	Doc:  doc,
	Run:  run,
}

const doc = "dupimport finds duplicated imports in same file"

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		paths := map[string]bool{}
		for _, ip := range f.Imports {
			path, err := strconv.Unquote(ip.Path.Value)
			if err != nil {
				return nil, err
			}
			if paths[path] {
				pass.Reportf(ip.Pos(), "%s is duplicated import", path)
			} else {
				paths[path] = true
			}
		}
	}
	return nil, nil
}
