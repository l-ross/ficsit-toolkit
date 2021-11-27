// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableAttachmentMerger

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableAttachmentMerger struct {
	Name                                 string
	ClassName                            string
	MAddToSignificanceManager            bool
	MAllowColoring                       bool
	MAttachmentPoints                    string
	MBuildEffectSpeed                    float64
	MCachedSkeletalMeshes                string
	MCanChangePotential                  bool
	MCreateClearanceMeshRepresentation   bool
	MCurrentInputIndex                   int
	MCurrentInventoryIndex               int
	MDescription                         string
	MDisplayName                         string
	MDoesHaveShutdownAnimation           bool
	MEffectUpdateInterval                float64
	MFluidStackSizeDefault               resource.StackSize
	MFluidStackSizeMultiplier            int
	MForceNetUpdateOnRegisterPlayer      bool
	MHideOnBuildEffectStart              bool
	MHighlightVector                     string
	MInteractingPlayers                  string
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
	MPowerConsumption                    float64
	MPowerConsumptionExponent            float64
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
	ConveyorAttachmentMerger = FGBuildableAttachmentMerger{
		Name:                                 "ConveyorAttachmentMerger",
		ClassName:                            "Build_ConveyorAttachmentMerger_C",
		MAddToSignificanceManager:            false,
		MAllowColoring:                       true,
		MAttachmentPoints:                    ``,
		MBuildEffectSpeed:                    0.000000,
		MCachedSkeletalMeshes:                ``,
		MCanChangePotential:                  false,
		MCreateClearanceMeshRepresentation:   true,
		MCurrentInputIndex:                   -1,
		MCurrentInventoryIndex:               0,
		MDescription:                         `Merges up to three conveyor belts into one.`,
		MDisplayName:                         `Conveyor Merger`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           false,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                -1.000000,
		MMinimumStoppedTime:                  -1.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (FGBuildableAttachmentMerger, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableAttachmentMerger{}, fmt.Errorf("failed to find FGBuildableAttachmentMerger with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableAttachmentMerger{
	"Build_ConveyorAttachmentMerger_C": ConveyorAttachmentMerger,
}
