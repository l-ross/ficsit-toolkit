package BuildableWire

import (
	"fmt"
)

type FGBuildableWire struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MConnections                            string
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MLengthPerCost                          float64
	MMaterialNameToInstanceManager          string
	MMaxLength                              float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
}

var (
	PowerLine = FGBuildableWire{
		ClassName:                               "Build_PowerLine_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MConnections:                            `None`,
		MDescription:                            `Used to connect Power Poles, Power Generators and Factory buildings.`,
		MDisplayName:                            `Power Line`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MLengthPerCost:                          2500.000000,
		MMaterialNameToInstanceManager:          `()`,
		MMaxLength:                              10000.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       1000.000000,
	}

	XmassLightsLine = FGBuildableWire{
		ClassName:                               "Build_XmassLightsLine_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MConnections:                            `None`,
		MDescription:                            `Used to connect Power Poles, Power Generators and Factory buildings. Has pretty lights to boot!`,
		MDisplayName:                            `FICSMAS Power Light`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MLengthPerCost:                          2500.000000,
		MMaterialNameToInstanceManager:          `()`,
		MMaxLength:                              10000.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       1000.000000,
	}
)

func GetByClassName(className string) (*FGBuildableWire, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableWire with name %s", className)
}

var classNameToVar = map[string]*FGBuildableWire{
	"Build_PowerLine_C":       &PowerLine,
	"Build_XmassLightsLine_C": &XmassLightsLine,
}