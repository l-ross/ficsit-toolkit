// Package save is a parser and serializer for the Satisfactory save file format.
package save

// A Save file for Satisfactory
type Save struct {
	// Header of the save file.
	Header *Header `json:"header"`
	// Components within the save file. The key is the InstanceName of the Component.
	Components map[string]*Component `json:"components"`
	// Entities within the save file. The key is the InstanceName of the Entity.
	Entities         map[string]*Entity `json:"entities"`
	CollectedObjects []*ObjectReference `json:"collected_objects"`

	// Store the order we parsed entities / components in so that when serializing we
	// write them back in the same order.
	// TODO: If we add a new entity or component it won't be serialized as it's not
	//  in the below order slice.
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
