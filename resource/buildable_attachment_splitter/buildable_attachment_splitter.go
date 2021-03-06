// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableAttachmentSplitter

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableAttachmentSplitter struct {
	Name                                 string
	ClassName                            string
	MAddToSignificanceManager            bool
	MAllowColoring                       bool
	MAttachmentPoints                    string
	MBuildEffectSpeed                    float64
	MCachedSkeletalMeshes                string
	MCanChangePotential                  bool
	MCreateClearanceMeshRepresentation   bool
	MCurrentInventoryIndex               int
	MCurrentOutputIndex                  int
	MDescription                         string
	MDisplayName                         string
	MDistributionTable                   string
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
	ConveyorAttachmentSplitter = FGBuildableAttachmentSplitter{
		Name:                               "ConveyorAttachmentSplitter",
		ClassName:                          "Build_ConveyorAttachmentSplitter_C",
		MAddToSignificanceManager:          false,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                false,
		MCreateClearanceMeshRepresentation: true,
		MCurrentInventoryIndex:             0,
		MCurrentOutputIndex:                -1,
		MDescription: `Splits conveyor belts in three. 
Useful to move parts and resources from oversaturated conveyor belts.`,
		MDisplayName:                         `Conveyor Splitter`,
		MDistributionTable:                   ``,
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

func GetByClassName(className string) (FGBuildableAttachmentSplitter, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableAttachmentSplitter{}, fmt.Errorf("failed to find FGBuildableAttachmentSplitter with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableAttachmentSplitter{
	"Build_ConveyorAttachmentSplitter_C": ConveyorAttachmentSplitter,
}
