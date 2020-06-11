package domain

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDependencyRelation(t *testing.T) {
	currentDir, _ := os.Getwd()
	projectPath := filepath.Join(filepath.Dir(currentDir), "test", "example")
	filepaths, _ := LoadGoFiles(projectPath)
	analyzer := NewAnalyzer(filepaths, projectPath)
	codes := analyzer.Analyse()
	dr := codes.DependecyRelation()

	for _, dep := range dr {
		exampleC := filepath.Join("example", "example.go")
		if dep.Consumer.Filename == exampleC {
			vendors := dep.Vendors
			externalGO := filepath.Join("example", "library", "external.go")
			if len(vendors) != 1 {
				t.Errorf("exmaple/example.g's vendor num is wrong: %v", len(vendors))
			}
			if vendors[0].Filename != externalGO {
				t.Errorf("no vendor: %v, %v", vendors[0].Filename, externalGO)
			}
		}
	}

}
