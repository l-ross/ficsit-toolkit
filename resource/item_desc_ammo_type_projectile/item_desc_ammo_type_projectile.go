// Code generated by ../../gen/docs_json. DO NOT EDIT.

package ItemDescAmmoTypeProjectile

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGItemDescAmmoTypeProjectile struct {
	Name                             string
	ClassName                        string
	AmmoFiredDelegate                string
	MAbbreviatedDisplayName          string
	MAmmoDamageFalloff               string
	MBuildMenuPriority               float64
	MCanBeDiscarded                  bool
	MDescription                     string
	MDispersionPerShotChangeModifier float64
	MDisplayName                     string
	MEnergyValue                     float64
	MFireRateModifier                float64
	MFiringTransform                 string
	MFluidColor                      string
	MForm                            resource.Form
	MGasColor                        string
	MHasBeenInitialized              bool
	MMagazineSize                    int
	MMaxAmmoEffectiveRange           float64
	MMaxWeaponDispersionModifier     float64
	MMenuPriority                    float64
	MMinWeaponDispersionModifier     float64
	MPersistentBigIcon               string
	MProjectileData                  string
	MRadioactiveDecay                float64
	MReloadTimeModifier              float64
	MRememberPickUp                  bool
	MResourceSinkPoints              int
	MSmallIcon                       string
	MStackSize                       resource.StackSize
	MSubCategories                   string
	MWeaponAimTimeModifier           float64
	MWeaponDamageMultiplier          float64
}

