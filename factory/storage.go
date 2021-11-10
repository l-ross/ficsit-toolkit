package factory

import (
	"fmt"

	BuildableStorage "github.com/l-ross/ficsit-toolkit/resource/buildable_storage"
	"github.com/l-ross/ficsit-toolkit/save"
)

type StorageContainer struct {
	Docs BuildableStorage.FGBuildableStorage
	*building
	*storage
}

func (f *Factory) LoadStorageContainer(e *save.Entity, save *save.Save) (*StorageContainer, error) {
	b, err := f.loadBuilding(e, save)
	if err != nil {
		return nil, err
	}

	st, err := f.loadStorage(b, save)
	if err != nil {
		return nil, err
	}

	s := &StorageContainer{
		building: b,
		storage:  st,
	}

	f.buildings[s.node.ID()] = b

	return s, nil
}

type storage struct {
	inventory []InventoryStack
}

func (s *storage) Inventory() []InventoryStack {
	return s.inventory
}

func (f *Factory) loadStorage(b *building, save *save.Save) (*storage, error) {
	s := &storage{}

	for _, prop := range b.entity.Properties {
		var err error

		switch prop.Name {
		case "mStorageInventory":
			err = s.setInventory(prop, save)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to handle prop %q: %w", prop.Name, err)
		}
	}

	return s, nil
}

func (s *storage) setInventory(prop *save.Property, save *save.Save) error {
	inv, err := getInventoryStacks(prop, save)
	if err != nil {
		return err
	}

	s.inventory = inv

	return nil
}
