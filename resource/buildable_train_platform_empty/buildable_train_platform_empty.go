// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableTrainPlatformEmpty

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableTrainPlatformEmpty struct {
	Name                                 string
	ClassName                            string
	MAddToSignificanceManager            bool
	MAllowColoring                       bool
	MAttachmentPoints                    string
	MBuildEffectSpeed                    float64
	MCachedSkeletalMeshes                string
	MCanChangePotential                  bool
	MCreateClearanceMeshRepresentation   bool
	MDescription                         string
	MDisplayName                         string
	MDockWasCancelled                    bool
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
	MSignificanceRange                   float64
	MSkipBuildEffect                     bool
	MToggleDormancyOnInteraction         bool
	MaxRenderDistance                    float64
	OnReplicationDetailActorCreatedEvent string
}

var (
	TrainPlatformEmpty = FGBuildableTrainPlatformEmpty{
		Name:                                 "TrainPlatformEmpty",
		ClassName:                            "Build_TrainPlatformEmpty_C",
		MAddToSignificanceManager:            true,
		MAllowColoring:                       true,
		MAttachmentPoints:                    ``,
		MBuildEffectSpeed:                    0.000000,
		MCachedSkeletalMeshes:                ``,
		MCanChangePotential:                  false,
		MCreateClearanceMeshRepresentation:   true,
		MDescription:                         `An empty train platform for when you need to create some empty space.`,
		MDisplayName:                         `Empty Platform`,
		MDockWasCancelled:                    false,
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
		MIsUseable:                           false,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPlatformConnections:                 `(FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainPlatformEmpty.Default__Build_TrainPlatformEmpty_C:PlatformConnection0"',FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainPlatformEmpty.Default__Build_TrainPlatformEmpty_C:PlatformConnection1"')`,
		MPlatformDockingStatus:               `ETPDS_None`,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MSavedDockingStatus:                  `ETPDS_None`,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	TrainPlatformEmpty02 = FGBuildableTrainPlatformEmpty{
		Name:                                 "TrainPlatformEmpty02",
		ClassName:                            "Build_TrainPlatformEmpty_02_C",
		MAddToSignificanceManager:            true,
		MAllowColoring:                       true,
		MAttachmentPoints:                    ``,
		MBuildEffectSpeed:                    0.000000,
		MCachedSkeletalMeshes:                ``,
		MCanChangePotential:                  false,
		MCreateClearanceMeshRepresentation:   true,
		MDescription:                         `An empty train platform for when you need to create some empty space.`,
		MDisplayName:                         `Empty Platform With Catwalk`,
		MDockWasCancelled:                    false,
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
		MIsUseable:                           false,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPlatformConnections:                 `(FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainPlatformEmpty_02.Default__Build_TrainPlatformEmpty_02_C:PlatformConnection0"',FGTrainPlatformConnection'"/Game/FactoryGame/Buildable/Factory/Train/Station/Build_TrainPlatformEmpty_02.Default__Build_TrainPlatformEmpty_02_C:PlatformConnection1"')`,
		MPlatformDockingStatus:               `ETPDS_None`,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MSavedDockingStatus:                  `ETPDS_None`,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (FGBuildableTrainPlatformEmpty, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableTrainPlatformEmpty{}, fmt.Errorf("failed to find FGBuildableTrainPlatformEmpty with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableTrainPlatformEmpty{
	"Build_TrainPlatformEmpty_C":    TrainPlatformEmpty,
	"Build_TrainPlatformEmpty_02_C": TrainPlatformEmpty02,
}
