package factory

import (
	"strings"

	"gonum.org/v1/gonum/graph"

	"gonum.org/v1/gonum/graph/simple"

	"github.com/l-ross/ficsit-toolkit/save"
)

type Conveyor struct {
	*building
}

func (f *Factory) LoadConveyor(e *save.Entity, s *save.Save) (*Conveyor, error) {
	b, err := f.loadBuilding(e, s)
	if err != nil {
		return nil, err
	}

	c := &Conveyor{
		building: b,
	}

	f.buildings[c.node.ID()] = c

	return c, nil
}

func (f *Factory) LoadConnection(c *save.Component) (graph.Edge, error) {
	n1, err := getID(c.InstanceName)
	if err != nil {
		return nil, err
	}
	var n2 int64

	// Assume we are going from node 1 to node 2.
	n1ToN2 := true

	if inputRegexp.MatchString(c.InstanceName) {
		// In the Component is an input then our assumption about direction is wrong.
		n1ToN2 = false
	}

	for _, prop := range c.Properties {
		switch prop.Name {
		case "mDirection":
			e, err := prop.GetEnumPropertyValue()
			if err != nil {
				return nil, err
			}
			if strings.Contains(e.Value, "FCD_INPUT") {
				// If the direction is input then our assumption about direction is wrong.
				n1ToN2 = false
			}
		case "mConnectedComponent":
			o, err := prop.GetObjectValue()
			if err != nil {
				return nil, err
			}
			n2, err = getID(o.PathName)
			if err != nil {
				return nil, err
			}
		}
	}

	if n2 == 0 {
		return nil, nil
	}

	var e graph.Edge

	// Create edge based on direction.
	if n1ToN2 {
		e = f.conveyorGraph.NewEdge(simple.Node(n1), simple.Node(n2))
	} else {
		e = f.conveyorGraph.NewEdge(simple.Node(n2), simple.Node(n1))
	}

	f.conveyorGraph.SetEdge(e)

	return e, nil
}
