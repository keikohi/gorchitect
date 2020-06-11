package domain

type paths []Path

// Chain is a list of paths to depend on when trying to analyze specific packages.
type Chain struct {
	paths paths
}

func (c *Chain) add(p Path) { (*c).paths = append((*c).paths, p) }

func generateChain(deps DependecyRelation) Chain {
	var chain Chain
	existense := map[Path]bool{}
	for _, dep := range deps {
		for _, depFile := range dep.Vendors {
			path := Path{depFile.Filename}
			if !existense[path] {
				chain.add(path)
				existense[path] = true
			}
		}
		path := Path{dep.Consumer.Filename}
		if !existense[path] {
			chain.add(path)
			existense[path] = true
		}
	}
	return chain
}

func (c Chain) Contain(path Path) bool {
	for _, p := range c.paths {
		if p.contain(path) {
			return true
		}
	}
	return false
}
