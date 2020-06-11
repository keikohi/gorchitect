package domain

import (
	"log"
)

type nodeType string

// FileType is file path type i.e) xxx/xxx/xxx.go
const FileType nodeType = nodeType("file")

// DirType is dir path type. i.e) xxx/xxx/xxx
const DirType nodeType = nodeType("dir")

type Node interface {
	GetName() string
	IsDir() bool
	GetDepth() int
	GetChileds() []Node
	GetPath() Path
	getParent() Node
	isSamePath(Path) bool
	addChildren(Node)
	hasChilds() bool
}

type file struct {
	path   Path   //fullpath: /xxx/xxx/xxx/xx.go
	name   string //xx.go or xxx
	parent Node
	childs []Node
	depth  int // Level of hierarchy
}

func (file file) getParent() Node { return file.parent }

func (file file) GetChileds() []Node { return file.childs }

func (file file) isSamePath(path Path) bool { return file.path.Value == path.Value }

func (file *file) addChildren(node Node) {
	log.Fatal("Cannot add node to file nodes.", node.GetPath().Value)
}

func (file file) GetName() string { return file.name }

func (file file) hasChilds() bool { return false }

func (file file) IsDir() bool { return false }

func (file file) GetDepth() int { return file.depth }

func (file file) GetPath() Path { return file.path }

type Dir struct {
	path   Path
	name   string
	parent Node
	childs []Node
	depth  int // Level of hierarchy
}

func (dir Dir) getParent() Node { return dir.parent }

func (dir Dir) GetChileds() []Node { return dir.childs }

func (dir Dir) isSamePath(path Path) bool { return dir.path.Value == path.Value }

func (dir *Dir) addChildren(node Node) { dir.childs = append(dir.childs, node) }

func (dir Dir) GetName() string { return dir.name }

func (dir Dir) hasChilds() bool { return len(dir.GetChileds()) > 0 }

func (dir Dir) IsDir() bool { return true }

func (dir Dir) GetDepth() int { return dir.depth }

func (dir Dir) GetPath() Path { return dir.path }
