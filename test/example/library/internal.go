package library

type internal struct {
	a int
}

func internalFunc() {
	e := External{a: 0, B: 0}
	e.ExternalFunc()
}
