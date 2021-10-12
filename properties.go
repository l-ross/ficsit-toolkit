package satisfactorysave

import (
	"bytes"
	"fmt"
)

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

type Property struct {
	Name     string        `json:"name"`
	Type     PropertyType  `json:"type"`
	ValueLen int32         `json:"value_len"`
	Index    int32         `json:"index"`
	Value    PropertyValue `json:"value"`
}

type PropertyValue interface {
	Parse(r *bytes.Reader) error
}

//
// ArrayProperty
//

//
// BoolProperty
//

type BoolPropertyValue bool

func (p *Property) GetBoolValue() (bool, error) {
	if v, ok := p.Value.(*BoolPropertyValue); ok {
		return bool(*v), nil
	}

	return false, fmt.Errorf("wrong type %s", p.Type)
}

func (v *BoolPropertyValue) Parse(r *bytes.Reader) error {
	b, err := readBool(r)
	if err != nil {
		return err
	}
	*v = BoolPropertyValue(b)
	return nil
}

//
// ByteProperty
//

type BytePropertyValue struct {
	Type  string
	Value []byte
}

func (p *Property) GetByteValue() ([]byte, error) {
	if v, ok := p.Value.(*BytePropertyValue); ok {
		return v.Value, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *BytePropertyValue) Parse(r *bytes.Reader) error {

	var err error
	v.Type, err = readString(r)
	if err != nil {
		return err
	}

	switch v.Type {
	case "None":
		b, err := readByte(r)
		if err != nil {
			return err
		}

		v.Value = []byte{b}
	default:
		bytesLen, err := readInt32(r)
		if err != nil {
			return err
		}

		v.Value, err = readBytes(r, bytesLen)
		if err != nil {
			return err
		}
	}

	err = nextByteIsNull(r)
	if err != nil {
		return err
	}

	return nil
}

//
// DoubleProperty
//

type DoublePropertyValue float64

func (p *Property) GetDoubleValue() (float64, error) {
	if v, ok := p.Value.(*DoublePropertyValue); ok {
		return float64(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *DoublePropertyValue) Parse(r *bytes.Reader) error {
	// TODO: What is this byte for?
	err := nextByteIsNull(r)
	if err != nil {
		return err
	}

	f, err := readFloat64(r)
	if err != nil {
		return err
	}

	*v = DoublePropertyValue(f)
	return nil
}

//
// FloatProperty
//

type FloatPropertyValue float32

func (p *Property) GetFloatValue() (float32, error) {
	if v, ok := p.Value.(*FloatPropertyValue); ok {
		return float32(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *FloatPropertyValue) Parse(r *bytes.Reader) error {
	// TODO: What is this byte for?
	err := nextByteIsNull(r)
	if err != nil {
		return err
	}

	f, err := readFloat32(r)
	if err != nil {
		return err
	}

	*v = FloatPropertyValue(f)
	return nil
}

//
// Int8Property
//

type Int8PropertyValue int8

func (p *Property) GetInt8Value() (int8, error) {
	if v, ok := p.Value.(*Int8PropertyValue); ok {
		return int8(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *Int8PropertyValue) Parse(r *bytes.Reader) error {
	// TODO: What is this byte for?
	err := nextByteIsNull(r)
	if err != nil {
		return err
	}

	f, err := readInt8(r)
	if err != nil {
		return err
	}
	*v = Int8PropertyValue(f)
	return nil
}

//
// Int64Property
//

type Int64PropertyValue int8

func (p *Property) GetInt64Value() (int64, error) {
	if v, ok := p.Value.(*Int64PropertyValue); ok {
		return int64(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *Int64PropertyValue) Parse(r *bytes.Reader) error {
	// TODO: What is this byte for?
	err := nextByteIsNull(r)
	if err != nil {
		return err
	}

	f, err := readInt64(r)
	if err != nil {
		return err
	}
	*v = Int64PropertyValue(f)
	return nil
}

//
// InterfaceProperty
//

type InterfacePropertyValue struct {
	LevelName string
	PathName  string
}

func (p *Property) GetInterfaceValue() (*InterfacePropertyValue, error) {
	if v, ok := p.Value.(*InterfacePropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *InterfacePropertyValue) Parse(r *bytes.Reader) error {
	// TODO: What is this byte for?
	err := nextByteIsNull(r)
	if err != nil {
		return err
	}

	v.LevelName, err = readString(r)
	if err != nil {
		return err
	}

	v.PathName, err = readString(r)
	if err != nil {
		return err
	}

	return nil
}

//
// IntProperty
//

type IntPropertyValue int32

func (p *Property) GetIntValue() (int32, error) {
	if v, ok := p.Value.(*IntPropertyValue); ok {
		return int32(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *IntPropertyValue) Parse(r *bytes.Reader) error {
	// TODO: What is this byte for?
	err := nextByteIsNull(r)
	if err != nil {
		return err
	}

	i, err := readInt32(r)
	if err != nil {
		return err
	}
	*v = IntPropertyValue(i)
	return nil
}

//
// MapProperty
//

//type MapPropertyValue struct {
//}
//
//func (v *MapPropertyValue) Parse(r *bytes.Reader) error {
//	keyType, err := readString(r)
//	if err != nil {
//		return err
//	}
//
//	valueType, err := readString(r)
//	if err != nil {
//		return err
//	}
//
//	// TODO: What is this?
//	_, err = readInt32(r)
//	if err != nil {
//		return err
//	}
//
//	// TODO: What is this byte for?
//	err = nextByteIsNull(r)
//	if err != nil {
//		return err
//	}
//
//	itemCount, err := readInt32(r)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//
// NameProperty
//

type NamePropertyValue string

func (p *Property) GetNameValue() (string, error) {
	if v, ok := p.Value.(*NamePropertyValue); ok {
		return string(*v), nil
	}

	return "", fmt.Errorf("wrong type %s", p.Type)
}

func (v *NamePropertyValue) Parse(r *bytes.Reader) error {
	// TODO: What is this byte for?
	err := nextByteIsNull(r)
	if err != nil {
		return err
	}

	s, err := readString(r)
	if err != nil {
		return err
	}
	*v = NamePropertyValue(s)
	return nil
}

//
// ObjectProperty
//

type ObjectPropertyValue struct {
	LevelName string
	PathName  string
}

func (p *Property) GetObjectValue() (*ObjectPropertyValue, error) {
	if v, ok := p.Value.(*ObjectPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *ObjectPropertyValue) Parse(r *bytes.Reader) error {
	// TODO: What is this byte for?
	err := nextByteIsNull(r)
	if err != nil {
		return err
	}

	v.LevelName, err = readString(r)
	if err != nil {
		return err
	}

	v.PathName, err = readString(r)
	if err != nil {
		return err
	}

	return nil
}

//
// StringProperty
//

type StringPropertyValue string

func (p *Property) GetStringValue() (string, error) {
	if v, ok := p.Value.(*StringPropertyValue); ok {
		return string(*v), nil
	}

	return "", fmt.Errorf("wrong type %s", p.Type)
}

func (v *StringPropertyValue) Parse(r *bytes.Reader) error {
	// TODO: What is this byte for?
	err := nextByteIsNull(r)
	if err != nil {
		return err
	}

	s, err := readString(r)
	if err != nil {
		return err
	}
	*v = StringPropertyValue(s)
	return nil
}

//
// StructProperty
//

//
// TextProperty
//
