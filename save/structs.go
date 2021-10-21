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

func (s *ArbitraryStruct) parse(p *parser) error {
	for {
		prop, err := p.parseProperty()
		if err != nil {
			return err
		}
		if prop == nil {
			break
		}

		s.Properties = append(s.Properties, prop)
	}

	return nil
}

func (s *ArbitraryStruct) serialize(p *parser) error {
	err := p.serializeProperties(s.Properties)
	if err != nil {
		return err
	}

	return nil
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

func (s *BoxStruct) parse(p *parser) error {
	var err error
	s.Min, err = p.readFloat32Array(3)
	if err != nil {
		return err
	}

	s.Max, err = p.readFloat32Array(3)
	if err != nil {
		return err
	}

	// TODO: Is this definitely a bool?
	s.IsValid, err = p.readBool()
	if err != nil {
		return err
	}

	return nil
}

func (s *BoxStruct) serialize(p *parser) error {
	err := p.writeFloat32Array(s.Min)
	if err != nil {
		return err
	}

	err = p.writeFloat32Array(s.Max)
	if err != nil {
		return err
	}

	err = p.writeBool(s.IsValid)
	if err != nil {
		return err
	}

	return nil
}

//
// FluidBox
//

type FluidBoxStruct float32

func (s *FluidBoxStruct) parse(p *parser) error {
	v, err := p.readFloat32()
	if err != nil {
		return err
	}

	*s = FluidBoxStruct(v)

	return nil
}

func (s *FluidBoxStruct) serialize(p *parser) error {
	return p.writeFloat32(float32(*s))
}

//
// InventoryItem
//

type InventoryItemStruct struct {
	ItemName string
	NumItems int32
}

func (v *StructPropertyValue) GetInventoryItemStruct() (*InventoryItemStruct, error) {
	if v, ok := v.Value.(*InventoryItemStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", v.Type)
}

func (s *InventoryItemStruct) parse(p *parser) error {
	// UNKNOWN_DATA
	_, err := p.readInt32()
	if err != nil {
		return err
	}

	s.ItemName, err = p.readString()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = p.readString()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = p.readString()
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

	// UNKNOWN_DATA. Possibly value length.
	_, err = p.readInt32()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = p.readBytes(5)
	if err != nil {
		return err
	}

	s.NumItems, err = p.readInt32()
	if err != nil {
		return err
	}

	return nil
}

func (s *InventoryItemStruct) serialize(p *parser) error {
	// UNKNOWN_DATA
	err := p.writeInt32(0)
	if err != nil {
		return err
	}

	err = p.writeString(s.ItemName)
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	err = p.writeInt32(0)
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	err = p.writeInt32(0)
	if err != nil {
		return err
	}

	err = p.writeString("NumItems")
	if err != nil {
		return err
	}

	err = p.writeString("IntProperty")
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	err = p.writeInt32(0)
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	err = p.writeNulls(5)
	if err != nil {
		return err
	}

	err = p.writeInt32(s.NumItems)
	if err != nil {
		return err
	}

	return nil
}

//
// LinearColor
//

type LinearColor struct {
	R float32
	G float32
	B float32
	A float32
}

func (s *LinearColor) parse(p *parser) error {
	var err error
	s.R, err = p.readFloat32()
	if err != nil {
		return err
	}

	s.G, err = p.readFloat32()
	if err != nil {
		return err
	}

	s.B, err = p.readFloat32()
	if err != nil {
		return err
	}

	s.A, err = p.readFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (s *LinearColor) serialize(p *parser) error {
	err := p.writeFloat32(s.R)
	if err != nil {
		return err
	}

	err = p.writeFloat32(s.G)
	if err != nil {
		return err
	}

	err = p.writeFloat32(s.B)
	if err != nil {
		return err
	}

	err = p.writeFloat32(s.A)
	if err != nil {
		return err
	}

	return nil
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

func (s *QuatStruct) parse(p *parser) error {
	var err error
	s.X, err = p.readFloat32()
	if err != nil {
		return err
	}

	s.Y, err = p.readFloat32()
	if err != nil {
		return err
	}

	s.Z, err = p.readFloat32()
	if err != nil {
		return err
	}

	s.W, err = p.readFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (s *QuatStruct) serialize(p *parser) error {
	err := p.writeFloat32(s.X)
	if err != nil {
		return err
	}

	err = p.writeFloat32(s.Y)
	if err != nil {
		return err
	}

	err = p.writeFloat32(s.Z)
	if err != nil {
		return err
	}

	err = p.writeFloat32(s.W)
	if err != nil {
		return err
	}

	return nil
}

//
// VectorStruct
//

type VectorStruct struct {
	X float32
	Y float32
	Z float32
}

func (s *VectorStruct) parse(p *parser) error {
	var err error
	s.X, err = p.readFloat32()
	if err != nil {
		return err
	}

	s.Y, err = p.readFloat32()
	if err != nil {
		return err
	}

	s.Z, err = p.readFloat32()
	if err != nil {
		return err
	}

	return nil
}

func (s *VectorStruct) serialize(p *parser) error {
	err := p.writeFloat32(s.X)
	if err != nil {
		return err
	}

	err = p.writeFloat32(s.Y)
	if err != nil {
		return err
	}

	err = p.writeFloat32(s.Z)
	if err != nil {
		return err
	}

	return nil
}
