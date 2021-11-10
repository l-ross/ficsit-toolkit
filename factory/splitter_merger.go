package factory

import "github.com/l-ross/ficsit-toolkit/save"

type Splitter struct {
	*building
	*storage
}

func (f *Factory) LoadSplitter(e *save.Entity, s *save.Save) (*Splitter, error) {
	b, err := f.loadBuilding(e, s)
	if err != nil {
		return nil, err
	}

	st, err := f.loadStorage(b, s)
	if err != nil {
		return nil, err
	}

	spl := &Splitter{
		building: b,
		storage:  st,
	}

	f.buildings[spl.node.ID()] = spl

	return spl, nil
}

type Merger struct {
	*building
	*storage
}

func (f *Factory) LoadMerger(e *save.Entity, s *save.Save) (*Merger, error) {
	b, err := f.loadBuilding(e, s)
	if err != nil {
		return nil, err
	}

	st, err := f.loadStorage(b, s)
	if err != nil {
		return nil, err
	}

	m := &Merger{
		building: b,
		storage:  st,
	}

	f.buildings[m.node.ID()] = m

	return m, nil
}
