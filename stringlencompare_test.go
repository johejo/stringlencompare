package stringlencompare_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/johejo/stringlencompare"
)

func Test(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), stringlencompare.Analyzer, "a")
}
