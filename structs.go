package satisfactorysave

import "fmt"

type StructType string

const (
	BoxStructType           StructType = "Box"
	InventoryItemStructType StructType = "InventoryItem"
	LinearColorStructType   StructType = "LinearColor"
	QuatStructType          StructType = "Quat"
	VectorStructType        StructType = "Vector"
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
	for i := int32(0); ; i++ {
		if s.numProps != 0 && i < s.numProps {
			break
		}

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
	// Unknown
	_, err := p.readInt32()
	if err != nil {
		return err
	}

	s.ItemName, err = p.readString()
	if err != nil {
		return err
	}

	// Unknown
	_, err = p.readString()
	if err != nil {
		return err
	}

	// Unknown
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

	s.NumItems, err = p.readInt32()
	if err != nil {
		return err
	}

	_, err = p.readBytes(9)
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
