package save

import (
	"fmt"
	"io"

	"github.com/l-ross/ficsit-toolkit/save/data"
)

type parser struct {
	r io.Reader

	*data.Data
}

func newParser(r io.Reader) (*parser, error) {
	d, err := data.NewFromReader(r)
	if err != nil {
		return nil, err
	}

	return &parser{
		Data: d,
	}, nil
}

// Parse a Satisfactory save file
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
		Components:       make(map[string]*Component),
		Entities:         make(map[string]*Entity),
		CollectedObjects: make([]*ObjectReference, 0),

		entityOrder:    make([]string, 0),
		componentOrder: make([]string, 0),
		objects:        make([]object, 0),
	}

	// Decompress the save file and replace p.body with the decompressed version.
	p.Data, err = p.decompressBody()
	if err != nil {
		return nil, err
	}

	err = p.parseBody(s)
	if err != nil {
		return nil, fmt.Errorf("parse error at %d: %w", p.Index(), err)
	}

	return s, nil
}

// DumpBody will decompress the body of a Satisfactory save file and return
// the body as a byte slice.
//
// Primarily used for debugging.
func DumpBody(r io.Reader) ([]byte, error) {
	p, err := newParser(r)
	if err != nil {
		return nil, err
	}

	_, err = p.parseHeader()
	if err != nil {
		return nil, err
	}

	body, err := p.decompressBody()
	if err != nil {
		return nil, err
	}

	return body.Bytes(), nil
}

func (p *parser) parseBody(s *Save) error {
	bodyLen, err := p.ReadInt32()
	if err != nil {
		return err
	}

	// Add 4 to the body length to account for itself
	bodyLen += 4

	// Verify the body is the expected length.
	actualLen := p.Len()
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
	objectDataCount, err := p.ReadInt32()
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
	if p.Index() != int64(bodyLen) {
		return fmt.Errorf("found %d unparsed bytes at the end of the body", p.Index())
	}

	return nil
}

func (p *parser) parseObjects(s *Save) error {
	var err error
	s.objectCount, err = p.ReadInt32()
	if err != nil {
		return err
	}

	for i := int32(0); i < s.objectCount; i++ {
		objectType, err := p.ReadInt32()
		if err != nil {
			return err
		}

		switch objectType {
		case componentType:
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
			s.componentOrder = append(s.componentOrder, c.InstanceName)
		case entityType:
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
			s.entityOrder = append(s.entityOrder, e.InstanceName)
		default:
			return fmt.Errorf("unknown object type %d", objectType)
		}
	}

	return nil
}

func (p *parser) parseCollectedObjects(s *Save) error {
	collectedObjectCount, err := p.ReadInt32()
	if err != nil {
		return err
	}

	for i := int32(0); i < collectedObjectCount; i++ {
		o := &ObjectReference{}

		o.LevelName, err = p.ReadString()
		if err != nil {
			return err
		}

		o.PathName, err = p.ReadString()
		if err != nil {
			return err
		}

		s.CollectedObjects = append(s.CollectedObjects, o)
	}

	return nil
}
