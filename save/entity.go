package save

type Entity struct {
	TypePath         string             `json:"typePath"`
	RootObject       string             `json:"rootObject"`
	InstanceName     string             `json:"instanceName"`
	NeedTransform    int32              `json:"needTransform"`
	Rotation         []float32          `json:"rotation"`
	Position         []float32          `json:"position"`
	Scale            []float32          `json:"scale"`
	WasPlacedInLevel int32              `json:"wasPlacedInLevel"`
	ParentObjectRoot string             `json:"parentObjectRoot"`
	ParentObjectName string             `json:"parentObjectName"`
	References       []*ObjectReference `json:"component"`
	Properties       []*Property        `json:"properties"`
	Extra            *Extra             `json:"extras"`

	objectDataPos int64
	objectData    []byte
}

func (e *Entity) parse(p *Parser) error {
	var err error
	e.TypePath, err = p.readString()
	if err != nil {
		return err
	}

	e.RootObject, err = p.readString()
	if err != nil {
		return err
	}

	e.InstanceName, err = p.readString()
	if err != nil {
		return err
	}

	e.NeedTransform, err = p.readInt32()
	if err != nil {
		return err
	}

	e.Rotation, err = p.readFloat32Array(4)
	if err != nil {
		return err
	}

	e.Position, err = p.readFloat32Array(3)
	if err != nil {
		return err
	}

	e.Scale, err = p.readFloat32Array(3)
	if err != nil {
		return err
	}

	e.WasPlacedInLevel, err = p.readInt32()
	if err != nil {
		return err
	}

	return nil
}

func (e *Entity) parseData(p *Parser) error {
	dataSize, err := p.readInt32()
	if err != nil {
		return err
	}

	// Capture all bytes if debug is enabled.
	if debug {
		e.objectDataPos = p.body.Index - 4
		e.objectData, err = p.debugReadDataChunk(dataSize)
	}

	m := p.measure()

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

	e.References = make([]*ObjectReference, childCount)

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

		e.References[i] = o
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
	extraLen := dataSize - m()
	if extraLen > 0 {
		e.Extra = getExtra(e.TypePath)(extraLen)
		err = e.Extra.Value.parse(p)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Entity) serialize(s *Serializer) error {
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

func (e *Entity) serializeData(s *Serializer) error {
	// Write placeholder length
	lenPos := s.body.Index
	err := s.writeInt32(0)
	if err != nil {
		return err
	}

	m := s.measure()

	err = s.writeString(e.ParentObjectRoot)
	if err != nil {
		return err
	}

	err = s.writeString(e.ParentObjectName)
	if err != nil {
		return err
	}

	err = s.writeInt32(int32(len(e.References)))
	if err != nil {
		return err
	}

	for _, c := range e.References {
		err = s.writeString(c.LevelName)
		if err != nil {
			return err
		}

		err = s.writeString(c.PathName)
		if err != nil {
			return err
		}
	}

	err = s.serializeProperties(e.Properties)
	if err != nil {
		return err
	}

	// Write ExtraCount
	// Is this always zero?
	err = s.writeInt32(0)
	if err != nil {
		return err
	}

	if e.Extra != nil {
		err = e.Extra.Value.serialize(s)
		if err != nil {
			return err
		}
	}

	// Update length
	err = s.writeLen(m(), lenPos)
	if err != nil {
		return err
	}

	return nil
}
