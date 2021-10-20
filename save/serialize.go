package save

import (
	"io"

	"github.com/ViRb3/slicewriteseek"
)

func Serialize(s *Save, w io.Writer) error {
	p := &parser{
		body: slicewriteseek.New(),
	}

	// Write placeholder length.
	err := p.writeInt32(0)
	if err != nil {
		return err
	}

	err = p.serializeObjects(s)
	if err != nil {
		return err
	}

	err = p.serializeObjectData(s)
	if err != nil {
		return err
	}

	//_, err = w.Write(p.body.Buff.Bytes())
	//if err != nil {
	//	return err
	//}

	return nil
}

func (p *parser) serializeObjects(s *Save) error {
	objectCount := len(s.Components) + len(s.Entities)
	err := p.writeInt32(int32(objectCount))
	if err != nil {
		return err
	}

	for _, e := range s.Entities {
		err = p.writeInt32(int32(EntityType))
		if err != nil {
			return err
		}

		err = p.serializeEntity(e)
		if err != nil {
			return err
		}
	}

	for _, c := range s.Components {
		err = p.writeInt32(int32(ComponentType))
		if err != nil {
			return err
		}

		err = p.serializeComponent(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *parser) serializeObjectData(s *Save) error {
	objectCount := len(s.Components) + len(s.Entities)
	err := p.writeInt32(int32(objectCount))
	if err != nil {
		return err
	}

	//for _, e := range save.Entities {
	//	// Record location of the length.
	//	startPos := p.body.Index
	//
	//	// Write placeholder length
	//	err := p.writeInt32(0)
	//	if err != nil {
	//		return err
	//	}
	//
	//
	//}

	return nil
}
