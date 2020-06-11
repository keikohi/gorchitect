package dot

import "github.com/keikohi/gorchitect/domain"

type GraphConverter struct {
	rootNode domain.Node
}

func NewGraphConverter(root domain.Node) GraphConverter {
	return GraphConverter{rootNode: root}
}

func (gf GraphConverter) Convert() Graph {
	root := gf.toGraph(gf.rootNode)
	return graph{
		label:     (root).Label(),
		path:      (root).Path(),
		child:     (root).Childs(),
		graphtype: Digraph}
}

func (gf GraphConverter) focusPackage(chain domain.Chain) Graph {
	root := gf.toGraphByChain(gf.rootNode, chain)
	return graph{
		label:     (root).Label(),
		path:      (root).Path(),
		child:     (root).Childs(),
		graphtype: Digraph}
}

func (gf GraphConverter) toGraph(node domain.Node) Graph {

	if node.IsDir() {
		childs := []Graph{}
		for _, child := range node.GetChileds() {
			childs = append(childs, gf.toGraph(child))
		}
		return subgraph{
			graph{
				label:     node.GetName(),
				path:      node.GetPath(),
				child:     childs,
				graphtype: Subgraph}}
	}
	return item{
		graph{
			label:     node.GetName(),
			path:      node.GetPath(),
			child:     nil,
			graphtype: Item,
		}}
}

func (gf GraphConverter) toGraphByChain(node domain.Node, chain domain.Chain) Graph {

	if !chain.Contain(node.GetPath()) {
		return nil
	}

	if node.IsDir() {
		childs := []Graph{}
		for _, child := range node.GetChileds() {
			graph := gf.toGraphByChain(child, chain)
			if graph != nil {
				childs = append(childs, graph)
			}
		}
		return subgraph{
			graph{
				label:     node.GetName(),
				path:      node.GetPath(),
				child:     childs,
				graphtype: Subgraph}}
	}
	return item{
		graph{
			label:     node.GetName(),
			path:      node.GetPath(),
			child:     nil,
			graphtype: Item,
		}}
}
