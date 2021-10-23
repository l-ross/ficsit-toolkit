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
	Children         []*ObjectReference `json:"children"`
	Properties       []*Property        `json:"properties"`
	Extra            *Extra             `json:"extras"`
}

func (e *Entity) parse(p *parser) error {
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

func (e *Entity) parseData(p *parser) error {
	dataSize, err := p.readInt32()
	if err != nil {
		return err
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

func (e *Entity) serialize(p *parser) error {
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

func (e *Entity) serializeData(p *parser) error {
	// Write placeholder length
	lenPos := p.body.Index
	err := p.writeInt32(0)
	if err != nil {
		return err
	}

	m := p.measure()

	err = p.writeString(e.ParentObjectRoot)
	if err != nil {
		return err
	}

	err = p.writeString(e.ParentObjectName)
	if err != nil {
		return err
	}

	err = p.writeInt32(int32(len(e.Children)))
	if err != nil {
		return err
	}

	for _, c := range e.Children {
		err = p.writeString(c.LevelName)
		if err != nil {
			return err
		}

		err = p.writeString(c.PathName)
		if err != nil {
			return err
		}
	}

	err = p.serializeProperties(e.Properties)
	if err != nil {
		return err
	}

	// Write ExtraCount
	// Is this always zero?
	err = p.writeInt32(0)
	if err != nil {
		return err
	}

	if e.Extra != nil {
		err = e.Extra.Value.serialize(p)
		if err != nil {
			return err
		}
	}

	// Update length
	err = p.writeLen(m(), lenPos)
	if err != nil {
		return err
	}

	return nil
}
