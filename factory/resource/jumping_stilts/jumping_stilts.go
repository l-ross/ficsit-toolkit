// Code generated by ../../gen/docs_json. DO NOT EDIT.

package JumpingStilts

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGJumpingStilts struct {
	ClassName              string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MCostToUse             string
	MEquipmentSlot         resource.EquipmentSlot
	MHasPersistentOwner    bool
	MJumpSpeedFactor       float64
	MSprintSpeedFactor     float64
	MUseDefaultPrimaryFire bool
}

var (
	JumpingStilts = FGJumpingStilts{
		ClassName:              "Equip_JumpingStilts_C",
		MArmAnimation:          `AE_None`,
		MAttachSocket:          `jumpingStilt_lSocket`,
		MBackAnimation:         `BE_None`,
		MCostToUse:             ``,
		MEquipmentSlot:         resource.Back,
		MHasPersistentOwner:    false,
		MJumpSpeedFactor:       1.500000,
		MSprintSpeedFactor:     1.500000,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (*FGJumpingStilts, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGJumpingStilts with class name %s", className)
}

var classNameToVar = map[string]*FGJumpingStilts{
	"Equip_JumpingStilts_C": &JumpingStilts,
}
