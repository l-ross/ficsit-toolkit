package satisfactorysave

type Save struct {
	Header           *Header            `json:"header"`
	Components       []*Component       `json:"components"`
	Entities         []*Entity          `json:"entities"`
	CollectedObjects []*CollectedObject `json:"collected_objects"`

	objData     map[int32]*dataLoc
	objectCount int32
}

type dataLoc struct {
	offset int64
	len    int32
}

type Header struct {
	HeaderVersion       int32  `json:"header_version"`
	SaveVersion         int32  `json:"save_version"`
	BuildVersion        int32  `json:"build_version"`
	MapName             string `json:"map_name"`
	MapOptions          string `json:"map_options"`
	SessionName         string `json:"session_name"`
	PlayTime            int32  `json:"play_time"`
	SaveDate            int64  `json:"save_date"`
	SessionVisibility   byte   `json:"session_visibility"`
	EditorObjectVersion int32  `json:"editor_object_version,omitempty"`
	ModMetadata         string `json:"mod_metadata,omitempty"`
	ModFlags            int32  `json:"mod_flags,omitempty"`
}

type ObjectType int32

const (
	ComponentType ObjectType = 0
	EntityType    ObjectType = 1
)

type Component struct {
	ClassName     string `json:"class_name,omitempty"`
	LevelName     string `json:"level_name,omitempty"`
	PathName      string `json:"path_name,omitempty"`
	OuterPathName string `json:"outer_path_name,omitempty"`

	order  int32
	offset int64
	len    int32

	dataOffset int64
	dataLen    int32
}

type Entity struct {
	ClassName        string    `json:"class_name,omitempty"`
	LevelName        string    `json:"level_name,omitempty"`
	PathName         string    `json:"path_name,omitempty"`
	NeedTransform    int32     `json:"need_transform,omitempty"`
	Rotation         []float32 `json:"rotation,omitempty"`
	Translation      []float32 `json:"translation,omitempty"`
	Scale3D          []float32 `json:"scale_3_d,omitempty"`
	WasPlacedInLevel int32     `json:"was_placed_in_level,omitempty"`

	order  int32
	offset int64
	len    int32

	dataOffset int64
	dataLen    int32
}

type ObjectData struct {
	LevelName  string      `json:"level_name,omitempty"`
	PathName   string      `json:"path_name,omitempty"`
	Children   []*Child    `json:"children,omitempty"`
	Properties []*Property `json:"properties,omitempty"`
}

type Child struct {
	LevelName string
	PathName  string
}

type CollectedObject struct {
	LevelName string `json:"level_name,omitempty"`
	PathName  string `json:"path_name,omitempty"`
}
