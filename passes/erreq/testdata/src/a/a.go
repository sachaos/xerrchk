package a

import (
	"fmt"
	"golang.org/x/xerrors"
)

var sentinelErr = fmt.Errorf("sentinel err")

func main() {
	_ = alwaysErr() == nil         // OK
	_ = alwaysErr() == sentinelErr // want `do not compare errors with binary ops.`
	_ = alwaysErr() != nil         // OK
	_ = alwaysErr() != sentinelErr // want `do not compare errors with binary ops.`
}

func alwaysErr() error {
	return sentinelErr
}

func switchErrCase1() string {
	err := xerrors.New("write failed")
	switch err { // want `do not use not unwrapped errors as a tag of switch statement.`
	case sentinelErr:
		return "true"
	default:
		return "false"
	}
}

func switchErrCase2() string {
	err := xerrors.New("write failed")
	switch TestCauseFunc(err) { // OK
	case sentinelErr:
		return "true"
	default:
		return "false"
	}
}

func switchErrCase3() string {
	err := xerrors.New("write failed")
	switch xerrors.Unwrap(err) { // OK
	case sentinelErr:
		return "true"
	default:
		return "false"
	}
}

func TestCauseFunc(err error) error {
	return err
}
