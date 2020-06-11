package domain

import (
	"strings"

	"github.com/keikohi/gorchitect/util"
)

type dependencyMap map[string][]string

func newDependencyMap() dependencyMap {
	return make(dependencyMap)
}

func (sm dependencyMap) isEmpty(x string) bool {
	// value, isExist
	_, ok := sm[x]
	return !ok
}

func (sm dependencyMap) add(x string, selector string) {
	if sm.isEmpty(x) {
		var selectors []string
		selectors = append(selectors, selector)
		sm[x] = selectors
	} else {
		sm[x] = append(sm[x], selector)
	}
}

func (sm dependencyMap) distinct() dependencyMap {
	newSelectormap := newDependencyMap()
	for x, selectors := range sm {
		distinctmap := make(map[string]bool)
		var newSelectors []string
		for _, selector := range selectors {
			if _, ok := distinctmap[selector]; !ok {
				distinctmap[selector] = true
				newSelectors = append(newSelectors, selector)
			}
		}
		newSelectormap[x] = newSelectors
	}
	return newSelectormap
}

func (sm dependencyMap) outerImports(imports imprts) dependencyMap {
	newSelectormap := newDependencyMap()
	for x, selectors := range sm {
		for _, imprt := range imports.decls {
			//imprt is slashed path (xxx/xxx/xxx/)
			lastDir := lastItem(strings.Split(string(imprt), "/"))
			if x == lastDir {
				newSelectormap[x] = selectors
			}
		}
	}
	return newSelectormap
}

func (sm dependencyMap) typemap(funcmap dependencyMap) dependencyMap {
	newSelectorMap := newDependencyMap()
	for x, functypes := range sm {
		if funcs, ok := funcmap[x]; ok {
			var types []string
			for _, functype := range functypes {
				if isType(funcs, functype) {
					types = append(types, functype)
				}
			}
			if len(types) > 0 {
				newSelectorMap[x] = types
			}
		} else {
			// functypes are type!
			newSelectorMap[x] = functypes
		}
	}
	return newSelectorMap
}

func lastItem(ss []string) string {
	return ss[len(ss)-1]
}

func isType(funcs []string, functype string) bool {
	return !util.Contains(funcs, functype)
}
