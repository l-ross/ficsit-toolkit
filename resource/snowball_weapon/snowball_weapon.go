package SnowballWeapon

import (
	"fmt"
)

type FGSnowballWeapon struct {
	ClassName               string
	MArmAnimation           string
	MAttachSocket           string
	MBackAnimation          string
	MBlockSprintWhenFiring  bool
	MCostToUse              string
	MCurrentAmmo            int
	MDelayBetweenExplosions float64
	MDispensedExplosives    string
	MEquipmentSlot          string
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
	SnowballWeaponMittens = FGSnowballWeapon{
		ClassName:               "Equip_SnowballWeaponMittens_C",
		MArmAnimation:           `AE_Generic1Hand`,
		MAttachSocket:           `hand_lSocket`,
		MBackAnimation:          `BE_None`,
		MBlockSprintWhenFiring:  true,
		MCostToUse:              ``,
		MCurrentAmmo:            0,
		MDelayBetweenExplosions: 0.000000,
		MDispensedExplosives:    ``,
		MEquipmentSlot:          `ES_ARMS`,
		MExplosiveData:          `(ProjectileClass=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/Holiday/SnowballWeapon/BP_SnowballProjectile.BP_SnowballProjectile_C"',ProjectileLifeSpan=-1.000000,ProjectileStickSpan=-1.000000,ExplosionDamage=50,ExplosionRadius=750.000000,ImpactDamage=1,ExplodeAtEndOfLife=True,DamageType=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/Holiday/SnowballWeapon/DamageType_SnowballImpact.DamageType_SnowballImpact_C"',DamageTypeExplode=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/Holiday/SnowballWeapon/DamageType_SnowballExplosive.DamageType_SnowballExplosive_C"')`,
		MFireRate:               0.500000,
		MHasPersistentOwner:     false,
		MIsPendingExecuteFire:   false,
		MMagSize:                1,
		MMaxChargeTime:          0.300000,
		MMaxThrowForce:          3500,
		MMinThrowForce:          800,
		MReloadTime:             0.200000,
		MUseDefaultPrimaryFire:  false,
	}
)

func GetByClassName(className string) (*FGSnowballWeapon, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGSnowballWeapon with name %s", className)
}

var classNameToVar = map[string]*FGSnowballWeapon{
	"Equip_SnowballWeaponMittens_C": &SnowballWeaponMittens,
}
