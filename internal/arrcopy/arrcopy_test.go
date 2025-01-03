package arrcopy_test

import (
	"testing"

	"github.com/m0t9/goperflint/internal/arrcopy"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestArrcopyAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), arrcopy.NewAnalyzer())
}
