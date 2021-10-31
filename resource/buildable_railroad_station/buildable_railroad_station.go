package BuildableRailroadStation

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableRailroadStation struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MDescription                            string
	MDisplayName                            string
	MDockedPlatformList                     string
	MDockingSequenceTimerHandle             string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsOrientationReversed                  bool
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
	MPlatformConnections                    string
	MPlatformDockingStatus                  string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MSavedDockingStatus                     string
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	TrainStation = FGBuildableRailroadStation{
		ClassName:                               "Build_TrainStation_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDescription: `Locomotives can be set to drive to and stop at the Train Station.
You can connect power to the train station to power up the trains on the railway as well as transport the power to other stations.`,
		MDisplayName:                         `Train Station`,
		MDockedPlatformList:                  ``,
		MDockingSequenceTimerHandle:          `()`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsOrientationReversed:               false,
		MIsUseable:                           true,
		MMapText:                             `Train Station`,
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
		MPlatformConnections:                 `(FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainStation.Default__Build_TrainStation_C:PlatformConnection0"',FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainStation.Default__Build_TrainStation_C:PlatformConnection1"')`,
		MPlatformDockingStatus:               `ETPDS_None`,
		MPowerConsumption:                    50.000000,
		MPowerConsumptionExponent:            1.600000,
		MSavedDockingStatus:                  `ETPDS_None`,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableRailroadStation, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableRailroadStation with name %s", className)
}

var classNameToVar = map[string]*FGBuildableRailroadStation{
	"Build_TrainStation_C": &TrainStation,
}