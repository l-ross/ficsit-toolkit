package save

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/ViRb3/slicewriteseek"
)

type parser struct {
	// Body of the save file.
	body *slicewriteseek.SliceWriteSeeker
}

func newParser(r io.Reader) (*parser, error) {
	p := &parser{}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	p.body = &slicewriteseek.SliceWriteSeeker{
		Buffer: data,
		Index:  0,
	}

	return p, nil
}

// Parse will parse the entire file and return a Save object that contains
// the entire data structure of the file.
func Parse(r io.Reader) (*Save, error) {
	p, err := newParser(r)
	if err != nil {
		return nil, err
	}

	h, err := p.parseHeader()
	if err != nil {
		return nil, err
	}

	s := &Save{
		Header:           h,
		Components:       make([]*Component, 0),
		Entities:         make([]*Entity, 0),
		CollectedObjects: make([]*ObjectReference, 0),

		objects: make([]object, 0),
	}

	// Decompress the save file and replace p.body with the decompressed version.
	p.body, err = p.decompressBody()
	if err != nil {
		return nil, err
	}

	err = p.parseBody(s)
	if err != nil {
		return nil, fmt.Errorf("parse error at %d: %w", p.body.Index, err)
	}

	return s, nil
}

func (p *parser) parseBody(s *Save) error {
	bodyLen, err := p.readInt32()
	if err != nil {
		return err
	}

	// Add 4 to the body length to account for itself
	bodyLen += 4

	// Verify the body is the expected length.
	actualLen := p.body.Len()
	if bodyLen != int32(actualLen) {
		return fmt.Errorf("expected decompressed body to be %d but was %d", actualLen, bodyLen)
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
	if p.body.Index != int64(bodyLen) {
		return fmt.Errorf("found %d unparsed bytes at the end of the body", p.body.Index)
	}

	// Parse data for all entities.
	for _, e := range s.Entities {
		err = p.parseEntityData(e)
		if err != nil {
			return err
		}
	}

	// Parse data for all components.
	for _, c := range s.Components {
		err = p.parseComponentData(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *parser) parseObjects(s *Save) error {
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
			s.objects = append(s.objects, c)
		case EntityType:
			e, err := p.parseEntity()
			if err != nil {
				return err
			}

			s.Entities = append(s.Entities, e)
			s.objects = append(s.objects, e)
		default:
			return fmt.Errorf("unknown object type %d", objectType)
		}
	}

	return nil
}

func (p *parser) scanObjectData(s *Save) error {
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

		offset := p.body.Index

		// Set location of this object's data.
		s.objects[i].setLoc(offset, l)

		// Move body to start of next object data chunk.
		_, err = p.body.Seek(offset+int64(l), io.SeekStart)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *parser) parseComponentData(c *Component) error {
	_, err := p.body.Seek(c.offset, io.SeekStart)
	if err != nil {
		return err
	}

	c.Properties, err = p.parseProperties()
	if err != nil {
		return err
	}

	// UNKNOWN_DATA
	_, err = p.readInt32()
	if err != nil {
		return err
	}

	return nil
}

func (p *parser) parseCollectedObjects(s *Save) error {
	collectedObjectCount, err := p.readInt32()
	if err != nil {
		return err
	}

	for i := int32(0); i < collectedObjectCount; i++ {
		o := &ObjectReference{}

		o.LevelName, err = p.readString()
		if err != nil {
			return err
		}

		o.PathName, err = p.readString()
		if err != nil {
			return err
		}

		s.CollectedObjects = append(s.CollectedObjects, o)
	}

	return nil
}
