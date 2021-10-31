package GasMask

import (
	"fmt"
)

type FGGasMask struct {
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
	MIsWorking             bool
	MPostProcessEnabled    bool
	MUseDefaultPrimaryFire bool
}

var (
	GasMask = FGGasMask{
		ClassName:              "Equip_GasMask_C",
		MArmAnimation:          `AE_None`,
		MAttachSocket:          `helmetSocket`,
		MBackAnimation:         `BE_None`,
		MCostToUse:             `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Filter/Desc_Filter.Desc_Filter_C"',Amount=1))`,
		MCountdown:             0.000000,
		MDamageNegated:         0.000000,
		MDisableEffectTimer:    0.000000,
		MEquipmentSlot:         `ES_BACK`,
		MFilterDuration:        240.000000,
		MHasNegatedDamage:      false,
		MHasPersistentOwner:    false,
		MIsWorking:             false,
		MPostProcessEnabled:    false,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (*FGGasMask, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGGasMask with name %s", className)
}

var classNameToVar = map[string]*FGGasMask{
	"Equip_GasMask_C": &GasMask,
}
