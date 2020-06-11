package domain

import (
	"errors"
	"path/filepath"
	"strconv"
	"strings"
)

type Path struct {
	Value string
}

func (p Path) elementsUptoN(n int) (Path, error) {
	elements := strings.Split(p.Value, "\\")
	uptoPath := ""
	if n <= 0 {
		return Path{Value: uptoPath}, errors.New("An argument is minus: " + strconv.Itoa(n))
	}
	if n > len(elements) {
		return Path{Value: uptoPath}, errors.New("An argument is bigger than path elements:" + p.Value + ", num:" + strconv.Itoa(n))
	}
	uptoPath = strings.Join(elements[:n], "\\")
	return Path{Value: uptoPath}, nil

}

func (p Path) nodeType() nodeType {
	if filepath.Ext(p.Value) == ".go" {
		return FileType
	}
	return DirType
}

func (p Path) last() string {
	ss := strings.Split(p.Value, "\\")
	return strings.TrimSpace(ss[len(ss)-1])
}

// contain return true if p contain path ( p: xxx/yyy/zzz, path xxx/yyy -> true)
func (p Path) contain(path Path) bool {
	return strings.Contains(p.Value, path.Value)
}

func (p Path) first() string {
	ss := strings.Split(p.Value, "\\")
	return ss[0]
}

func (p Path) Depth() int {
	return len(strings.Split(p.Value, "\\"))
}

func (p Path) IsTest() bool {
	return strings.Contains(strings.ToLower(p.Value), "test")
}

func (p Path) fullPath(projectRoot string) string {
	return filepath.Dir(projectRoot) + filepath.FromSlash("/") + p.Value
}
