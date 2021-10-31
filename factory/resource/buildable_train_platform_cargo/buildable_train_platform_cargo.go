// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableTrainPlatformCargo

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/factory/resource"
)

type FGBuildableTrainPlatformCargo struct {
	Name                                    string
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MCanLoadAny                             bool
	MCanUnloadAny                           bool
	MDescription                            string
	MDisplayName                            string
	MDockingSequenceTimerHandle             string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MFreightCargoType                       string
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsFullLoad                             bool
	MIsFullUnload                           bool
	MIsOrientationReversed                  bool
	MIsUseable                              bool
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
	MPipeInputConnections                   string
	MPipeOutputConnections                  string
	MPlatformConnections                    string
	MPlatformDockingStatus                  string
	MPotentialDockers                       string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MReplicatedInflowRate                   float64
	MReplicatedOutflowRate                  float64
	MSavedDockingStatus                     string
	MShouldExecuteLoadOrUnload              bool
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MStorageInputConnections                string
	MStorageSizeX                           int
	MStorageSizeY                           int
	MSwapCargoVisibilityTimerHandle         string
	MTimeToCompleteLoad                     float64
	MTimeToCompleteUnload                   float64
	MTimeToSwapLoadVisibility               float64
	MTimeToSwapUnloadVisibility             float64
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	TrainDockingStation = FGBuildableTrainPlatformCargo{
		Name:                                    "TrainDockingStation",
		ClassName:                               "Build_TrainDockingStation_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MCanLoadAny:                             false,
		MCanUnloadAny:                           false,
		MDescription: `Freight Cars that stop at the Freight Platform will be loaded or unloaded by the Freight Platform.
Loading and unloading options can be set inside the building.
Snaps to other Platforms and Stations.
Needs to be connected to a powered Railway to function.`,
		MDisplayName:                         `Freight Platform`,
		MDockingSequenceTimerHandle:          `()`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Translation=(X=300.000000,Y=1666.997192,Z=559.999207)),(Translation=(X=300.622345,Y=1655.377319,Z=162.239334)),(Translation=(X=-300.000000,Y=1666.186157,Z=160.540833)),(Translation=(X=-300.000000,Y=1660.841187,Z=562.163086)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MFreightCargoType:                    `FCT_Standard`,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsFullLoad:                          false,
		MIsFullUnload:                        false,
		MIsOrientationReversed:               false,
		MIsUseable:                           true,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPlatformConnections:                 `(FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainDockingStation.Default__Build_TrainDockingStation_C:PlatformConnection0"',FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainDockingStation.Default__Build_TrainDockingStation_C:PlatformConnection1"')`,
		MPlatformDockingStatus:               `ETPDS_None`,
		MPotentialDockers:                    ``,
		MPowerConsumption:                    50.000000,
		MPowerConsumptionExponent:            1.600000,
		MReplicatedInflowRate:                0.000000,
		MReplicatedOutflowRate:               0.000000,
		MSavedDockingStatus:                  `ETPDS_None`,
		MShouldExecuteLoadOrUnload:           false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MStorageInputConnections:             ``,
		MStorageSizeX:                        8,
		MStorageSizeY:                        6,
		MSwapCargoVisibilityTimerHandle:      `()`,
		MTimeToCompleteLoad:                  27.000000,
		MTimeToCompleteUnload:                27.000000,
		MTimeToSwapLoadVisibility:            20.000000,
		MTimeToSwapUnloadVisibility:          6.000000,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	TrainDockingStationLiquid = FGBuildableTrainPlatformCargo{
		Name:                                    "TrainDockingStationLiquid",
		ClassName:                               "Build_TrainDockingStationLiquid_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MCanLoadAny:                             false,
		MCanUnloadAny:                           false,
		MDescription: `Freight Cars that stop at the Freight Platform will be loaded or unloaded by the Freight Platform.
Loading and unloading options can be set inside the building.
Snaps to other Platforms and Stations.
Needs to be connected to a powered Railway to function.`,
		MDisplayName:                         `Fluid Freight Platform`,
		MDockingSequenceTimerHandle:          `()`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            48,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MFreightCargoType:                    `FCT_Liquid`,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsFullLoad:                          false,
		MIsFullUnload:                        false,
		MIsOrientationReversed:               false,
		MIsUseable:                           true,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPlatformConnections:                 `(FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainDockingStationLiquid.Default__Build_TrainDockingStationLiquid_C:PlatformConnection0"',FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainDockingStationLiquid.Default__Build_TrainDockingStationLiquid_C:PlatformConnection1"')`,
		MPlatformDockingStatus:               `ETPDS_None`,
		MPotentialDockers:                    ``,
		MPowerConsumption:                    50.000000,
		MPowerConsumptionExponent:            1.600000,
		MReplicatedInflowRate:                0.000000,
		MReplicatedOutflowRate:               0.000000,
		MSavedDockingStatus:                  `ETPDS_None`,
		MShouldExecuteLoadOrUnload:           false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MStorageInputConnections:             ``,
		MStorageSizeX:                        1,
		MStorageSizeY:                        1,
		MSwapCargoVisibilityTimerHandle:      `()`,
		MTimeToCompleteLoad:                  27.000000,
		MTimeToCompleteUnload:                27.000000,
		MTimeToSwapLoadVisibility:            20.000000,
		MTimeToSwapUnloadVisibility:          6.000000,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableTrainPlatformCargo, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableTrainPlatformCargo with class name %s", className)
}

var classNameToVar = map[string]*FGBuildableTrainPlatformCargo{
	"Build_TrainDockingStation_C":       &TrainDockingStation,
	"Build_TrainDockingStationLiquid_C": &TrainDockingStationLiquid,
}
