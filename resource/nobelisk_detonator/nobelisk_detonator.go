// Code generated by ../../gen/docs_json. DO NOT EDIT.

package NobeliskDetonator

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGNobeliskDetonator struct {
	Name                    string
	ClassName               string
	MArmAnimation           string
	MAttachSocket           string
	MBackAnimation          string
	MBlockSprintWhenFiring  bool
	MCostToUse              string
	MCurrentAmmo            int
	MDelayBetweenExplosions float64
	MDispensedExplosives    string
	MEquipmentSlot          resource.EquipmentSlot
	MExplosiveData          string
	MFireRate               float64
	MHasPersistentOwner     bool
	MIsPendingExecuteFire   bool
	MMagSize                int
	MMaxChargeTime          float64
	MMaxThrowForce          int
	MMinThrowForce          int
	MReloadTime             float64
	MUseDefaultPrimaryFire  bool
}

var (
	NobeliskDetonator = FGNobeliskDetonator{
		Name:                    "NobeliskDetonator",
		ClassName:               "Equip_NobeliskDetonator_C",
		MArmAnimation:           `AE_Nobelisk`,
		MAttachSocket:           `hand_lSocket`,
		MBackAnimation:          `BE_None`,
		MBlockSprintWhenFiring:  true,
		MCostToUse:              ``,
		MCurrentAmmo:            0,
		MDelayBetweenExplosions: 0.150000,
		MDispensedExplosives:    ``,
		MEquipmentSlot:          resource.Arms,
		MExplosiveData:          `(ProjectileClass=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/NobeliskDetonator/BP_NobeliskExplosive.BP_NobeliskExplosive_C"',ProjectileLifeSpan=-1.000000,ProjectileStickSpan=-1.000000,ExplosionDamage=50,ExplosionRadius=750.000000,ImpactDamage=1,ExplodeAtEndOfLife=True,DamageType=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/NobeliskDetonator/DamageType_NobeliskExplosiveImpact.DamageType_NobeliskExplosiveImpact_C"',DamageTypeExplode=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/NobeliskDetonator/DamageType_NobeliskExplosive.DamageType_NobeliskExplosive_C"')`,
		MFireRate:               0.500000,
		MHasPersistentOwner:     false,
		MIsPendingExecuteFire:   false,
		MMagSize:                1,
		MMaxChargeTime:          1.000000,
		MMaxThrowForce:          3000,
		MMinThrowForce:          200,
		MReloadTime:             1.000000,
		MUseDefaultPrimaryFire:  false,
	}
)

func GetByClassName(className string) (FGNobeliskDetonator, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGNobeliskDetonator{}, fmt.Errorf("failed to find FGNobeliskDetonator with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGNobeliskDetonator{
	"Equip_NobeliskDetonator_C": NobeliskDetonator,
}
