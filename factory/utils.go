package factory

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	ItemDescriptor "github.com/l-ross/ficsit-toolkit/resource/item_descriptor"

	"github.com/l-ross/ficsit-toolkit/save"
)

type InventoryStack struct {
	Item     ItemDescriptor.FGItemDescriptor
	Quantity int
}

func getPropFromArray(name string, props []*save.Property) (*save.Property, error) {
	for _, p := range props {
		if p.Name == name {
			return p, nil
		}
	}

	return nil, fmt.Errorf("failed to find property with name %q", name)
}

func getValuesFromStructProperty(name string, s []*save.StructPropertyValue) []*save.StructPropertyValue {
	values := make([]*save.StructPropertyValue, 0)

	for _, x := range s {
		if string(x.Type) == name {
			values = append(values, x)
		}
	}

	return values
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
