package factory

import (
	"fmt"
	"strings"

	"github.com/l-ross/ficsit-toolkit/save"

	ItemDescriptorBiomass "github.com/l-ross/ficsit-toolkit/resource/item_descriptor_biomass"

	ItemDescriptor "github.com/l-ross/ficsit-toolkit/resource/item_descriptor"
	ResourceDescriptor "github.com/l-ross/ficsit-toolkit/resource/resource_descriptor"
	"github.com/l-ross/ficsit-toolkit/save/property"
)

type InventoryStackType string

const (
	ItemStackType     InventoryStackType = "ItemStack"
	ResourceStackType InventoryStackType = "ResourceStack"
	BiomassStackType  InventoryStackType = "BiomassStack"
)

type InventoryStack struct {
	Type InventoryStackType

	Item     *ItemDescriptor.FGItemDescriptor
	Resource *ResourceDescriptor.FGResourceDescriptor
	Biomass  *ItemDescriptorBiomass.FGItemDescriptorBiomass

	Quantity int
}

func getPropFromArray(name string, props []*property.Property) (*property.Property, error) {
	for _, p := range props {
		if p.Name == name {
			return p, nil
		}
	}

	return nil, fmt.Errorf("failed to find property with name %q", name)
}

func getValuesFromStructProperty(name string, s []*property.StructPropertyValue) []*property.StructPropertyValue {
	values := make([]*property.StructPropertyValue, 0)

	for _, x := range s {
		if string(x.Type) == name {
			values = append(values, x)
		}
	}

	return values
}

func getInventoryStack(itemName string, quantity int) (*InventoryStack, error) {
	split := strings.Split(itemName, ".")
	if len(split) != 2 {
		return nil, fmt.Errorf("failed to split item name %q to get class name", itemName)
	}

	className := split[1]

	is := &InventoryStack{
		Quantity: quantity,
	}

	if d, ok := ItemDescriptor.ClassNameToDescriptor[className]; ok {
		is.Item = &d
		is.Type = ItemStackType
	} else if d, ok := ResourceDescriptor.ClassNameToDescriptor[className]; ok {
		is.Resource = &d
		is.Type = ResourceStackType
	} else if d, ok := ItemDescriptorBiomass.ClassNameToDescriptor[className]; ok {
		is.Biomass = &d
		is.Type = BiomassStackType
	} else {
		return nil, fmt.Errorf("failed to find descriptor %q", className)
	}

	return is, nil
}

func getInventoryStacks(prop *property.Property, s *save.Save) ([]*InventoryStack, error) {
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

	stacks := make([]*InventoryStack, 0)

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

		is, err := getInventoryStack(iiStruct.ItemName, int(iiStruct.NumItems))
		if err != nil {
			return nil, err
		}

		stacks = append(stacks, is)
	}

	return stacks, nil
}
