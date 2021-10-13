package satisfactorysave

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type Parser struct {
	// Decompressed body
	body *bytes.Reader
	// Counter of how many bytes have been read. Can be reset.
	counter int32
}

// NewParser constructs a new Parser to parse a Satisfactory save file.
func NewParser(r io.Reader) (*Parser, error) {
	p := &Parser{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	p.body = bytes.NewReader(data)

	return p, nil
}

// Parse will parse the entire file and return a Save object that contains
// the entire data structure of the file.
func (p *Parser) Parse() (*Save, error) {
	h, err := p.ParseHeader()
	if err != nil {
		return nil, err
	}

	s := &Save{
		Header:           h,
		Components:       make([]*Component, 0),
		Entities:         make([]*Entity, 0),
		CollectedObjects: make([]*CollectedObject, 0),
	}

	// Decompress the save body and replace p.body with it.
	p.body, err = p.decompressBody()
	if err != nil {
		return nil, err
	}

	err = p.parseBody(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (p *Parser) parseBody(s *Save) error {
	bodyLen, err := p.readInt32()
	if err != nil {
		return err
	}

	// Verify the body is the expected length.
	// Account for the fact that the specified body length does not include itself (+4 bytes).
	if bodyLen+4 != int32(p.body.Size()) {
		return fmt.Errorf("expected decompressed body to be %d but was %d", p.body.Size(), bodyLen)
	}

	err = p.parseObjects(s)
	if err != nil {
		return err
	}

	err = p.parseObjectsData(s)
	if err != nil {
		return err
	}

	err = p.parseCollectedObjects(s)
	if err != nil {
		return err
	}

	// At this point we should have reached the end of the file.
	// Check that there is no data left.
	if p.body.Len() != 0 {
		return fmt.Errorf("found %d unparsed bytes at the end of the body", p.body.Len())
	}

	return nil
}

func (p *Parser) parseObjects(s *Save) error {
	var err error
	s.objectCount, err = p.readInt32()
	if err != nil {
		return err
	}

	for i := int32(0); i < s.objectCount; i++ {
		objectType, err := p.readInt32()
		if err != nil {
			return err
		}

		switch ObjectType(objectType) {
		case ComponentType:
			c, err := p.parseComponent()
			if err != nil {
				return err
			}

			s.Components = append(s.Components, c)
		case EntityType:
			e, err := p.parseEntity()
			if err != nil {
				return err
			}

			s.Entities = append(s.Entities, e)
		default:
			return fmt.Errorf("unknown object type %d", objectType)
		}
	}

	return nil
}

func (p *Parser) parseObjectsData(s *Save) error {
	objectDataCount, err := p.readInt32()
	if err != nil {
		return err
	}

	if s.objectCount != objectDataCount {
		return fmt.Errorf("total objects and object data counts should be the same but were %d and %d", s.objectCount, objectDataCount)
	}

	for i := int32(0); i < objectDataCount; i++ {
		fmt.Println(i)
		_, err := p.parseObjectData()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseObjectData() (*ObjectData, error) {
	p.resetCounter()

	o := &ObjectData{
		offset: p.body.Size() - int64(p.body.Len()),
	}

	var err error
	o.len, err = p.readInt32()
	if err != nil {
		return nil, err
	}
	if o.len == 0 {
		return nil, nil
	}

	o.LevelName, err = p.readString()
	if err != nil {
		return nil, err
	}

	o.PathName, err = p.readString()
	if err != nil {
		return nil, err
	}

	childCount, err := p.readInt32()
	if err != nil {
		return nil, err
	}

	if childCount != 0 {
		return nil, fmt.Errorf("non zero child count %d", childCount)
	}

	o.Properties = make([]*Property, 0)

	for o.len-p.counter > 0 {
		prop, err := p.parseProperty()
		if err != nil {
			return nil, err
		}
		if prop == nil {
			// TODO: Can this happen?
			// We should only get nil if we are parsing props in a child.
			continue
		}

		o.Properties = append(o.Properties, prop)
	}

	return o, nil
}

func (p *Parser) parseProperty() (*Property, error) {
	var err error
	prop := &Property{}

	// PropName
	prop.Name, err = p.readString()
	if err != nil {
		return nil, err
	}

	if prop.Name == "None" {
		// Parsing props inside a structure of some kind and we have reached the end.
		return nil, nil
	}

	propType, err := p.readString()
	if err != nil {
		return nil, err
	}
	prop.Type = PropertyType(propType)

	prop.ValueLen, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	prop.Index, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	switch prop.Type {
	case ArrayPropertyType:
		prop.Value = &ArrayPropertyValue{}
	case BoolPropertyType:
		v := BoolPropertyValue(false)
		prop.Value = &v
	case BytePropertyType:
		prop.Value = &BytePropertyValue{}
	case DoublePropertyType:
		v := DoublePropertyValue(0)
		prop.Value = &v
	case FloatPropertyType:
		v := FloatPropertyValue(0)
		prop.Value = &v
	case Int8PropertyType:
		v := Int8PropertyValue(0)
		prop.Value = &v
	case Int64PropertyType:
		v := Int64PropertyValue(0)
		prop.Value = &v
	case IntPropertyType:
		v := IntPropertyValue(0)
		prop.Value = &v
	case InterfacePropertyType:
		prop.Value = &InterfacePropertyValue{}
	case NamePropertyType:
		v := NamePropertyValue("")
		prop.Value = &v
	case ObjectPropertyType:
		prop.Value = &ObjectPropertyValue{}
	case StringPropertyType:
		v := StringPropertyValue("")
		prop.Value = &v
	case StructPropertyType:
		prop.Value = &StructPropertyValue{}
	default:
		// TODO: Have a UnknownPropertyType where we just store the value as a byte slice.
		return nil, fmt.Errorf("unknown property type %s", propType)
	}

	err = prop.Value.parse(p, false)
	if err != nil {
		return nil, err
	}

	return prop, nil
}

func (p *Parser) parseCollectedObjects(s *Save) error {
	collectedObjectCount, err := p.readInt32()
	if err != nil {
		return err
	}

	for i := int32(0); i < collectedObjectCount; i++ {
		co := &CollectedObject{}

		co.LevelName, err = p.readString()
		if err != nil {
			return err
		}

		co.PathName, err = p.readString()
		if err != nil {
			return err
		}

		s.CollectedObjects = append(s.CollectedObjects, co)
	}

	return nil
}

func (p *Parser) parseEntity() (*Entity, error) {
	e := &Entity{}

	var err error
	e.ClassName, err = p.readString()
	if err != nil {
		return nil, err
	}

	e.LevelName, err = p.readString()
	if err != nil {
		return nil, err
	}

	e.PathName, err = p.readString()
	if err != nil {
		return nil, err
	}

	e.NeedTransform, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	e.Rotation, err = p.readFloat32Array(4)
	if err != nil {
		return nil, err
	}

	e.Translation, err = p.readFloat32Array(3)
	if err != nil {
		return nil, err
	}

	e.Scale3D, err = p.readFloat32Array(3)
	if err != nil {
		return nil, err
	}

	e.WasPlacedInLevel, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (p *Parser) parseComponent() (*Component, error) {
	c := &Component{
		// We already read the type so shift the offset back 4 bytes.
		offset: p.body.Size() - int64(p.body.Len()) - 4,
	}

	var err error

	c.ClassName, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.LevelName, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.PathName, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.OuterPathName, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.len = p.body.Size() - c.offset - int64(p.body.Len())

	return c, nil
}
