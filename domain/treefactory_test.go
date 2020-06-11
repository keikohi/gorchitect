package domain

import (
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// gomeba
// └─test
//     └─example
//         │  example.go
//         │
//         └─library
//             │  external.go
//             │  internal.go
//             │
//             └─empty

func TestTreeFactory(t *testing.T) {
	currentDir, _ := os.Getwd()
	srcDir := filepath.Dir(currentDir)

	filepaths, _ := LoadGoFiles(filepath.Join(srcDir, "test", "example"))
	nf := NewTreeFactory(filepaths)
	example := nf.Build()

	if example.GetName() != "example" {
		t.Fatalf("Node is not example: %v", example.GetName())
	}

	exampleChilds := example.GetChileds()
	if len(exampleChilds) != 2 {
		t.Fatalf("exampleChilds num is wrong: %v", len(exampleChilds))
	}

	for _, exampleChild := range exampleChilds {
		if exampleChild.IsDir() {
			if exampleChild.GetName() != "library" {
				t.Fatalf("Node is not library: %v", exampleChild.GetName())
			}

			library := exampleChild.GetChileds()
			sort.Slice(library, func(i, j int) bool {
				return library[i].GetName() < library[j].GetName()
			})

			expects := []string{"const.go", "external.go", "func.go", "internal.go", "type.go"}
			actuals, ok := containAll(library, expects)
			if !ok {
				t.Logf("actuals: %+v", actuals)
				t.Fatalf("expexts: %+v", expects)
			}
		} else {
			if exampleChild.GetName() != "example.go" {
				t.Fatalf("Node is not example.go: %v", exampleChild.GetName())
			}
		}
	}

}

func containAll(nodes []Node, expects []string) ([]string, bool) {
	nodeNames := make([]string, 0, len(nodes))
	for _, node := range nodes {
		nodeNames = append(nodeNames, node.GetName())
	}
	sort.Slice(nodeNames, func(i, j int) bool {
		return nodeNames[i] < nodeNames[j]
	})

	sort.Slice(expects, func(i, j int) bool {
		return expects[i] < expects[j]
	})

	for i := range expects {
		if nodeNames[i] != expects[i] {
			return nodeNames, false
		}
	}
	return nodeNames, true
}
