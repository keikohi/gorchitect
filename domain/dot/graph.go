package dot

import "github.com/keikohi/gorchitect/domain"

// Graph coresponds to rootNode
type Graph interface {
	Label() string
	Depth() int
	Childs() []Graph
	Path() domain.Path
	PathVal() string
	GraphType() GraphType
	IsTestPath() bool
}

type graph struct {
	label     string
	path      domain.Path
	child     []Graph
	graphtype GraphType
}

func (g graph) Label() string        { return g.label }
func (g graph) Childs() []Graph      { return g.child }
func (g graph) Path() domain.Path    { return g.path }
func (g graph) PathVal() string      { return g.Path().Value }
func (g graph) GraphType() GraphType { return g.graphtype }
func (g graph) Depth() int           { return g.path.Depth() }
func (g graph) IsTestPath() bool     { return g.path.IsTest() }

//digraph coresponds to dir
type digraph struct {
	graph
}

func newDigraph(label string) Graph {
	return digraph{
		graph{
			label:     label,
			path:      domain.Path{},
			child:     []Graph{},
			graphtype: Digraph}}
}

type subgraph struct {
	graph
}

func newSubgraph(label string, depth int, child []Graph) Graph {
	return subgraph{
		graph{
			label:     label,
			path:      domain.Path{},
			child:     []Graph{},
			graphtype: Subgraph}}
}

//digraph coresponds to dir
type item struct {
	graph
}

func newItem(label string, child []Graph) Graph {
	return item{
		graph{
			label:     label,
			path:      domain.Path{},
			child:     []Graph{},
			graphtype: Subgraph}}
}

// GraphType is a type of graph in dot languages
type GraphType string

// Digraph is a type of digraph
const Digraph GraphType = "digraph"

// Subgraph is a type of subgraph
const Subgraph GraphType = "subgraph"

// Item is a type of item
const Item GraphType = "item"
