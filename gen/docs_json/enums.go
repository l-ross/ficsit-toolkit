package main

var stackSize = map[string]string{
	"SS_ONE":    "resource.One",
	"SS_SMALL":  "resource.Small",
	"SS_MEDIUM": "resource.Medium",
	"SS_BIG":    "resource.Big",
	"SS_HUGE":   "resource.Huge",
	"SS_FLUID":  "resource.Fluid",
}

var resourceForm = map[string]string{
	"RF_INVALID": "resource.Invalid",
	"RF_SOLID":   "resource.Solid",
	"RF_LIQUID":  "resource.Liquid",
	"RF_GAS":     "resource.Gas",
	"RF_HEAT":    "resource.Heat",
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
