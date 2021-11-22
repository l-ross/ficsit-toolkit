// Package save is a parser and serializer for the Satisfactory save file format.
package save

// A Save file for Satisfactory
type Save struct {
	Header           *Header               `json:"header"`
	Components       map[string]*Component `json:"components"`
	Entities         map[string]*Entity    `json:"entities"`
	CollectedObjects []*ObjectReference    `json:"collected_objects"`

	// Store the order we parsed entities / components in so that when serializing we
	// write them back in the same order.
	entityOrder    []string
	componentOrder []string

	objects                  []object
	objectCount              int32
	componentsByInstanceName map[string]*Component
}

const (
	componentType int32 = 0
	entityType    int32 = 1
)

type object interface {
	parseData(p *parser) error
}

type ObjectReference struct {
	LevelName string `json:"levelName"`
	PathName  string `json:"pathName"`
}
