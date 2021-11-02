package save

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
	parseData(p *parser) error
}

type ObjectReference struct {
	LevelName string `json:"levelName"`
	PathName  string `json:"pathName"`
}
