package factory

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	ItemDescriptor "github.com/l-ross/ficsit-toolkit/resource/item_descriptor"
	Recipe "github.com/l-ross/ficsit-toolkit/resource/recipe"

	"github.com/l-ross/ficsit-toolkit/save"
)

type Production struct {
	PowerConsumption              float32
	CurrentRecipe                 *Recipe.FGRecipe
	InputInventory                []InventoryStack
	OutputInventory               []InventoryStack
	IsProducing                   bool
	TimeSinceStartOrStopProducing time.Duration

	Inputs  []*Connection
	Outputs []*Connection

	*Building
}

type InventoryStack struct {
	Item     ItemDescriptor.FGItemDescriptor
	Quantity int
}

func (f *Factory) loadProduction(e *save.Entity, s *save.Save) (*Production, error) {
	b, err := f.loadBuilding(e, s)
	if err != nil {
		return nil, err
	}

	p := &Production{
		Building: b,
	}

	for _, prop := range e.Properties {
		var err error
		switch prop.Name {
		case "mInputInventory":
			err = p.setInputInventory(prop, s)
		case "mOutputInventory":
			err = p.setOutputInventory(prop, s)
		case "mCurrentRecipe":
			err = p.setRecipe(prop)
		case "mPowerInfo":
			err = p.setPowerInfo(prop, s)
		case "mIsProducing":
			err = p.setIsProducing(prop)
		case "mTimeSinceStartStopProducing":
			// TODO
		}

		if err != nil {
			return nil, fmt.Errorf("failed to handle prop %q: %w", prop.Name, err)
		}
	}

	err = p.setInputs(e, s)
	if err != nil {
		return nil, fmt.Errorf("failed to set inputs: %w", err)
	}

	return p, nil
}

// Case-insensitive as sometimes 'Input' is spelt 'InPut'.
var inputRegexp = regexp.MustCompile(`(?i)Input\d$`)

func (p *Production) setInputs(e *save.Entity, s *save.Save) error {
	inputs := getObjectsThatMatch(e.References, inputRegexp)
	if len(inputs) == 0 {
		return nil
	}

	return nil
}

func (p *Production) setInputInventory(prop *save.Property, s *save.Save) error {
	stacks, err := getInventoryStacks(prop, s)
	if err != nil {
		return err
	}

	p.InputInventory = stacks

	return nil
}

func (p *Production) setOutputInventory(prop *save.Property, s *save.Save) error {
	stacks, err := getInventoryStacks(prop, s)
	if err != nil {
		return err
	}

	p.OutputInventory = stacks

	return nil
}

func getInventoryStacks(prop *save.Property, s *save.Save) ([]InventoryStack, error) {
	objValue, err := prop.GetObjectValue()
	if err != nil {
		return nil, err
	}

	component, ok := s.Components[objValue.PathName]
	if !ok {
		return nil, fmt.Errorf("failed to find component %s", objValue.PathName)
	}

	p, err := getPropFromArray("mInventoryStacks", component.Properties)
	if err != nil {
		return nil, err
	}

	arrayProp, err := p.GetArrayValue()
	if err != nil {
		return nil, err
	}

	structProps, err := arrayProp.GetStructValues()
	if err != nil {
		return nil, err
	}

	structProp := getValueFromStructProperty("InventoryStack", structProps)
	if structProp == nil {
		return nil, fmt.Errorf("failed to find InventoryStack")
	}

	as, err := structProp.GetArbitraryStruct()
	if err != nil {
		return nil, err
	}

	itemProp, err := getPropFromArray("Item", as.Properties)
	if err != nil {
		return nil, err
	}

	itemStruct, err := itemProp.GetStructValue()
	if err != nil {
		return nil, err
	}

	iiStruct, err := itemStruct.GetInventoryItemStruct()
	if err != nil {
		return nil, err
	}

	className := strings.Split(iiStruct.ItemName, ".")[1]

	i, err := ItemDescriptor.GetByClassName(className)
	if err != nil {
		return nil, err
	}

	is := InventoryStack{
		Item:     i,
		Quantity: int(iiStruct.NumItems),
	}

	return []InventoryStack{is}, nil
}

func (p *Production) setPowerInfo(prop *save.Property, s *save.Save) error {
	obj, err := prop.GetObjectValue()
	if err != nil {
		return err
	}

	component, ok := s.Components[obj.PathName]
	if !ok {
		return fmt.Errorf("failed to find component %s", obj.PathName)
	}

	consumptionProp, err := getPropFromArray("mTargetConsumption", component.Properties)
	if err != nil {
		return err
	}

	f, err := consumptionProp.GetFloatValue()
	if err != nil {
		return err
	}

	p.PowerConsumption = f

	return nil
}

func (p *Production) setRecipe(prop *save.Property) error {
	obj, err := prop.GetObjectValue()
	if err != nil {
		return err
	}

	recipe, err := Recipe.GetByFullName(obj.PathName)
	if err != nil {
		return err
	}

	p.CurrentRecipe = &recipe

	return nil
}

func (p *Production) setIsProducing(prop *save.Property) error {
	b, err := prop.GetBoolValue()
	if err != nil {
		return err
	}

	p.IsProducing = b

	return nil
}

/*
	Connector
	- Input
	- Output
*/

/*
	Connection
	Input
	Output
	Belts []Conveyor // Organised output to input
*/

type Connection struct {
	Input  Connectable
	Output Connectable
	Belts  []Conveyor
}

type Connectable interface {
	Type() string
}
