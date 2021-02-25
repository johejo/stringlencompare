package main

import (
	"github.com/johejo/stringlencompare"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(stringlencompare.Analyzer) }
