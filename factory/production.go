package factory

import (
	"fmt"
	"time"

	Recipe "github.com/l-ross/ficsit-toolkit/resource/recipe"
	"github.com/l-ross/ficsit-toolkit/save"
	"github.com/l-ross/ficsit-toolkit/save/property"
)

// Production is implemented by all Satisfactory production buildings.
// e.g. Constructor and Assembler
type Production interface {
	PowerConsumption() float32
	CurrentRecipe() *Recipe.FGRecipe
	InputInventory() []InventoryStack
	OutputInventory() []InventoryStack
	IsProducing() bool
	TimeSinceStartOrStopProducing() time.Duration
}

type production struct {
	powerConsumption              float32
	currentRecipe                 *Recipe.FGRecipe
	inputInventory                []InventoryStack
	outputInventory               []InventoryStack
	isProducing                   bool
	timeSinceStartOrStopProducing time.Duration
}

// PowerConsumption returns the power consumption of this building.
func (p *production) PowerConsumption() float32 {
	return p.powerConsumption
}

// CurrentRecipe returns the currently assigned recipe for this production machine.
// Returns nil if the machine has no recipe assigned.
func (p *production) CurrentRecipe() *Recipe.FGRecipe {
	return p.currentRecipe
}

// InputInventory returns the input inventory of this building.
func (p *production) InputInventory() []InventoryStack {
	return p.inputInventory
}

// OutputInventory returns the output inventory of this building.
func (p *production) OutputInventory() []InventoryStack {
	return p.outputInventory
}

// IsProducing returns true if this building is currently producing.
func (p *production) IsProducing() bool {
	return p.isProducing
}

// TimeSinceStartOrStopProducing returns the amount of time that has elapsed
// since the building started or stopped producing.
func (p *production) TimeSinceStartOrStopProducing() time.Duration {
	return p.timeSinceStartOrStopProducing
}

func (f *Factory) loadProduction(b *building, s *save.Save) (*production, error) {
	p := &production{}

	for _, prop := range b.entity.Properties {
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
			err = p.setTimeSinceStartOrStopProducing(prop)
		}

		if err != nil {
			return nil, fmt.Errorf("failed to handle prop %q: %w", prop.Name, err)
		}
	}

	// TODO: Check all property have been loaded.

	return p, nil
}

func (p *production) setInputInventory(prop *property.Property, s *save.Save) error {
	stacks, err := getInventoryStacks(prop, s)
	if err != nil {
		return err
	}

	p.inputInventory = stacks

	return nil
}

func (p *production) setOutputInventory(prop *property.Property, s *save.Save) error {
	stacks, err := getInventoryStacks(prop, s)
	if err != nil {
		return err
	}

	p.outputInventory = stacks

	return nil
}

func getInventoryStacks(prop *property.Property, s *save.Save) ([]InventoryStack, error) {
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

	stackProps := getValuesFromStructProperty("InventoryStack", structProps)
	if len(structProps) == 0 {
		return nil, fmt.Errorf("failed to find InventoryStack")
	}

	stacks := make([]InventoryStack, 0)

	for _, p := range stackProps {
		as, err := p.GetArbitraryStruct()
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

		if iiStruct.ItemName == "" {
			continue
		}

		i, err := getItemDescriptorByName(iiStruct.ItemName)
		if err != nil {
			return nil, err
		}

		is := InventoryStack{
			Item:     i,
			Quantity: int(iiStruct.NumItems),
		}

		stacks = append(stacks, is)
	}

	return stacks, nil
}

func (p *production) setPowerInfo(prop *property.Property, s *save.Save) error {
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

	p.powerConsumption = f

	return nil
}

func (p *production) setRecipe(prop *property.Property) error {
	obj, err := prop.GetObjectValue()
	if err != nil {
		return err
	}

	recipe, err := Recipe.GetByFullName(obj.PathName)
	if err != nil {
		return err
	}

	p.currentRecipe = &recipe

	return nil
}

func (p *production) setIsProducing(prop *property.Property) error {
	b, err := prop.GetBoolValue()
	if err != nil {
		return err
	}

	p.isProducing = b

	return nil
}

func (p *production) setTimeSinceStartOrStopProducing(prop *property.Property) error {
	f, err := prop.GetFloatValue()
	if err != nil {
		return err
	}

	p.timeSinceStartOrStopProducing = time.Duration(f * timeMultiplier)

	return nil
}
