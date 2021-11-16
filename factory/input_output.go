package factory

import (
	"regexp"

	"gonum.org/v1/gonum/graph"

	"github.com/l-ross/ficsit-toolkit/save"
)

type Connection struct {
	Connected     Building
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
	// TODO: Error if we find multiple inputs

	i := &input{}

	for _, ref := range b.entity.References {
		if inputRegexp.MatchString(ref.PathName) {
			c := s.Components[ref.PathName]
			if c == nil {
				return nil, nil
			}

			conn, err := getPropFromArray("mConnectedComponent", c.Properties)
			if err != nil {
				return nil, err
			}

			obj, err := conn.GetObjectValue()
			if err != nil {
				return nil, err
			}

			id, err := getID(obj.PathName)
			if err != nil {
				return nil, err
			}

			i.c = f.next(&Connection{}, id, f.conveyorGraph.To)
		}
	}

	return i, nil
}

// GetInput returns the
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

	for _, ref := range b.entity.References {
		if outputRegexp.MatchString(ref.PathName) {
			c := s.Components[ref.PathName]
			if c == nil {
				return nil, nil
			}

			conn, err := getPropFromArray("mConnectedComponent", c.Properties)
			if err != nil {
				return nil, err
			}

			obj, err := conn.GetObjectValue()
			if err != nil {
				return nil, err
			}

			id, err := getID(obj.PathName)
			if err != nil {
				return nil, err
			}

			o.c = f.next(&Connection{}, id, f.conveyorGraph.From)
		}
	}

	return o, nil
}

func (o *output) GetOutput() *Connection {
	return o.c
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
