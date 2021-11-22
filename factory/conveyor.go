package factory

import (
	"fmt"
	"strings"

	"github.com/l-ross/ficsit-toolkit/factory/typepath"
	BuildableConveyorBelt "github.com/l-ross/ficsit-toolkit/resource/buildable_conveyor_belt"
	BuildableConveyorLift "github.com/l-ross/ficsit-toolkit/resource/buildable_conveyor_lift"
	"github.com/l-ross/ficsit-toolkit/save"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

// Conveyor is implemented by ConveyorBelt and ConveyorLift
type Conveyor interface {
	BeltContents() []InventoryStack
}

type ConveyorBelt struct {
	Mark int

	BeltDescriptor BuildableConveyorBelt.FGBuildableConveyorBelt

	beltContents []InventoryStack
	*building
}

// BeltContents will return the items that are currently on this ConveyorBelt
func (c *ConveyorBelt) BeltContents() []InventoryStack {
	return c.beltContents
}

func (f *Factory) loadConveyorBelt(b *building, s *save.Save) (Building, error) {
	c := &ConveyorBelt{
		building: b,
	}

	switch b.typePath {
	case typepath.ConveyorBeltMk1:
		c.BeltDescriptor = BuildableConveyorBelt.ConveyorBeltMk1
		c.Mark = 1
	case typepath.ConveyorBeltMk2:
		c.BeltDescriptor = BuildableConveyorBelt.ConveyorBeltMk2
		c.Mark = 2
	case typepath.ConveyorBeltMk3:
		c.BeltDescriptor = BuildableConveyorBelt.ConveyorBeltMk3
		c.Mark = 3
	case typepath.ConveyorBeltMk4:
		c.BeltDescriptor = BuildableConveyorBelt.ConveyorBeltMk4
		c.Mark = 4
	case typepath.ConveyorBeltMk5:
		c.BeltDescriptor = BuildableConveyorBelt.ConveyorBeltMk5
		c.Mark = 5
	default:
		return nil, fmt.Errorf("unknown conveyor belt type path %s", b.typePath)
	}

	var err error
	c.beltContents, err = loadBeltContents(b)
	if err != nil {
		return nil, err
	}

	f.Conveyors[c.ID()] = c

	return c, nil
}

type ConveyorLift struct {
	Mark           int
	LiftDescriptor BuildableConveyorLift.FGBuildableConveyorLift

	beltContents []InventoryStack
	*building
}

// BeltContents will return the items that are currently on this ConveyorLift
func (c *ConveyorLift) BeltContents() []InventoryStack {
	return c.beltContents
}

func (f *Factory) loadConveyorLift(b *building, s *save.Save) (Building, error) {
	c := &ConveyorLift{
		building: b,
	}

	switch b.typePath {
	case typepath.ConveyorLiftMk1:
		c.LiftDescriptor = BuildableConveyorLift.ConveyorLiftMk1
		c.Mark = 1
	case typepath.ConveyorLiftMk2:
		c.LiftDescriptor = BuildableConveyorLift.ConveyorLiftMk2
		c.Mark = 2
	case typepath.ConveyorLiftMk3:
		c.LiftDescriptor = BuildableConveyorLift.ConveyorLiftMk3
		c.Mark = 3
	case typepath.ConveyorLiftMk4:
		c.LiftDescriptor = BuildableConveyorLift.ConveyorLiftMk4
		c.Mark = 4
	case typepath.ConveyorLiftMk5:
		c.LiftDescriptor = BuildableConveyorLift.ConveyorLiftMk5
		c.Mark = 5
	default:
		return nil, fmt.Errorf("unknown conveyor lift type path %s", b.typePath)
	}

	var err error
	c.beltContents, err = loadBeltContents(b)
	if err != nil {
		return nil, err
	}

	f.Conveyors[c.ID()] = c

	return c, nil
}

func loadBeltContents(b *building) ([]InventoryStack, error) {
	e, err := b.entity.Extra.GetConveyorBelt()
	if err != nil {
		return nil, err
	}

	itemCount := make(map[string]int)

	for _, item := range e.Items {
		itemCount[item.ResourceName]++
	}

	is := make([]InventoryStack, 0)

	for item, count := range itemCount {
		iDesc, err := getItemDescriptorByName(item)
		if err != nil {
			return nil, err
		}

		is = append(is, InventoryStack{
			Item:     iDesc,
			Quantity: count,
		})
	}

	return is, nil
}

func (f *Factory) loadConveyorConnection(c *save.Component) error {
	n1, err := getID(c.InstanceName)
	if err != nil {
		return err
	}
	var n2 int64

	// Assume we are going from node 1 to node 2.
	n1ToN2 := true

	if inputRegexp.MatchString(c.InstanceName) {
		// If the Component is an input then our assumption about direction is wrong.
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
