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

	conns, err := f.getConnections(b, s, inputRegexp, f.conveyorGraph.To)
	if err != nil {
		return nil, fmt.Errorf("failed to load input: %w", err)
	}

	switch len(conns) {
	case 0:
	case 1:
		i.c = conns[0]
	default:
		return nil, fmt.Errorf("expected to only have 1 input but had %d", len(conns))
	}

	return i, nil
}

// GetInput returns the Connection for this input.
// If the input is not connected to anything then nil will be returned.
func (i *input) GetInput() *Connection {
	return i.c
}

type inputs struct {
	c []*Connection
}

func (f *Factory) loadInputs(b *building, s *save.Save) (*inputs, error) {
	i := &inputs{}

	conns, err := f.getConnections(b, s, inputRegexp, f.conveyorGraph.To)
	if err != nil {
		return nil, fmt.Errorf("failed to load input: %w", err)
	}

	i.c = conns

	return i, nil
}

func (i *inputs) GetInputs() []*Connection {
	return i.c
}

type output struct {
	c *Connection
}

func (f *Factory) loadOutput(b *building, s *save.Save) (*output, error) {
	o := &output{}

	conns, err := f.getConnections(b, s, outputRegexp, f.conveyorGraph.From)
	if err != nil {
		return nil, fmt.Errorf("failed to load output: %w", err)
	}

	switch len(conns) {
	case 0:
	case 1:
		o.c = conns[0]
	default:
		return nil, fmt.Errorf("expected to only have 1 output but had %d", len(conns))
	}

	return o, nil
}

// GetOutput returns the Connection for this output.
// If the output is not connected to anything then nil will be returned.
func (o *output) GetOutput() *Connection {
	return o.c
}

type outputs struct {
	c []*Connection
}

func (f *Factory) loadOutputs(b *building, s *save.Save) (*outputs, error) {
	o := &outputs{}

	conns, err := f.getConnections(b, s, outputRegexp, f.conveyorGraph.From)
	if err != nil {
		return nil, fmt.Errorf("failed to load input: %w", err)
	}

	o.c = conns

	return o, nil
}

func (o *outputs) GetOutputs() []*Connection {
	return o.c
}

func (f *Factory) getConnections(b *building, s *save.Save, re *regexp.Regexp, direction func(int64) graph.Nodes) ([]*Connection, error) {
	conns := make([]*Connection, 0)

	for _, ref := range b.entity.References {
		if re.MatchString(ref.PathName) {
			c := s.Components[ref.PathName]

			connProp, err := getPropFromArray("mConnectedComponent", c.Properties)
			if err != nil {
				continue
			}

			obj, err := connProp.GetObjectValue()
			if err != nil {
				return nil, err
			}

			id, err := getID(obj.PathName)
			if err != nil {
				return nil, err
			}

			conn := &Connection{}
			conn = f.next(conn, id, direction)

			if len(conn.ConveyorBelts) == 0 {
				continue
			}

			conns = append(conns, conn)
		}
	}

	return conns, nil
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
