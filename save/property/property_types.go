package property

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/l-ross/ficsit-toolkit/save/data"
)

const (
	ArrayPropertyType     Type = "ArrayProperty"
	BoolPropertyType      Type = "BoolProperty"
	BytePropertyType      Type = "ByteProperty"
	DoublePropertyType    Type = "DoubleProperty"
	EnumPropertyType      Type = "EnumProperty"
	FloatPropertyType     Type = "FloatProperty"
	Int8PropertyType      Type = "Int8Property"
	Int64PropertyType     Type = "Int64Property"
	InterfacePropertyType Type = "InterfaceProperty"
	IntPropertyType       Type = "IntProperty"
	MapPropertyType       Type = "MapProperty"
	NamePropertyType      Type = "NameProperty"
	ObjectPropertyType    Type = "ObjectProperty"
	StringPropertyType    Type = "StrProperty"
	StructPropertyType    Type = "StructProperty"
	TextPropertyType      Type = "TextProperty"
)

//
// ArrayProperty
//

type ArrayPropertyValue struct {
	ValueType Type    `json:"value_type,omitempty"`
	Values    []Value `json:"values,omitempty"`

	//
	// The following fields are only set if the ValueType is StructProperty
	//

	StructName      string
	StructBytes     []byte
	StructInnerType string
	StructGUID      []int32
}

func newArrayPropertyValue() Value {
	return &ArrayPropertyValue{}
}

