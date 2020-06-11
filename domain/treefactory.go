package domain

import (
	"fmt"
)

func newRootNode(rootName string) *Dir {
	path := Path{rootName}
	return &Dir{
		path:   path,
		name:   rootName,
		parent: nil,
		childs: []Node{},
		depth:  path.Depth()}
}

type treefactory struct {
	filenames []Path
	RootNode  *Dir
	nodeNum   int
}

func NewTreeFactory(filenames []string) treefactory {
	paths := toPaths(filenames)
	return treefactory{paths, newRootNode(paths[0].first()), 1}
}

func toPaths(fullnames []string) []Path {
	var paths []Path
	for _, p := range fullnames {
		paths = append(paths, Path{p})
	}
	return paths
}

func (nf treefactory) Build() Node {
	for i, path := range nf.filenames {
		if i == 0 {
			nf.createNode(nf.RootNode, path, 2)
		} else {
			nf.createNode(nf.RootNode, path, 1)
		}
	}
	return nf.RootNode
}

func newNode(parent Node, pth Path) Node {
	if pth.nodeType() == FileType {
		return &file{
			path:   pth,
			name:   pth.last(),
			parent: parent,
			childs: []Node{},
			depth:  pth.Depth()}
	}
	return &Dir{
		path:   pth,
		name:   pth.last(),
		parent: parent,
		childs: []Node{},
		depth:  pth.Depth()}
}

//DFS
func (nf treefactory) createNode(node Node, filepath Path, i int) {
	if curPath, err := filepath.elementsUptoN(i); err == nil {
		if !node.isSamePath(curPath) {
			depth := node.GetDepth()
			if depth < curPath.Depth() {
				child := newNode(node, curPath)
				nf.nodeNum++
				node.addChildren(child)
				nf.createNode(child, filepath, i+1)
				return
			}
			if depth == node.GetDepth() {
				parent := node.getParent()
				child := newNode(parent, curPath)
				nf.nodeNum++
				parent.addChildren(child)
				nf.createNode(child, filepath, i+1)
				return
			}
		} else {
			if node.hasChilds() {
				for _, child := range node.GetChileds() {
					if nextPath, err := filepath.elementsUptoN(i + 1); err == nil {
						if child.isSamePath(nextPath) {
							nf.createNode(child, filepath, i+1)
							return
						}
					}
				}
				nf.createNode(node, filepath, i+1)
			} else {
				nf.createNode(node, filepath, i+1)
			}
		}
	}
}

func (nf treefactory) printNodeNameRecursively(node Node) {
	fmt.Println(node.GetPath().Value, " : ", node.GetName())
	if len(node.GetChileds()) == 0 {
		return
	}
	for _, child := range node.GetChileds() {
		nf.printNodeNameRecursively(child)
	}
}
