package factory

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/save"
	"github.com/l-ross/ficsit-toolkit/save/property"
)

type storage struct {
	inventory []*InventoryStack
}

// Inventory returns the contents of this building.
func (s *storage) Inventory() []*InventoryStack {
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

func (s *storage) setInventory(prop *property.Property, save *save.Save) error {
	inv, err := getInventoryStacks(prop, save)
	if err != nil {
		return err
	}

	s.inventory = inv

	return nil
}
