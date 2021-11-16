package factory

import (
	"io"
	"strings"

	"github.com/l-ross/ficsit-toolkit/save"
	"gonum.org/v1/gonum/graph/simple"
)

type Factory struct {
	s *save.Save

	conveyorGraph *simple.DirectedGraph
	Buildings     map[int64]Building

	Mergers           map[int64]*Merger
	Splitters         map[int64]*Splitter
	Constructors      map[int64]*Constructor
	Conveyors         map[int64]*Conveyor
	StorageContainers map[int64]*StorageContainer
}

func Load(r io.Reader) (*Factory, error) {
	s, err := save.Parse(r)
	if err != nil {
		return nil, err
	}

	f := &Factory{
		s:             s,
		conveyorGraph: simple.NewDirectedGraph(),
		Buildings:     make(map[int64]Building),
		Constructors:  make(map[int64]*Constructor),
	}

	// Load all buildings
	for _, e := range s.Entities {
		if isBuilding(e) {
			b, err := f.loadBuilding(e, s)
			if err != nil {
				return nil, err
			}

			f.Buildings[b.ID()] = b
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
	for _, loaders := range prioritisedLoading {
		for _, b := range f.Buildings {
			if loader, ok := loaders[b.TypePath()]; ok {
				b2, err := loader(f, b.(*building), s)
				if err != nil {
					return nil, err
				}

				f.Buildings[b2.ID()] = b2
			}
		}
	}

	return f, nil
}

func isBuilding(e *save.Entity) bool {
	if e.ParentObjectName != "Persistent_Level:PersistentLevel.BuildableSubsystem" {
		return false
	}

	if !strings.HasPrefix(e.TypePath, "/Game/FactoryGame/Buildable/Factory/") {
		// We don't want to load things like foundations which start
		// "/Game/FactoryGame/Buildable/Building/"
		return false
	}

	return true
}
