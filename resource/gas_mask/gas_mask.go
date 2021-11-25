// Code generated by ../../gen/docs_json. DO NOT EDIT.

package GasMask

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGGasMask struct {
	Name                   string
	ClassName              string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MCostToUse             string
	MCountdown             float64
	MDamageNegated         float64
	MDisableEffectTimer    float64
	MEquipmentSlot         resource.EquipmentSlot
	MFilterDuration        float64
	MHasNegatedDamage      bool
	MHasPersistentOwner    bool
	MIsWorking             bool
	MPostProcessEnabled    bool
	MUseDefaultPrimaryFire bool
}

var (
	GasMask = FGGasMask{
		Name:                   "GasMask",
		ClassName:              "Equip_GasMask_C",
		MArmAnimation:          `AE_None`,
		MAttachSocket:          `helmetSocket`,
		MBackAnimation:         `BE_None`,
		MCostToUse:             `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Filter/Desc_Filter.Desc_Filter_C"',Amount=1))`,
		MCountdown:             0.000000,
		MDamageNegated:         0.000000,
		MDisableEffectTimer:    0.000000,
		MEquipmentSlot:         resource.Back,
		MFilterDuration:        240.000000,
		MHasNegatedDamage:      false,
		MHasPersistentOwner:    false,
		MIsWorking:             false,
		MPostProcessEnabled:    false,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (FGGasMask, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGGasMask{}, fmt.Errorf("failed to find FGGasMask with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGGasMask{
	"Equip_GasMask_C": GasMask,
}
