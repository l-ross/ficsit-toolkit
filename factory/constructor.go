package factory

import (
	"github.com/l-ross/ficsit-toolkit/save"
)

type Constructor struct {
	*production
	*building
	//Output *Connection
	//Input  *Connection
	*input
	*output
}

func (f *Factory) LoadConstructor(e *save.Entity, s *save.Save) (*Constructor, error) {
	b, err := f.loadBuilding(e, s)
	if err != nil {
		return nil, err
	}

	p, err := f.loadProduction(b, s)
	if err != nil {
		return nil, err
	}

	i, err := f.loadInput(b, s)
	if err != nil {
		return nil, err
	}

	o, err := f.loadOutput(b)
	if err != nil {
		return nil, err
	}

	c := &Constructor{
		building:   b,
		production: p,
		input:      i,
		output:     o,
	}

	f.buildings[c.node.ID()] = c

	return c, nil
}

type Connection struct {
	Connected     Building
	ConveyorBelts []Conveyor
}
