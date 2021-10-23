package save

type Component struct {
	TypePath         string      `json:"typePath"`
	RootObject       string      `json:"rootObject"`
	InstanceName     string      `json:"instanceName"`
	ParentEntityName string      `json:"parentEntityName"`
	Properties       []*Property `json:"properties"`
}

func (c *Component) parse(p *parser) error {
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

func (c *Component) parseData(p *parser) error {
	// Data length
	_, err := p.readInt32()
	if err != nil {
		return err
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

func (c *Component) serialize(p *parser) error {
	err := p.writeString(c.TypePath)
	if err != nil {
		return err
	}

	err = p.writeString(c.RootObject)
	if err != nil {
		return err
	}

	err = p.writeString(c.InstanceName)
	if err != nil {
		return err
	}

	err = p.writeString(c.ParentEntityName)
	if err != nil {
		return err
	}

	return nil
}

func (c *Component) serializeData(p *parser) error {
	// Write placeholder length
	lenPos := p.body.Index
	err := p.writeInt32(0)
	if err != nil {
		return err
	}

	m := p.measure()

	err = p.serializeProperties(c.Properties)
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	err = p.writeInt32(0)
	if err != nil {
		return err
	}

	err = p.writeLen(m(), lenPos)
	if err != nil {
		return err
	}

	return nil
}
