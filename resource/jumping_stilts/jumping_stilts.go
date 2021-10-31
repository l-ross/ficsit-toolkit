package JumpingStilts

import (
	"fmt"
)

type FGJumpingStilts struct {
	ClassName              string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MCostToUse             string
	MEquipmentSlot         string
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
		MEquipmentSlot:         `ES_BACK`,
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

	return nil, fmt.Errorf("failed to find FGJumpingStilts with name %s", className)
}

var classNameToVar = map[string]*FGJumpingStilts{
	"Equip_JumpingStilts_C": &JumpingStilts,
}
