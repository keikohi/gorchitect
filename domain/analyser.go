package domain

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// Analyzer analyses go file package
type Analyzer struct {
	projectRoot string
	filepaths   []string
}

func NewAnalyzer(filepath []string, projectPath string) Analyzer {
	return Analyzer{
		projectRoot: projectPath,
		filepaths:   filepath}
}

//Analyse analyse go files and return dependencies between go files
func (da Analyzer) Analyse() Codes {
	var gofiles Codes
	for _, filepath := range da.filepaths {
		path := Path{filepath}
		if path.IsTest() {
			continue
		}
		decl, depMap := da.inspect(path)
		gofiles.add(decl, depMap)
	}
	return gofiles
}

func (da Analyzer) inspect(filepath Path) (declaration, dependencyMaps) {
	imports := imprts{}
	var funcs fncs
	var types typs

	depmap := newDependencyMap()
	funcmap := newDependencyMap()

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, filepath.fullPath(da.projectRoot), nil, parser.Mode(0))
	ast.Inspect(f, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.FuncDecl:
			v, _ := node.(*ast.FuncDecl)
			if v.Name != nil {
				funcs.add(v.Name.Name)
			}

		case *ast.ImportSpec:
			v, _ := node.(*ast.ImportSpec)
			if v.Path != nil {
				// v.Name.Value has double-quotation
				imprt := strings.Replace(v.Path.Value, "\"", "", -1)
				imports.add(imprt)
			}

		//type struct
		case *ast.TypeSpec:
			v, _ := node.(*ast.TypeSpec)
			if v.Name != nil {
				types.add(v.Name.Name)
			}

		//dependency map : package, function
		case *ast.CallExpr:
			callexpr, _ := node.(*ast.CallExpr)
			if selector, ok := (callexpr.Fun).(*ast.SelectorExpr); ok {
				if x, ok := selector.X.(*ast.Ident); ok {
					funcmap.add(x.Name, selector.Sel.Name)
				}
			}
		//xxx.yyy()
		case *ast.SelectorExpr:
			if se, ok := node.(*ast.SelectorExpr); ok {
				if x, ok := se.X.(*ast.Ident); ok {
					depmap.add(x.Name, se.Sel.Name)
				}
			}
		default:
			_ = n
		}
		return true
	})
	depmap = depmap.distinct().outerImports(imports)
	funcmap = funcmap.distinct().outerImports(imports)
	//  typemap = depmap - funcmap
	typemap := depmap.typemap(funcmap)

	dependencyMaps := dependencyMaps{
		funcs: funcmap,
		types: typemap,
		deps:  depmap}

	declaration := declaration{
		filename:    filepath,
		packageName: f.Name.Name,
		imports:     imports,
		funcs:       funcs,
		types:       types}
	return declaration, dependencyMaps
}
