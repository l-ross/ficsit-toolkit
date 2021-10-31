package BuildablePipelineSupport

import (
	"fmt"
)

type FGBuildablePipelineSupport struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCanStack                               bool
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MLength                                 float64
	MMaterialNameToInstanceManager          string
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MStackHeight                            float64
	MToggleDormancyOnInteraction            bool
	MUseStaticHeight                        bool
	MVerticalAngle                          float64
	MaxRenderDistance                       float64
}

var (
	HyperPoleStackable = FGBuildablePipelineSupport{
		ClassName:                               "Build_HyperPoleStackable_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCanStack:                               true,
		MDescription:                            `Support for Hyper Tubes. Can be stacked on other stackable supports.`,
		MDisplayName:                            `Stackable Hyper Tube Support`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MLength:                                 200.000000,
		MMaterialNameToInstanceManager:          `()`,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MStackHeight:                            200.000000,
		MToggleDormancyOnInteraction:            false,
		MUseStaticHeight:                        true,
		MVerticalAngle:                          0.000000,
		MaxRenderDistance:                       -1.000000,
	}

	PipeHyperSupport = FGBuildablePipelineSupport{
		ClassName:                               "Build_PipeHyperSupport_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCanStack:                               false,
		MDescription:                            `Supports for Hyper Tubes to allow for longer distances.`,
		MDisplayName:                            `Hyper Tube Support`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MLength:                                 200.000000,
		MMaterialNameToInstanceManager:          `()`,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MStackHeight:                            200.000000,
		MToggleDormancyOnInteraction:            false,
		MUseStaticHeight:                        false,
		MVerticalAngle:                          0.000000,
		MaxRenderDistance:                       -1.000000,
	}

	PipeSupportStackable = FGBuildablePipelineSupport{
		ClassName:                               "Build_PipeSupportStackable_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCanStack:                               true,
		MDescription:                            `Support for pipelines. Can be stacked on other stackable supports.`,
		MDisplayName:                            `Stackable Pipeline Support`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MLength:                                 200.000000,
		MMaterialNameToInstanceManager:          `()`,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MStackHeight:                            200.000000,
		MToggleDormancyOnInteraction:            false,
		MUseStaticHeight:                        true,
		MVerticalAngle:                          0.000000,
		MaxRenderDistance:                       -1.000000,
	}

	PipelineSupport = FGBuildablePipelineSupport{
		ClassName:                               "Build_PipelineSupport_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCanStack:                               false,
		MDescription: `Can be used as a connection for pipelines. The height of the support can be adjusted.
Useful to route pipelines in a more controlled manner and over long distances.`,
		MDisplayName:                    `Pipeline Support`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MLength:                         200.000000,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MStackHeight:                    200.000000,
		MToggleDormancyOnInteraction:    false,
		MUseStaticHeight:                false,
		MVerticalAngle:                  0.000000,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildablePipelineSupport, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildablePipelineSupport with name %s", className)
}

var classNameToVar = map[string]*FGBuildablePipelineSupport{
	"Build_HyperPoleStackable_C":   &HyperPoleStackable,
	"Build_PipeHyperSupport_C":     &PipeHyperSupport,
	"Build_PipeSupportStackable_C": &PipeSupportStackable,
	"Build_PipelineSupport_C":      &PipelineSupport,
}