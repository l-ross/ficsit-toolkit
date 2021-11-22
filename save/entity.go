package save

import (
	"github.com/l-ross/ficsit-toolkit/save/extra"
	"github.com/l-ross/ficsit-toolkit/save/property"
)

type Entity struct {
	TypePath         string               `json:"typePath"`
	RootObject       string               `json:"rootObject"`
	InstanceName     string               `json:"instanceName"`
	NeedTransform    int32                `json:"needTransform"`
	Rotation         []float32            `json:"rotation"`
	Position         []float32            `json:"position"`
	Scale            []float32            `json:"scale"`
	WasPlacedInLevel int32                `json:"wasPlacedInLevel"`
	ParentObjectRoot string               `json:"parentObjectRoot"`
	ParentObjectName string               `json:"parentObjectName"`
	References       []*ObjectReference   `json:"component"`
	Properties       []*property.Property `json:"property"`
	Extra            *extra.Extra         `json:"extras"`
}

func (e *Entity) parse(p *parser) error {
	var err error
	e.TypePath, err = p.ReadString()
	if err != nil {
		return err
	}

	e.RootObject, err = p.ReadString()
	if err != nil {
		return err
	}

	e.InstanceName, err = p.ReadString()
	if err != nil {
		return err
	}

	e.NeedTransform, err = p.ReadInt32()
	if err != nil {
		return err
	}

	e.Rotation, err = p.ReadFloat32Array(4)
	if err != nil {
		return err
	}

	e.Position, err = p.ReadFloat32Array(3)
	if err != nil {
		return err
	}

	e.Scale, err = p.ReadFloat32Array(3)
	if err != nil {
		return err
	}

	e.WasPlacedInLevel, err = p.ReadInt32()
	if err != nil {
		return err
	}

	return nil
}

func (e *Entity) parseData(p *parser) error {
	dataSize, err := p.ReadInt32()
	if err != nil {
		return err
	}

	m := p.Measure()

	e.ParentObjectRoot, err = p.ReadString()
	if err != nil {
		return err
	}

	e.ParentObjectName, err = p.ReadString()
	if err != nil {
		return err
	}

	childCount, err := p.ReadInt32()
	if err != nil {
		return err
	}

	e.References = make([]*ObjectReference, childCount)

	for i := int32(0); i < childCount; i++ {
		o := &ObjectReference{}

		o.LevelName, err = p.ReadString()
		if err != nil {
			return err
		}

		o.PathName, err = p.ReadString()
		if err != nil {
			return err
		}

		e.References[i] = o
	}

	e.Properties, err = property.ParseProperties(p.Data)
	if err != nil {
		return err
	}

	// ExtraCount
	// Is this ever not zero?
	_, err = p.ReadInt32()
	if err != nil {
		return err
	}

	// If we have extra data then parse it.
	extraLen := dataSize - m()
	if extraLen > 0 {
		e.Extra, err = extra.Parse(e.TypePath, extraLen, p.Data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Entity) serialize(s *serializer) error {
	err := s.WriteString(e.TypePath)
	if err != nil {
		return err
	}

	err = s.WriteString(e.RootObject)
	if err != nil {
		return err
	}

	err = s.WriteString(e.InstanceName)
	if err != nil {
		return err
	}

	err = s.WriteInt32(e.NeedTransform)
	if err != nil {
		return err
	}

	err = s.WriteFloat32Array(e.Rotation)
	if err != nil {
		return err
	}

	err = s.WriteFloat32Array(e.Position)
	if err != nil {
		return err
	}

	err = s.WriteFloat32Array(e.Scale)
	if err != nil {
		return err
	}

	err = s.WriteInt32(e.WasPlacedInLevel)
	if err != nil {
		return err
	}

	return nil
}

func (e *Entity) serializeData(s *serializer) error {
	// Write placeholder length
	lenPos := s.Index()
	err := s.WriteInt32(0)
	if err != nil {
		return err
	}

	m := s.Measure()

	err = s.WriteString(e.ParentObjectRoot)
	if err != nil {
		return err
	}

	err = s.WriteString(e.ParentObjectName)
	if err != nil {
		return err
	}

	err = s.WriteInt32(int32(len(e.References)))
	if err != nil {
		return err
	}

	for _, c := range e.References {
		err = s.WriteString(c.LevelName)
		if err != nil {
			return err
		}

		err = s.WriteString(c.PathName)
		if err != nil {
			return err
		}
	}

	err = property.SerializeProperties(e.Properties, s.Data)
	if err != nil {
		return err
	}

	// Write ExtraCount
	// Is this always zero?
	err = s.WriteInt32(0)
	if err != nil {
		return err
	}

	if e.Extra != nil {
		err = e.Extra.Serialize(s.Data)
		if err != nil {
			return err
		}
	}

	// Update length
	err = s.WriteLen(m(), lenPos)
	if err != nil {
		return err
	}

	return nil
}
