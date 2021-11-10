package factory

import (
	"io"

	"gonum.org/v1/gonum/graph"

	"gonum.org/v1/gonum/graph/simple"

	"github.com/l-ross/ficsit-toolkit/save"
)

type Factory struct {
	s *save.Save

	conveyorGraph *simple.DirectedGraph
	buildings     map[int64]Building

	mergers           []*Merger
	splitters         []*Splitter
	constructors      []*Constructor
	conveyors         []*Conveyor
	storageContainers []*StorageContainer
}

func Load(r io.Reader) (*Factory, error) {
	s, err := save.Parse(r)
	if err != nil {
		return nil, err
	}

	f := &Factory{
		s:             s,
		conveyorGraph: simple.NewDirectedGraph(),
		buildings:     make(map[int64]Building),
	}

	//for _, c := range s.Components {
	//	switch c.TypePath {
	//	case "/Script/FactoryGame.FGFactoryConnectionComponent":
	//		err := f.LoadConnection(c)
	//		if err != nil {
	//			return nil, err
	//		}
	//	}
	//}

	for _, e := range s.Entities {
		switch e.TypePath {
		case "/Game/FactoryGame/Buildable/Factory/CA_Merger/Build_ConveyorAttachmentMerger.Build_ConveyorAttachmentMerger_C":
			m, err := f.LoadMerger(e, s)
			if err != nil {
				return nil, err
			}
			f.mergers = append(f.mergers, m)
		case "/Game/FactoryGame/Buildable/Factory/CA_Splitter/Build_ConveyorAttachmentSplitter.Build_ConveyorAttachmentSplitter_C":
			spl, err := f.LoadSplitter(e, s)
			if err != nil {
				return nil, err
			}
			f.splitters = append(f.splitters, spl)
		case "/Game/FactoryGame/Buildable/Factory/StorageContainerMk2/Build_StorageContainerMk2.Build_StorageContainerMk2_C":
			st, err := f.LoadStorageContainer(e, s)
			if err != nil {
				return nil, err
			}
			f.storageContainers = append(f.storageContainers, st)
		case "/Game/FactoryGame/Buildable/Factory/ConstructorMk1/Build_ConstructorMk1.Build_ConstructorMk1_C":
			c, err := f.LoadConstructor(e, s)
			if err != nil {
				return nil, err
			}
			f.constructors = append(f.constructors, c)
		case "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk1/Build_ConveyorBeltMk1.Build_ConveyorBeltMk1_C",
			"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk2/Build_ConveyorBeltMk2.Build_ConveyorBeltMk2_C":
			c, err := f.LoadConveyor(e, s)
			if err != nil {
				return nil, err
			}
			f.conveyors = append(f.conveyors, c)
		}
	}

	//for _, b := range f.buildings {
	//	err := f.setInputs(b, s)
	//	if err != nil {
	//		return nil, err
	//	}
	//}

	//for _, c := range f.splitters {
	//	next := f.next(c.node)
	//	spew.Dump(next)
	//}

	return f, nil
}

func (f *Factory) next(n graph.Node) Building {
	nodes := f.conveyorGraph.From(n.ID())
	if !nodes.Next() {
		return nil
	}

	node := nodes.Node()
	b := f.buildings[node.ID()]
	switch b.(type) {
	case nil:
		return nil
	case *Conveyor:
		// TODO: Stop possible infinite loop
		return f.next(b.Node())
	}

	return b
}
