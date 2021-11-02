package save

import (
	"encoding/json"
	"fmt"
	"strings"
)

type newPropValueFunc func() PropertyValue

//
// ArrayProperty
//

type ArrayPropertyValue struct {
	ValueType PropertyType    `json:"value_type,omitempty"`
	Values    []PropertyValue `json:"values,omitempty"`

	// TODO: Explain the fields below.

	StructName      string
	StructBytes     []byte
	StructInnerType string
	StructGUID      []int32
}

func newArrayPropertyValue() PropertyValue {
	return &ArrayPropertyValue{}
}

func (p *Property) GetArrayValue() (*ArrayPropertyValue, error) {
	if v, ok := p.PropertyValue.(*ArrayPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *ArrayPropertyValue) GetStructValues() ([]*StructPropertyValue, error) {
	if v.ValueType != StructPropertyType {
		return nil, fmt.Errorf("wrong value type %s", v.ValueType)
	}

	values := make([]*StructPropertyValue, len(v.Values))

	for i, pv := range v.Values {
		values[i] = pv.(*StructPropertyValue)
	}

	return values, nil
}

func (v *ArrayPropertyValue) parse(p *parser, inner bool) error {
	valueType, err := p.readString()
	if err != nil {
		return err
	}
	v.ValueType = PropertyType(valueType)

	// UNKNOWN_DATA
	_, err = p.readByte()
	if err != nil {
		return err
	}

	elemCount, err := p.readInt32()
	if err != nil {
		return err
	}

	// Special handling for BytePropertyType.
	// Rather than creating 1 entry in the array per byte we just want to parse all
	// the data in to a single BytePropertyValue.
	if v.ValueType == BytePropertyType {
		v.Values = make([]PropertyValue, 1)

		elem := &BytePropertyValue{
			len: elemCount,
		}
		err := elem.parse(p, true)
		if err != nil {
			return err
		}

		v.Values[0] = elem

		return nil
	}

	v.Values = make([]PropertyValue, elemCount)

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
		// TODO: Do these elements serve any purpose? Will they ever be anything other than expected values?
		//  Possibly add some checks and log if they are found to differ.
		// Name
		v.StructName, err = p.readString()
		if err != nil {
			return err
		}

		// Type
		// Is this always StructProperty?
		_, err = p.readString()
		if err != nil {
			return err
		}

		// UNKNOWN_DATA
		v.StructBytes, err = p.readBytes(8)
		if err != nil {
			return err
		}

		v.StructInnerType, err = p.readString()
		if err != nil {
			return err
		}

		// GUID
		v.StructGUID, err = p.readInt32Array(4)
		if err != nil {
			return err
		}

		err = p.nextByteIsNull()
		if err != nil {
			return err
		}

		newPropValue = func() PropertyValue {
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

		err := propVal.parse(p, true)
		if err != nil {
			return err
		}

		v.Values[i] = propVal
	}

	return nil
}

func (v *ArrayPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	err := s.writeString(string(v.ValueType))
	if err != nil {
		return 0, err
	}

	// UNKNOWN_DATA
	err = s.writeNull()
	if err != nil {
		return 0, err
	}

	m := s.measure()

	// Special handling for BytePropertyType.
	// See parse() for more details.
	if v.ValueType == BytePropertyType {
		if len(v.Values) != 1 {
			return 0, fmt.Errorf("arrayproperty containg byteproperty, expected 1 value got %d", len(v.Values))
		}

		// Write placeholder element count.
		elemCountPos := s.body.Index
		err = s.writeInt32(0)
		if err != nil {
			return 0, err
		}

		// Write all bytes.
		l, err := v.Values[0].serialize(s, true)
		if err != nil {
			return 0, err
		}

		// Update element count.
		err = s.writeLen(l, elemCountPos)
		if err != nil {
			return 0, err
		}

		return m(), nil
	}

	err = s.writeInt32(int32(len(v.Values)))
	if err != nil {
		return 0, err
	}

	if v.ValueType == StructPropertyType {
		err = s.writeString(v.StructName)
		if err != nil {
			return 0, err
		}

		err = s.writeString("StructProperty")
		if err != nil {
			return 0, err
		}

		err = s.writeBytes(v.StructBytes)
		if err != nil {
			return 0, err
		}

		err = s.writeString(v.StructInnerType)
		if err != nil {
			return 0, err
		}

		err = s.writeInt32Array(v.StructGUID)
		if err != nil {
			return 0, err
		}

		err = s.writeNull()
		if err != nil {
			return 0, err
		}
	}

	for _, pv := range v.Values {
		_, err := pv.serialize(s, true)
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

func newBoolPropertyValue() PropertyValue {
	v := BoolPropertyValue(false)
	return &v
}

func (p *Property) GetBoolValue() (bool, error) {
	if v, ok := p.PropertyValue.(*BoolPropertyValue); ok {
		return bool(*v), nil
	}

	return false, fmt.Errorf("wrong type %s", p.Type)
}

func (v *BoolPropertyValue) parse(p *parser, inner bool) error {
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

func (v *BoolPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	var err error
	if *v {
		err = s.writeBool(true)
	} else {
		err = s.writeBool(false)
	}
	if err != nil {
		return 0, err
	}

	return 0, s.writeNull()
}

//
// ByteProperty
//

type BytePropertyValue struct {
	Type  string `json:"text,omitempty"`
	Value []byte `json:"value,omitempty"`

	len int32
}

func newBytePropertyValue() PropertyValue {
	return &BytePropertyValue{}
}

func (p *Property) GetByteValue() ([]byte, error) {
	if v, ok := p.PropertyValue.(*BytePropertyValue); ok {
		return v.Value, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *BytePropertyValue) parse(p *parser, inner bool) error {
	if inner {
		return v.parseInner(p)
	}

	var err error
	v.Type, err = p.readString()
	if err != nil {
		return err
	}

	err = p.nextByteIsNull()
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
		value, err := p.readString()
		if err != nil {
			return err
		}

		v.Value = []byte(value)
	}

	return nil
}

func (v *BytePropertyValue) parseInner(p *parser) error {
	// If len is 0 then we must be parsing a ByteProperty in the value of a MapProperty.
	// Best guess is that it's only a single byte.
	if v.len == 0 {
		v.len = 1
	}

	v.Type = "None"

	var err error
	v.Value, err = p.readBytes(v.len)
	if err != nil {
		return err
	}

	return nil
}

func (v *BytePropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	if inner {
		return v.serializeInner(s)
	}

	err := s.writeString(v.Type)
	if err != nil {
		return 0, err
	}

	err = s.writeNull()
	if err != nil {
		return 0, err
	}

	m := s.measure()

	switch v.Type {
	case "None":
		err = s.writeByte(v.Value[0])
		if err != nil {
			return 0, err
		}
	default:
		err = s.writeString(string(v.Value))
		if err != nil {
			return 0, err
		}
	}

	return m(), nil
}

func (v *BytePropertyValue) serializeInner(s *serializer) (int32, error) {
	m := s.measure()

	err := s.writeBytes(v.Value)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// DoubleProperty
//

type DoublePropertyValue float64

func newDoublePropertyValue() PropertyValue {
	v := DoublePropertyValue(0)
	return &v
}

func (p *Property) GetDoubleValue() (float64, error) {
	if v, ok := p.PropertyValue.(*DoublePropertyValue); ok {
		return float64(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *DoublePropertyValue) parse(p *parser, inner bool) error {
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

func (v *DoublePropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	err := s.writeNull()
	if err != nil {
		return 0, err
	}

	err = s.writeFloat64(float64(*v))
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
	Value string `json:"value,omitempty"`
}

func newEnumPropertyValue() PropertyValue {
	return &EnumPropertyValue{}
}

func (p *Property) GetEnumPropertyValue() (*EnumPropertyValue, error) {
	if v, ok := p.PropertyValue.(*EnumPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *EnumPropertyValue) parse(p *parser, inner bool) error {
	if inner {
		return v.parseInner(p)
	}

	var err error
	v.Type, err = p.readString()
	if err != nil {
		return err
	}

	err = p.nextByteIsNull()
	if err != nil {
		return err
	}

	v.Value, err = p.readString()
	if err != nil {
		return err
	}

	return nil
}

// When inside an array the entire enum is stored as a single string.
// Split based on the delimiter ::
func (v *EnumPropertyValue) parseInner(p *parser) error {
	enum, err := p.readString()
	if err != nil {
		return err
	}

	enumSplit := strings.SplitN(enum, "::", 2)
	if len(enumSplit) != 2 {
		return fmt.Errorf("failed to parse enum value %s", enumSplit)
	}

	v.Type, v.Value = enumSplit[0], enumSplit[1]

	return nil
}

func (v *EnumPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	if inner {
		return 0, s.writeString(fmt.Sprintf("%s::%s", v.Type, v.Value))
	}

	err := s.writeString(v.Type)
	if err != nil {
		return 0, err
	}

	err = s.writeNull()
	if err != nil {
		return 0, err
	}

	m := s.measure()

	err = s.writeString(v.Value)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// FloatProperty
//

type FloatPropertyValue float32

func newFloatPropertyValue() PropertyValue {
	v := FloatPropertyValue(0)
	return &v
}

func (p *Property) GetFloatValue() (float32, error) {
	if v, ok := p.PropertyValue.(*FloatPropertyValue); ok {
		return float32(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *FloatPropertyValue) parse(p *parser, inner bool) error {
	if !inner {
		err := p.nextByteIsNull()
		if err != nil {
			return err
		}
	}

	f, err := p.readFloat32()
	if err != nil {
		return err
	}

	*v = FloatPropertyValue(f)
	return nil
}

func (v *FloatPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	if !inner {
		err := s.writeNull()
		if err != nil {
			return 0, err
		}
	}

	err := s.writeFloat32(float32(*v))
	if err != nil {
		return 0, err
	}

	return 4, nil
}

//
// Int8Property
//

type Int8PropertyValue int8

func newInt8PropertyValue() PropertyValue {
	v := Int8PropertyValue(0)
	return &v
}

func (p *Property) GetInt8Value() (int8, error) {
	if v, ok := p.PropertyValue.(*Int8PropertyValue); ok {
		return int8(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *Int8PropertyValue) parse(p *parser, inner bool) error {
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

func (v *Int8PropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	err := s.writeNull()
	if err != nil {
		return 0, err
	}

	err = s.writeInt8(int8(*v))
	if err != nil {
		return 0, err
	}

	return 1, nil
}

//
// Int64Property
//

type Int64PropertyValue int64

func newInt64PropertyValue() PropertyValue {
	v := Int64PropertyValue(0)
	return &v
}

func (p *Property) GetInt64Value() (int64, error) {
	if v, ok := p.PropertyValue.(*Int64PropertyValue); ok {
		return int64(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *Int64PropertyValue) parse(p *parser, inner bool) error {
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

func (v *Int64PropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	err := s.writeNull()
	if err != nil {
		return 0, err
	}

	err = s.writeInt64(int64(*v))
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

func newInterfacePropertyValue() PropertyValue {
	return &InterfacePropertyValue{}
}

func (p *Property) GetInterfaceValue() (*InterfacePropertyValue, error) {
	if v, ok := p.PropertyValue.(*InterfacePropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *InterfacePropertyValue) parse(p *parser, inner bool) error {
	var err error

	if !inner {
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

func (v *InterfacePropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	if !inner {
		err := s.writeNull()
		if err != nil {
			return 0, err
		}
	}

	m := s.measure()

	err := s.writeString(v.LevelName)
	if err != nil {
		return 0, err
	}

	err = s.writeString(v.PathName)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// IntProperty
//

type IntPropertyValue int32

func newIntPropertyValue() PropertyValue {
	v := IntPropertyValue(0)
	return &v
}

func (p *Property) GetIntValue() (int32, error) {
	if v, ok := p.PropertyValue.(*IntPropertyValue); ok {
		return int32(*v), nil
	}

	return 0, fmt.Errorf("wrong type %s", p.Type)
}

func (v *IntPropertyValue) parse(p *parser, inner bool) error {
	if !inner {
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

func (v *IntPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	if !inner {
		err := s.writeNull()
		if err != nil {
			return 0, err
		}
	}

	err := s.writeInt32(int32(*v))
	if err != nil {
		return 0, err
	}

	return 4, nil
}

//
// MapProperty
//

type MapPropertyValue struct {
	KeyType   PropertyType                    `json:"key_type,omitempty"`
	ValueType PropertyType                    `json:"value_type,omitempty"`
	Values    map[PropertyValue]PropertyValue `json:"values,omitempty"`

	keyOrder []PropertyValue
}

func newMapPropertyValue() PropertyValue {
	return &MapPropertyValue{}
}

func (p *Property) GetMapPropertyValue() (*MapPropertyValue, error) {
	if v, ok := p.PropertyValue.(*MapPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *MapPropertyValue) parse(p *parser, inner bool) error {
	if v.Values == nil {
		v.Values = make(map[PropertyValue]PropertyValue)
		v.keyOrder = make([]PropertyValue, 0)
	}

	keyType, err := p.readString()
	if err != nil {
		return err
	}
	v.KeyType = PropertyType(keyType)

	valueType, err := p.readString()
	if err != nil {
		return err
	}
	v.ValueType = PropertyType(valueType)

	err = p.nextByteIsNull()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = p.readInt32()
	if err != nil {
		return err
	}

	count, err := p.readInt32()
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
		return fmt.Errorf("unsupported property type in map value %s", v.ValueType)
	}

	for i := int32(0); i < count; i++ {
		key := newKey()
		err = key.parse(p, true)
		if err != nil {
			return err
		}

		v.keyOrder = append(v.keyOrder, key)

		val := newValue()
		err = val.parse(p, true)
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

func (v *MapPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	err := s.writeString(string(v.KeyType))
	if err != nil {
		return 0, err
	}

	err = s.writeString(string(v.ValueType))
	if err != nil {
		return 0, err
	}

	err = s.writeNull()
	if err != nil {
		return 0, err
	}

	m := s.measure()

	// UNKNOWN_DATA
	err = s.writeInt32(0)
	if err != nil {
		return 0, err
	}

	err = s.writeInt32(int32(len(v.keyOrder)))
	if err != nil {
		return 0, err
	}

	for _, k := range v.keyOrder {
		value, ok := v.Values[k]
		if !ok {
			return 0, fmt.Errorf("failed to find value for key %v", k)
		}

		_, err = k.serialize(s, true)
		if err != nil {
			return 0, err
		}

		_, err = value.serialize(s, true)
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

func newNamePropertyValue() PropertyValue {
	v := NamePropertyValue("")
	return &v
}

func (p *Property) GetNameValue() (string, error) {
	if v, ok := p.PropertyValue.(*NamePropertyValue); ok {
		return string(*v), nil
	}

	return "", fmt.Errorf("wrong type %s", p.Type)
}

func (v *NamePropertyValue) parse(p *parser, inner bool) error {
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

func (v *NamePropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	err := s.writeNull()
	if err != nil {
		return 0, err
	}

	m := s.measure()

	err = s.writeString(string(*v))
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

func newObjectPropertyValue() PropertyValue {
	return &ObjectPropertyValue{}
}

func (p *Property) GetObjectValue() (*ObjectPropertyValue, error) {
	if v, ok := p.PropertyValue.(*ObjectPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *ObjectPropertyValue) parse(p *parser, inner bool) error {
	var err error

	if !inner {
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

func (v *ObjectPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	if !inner {
		err := s.writeNull()
		if err != nil {
			return 0, err
		}
	}

	m := s.measure()

	err := s.writeString(v.LevelName)
	if err != nil {
		return 0, err
	}

	err = s.writeString(v.PathName)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// StringProperty
//

type StringPropertyValue string

func newStringPropertyValue() PropertyValue {
	v := StringPropertyValue("")
	return &v
}

func (p *Property) GetStringValue() (string, error) {
	if v, ok := p.PropertyValue.(*StringPropertyValue); ok {
		return string(*v), nil
	}

	return "", fmt.Errorf("wrong type %s", p.Type)
}

func (v *StringPropertyValue) parse(p *parser, inner bool) error {
	if !inner {
		err := p.nextByteIsNull()
		if err != nil {
			return err
		}
	}

	s, err := p.readString()
	if err != nil {
		return err
	}
	*v = StringPropertyValue(s)
	return nil
}

func (v *StringPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	if !inner {
		err := s.writeNull()
		if err != nil {
			return 0, err
		}
	}

	m := s.measure()

	err := s.writeString(string(*v))
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// StructProperty
//

type StructValue interface {
	parse(p *parser) error
	serialize(s *serializer) (int32, error)
}

type StructPropertyValue struct {
	GUID  []int32     `json:"guid,omitempty"`
	Type  StructType  `json:"type,omitempty"`
	Value StructValue `json:"value,omitempty"`
}

func newStructPropertyValue() PropertyValue {
	return &StructPropertyValue{}
}

func (p *Property) GetStructValue() (*StructPropertyValue, error) {
	if v, ok := p.PropertyValue.(*StructPropertyValue); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", p.Type)
}

func (v *StructPropertyValue) parse(p *parser, inner bool) error {
	if !inner {
		structType, err := p.readString()
		if err != nil {
			return err
		}
		v.Type = StructType(structType)

		v.GUID, err = p.readInt32Array(4)
		if err != nil {
			return err
		}

		err = p.nextByteIsNull()
		if err != nil {
			return err
		}
	}

	switch v.Type {
	case BoxStructType:
		v.Value = &BoxStruct{}
	case ColorStructType:
		panic("TODO ColorStructType")
	case DateTimeStructType:
		panic("TODO DateTimeStructType")
	case FluidBoxStructType:
		s := FluidBoxStruct(0)
		v.Value = &s
	case GUIDStructType:
		panic("TODO GUIDStructType")
	case InventoryItemStructType:
		v.Value = &InventoryItemStruct{}
	case LinearColorStructType:
		v.Value = &LinearColor{}
	case QuatStructType:
		v.Value = &QuatStruct{}
	case RailroadTrackPositionStructType:
		panic("TODO RailroadTrackPositionStructType")
	case VectorStructType:
		v.Value = &VectorStruct{}
	case Vector2DStructType:
		panic("TODO Vector2DStructType")
	default:
		v.Value = &ArbitraryStruct{}
	}

	err := v.Value.parse(p)
	if err != nil {
		return err
	}

	return nil
}

func (v *StructPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	if !inner {
		err := s.writeString(string(v.Type))
		if err != nil {
			return 0, err
		}

		err = s.writeInt32Array(v.GUID)
		if err != nil {
			return 0, err
		}

		err = s.writeNull()
		if err != nil {
			return 0, err
		}
	}

	l, err := v.Value.serialize(s)
	if err != nil {
		return 0, err
	}

	return l, nil
}

//
// TextProperty
//

// TODO: TextProperty

type TextPropertyValue struct{}

func newTextPropertyValue() PropertyValue {
	return &TextPropertyValue{}
}

func (v *TextPropertyValue) parse(p *parser, inner bool) error {
	panic("implement me")
}

func (v *TextPropertyValue) serialize(s *serializer, inner bool) (int32, error) {
	panic("implement me")
}
