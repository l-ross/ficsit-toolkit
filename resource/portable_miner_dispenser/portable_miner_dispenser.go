package PortableMinerDispenser

import (
	"fmt"
)

type FGPortableMinerDispenser struct {
	ClassName              string
	MAllowedResourceForms  string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MCostToUse             string
	MEquipmentSlot         string
	MHasPersistentOwner    bool
	MPlaceDistanceMax      float64
	MUseDefaultPrimaryFire bool
}

var (
	PortableMinerDispenser = FGPortableMinerDispenser{
		ClassName:              "Equip_PortableMinerDispenser_C",
		MAllowedResourceForms:  `(RF_SOLID)`,
		MArmAnimation:          `AE_Generic2Hand`,
		MAttachSocket:          `hand_rSocket`,
		MBackAnimation:         `BE_None`,
		MCostToUse:             ``,
		MEquipmentSlot:         `ES_ARMS`,
		MHasPersistentOwner:    false,
		MPlaceDistanceMax:      1000.000000,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (*FGPortableMinerDispenser, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGPortableMinerDispenser with name %s", className)
}

var classNameToVar = map[string]*FGPortableMinerDispenser{
	"Equip_PortableMinerDispenser_C": &PortableMinerDispenser,
}