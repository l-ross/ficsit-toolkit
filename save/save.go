// Package save is a parser and serializer for the Satisfactory save file format.
package save

import "fmt"

// A Save file for Satisfactory
type Save struct {
	// Header of the save file.
	Header *Header `json:"header"`

	// Components within the save file. The key is the InstanceName of the Component.
	//
	// If you want to manually add a Component you need to use Save.AddComponent.
	Components map[string]*Component `json:"components"`

	// Entities within the save file. The key is the InstanceName of the Entity.
	//
	// If you want to manually add an Entity you need to use Save.AddEntity.
	Entities map[string]*Entity `json:"entities"`

	// CollectedObjects is a list of all the objects the player has collected in the level.
	CollectedObjects []*ObjectReference `json:"collected_objects"`

	// Store the order we parsed entities / components in so that when serializing we
	// write them back in the same order. Makes testing easier.
	entityOrder    []string
	componentOrder []string

	objects     []object
	objectCount int32
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

// AddEntity will add the provided Entity to the Save.
//
// Entities must be added using this function and cannot just be manually added to Save.Entities.
func (s *Save) AddEntity(e *Entity) error {
	if _, exist := s.Entities[e.InstanceName]; exist {
		return fmt.Errorf("cannot add duplicate entity with InstanceName %q", e.InstanceName)
	}

	s.Entities[e.InstanceName] = e
	s.entityOrder = append(s.entityOrder, e.InstanceName)

	return nil
}

// AddComponent will add the provided Component to the Save.
//
// Components must be added using this function and cannot just be manually added to Save.Components.
func (s *Save) AddComponent(c *Component) error {
	if _, exist := s.Components[c.InstanceName]; exist {
		return fmt.Errorf("cannot add duplicate component with InstanceName %q", c.InstanceName)
	}

	s.Components[c.InstanceName] = c
	s.componentOrder = append(s.componentOrder, c.InstanceName)

	return nil
}
