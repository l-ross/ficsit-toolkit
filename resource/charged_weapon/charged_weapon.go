// Code generated by ../../gen/docs_json. DO NOT EDIT.

package ChargedWeapon

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGChargedWeapon struct {
	Name                             string
	ClassName                        string
	MAimTime                         float64
	MAllowedAmmoClasses              string
	MAmmoCycleMode                   string
	MAmmoSwitchUsedRadialMenu        bool
	MArmAnimation                    string
	MAttachSocket                    string
	MAutoReloadDelay                 float64
	MAutoReloadTimerHandle           string
	MAutomaticallyReload             bool
	MBackAnimation                   string
	MBaseAimTime                     float64
	MBaseFiringDispersion            float64
	MBaseRestingDispersion           float64
	MBlockSprintWhenFiring           bool
	MCostToUse                       string
	MCurrentAmmoCount                int
	MCurrentDispersion               float64
	MDelayBetweenSecondaryTriggers   float64
	MDispensedProjectiles            string
	MDispersionChangePerShot         float64
	MEquipmentSlot                   resource.EquipmentSlot
	MFireRate                        float64
	MFiringBlocksDispersionReduction bool
	MFiringDispersion                float64
	MHasPersistentOwner              bool
	MIsPendingExecuteFire            bool
	MMaxChargeTime                   float64
	MMaxThrowForce                   int
	MMinThrowForce                   int
	MOldState                        string
	MOnAmmoCyclingPressed            string
	MOnAmmoCyclingReleased           string
	MOnWeaponStateChanged            string
	MRadialMenuShowUpTime            float64
	MReloadTime                      float64
	MRestingDispersion               float64
	MShowCycleAmmoRadialMenuTimer    string
	MUseDefaultPrimaryFire           bool
	MWeaponDamageMultiplier          float64
	MWeaponState                     string
}

var (
	NobeliskDetonator = FGChargedWeapon{
		Name:                             "NobeliskDetonator",
		ClassName:                        "Equip_NobeliskDetonator_C",
		MAimTime:                         1.000000,
		MAllowedAmmoClasses:              `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/NobeliskExplosive/Desc_NobeliskExplosive.Desc_NobeliskExplosive_C"')`,
		MAmmoCycleMode:                   `IE_Released`,
		MAmmoSwitchUsedRadialMenu:        false,
		MArmAnimation:                    `AE_Nobelisk`,
		MAttachSocket:                    `hand_lSocket`,
		MAutoReloadDelay:                 1.000000,
		MAutoReloadTimerHandle:           `()`,
		MAutomaticallyReload:             false,
		MBackAnimation:                   `BE_None`,
		MBaseAimTime:                     0.000000,
		MBaseFiringDispersion:            0.000000,
		MBaseRestingDispersion:           0.000000,
		MBlockSprintWhenFiring:           true,
		MCostToUse:                       ``,
		MCurrentAmmoCount:                0,
		MCurrentDispersion:               0.000000,
		MDelayBetweenSecondaryTriggers:   0.150000,
		MDispensedProjectiles:            ``,
		MDispersionChangePerShot:         0.000000,
		MEquipmentSlot:                   resource.Arms,
		MFireRate:                        0.500000,
		MFiringBlocksDispersionReduction: false,
		MFiringDispersion:                3.000000,
		MHasPersistentOwner:              false,
		MIsPendingExecuteFire:            false,
		MMaxChargeTime:                   1.000000,
		MMaxThrowForce:                   3000,
		MMinThrowForce:                   200,
		MOldState:                        `(EWS_Unequipped)`,
		MOnAmmoCyclingPressed:            `()`,
		MOnAmmoCyclingReleased:           `()`,
		MOnWeaponStateChanged:            `()`,
		MRadialMenuShowUpTime:            0.200000,
		MReloadTime:                      1.500000,
		MRestingDispersion:               0.200000,
		MShowCycleAmmoRadialMenuTimer:    `()`,
		MUseDefaultPrimaryFire:           false,
		MWeaponDamageMultiplier:          1.000000,
		MWeaponState:                     `EWS_Unequipped`,
	}

	SnowballWeaponMittens = FGChargedWeapon{
		Name:                             "SnowballWeaponMittens",
		ClassName:                        "Equip_SnowballWeaponMittens_C",
		MAimTime:                         1.000000,
		MAllowedAmmoClasses:              `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SnowballProjectile/Desc_SnowballProjectile.Desc_SnowballProjectile_C"')`,
		MAmmoCycleMode:                   `IE_Released`,
		MAmmoSwitchUsedRadialMenu:        false,
		MArmAnimation:                    `AE_Generic1Hand`,
		MAttachSocket:                    `hand_lSocket`,
		MAutoReloadDelay:                 1.000000,
		MAutoReloadTimerHandle:           `()`,
		MAutomaticallyReload:             false,
		MBackAnimation:                   `BE_None`,
		MBaseAimTime:                     0.000000,
		MBaseFiringDispersion:            0.000000,
		MBaseRestingDispersion:           0.000000,
		MBlockSprintWhenFiring:           true,
		MCostToUse:                       ``,
		MCurrentAmmoCount:                0,
		MCurrentDispersion:               0.000000,
		MDelayBetweenSecondaryTriggers:   0.000000,
		MDispensedProjectiles:            ``,
		MDispersionChangePerShot:         0.000000,
		MEquipmentSlot:                   resource.Arms,
		MFireRate:                        0.500000,
		MFiringBlocksDispersionReduction: false,
		MFiringDispersion:                3.000000,
		MHasPersistentOwner:              false,
		MIsPendingExecuteFire:            false,
		MMaxChargeTime:                   0.300000,
		MMaxThrowForce:                   3500,
		MMinThrowForce:                   800,
		MOldState:                        `(EWS_Unequipped)`,
		MOnAmmoCyclingPressed:            `()`,
		MOnAmmoCyclingReleased:           `()`,
		MOnWeaponStateChanged:            `()`,
		MRadialMenuShowUpTime:            0.200000,
		MReloadTime:                      0.200000,
		MRestingDispersion:               0.200000,
		MShowCycleAmmoRadialMenuTimer:    `()`,
		MUseDefaultPrimaryFire:           false,
		MWeaponDamageMultiplier:          1.000000,
		MWeaponState:                     `EWS_Unequipped`,
	}
)

func GetByClassName(className string) (FGChargedWeapon, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGChargedWeapon{}, fmt.Errorf("failed to find FGChargedWeapon with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGChargedWeapon{
	"Equip_NobeliskDetonator_C":     NobeliskDetonator,
	"Equip_SnowballWeaponMittens_C": SnowballWeaponMittens,
}