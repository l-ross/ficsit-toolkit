package factory

import (
	"fmt"
	"strings"
	"time"

	ItemDescriptor "github.com/l-ross/ficsit-toolkit/resource/item_descriptor"
	Recipe "github.com/l-ross/ficsit-toolkit/resource/recipe"

	"github.com/l-ross/ficsit-toolkit/save"
)

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

func (p *production) PowerConsumption() float32 {
	return p.powerConsumption
}

// CurrentRecipe returns the currently assigned recipe for this production machine.
// Returns nil if the machine has no recipe assigned.
func (p *production) CurrentRecipe() *Recipe.FGRecipe {
	return p.currentRecipe
}

func (p *production) InputInventory() []InventoryStack {
	return p.inputInventory
}

func (p *production) OutputInventory() []InventoryStack {
	return p.outputInventory
}

func (p *production) IsProducing() bool {
	return p.isProducing
}

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
			// TODO
		}

		if err != nil {
			return nil, fmt.Errorf("failed to handle prop %q: %w", prop.Name, err)
		}
	}

	err := p.setInputs(b.entity, s)
	if err != nil {
		return nil, fmt.Errorf("failed to set inputs: %w", err)
	}

	//outputNodes := f.conveyorGraph.From(b.node.ID())
	//for outputNodes.Next() {
	//	n := outputNodes.Node()
	//	fmt.Println(n.ID())
	//	// Wind forward until we get something that is not a Conveyor
	//}
	//
	//inputNodes := f.conveyorGraph.To(b.node.ID())
	//for inputNodes.Next() {
	//	n := inputNodes.Node()
	//	fmt.Println(n.ID())
	//	// Wind back until we get something that is not a Conveyor
	//}

	return p, nil
}

func (p *production) setInputs(e *save.Entity, s *save.Save) error {
	inputs := getObjectsThatMatch(e.References, inputRegexp)
	if len(inputs) == 0 {
		return nil
	}

	return nil
}

func (p *production) setInputInventory(prop *save.Property, s *save.Save) error {
	stacks, err := getInventoryStacks(prop, s)
	if err != nil {
		return err
	}

	p.inputInventory = stacks

	return nil
}

func (p *production) setOutputInventory(prop *save.Property, s *save.Save) error {
	stacks, err := getInventoryStacks(prop, s)
	if err != nil {
		return err
	}

	p.outputInventory = stacks

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

		className := strings.Split(iiStruct.ItemName, ".")[1]

		i, err := ItemDescriptor.GetByClassName(className)
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

func (p *production) setPowerInfo(prop *save.Property, s *save.Save) error {
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

func (p *production) setRecipe(prop *save.Property) error {
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

func (p *production) setIsProducing(prop *save.Property) error {
	b, err := prop.GetBoolValue()
	if err != nil {
		return err
	}

	p.isProducing = b

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
