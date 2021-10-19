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

func (s *StructPropertyValue) GetArbitraryStruct() (*ArbitraryStruct, error) {
	if v, ok := s.Value.(*ArbitraryStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", s.Type)
}

func (s *ArbitraryStruct) parse(p *Parser) error {
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

func (s *ArbitraryStruct) serialize(p *Parser) error {
	for _, prop := range s.Properties {
		err := p.serializeProperty(prop)
		if err != nil {
			return err
		}
	}

	err := p.writeNoneProp()
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

func (s *StructPropertyValue) GetBoxStruct() (*BoxStruct, error) {
	if v, ok := s.Value.(*BoxStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", s.Type)
}

func (s *BoxStruct) parse(p *Parser) error {
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

func (s *BoxStruct) serialize(p *Parser) error {
	return nil
}

//
// FluidBox
//

type FluidBoxStruct float32

func (s *FluidBoxStruct) parse(p *Parser) error {
	v, err := p.readFloat32()
	if err != nil {
		return err
	}

	*s = FluidBoxStruct(v)

	return nil
}

func (s *FluidBoxStruct) serialize(p *Parser) error {
	return nil
}

//
// InventoryItem
//

type InventoryItemStruct struct {
	ItemName string
	NumItems int32
}

func (s *StructPropertyValue) GetInventoryItemStruct() (*InventoryItemStruct, error) {
	if v, ok := s.Value.(*InventoryItemStruct); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong type %s", s.Type)
}

func (s *InventoryItemStruct) parse(p *Parser) error {
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

func (s *InventoryItemStruct) serialize(p *Parser) error {
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

func (s *LinearColor) parse(p *Parser) error {
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

func (s *LinearColor) serialize(p *Parser) error {
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

func (s *QuatStruct) parse(p *Parser) error {
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

func (s *QuatStruct) serialize(p *Parser) error {
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

func (s *VectorStruct) parse(p *Parser) error {
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

func (s *VectorStruct) serialize(p *Parser) error {
	return nil
}
