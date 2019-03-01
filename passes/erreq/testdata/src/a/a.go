package a

import "fmt"

var sentinelErr = fmt.Errorf("sentinel err")

func main() {
	_ = alwaysErr() == nil         // OK
	_ = alwaysErr() == sentinelErr // want `do not compare error with "==" or "!="`
	_ = alwaysErr() != nil         // OK
	_ = alwaysErr() != sentinelErr // want `do not compare error with "==" or "!="`
}

func alwaysErr() error {
	return sentinelErr
}
