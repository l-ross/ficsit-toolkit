// Code generated by ../../gen/docs_json. DO NOT EDIT.

package WeaponProjectileFire

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGWeaponProjectileFire struct {
	Name                   string
	ClassName              string
	MArmAnimation          string
	MAttachSocket          string
	MBackAnimation         string
	MBlockSprintWhenFiring bool
	MCostToUse             string
	MCurrentAmmo           int
	MEquipmentSlot         resource.EquipmentSlot
	MFireRate              float64
	MHasPersistentOwner    bool
	MMagSize               int
	MMuteDryFire           bool
	MProjectileData        string
	MRandomReloadAnim      int
	MRandomStingerAnim     int
	MReloadTime            float64
	MUseDefaultPrimaryFire bool
}

var (
	RebarGunProjectile = FGWeaponProjectileFire{
		Name:                   "RebarGunProjectile",
		ClassName:              "Equip_RebarGun_Projectile_C",
		MArmAnimation:          `AE_RebarGun`,
		MAttachSocket:          `hand_rSocket`,
		MBackAnimation:         `BE_None`,
		MBlockSprintWhenFiring: true,
		MCostToUse:             ``,
		MCurrentAmmo:           0,
		MEquipmentSlot:         resource.Arms,
		MFireRate:              1.000000,
		MHasPersistentOwner:    false,
		MMagSize:               1,
		MMuteDryFire:           false,
		MProjectileData:        `(ProjectileClass=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/RebarGun/BP_RebarProjectile.BP_RebarProjectile_C"',ProjectileLifeSpan=10.000000,ProjectileStickSpan=5.000000,ExplosionDamage=15,ImpactDamage=15,CanTriggerExplodeBySameClass=True,DamageType=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/RebarGun/DamageType_RebarGunProjectile.DamageType_RebarGunProjectile_C"',DamageTypeExplode=Class'"/Script/FactoryGame.FGDamageType"')`,
		MRandomReloadAnim:      0,
		MRandomStingerAnim:     0,
		MReloadTime:            1.800000,
		MUseDefaultPrimaryFire: false,
	}
)

func GetByClassName(className string) (FGWeaponProjectileFire, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGWeaponProjectileFire{}, fmt.Errorf("failed to find FGWeaponProjectileFire with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGWeaponProjectileFire{
	"Equip_RebarGun_Projectile_C": RebarGunProjectile,
}
