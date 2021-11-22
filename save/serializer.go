package save

import (
	"io"

	"github.com/l-ross/ficsit-toolkit/save/data"
)

type serializer struct {
	w io.Writer

	*data.Data
}

// Serialize the Save to the provided writer.
func (s *Save) Serialize(w io.Writer) error {
	sr := &serializer{
		w:    w,
		Data: data.New(),
	}

	// Serialize the uncompressed body.
	err := sr.serializeBody(s)
	if err != nil {
		return err
	}

	// Take a copy of the uncompressed save
	uncompressed := make([]byte, sr.Len())
	copy(uncompressed, sr.Bytes())

	// Empty out p.body as we will now write the header / compressed data
	sr.Data = data.New()

	err = sr.serializeHeader(s.Header)
	if err != nil {
		return err
	}

	err = sr.compressBody(uncompressed)
	if err != nil {
		return err
	}

	_, err = sr.w.Write(sr.Bytes())
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

	for _, eName := range save.entityOrder {
		e := save.Entities[eName]

		err = s.WriteInt32(entityType)
		if err != nil {
			return err
		}

		err = e.serialize(s)
		if err != nil {
			return err
		}
	}

	for _, cName := range save.componentOrder {
		c := save.Components[cName]

		err = s.WriteInt32(componentType)
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

	for _, eName := range save.entityOrder {
		e := save.Entities[eName]

		err = e.serializeData(s)
		if err != nil {
			return err
		}
	}

	for _, cName := range save.componentOrder {
		c := save.Components[cName]

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
