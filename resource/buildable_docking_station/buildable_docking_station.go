package BuildableDockingStation

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableDockingStation struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MDescription                            string
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MFuelTransferSpeed                      float64
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MMapText                                string
	MMaterialNameToInstanceManager          string
	MMaxPotential                           float64
	MMaxPotentialIncreasePerCrystal         float64
	MMinPotential                           float64
	MMinimumProducingTime                   float64
	MMinimumStoppedTime                     float64
	MNumCyclesForProductivity               int
	MOnHasPowerChanged                      string
	MOnHasProductionChanged                 string
	MOnHasStandbyChanged                    string
	MOnProductionStatusChanged              string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MStackTransferSize                      float64
	MStorageSizeX                           int
	MStorageSizeY                           int
	MToggleDormancyOnInteraction            bool
	MTransferSpeed                          float64
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	TruckStation = FGBuildableDockingStation{
		ClassName:                               "Build_TruckStation_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDescription: `Either send or receive resources to vehicles. Has an inventory with 48 slots.
Transfers up to 120 stacks per minute to/from docked vehicle. 
Always refuels vehicles if it has access to a matching fuel type.
`,
		MDisplayName:                         `Truck Station`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000021),Translation=(X=300.003357,Y=-1140.064819,Z=175.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000021),Translation=(X=700.003357,Y=-1140.063354,Z=177.326416)),(Rotation=(X=0.000015,Y=-0.707107,Z=-0.707106,W=0.000015),Translation=(X=0.000316,Y=-770.000000,Z=416.844025)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000021),Translation=(X=-299.996643,Y=-1140.063354,Z=175.000000)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MFuelTransferSpeed:                   5.000000,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MMapText:                             ``,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPowerConsumption:                    20.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MStackTransferSize:                   1.000000,
		MStorageSizeX:                        8,
		MStorageSizeY:                        6,
		MToggleDormancyOnInteraction:         false,
		MTransferSpeed:                       0.500000,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableDockingStation, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableDockingStation with name %s", className)
}

var classNameToVar = map[string]*FGBuildableDockingStation{
	"Build_TruckStation_C": &TruckStation,
}
