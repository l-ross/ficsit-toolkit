package factory

import (
	"io"

	"gonum.org/v1/gonum/graph/simple"

	"github.com/l-ross/ficsit-toolkit/save"
)

type Factory struct {
	s *save.Save

	g *simple.DirectedGraph
}

func Load(r io.Reader) (*Factory, error) {
	s, err := save.Parse(r)
	if err != nil {
		return nil, err
	}

	f := &Factory{
		s: s,
		g: simple.NewDirectedGraph(),
	}

	for _, e := range s.Entities {
		switch e.TypePath {
		case "/Game/FactoryGame/Buildable/Factory/ConstructorMk1/Build_ConstructorMk1.Build_ConstructorMk1_C":
			_, err := f.LoadConstructor(e, s)
			if err != nil {
				return nil, err
			}
		case "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk1/Build_ConveyorBeltMk1.Build_ConveyorBeltMk1_C",
			"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk2/Build_ConveyorBeltMk2.Build_ConveyorBeltMk2_C":
			_, err := f.LoadConveyor(e, s)
			if err != nil {
				return nil, err
			}
		}
	}

	for _, c := range s.Components {
		switch c.TypePath {
		case "/Script/FactoryGame.FGFactoryConnectionComponent":
			err := f.LoadConnection(c)
			if err != nil {
				return nil, err
			}
		}
	}

	return f, nil
}
