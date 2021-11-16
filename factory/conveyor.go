package factory

import (
	"strings"

	BuildableConveyorLift "github.com/l-ross/ficsit-toolkit/resource/buildable_conveyor_lift"

	BuildableConveyorBelt "github.com/l-ross/ficsit-toolkit/resource/buildable_conveyor_belt"

	"github.com/l-ross/ficsit-toolkit/save"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type ConveyorBelt struct {
	tier           int
	BeltDescriptor BuildableConveyorBelt.FGBuildableConveyorBelt
	*building
}

func (c *ConveyorBelt) Tier() int {
	return c.tier
}

type ConveyorLift struct {
	Tier           int
	LiftDescriptor BuildableConveyorLift.FGBuildableConveyorLift
	*building
}

type Conveyor interface {
	Tier() int
}

func (f *Factory) loadConveyorBelt(b *building, s *save.Save) (Building, error) {
	c := &ConveyorBelt{
		building: b,
	}

	return c, nil
}

//func (f *Factory) loadConveyor(b *building, s *save.Save) (Building, error) {
//	c := &Conveyor{
//		building: b,
//	}
//
//	return c, nil
//}

func (f *Factory) loadConveyorConnection(c *save.Component) error {
	n1, err := getID(c.InstanceName)
	if err != nil {
		return err
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
				return err
			}
			if strings.Contains(e.Value, "FCD_INPUT") {
				// If the direction is input then our assumption about direction is wrong.
				n1ToN2 = false
			}
		case "mConnectedComponent":
			o, err := prop.GetObjectValue()
			if err != nil {
				return err
			}
			n2, err = getID(o.PathName)
			if err != nil {
				return err
			}
		}
	}

	if n2 == 0 {
		return nil
	}

	var e graph.Edge

	// Create edge based on direction.
	if n1ToN2 {
		e = f.conveyorGraph.NewEdge(simple.Node(n1), simple.Node(n2))
	} else {
		e = f.conveyorGraph.NewEdge(simple.Node(n2), simple.Node(n1))
	}

	f.conveyorGraph.SetEdge(e)

	return nil
}
