package factory

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/factory/typepath"
	BuildableStorage "github.com/l-ross/ficsit-toolkit/resource/buildable_storage"
	"github.com/l-ross/ficsit-toolkit/save"
)

type StorageContainer struct {
	Mark              int
	StorageDescriptor BuildableStorage.FGBuildableStorage

	*building
	*storage
}

func (f *Factory) loadStorageContainer(b *building, save *save.Save) (Building, error) {
	st, err := f.loadStorage(b, save)
	if err != nil {
		return nil, err
	}

	s := &StorageContainer{
		building: b,
		storage:  st,
	}

	switch b.typePath {
	case typepath.StorageContainerMk1:
		s.StorageDescriptor = BuildableStorage.StorageContainerMk1
		s.Mark = 1
	case typepath.StorageContainerMk2:
		s.StorageDescriptor = BuildableStorage.StorageContainerMk2
		s.Mark = 2
	default:
		return nil, fmt.Errorf("unknown storage type path %s", b.typePath)
	}

	f.StorageContainers[b.InstanceName()] = s

	return s, nil
}
