package save

func (p *parser) parseComponent() (*Component, error) {
	c := &Component{}

	var err error
	c.TypePath, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.RootObject, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.InstanceName, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.ParentEntityName, err = p.readString()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (p *parser) serializeComponent(c *Component) error {
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
