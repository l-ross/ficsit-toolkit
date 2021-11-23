package factory

import "github.com/l-ross/ficsit-toolkit/save"

type Splitter struct {
	*building
	*storage
	*input
	*outputs
}

func (f *Factory) loadSplitter(b *building, s *save.Save) (Building, error) {
	st, err := f.loadStorage(b, s)
	if err != nil {
		return nil, err
	}

	i, err := f.loadInput(b, s)
	if err != nil {
		return nil, err
	}

	o, err := f.loadOutputs(b, s)
	if err != nil {
		return nil, err
	}

	spl := &Splitter{
		building: b,
		storage:  st,
		input:    i,
		outputs:  o,
	}

	f.Splitters[b.id] = spl

	return spl, nil
}

type Merger struct {
	*building
	*storage
	*inputs
	*output
}

func (f *Factory) loadMerger(b *building, s *save.Save) (Building, error) {
	st, err := f.loadStorage(b, s)
	if err != nil {
		return nil, err
	}

	i, err := f.loadInputs(b, s)
	if err != nil {
		return nil, err
	}

	o, err := f.loadOutput(b, s)
	if err != nil {
		return nil, err
	}

	m := &Merger{
		building: b,
		storage:  st,
		inputs:   i,
		output:   o,
	}

	f.Mergers[b.id] = m

	return m, nil
}
