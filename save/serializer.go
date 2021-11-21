package save

import (
	"io"

	"github.com/l-ross/ficsit-toolkit/save/data"
)

type serializer struct {
	w io.Writer

	*data.Data
}

func Serialize(w io.Writer, save *Save) error {
	s := &serializer{
		w: w,
	}

	// Serialize the uncompressed body.
	err := s.serializeBody(save)
	if err != nil {
		return err
	}

	// Take a copy of the uncompressed save
	uncompressed := make([]byte, s.Len())
	copy(uncompressed, s.Bytes())

	// Empty out p.body as we will now write the header / compressed data
	s.Data = data.New()

	err = s.serializeHeader(save.Header)
	if err != nil {
		return err
	}

	err = s.compressBody(uncompressed)
	if err != nil {
		return err
	}

	_, err = s.w.Write(s.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (s *serializer) serializeBody(save *Save) error {
	// Write placeholder length.
	err := s.WriteInt32(0)
	if err != nil {
		return err
	}

	m := s.Measure()

	err = s.serializeObjects(save)
	if err != nil {
		return err
	}

	err = s.serializeObjectData(save)
	if err != nil {
		return err
	}

	err = s.serializeCollectedObjects(save)
	if err != nil {
		return err
	}

	err = s.WriteLen(m(), 0)
	if err != nil {
		return err
	}

	return nil
}

func (s *serializer) serializeObjects(save *Save) error {
	objectCount := len(save.Components) + len(save.Entities)
	err := s.WriteInt32(int32(objectCount))
	if err != nil {
		return err
	}

	for _, e := range save.Entities {
		err = s.WriteInt32(int32(EntityType))
		if err != nil {
			return err
		}

		err = e.serialize(s)
		if err != nil {
			return err
		}
	}

	for _, c := range save.Components {
		err = s.WriteInt32(int32(ComponentType))
		if err != nil {
			return err
		}

		err = c.serialize(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *serializer) serializeObjectData(save *Save) error {
	objectCount := len(save.Components) + len(save.Entities)
	err := s.WriteInt32(int32(objectCount))
	if err != nil {
		return err
	}

	for _, e := range save.Entities {
		err = e.serializeData(s)
		if err != nil {
			return err
		}
	}

	for _, c := range save.Components {
		err = c.serializeData(s)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *serializer) serializeCollectedObjects(save *Save) error {
	err := s.WriteInt32(int32(len(save.CollectedObjects)))
	if err != nil {
		return err
	}

	for _, c := range save.CollectedObjects {
		err = s.WriteString(c.LevelName)
		if err != nil {
			return err
		}

		err = s.WriteString(c.PathName)
		if err != nil {
			return err
		}
	}

	return nil
}
