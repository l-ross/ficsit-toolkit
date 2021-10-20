package save

import "io"

func (p *parser) parseEntity() (*Entity, error) {
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

func (p *parser) serializeEntity(e *Entity) error {
	err := p.writeString(e.TypePath)
	if err != nil {
		return err
	}

	err = p.writeString(e.RootObject)
	if err != nil {
		return err
	}

	err = p.writeString(e.InstanceName)
	if err != nil {
		return err
	}

	err = p.writeInt32(e.NeedTransform)
	if err != nil {
		return err
	}

	err = p.writeFloat32Array(e.Rotation)
	if err != nil {
		return err
	}

	err = p.writeFloat32Array(e.Position)
	if err != nil {
		return err
	}

	err = p.writeFloat32Array(e.Scale)
	if err != nil {
		return err
	}

	err = p.writeInt32(e.WasPlacedInLevel)
	if err != nil {
		return err
	}

	return nil
}

func (p *parser) parseEntityData(e *Entity) error {
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

	// If we have extra data then parse it.
	extraLen := (int32(e.offset) + e.len) - int32(p.body.Index)
	if extraLen > 0 {
		e.Extra = getExtra(e.TypePath)(extraLen)
		err = e.Extra.Value.parse(p)
		if err != nil {
			return err
		}
	}

	return nil
}
