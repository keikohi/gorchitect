package writer

import (
	"fmt"
	"io"

	"github.com/keikohi/gorchitect/domain"
	"github.com/keikohi/gorchitect/domain/dot"
)

// Dotwriter write go file dependencies in dot language.
type Dotwriter struct {
	W io.Writer
}

// FileAttribute is file attribute options
const FileAttribute string = "penwidth=0.6, color=\"#395070\", style=filled, fontcolor=white, fontsize=18, size=30, fontname=meiryo"

func (dw Dotwriter) Write(rootNode domain.Node, dr *domain.DependecyRelation) {
	dotConverter := dot.NewGraphConverter(rootNode)
	rootDotGraph := dotConverter.Convert()
	dw.writeRecursively(rootDotGraph, dr)
}

// Write write project dependencies to a file
func (dw Dotwriter) writeRecursively(graph dot.Graph, dr *domain.DependecyRelation) {
	// filtering dot.Graphs has a path include "test" ( regardless of Upper or Lower "test"
	if graph.IsTestPath() {
		return
	}
	if graph.GraphType() == dot.Digraph {
		dw.digraph(graph, "digraph")
		for _, child := range graph.Childs() {
			dw.writeRecursively(child, dr)
		}
		dw.dependencies(dr)
		dw.end(graph)
	}
	if graph.GraphType() == dot.Subgraph {
		dw.subgraph(graph, "subgraph")
		for _, child := range graph.Childs() {
			dw.writeRecursively(child, dr)
		}
		dw.end(graph)
	}
	if graph.GraphType() == dot.Item {
		dw.file(graph)
	}
}

func (dw Dotwriter) digraph(graph dot.Graph, graphType string) {
	fmt.Fprintln(dw.W, dw.space(graph.Depth()), graphType, "\"", graph.Label(), "\"", " {")
	// TODO
	space := dw.space(graph.Depth() + 1)
	fmt.Fprintln(dw.W, space, "rankdir=LR")
	fmt.Fprintln(dw.W, space, "bgcolor =\"#061A2B\"")
	fmt.Fprintln(dw.W, space, "fontcolor = white")
	fmt.Fprintln(dw.W, space, "fontname = meiryo")
	fmt.Fprintln(dw.W, space, "fontsize =20")
	fmt.Fprintln(dw.W, space, "ranksep = 4")
	fmt.Fprintln(dw.W, space, "nodesep = 0.2")

}

func (dw Dotwriter) subgraph(graph dot.Graph, graphType string) {
	fmt.Fprintln(dw.W, dw.space(graph.Depth()), graphType, "\"cluster_"+graph.Label()+"\" {")
}

// "archive\tar\example_test.go" [label="example_test.go"]
func (dw Dotwriter) file(graph dot.Graph) {
	fmt.Fprintln(dw.W, dw.space(graph.Depth()), "\""+graph.PathVal()+"\""+" [label=\""+graph.Label()+"\","+FileAttribute+"]")
}

// [label="xxx",color="#3B9FED"]
func (dw Dotwriter) end(graph dot.Graph) {
	fmt.Fprintln(dw.W, dw.space(graph.Depth()+1), "label="+"\""+graph.Label()+"\" color=\"#3B9FED\"")
	fmt.Fprintln(dw.W, dw.space(graph.Depth()), "}")
}

func (dw Dotwriter) dependencies(dr *domain.DependecyRelation) {
	for _, dependency := range *dr {
		client := dependency.Consumer
		for _, vendors := range dependency.Vendors {
			fmt.Fprintf(dw.W, "    \"%s\" -> \"%s\" [color=white, penwidth=0.4]\n", client.Filename, vendors.Filename)
		}
	}
}

func (dw Dotwriter) space(depth int) string {
	space := ""
	for i := 0; i < depth; i++ {
		space += "  "
	}
	return space
}