func (p *Property) GetArrayValue() (*ArrayPropertyValue, error) {
	if v, ok := p.Value.(*ArrayPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *ArrayPropertyValue) GetStructValues() ([]*StructPropertyValue, error) {
	if v.ValueType != StructPropertyType {
		return nil, fmt.Errorf("wrong Value type %s", v.ValueType)
	}

	values := make([]*StructPropertyValue, len(v.Values))

	for i, pv := range v.Values {
		values[i] = pv.(*StructPropertyValue)
	}

	return values, nil
}

func (v *ArrayPropertyValue) parse(d *data.Data, inner bool) error {
	valueType, err := d.ReadString()
	if err != nil {
		return err
	}
	v.ValueType = Type(valueType)

	// UNKNOWN_DATA
	_, err = d.ReadByte()
	if err != nil {
		return err
	}

	elemCount, err := d.ReadInt32()
	if err != nil {
		return err
	}

	// Special handling for BytePropertyType.
	// Rather than creating 1 entry in the array per byte we just want to parse all
	// the data in to a single BytePropertyValue.
	if v.ValueType == BytePropertyType {
		v.Values = make([]Value, 1)

		elem := &BytePropertyValue{
			len: elemCount,
		}
		err := elem.parse(d, true)
		if err != nil {
			return err
		}

		v.Values[0] = elem

		return nil
	}

	v.Values = make([]Value, elemCount)

	var newPropValue newPropValueFunc

	switch v.ValueType {
	case EnumPropertyType:
		newPropValue = newEnumPropertyValue
	case FloatPropertyType:
		newPropValue = newFloatPropertyValue
	case InterfacePropertyType:
		newPropValue = newInterfacePropertyValue
	case IntPropertyType:
		newPropValue = newIntPropertyValue
	case ObjectPropertyType:
		newPropValue = newObjectPropertyValue
	case StringPropertyType:
		newPropValue = newStringPropertyValue
	case StructPropertyType:
		// Name
		v.StructName, err = d.ReadString()
		if err != nil {
			return err
		}

		// Type
		// Is this always StructProperty?
		_, err = d.ReadString()
		if err != nil {
			return err
		}

		v.StructBytes, err = d.ReadBytes(8)
		if err != nil {
			return err
		}

		v.StructInnerType, err = d.ReadString()
		if err != nil {
			return err
		}

		v.StructGUID, err = d.ReadInt32Array(4)
		if err != nil {
			return err
		}

		err = d.NextByteIsNull()
		if err != nil {
			return err
		}

		newPropValue = func() Value {
			return &StructPropertyValue{
				Type: StructType(v.StructInnerType),
			}
		}
	case TextPropertyType:
		panic("TODO TextProperty")
	default:
		return fmt.Errorf("unsupported array type %s", v.ValueType)
	}

	for i := int32(0); i < elemCount; i++ {
		propVal := newPropValue()

		err := propVal.parse(d, true)
		if err != nil {
			return err
		}

		v.Values[i] = propVal
	}

	return nil
}

func (v *ArrayPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	err := d.WriteString(string(v.ValueType))
	if err != nil {
		return 0, err
	}

	// UNKNOWN_DATA
	err = d.WriteNull()
	if err != nil {
		return 0, err
	}

	m := d.Measure()

	// Special handling for BytePropertyType.
	// See Parse() for more detaild.
	if v.ValueType == BytePropertyType {
		if len(v.Values) != 1 {
			return 0, fmt.Errorf("arrayproperty containg byteproperty, expected 1 Value got %d", len(v.Values))
		}

		// Write placeholder element count.
		elemCountPos := d.Index()
		err = d.WriteInt32(0)
		if err != nil {
			return 0, err
		}

		// Write all byted.
		l, err := v.Values[0].serialize(d, true)
		if err != nil {
			return 0, err
		}

		// Update element count.
		err = d.WriteLen(l, elemCountPos)
		if err != nil {
			return 0, err
		}

		return m(), nil
	}

	err = d.WriteInt32(int32(len(v.Values)))
	if err != nil {
		return 0, err
	}

	if v.ValueType == StructPropertyType {
		err = d.WriteString(v.StructName)
		if err != nil {
			return 0, err
		}

		err = d.WriteString("StructProperty")
		if err != nil {
			return 0, err
		}

		err = d.WriteBytes(v.StructBytes)
		if err != nil {
			return 0, err
		}

		err = d.WriteString(v.StructInnerType)
		if err != nil {
			return 0, err
		}

		err = d.WriteInt32Array(v.StructGUID)
		if err != nil {
			return 0, err
		}

		err = d.WriteNull()
		if err != nil {
			return 0, err
		}
	}

	for _, pv := range v.Values {
		_, err := pv.serialize(d, true)
		if err != nil {
			return 0, err
		}
	}

	return m(), nil
}

//
// BoolProperty
//

type BoolPropertyValue bool

func newBoolPropertyValue() Value {
	v := BoolPropertyValue(false)
	return &v
}

func (p *Property) GetBoolValue() (bool, error) {
	if v, ok := p.Value.(*BoolPropertyValue); ok {
		return bool(*v), nil
	}

	return false, fmt.Errorf("wrong type %s", p.Type)
}

func (v *BoolPropertyValue) parse(d *data.Data, inner bool) error {
	b, err := d.ReadBool()
	if err != nil {
		return err
	}

	err = d.NextByteIsNull()
	if err != nil {
		return err
	}

	*v = BoolPropertyValue(b)
	return nil
}

func (v *BoolPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	var err error
	if *v {
		err = d.WriteBool(true)
	} else {
		err = d.WriteBool(false)
	}
	if err != nil {
		return 0, err
	}

	return 0, d.WriteNull()
}

//
// ByteProperty
//

type BytePropertyValue struct {
	Type  string `json:"text,omitempty"`
	Value []byte `json:"Value,omitempty"`

	len int32
}

func newBytePropertyValue() Value {
	return &BytePropertyValue{}
}

func (p *Property) GetByteValue() ([]byte, error) {
	if v, ok := p.Value.(*BytePropertyValue); ok {
		return v.Value, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *BytePropertyValue) parse(d *data.Data, inner bool) error {
	if inner {
		return v.parseInner(d)
	}

	var err error
	v.Type, err = d.ReadString()
	if err != nil {
		return err
	}

	err = d.NextByteIsNull()
	if err != nil {
		return err
	}

	switch v.Type {
	case "None":
		b, err := d.ReadByte()
		if err != nil {
			return err
		}

		v.Value = []byte{b}
	default:
		value, err := d.ReadString()
		if err != nil {
			return err
		}

		v.Value = []byte(value)
	}

	return nil
}

func (v *BytePropertyValue) parseInner(d *data.Data) error {
	// If len is 0 then we must be parsing a ByteProperty in the Value of a MapProperty.
	// Best guess is that it's only a single byte.
	if v.len == 0 {
		v.len = 1
	}

	v.Type = "None"

	var err error
	v.Value, err = d.ReadBytes(v.len)
	if err != nil {
		return err
	}

	return nil
}

func (v *BytePropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	if inner {
		return v.serializeInner(d)
	}

	err := d.WriteString(v.Type)
	if err != nil {
		return 0, err
	}

	err = d.WriteNull()
	if err != nil {
		return 0, err
	}

	m := d.Measure()

	switch v.Type {
	case "None":
		err = d.WriteByte(v.Value[0])
		if err != nil {
			return 0, err
		}
	default:
		err = d.WriteString(string(v.Value))
		if err != nil {
			return 0, err
		}
	}

	return m(), nil
}

func (v *BytePropertyValue) serializeInner(d *data.Data) (int32, error) {
	m := d.Measure()

	err := d.WriteBytes(v.Value)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// DoubleProperty
//

type DoublePropertyValue float64

func newDoublePropertyValue() Value {
	v := DoublePropertyValue(0)
	return &v
}

func (p *Property) GetDoubleValue() (float64, error) {
	if v, ok := p.Value.(*DoublePropertyValue); ok {
		return float64(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *DoublePropertyValue) parse(d *data.Data, inner bool) error {
	err := d.NextByteIsNull()
	if err != nil {
		return err
	}

	f, err := d.ReadFloat64()
	if err != nil {
		return err
	}

	*v = DoublePropertyValue(f)
	return nil
}

func (v *DoublePropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	err := d.WriteNull()
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat64(float64(*v))
	if err != nil {
		return 0, err
	}

	return 8, nil
}

//
// EnumProperty
//

type EnumPropertyValue struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"Value,omitempty"`
}

func newEnumPropertyValue() Value {
	return &EnumPropertyValue{}
}

func (p *Property) GetEnumPropertyValue() (*EnumPropertyValue, error) {
	if v, ok := p.Value.(*EnumPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *EnumPropertyValue) parse(d *data.Data, inner bool) error {
	if inner {
		return v.parseInner(d)
	}

	var err error
	v.Type, err = d.ReadString()
	if err != nil {
		return err
	}

	err = d.NextByteIsNull()
	if err != nil {
		return err
	}

	v.Value, err = d.ReadString()
	if err != nil {
		return err
	}

	return nil
}

// When inside an array the entire enum is stored as a single string.
// Split based on the delimiter ::
func (v *EnumPropertyValue) parseInner(d *data.Data) error {
	enum, err := d.ReadString()
	if err != nil {
		return err
	}

	enumSplit := strings.SplitN(enum, "::", 2)
	if len(enumSplit) != 2 {
		return fmt.Errorf("failed to parse enum Value %s", enumSplit)
	}

	v.Type, v.Value = enumSplit[0], enumSplit[1]

	return nil
}

func (v *EnumPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	if inner {
		return 0, d.WriteString(fmt.Sprintf("%s::%s", v.Type, v.Value))
	}

	err := d.WriteString(v.Type)
	if err != nil {
		return 0, err
	}

	err = d.WriteNull()
	if err != nil {
		return 0, err
	}

	m := d.Measure()

	err = d.WriteString(v.Value)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// FloatProperty
//

type FloatPropertyValue float32

func newFloatPropertyValue() Value {
	v := FloatPropertyValue(0)
	return &v
}

func (p *Property) GetFloatValue() (float32, error) {
	if v, ok := p.Value.(*FloatPropertyValue); ok {
		return float32(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *FloatPropertyValue) parse(d *data.Data, inner bool) error {
	if !inner {
		err := d.NextByteIsNull()
		if err != nil {
			return err
		}
	}

	f, err := d.ReadFloat32()
	if err != nil {
		return err
	}

	*v = FloatPropertyValue(f)
	return nil
}

func (v *FloatPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	if !inner {
		err := d.WriteNull()
		if err != nil {
			return 0, err
		}
	}

	err := d.WriteFloat32(float32(*v))
	if err != nil {
		return 0, err
	}

	return 4, nil
}

//
// Int8Property
//

type Int8PropertyValue int8

func newInt8PropertyValue() Value {
	v := Int8PropertyValue(0)
	return &v
}

func (p *Property) GetInt8Value() (int8, error) {
	if v, ok := p.Value.(*Int8PropertyValue); ok {
		return int8(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *Int8PropertyValue) parse(d *data.Data, inner bool) error {
	err := d.NextByteIsNull()
	if err != nil {
		return err
	}

	f, err := d.ReadInt8()
	if err != nil {
		return err
	}
	*v = Int8PropertyValue(f)
	return nil
}

func (v *Int8PropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	err := d.WriteNull()
	if err != nil {
		return 0, err
	}

	err = d.WriteInt8(int8(*v))
	if err != nil {
		return 0, err
	}

	return 1, nil
}

//
// Int64Property
//

type Int64PropertyValue int64

func newInt64PropertyValue() Value {
	v := Int64PropertyValue(0)
	return &v
}

func (p *Property) GetInt64Value() (int64, error) {
	if v, ok := p.Value.(*Int64PropertyValue); ok {
		return int64(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *Int64PropertyValue) parse(d *data.Data, inner bool) error {
	err := d.NextByteIsNull()
	if err != nil {
		return err
	}

	f, err := d.ReadInt64()
	if err != nil {
		return err
	}
	*v = Int64PropertyValue(f)
	return nil
}

func (v *Int64PropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	err := d.WriteNull()
	if err != nil {
		return 0, err
	}

	err = d.WriteInt64(int64(*v))
	if err != nil {
		return 0, err
	}

	return 8, nil
}

//
// InterfaceProperty
//

type InterfacePropertyValue struct {
	LevelName string `json:"level_name"`
	PathName  string `json:"path_name"`
}

func newInterfacePropertyValue() Value {
	return &InterfacePropertyValue{}
}

func (p *Property) GetInterfaceValue() (*InterfacePropertyValue, error) {
	if v, ok := p.Value.(*InterfacePropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *InterfacePropertyValue) parse(d *data.Data, inner bool) error {
	var err error

	if !inner {
		err := d.NextByteIsNull()
		if err != nil {
			return err
		}
	}

	v.LevelName, err = d.ReadString()
	if err != nil {
		return err
	}

	v.PathName, err = d.ReadString()
	if err != nil {
		return err
	}

	return nil
}

func (v *InterfacePropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	if !inner {
		err := d.WriteNull()
		if err != nil {
			return 0, err
		}
	}

	m := d.Measure()

	err := d.WriteString(v.LevelName)
	if err != nil {
		return 0, err
	}

	err = d.WriteString(v.PathName)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// IntProperty
//

type IntPropertyValue int32

func newIntPropertyValue() Value {
	v := IntPropertyValue(0)
	return &v
}

func (p *Property) GetIntValue() (int32, error) {
	if v, ok := p.Value.(*IntPropertyValue); ok {
		return int32(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *IntPropertyValue) parse(d *data.Data, inner bool) error {
	if !inner {
		err := d.NextByteIsNull()
		if err != nil {
			return err
		}
	}

	i, err := d.ReadInt32()
	if err != nil {
		return err
	}
	*v = IntPropertyValue(i)
	return nil
}

func (v *IntPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	if !inner {
		err := d.WriteNull()
		if err != nil {
			return 0, err
		}
	}

	err := d.WriteInt32(int32(*v))
	if err != nil {
		return 0, err
	}

	return 4, nil
}

//
// MapProperty
//

type MapPropertyValue struct {
	KeyType   Type            `json:"key_type,omitempty"`
	ValueType Type            `json:"value_type,omitempty"`
	Values    map[Value]Value `json:"values,omitempty"`

	keyOrder []Value
}

func newMapPropertyValue() Value {
	return &MapPropertyValue{}
}

func (p *Property) GetMapPropertyValue() (*MapPropertyValue, error) {
	if v, ok := p.Value.(*MapPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *MapPropertyValue) parse(d *data.Data, inner bool) error {
	if v.Values == nil {
		v.Values = make(map[Value]Value)
		v.keyOrder = make([]Value, 0)
	}

	keyType, err := d.ReadString()
	if err != nil {
		return err
	}
	v.KeyType = Type(keyType)

	valueType, err := d.ReadString()
	if err != nil {
		return err
	}
	v.ValueType = Type(valueType)

	err = d.NextByteIsNull()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = d.ReadInt32()
	if err != nil {
		return err
	}

	count, err := d.ReadInt32()
	if err != nil {
		return err
	}

	var newKey newPropValueFunc
	var newValue newPropValueFunc

	switch v.KeyType {
	case EnumPropertyType:
		newKey = newEnumPropertyValue
	case Int64PropertyType:
		newKey = newInt64PropertyValue
	case IntPropertyType:
		newKey = newIntPropertyValue
	case NamePropertyType:
		newKey = newNamePropertyValue
	case ObjectPropertyType:
		newKey = newObjectPropertyValue
	case StringPropertyType:
		newKey = newStringPropertyValue
	default:
		return fmt.Errorf("unsupported property type in map key %s", v.KeyType)
	}

	switch v.ValueType {
	case BytePropertyType:
		newValue = newBytePropertyValue
	case EnumPropertyType:
		newValue = newEnumPropertyValue
	case IntPropertyType:
		newValue = newIntPropertyValue
	case ObjectPropertyType:
		newValue = newObjectPropertyValue
	case StringPropertyType:
		newValue = newStringPropertyValue
	case StructPropertyType:
		newValue = newStructPropertyValue
	default:
		return fmt.Errorf("unsupported property type in map Value %s", v.ValueType)
	}

	for i := int32(0); i < count; i++ {
		key := newKey()
		err = key.parse(d, true)
		if err != nil {
			return err
		}

		v.keyOrder = append(v.keyOrder, key)

		val := newValue()
		err = val.parse(d, true)
		if err != nil {
			return err
		}

		v.Values[key] = val
	}

	return nil
}

func (v *MapPropertyValue) MarshalJSON() ([]byte, error) {
	out := struct {
		KeyType  string `json:"key_type"`
		KeyValue string `json:"key_value"`
	}{
		KeyType:  string(v.KeyType),
		KeyValue: string(v.ValueType),
	}

	return json.Marshal(out)
}

func (v *MapPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	err := d.WriteString(string(v.KeyType))
	if err != nil {
		return 0, err
	}

	err = d.WriteString(string(v.ValueType))
	if err != nil {
		return 0, err
	}

	err = d.WriteNull()
	if err != nil {
		return 0, err
	}

	m := d.Measure()

	// UNKNOWN_DATA
	err = d.WriteInt32(0)
	if err != nil {
		return 0, err
	}

	err = d.WriteInt32(int32(len(v.keyOrder)))
	if err != nil {
		return 0, err
	}

	for _, k := range v.keyOrder {
		value, ok := v.Values[k]
		if !ok {
			return 0, fmt.Errorf("failed to find Value for key %v", k)
		}

		_, err = k.serialize(d, true)
		if err != nil {
			return 0, err
		}

		_, err = value.serialize(d, true)
		if err != nil {
			return 0, err
		}
	}

	return m(), nil
}

//
// NameProperty
//

type NamePropertyValue string

func newNamePropertyValue() Value {
	v := NamePropertyValue("")
	return &v
}

func (p *Property) GetNameValue() (string, error) {
	if v, ok := p.Value.(*NamePropertyValue); ok {
		return string(*v), nil
	}

	return "", fmt.Errorf("wrong type %s", p.Type)
}

func (v *NamePropertyValue) parse(d *data.Data, inner bool) error {
	err := d.NextByteIsNull()
	if err != nil {
		return err
	}

	s, err := d.ReadString()
	if err != nil {
		return err
	}
	*v = NamePropertyValue(s)
	return nil
}

func (v *NamePropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	err := d.WriteNull()
	if err != nil {
		return 0, err
	}

	m := d.Measure()

	err = d.WriteString(string(*v))
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// ObjectProperty
//

type ObjectPropertyValue struct {
	LevelName string `json:"level_name,omitempty"`
	PathName  string `json:"path_name,omitempty"`
}

func newObjectPropertyValue() Value {
	return &ObjectPropertyValue{}
}

func (p *Property) GetObjectValue() (*ObjectPropertyValue, error) {
	if v, ok := p.Value.(*ObjectPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *ObjectPropertyValue) parse(d *data.Data, inner bool) error {
	var err error

	if !inner {
		err := d.NextByteIsNull()
		if err != nil {
			return err
		}
	}

	v.LevelName, err = d.ReadString()
	if err != nil {
		return err
	}

	v.PathName, err = d.ReadString()
	if err != nil {
		return err
	}

	return nil
}

func (v *ObjectPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	if !inner {
		err := d.WriteNull()
		if err != nil {
			return 0, err
		}
	}

	m := d.Measure()

	err := d.WriteString(v.LevelName)
	if err != nil {
		return 0, err
	}

	err = d.WriteString(v.PathName)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// StringProperty
//

type StringPropertyValue string

func newStringPropertyValue() Value {
	v := StringPropertyValue("")
	return &v
}

func (p *Property) GetStringValue() (string, error) {
	if v, ok := p.Value.(*StringPropertyValue); ok {
		return string(*v), nil
	}

	return "", fmt.Errorf("wrong type %s", p.Type)
}

func (v *StringPropertyValue) parse(d *data.Data, inner bool) error {
	if !inner {
		err := d.NextByteIsNull()
		if err != nil {
			return err
		}
	}

	s, err := d.ReadString()
	if err != nil {
		return err
	}
	*v = StringPropertyValue(s)
	return nil
}

func (v *StringPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	if !inner {
		err := d.WriteNull()
		if err != nil {
			return 0, err
		}
	}

	m := d.Measure()

	err := d.WriteString(string(*v))
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// StructProperty
//

type StructPropertyValue struct {
	GUID []int32    `json:"guid,omitempty"`
	Type StructType `json:"type,omitempty"`

	// Value of the Struct
	//
	// Accessing the value can be achieved by calling the appropriate Get method
	// on the StructPropertyValue based on its Type.
	Value StructValue `json:"value,omitempty"`
}

func newStructPropertyValue() Value {
	return &StructPropertyValue{}
}

func (p *Property) GetStructValue() (*StructPropertyValue, error) {
	if v, ok := p.Value.(*StructPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *StructPropertyValue) parse(d *data.Data, inner bool) error {
	if !inner {
		structType, err := d.ReadString()
		if err != nil {
			return err
		}
		v.Type = StructType(structType)

		v.GUID, err = d.ReadInt32Array(4)
		if err != nil {
			return err
		}

		err = d.NextByteIsNull()
		if err != nil {
			return err
		}
	}

	switch v.Type {
	case BoxStructType:
		v.Value = &BoxStruct{}
	case ColorStructType:
		v.Value = &ColorStruct{}
	case DateTimeStructType:
		ds := DateTimeStruct(0)
		v.Value = &ds
	case FluidBoxStructType:
		s := FluidBoxStruct(0)
		v.Value = &s
	case GUIDStructType:
		v.Value = &GUIDStruct{}
	case InventoryItemStructType:
		v.Value = &InventoryItemStruct{}
	case LinearColorStructType:
		v.Value = &LinearColorStruct{}
	case QuatStructType:
		v.Value = &QuatStruct{}
	case RailroadTrackPositionStructType:
		v.Value = &RailroadTrackPositionStruct{}
	case VectorStructType:
		v.Value = &VectorStruct{}
	case Vector2DStructType:
		v.Value = &Vector2DStruct{}
	default:
		v.Value = &ArbitraryStruct{}
	}

	err := v.Value.parse(d)
	if err != nil {
		return err
	}

	return nil
}

func (v *StructPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	if !inner {
		err := d.WriteString(string(v.Type))
		if err != nil {
			return 0, err
		}

		err = d.WriteInt32Array(v.GUID)
		if err != nil {
			return 0, err
		}

		err = d.WriteNull()
		if err != nil {
			return 0, err
		}
	}

	l, err := v.Value.serialize(d)
	if err != nil {
		return 0, err
	}

	return l, nil
}

//
// TextProperty
//

type TextPropertyValue struct {
	Flags int32
	Type  TextType
	value textValue
}

func newTextPropertyValue() Value {
	return &TextPropertyValue{}
}

func (p *Property) GetTextValue() (*TextPropertyValue, error) {
	if v, ok := p.Value.(*TextPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *TextPropertyValue) parse(d *data.Data, inner bool) error {
	err := d.NextByteIsNull()
	if err != nil {
		return err
	}

	v.Flags, err = d.ReadInt32()
	if err != nil {
		return err
	}

	textType, err := d.ReadByte()
	if err != nil {
		return err
	}

	v.Type = TextType(textType)

	// TODO: Handle other text types
	switch v.Type {
	case BaseTextType:
		v.value = &TextBase{}
	case NoneTextType:
		v.value = &TextNone{}
	default:
		return fmt.Errorf("unknown text type %v", v.Type)
	}

	err = v.value.parse(d)
	if err != nil {
		return err
	}

	return nil
}

func (v *TextPropertyValue) serialize(d *data.Data, inner bool) (int32, error) {
	err := d.WriteNull()
	if err != nil {
		return 0, err
	}

	m := d.Measure()

	err = d.WriteInt32(v.Flags)
	if err != nil {
		return 0, err
	}

	err = d.WriteByte(byte(v.Type))
	if err != nil {
		return 0, err
	}

	err = v.value.serialize(d)
	if err != nil {
		return 0, err
	}

	return m(), nil
}
