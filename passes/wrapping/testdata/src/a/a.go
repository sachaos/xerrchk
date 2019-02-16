package a

import "golang.org/x/xerrors"

var sentinelErr = xerrors.New("sentinel err")

func main() {
	xerrors.Errorf("hogehoge: %w", sentinelErr)
}