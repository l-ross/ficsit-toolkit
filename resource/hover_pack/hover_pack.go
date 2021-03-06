// Code generated by ../../gen/docs_json. DO NOT EDIT.

package HoverPack

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGHoverPack struct {
	Name                                     string
	ClassName                                string
	ConnectionLocationUpdatedDelegate        string
	ConnectionStatusUpdatedDelegate          string
	HoverModeChangedDelegate                 string
	MArmAnimation                            string
	MAttachSocket                            string
	MBackAnimation                           string
	MCharacterUseDistanceWhenActive          float64
	MCostToUse                               string
	MCrouchHoverCancelTime                   float64
	MCurrentBatteryPowerLevel                float64
	MCurrentConnectionLocation               string
	MCurrentHoverMode                        string
	MCurrentMouseDelta                       float64
	MCurrentPlayerVelocity                   float64
	MCurrentPowerLevel                       float64
	MEquipmentSlot                           resource.EquipmentSlot
	MFallSpeedLimitWhenPowered               float64
	MHasConnection                           bool
	MHasPersistentOwner                      bool
	MHoverAccelerationSpeed                  float64
	MHoverFriction                           float64
	MHoverPackActiveTimer                    string
	MHoverSpeed                              float64
	MHoverSprintMultiplier                   float64
	MHoverpackJoystickTimer                  string
	MJumpKeyHoldActivationTime               float64
	MPowerCapacity                           float64
	MPowerConnectionDisconnectionTime        float64
	MPowerConnectionSearchRadius             float64
	MPowerConnectionSearchTickRate           float64
	MPowerConsumption                        float64
	MPowerDrainRate                          float64
	MRailRoadSurfSpeed                       float64
	MRailroadSurfSensitivity                 float64
	MRangeWarningNormalizedDistanceThreshold float64
	MShouldAutomaticallyHoverWhenConnected   bool
	MUseDefaultPrimaryFire                   bool
	M_PreviousHoverMode                      string
	RangeWarningToggleDelegate               string
}

var (
	HoverPack = FGHoverPack{
		Name:                                     "HoverPack",
		ClassName:                                "Equip_HoverPack_C",
		ConnectionLocationUpdatedDelegate:        `()`,
		ConnectionStatusUpdatedDelegate:          `()`,
		HoverModeChangedDelegate:                 `()`,
		MArmAnimation:                            `AE_None`,
		MAttachSocket:                            `root`,
		MBackAnimation:                           `BE_Jetpack`,
		MCharacterUseDistanceWhenActive:          2000.000000,
		MCostToUse:                               ``,
		MCrouchHoverCancelTime:                   0.300000,
		MCurrentBatteryPowerLevel:                0.000000,
		MCurrentConnectionLocation:               `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MCurrentHoverMode:                        `HPM_Inactive`,
		MCurrentMouseDelta:                       0.000000,
		MCurrentPlayerVelocity:                   0.000000,
		MCurrentPowerLevel:                       0.000000,
		MEquipmentSlot:                           resource.Back,
		MFallSpeedLimitWhenPowered:               600.000000,
		MHasConnection:                           false,
		MHasPersistentOwner:                      false,
		MHoverAccelerationSpeed:                  2000.000000,
		MHoverFriction:                           0.990000,
		MHoverPackActiveTimer:                    `()`,
		MHoverSpeed:                              650.000000,
		MHoverSprintMultiplier:                   2.000000,
		MHoverpackJoystickTimer:                  `()`,
		MJumpKeyHoldActivationTime:               0.300000,
		MPowerCapacity:                           60.000000,
		MPowerConnectionDisconnectionTime:        0.100000,
		MPowerConnectionSearchRadius:             3200.000000,
		MPowerConnectionSearchTickRate:           0.100000,
		MPowerConsumption:                        100.000000,
		MPowerDrainRate:                          0.500000,
		MRailRoadSurfSpeed:                       2500.000000,
		MRailroadSurfSensitivity:                 0.100000,
		MRangeWarningNormalizedDistanceThreshold: 0.500000,
		MShouldAutomaticallyHoverWhenConnected:   false,
		MUseDefaultPrimaryFire:                   false,
		M_PreviousHoverMode:                      `HPM_Inactive`,
		RangeWarningToggleDelegate:               `()`,
	}
)

func GetByClassName(className string) (FGHoverPack, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGHoverPack{}, fmt.Errorf("failed to find FGHoverPack with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGHoverPack{
	"Equip_HoverPack_C": HoverPack,
}
