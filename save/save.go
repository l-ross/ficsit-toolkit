package save

import (
	"io"
)

type Save struct {
	Header           *Header               `json:"header"`
	Components       map[string]*Component `json:"components"`
	Entities         map[string]*Entity    `json:"entities"`
	CollectedObjects []*ObjectReference    `json:"collected_objects"`

	objects                  []object
	objectCount              int32
	componentsByInstanceName map[string]*Component
}

type ObjectType int32

const (
	ComponentType ObjectType = 0
	EntityType    ObjectType = 1
)

type object interface {
	parseData(p *Parser) error
}

type ObjectReference struct {
	LevelName string `json:"levelName"`
	PathName  string `json:"pathName"`
}

// Parse will parse the entire file and return a Save object that contains
// the entire data structure of the file.
func Parse(r io.Reader) (*Save, error) {
	p, err := NewParser(r)
	if err != nil {
		return nil, err
	}

	return p.Parse()
}

func Serialize(save *Save, w io.Writer) error {
	s, err := NewSerializer(w)
	if err != nil {
		return err
	}

	return s.Serialize(save)
}
