// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableFrackingActivator

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableFrackingActivator struct {
	Name                                   string
	ClassName                              string
	ConnectedExtractorCountChangedDelegate string
	CurrentPotentialChangedDelegate        string
	MActivationStartupTime                 float64
	MActivationStartupTimer                float64
	MAddToSignificanceManager              bool
	MAllowColoring                         bool
	MAllowedResourceForms                  string
	MAllowedResources                      string
	MAttachmentPoints                      string
	MBuildEffectSpeed                      float64
	MCachedSkeletalMeshes                  string
	MCanChangePotential                    bool
	MConnectedExtractorCount               int
	MCreateClearanceMeshRepresentation     bool
	MDefaultPotentialExtractionPerMinute   float64
	MDescription                           string
	MDisplayName                           string
	MDoesHaveShutdownAnimation             bool
	MEffectUpdateInterval                  float64
	MExtractorTypeName                     string
	MFluidStackSizeDefault                 resource.StackSize
	MFluidStackSizeMultiplier              int
	MForceNetUpdateOnRegisterPlayer        bool
	MHideOnBuildEffectStart                bool
	MHighlightVector                       string
	MInteractingPlayers                    string
	MIsUseable                             bool
	MMaxPotential                          float64
	MMaxPotentialIncreasePerCrystal        float64
	MMinPotential                          float64
	MMinimumProducingTime                  float64
	MMinimumStoppedTime                    float64
	MNumCyclesForProductivity              int
	MOnHasPowerChanged                     string
	MOnHasProductionChanged                string
	MOnHasStandbyChanged                   string
	MOnlyAllowCertainResources             bool
	MPowerConsumption                      float64
	MPowerConsumptionExponent              float64
	MSatelliteActivationComplete           bool
	MSatelliteNodeCount                    int
	MShouldModifyWorldGrid                 bool
	MShouldShowAttachmentPointVisuals      bool
	MShouldShowHighlight                   bool
	MSignificanceRange                     float64
	MSkipBuildEffect                       bool
	MToggleDormancyOnInteraction           bool
	MaxRenderDistance                      float64
	OnReplicationDetailActorCreatedEvent   string
}

var (
	FrackingSmasher = FGBuildableFrackingActivator{
		Name:                                   "FrackingSmasher",
		ClassName:                              "Build_FrackingSmasher_C",
		ConnectedExtractorCountChangedDelegate: `()`,
		CurrentPotentialChangedDelegate:        `()`,
		MActivationStartupTime:                 16.500000,
		MActivationStartupTimer:                0.000000,
		MAddToSignificanceManager:              true,
		MAllowColoring:                         true,
		MAllowedResourceForms:                  `(RF_LIQUID,RF_GAS,RF_HEAT)`,
		MAllowedResources:                      `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/CrudeOil/Desc_LiquidOil.Desc_LiquidOil_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/NitrogenGas/Desc_NitrogenGas.Desc_NitrogenGas_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Water/Desc_Water.Desc_Water_C"')`,
		MAttachmentPoints:                      ``,
		MBuildEffectSpeed:                      0.000000,
		MCachedSkeletalMeshes:                  ``,
		MCanChangePotential:                    true,
		MConnectedExtractorCount:               0,
		MCreateClearanceMeshRepresentation:     true,
		MDefaultPotentialExtractionPerMinute:   0.000000,
		MDescription: `Can be placed on a Resource Well to activate it by pressurizing the underground resource.
Once activated, Resource Well Extractors can be placed on the surrounding sub-nodes to extract the resource.
Requires Power. Overclocking increases the output potential of the entire Resource Well.`,
		MDisplayName:                         `Resource Well Pressurizer`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MExtractorTypeName:                   `None`,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnlyAllowCertainResources:           true,
		MPowerConsumption:                    150.000000,
		MPowerConsumptionExponent:            1.600000,
		MSatelliteActivationComplete:         false,
		MSatelliteNodeCount:                  0,
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

func GetByClassName(className string) (FGBuildableFrackingActivator, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableFrackingActivator{}, fmt.Errorf("failed to find FGBuildableFrackingActivator with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableFrackingActivator{
	"Build_FrackingSmasher_C": FrackingSmasher,
}
