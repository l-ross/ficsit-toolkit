package satisfactorysave

import (
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
	// Parse the property value
	//
	// If inner is true then the property value is a child element of another property value
	// e.g. a StringProperty inside an ArrayProperty. In some cases this can change the format
	// of the property value slightly.
	parse(p *Parser, inner bool) error
}

//
// ArrayProperty
//

type ArrayPropertyValue struct {
	ValueType PropertyType
	Values    []PropertyValue
}

func (p *Property) GetArrayValue() (*ArrayPropertyValue, error) {
	if v, ok := p.Value.(*ArrayPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *ArrayPropertyValue) parse(p *Parser, inner bool) error {
	valueType, err := p.readString()
	if err != nil {
		return err
	}
	v.ValueType = PropertyType(valueType)

	// TODO: What is this?
	_, err = p.readByte()
	if err != nil {
		return err
	}

	elemCount, err := p.readInt32()
	if err != nil {
		return err
	}

	v.Values = make([]PropertyValue, elemCount)

	var newInner func() PropertyValue

	switch v.ValueType {
	case InterfacePropertyType:
		newInner = func() PropertyValue {
			return &InterfacePropertyValue{}
		}
	case ObjectPropertyType:
		newInner = func() PropertyValue {
			return &ObjectPropertyValue{}
		}
	default:
		return fmt.Errorf("unsupported array type %s", v.ValueType)
	}

	for i := int32(0); i < elemCount; i++ {
		inner := newInner()

		err := inner.parse(p, true)
		if err != nil {
			return err
		}

		v.Values[i] = inner
	}

	return nil
}

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

func (v *BoolPropertyValue) parse(p *Parser, inner bool) error {
	b, err := p.readBool()
	if err != nil {
		return err
	}

	err = p.nextByteIsNull()
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

func (v *BytePropertyValue) parse(p *Parser, inner bool) error {
	var err error
	v.Type, err = p.readString()
	if err != nil {
		return err
	}

	switch v.Type {
	case "None":
		b, err := p.readByte()
		if err != nil {
			return err
		}

		v.Value = []byte{b}
	default:
		bytesLen, err := p.readInt32()
		if err != nil {
			return err
		}

		v.Value, err = p.readBytes(bytesLen)
		if err != nil {
			return err
		}
	}

	err = p.nextByteIsNull()
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

func (v *DoublePropertyValue) parse(p *Parser, inner bool) error {
	// TODO: What is this byte for?
	err := p.nextByteIsNull()
	if err != nil {
		return err
	}

	f, err := p.readFloat64()
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

func (v *FloatPropertyValue) parse(p *Parser, inner bool) error {
	// TODO: What is this byte for?
	err := p.nextByteIsNull()
	if err != nil {
		return err
	}

	f, err := p.readFloat32()
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

func (v *Int8PropertyValue) parse(p *Parser, inner bool) error {
	// TODO: What is this byte for?
	err := p.nextByteIsNull()
	if err != nil {
		return err
	}

	f, err := p.readInt8()
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

func (v *Int64PropertyValue) parse(p *Parser, inner bool) error {
	// TODO: What is this byte for?
	err := p.nextByteIsNull()
	if err != nil {
		return err
	}

	f, err := p.readInt64()
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

func (v *InterfacePropertyValue) parse(p *Parser, inner bool) error {
	var err error

	if !inner {
		// TODO: What is this byte for?
		err := p.nextByteIsNull()
		if err != nil {
			return err
		}
	}

	v.LevelName, err = p.readString()
	if err != nil {
		return err
	}

	v.PathName, err = p.readString()
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

func (v *IntPropertyValue) parse(p *Parser, inner bool) error {
	if !inner {
		// TODO: What is this byte for?
		err := p.nextByteIsNull()
		if err != nil {
			return err
		}
	}

	i, err := p.readInt32()
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
//func (v *MapPropertyValue) parse(r *bytes.Reader) error {
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

func (v *NamePropertyValue) parse(p *Parser, inner bool) error {
	// TODO: What is this byte for?
	err := p.nextByteIsNull()
	if err != nil {
		return err
	}

	s, err := p.readString()
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

func (v *ObjectPropertyValue) parse(p *Parser, inner bool) error {
	var err error

	if !inner {
		// TODO: What is this byte for?
		err := p.nextByteIsNull()
		if err != nil {
			return err
		}
	}

	v.LevelName, err = p.readString()
	if err != nil {
		return err
	}

	v.PathName, err = p.readString()
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

func (v *StringPropertyValue) parse(p *Parser, inner bool) error {
	// TODO: What is this byte for?
	err := p.nextByteIsNull()
	if err != nil {
		return err
	}

	s, err := p.readString()
	if err != nil {
		return err
	}
	*v = StringPropertyValue(s)
	return nil
}

//
// StructProperty
//

type StructValue interface {
	parse(p *Parser) error
}

type StructPropertyValue struct {
	GUID  []int32
	Type  StructType
	Value StructValue
}

func (p *Property) GetStructValue() (*StructPropertyValue, error) {
	if v, ok := p.Value.(*StructPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *StructPropertyValue) parse(p *Parser, inner bool) error {
	structType, err := p.readString()
	if err != nil {
		return err
	}
	v.Type = StructType(structType)

	fmt.Println(structType)

	v.GUID, err = p.readInt32Array(4)
	if err != nil {
		return err
	}

	// TODO: What is this byte for?
	err = p.nextByteIsNull()
	if err != nil {
		return err
	}

	switch v.Type {
	case BoxStructType:
		v.Value = &BoxStruct{}
	default:
		v.Value = &ArbitraryStruct{}
	}

	err = v.Value.parse(p)
	if err != nil {
		return err
	}

	return nil
}

//
// TextProperty
//
