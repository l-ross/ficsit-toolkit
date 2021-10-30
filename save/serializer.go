package save

import (
	"io"

	"github.com/ViRb3/slicewriteseek"
)

type Serializer struct {
	w io.Writer

	// Body of the save file.
	body *slicewriteseek.SliceWriteSeeker

	serializerOptions
}

func NewSerializer(w io.Writer, opts ...SerializerOption) (*Serializer, error) {
	s := &Serializer{
		body: &slicewriteseek.SliceWriteSeeker{
			Buffer: make([]byte, 0),
			Index:  0,
		},
	}

	for _, opt := range opts {
		if opt != nil {
			if err := opt(&s.serializerOptions); err != nil {
				return nil, err
			}
		}
	}

	return s, nil
}

func (s *Serializer) Serialize(save *Save) error {
	// Serialize the uncompressed body.
	err := s.serializeBody(save)
	if err != nil {
		return err
	}

	// Take a copy of the uncompressed save
	uncompressed := make([]byte, len(s.body.Buffer))
	copy(uncompressed, s.body.Buffer)

	// Empty out p.body as we will now write the header / compressed data
	s.body = slicewriteseek.New()

	err = s.serializeHeader(save.Header)
	if err != nil {
		return err
	}

	err = s.compressBody(uncompressed)
	if err != nil {
		return err
	}

	_, err = s.w.Write(s.body.Buffer)
	if err != nil {
		return err
	}

	return nil
}

func (s *Serializer) serializeBody(save *Save) error {
	// Write placeholder length.
	err := s.writeInt32(0)
	if err != nil {
		return err
	}

	m := s.measure()

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

	err = s.writeLen(m(), 0)
	if err != nil {
		return err
	}

	return nil
}

func (s *Serializer) serializeObjects(save *Save) error {
	objectCount := len(save.Components) + len(save.Entities)
	err := s.writeInt32(int32(objectCount))
	if err != nil {
		return err
	}

	for _, e := range save.Entities {
		err = s.writeInt32(int32(EntityType))
		if err != nil {
			return err
		}

		err = e.serialize(s)
		if err != nil {
			return err
		}
	}

	for _, c := range save.Components {
		err = s.writeInt32(int32(ComponentType))
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

func (s *Serializer) serializeObjectData(save *Save) error {
	objectCount := len(save.Components) + len(save.Entities)
	err := s.writeInt32(int32(objectCount))
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

func (s *Serializer) serializeCollectedObjects(save *Save) error {
	err := s.writeInt32(int32(len(save.CollectedObjects)))
	if err != nil {
		return err
	}

	for _, c := range save.CollectedObjects {
		err = s.writeString(c.LevelName)
		if err != nil {
			return err
		}

		err = s.writeString(c.PathName)
		if err != nil {
			return err
		}
	}

	return nil
}
