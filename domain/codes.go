package domain

import "github.com/keikohi/gorchitect/util"

// code has declarations(type, methods...) and dependencyMap()
type code struct {
	decl    declaration
	depmaps dependencyMaps
}

func (c code) getPackageName() string {
	return c.decl.packageName
}

// Codes is a array of gofile
type Codes []code

func (cs *Codes) add(decl declaration, depmaps dependencyMaps) {
	gofile := code{decl, depmaps}
	*cs = append(*cs, gofile)
}

func (cs Codes) DependecyRelation() DependecyRelation {
	var dependencies DependecyRelation
	for _, gfi := range cs {
		// var depFiles []string
		var depFiles Vendors
		// gfi are fixed, gfj id looped.
		for _, gfj := range cs {
			for pkg, selectors := range gfi.depmaps.deps {
				if pkg != gfj.getPackageName() {
					continue
				}
				//k:v = pkg:
				for _, s := range selectors {
					if util.Contains(gfj.decl.funcs.toStrings(), s) || util.Contains(gfj.decl.types.toStrings(), s) {
						depFiles.add(gfj.decl.filename.Value, gfj.decl.packageName)
					}
				}
			}
		}
		dep := newDependency(gfi.decl.filename.Value, gfi.decl.packageName, depFiles)
		dependencies.add(dep)
	}
	return dependencies
}
