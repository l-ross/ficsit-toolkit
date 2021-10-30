package save

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/ViRb3/slicewriteseek"
)

var debug = false

type Parser struct {
	r io.Reader

	// Body of the save file.
	body *slicewriteseek.SliceWriteSeeker

	parserOptions
}

func NewParser(r io.Reader, opts ...ParserOption) (*Parser, error) {
	p := &Parser{
		r: r,
	}

	for _, opt := range opts {
		if opt != nil {
			if err := opt(&p.parserOptions); err != nil {
				return nil, err
			}
		}
	}

	return p, nil
}

func (p *Parser) Parse() (*Save, error) {
	data, err := ioutil.ReadAll(p.r)
	if err != nil {
		return nil, err
	}

	p.body = &slicewriteseek.SliceWriteSeeker{
		Buffer: data,
		Index:  0,
	}

	h, err := p.parseHeader()
	if err != nil {
		return nil, err
	}

	s := &Save{
		Header:           h,
		Components:       make(map[string]*Component),
		Entities:         make(map[string]*Entity),
		CollectedObjects: make([]*ObjectReference, 0),

		objects:                  make([]object, 0),
		componentsByInstanceName: make(map[string]*Component),
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

func (p *Parser) parseBody(s *Save) error {
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

	// Parse the components and entities.
	err = p.parseObjects(s)
	if err != nil {
		return err
	}

	// Read the number of object data chunks.
	// Should be the same as the number of objects
	objectDataCount, err := p.readInt32()
	if err != nil {
		return err
	}

	if s.objectCount != objectDataCount {
		return fmt.Errorf("total objects and object data counts should be the same but were %d and %d", s.objectCount, objectDataCount)
	}

	// Parse the data for each entity and component.
	for _, o := range s.objects {
		err = o.parseData(p)
		if err != nil {
			return err
		}
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
			c := &Component{}
			err = c.parse(p)
			if err != nil {
				return err
			}

			if _, exist := s.Components[c.InstanceName]; exist {
				return fmt.Errorf("duplicate components with instance name %s", c.InstanceName)
			}

			s.Components[c.InstanceName] = c
			s.objects = append(s.objects, c)
		case EntityType:
			e := &Entity{}
			err := e.parse(p)
			if err != nil {
				return err
			}

			if _, exist := s.Entities[e.InstanceName]; exist {
				return fmt.Errorf("duplicate entities with instance name %s", e.InstanceName)
			}

			s.Entities[e.InstanceName] = e
			s.objects = append(s.objects, e)
		default:
			return fmt.Errorf("unknown object type %d", objectType)
		}
	}

	return nil
}

func (p *Parser) parseCollectedObjects(s *Save) error {
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
