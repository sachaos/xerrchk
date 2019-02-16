package a

import "golang.org/x/xerrors"

var sentinelErr = xerrors.New("sentinel err")

func privateFunc1() error {
	return errFunc1() // want `wrap with xerrros.Wrap or hide with xerrors.Opaque`
}

func privateFunc2() error {
	return xerrors.Errorf("wrap: %w", errFunc1()) // OK
}

func privateFunc3() error {
	_, err := errFunc2()
	return err // want `wrap with xerrros.Wrap or hide with xerrors.Opaque`
}

func privateFunc4() error {
	_, err := errFunc2()
	return xerrors.Errorf("wrap: %w", err) // OK
}

func errFunc1() error {
	return xerrors.New("foo")
}

func errFunc2() (int, error) {
	return 0, xerrors.New("foo")
}
