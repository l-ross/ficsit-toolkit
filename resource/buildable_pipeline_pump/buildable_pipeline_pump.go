package BuildablePipelinePump

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildablePipelinePump struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MAnimSpeed                              float64
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MCurrentAudioHeadLift                   float64
	MDefaultFlowLimit                       float64
	MDescription                            string
	MDesignPressure                         float64
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
	MIndicatorData                          string
	MInteractingPlayers                     string
	MIsExceedingHeadLift                    bool
	MIsPipePumpPlaying                      bool
	MIsUseable                              bool
	MLastFlowUpdate                         float64
	MLastFlowValue                          float64
	MMaterialNameToInstanceManager          string
	MMaxPotential                           float64
	MMaxPotentialIncreasePerCrystal         float64
	MMaxPressure                            float64
	MMinPotential                           float64
	MMinimumFlowPercentForStandby           float64
	MMinimumProducingTime                   float64
	MMinimumStoppedTime                     float64
	MNumCyclesForProductivity               int
	MOnHasPowerChanged                      string
	MOnHasProductionChanged                 string
	MOnHasStandbyChanged                    string
	MOnProductionStatusChanged              string
	MPipeConnections                        string
	MPistonAudioTimer                       string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MRadius                                 float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MTimeScaleOffset                        float64
	MToggleDormancyOnInteraction            bool
	MUpdateFlowTime                         float64
	MUserFlowLimit                          float64
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	PipelinePump = FGBuildablePipelinePump{
		ClassName:                               "Build_PipelinePump_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAnimSpeed:                              0.000000,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MCurrentAudioHeadLift:                   0.000000,
		MDefaultFlowLimit:                       10.000000,
		MDescription: `Can be attached to a Pipeline to apply Head Lift.
Maximum Head Lift: 20 meters
(Allows fluids to be transported 20 meters upwards.)

NOTE: Has an in- and output direction.
NOTE: Head Lift does not stack, so space between Pumps is recommended.`,
		MDesignPressure:                      20.000000,
		MDisplayName:                         `Pipeline Pump Mk.1`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidBox:                            `()`,
		MFluidBoxVolume:                      5.000000,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MIndicatorData:                       `()`,
		MInteractingPlayers:                  ``,
		MIsExceedingHeadLift:                 false,
		MIsPipePumpPlaying:                   false,
		MIsUseable:                           true,
		MLastFlowUpdate:                      0.000000,
		MLastFlowValue:                       0.000000,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMaxPressure:                         22.000000,
		MMinPotential:                        0.010000,
		MMinimumFlowPercentForStandby:        0.050000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPipeConnections:                     ``,
		MPowerConsumption:                    4.000000,
		MPowerConsumptionExponent:            1.600000,
		MRadius:                              65.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   12000.000000,
		MSkipBuildEffect:                     false,
		MTimeScaleOffset:                     0.000000,
		MToggleDormancyOnInteraction:         false,
		MUpdateFlowTime:                      0.500000,
		MUserFlowLimit:                       -1.000000,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	PipelinePumpMk2 = FGBuildablePipelinePump{
		ClassName:                               "Build_PipelinePumpMk2_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAnimSpeed:                              0.000000,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MCurrentAudioHeadLift:                   0.000000,
		MDefaultFlowLimit:                       10.000000,
		MDescription: `Can be attached to a Pipeline to apply Head Lift.
Maximum Head Lift: 50 meters
(Allows fluids to be transported 50 meters upwards.)

NOTE: Has an in- and output direction.
NOTE: Head Lift does not stack, so space between Pumps is recommended.`,
		MDesignPressure:                      50.000000,
		MDisplayName:                         `Pipeline Pump Mk.2`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidBox:                            `()`,
		MFluidBoxVolume:                      5.000000,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MIndicatorData:                       `()`,
		MInteractingPlayers:                  ``,
		MIsExceedingHeadLift:                 false,
		MIsPipePumpPlaying:                   false,
		MIsUseable:                           true,
		MLastFlowUpdate:                      0.000000,
		MLastFlowValue:                       0.000000,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMaxPressure:                         55.000000,
		MMinPotential:                        0.010000,
		MMinimumFlowPercentForStandby:        0.050000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPipeConnections:                     ``,
		MPistonAudioTimer:                    `()`,
		MPowerConsumption:                    8.000000,
		MPowerConsumptionExponent:            1.600000,
		MRadius:                              65.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MTimeScaleOffset:                     0.000000,
		MToggleDormancyOnInteraction:         false,
		MUpdateFlowTime:                      0.500000,
		MUserFlowLimit:                       -1.000000,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	Valve = FGBuildablePipelinePump{
		ClassName:                               "Build_Valve_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAnimSpeed:                              0.000000,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDefaultFlowLimit:                       10.000000,
		MDescription: `Used to limit Pipeline flow rates.
Can be attached to a Pipeline.

NOTE: Has an in- and output direction.`,
		MDesignPressure:                      0.000000,
		MDisplayName:                         `Valve`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidBox:                            `()`,
		MFluidBoxVolume:                      5.000000,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MIndicatorData:                       `()`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MLastFlowUpdate:                      0.000000,
		MLastFlowValue:                       0.000000,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMaxPressure:                         0.000000,
		MMinPotential:                        0.010000,
		MMinimumFlowPercentForStandby:        0.050000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPipeConnections:                     ``,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MRadius:                              65.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   12000.000000,
		MSkipBuildEffect:                     false,
		MTimeScaleOffset:                     0.000000,
		MToggleDormancyOnInteraction:         false,
		MUpdateFlowTime:                      0.500000,
		MUserFlowLimit:                       -1.000000,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildablePipelinePump, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildablePipelinePump with name %s", className)
}

var classNameToVar = map[string]*FGBuildablePipelinePump{
	"Build_PipelinePump_C":    &PipelinePump,
	"Build_PipelinePumpMk2_C": &PipelinePumpMk2,
	"Build_Valve_C":           &Valve,
}