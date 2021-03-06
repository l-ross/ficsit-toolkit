// Code generated by ../../gen/docs_json. DO NOT EDIT.

package GolfCartDispenser

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGGolfCartDispenser struct {
	Name                   string
	ClassName              string
	CanDisplayDisqualifier bool
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MBuildDisqualifierText string
	MCostToUse             string
	MEquipmentSlot         resource.EquipmentSlot
	MHasPersistentOwner    bool
	MMaxSpawnDistance      float64
	MPlaceDistanceMax      float64
	MSpawningClearance     float64
	MUseDefaultPrimaryFire bool
}

var (
	GoldGolfCartDispenser = FGGolfCartDispenser{
		Name:                   "GoldGolfCartDispenser",
		ClassName:              "Equip_GoldGolfCartDispenser_C",
		CanDisplayDisqualifier: true,
		MArmAnimation:          `AE_Generic2Hand`,
		MAttachSocket:          `hand_rSocket`,
		MBackAnimation:         `BE_None`,
		MBuildDisqualifierText: `Vehicles cannot be built or deployed on top of existing vehicles.`,
		MCostToUse:             ``,
		MEquipmentSlot:         resource.Arms,
		MHasPersistentOwner:    false,
		MMaxSpawnDistance:      3000.000000,
		MPlaceDistanceMax:      1000.000000,
		MSpawningClearance:     1.050000,
		MUseDefaultPrimaryFire: false,
	}

	GolfCartDispenser = FGGolfCartDispenser{
		Name:                   "GolfCartDispenser",
		ClassName:              "Equip_GolfCartDispenser_C",
		CanDisplayDisqualifier: true,
		MArmAnimation:          `AE_Generic2Hand`,
		MAttachSocket:          `hand_rSocket`,
		MBackAnimation:         `BE_None`,
		MBuildDisqualifierText: `Vehicles cannot be built or deployed on top of existing vehicles.`,
		MCostToUse:             ``,
		MEquipmentSlot:         resource.Arms,
		MHasPersistentOwner:    false,
		MMaxSpawnDistance:      3000.000000,
		MPlaceDistanceMax:      1000.000000,
		MSpawningClearance:     1.050000,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (FGGolfCartDispenser, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGGolfCartDispenser{}, fmt.Errorf("failed to find FGGolfCartDispenser with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGGolfCartDispenser{
	"Equip_GoldGolfCartDispenser_C": GoldGolfCartDispenser,
	"Equip_GolfCartDispenser_C":     GolfCartDispenser,
}
