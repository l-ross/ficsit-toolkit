package factory

import (
	BuildingDescriptor "github.com/l-ross/ficsit-toolkit/resource/building_descriptor"
	"github.com/l-ross/ficsit-toolkit/save"
)

type Constructor struct {
	*production
	*building
	*input
	*output
}

func (f *Factory) loadConstructor(b *building, s *save.Save) (Building, error) {
	b.descriptor = BuildingDescriptor.ConstructorMk1

	p, err := f.loadProduction(b, s)
	if err != nil {
		return nil, err
	}

	i, err := f.loadInput(b, s)
	if err != nil {
		return nil, err
	}

	o, err := f.loadOutput(b, s)
	if err != nil {
		return nil, err
	}

	c := &Constructor{
		building:   b,
		production: p,
		input:      i,
		output:     o,
	}

	f.Production[b.InstanceName()] = c

	return c, err
}
