package save

import "fmt"

type TextType byte

const (
	NoneTextType TextType = 0xFF
	BaseTextType TextType = 0x00
)

type TextValue interface {
	parse(p *parser) error

	serialize(s *serializer) error
}

func getText(t TextType) func() TextValue {
	return nil
}

//
// None
//

type NoneText struct {
	CultureInvariantString int32
	String                 string
}

func (t *TextPropertyValue) GetNoneText() (*NoneText, error) {
	if v, ok := t.Value.(*NoneText); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong text type: %v", t.Type)
}

func (t *NoneText) parse(p *parser) error {
	var err error
	t.CultureInvariantString, err = p.readInt32()
	if err != nil {
		return err
	}

	// TODO: Switch on above?
	t.String, err = p.readString()
	if err != nil {
		return err
	}

	return nil
}

func (t *NoneText) serialize(s *serializer) error {
	err := s.writeInt32(t.CultureInvariantString)
	if err != nil {
		return err
	}

	// TODO: Switch on above?
	err = s.writeString(t.String)
	if err != nil {
		return err
	}

	return nil
}

//
// Base
//

type BaseText struct {
	Namespace string
	Key       string
	Value     string
}

func (t *TextPropertyValue) GetBaseText() (*BaseText, error) {
	if v, ok := t.Value.(*BaseText); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong text type: %v", t.Type)
}

func (t *BaseText) parse(p *parser) error {
	var err error
	t.Namespace, err = p.readString()
	if err != nil {
		return err
	}

	t.Key, err = p.readString()
	if err != nil {
		return err
	}

	t.Value, err = p.readString()
	if err != nil {
		return err
	}

	return nil
}

func (t *BaseText) serialize(s *serializer) error {
	err := s.writeString(t.Namespace)
	if err != nil {
		return err
	}

	err = s.writeString(t.Key)
	if err != nil {
		return err
	}

	err = s.writeString(t.Value)
	if err != nil {
		return err
	}

	return nil
}
