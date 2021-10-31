package EquipmentZipline

import (
	"fmt"
)

type FGEquipmentZipline struct {
	ClassName                  string
	MArmAnimation              string
	MAttachSocket              string
	MBackAnimation             string
	MCostToUse                 string
	MEquipmentSlot             string
	MHasPersistentOwner        bool
	MMaxZiplineAngle           float64
	MShouldPlayDeactivateSound bool
	MTraceDistance             float64
	MTraceRadius               float64
	MTraceStartOffset          float64
	MUseDefaultPrimaryFire     bool
	MVisualizeTraceDistance    bool
	MZiplineJumpLaunchVelocity float64
	MZiplineReattachCooldown   float64
}

var (
	Zipline = FGEquipmentZipline{
		ClassName:                  "Equip_Zipline_C",
		MArmAnimation:              `AE_Generic1Hand`,
		MAttachSocket:              `hand_rSocket`,
		MBackAnimation:             `BE_None`,
		MCostToUse:                 ``,
		MEquipmentSlot:             `ES_ARMS`,
		MHasPersistentOwner:        false,
		MMaxZiplineAngle:           0.950000,
		MShouldPlayDeactivateSound: false,
		MTraceDistance:             180.000000,
		MTraceRadius:               100.000000,
		MTraceStartOffset:          250.000000,
		MUseDefaultPrimaryFire:     true,
		MVisualizeTraceDistance:    false,
		MZiplineJumpLaunchVelocity: 700.000000,
		MZiplineReattachCooldown:   0.300000,
	}
)

func GetByClassName(className string) (*FGEquipmentZipline, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGEquipmentZipline with name %s", className)
}

var classNameToVar = map[string]*FGEquipmentZipline{
	"Equip_Zipline_C": &Zipline,
}
