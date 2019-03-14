package a

type Error struct{}

func (e Error) Error() string {
	return "something error"
}

func (e Error) Wrap(next error) error { // want `wrap with xerrros.Errorf or xerrors.Opaque`
	return &e
}
