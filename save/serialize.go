package save

import (
	"io"

	"github.com/mattetti/filebuffer"
)

func (s *Save) Serialize(w io.Writer) error {
	s.body = filebuffer.New([]byte{})

	// Write placeholder length.
	err := s.writeInt32(0)
	if err != nil {
		return err
	}

	err = s.serializeObjects()
	if err != nil {
		return err
	}

	err = s.serializeObjectData()
	if err != nil {
		return err
	}

	//_, err = w.Write(s.body.Buff.Bytes())
	//if err != nil {
	//	return err
	//}

	return nil
}

func (s *Save) serializeObjects() error {
	objectCount := len(s.Components) + len(s.Entities)
	err := s.writeInt32(int32(objectCount))
	if err != nil {
		return err
	}

	for _, e := range s.Entities {
		err = s.writeInt32(int32(EntityType))
		if err != nil {
			return err
		}

		err = s.serializeEntity(e)
		if err != nil {
			return err
		}
	}

	for _, c := range s.Components {
		err = s.writeInt32(int32(ComponentType))
		if err != nil {
			return err
		}

		err = s.serializeComponent(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Save) serializeObjectData() error {
	objectCount := len(s.Components) + len(s.Entities)
	err := s.writeInt32(int32(objectCount))
	if err != nil {
		return err
	}

	//for _, e := range save.Entities {
	//	// Record location of the length.
	//	startPos := s.body.Index
	//
	//	// Write placeholder length
	//	err := s.writeInt32(0)
	//	if err != nil {
	//		return err
	//	}
	//
	//
	//}

	return nil
}
