package main

import (
	"fmt"
	"os"

	"github.com/keikohi/gorchitect/driver/cmd"
	"github.com/keikohi/gorchitect/registry"
)

func main() {

	fp, err := os.Create("./result.dot")
	if err != nil {
		fmt.Fprintln(os.Stderr, "counld not create a result file.")
		os.Exit(1)
	}
	writer := registry.NewWriter(fp)

	errCode := ErrorHandler(cmd.Run(writer))
	os.Exit(errCode)
}

type ExitCoder interface {
	ExitCode() int
	Error() string
}

const ExitCodeOk int = 0
const ExitCodeError int = 1
const UnexpectedExitCode int = 2

func ErrorHandler(err error) int {
	if err == nil {
		return ExitCodeOk
	}
	if exitErr, ok := err.(ExitCoder); ok {
		fmt.Fprintln(os.Stderr, exitErr.Error())
		return exitErr.ExitCode()
	}

	if _, ok := err.(error); ok {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return ExitCodeError
	}

	fmt.Fprintf(os.Stderr, "Unexpected Error: %+v\n", err)
	return UnexpectedExitCode
}
