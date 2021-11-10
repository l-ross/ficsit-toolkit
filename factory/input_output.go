package factory

import "github.com/l-ross/ficsit-toolkit/save"

type input struct{}

func (f *Factory) loadInput(b *building, s *save.Save) (*input, error) {
	i := &input{}

	return i, nil
}

func (i *input) GetInput() *Connection {
	return nil
}

type output struct {
	conn *Connection
}

func (f *Factory) loadOutput(b *building) (*output, error) {
	// TODO: Length check
	conn, err := f.createConnection(b.outputs[0])
	if err != nil {
		return nil, err
	}

	o := &output{
		conn: conn,
	}

	return o, nil
}

func (o *output) GetOutput() *Connection {
	return o.conn
}
