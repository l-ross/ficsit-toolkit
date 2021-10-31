package Parachute

import (
	"fmt"
)

type FGParachute struct {
	ClassName              string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MCostToUse             string
	MEquipmentSlot         string
	MHasPersistentOwner    bool
	MIsDeployed            bool
	MTerminalVelocityZ     float64
	MUseDefaultPrimaryFire bool
}

var (
	Parachute = FGParachute{
		ClassName:              "Equip_Parachute_C",
		MArmAnimation:          `AE_None`,
		MAttachSocket:          `root`,
		MBackAnimation:         `BE_None`,
		MCostToUse:             `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Beacon/Desc_Parachute.Desc_Parachute_C"',Amount=1))`,
		MEquipmentSlot:         `ES_BACK`,
		MHasPersistentOwner:    false,
		MIsDeployed:            false,
		MTerminalVelocityZ:     300.000000,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (*FGParachute, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGParachute with name %s", className)
}

var classNameToVar = map[string]*FGParachute{
	"Equip_Parachute_C": &Parachute,
}