package save

import (
	"io"

	"github.com/ViRb3/slicewriteseek"
)

type serializer struct {
	w io.Writer

	// Body of the save file.
	body *slicewriteseek.SliceWriteSeeker
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

func (s *serializer) serializeBody(save *Save) error {
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

func (s *serializer) serializeObjects(save *Save) error {
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

func (s *serializer) serializeObjectData(save *Save) error {
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

func (s *serializer) serializeCollectedObjects(save *Save) error {
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
