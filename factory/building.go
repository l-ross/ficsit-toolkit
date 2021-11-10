package factory

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	Recipe "github.com/l-ross/ficsit-toolkit/resource/recipe"
	"github.com/l-ross/ficsit-toolkit/save"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type Building interface {
	BuildTimestamp() time.Duration
	BuiltWithRecipe() Recipe.FGRecipe
	PrimaryColor() save.LinearColor
	SecondaryColor() save.LinearColor
	Entity() save.Entity
	Node() graph.Node
}

type building struct {
	entity *save.Entity

	buildTimestamp  time.Duration
	builtWithRecipe Recipe.FGRecipe
	primaryColor    save.LinearColor
	secondaryColor  save.LinearColor

	inputs  []graph.Edge
	outputs []graph.Edge

	node graph.Node
}

func (b *building) Node() graph.Node {
	return b.node
}

func (b *building) Entity() save.Entity {
	return *b.entity
}

func (b *building) BuildTimestamp() time.Duration {
	return b.buildTimestamp
}

func (b *building) BuiltWithRecipe() Recipe.FGRecipe {
	return b.builtWithRecipe
}

func (b *building) PrimaryColor() save.LinearColor {
	return b.primaryColor
}

func (b *building) SecondaryColor() save.LinearColor {
	return b.secondaryColor
}

func (f *Factory) loadBuilding(e *save.Entity, s *save.Save) (*building, error) {
	id, err := getID(e.InstanceName)
	if err != nil {
		return nil, err
	}

	b := &building{
		entity: e,
		node:   simple.Node(id),
	}

	// Add node if it hasn't already been added.
	if f.conveyorGraph.Node(id) == nil {
		f.conveyorGraph.AddNode(b.node)
	}

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

	err = f.setInputs(b, s)
	if err != nil {
		return nil, err
	}

	err = f.setOutputs(b, s)
	if err != nil {
		return nil, err
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

func (b *building) setBuiltWithRecipe(p *save.Property) error {
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

// Case-insensitive as sometimes 'Input' is spelt 'InPut'.
var inputRegexp = regexp.MustCompile(`(?i)(Input\d|ConveyorAny0)$`)
var outputRegexp = regexp.MustCompile(`(?i)(Output\d|ConveyorAny1)$`)

func (f *Factory) setInputs(b *building, s *save.Save) error {
	inputs := getObjectsThatMatch(b.entity.References, inputRegexp)
	for _, i := range inputs {
		c := s.Components[i]
		e, err := f.LoadConnection(c)
		if err != nil {
			return err
		}
		b.inputs = append(b.inputs, e)
	}
	return nil
}

func (f *Factory) setOutputs(b *building, s *save.Save) error {
	outputs := getObjectsThatMatch(b.entity.References, outputRegexp)
	for _, o := range outputs {
		c := s.Components[o]
		e, err := f.LoadConnection(c)
		if err != nil {
			return err
		}
		b.outputs = append(b.outputs, e)
	}
	return nil
}

func (f *Factory) createConnection(e graph.Edge) (*Connection, error) {
	b := f.next(e.From())

	c := &Connection{
		Connected: b,
	}

	return c, nil
}
