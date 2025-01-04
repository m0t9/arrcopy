package arrcopy_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/m0t9/goperflint/pkg/arrcopy"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestArrcopyAnalyzer(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, arrcopy.NewAnalyzer(), "arrcopy")
}
