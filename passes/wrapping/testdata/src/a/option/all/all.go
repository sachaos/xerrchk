package all

import "golang.org/x/xerrors"

var sentinelErr = xerrors.New("sentinel err")

func privateFunc1() error {
	return errFunc1() // want `wrap with xerrros.Errorf or xerrors.Opaque`
}

func privateFunc2() error {
	return xerrors.Errorf("wrap: %w", errFunc1()) // OK
}

func privateFunc3() error {
	_, err := errFunc2() // want `wrap with xerrros.Errorf or xerrors.Opaque`
	return err
}

func privateFunc4() error {
	_, err := errFunc2() // OK
	return xerrors.Errorf("wrap: %w", err)
}

func privateFunc5() error {
	return sentinelErr // want `wrap with xerrros.Errorf or xerrors.Opaque`
}

func PublicFunc1() error {
	return errFunc1() // want `wrap with xerrros.Errorf or xerrors.Opaque`
}

func PublicFunc2() error {
	return xerrors.Errorf("wrap: %w", errFunc1()) // OK
}

func PublicFunc3() error {
	_, err := errFunc2() // want `wrap with xerrros.Errorf or xerrors.Opaque`
	return err
}

func PublicFunc4() error {
	_, err := errFunc2() // OK
	return xerrors.Errorf("wrap: %w", err)
}

func PublicFunc5() error {
	return sentinelErr // want `wrap with xerrros.Errorf or xerrors.Opaque`
}

func errFunc1() error {
	return xerrors.New("foo")
}

func errFunc2() (int, error) {
	return 0, xerrors.New("foo")
}
