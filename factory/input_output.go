package factory

import (
	"fmt"
	"regexp"

	"github.com/l-ross/ficsit-toolkit/save"
	"gonum.org/v1/gonum/graph"
)

// Connection details what a buildings input or output is connected to.
type Connection struct {
	// Connected is the non-conveyor belt Building that this input or output
	// is connected to.
	//
	// Will be nil if there is no terminating Building.
	Connected Building
	// ConveyorBelts contains all the Conveyor objects in order from the input or output to
	// the Connected Building.
	ConveyorBelts []Conveyor
}

var (
	// Case-insensitive as sometimes 'Input' is spelt 'InPut'.
	inputRegexp  = regexp.MustCompile(`(?i)(Input\d|ConveyorAny0)$`)
	outputRegexp = regexp.MustCompile(`(?i)(Output\d|ConveyorAny1)$`)
)

type input struct {
	c *Connection
}

func (f *Factory) loadInput(b *building, s *save.Save) (*input, error) {
	i := &input{}

	var err error
	i.c, err = f.getConnection(b, s, f.conveyorGraph.To)
	if err != nil {
		return nil, fmt.Errorf("failed to load input: %w", err)
	}

	return i, nil
}

// GetInput returns the Connection for this input.
// If the input is not connected to anything then nil will be returned.
func (i *input) GetInput() *Connection {
	return i.c
}

type inputs []input

func (f *Factory) loadInputs(b *building, s *save.Save) (inputs, error) {
	return nil, nil
}

type output struct {
	c *Connection
}

func (f *Factory) loadOutput(b *building, s *save.Save) (*output, error) {
	o := &output{}

	var err error
	o.c, err = f.getConnection(b, s, f.conveyorGraph.From)
	if err != nil {
		return nil, fmt.Errorf("failed to load output: %w", err)
	}

	return o, nil
}

// GetOutput returns the Connection for this output.
// If the output is not connected to anything then nil will be returned.
func (o *output) GetOutput() *Connection {
	return o.c
}

func (f *Factory) getConnection(b *building, s *save.Save, direction func(int64) graph.Nodes) (*Connection, error) {
	conn := &Connection{}

	for _, ref := range b.entity.References {
		if outputRegexp.MatchString(ref.PathName) {
			c := s.Components[ref.PathName]
			if c == nil {
				return nil, nil
			}

			connProp, err := getPropFromArray("mConnectedComponent", c.Properties)
			if err != nil {
				return nil, err
			}

			obj, err := connProp.GetObjectValue()
			if err != nil {
				return nil, err
			}

			id, err := getID(obj.PathName)
			if err != nil {
				return nil, err
			}

			conn = f.next(conn, id, direction)
		}
	}

	if len(conn.ConveyorBelts) == 0 {
		return nil, nil
	}

	return conn, nil
}

func (f *Factory) next(c *Connection, n int64, direction func(int64) graph.Nodes) *Connection {
	b := f.Buildings[n]
	switch b.(type) {
	case nil:
		return c
	case Conveyor:
		c.ConveyorBelts = append(c.ConveyorBelts, b.(Conveyor))

		// TODO: Check we only have 1 connected node
		nodes := direction(n)
		if !nodes.Next() {
			return c
		}

		node := nodes.Node()

		// TODO: Stop possible infinite loop
		return f.next(c, node.ID(), direction)
	}

	c.Connected = b

	return c
}
