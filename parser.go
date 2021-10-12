package satisfactorysave

import (
	"bytes"
	"fmt"
	"io"
)

type Parser struct {
	r io.Reader
}

// NewParser constructs a new Parser to parse a Satisfactory save file.
func NewParser(r io.Reader) (*Parser, error) {
	p := &Parser{
		r: r,
	}

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

	b, err := p.decompressBody()
	if err != nil {
		return nil, err
	}

	err = p.parseBody(b, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (p *Parser) parseBody(r *bytes.Reader, s *Save) error {
	err := parseObjects(r, s)
	if err != nil {
		return err
	}

	fmt.Println(r.Size() - int64(r.Len()))

	err = parseObjectsData(r, s)
	if err != nil {
		return err
	}

	err = parseCollectedObjects(r, s)
	if err != nil {
		return err
	}

	// At this point we should have reached the end of the file.
	// Check that there is no data left.
	if r.Len() != 0 {
		return fmt.Errorf("found %d unparsed bytes at the end of the body", r.Len())
	}

	return nil
}

func parseObjects(b *bytes.Reader, s *Save) error {
	var err error
	s.objectCount, err = readInt32(b)
	if err != nil {
		return err
	}

	for i := int32(0); i < s.objectCount; i++ {
		objectType, err := readInt32(b)
		if err != nil {
			return err
		}

		switch ObjectType(objectType) {
		case ComponentType:
			c, err := parseComponent(b)
			if err != nil {
				return err
			}

			s.Components = append(s.Components, c)
		case EntityType:
			e, err := parseEntity(b)
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

func parseObjectsData(r *bytes.Reader, s *Save) error {
	objectDataCount, err := readInt32(r)
	if err != nil {
		return err
	}

	if s.objectCount != objectDataCount {
		return fmt.Errorf("total objects and object data counts should be the same but were %d and %d", s.objectCount, objectDataCount)
	}

	for i := int32(0); i < objectDataCount; i++ {
		_, err := parseObjectData(r)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseObjectData(r *bytes.Reader) (*ObjectData, error) {
	o := &ObjectData{}

	o.offset = r.Size() - int64(r.Len())
	startPos := r.Len()

	var err error
	o.len, err = readInt32(r)
	if err != nil {
		return nil, err
	}

	o.LevelName, err = readString(r)
	if err != nil {
		return nil, err
	}

	o.PathName, err = readString(r)
	if err != nil {
		return nil, err
	}

	childCount, err := readInt32(r)
	if err != nil {
		return nil, err
	}

	if childCount != 0 {
		return nil, fmt.Errorf("non zero child count %d", childCount)
	}

	o.Properties = make([]*Property, 0)

	rem := startPos - r.Len()

	for rem > 0 {
		pos := r.Len()

		// PropName
		name, err := readString(r)
		if err != nil {
			return nil, err
		}

		propType, err := readString(r)
		if err != nil {
			return nil, err
		}

		valueLen, err := readInt32(r)
		if err != nil {
			return nil, err
		}

		index, err := readInt32(r)
		if err != nil {
			return nil, err
		}

		p := &Property{
			Name:     name,
			Type:     PropertyType(propType),
			ValueLen: valueLen,
			Index:    index,
		}

		switch p.Type {
		case BoolPropertyType:
			v := BoolPropertyValue(false)
			p.Value = &v
		case BytePropertyType:
			p.Value = &BytePropertyValue{}
		case DoublePropertyType:
			v := DoublePropertyValue(0)
			p.Value = &v
		case FloatPropertyType:
			v := FloatPropertyValue(0)
			p.Value = &v
		case Int8PropertyType:
			v := Int8PropertyValue(0)
			p.Value = &v
		case Int64PropertyType:
			v := Int64PropertyValue(0)
			p.Value = &v
		case IntPropertyType:
			v := IntPropertyValue(0)
			p.Value = &v
		case NamePropertyType:
			v := NamePropertyValue("")
			p.Value = &v
		case ObjectPropertyType:
			p.Value = &ObjectPropertyValue{}
		case StringPropertyType:
			v := StringPropertyValue("")
			p.Value = &v
		default:
			// TODO: Have a UnknownPropertyType where we just store the value as a byte slice.
			return nil, fmt.Errorf("unknown property type %s", propType)
		}

		err = p.Value.Parse(r)
		if err != nil {
			return nil, err
		}

		o.Properties = append(o.Properties, p)
		rem = rem - (pos - r.Len())
	}

	//// For now just skip parsing the entity internals.
	//seekLoc := o.offset + int64(o.len) + 4
	//_, err = r.Seek(seekLoc, io.SeekStart)
	//if err != nil {
	//	return nil, err
	//}

	return o, nil
}

func parseCollectedObjects(r *bytes.Reader, s *Save) error {
	collectedObjectCount, err := readInt32(r)
	if err != nil {
		return err
	}

	for i := int32(0); i < collectedObjectCount; i++ {
		co := &CollectedObject{}

		co.LevelName, err = readString(r)
		if err != nil {
			return err
		}

		co.PathName, err = readString(r)
		if err != nil {
			return err
		}

		s.CollectedObjects = append(s.CollectedObjects, co)
	}

	return nil
}

func parseEntity(r io.Reader) (*Entity, error) {
	e := &Entity{}

	var err error
	e.ClassName, err = readString(r)
	if err != nil {
		return nil, err
	}

	e.LevelName, err = readString(r)
	if err != nil {
		return nil, err
	}

	e.PathName, err = readString(r)
	if err != nil {
		return nil, err
	}

	e.NeedTransform, err = readInt32(r)
	if err != nil {
		return nil, err
	}

	e.Rotation, err = readFloat32Array(r, 4)
	if err != nil {
		return nil, err
	}

	e.Translation, err = readFloat32Array(r, 3)
	if err != nil {
		return nil, err
	}

	e.Scale3D, err = readFloat32Array(r, 3)
	if err != nil {
		return nil, err
	}

	e.WasPlacedInLevel, err = readInt32(r)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func parseComponent(b *bytes.Reader) (*Component, error) {
	c := &Component{
		// We already read the type so shift the offset back 4 bytes.
		offset: b.Size() - int64(b.Len()) - 4,
	}

	var err error

	c.ClassName, err = readString(b)
	if err != nil {
		return nil, err
	}

	c.LevelName, err = readString(b)
	if err != nil {
		return nil, err
	}

	c.PathName, err = readString(b)
	if err != nil {
		return nil, err
	}

	c.OuterPathName, err = readString(b)
	if err != nil {
		return nil, err
	}

	c.len = b.Size() - c.offset - int64(b.Len())

	return c, nil
}
