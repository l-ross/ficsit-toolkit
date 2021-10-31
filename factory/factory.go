package factory

import (
	"io"

	"github.com/l-ross/ficsit-toolkit/save"
)

type Factory struct {
	s *save.Save
}

func Load(r io.Reader) (*Factory, error) {
	s, err := save.Parse(r)
	if err != nil {
		return nil, err
	}

	f := &Factory{
		s: s,
	}

	return f, nil
}
