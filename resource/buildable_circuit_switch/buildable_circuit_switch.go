package BuildableCircuitSwitch

import (
	"fmt"
)

type FGBuildableCircuitSwitch struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MBuildingTag                            string
	MConnections                            string
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHasBuildingTag                         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsBridgeConnected                      bool
	MIsSwitchOn                             bool
	MIsUseable                              bool
	MMaterialNameToInstanceManager          string
	MMaxCharacters                          int
	MOnIsConnectedChanged                   string
	MOnIsSwitchOnChanged                    string
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MTextRenderers                          string
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	bIsSignificant                          bool
}

var (
	PowerSwitch = FGBuildableCircuitSwitch{
		ClassName:                               "Build_PowerSwitch_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MBuildingTag:                            ``,
		MConnections:                            ``,
		MDescription: `Can be switched ON/OFF to enable/disable the connection between two Power Grids.

Note the location of the A and B connection.`,
		MDisplayName:                    `Power Switch`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHasBuildingTag:                 false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsBridgeConnected:              false,
		MIsSwitchOn:                     false,
		MIsUseable:                      true,
		MMaterialNameToInstanceManager:  `()`,
		MMaxCharacters:                  20,
		MOnIsConnectedChanged:           `()`,
		MOnIsSwitchOnChanged:            `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MTextRenderers:                  ``,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
		bIsSignificant:                  false,
	}
)

func GetByClassName(className string) (*FGBuildableCircuitSwitch, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableCircuitSwitch with name %s", className)
}

var classNameToVar = map[string]*FGBuildableCircuitSwitch{
	"Build_PowerSwitch_C": &PowerSwitch,
}
