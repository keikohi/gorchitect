package domain

type DependecyRelation []Dependency

func (dr *DependecyRelation) add(d Dependency) {
	if d.hasVendors() {
		*dr = append(*dr, d)
	}
}

func (dr *DependecyRelation) distinct() DependecyRelation {
	var dependencies DependecyRelation
	for _, dep := range *dr {
		merged := make(Vendors, 0, len(dep.Vendors))
		existence := map[string]bool{}
		for _, df := range dep.Vendors {
			if !existence[df.Filename] {
				existence[df.Filename] = true
				merged.add(df.Filename, df.packageName)
			}
		}
		dependencies.add(newDependency(dep.Consumer.Filename, dep.Consumer.packageName, merged))
	}
	return dependencies
}

func (dr *DependecyRelation) focusPackage(focusPackage string) (DependecyRelation, Chain) {
	var dependencies DependecyRelation
	for _, dep := range *dr {
		if dep.Consumer.packageName == focusPackage {
			dependencies.add(newDependency(dep.Consumer.Filename, dep.Consumer.packageName, dep.Vendors))
		}
	}
	return dependencies, generateChain(dependencies)
}

func (dr *DependecyRelation) GetDependency() map[string][]string {
	m := make(map[string][]string)
	for _, dep := range *dr {
		tails := make([]string, len(dep.Vendors))
		for _, depFile := range dep.Vendors {
			tails = append(tails, depFile.Filename)
		}
		m[dep.Consumer.Filename] = tails
	}
	return m
}

type vendor struct {
	Filename    string
	packageName string
}

// consumer is a file which depend on vendor.
type consumer struct {
	Filename    string
	packageName string
}

type Vendors []vendor

func (dfs *Vendors) add(fileName string, packageName string) {
	df := vendor{Filename: fileName, packageName: packageName}
	*dfs = append(*dfs, df)
}

// consumer : vendor = 1 : N
type Dependency struct {
	Consumer consumer
	Vendors  Vendors
}

// hasVendors returns true if d has one more vendors.
func (d Dependency) hasVendors() bool { return d.Vendors != nil }

func newDependency(filename string, packageName string, vendors Vendors) Dependency {

	return Dependency{
		Consumer: consumer{
			Filename:    filename,
			packageName: packageName,
		},
		Vendors: vendors}
}
