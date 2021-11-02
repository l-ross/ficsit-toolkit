// Code generated by ../../gen/docs_json. DO NOT EDIT.

package Parachute

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/factory/resource"
)

type FGParachute struct {
	Name                   string
	ClassName              string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MCostToUse             string
	MEquipmentSlot         resource.EquipmentSlot
	MHasPersistentOwner    bool
	MIsDeployed            bool
	MTerminalVelocityZ     float64
	MUseDefaultPrimaryFire bool
}

var (
	Parachute = FGParachute{
		Name:                   "Parachute",
		ClassName:              "Equip_Parachute_C",
		MArmAnimation:          `AE_None`,
		MAttachSocket:          `root`,
		MBackAnimation:         `BE_None`,
		MCostToUse:             `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Beacon/Desc_Parachute.Desc_Parachute_C"',Amount=1))`,
		MEquipmentSlot:         resource.Back,
		MHasPersistentOwner:    false,
		MIsDeployed:            false,
		MTerminalVelocityZ:     300.000000,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (FGParachute, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return FGParachute{}, fmt.Errorf("failed to find FGParachute with class name %s", className)
}

var classNameToVar = map[string]FGParachute{
	"Equip_Parachute_C": Parachute,
}
