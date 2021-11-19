package factory

import (
	"github.com/l-ross/ficsit-toolkit/factory/typepath"
	"github.com/l-ross/ficsit-toolkit/save"
)

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
