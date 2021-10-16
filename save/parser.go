package save

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type Parser struct {
	// Decompressed body
	body []byte

	buf *bytes.Reader
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

	p.buf = bytes.NewReader(data)

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

		objData: make(map[int32]*dataLoc),
	}

	// Decompress the save body and replace p.body with it.
	p.body, err = p.decompressBody()
	if err != nil {
		return nil, err
	}

	p.buf = bytes.NewReader(p.body)

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
	if bodyLen+4 != int32(p.buf.Size()) {
		return fmt.Errorf("expected decompressed body to be %d but was %d", p.buf.Size(), bodyLen)
	}

	err = p.parseObjects(s)
	if err != nil {
		return err
	}

	// Don't fully parse the object data. Just scan over to find offsets and lengths.
	err = p.scanObjectData(s)
	if err != nil {
		return err
	}

	err = p.parseCollectedObjects(s)
	if err != nil {
		return err
	}

	// At this point we should have reached the end of the file.
	// Check that there is no data left.
	if p.buf.Len() != 0 {
		return fmt.Errorf("found %d unparsed bytes at the end of the body", p.buf.Len())
	}

	for _, e := range s.Entities {
		dataLoc := s.objData[e.order]

		err = p.parseEntityData(e, dataLoc.offset, dataLoc.len)
		if err != nil {
			return err
		}
	}

	for _, c := range s.Components {
		fmt.Println(c.order)
		dataLoc := s.objData[c.order]

		err = p.parseComponentData(c, dataLoc.offset, dataLoc.len)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseComponentData(c *Component, offset int64, l int32) error {
	p.buf = bytes.NewReader(p.body[offset : offset+int64(l)])

	var err error
	//c.ParentEntityName, err = p.readString()
	//if err != nil {
	//	return err
	//}

	c.Properties, err = p.parseProperties()
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) parseEntityData(e *Entity, offset int64, l int32) error {
	p.buf = bytes.NewReader(p.body[offset : offset+int64(l)])

	var err error
	e.ParentObjectRoot, err = p.readString()
	if err != nil {
		return err
	}

	e.ParentObjectName, err = p.readString()
	if err != nil {
		return err
	}

	childCount, err := p.readInt32()
	if err != nil {
		return err
	}

	e.Children = make([]*Child, childCount)

	for i := int32(0); i < childCount; i++ {
		c := &Child{}

		c.LevelName, err = p.readString()
		if err != nil {
			return err
		}

		c.PathName, err = p.readString()
		if err != nil {
			return err
		}

		e.Children[i] = c
	}

	e.Properties, err = p.parseProperties()
	if err != nil {
		return err
	}

	// ExtraCount
	_, err = p.readInt32()
	if err != nil {
		return err
	}

	// TODO: Handle ExtraCount

	return nil
}

func (p *Parser) parseProperties() ([]*Property, error) {
	props := make([]*Property, 0)

	for {
		prop, err := p.parseProperty()
		if err != nil {
			return nil, err
		}
		if prop == nil {
			// Reached the end of the property list.
			break
		}

		props = append(props, prop)
	}

	return props, nil
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
			c.order = i
			s.Components = append(s.Components, c)
		case EntityType:
			e, err := p.parseEntity()
			if err != nil {
				return err
			}
			e.order = i
			s.Entities = append(s.Entities, e)
		default:
			return fmt.Errorf("unknown object type %d", objectType)
		}
	}

	return nil
}

func (p *Parser) scanObjectData(s *Save) error {
	objectDataCount, err := p.readInt32()
	if err != nil {
		return err
	}

	if s.objectCount != objectDataCount {
		return fmt.Errorf("total objects and object data counts should be the same but were %d and %d", s.objectCount, objectDataCount)
	}

	for i := int32(0); i < objectDataCount; i++ {
		l, err := p.readInt32()
		if err != nil {
			return err
		}
		offset := p.offset()

		s.objData[i] = &dataLoc{
			offset: offset,
			len:    l,
		}

		_, err = p.buf.Seek(offset+int64(l), io.SeekStart)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Parser) parseProperty() (*Property, error) {
	var err error
	prop := &Property{}

	prop.Name, err = p.readString()
	if err != nil {
		return nil, err
	}

	if prop.Name == "None" {
		// End of the property list.
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

	var newPropValue newPropValueFunc

	switch prop.Type {
	case ArrayPropertyType:
		newPropValue = newArrayPropertyValue
	case BoolPropertyType:
		newPropValue = newBoolPropertyValue
	case BytePropertyType:
		newPropValue = newBytePropertyValue
	case DoublePropertyType:
		newPropValue = newDoublePropertyValue
	case EnumPropertyType:
		newPropValue = newEnumPropertyValue
	case FloatPropertyType:
		newPropValue = newFloatPropertyValue
	case Int8PropertyType:
		newPropValue = newInt8PropertyValue
	case Int64PropertyType:
		newPropValue = newInt64PropertyValue
	case IntPropertyType:
		newPropValue = newIntPropertyValue
	case InterfacePropertyType:
		newPropValue = newInterfacePropertyValue
	case NamePropertyType:
		newPropValue = newNamePropertyValue
	case MapPropertyType:
		newPropValue = newMapPropertyValue
	case ObjectPropertyType:
		newPropValue = newObjectPropertyValue
	case StringPropertyType:
		newPropValue = newStringPropertyValue
	case StructPropertyType:
		newPropValue = newStructPropertyValue
	default:
		// TODO: Possibly have a UnknownPropertyType where we just store the value as a byte slice.
		return nil, fmt.Errorf("unknown property type %s", propType)
	}

	prop.PropertyValue = newPropValue()

	err = prop.PropertyValue.parse(p, false)
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
	e := &Entity{
		offset: p.offset(),
	}

	p.resetCounter()

	var err error
	e.TypePath, err = p.readString()
	if err != nil {
		return nil, err
	}

	e.RootObject, err = p.readString()
	if err != nil {
		return nil, err
	}

	e.InstanceName, err = p.readString()
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

	e.Position, err = p.readFloat32Array(3)
	if err != nil {
		return nil, err
	}

	e.Scale, err = p.readFloat32Array(3)
	if err != nil {
		return nil, err
	}

	e.WasPlacedInLevel, err = p.readInt32()
	if err != nil {
		return nil, err
	}

	e.len = p.counter + 4

	return e, nil
}

func (p *Parser) parseComponent() (*Component, error) {
	c := &Component{
		offset: p.offset(),
	}

	p.resetCounter()

	var err error

	c.TypePath, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.RootObject, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.InstanceName, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.ParentEntityName, err = p.readString()
	if err != nil {
		return nil, err
	}

	c.len = p.counter + 4

	return c, nil
}
