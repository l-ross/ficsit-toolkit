// Package property provides handling for the different property types that may be encountered within
// a Satisfactory save file.
package property

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/save/data"
)

// A Property within a save.Component or save.Entity
type Property struct {
	Name  string `json:"name"`
	Type  Type   `json:"type"`
	Index int32  `json:"index"`

	// The Value of the Property.
	//
	// Accessing the value can be achieved by calling the appropriate Get method
	// on the Property based on its Type.
	Value Value `json:"value"`
}

type Type string

type Value interface {
	// parse the property Value
	//
	// If inner is true then the property Value is inside an ArrayProperty or MapProperty.
	// In some cases this can change the format of the property Value.
	parse(d *data.Data, inner bool) error

	// serialize the property Value and return the length of the property data.
	// Some PropertyValue implementations always return a zero length as they aren't
	// always set.
	//
	// If inner is true then the property Value is inside an ArrayProperty or MapProperty.
	// In some cases this can change the format of the property Value.
	serialize(d *data.Data, inner bool) (int32, error)
}

type newPropValueFunc func() Value

func ParseProperties(d *data.Data) ([]*Property, error) {
	props := make([]*Property, 0)

	for {
		prop, err := ParseProperty(d)
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

func ParseProperty(d *data.Data) (*Property, error) {
	var err error
	prop := &Property{}

	prop.Name, err = d.ReadString()
	if err != nil {
		return nil, err
	}

	if prop.Name == "None" {
		// End of the property list.
		return nil, nil
	}

	propType, err := d.ReadString()
	if err != nil {
		return nil, err
	}
	prop.Type = Type(propType)

	// Value length
	_, err = d.ReadInt32()
	if err != nil {
		return nil, err
	}

	prop.Index, err = d.ReadInt32()
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

	prop.Value = newPropValue()

	err = prop.Value.parse(d, false)
	if err != nil {
		return nil, err
	}

	return prop, nil
}

func SerializeProperties(props []*Property, d *data.Data) error {
	for _, prop := range props {
		err := prop.SerializeProperty(d)
		if err != nil {
			return err
		}
	}

	err := d.WriteNoneProp()
	if err != nil {
		return err
	}

	return nil
}

func (p *Property) SerializeProperty(d *data.Data) error {
	err := d.WriteString(p.Name)
	if err != nil {
		return err
	}

	err = d.WriteString(string(p.Type))
	if err != nil {
		return err
	}

	// Write placeholder length and record position.
	lenPos := d.Index()
	err = d.WriteInt32(0)
	if err != nil {
		return err
	}

	err = d.WriteInt32(p.Index)
	if err != nil {
		return err
	}

	l, err := p.Value.serialize(d, false)
	if err != nil {
		return err
	}

	// Write length at the recorded position.
	err = d.WriteLen(l, lenPos)
	if err != nil {
		return err
	}

	return nil
}
