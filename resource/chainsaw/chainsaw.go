package Chainsaw

import (
	"fmt"
)

type FGChainsaw struct {
	ClassName                   string
	MArmAnimation               string
	MAttachSocket               string
	MBackAnimation              string
	MCollateralPickupRadius     float64
	MCostToUse                  string
	MCurrentAudioState          string
	MCurrentState               string
	MEnergyConsumption          float64
	MEnergyStored               float64
	MEquipmentSlot              string
	MExcludeChainsawableFoliage bool
	MHasPersistentOwner         bool
	MInterpSawProgress          float64
	MMontageLength              float64
	MPlayingSound               bool
	MPreviousAudioState         string
	MSawDownTreeTime            float64
	MUseDefaultPrimaryFire      bool
	MWasSawing                  bool
}

var (
	Chainsaw = FGChainsaw{
		ClassName:                   "Equip_Chainsaw_C",
		MArmAnimation:               `AE_ChainSaw`,
		MAttachSocket:               `hand_lSocket`,
		MBackAnimation:              `BE_None`,
		MCollateralPickupRadius:     500.000000,
		MCostToUse:                  ``,
		MCurrentAudioState:          `NewEnumerator3`,
		MCurrentState:               `NewEnumerator3`,
		MEnergyConsumption:          75.000000,
		MEnergyStored:               0.000000,
		MEquipmentSlot:              `ES_ARMS`,
		MExcludeChainsawableFoliage: false,
		MHasPersistentOwner:         false,
		MInterpSawProgress:          0.000000,
		MMontageLength:              0.000000,
		MPlayingSound:               false,
		MPreviousAudioState:         `NewEnumerator3`,
		MSawDownTreeTime:            2.000000,
		MUseDefaultPrimaryFire:      false,
		MWasSawing:                  false,
	}
)

func GetByClassName(className string) (*FGChainsaw, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGChainsaw with name %s", className)
}

var classNameToVar = map[string]*FGChainsaw{
	"Equip_Chainsaw_C": &Chainsaw,
}