var (
	NobeliskExplosive = FGItemDescAmmoTypeProjectile{
		Name:                             "NobeliskExplosive",
		ClassName:                        "Desc_NobeliskExplosive_C",
		AmmoFiredDelegate:                `()`,
		MAbbreviatedDisplayName:          ``,
		MAmmoDamageFalloff:               `(EditorCurveData=(DefaultValue=340282346638528859811704183484516925440.000000,PreInfinityExtrap=RCCE_Constant,PostInfinityExtrap=RCCE_Constant),ExternalCurve=CurveFloat'"/Game/FactoryGame/Resource/Parts/CartridgeStandard/LinearDamageFalloff.LinearDamageFalloff"')`,
		MBuildMenuPriority:               0.000000,
		MCanBeDiscarded:                  true,
		MDescription:                     `Can be used with the Nobelisk Detonator to blow up cracked boulders, vegetation or other vulnerable targets.`,
		MDispersionPerShotChangeModifier: 0.000000,
		MDisplayName:                     `Nobelisk`,
		MEnergyValue:                     0.000000,
		MFireRateModifier:                1.000000,
		MFiringTransform:                 `(Rotation=(X=0.000000,Y=0.000000,Z=0.000000,W=1.000000),Translation=(X=0.000000,Y=0.000000,Z=0.000000),Scale3D=(X=1.000000,Y=1.000000,Z=1.000000))`,
		MFluidColor:                      `(B=0,G=0,R=0,A=0)`,
		MForm:                            resource.Solid,
		MGasColor:                        `(B=0,G=0,R=0,A=0)`,
		MHasBeenInitialized:              false,
		MMagazineSize:                    1,
		MMaxAmmoEffectiveRange:           5000.000000,
		MMaxWeaponDispersionModifier:     0.000000,
		MMenuPriority:                    0.000000,
		MMinWeaponDispersionModifier:     0.000000,
		MPersistentBigIcon:               `Texture2D /Game/FactoryGame/Resource/Parts/NobeliskExplosive/UI/IconDesc_Explosive_256.IconDesc_Explosive_256`,
		MProjectileData:                  `(ProjectileClass=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/NobeliskDetonator/BP_NobeliskExplosive.BP_NobeliskExplosive_C"',ProjectileLifeSpan=-1.000000,ProjectileStickSpan=-1.000000,ExplosionDamage=50,ExplosionRadius=750.000000,ImpactDamage=1,ExplodeAtEndOfLife=True,DamageType=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/NobeliskDetonator/DamageType_NobeliskExplosiveImpact.DamageType_NobeliskExplosiveImpact_C"',DamageTypeExplode=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/NobeliskDetonator/DamageType_NobeliskExplosive.DamageType_NobeliskExplosive_C"',DamageFalloffCurve=(EditorCurveData=(DefaultValue=340282346638528859811704183484516925440.000000,PreInfinityExtrap=RCCE_Constant,PostInfinityExtrap=RCCE_Constant)),WeaponDamageMultiplier=1.000000)`,
		MRadioactiveDecay:                0.000000,
		MReloadTimeModifier:              1.000000,
		MRememberPickUp:                  false,
		MResourceSinkPoints:              980,
		MSmallIcon:                       `Texture2D /Game/FactoryGame/Resource/Parts/NobeliskExplosive/UI/IconDesc_Explosive_64.IconDesc_Explosive_64`,
		MStackSize:                       resource.Small,
		MSubCategories:                   ``,
		MWeaponAimTimeModifier:           0.000000,
		MWeaponDamageMultiplier:          1.000000,
	}

	SnowballProjectile = FGItemDescAmmoTypeProjectile{
		Name:                    "SnowballProjectile",
		ClassName:               "Desc_SnowballProjectile_C",
		AmmoFiredDelegate:       `()`,
		MAbbreviatedDisplayName: ``,
		MAmmoDamageFalloff:      `(EditorCurveData=(DefaultValue=340282346638528859811704183484516925440.000000,PreInfinityExtrap=RCCE_Constant,PostInfinityExtrap=RCCE_Constant),ExternalCurve=CurveFloat'"/Game/FactoryGame/Resource/Parts/CartridgeStandard/LinearDamageFalloff.LinearDamageFalloff"')`,
		MBuildMenuPriority:      0.000000,
		MCanBeDiscarded:         true,
		MDescription: `Compressed dihydrogen monoxide crystals.

Alternative Nobelisk Ammo. Use G to swap!`,
		MDispersionPerShotChangeModifier: 0.000000,
		MDisplayName:                     `Snowball`,
		MEnergyValue:                     0.000000,
		MFireRateModifier:                1.000000,
		MFiringTransform:                 `(Rotation=(X=0.000000,Y=0.000000,Z=0.000000,W=1.000000),Translation=(X=0.000000,Y=0.000000,Z=0.000000),Scale3D=(X=1.000000,Y=1.000000,Z=1.000000))`,
		MFluidColor:                      `(B=0,G=0,R=0,A=0)`,
		MForm:                            resource.Solid,
		MGasColor:                        `(B=0,G=0,R=0,A=0)`,
		MHasBeenInitialized:              false,
		MMagazineSize:                    1,
		MMaxAmmoEffectiveRange:           5000.000000,
		MMaxWeaponDispersionModifier:     0.000000,
		MMenuPriority:                    0.000000,
		MMinWeaponDispersionModifier:     0.000000,
		MPersistentBigIcon:               `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_SnowballProjectile_256.IconDesc_SnowballProjectile_256`,
		MProjectileData:                  `(ProjectileClass=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/Holiday/SnowballWeapon/BP_SnowballProjectile.BP_SnowballProjectile_C"',ProjectileLifeSpan=-1.000000,ProjectileStickSpan=-1.000000,ExplosionDamage=50,ExplosionRadius=750.000000,ImpactDamage=1,ExplodeAtEndOfLife=True,DamageType=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/Holiday/SnowballWeapon/DamageType_SnowballImpact.DamageType_SnowballImpact_C"',DamageTypeExplode=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/Holiday/SnowballWeapon/DamageType_SnowballExplosive.DamageType_SnowballExplosive_C"',DamageFalloffCurve=(EditorCurveData=(DefaultValue=340282346638528859811704183484516925440.000000,PreInfinityExtrap=RCCE_Constant,PostInfinityExtrap=RCCE_Constant)),WeaponDamageMultiplier=1.000000)`,
		MRadioactiveDecay:                0.000000,
		MReloadTimeModifier:              1.000000,
		MRememberPickUp:                  false,
		MResourceSinkPoints:              1,
		MSmallIcon:                       `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_SnowballProjectile_64.IconDesc_SnowballProjectile_64`,
		MStackSize:                       resource.Huge,
		MSubCategories:                   ``,
		MWeaponAimTimeModifier:           0.000000,
		MWeaponDamageMultiplier:          1.000000,
	}

	SpikedRebar = FGItemDescAmmoTypeProjectile{
		Name:                             "SpikedRebar",
		ClassName:                        "Desc_SpikedRebar_C",
		AmmoFiredDelegate:                `()`,
		MAbbreviatedDisplayName:          ``,
		MAmmoDamageFalloff:               `(EditorCurveData=(DefaultValue=340282346638528859811704183484516925440.000000,PreInfinityExtrap=RCCE_Constant,PostInfinityExtrap=RCCE_Constant),ExternalCurve=CurveFloat'"/Game/FactoryGame/Resource/Parts/CartridgeStandard/LinearDamageFalloff.LinearDamageFalloff"')`,
		MBuildMenuPriority:               0.000000,
		MCanBeDiscarded:                  true,
		MDescription:                     `Ammo for the Rebar Gun.`,
		MDispersionPerShotChangeModifier: 0.000000,
		MDisplayName:                     `Spiked Rebar`,
		MEnergyValue:                     0.000000,
		MFireRateModifier:                1.000000,
		MFiringTransform:                 `(Rotation=(X=0.000000,Y=0.000000,Z=0.000000,W=1.000000),Translation=(X=0.000000,Y=0.000000,Z=0.000000),Scale3D=(X=1.000000,Y=1.000000,Z=1.000000))`,
		MFluidColor:                      `(B=0,G=0,R=0,A=0)`,
		MForm:                            resource.Solid,
		MGasColor:                        `(B=0,G=0,R=0,A=0)`,
		MHasBeenInitialized:              false,
		MMagazineSize:                    1,
		MMaxAmmoEffectiveRange:           5000.000000,
		MMaxWeaponDispersionModifier:     0.000000,
		MMenuPriority:                    0.000000,
		MMinWeaponDispersionModifier:     0.000000,
		MPersistentBigIcon:               `Texture2D /Game/FactoryGame/Resource/Parts/SpikedRebar/UI/IconDesc_SpikedRebar_256.IconDesc_SpikedRebar_256`,
		MProjectileData:                  `(ProjectileClass=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/RebarGun/BP_RebarProjectile.BP_RebarProjectile_C"',ProjectileLifeSpan=10.000000,ProjectileStickSpan=5.000000,ImpactDamage=15,CanTriggerExplodeBySameClass=True,DamageType=BlueprintGeneratedClass'"/Game/FactoryGame/Equipment/RebarGun/DamageType_RebarGunProjectile.DamageType_RebarGunProjectile_C"',DamageTypeExplode=Class'"/Script/FactoryGame.FGDamageType"',DamageFalloffCurve=(EditorCurveData=(DefaultValue=340282346638528859811704183484516925440.000000,PreInfinityExtrap=RCCE_Constant,PostInfinityExtrap=RCCE_Constant)),WeaponDamageMultiplier=1.000000)`,
		MRadioactiveDecay:                0.000000,
		MReloadTimeModifier:              1.000000,
		MRememberPickUp:                  false,
		MResourceSinkPoints:              8,
		MSmallIcon:                       `Texture2D /Game/FactoryGame/Resource/Parts/SpikedRebar/UI/IconDesc_SpikedRebar_64.IconDesc_SpikedRebar_64`,
		MStackSize:                       resource.Small,
		MSubCategories:                   ``,
		MWeaponAimTimeModifier:           0.000000,
		MWeaponDamageMultiplier:          1.000000,
	}
)

func GetByClassName(className string) (FGItemDescAmmoTypeProjectile, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGItemDescAmmoTypeProjectile{}, fmt.Errorf("failed to find FGItemDescAmmoTypeProjectile with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGItemDescAmmoTypeProjectile{
	"Desc_NobeliskExplosive_C":  NobeliskExplosive,
	"Desc_SnowballProjectile_C": SnowballProjectile,
	"Desc_SpikedRebar_C":        SpikedRebar,
}
