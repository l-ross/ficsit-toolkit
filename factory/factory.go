package factory

import (
	"io"

	"github.com/l-ross/ficsit-toolkit/save"
	"gonum.org/v1/gonum/graph/simple"
)

const (
	timeMultiplier = 1_000_000_000
)

type Factory struct {
	s                 *save.Save
	conveyorGraph     *simple.DirectedGraph
	Buildings         map[string]Building
	Mergers           map[string]*Merger
	Splitters         map[string]*Splitter
	StorageContainers map[string]*StorageContainer
	Conveyors         map[string]Conveyor
	Production        map[string]Production

	// Stores the next ID that should be assigned for a building in the conveyorGraph.
	nextID int64
}

// Load the provided Satisfactory save file.
func Load(r io.Reader) (*Factory, error) {
	s, err := save.Parse(r)
	if err != nil {
		return nil, err
	}

	f := &Factory{
		s:             s,
		conveyorGraph: simple.NewDirectedGraph(),
		Buildings:     make(map[string]Building),
		Production:    make(map[string]Production),
		Conveyors:     make(map[string]Conveyor),

		Mergers:           make(map[string]*Merger),
		Splitters:         make(map[string]*Splitter),
		StorageContainers: make(map[string]*StorageContainer),
	}

	// Load all buildings
	for _, e := range s.Entities {
		if isLoadableBuilding(e) {
			b, err := f.loadBuilding(e, s)
			if err != nil {
				return nil, err
			}

			f.Buildings[b.InstanceName()] = b
		}
	}

	// Load all connections
	for _, c := range s.Components {
		switch c.TypePath {
		case "/Script/FactoryGame.FGFactoryConnectionComponent":
			err := f.loadConveyorConnection(c)
			if err != nil {
				return nil, err
			}
		}
	}

	// Prioritised loading of all buildings
	for _, loaders := range prioritizedLoading {
		for _, b := range f.Buildings {
			if loader, ok := loaders[b.TypePath()]; ok {
				b2, err := loader(f, b.(*building), s)
				if err != nil {
					return nil, err
				}

				f.Buildings[b2.InstanceName()] = b2
			}
		}
	}

	return f, nil
}

func (f *Factory) id() int64 {
	f.nextID++
	return f.nextID
}
