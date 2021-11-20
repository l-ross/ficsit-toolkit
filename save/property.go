package save

import "fmt"

type Property struct {
	Name          string       `json:"name"`
	Type          PropertyType `json:"type"`
	Index         int32        `json:"index"`
	PropertyValue `json:"value"`
}

type PropertyType string

const (
	ArrayPropertyType     PropertyType = "ArrayProperty"
	BoolPropertyType      PropertyType = "BoolProperty"
	BytePropertyType      PropertyType = "ByteProperty"
	DoublePropertyType    PropertyType = "DoubleProperty"
	EnumPropertyType      PropertyType = "EnumProperty"
	FloatPropertyType     PropertyType = "FloatProperty"
	Int8PropertyType      PropertyType = "Int8Property"
	Int64PropertyType     PropertyType = "Int64Property"
	InterfacePropertyType PropertyType = "InterfaceProperty"
	IntPropertyType       PropertyType = "IntProperty"
	MapPropertyType       PropertyType = "MapProperty"
	NamePropertyType      PropertyType = "NameProperty"
	ObjectPropertyType    PropertyType = "ObjectProperty"
	StringPropertyType    PropertyType = "StrProperty"
	StructPropertyType    PropertyType = "StructProperty"
	TextPropertyType      PropertyType = "TextProperty"
)

type PropertyValue interface {
	// parse the property value
	//
	// If inner is true then the property value is inside an ArrayProperty or MapProperty.
	// In some cases this can change the format of the property value.
	parse(p *parser, inner bool) error

	// serialize the property value and return the length of the property data.
	// Some PropertyValue implementations always return a zero length as they aren't
	// always set.
	//
	// If inner is true then the property value is inside an ArrayProperty or MapProperty.
	// In some cases this can change the format of the property value.
	serialize(s *serializer, inner bool) (int32, error)
}

func (p *parser) parseProperties() ([]*Property, error) {
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

func (p *parser) parseProperty() (*Property, error) {
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

	// Value length
	_, err = p.readInt32()
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
		return nil, fmt.Errorf("unknown property type %s", propType)
	}

	prop.PropertyValue = newPropValue()

	err = prop.PropertyValue.parse(p, false)
	if err != nil {
		return nil, err
	}

	return prop, nil
}

func (s *serializer) serializeProperties(props []*Property) error {
	for _, prop := range props {
		err := s.serializeProperty(prop)
		if err != nil {
			return err
		}
	}

	err := s.writeNoneProp()
	if err != nil {
		return err
	}

	return nil
}

func (s *serializer) serializeProperty(prop *Property) error {
	err := s.writeString(prop.Name)
	if err != nil {
		return err
	}

	err = s.writeString(string(prop.Type))
	if err != nil {
		return err
	}

	// Write placeholder length and record position.
	lenPos := s.body.Index
	err = s.writeInt32(0)
	if err != nil {
		return err
	}

	err = s.writeInt32(prop.Index)
	if err != nil {
		return err
	}

	l, err := prop.serialize(s, false)
	if err != nil {
		return err
	}

	// Write length at the recorded position.
	err = s.writeLen(l, lenPos)
	if err != nil {
		return err
	}

	return nil
}
