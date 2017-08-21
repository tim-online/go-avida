package avida

import "io"

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}
