package err

import (
	"fmt"
	"os"
)

type ExitCode int

const Error ExitCode = ExitCode(1)

type ExitError struct {
	exitCode ExitCode
	err      error
}

func (ee ExitError) ExitCode() int {
	return int(ee.exitCode)
}

func (ee *ExitError) Error() string {
	if ee.err == nil {
		return ""
	}
	return fmt.Sprintf("%v", ee.err)
}

func NewExitError(exitCode ExitCode, err error) *ExitError {
	return &ExitError{
		exitCode: exitCode,
		err:      err,
	}
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}
