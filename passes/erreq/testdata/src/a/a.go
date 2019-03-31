package a

import (
	"fmt"

	"golang.org/x/xerrors"
)

var sentinelErr = fmt.Errorf("sentinel err")

func main() {
	_ = alwaysErr() == nil         // OK
	_ = alwaysErr() == sentinelErr // want `do not compare error with \"==\" or \"!=\"`
	_ = alwaysErr() != nil         // OK
	_ = alwaysErr() != sentinelErr // want `do not compare error with \"==\" or \"!=\"`
}

func alwaysErr() error {
	return sentinelErr
}

func switchErrCase1() string {
	err := xerrors.New("write failed")
	switch err { // want `do not use error on switch statement`
	case sentinelErr:
		return "true"
	default:
		return "false"
	}
}

func switchErrCase2() string {
	err := xerrors.New("write failed")
	switch xerrors.Unwrap(err) { // want `do not use error on switch statement`
	case sentinelErr:
		return "true"
	default:
		return "false"
	}
}
