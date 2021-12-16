package factory

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/l-ross/ficsit-toolkit/factory/typepath"
	BuildingDescriptor "github.com/l-ross/ficsit-toolkit/resource/building_descriptor"
	Recipe "github.com/l-ross/ficsit-toolkit/resource/recipe"
	"github.com/l-ross/ficsit-toolkit/save"
	"github.com/l-ross/ficsit-toolkit/save/property"
)

// Building is implemented by all Satisfactory buildings
type Building interface {
	InstanceName() string
	BuildTimestamp() time.Duration
	Recipe() Recipe.FGRecipe
	PrimaryColor() property.LinearColorStruct
	SecondaryColor() property.LinearColorStruct
	Entity() save.Entity
	ID() int64
	TypePath() typepath.TypePath
}

type building struct {
	id              int64
	instanceName    string
	descriptor      BuildingDescriptor.FGBuildingDescriptor
	typePath        typepath.TypePath
	entity          *save.Entity
	buildTimestamp  time.Duration
	builtWithRecipe Recipe.FGRecipe
	primaryColor    property.LinearColorStruct
	secondaryColor  property.LinearColorStruct
}

func (b *building) InstanceName() string {
	return b.instanceName
}

// TypePath returns the type path for this building.
func (b *building) TypePath() typepath.TypePath {
	return b.typePath
}

// Descriptor returns the descriptor for this building based on the data provided in
// the satisfactory Docs.json that ships with the game.
func (b *building) Descriptor() BuildingDescriptor.FGBuildingDescriptor {
	return b.descriptor
}

// ID returns the unique identifier for this building
func (b *building) ID() int64 {
	return b.id
}

// Entity returns the entity from the save file for this building.
func (b *building) Entity() save.Entity {
	return *b.entity
}

// BuildTimestamp returns the duration between when the game started and when the building
// was constructed.
func (b *building) BuildTimestamp() time.Duration {
	return b.buildTimestamp
}

// Recipe returns the recipe for constructing this building.
func (b *building) Recipe() Recipe.FGRecipe {
	return b.builtWithRecipe
}

// PrimaryColor returns the RGBA of the buildings primary color.
func (b *building) PrimaryColor() property.LinearColorStruct {
	return b.primaryColor
}

// SecondaryColor returns the RGBA of the buildings secondary color.
func (b *building) SecondaryColor() property.LinearColorStruct {
	return b.secondaryColor
}

func (f *Factory) loadBuilding(e *save.Entity, s *save.Save) (Building, error) {
	b := &building{
		id:           f.id(),
		instanceName: e.InstanceName,
		typePath:     typepath.TypePath(e.TypePath),
		entity:       e,
	}

	// Only add the building to the conveyor graph if it has a reference to an input or output.
	hasConveyor := false

	for _, ref := range e.References {
		if inputRegexp.MatchString(ref.PathName) || outputRegexp.MatchString(ref.PathName) {
			hasConveyor = true
			break
		}
	}

	if hasConveyor {
		f.conveyorGraph.AddNode(b)
	}

	for _, prop := range e.Properties {
		var err error

		switch prop.Name {
		case "mPrimaryColor":
			err = b.setPrimaryColor(prop)
		case "mSecondaryColor":
			err = b.setSecondaryColor(prop)
		case "mBuiltWithRecipe":
			err = b.setRecipe(prop)
		case "mBuildTimeStamp":
			err = b.setTimestamp(prop, s)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to handle prop %q: %w", prop.Name, err)
		}
	}

	// TODO: Check all properties have been loaded.

	return b, nil
}

var idRegexp = regexp.MustCompile(`_C_(\d+)`)

func getID(instanceName string) (int64, error) {
	match := idRegexp.FindStringSubmatch(instanceName)
	if len(match) != 2 {
		return 0, fmt.Errorf("failed to get instance ID from: %q", instanceName)
	}

	id, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, fmt.Errorf("failed to convert instance ID %q to int", match[1])
	}

	return int64(id), nil
}

func (b *building) setTimestamp(p *property.Property, s *save.Save) error {
	f, err := p.GetFloatValue()
	if err != nil {
		return err
	}

	// Timestamp is stored as a negative value, add it to the playtime
	// to get the time when the building was constructed.
	playTime := time.Duration(int64(s.Header.PlayTime) * timeMultiplier)
	t := time.Duration(f * timeMultiplier)

	b.buildTimestamp = t + playTime

	return nil
}

func (b *building) setRecipe(p *property.Property) error {
	o, err := p.GetObjectValue()
	if err != nil {
		return err
	}

	r, err := Recipe.GetByFullName(o.PathName)
	if err != nil {
		return err
	}

	b.builtWithRecipe = r

	return nil
}

func (b *building) setPrimaryColor(p *property.Property) error {
	l, err := b.getLinearColor(p)
	if err != nil {
		return err
	}

	b.primaryColor = *l

	return nil
}

func (b *building) setSecondaryColor(p *property.Property) error {
	l, err := b.getLinearColor(p)
	if err != nil {
		return err
	}

	b.secondaryColor = *l

	return nil
}

func (b *building) getLinearColor(p *property.Property) (*property.LinearColorStruct, error) {
	structVal, err := p.GetStructValue()
	if err != nil {
		return nil, err
	}

	l, err := structVal.GetLinearColor()
	if err != nil {
		return nil, err
	}

	return l, err
}
