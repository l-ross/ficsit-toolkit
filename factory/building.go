package factory

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	BuildingDescriptor "github.com/l-ross/ficsit-toolkit/resource/building_descriptor"

	"gonum.org/v1/gonum/graph/simple"

	Recipe "github.com/l-ross/ficsit-toolkit/resource/recipe"
	"github.com/l-ross/ficsit-toolkit/save"
)

type Building interface {
	BuildTimestamp() time.Duration
	Recipe() Recipe.FGRecipe
	PrimaryColor() save.LinearColor
	SecondaryColor() save.LinearColor
	Entity() save.Entity
	ID() int64
	TypePath() string
}

type building struct {
	id              int64
	descriptor      BuildingDescriptor.FGBuildingDescriptor
	typePath        string
	entity          *save.Entity
	buildTimestamp  time.Duration
	builtWithRecipe Recipe.FGRecipe
	primaryColor    save.LinearColor
	secondaryColor  save.LinearColor
}

// TypePath returns the type path for this building.
func (b *building) TypePath() string {
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
func (b *building) PrimaryColor() save.LinearColor {
	return b.primaryColor
}

// SecondaryColor returns the RGBA of the buildings secondary color.
func (b *building) SecondaryColor() save.LinearColor {
	return b.secondaryColor
}

func (f *Factory) loadBuilding(e *save.Entity, s *save.Save) (Building, error) {
	id, err := getID(e.InstanceName)
	if err != nil {
		return nil, err
	}

	b := &building{
		id:       id,
		typePath: e.TypePath,
		entity:   e,
	}

	// TODO: This will currently add things that don't have conveyor inputs / outputs
	//  to the conveyorGraph e.g. Power poles.
	f.conveyorGraph.AddNode(simple.Node(id))

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

func (b *building) setTimestamp(p *save.Property, s *save.Save) error {
	f, err := p.GetFloatValue()
	if err != nil {
		return err
	}

	// Timestamp is stored as a negative value, add it to the playtime
	// to get the time when the building was constructed.
	playTime := time.Duration(int64(s.Header.PlayTime) * 1_000_000_000)
	t := time.Duration(f * 1_000_000_000)

	b.buildTimestamp = t + playTime

	return nil
}

func (b *building) setRecipe(p *save.Property) error {
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

func (b *building) setPrimaryColor(p *save.Property) error {
	l, err := b.getLinearColor(p)
	if err != nil {
		return err
	}

	b.primaryColor = *l

	return nil
}

func (b *building) setSecondaryColor(p *save.Property) error {
	l, err := b.getLinearColor(p)
	if err != nil {
		return err
	}

	b.secondaryColor = *l

	return nil
}

func (b *building) getLinearColor(p *save.Property) (*save.LinearColor, error) {
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
