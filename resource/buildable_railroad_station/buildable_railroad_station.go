// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableRailroadStation

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableRailroadStation struct {
	Name                                 string
	ClassName                            string
	MAddToSignificanceManager            bool
	MAllowColoring                       bool
	MAttachmentPoints                    string
	MBuildEffectSpeed                    float64
	MCachedSkeletalMeshes                string
	MCanChangePotential                  bool
	MCreateClearanceMeshRepresentation   bool
	MCurrentDockForDuration              float64
	MCurrentDockedWithRuleSet            string
	MDescription                         string
	MDisplayName                         string
	MDockWasCancelled                    bool
	MDockedPlatformList                  string
	MDockingSequenceTimerHandle          string
	MDoesHaveShutdownAnimation           bool
	MEffectUpdateInterval                float64
	MFluidStackSizeDefault               resource.StackSize
	MFluidStackSizeMultiplier            int
	MForceNetUpdateOnRegisterPlayer      bool
	MHideOnBuildEffectStart              bool
	MHighlightVector                     string
	MIdleUpdateTimerHandle               string
	MInteractingPlayers                  string
	MIsOrientationReversed               bool
	MIsUseable                           bool
	MMapText                             string
	MMaxPotential                        float64
	MMaxPotentialIncreasePerCrystal      float64
	MMinPotential                        float64
	MMinimumProducingTime                float64
	MMinimumStoppedTime                  float64
	MNumCyclesForProductivity            int
	MOnHasPowerChanged                   string
	MOnHasProductionChanged              string
	MOnHasStandbyChanged                 string
	MPlatformConnections                 string
	MPlatformDockingStatus               string
	MPowerConsumption                    float64
	MPowerConsumptionExponent            float64
	MSavedDockingStatus                  string
	MShouldModifyWorldGrid               bool
	MShouldShowAttachmentPointVisuals    bool
	MShouldShowHighlight                 bool
	MShouldTeleportHere                  bool
	MSignificanceRange                   float64
	MSkipBuildEffect                     bool
	MToggleDormancyOnInteraction         bool
	MaxRenderDistance                    float64
	OnReplicationDetailActorCreatedEvent string
}

var (
	TrainStation = FGBuildableRailroadStation{
		Name:                               "TrainStation",
		ClassName:                          "Build_TrainStation_C",
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                false,
		MCreateClearanceMeshRepresentation: true,
		MCurrentDockForDuration:            0.000000,
		MCurrentDockedWithRuleSet:          `(DockForDuration=15.000000)`,
		MDescription: `Locomotives can be set to drive to and stop at the Train Station.
You can connect power to the train station to power up the trains on the railway as well as transport the power to other stations.`,
		MDisplayName:                         `Train Station`,
		MDockWasCancelled:                    false,
		MDockedPlatformList:                  ``,
		MDockingSequenceTimerHandle:          `()`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MIdleUpdateTimerHandle:               `()`,
		MInteractingPlayers:                  ``,
		MIsOrientationReversed:               false,
		MIsUseable:                           true,
		MMapText:                             `Train Station`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPlatformConnections:                 `(FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainStation.Default__Build_TrainStation_C:PlatformConnection0"',FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainStation.Default__Build_TrainStation_C:PlatformConnection1"')`,
		MPlatformDockingStatus:               `ETPDS_None`,
		MPowerConsumption:                    50.000000,
		MPowerConsumptionExponent:            1.600000,
		MSavedDockingStatus:                  `ETPDS_None`,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    true,
		MShouldShowHighlight:                 false,
		MShouldTeleportHere:                  false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (FGBuildableRailroadStation, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableRailroadStation{}, fmt.Errorf("failed to find FGBuildableRailroadStation with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableRailroadStation{
	"Build_TrainStation_C": TrainStation,
}
