package save

import "io"

func (p *Parser) parseEntity() (*Entity, error) {
	e := &Entity{}

	var err error
	e.TypePath, err = p.readString()
	if err != nil {
		return nil, err
	}

	e.RootObject, err = p.readString()
	if err != nil {
		return nil, err
	}

	e.InstanceName, err = p.readString()
	if err != nil {
		return nil, err
	}

	e.NeedTransform, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	e.Rotation, err = p.readFloat32Array(4)
	if err != nil {
		return nil, err
	}

	e.Position, err = p.readFloat32Array(3)
	if err != nil {
		return nil, err
	}

	e.Scale, err = p.readFloat32Array(3)
	if err != nil {
		return nil, err
	}

	e.WasPlacedInLevel, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *Save) serializeEntity(e *Entity) error {
	err := s.writeString(e.TypePath)
	if err != nil {
		return err
	}

	err = s.writeString(e.RootObject)
	if err != nil {
		return err
	}

	err = s.writeString(e.InstanceName)
	if err != nil {
		return err
	}

	err = s.writeInt32(e.NeedTransform)
	if err != nil {
		return err
	}

	err = s.writeFloat32Array(e.Rotation)
	if err != nil {
		return err
	}

	err = s.writeFloat32Array(e.Position)
	if err != nil {
		return err
	}

	err = s.writeFloat32Array(e.Scale)
	if err != nil {
		return err
	}

	err = s.writeInt32(e.WasPlacedInLevel)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) parseEntityData(e *Entity) error {
	_, err := p.body.Seek(e.offset, io.SeekStart)
	if err != nil {
		return err
	}

	e.ParentObjectRoot, err = p.readString()
	if err != nil {
		return err
	}

	e.ParentObjectName, err = p.readString()
	if err != nil {
		return err
	}

	childCount, err := p.readInt32()
	if err != nil {
		return err
	}

	e.Children = make([]*ObjectReference, childCount)

	for i := int32(0); i < childCount; i++ {
		o := &ObjectReference{}

		o.LevelName, err = p.readString()
		if err != nil {
			return err
		}

		o.PathName, err = p.readString()
		if err != nil {
			return err
		}

		e.Children[i] = o
	}

	e.Properties, err = p.parseProperties()
	if err != nil {
		return err
	}

	// ExtraCount
	// Is this ever not zero?
	_, err = p.readInt32()
	if err != nil {
		return err
	}

	// Parse extra data
	// If we have a specific handler for extra data then use it, otherwise
	// treat as unknown data.
	extraFunc := hasExtra(e.TypePath)
	if extraFunc != nil {
		e.Extra = extraFunc()
		err = e.Extra.Value.parse(p)
		if err != nil {
			return err
		}
	} else {
		expPos := e.offset + int64(e.len)
		if rem := expPos - p.body.Index; rem > 0 {
			e.Extra = newUnknownExtra(int32(rem))
			err = e.Extra.Value.parse(p)
			if err != nil {
				return err
			}
		}
	}

	return nil
}