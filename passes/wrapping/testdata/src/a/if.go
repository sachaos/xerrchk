package a

import "golang.org/x/xerrors"

var sentinelErr = xerrors.New("sentinel error")

func IfStatement(cond bool) error {
	var err error // want `wrap with xerrros.Errorf or xerrors.Opaque`
	if cond {
		err = sentinelErr
	} else {
		err = xerrors.Errorf("wrap: %w", sentinelErr)
	}

	return err
}

func IfStatement2(cond bool) error {
	var err error // OK
	if cond {
		err = xerrors.Errorf("wrap1: %w", sentinelErr)
	} else {
		err = xerrors.Errorf("wrap2 %w", sentinelErr)
	}

	return err
}
