package domain

type declaration struct {
	filename    Path
	packageName string
	imports     imprts
	funcs       fncs
	types       typs
}

type decl string
type decls []decl

func (ds decls) toStrings() []string {
	strs := make([]string, len(ds), len(ds))
	for _, decl := range ds {
		strs = append(strs, string(decl))
	}
	return strs
}

type imprts struct{ decls }
type fncs struct{ decls }
type typs struct{ decls }

func newImoprts(imps []string) imprts {
	var decls decls
	for _, imp := range imps {
		decls.add(imp)
	}
	return imprts{decls}
}

func (ds *decls) add(s string) {
	decl := decl(s)
	*ds = append(*ds, decl)
}

type dependencyMaps struct {
	funcs dependencyMap
	types dependencyMap
	deps  dependencyMap
}
