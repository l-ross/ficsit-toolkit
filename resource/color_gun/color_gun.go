package ColorGun

import (
	"fmt"
)

type FGColorGun struct {
	ClassName                         string
	MArmAnimation                     string
	MAttachSocket                     string
	MBackAnimation                    string
	MBlockSprintWhenFiring            bool
	MColorSlot                        int
	MCostToUse                        string
	MCurrentAmmo                      int
	MEquipmentSlot                    string
	MFireRate                         float64
	MHasPersistentOwner               bool
	MInstantHitDamage                 float64
	MMagSize                          int
	MNoTargetCrosshairColor           string
	MNonColorableTargetCrosshairColor string
	MPrimaryColor                     string
	MRedundantTargetCrosshairColor    string
	MReloadTime                       float64
	MSecondaryColor                   string
	MUseDefaultPrimaryFire            bool
	MWeaponTraceLength                float64
}

var (
	ColorGun = FGColorGun{
		ClassName:                         "Equip_ColorGun_C",
		MArmAnimation:                     `AE_ColorGun`,
		MAttachSocket:                     `hand_rSocket`,
		MBackAnimation:                    `BE_None`,
		MBlockSprintWhenFiring:            false,
		MColorSlot:                        1,
		MCostToUse:                        ``,
		MCurrentAmmo:                      0,
		MEquipmentSlot:                    `ES_ARMS`,
		MFireRate:                         0.500000,
		MHasPersistentOwner:               false,
		MInstantHitDamage:                 10.000000,
		MMagSize:                          100,
		MNoTargetCrosshairColor:           `(R=1.000000,G=1.000000,B=1.000000,A=0.600000)`,
		MNonColorableTargetCrosshairColor: `(R=1.000000,G=1.000000,B=1.000000,A=0.600000)`,
		MPrimaryColor:                     `(R=1.000000,G=0.000000,B=0.000000,A=1.000000)`,
		MRedundantTargetCrosshairColor:    `(R=1.000000,G=1.000000,B=1.000000,A=1.000000)`,
		MReloadTime:                       0.100000,
		MSecondaryColor:                   `(R=0.000000,G=1.000000,B=0.000000,A=1.000000)`,
		MUseDefaultPrimaryFire:            false,
		MWeaponTraceLength:                10000.000000,
	}
)

func GetByClassName(className string) (*FGColorGun, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGColorGun with name %s", className)
}

var classNameToVar = map[string]*FGColorGun{
	"Equip_ColorGun_C": &ColorGun,
}