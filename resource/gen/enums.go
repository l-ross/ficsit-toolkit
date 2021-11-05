package main

import (
	"fmt"
	"regexp"
)

var enums = map[string]enum{
	"DS": {
		t:      "resource.DroneDockingState",
		values: droneDockingState,
	},
	"ES": {
		t:      "resource.EquipmentSlot",
		values: equipmentSlot,
	},
	"EST": {
		t:      "resource.SchematicType",
		values: schematicType,
	},
	"PPT": {
		t:      "resource.PowerPoleType",
		values: powerPoleType,
	},
	"RF": {
		t:      "resource.Form",
		values: resourceForm,
	},
	"SS": {
		t:      "resource.StackSize",
		values: stackSize,
	},
}

var enumRegexp = regexp.MustCompile(`^([A-Z]+)_[A-Za-z_]+$`)

func isEnum(s string) bool {
	if match := enumRegexp.FindStringSubmatch(s); match != nil {
		prefix := match[1]

		if _, exist := enums[prefix]; exist {
			return true
		}
	}

	return false
}

func getEnum(s string) (string, string, error) {
	if match := enumRegexp.FindStringSubmatch(s); match != nil {
		prefix := match[1]

		if e, exist := enums[prefix]; exist {
			if value, exist := e.values[s]; exist {
				return value, e.t, nil
			}
		}
	}

	return "", "", fmt.Errorf("failed to find enum for %s", s)
}

type enum struct {
	t      string
	values map[string]string
}

var schematicType = map[string]string{
	"EST_Custom":       "resource.Custom",
	"EST_Cheat":        "resource.Cheat",
	"EST_Tutorial":     "resource.Tutorial",
	"EST_Milestone":    "resource.Milestone",
	"EST_Alternate":    "resource.Alternate",
	"EST_Story":        "resource.Story",
	"EST_MAM":          "resource.MAM",
	"EST_ResourceSink": "resource.ResourceSink",
	"EST_HardDrive":    "resource.HardDrive",
	"EST_Prototype":    "resource.Prototype",
}

var resourceForm = map[string]string{
	"RF_INVALID": "resource.Invalid",
	"RF_SOLID":   "resource.Solid",
	"RF_LIQUID":  "resource.Liquid",
	"RF_GAS":     "resource.Gas",
	"RF_HEAT":    "resource.Heat",
}

var stackSize = map[string]string{
	"SS_ONE":    "resource.One",
	"SS_SMALL":  "resource.Small",
	"SS_MEDIUM": "resource.Medium",
	"SS_BIG":    "resource.Big",
	"SS_HUGE":   "resource.Huge",
	"SS_FLUID":  "resource.Fluid",
}

var equipmentSlot = map[string]string{
	"ES_NONE": "resource.None",
	"ES_ARMS": "resource.Arms",
	"ES_BACK": "resource.Back",
}

var droneDockingState = map[string]string{
	"DS_UNDOCKED": "resource.Undocked",
	"DS_DOCKING":  "resource.Docking",
	"DS_DOCKED":   "resource.Docked",
	"DS_TAKEOFF":  "resource.Takeoff",
}

var powerPoleType = map[string]string{
	"PPT_POLE":        "resource.Pole",
	"PPT_WALL":        "resource.Wall",
	"PPT_WALL_DOUBLE": "resource.WallDouble",
}
