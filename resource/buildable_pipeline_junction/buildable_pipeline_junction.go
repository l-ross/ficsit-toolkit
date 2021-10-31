package BuildablePipelineJunction

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildablePipelineJunction struct {
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
	MFluidBox                               string
	MFluidBoxVolume                         float64
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
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
	MPipeConnections                        string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MRadius                                 float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	PipelineJunctionCross = FGBuildablePipelineJunction{
		ClassName:                               "Build_PipelineJunction_Cross_C",
		MAddToSignificanceManager:               false,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDescription:                            `Can be attached to a pipeline to split it 4-way.`,
		MDisplayName:                            `Pipeline Junction Cross`,
		MEffectUpdateInterval:                   0.000000,
		MExcludeFromMaterialInstancing:          ``,
		MFluidBox:                               `()`,
		MFluidBoxVolume:                         5.000000,
		MFluidStackSizeDefault:                  resource.Fluid,
		MFluidStackSizeMultiplier:               1,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MMaterialNameToInstanceManager:          `()`,
		MMaxPotential:                           1.000000,
		MMaxPotentialIncreasePerCrystal:         0.500000,
		MMinPotential:                           0.010000,
		MMinimumProducingTime:                   2.000000,
		MMinimumStoppedTime:                     5.000000,
		MNumCyclesForProductivity:               20,
		MOnHasPowerChanged:                      `()`,
		MOnHasProductionChanged:                 `()`,
		MOnHasStandbyChanged:                    `()`,
		MOnProductionStatusChanged:              `()`,
		MPipeConnections:                        ``,
		MPowerConsumption:                       0.000000,
		MPowerConsumptionExponent:               1.600000,
		MRadius:                                 65.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSignificanceRange:                      18000.000000,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
		OnReplicationDetailActorCreatedEvent:    `()`,
	}
)

func GetByClassName(className string) (*FGBuildablePipelineJunction, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildablePipelineJunction with name %s", className)
}

var classNameToVar = map[string]*FGBuildablePipelineJunction{
	"Build_PipelineJunction_Cross_C": &PipelineJunctionCross,
}
