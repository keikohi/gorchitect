package todo

import (
	"fmt"

	"github.com/keikohi/gorchitect/test/example/library"
)

func testMain() {
	library.ExternalFunc()
	library.ConstTest
	library.Func()

	testType := library.TestType{}

	fmt.Println("fin")
}
