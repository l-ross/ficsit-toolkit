package save

type Component struct {
	TypePath         string      `json:"typePath"`
	RootObject       string      `json:"rootObject"`
	InstanceName     string      `json:"instanceName"`
	ParentEntityName string      `json:"parentEntityName"`
	Properties       []*Property `json:"properties"`

	objectDataPos int64
	objectData    []byte
}

func (c *Component) parse(p *Parser) error {
	var err error
	c.TypePath, err = p.readString()
	if err != nil {
		return err
	}

	c.RootObject, err = p.readString()
	if err != nil {
		return err
	}

	c.InstanceName, err = p.readString()
	if err != nil {
		return err
	}

	c.ParentEntityName, err = p.readString()
	if err != nil {
		return err
	}

	return nil
}

func (c *Component) parseData(p *Parser) error {
	// Data length
	dataSize, err := p.readInt32()
	if err != nil {
		return err
	}

	// Capture all bytes if debug is enabled.
	if debug {
		c.objectDataPos = p.body.Index - 4
		c.objectData, err = p.debugReadDataChunk(dataSize)
	}

	c.Properties, err = p.parseProperties()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = p.readInt32()
	if err != nil {
		return err
	}

	return nil
}

func (c *Component) serialize(s *Serializer) error {
	err := s.writeString(c.TypePath)
	if err != nil {
		return err
	}

	err = s.writeString(c.RootObject)
	if err != nil {
		return err
	}

	err = s.writeString(c.InstanceName)
	if err != nil {
		return err
	}

	err = s.writeString(c.ParentEntityName)
	if err != nil {
		return err
	}

	return nil
}

func (c *Component) serializeData(s *Serializer) error {
	// Write placeholder length
	lenPos := s.body.Index
	err := s.writeInt32(0)
	if err != nil {
		return err
	}

	m := s.measure()

	err = s.serializeProperties(c.Properties)
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	err = s.writeInt32(0)
	if err != nil {
		return err
	}

	err = s.writeLen(m(), lenPos)
	if err != nil {
		return err
	}

	return nil
}
