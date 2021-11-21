package save

import "fmt"

type StructType string

const (
	ArbitraryStructType             StructType = "Arbitrary"
	BoxStructType                   StructType = "Box"
	ColorStructType                 StructType = "Color"
	DateTimeStructType              StructType = "DateTime"
	FluidBoxStructType              StructType = "FluidBox"
	GUIDStructType                  StructType = "GUID"
	InventoryItemStructType         StructType = "InventoryItem"
	LinearColorStructType           StructType = "LinearColor"
	QuatStructType                  StructType = "Quat"
	RailroadTrackPositionStructType StructType = "RailroadTrackPosition"
	VectorStructType                StructType = "Vector"
	Vector2DStructType              StructType = "Vector2D"
)

//
// ArbitraryStruct
//

type ArbitraryStruct struct {
	Properties []*Property

	numProps int32
}

func (v *StructPropertyValue) GetArbitraryStruct() (*ArbitraryStruct, error) {
	if v, ok := v.Value.(*ArbitraryStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *ArbitraryStruct) parse(p *parser) error {
	for {
		prop, err := p.parseProperty()
		if err != nil {
			return err
		}
		if prop == nil {
			break
		}

		v.Properties = append(v.Properties, prop)
	}

	return nil
}

func (v *ArbitraryStruct) serialize(s *serializer) (int32, error) {
	m := s.measure()

	err := s.serializeProperties(v.Properties)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// BoxStruct
//

type BoxStruct struct {
	Min     []float32
	Max     []float32
	IsValid bool
}

func (v *StructPropertyValue) GetBoxStruct() (*BoxStruct, error) {
	if v, ok := v.Value.(*BoxStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *BoxStruct) parse(p *parser) error {
	var err error
	v.Min, err = p.readFloat32Array(3)
	if err != nil {
		return err
	}

	v.Max, err = p.readFloat32Array(3)
	if err != nil {
		return err
	}

	v.IsValid, err = p.readBool()
	if err != nil {
		return err
	}

	return nil
}

func (v *BoxStruct) serialize(s *serializer) (int32, error) {
	m := s.measure()

	err := s.writeFloat32Array(v.Min)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32Array(v.Max)
	if err != nil {
		return 0, err
	}

	err = s.writeBool(v.IsValid)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// ColorStruct
//

type ColorStruct struct {
	B byte
	G byte
	R byte
	A byte
}

func (v *StructPropertyValue) GetColorStruct() (*ColorStruct, error) {
	if v, ok := v.Value.(*ColorStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *ColorStruct) parse(p *parser) error {
	var err error
	v.B, err = p.readByte()
	if err != nil {
		return err
	}

	v.G, err = p.readByte()
	if err != nil {
		return err
	}

	v.R, err = p.readByte()
	if err != nil {
		return err
	}

	v.A, err = p.readByte()
	if err != nil {
		return err
	}

	return nil
}

func (v *ColorStruct) serialize(s *serializer) (int32, error) {
	m := s.measure()

	err := s.writeByte(v.B)
	if err != nil {
		return 0, err
	}

	err = s.writeByte(v.G)
	if err != nil {
		return 0, err
	}

	err = s.writeByte(v.R)
	if err != nil {
		return 0, err
	}

	err = s.writeByte(v.A)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// DateTimeStruct
//

type DateTimeStruct int64

func (v *StructPropertyValue) GetDateTimeStruct() (*DateTimeStruct, error) {
	if v, ok := v.Value.(*DateTimeStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *DateTimeStruct) parse(p *parser) error {
	i, err := p.readInt64()
	if err != nil {
		return err
	}

	*v = DateTimeStruct(i)

	return nil
}

func (v *DateTimeStruct) serialize(s *serializer) (int32, error) {
	err := s.writeInt64(int64(*v))
	if err != nil {
		return 0, err
	}

	return 8, nil
}

//
// FluidBoxStruct
//

type FluidBoxStruct float32

func (v *FluidBoxStruct) parse(p *parser) error {
	f, err := p.readFloat32()
	if err != nil {
		return err
	}

	*v = FluidBoxStruct(f)

	return nil
}

func (v *FluidBoxStruct) serialize(s *serializer) (int32, error) {
	return 4, s.writeFloat32(float32(*v))
}

//
// GUIDStruct
//

type GUIDStruct []byte

func (v *GUIDStruct) parse(p *parser) error {
	b, err := p.readBytes(16)
	if err != nil {
		return err
	}

	*v = b

	return nil
}

func (v *GUIDStruct) serialize(s *serializer) (int32, error) {
	err := s.writeBytes(*v)
	if err != nil {
		return 0, err
	}

	return 16, nil
}

//
// InventoryItemStruct
//

type InventoryItemStruct struct {
	ItemName  string
	LevelName string
	PathName  string
	NumItems  int32
}

func (v *StructPropertyValue) GetInventoryItemStruct() (*InventoryItemStruct, error) {
	if v, ok := v.Value.(*InventoryItemStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *InventoryItemStruct) parse(p *parser) error {
	// UNKNOWN_DATA
	_, err := p.readInt32()
	if err != nil {
		return err
	}

	v.ItemName, err = p.readString()
	if err != nil {
		return err
	}

	v.LevelName, err = p.readString()
	if err != nil {
		return err
	}

	v.PathName, err = p.readString()
	if err != nil {
		return err
	}

	propName, err := p.readString()
	if err != nil {
		return err
	}
	if propName != "NumItems" {
		return fmt.Errorf("found item property that was not NumItems %q", propName)
	}

	propType, err := p.readString()
	if err != nil {
		return err
	}
	if propType != "IntProperty" {
		return fmt.Errorf("found item property type that was not IntProperty %q", propType)
	}

	// Value length. Presuamble always 4.
	_, err = p.readInt32()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = p.readBytes(5)
	if err != nil {
		return err
	}

	v.NumItems, err = p.readInt32()
	if err != nil {
		return err
	}

	return nil
}

func (v *InventoryItemStruct) serialize(s *serializer) (int32, error) {
	m := s.measure()

	// UNKNOWN_DATA
	err := s.writeInt32(0)
	if err != nil {
		return 0, err
	}

	err = s.writeString(v.ItemName)
	if err != nil {
		return 0, err
	}

	err = s.writeString(v.LevelName)
	if err != nil {
		return 0, err
	}

	err = s.writeString(v.PathName)
	if err != nil {
		return 0, err
	}

	// Data after this is not included in the size.
	l := m()

	err = s.writeString("NumItems")
	if err != nil {
		return 0, err
	}

	err = s.writeString("IntProperty")
	if err != nil {
		return 0, err
	}

	// Value length.
	err = s.writeInt32(4)
	if err != nil {
		return 0, err
	}

	// UNKNOWN_DATA
	err = s.writeNulls(5)
	if err != nil {
		return 0, err
	}

	err = s.writeInt32(v.NumItems)
	if err != nil {
		return 0, err
	}

	return l, nil
}

//
// LinearColorStruct
//

type LinearColorStruct struct {
	R float32
	G float32
	B float32
	A float32
}

func (v *StructPropertyValue) GetLinearColor() (*LinearColorStruct, error) {
	if v, ok := v.Value.(*LinearColorStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *LinearColorStruct) parse(p *parser) error {
	var err error
	v.R, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.G, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.B, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.A, err = p.readFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *LinearColorStruct) serialize(s *serializer) (int32, error) {
	err := s.writeFloat32(v.R)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.G)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.B)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.A)
	if err != nil {
		return 0, err
	}

	return 16, nil
}

//
// RailroadTrackPositionStruct
//

type RailroadTrackPositionStruct struct {
	LevelName string  `json:"levelName"`
	PathName  string  `json:"pathName"`
	Offset    float32 `json:"offset"`
	Forward   float32 `json:"forward"`
}

func (v *RailroadTrackPositionStruct) parse(p *parser) error {
	var err error

	v.LevelName, err = p.readString()
	if err != nil {
		return err
	}

	v.PathName, err = p.readString()
	if err != nil {
		return err
	}

	v.Offset, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.Forward, err = p.readFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *RailroadTrackPositionStruct) serialize(s *serializer) (int32, error) {
	m := s.measure()

	err := s.writeString(v.LevelName)
	if err != nil {
		return 0, err
	}

	err = s.writeString(v.PathName)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.Offset)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.Forward)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// QuatStruct
//

type QuatStruct struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (v *QuatStruct) parse(p *parser) error {
	var err error
	v.X, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.Y, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.Z, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.W, err = p.readFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *QuatStruct) serialize(s *serializer) (int32, error) {
	err := s.writeFloat32(v.X)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.Y)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.Z)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.W)
	if err != nil {
		return 0, err
	}

	return 16, nil
}

//
// VectorStruct
//

type VectorStruct struct {
	X float32
	Y float32
	Z float32
}

func (v *VectorStruct) parse(p *parser) error {
	var err error
	v.X, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.Y, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.Z, err = p.readFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *VectorStruct) serialize(s *serializer) (int32, error) {
	err := s.writeFloat32(v.X)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.Y)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.Z)
	if err != nil {
		return 0, err
	}

	return 12, nil
}

//
// Vector2DStruct
//

type Vector2DStruct struct {
	X float32
	Y float32
}

func (v *Vector2DStruct) parse(p *parser) error {
	var err error
	v.X, err = p.readFloat32()
	if err != nil {
		return err
	}

	v.Y, err = p.readFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *Vector2DStruct) serialize(s *serializer) (int32, error) {
	err := s.writeFloat32(v.X)
	if err != nil {
		return 0, err
	}

	err = s.writeFloat32(v.Y)
	if err != nil {
		return 0, err
	}

	return 12, nil
}
