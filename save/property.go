package save

import "fmt"

func (p *Parser) parseProperties() ([]*Property, error) {
	props := make([]*Property, 0)

	for {
		prop, err := p.parseProperty()
		if err != nil {
			return nil, err
		}
		if prop == nil {
			// Reached the end of the property list.
			break
		}

		props = append(props, prop)
	}

	return props, nil
}

func (p *Parser) parseProperty() (*Property, error) {
	var err error
	prop := &Property{}

	prop.Name, err = p.readString()
	if err != nil {
		return nil, err
	}

	if prop.Name == "None" {
		// End of the property list.
		return nil, nil
	}

	propType, err := p.readString()
	if err != nil {
		return nil, err
	}
	prop.Type = PropertyType(propType)

	prop.ValueLen, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	prop.Index, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	var newPropValue newPropValueFunc

	switch prop.Type {
	case ArrayPropertyType:
		newPropValue = newArrayPropertyValue
	case BoolPropertyType:
		newPropValue = newBoolPropertyValue
	case BytePropertyType:
		newPropValue = newBytePropertyValue
	case DoublePropertyType:
		newPropValue = newDoublePropertyValue
	case EnumPropertyType:
		newPropValue = newEnumPropertyValue
	case FloatPropertyType:
		newPropValue = newFloatPropertyValue
	case Int8PropertyType:
		newPropValue = newInt8PropertyValue
	case Int64PropertyType:
		newPropValue = newInt64PropertyValue
	case IntPropertyType:
		newPropValue = newIntPropertyValue
	case InterfacePropertyType:
		newPropValue = newInterfacePropertyValue
	case NamePropertyType:
		newPropValue = newNamePropertyValue
	case MapPropertyType:
		newPropValue = newMapPropertyValue
	case ObjectPropertyType:
		newPropValue = newObjectPropertyValue
	case StringPropertyType:
		newPropValue = newStringPropertyValue
	case StructPropertyType:
		newPropValue = newStructPropertyValue
	case TextPropertyType:
		newPropValue = newTextPropertyValue
	default:
		// TODO: Possibly have a UnknownPropertyType where we just store the value as a byte slice.
		return nil, fmt.Errorf("unknown property type %s", propType)
	}

	prop.PropertyValue = newPropValue()

	err = prop.PropertyValue.parse(p, false)
	if err != nil {
		return nil, err
	}

	return prop, nil
}

func (s *Save) serializeProperty(p *Property) error {
	err := s.writeString(p.Name)
	if err != nil {
		return err
	}

	err = s.writeString(string(p.Type))
	if err != nil {
		return err
	}

	// TODO: Do we want to write the actual length?
	//  How often is it actually set by the game?
	err = s.writeInt32(0)
	if err != nil {
		return err
	}

	err = s.writeInt32(p.Index)
	if err != nil {
		return err
	}

	err = p.serialize(s, false)
	if err != nil {
		return err
	}

	err = s.writeNoneProp()
	if err != nil {
		return err
	}

	return nil
}
