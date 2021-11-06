package factory

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	Recipe "github.com/l-ross/ficsit-toolkit/resource/recipe"
	"github.com/l-ross/ficsit-toolkit/save"
	"gonum.org/v1/gonum/graph/simple"
)

type Building struct {
	Entity *save.Entity

	BuildTimestamp  time.Duration
	BuiltWithRecipe Recipe.FGRecipe
	PrimaryColor    save.LinearColor
	SecondaryColor  save.LinearColor

	n simple.Node
}

func (f *Factory) loadBuilding(e *save.Entity, s *save.Save) (*Building, error) {
	id, err := getID(e.InstanceName)
	if err != nil {
		return nil, err
	}

	b := &Building{
		Entity: e,
		n:      simple.Node(id),
	}

	f.g.AddNode(b.n)

	for _, prop := range e.Properties {
		var err error

		switch prop.Name {
		case "mPrimaryColor":
			err = b.setPrimaryColor(prop)
		case "mSecondaryColor":
			err = b.setSecondaryColor(prop)
		case "mBuiltWithRecipe":
			err = b.setBuiltWithRecipe(prop)
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

func (b *Building) setTimestamp(p *save.Property, s *save.Save) error {
	f, err := p.GetFloatValue()
	if err != nil {
		return err
	}

	// Timestamp is stored as a negative value, add it to the playtime
	// to get the time when the building was constructed.
	playTime := time.Duration(int64(s.Header.PlayTime) * 1_000_000_000)
	t := time.Duration(f * 1_000_000_000)

	b.BuildTimestamp = t + playTime

	return nil
}

func (b *Building) setBuiltWithRecipe(p *save.Property) error {
	o, err := p.GetObjectValue()
	if err != nil {
		return err
	}

	r, err := Recipe.GetByFullName(o.PathName)
	if err != nil {
		return err
	}

	b.BuiltWithRecipe = r

	return nil
}

func (b *Building) setPrimaryColor(p *save.Property) error {
	l, err := b.getLinearColor(p)
	if err != nil {
		return err
	}

	b.PrimaryColor = *l

	return nil
}

func (b *Building) setSecondaryColor(p *save.Property) error {
	l, err := b.getLinearColor(p)
	if err != nil {
		return err
	}

	b.SecondaryColor = *l

	return nil
}

func (b *Building) getLinearColor(p *save.Property) (*save.LinearColor, error) {
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
