package factory

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	ItemDescriptor "github.com/l-ross/ficsit-toolkit/resource/item_descriptor"
	"github.com/l-ross/ficsit-toolkit/save"
	"github.com/l-ross/ficsit-toolkit/save/property"
)

type InventoryStack struct {
	Item     ItemDescriptor.FGItemDescriptor
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

func getItemDescriptorByName(name string) (ItemDescriptor.FGItemDescriptor, error) {
	split := strings.Split(name, ".")
	if len(split) != 2 {
		return ItemDescriptor.FGItemDescriptor{}, fmt.Errorf("failed to split item name %q to get class name", name)
	}

	className := split[1]

	i, err := ItemDescriptor.GetByClassName(className)
	if err != nil {
		return ItemDescriptor.FGItemDescriptor{}, err
	}

	return i, nil
}

func getObjectsThatMatch(refs []*save.ObjectReference, re *regexp.Regexp) []string {
	matches := make([]string, 0)

	for _, o := range refs {
		if re.MatchString(o.PathName) {
			matches = append(matches, o.PathName)
		}
	}

	// We want to lower case the matches before comparing because
	// sometimes the cases differ.
	sort.Slice(matches, func(i, j int) bool {
		ii := strings.ToUpper(matches[i])
		jj := strings.ToUpper(matches[j])

		if ii < jj {
			return true
		}
		return false
	})

	return matches
}
