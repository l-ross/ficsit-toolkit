package save

import "io"

type Save struct {
	Header           *Header            `json:"header"`
	Components       []*Component       `json:"components"`
	Entities         []*Entity          `json:"entities"`
	CollectedObjects []*ObjectReference `json:"collected_objects"`

	objects     []object
	objectCount int32

	w io.Writer
}

type Header struct {
	HeaderVersion       int32  `json:"headerVersion"`
	SaveVersion         int32  `json:"saveVersion"`
	BuildVersion        int32  `json:"buildVersion"`
	MapName             string `json:"mapName"`
	MapOptions          string `json:"mapOptions"`
	SessionName         string `json:"sessionName"`
	PlayTime            int32  `json:"playTime"`
	SaveDate            int64  `json:"saveDate"`
	SessionVisibility   byte   `json:"sessionVisibility"`
	EditorObjectVersion int32  `json:"editorObjectVersion"`
	ModMetadata         string `json:"modMetadata"`
	ModFlags            int32  `json:"modFlags"`
}

type ObjectType int32

const (
	ComponentType ObjectType = 0
	EntityType    ObjectType = 1
)

type object interface {
	setLoc(offset int64, len int32)
}

type Component struct {
	TypePath         string      `json:"typePath"`
	RootObject       string      `json:"rootObject"`
	InstanceName     string      `json:"instanceName"`
	ParentEntityName string      `json:"parentEntityName"`
	Properties       []*Property `json:"properties"`

	offset int64
	len    int32
}

func (c *Component) setLoc(offset int64, len int32) {
	c.offset = offset
	c.len = len
}

type Entity struct {
	TypePath         string             `json:"typePath"`
	RootObject       string             `json:"rootObject"`
	InstanceName     string             `json:"instanceName"`
	NeedTransform    int32              `json:"needTransform"`
	Rotation         []float32          `json:"rotation"`
	Position         []float32          `json:"position"`
	Scale            []float32          `json:"scale"`
	WasPlacedInLevel int32              `json:"wasPlacedInLevel"`
	ParentObjectRoot string             `json:"parentObjectRoot"`
	ParentObjectName string             `json:"parentObjectName"`
	Children         []*ObjectReference `json:"children"`
	Properties       []*Property        `json:"properties"`
	Extra            *Extra             `json:"extras"`

	offset int64
	len    int32
}

func (e *Entity) setLoc(offset int64, len int32) {
	e.offset = offset
	e.len = len
}

type ObjectReference struct {
	LevelName string `json:"levelName"`
	PathName  string `json:"pathName"`
}
