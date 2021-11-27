package property

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/save/data"
)

type StructType string

type StructValue interface {
	parse(d *data.Data) error
	serialize(d *data.Data) (int32, error)
}

const (
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
	Properties []*Property `json:"properties,omitempty"`
}

func (v *StructPropertyValue) GetArbitraryStruct() (*ArbitraryStruct, error) {
	if v, ok := v.Value.(*ArbitraryStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *ArbitraryStruct) parse(d *data.Data) error {
	props, err := ParseProperties(d)
	if err != nil {
		return nil
	}

	v.Properties = props

	return nil
}

func (v *ArbitraryStruct) serialize(d *data.Data) (int32, error) {
	m := d.Measure()

	err := SerializeProperties(v.Properties, d)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// BoxStruct
//

type BoxStruct struct {
	Min     []float32 `json:"min,omitempty"`
	Max     []float32 `json:"max,omitempty"`
	IsValid bool      `json:"is_valid,omitempty"`
}

func (v *StructPropertyValue) GetBoxStruct() (*BoxStruct, error) {
	if v, ok := v.Value.(*BoxStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *BoxStruct) parse(d *data.Data) error {
	var err error
	v.Min, err = d.ReadFloat32Array(3)
	if err != nil {
		return err
	}

	v.Max, err = d.ReadFloat32Array(3)
	if err != nil {
		return err
	}

	v.IsValid, err = d.ReadBool()
	if err != nil {
		return err
	}

	return nil
}

func (v *BoxStruct) serialize(d *data.Data) (int32, error) {
	m := d.Measure()

	err := d.WriteFloat32Array(v.Min)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32Array(v.Max)
	if err != nil {
		return 0, err
	}

	err = d.WriteBool(v.IsValid)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// ColorStruct
//

type ColorStruct struct {
	B byte `json:"b,omitempty"`
	G byte `json:"g,omitempty"`
	R byte `json:"r,omitempty"`
	A byte `json:"a,omitempty"`
}

func (v *StructPropertyValue) GetColorStruct() (*ColorStruct, error) {
	if v, ok := v.Value.(*ColorStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *ColorStruct) parse(d *data.Data) error {
	var err error
	v.B, err = d.ReadByte()
	if err != nil {
		return err
	}

	v.G, err = d.ReadByte()
	if err != nil {
		return err
	}

	v.R, err = d.ReadByte()
	if err != nil {
		return err
	}

	v.A, err = d.ReadByte()
	if err != nil {
		return err
	}

	return nil
}

func (v *ColorStruct) serialize(d *data.Data) (int32, error) {
	m := d.Measure()

	err := d.WriteByte(v.B)
	if err != nil {
		return 0, err
	}

	err = d.WriteByte(v.G)
	if err != nil {
		return 0, err
	}

	err = d.WriteByte(v.R)
	if err != nil {
		return 0, err
	}

	err = d.WriteByte(v.A)
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

func (v *DateTimeStruct) parse(d *data.Data) error {
	i, err := d.ReadInt64()
	if err != nil {
		return err
	}

	*v = DateTimeStruct(i)

	return nil
}

func (v *DateTimeStruct) serialize(d *data.Data) (int32, error) {
	err := d.WriteInt64(int64(*v))
	if err != nil {
		return 0, err
	}

	return 8, nil
}

//
// FluidBoxStruct
//

type FluidBoxStruct float32

func (v *StructPropertyValue) GetFluidBoxStruct() (*FluidBoxStruct, error) {
	if v, ok := v.Value.(*FluidBoxStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *FluidBoxStruct) parse(d *data.Data) error {
	f, err := d.ReadFloat32()
	if err != nil {
		return err
	}

	*v = FluidBoxStruct(f)

	return nil
}

func (v *FluidBoxStruct) serialize(d *data.Data) (int32, error) {
	return 4, d.WriteFloat32(float32(*v))
}

//
// GUIDStruct
//

type GUIDStruct []byte

func (v *StructPropertyValue) GetGUIDStruct() (*GUIDStruct, error) {
	if v, ok := v.Value.(*GUIDStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *GUIDStruct) parse(d *data.Data) error {
	b, err := d.ReadBytes(16)
	if err != nil {
		return err
	}

	*v = b

	return nil
}

func (v *GUIDStruct) serialize(d *data.Data) (int32, error) {
	err := d.WriteBytes(*v)
	if err != nil {
		return 0, err
	}

	return 16, nil
}

//
// InventoryItemStruct
//

type InventoryItemStruct struct {
	ItemName  string `json:"item_name,omitempty"`
	LevelName string `json:"level_name,omitempty"`
	PathName  string `json:"path_name,omitempty"`
	NumItems  int32  `json:"num_items,omitempty"`
}

func (v *StructPropertyValue) GetInventoryItemStruct() (*InventoryItemStruct, error) {
	if v, ok := v.Value.(*InventoryItemStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *InventoryItemStruct) parse(d *data.Data) error {
	// UNKNOWN_DATA
	_, err := d.ReadInt32()
	if err != nil {
		return err
	}

	v.ItemName, err = d.ReadString()
	if err != nil {
		return err
	}

	v.LevelName, err = d.ReadString()
	if err != nil {
		return err
	}

	v.PathName, err = d.ReadString()
	if err != nil {
		return err
	}

	propName, err := d.ReadString()
	if err != nil {
		return err
	}
	if propName != "NumItems" {
		return fmt.Errorf("found item property that was not NumItems %q", propName)
	}

	propType, err := d.ReadString()
	if err != nil {
		return err
	}
	if propType != "IntProperty" {
		return fmt.Errorf("found item property type that was not IntProperty %q", propType)
	}

	// Value length. Presuamble always 4.
	_, err = d.ReadInt32()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = d.ReadBytes(5)
	if err != nil {
		return err
	}

	v.NumItems, err = d.ReadInt32()
	if err != nil {
		return err
	}

	return nil
}

func (v *InventoryItemStruct) serialize(d *data.Data) (int32, error) {
	m := d.Measure()

	// UNKNOWN_DATA
	err := d.WriteInt32(0)
	if err != nil {
		return 0, err
	}

	err = d.WriteString(v.ItemName)
	if err != nil {
		return 0, err
	}

	err = d.WriteString(v.LevelName)
	if err != nil {
		return 0, err
	}

	err = d.WriteString(v.PathName)
	if err != nil {
		return 0, err
	}

	// Data after this is not included in the size.
	l := m()

	err = d.WriteString("NumItems")
	if err != nil {
		return 0, err
	}

	err = d.WriteString("IntProperty")
	if err != nil {
		return 0, err
	}

	// Value length.
	err = d.WriteInt32(4)
	if err != nil {
		return 0, err
	}

	// UNKNOWN_DATA
	err = d.WriteNulls(5)
	if err != nil {
		return 0, err
	}

	err = d.WriteInt32(v.NumItems)
	if err != nil {
		return 0, err
	}

	return l, nil
}

//
// LinearColorStruct
//

type LinearColorStruct struct {
	R float32 `json:"r,omitempty"`
	G float32 `json:"g,omitempty"`
	B float32 `json:"b,omitempty"`
	A float32 `json:"a,omitempty"`
}

func (v *StructPropertyValue) GetLinearColor() (*LinearColorStruct, error) {
	if v, ok := v.Value.(*LinearColorStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *LinearColorStruct) parse(d *data.Data) error {
	var err error
	v.R, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.G, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.B, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.A, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *LinearColorStruct) serialize(d *data.Data) (int32, error) {
	err := d.WriteFloat32(v.R)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.G)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.B)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.A)
	if err != nil {
		return 0, err
	}

	return 16, nil
}

//
// RailroadTrackPositionStruct
//

type RailroadTrackPositionStruct struct {
	LevelName string  `json:"levelName,omitempty"`
	PathName  string  `json:"pathName,omitempty"`
	Offset    float32 `json:"offset,omitempty"`
	Forward   float32 `json:"forward,omitempty"`
}

func (v *StructPropertyValue) GetRailroadTrackPositionStruct() (*RailroadTrackPositionStruct, error) {
	if v, ok := v.Value.(*RailroadTrackPositionStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *RailroadTrackPositionStruct) parse(d *data.Data) error {
	var err error

	v.LevelName, err = d.ReadString()
	if err != nil {
		return err
	}

	v.PathName, err = d.ReadString()
	if err != nil {
		return err
	}

	v.Offset, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.Forward, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *RailroadTrackPositionStruct) serialize(d *data.Data) (int32, error) {
	m := d.Measure()

	err := d.WriteString(v.LevelName)
	if err != nil {
		return 0, err
	}

	err = d.WriteString(v.PathName)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.Offset)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.Forward)
	if err != nil {
		return 0, err
	}

	return m(), nil
}

//
// QuatStruct
//

type QuatStruct struct {
	X float32 `json:"x,omitempty"`
	Y float32 `json:"y,omitempty"`
	Z float32 `json:"z,omitempty"`
	W float32 `json:"w,omitempty"`
}

func (v *StructPropertyValue) GetQuatStruct() (*QuatStruct, error) {
	if v, ok := v.Value.(*QuatStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *QuatStruct) parse(d *data.Data) error {
	var err error
	v.X, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.Y, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.Z, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.W, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *QuatStruct) serialize(d *data.Data) (int32, error) {
	err := d.WriteFloat32(v.X)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.Y)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.Z)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.W)
	if err != nil {
		return 0, err
	}

	return 16, nil
}

//
// VectorStruct
//

type VectorStruct struct {
	X float32 `json:"x,omitempty"`
	Y float32 `json:"y,omitempty"`
	Z float32 `json:"z,omitempty"`
}

func (v *StructPropertyValue) GetVectorStruct() (*VectorStruct, error) {
	if v, ok := v.Value.(*VectorStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *VectorStruct) parse(d *data.Data) error {
	var err error
	v.X, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.Y, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.Z, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *VectorStruct) serialize(d *data.Data) (int32, error) {
	err := d.WriteFloat32(v.X)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.Y)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.Z)
	if err != nil {
		return 0, err
	}

	return 12, nil
}

//
// Vector2DStruct
//

type Vector2DStruct struct {
	X float32 `json:"x,omitempty"`
	Y float32 `json:"y,omitempty"`
}

func (v *StructPropertyValue) GetVector2DStruct() (*Vector2DStruct, error) {
	if v, ok := v.Value.(*Vector2DStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (v *Vector2DStruct) parse(d *data.Data) error {
	var err error
	v.X, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	v.Y, err = d.ReadFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (v *Vector2DStruct) serialize(d *data.Data) (int32, error) {
	err := d.WriteFloat32(v.X)
	if err != nil {
		return 0, err
	}

	err = d.WriteFloat32(v.Y)
	if err != nil {
		return 0, err
	}

	return 12, nil
}
