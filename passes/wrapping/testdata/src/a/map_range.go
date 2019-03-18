package a

// NOTE: Case https://github.com/sachaos/xerrchk/issues/6
func MapRange() error {
	maps := make(map[string]string)

	for range maps {
	}

	return nil
}
