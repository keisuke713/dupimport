package main

import (
	"github.com/keisuke713/dupimport"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(dupimport.Analyzer)
}
