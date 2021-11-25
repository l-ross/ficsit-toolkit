package factory

import (
	"github.com/l-ross/ficsit-toolkit/factory/typepath"
	"github.com/l-ross/ficsit-toolkit/save"
)

func init() {
	loadable = make(map[typepath.TypePath]bool)

	for _, p := range prioritizedLoading {
		for t := range p {
			loadable[t] = true
		}
	}
}

// Populated by flattening prioritizedLoading to determine all the building type paths
// we are interested in loading.
var loadable map[typepath.TypePath]bool

type loader func(f *Factory, b *building, s *save.Save) (Building, error)

// When loading we need to ensure that some buildings are loaded before others.
// e.g. When loading a constructor we want to ensure that all conveyor belts and lifts
// have been loaded first so that when loading the constructor's inputs and outputs
// we can find the first non conveyor building the constructor is connected to.
var prioritizedLoading = []map[typepath.TypePath]loader{
	// Load belts
	{
		typepath.ConveyorBeltMk1: (*Factory).loadConveyorBelt,
		typepath.ConveyorBeltMk2: (*Factory).loadConveyorBelt,
		typepath.ConveyorBeltMk3: (*Factory).loadConveyorBelt,
		typepath.ConveyorBeltMk4: (*Factory).loadConveyorBelt,
		typepath.ConveyorBeltMk5: (*Factory).loadConveyorBelt,
		typepath.ConveyorLiftMk1: (*Factory).loadConveyorLift,
		typepath.ConveyorLiftMk2: (*Factory).loadConveyorLift,
		typepath.ConveyorLiftMk3: (*Factory).loadConveyorLift,
		typepath.ConveyorLiftMk4: (*Factory).loadConveyorLift,
		typepath.ConveyorLiftMk5: (*Factory).loadConveyorLift,
	},
	// Load everything else
	{
		typepath.Splitter:            (*Factory).loadSplitter,
		typepath.Merger:              (*Factory).loadMerger,
		typepath.StorageContainerMk2: (*Factory).loadStorageContainer,
		typepath.Constructor:         (*Factory).loadConstructor,
	},
}

// isLoadableBuilding will return true if the entity is a building we are interested in loading.
func isLoadableBuilding(e *save.Entity) bool {
	if e.ParentObjectName != "Persistent_Level:PersistentLevel.BuildableSubsystem" {
		return false
	}

	// Only load buildings we are interested in.
	t := typepath.TypePath(e.TypePath)

	if _, ok := loadable[t]; ok {
		return true
	}

	return false
}
