package factory

import (
	"github.com/l-ross/ficsit-toolkit/save"
)

type Constructor struct {
	*Production
}

func (f *Factory) LoadConstructor(e *save.Entity, s *save.Save) (*Constructor, error) {
	p, err := f.loadProduction(e, s)
	if err != nil {
		return nil, err
	}

	c := &Constructor{
		Production: p,
	}

	return c, nil
}
