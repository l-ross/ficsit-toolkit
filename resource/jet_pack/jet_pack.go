// Code generated by ../../gen/docs_json. DO NOT EDIT.

package JetPack

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGJetPack struct {
	ClassName                string
	MArmAnimation            string
	MAttachSocket            string
	MBackAnimation           string
	MCostToUse               string
	MCurrentFuel             float64
	MDefaultAirControl       float64
	MEquipmentSlot           resource.EquipmentSlot
	MFuelConsumeRate         float64
	MFuelRegenRate           float64
	MFuelWorth               float64
	MHasPersistentOwner      bool
	MIsThrusting             bool
	MJumpBeforeThrustTime    float64
	MJumpTimeStamp           float64
	MMaxFuel                 float64
	MPaidForFuel             float64
	MRTPCInterval            float64
	MThrustAirControl        float64
	MThrustCooldown          float64
	MThrustPower             float64
	MUseDefaultPrimaryFire   bool
	MVelocityZExtreme        float64
	MVelocityZExtremesDamper float64
}

var (
	JetPack = FGJetPack{
		ClassName:                "Equip_JetPack_C",
		MArmAnimation:            `AE_None`,
		MAttachSocket:            `None`,
		MBackAnimation:           `BE_Jetpack`,
		MCostToUse:               `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Fuel/Desc_Fuel.Desc_Fuel_C"',Amount=1))`,
		MCurrentFuel:             1.000000,
		MDefaultAirControl:       0.000000,
		MEquipmentSlot:           resource.Back,
		MFuelConsumeRate:         0.200000,
		MFuelRegenRate:           0.500000,
		MFuelWorth:               0.500000,
		MHasPersistentOwner:      false,
		MIsThrusting:             false,
		MJumpBeforeThrustTime:    0.300000,
		MJumpTimeStamp:           0.000000,
		MMaxFuel:                 1.000000,
		MPaidForFuel:             0.000000,
		MRTPCInterval:            0.000000,
		MThrustAirControl:        0.300000,
		MThrustCooldown:          0.400000,
		MThrustPower:             4000.000000,
		MUseDefaultPrimaryFire:   false,
		MVelocityZExtreme:        350.000000,
		MVelocityZExtremesDamper: 0.900000,
	}
)

func GetByClassName(className string) (*FGJetPack, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGJetPack with name %s", className)
}

var classNameToVar = map[string]*FGJetPack{
	"Equip_JetPack_C": &JetPack,
}
