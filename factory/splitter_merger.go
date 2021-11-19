package factory

import "github.com/l-ross/ficsit-toolkit/save"

type Splitter struct {
	*building
	*storage
}

func (f *Factory) loadSplitter(b *building, s *save.Save) (Building, error) {
	st, err := f.loadStorage(b, s)
	if err != nil {
		return nil, err
	}

	spl := &Splitter{
		building: b,
		storage:  st,
	}

	f.Splitters[b.id] = spl

	return spl, nil
}

type Merger struct {
	*building
	*storage
}

func (f *Factory) loadMerger(b *building, s *save.Save) (Building, error) {
	st, err := f.loadStorage(b, s)
	if err != nil {
		return nil, err
	}

	m := &Merger{
		building: b,
		storage:  st,
	}

	f.Mergers[b.id] = m

	return m, nil
}
