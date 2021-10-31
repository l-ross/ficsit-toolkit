package BuildablePipeline

import (
	"fmt"
)

type FGBuildablePipeline struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCurrentFluid                           string
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFlowIndicatorMinimumPipeLength         float64
	MFlowLimit                              float64
	MFluidBox                               string
	MFluidNames                             string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MIgnoreActorsForIndicator               string
	MIndicatorData                          string
	MInteractingPlayers                     string
	MIsRattling                             bool
	MIsUseable                              bool
	MMaterialNameToInstanceManager          string
	MMaxIndicatorTurnAngle                  float64
	MMeshLength                             float64
	MPipeConnections                        string
	MQuantiziedContent                      float64
	MQuantiziedFlow                         float64
	MRadius                                 float64
	MRattleLimit                            float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MSplineData                             string
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
}

var (
	Pipeline = FGBuildablePipeline{
		ClassName:                               "Build_Pipeline_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCurrentFluid:                           ``,
		MDescription: `Outside indicators show volume, flow rate and direction.
Transports up to 300m³ of fluid per minute.
Used to transport fluids.`,
		MDisplayName:                    `Pipeline Mk.1`,
		MExcludeFromMaterialInstancing:  ``,
		MFlowIndicatorMinimumPipeLength: 400.000000,
		MFlowLimit:                      5.000000,
		MFluidBox:                       `()`,
		MFluidNames:                     `((WwiseSafeName="Crude_Oil",ActualName="Crude Oil"),(WwiseSafeName="Fuel",ActualName="Fuel"),(WwiseSafeName="Heavy_Oil_Residue",ActualName="Heavy Oil Residue"),(WwiseSafeName="Water",ActualName="Water"),(WwiseSafeName="No_Fluid"),(WwiseSafeName="Gas",ActualName="Nitrogen Gas"))`,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MIgnoreActorsForIndicator:       ``,
		MIndicatorData:                  `()`,
		MInteractingPlayers:             ``,
		MIsRattling:                     false,
		MIsUseable:                      true,
		MMaterialNameToInstanceManager:  `()`,
		MMaxIndicatorTurnAngle:          3.000000,
		MMeshLength:                     200.000000,
		MPipeConnections:                ``,
		MQuantiziedContent:              0.000000,
		MQuantiziedFlow:                 0.000000,
		MRadius:                         65.000000,
		MRattleLimit:                    0.500000,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MSplineData:                     ``,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PipelineMK2 = FGBuildablePipeline{
		ClassName:                               "Build_PipelineMK2_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCurrentFluid:                           ``,
		MDescription: `Outside indicators show volume, flow rate and direction.
Transports up to 600m³ of fluid per minute.
Used to transport fluids.`,
		MDisplayName:                    `Pipeline Mk.2`,
		MExcludeFromMaterialInstancing:  ``,
		MFlowIndicatorMinimumPipeLength: 400.000000,
		MFlowLimit:                      10.000000,
		MFluidBox:                       `()`,
		MFluidNames:                     `((WwiseSafeName="Crude_Oil",ActualName="Crude Oil"),(WwiseSafeName="Fuel",ActualName="Fuel"),(WwiseSafeName="Heavy_Oil_Residue",ActualName="Heavy Oil Residue"),(WwiseSafeName="Water",ActualName="Water"),(WwiseSafeName="No_Fluid"),(WwiseSafeName="Gas",ActualName="Nitrogen Gas"))`,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MIgnoreActorsForIndicator:       ``,
		MIndicatorData:                  `()`,
		MInteractingPlayers:             ``,
		MIsRattling:                     false,
		MIsUseable:                      true,
		MMaterialNameToInstanceManager:  `()`,
		MMaxIndicatorTurnAngle:          3.000000,
		MMeshLength:                     200.000000,
		MPipeConnections:                ``,
		MQuantiziedContent:              0.000000,
		MQuantiziedFlow:                 0.000000,
		MRadius:                         65.000000,
		MRattleLimit:                    0.500000,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MSplineData:                     ``,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildablePipeline, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildablePipeline with name %s", className)
}

var classNameToVar = map[string]*FGBuildablePipeline{
	"Build_Pipeline_C":    &Pipeline,
	"Build_PipelineMK2_C": &PipelineMK2,
}
