package a

import "fmt"

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

func switchErrCase1(err error) string {
	switch err { // want `a`
	case sentinelErr:
		return "true"
	default:
		return "false"
	}
}

func switchErrCase2(err error) string {
	switch TestCauseFunc(err) { // OK
	case sentinelErr:
		return "true"
	default:
		return "false"
	}
}

func TestCauseFunc(err error) error {
	return err
}
