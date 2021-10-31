package SuitBase

import (
	"fmt"
)

type FGSuitBase struct {
	ClassName              string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MCostToUse             string
	MCountdown             float64
	MDamageNegated         float64
	MDisableEffectTimer    float64
	MEquipmentSlot         string
	MFilterDuration        float64
	MHasNegatedDamage      bool
	MHasPersistentOwner    bool
	MImmunity              float64
	MIsWorking             bool
	MSuit1PMeshMaterials   string
	MUseDefaultPrimaryFire bool
}

var (
	HazmatSuit = FGSuitBase{
		ClassName:              "Equip_HazmatSuit_C",
		MArmAnimation:          `AE_None`,
		MAttachSocket:          `None`,
		MBackAnimation:         `BE_None`,
		MCostToUse:             `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IodineInfusedFilter/Desc_HazmatFilter.Desc_HazmatFilter_C"',Amount=1))`,
		MCountdown:             0.000000,
		MDamageNegated:         0.000000,
		MDisableEffectTimer:    0.000000,
		MEquipmentSlot:         `ES_BACK`,
		MFilterDuration:        240.000000,
		MHasNegatedDamage:      false,
		MHasPersistentOwner:    false,
		MImmunity:              0.000000,
		MIsWorking:             false,
		MSuit1PMeshMaterials:   `((SlotName="Body_Details",Material=MaterialInstanceConstant'"/Game/FactoryGame/Character/Player/Material/MI_Haz_Body_Details.MI_Haz_Body_Details"'),(SlotName="Body_01",Material=MaterialInstanceConstant'"/Game/FactoryGame/Character/Player/Material/MI_Haz_Body_01.MI_Haz_Body_01"'),(SlotName="Body_02",Material=MaterialInstanceConstant'"/Game/FactoryGame/Character/Player/Material/MI_Haz_Body_02.MI_Haz_Body_02"'),(SlotName="Body_Hands",Material=MaterialInstanceConstant'"/Game/FactoryGame/Character/Player/Material/MI_Haz_Body_Hands.MI_Haz_Body_Hands"'),(SlotName="Body_Backpack",Material=MaterialInstanceConstant'"/Game/FactoryGame/Character/Player/Material/MI_Haz_Body_Backpack.MI_Haz_Body_Backpack"'))`,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (*FGSuitBase, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGSuitBase with name %s", className)
}

var classNameToVar = map[string]*FGSuitBase{
	"Equip_HazmatSuit_C": &HazmatSuit,
}