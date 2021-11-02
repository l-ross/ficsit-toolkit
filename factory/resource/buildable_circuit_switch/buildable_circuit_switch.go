// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableCircuitSwitch

import (
	"fmt"
)

type FGBuildableCircuitSwitch struct {
	Name                                    string
	ClassName                               string
	BIsSignificant                          bool
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
}

var (
	PowerSwitch = FGBuildableCircuitSwitch{
		Name:                                    "PowerSwitch",
		ClassName:                               "Build_PowerSwitch_C",
		BIsSignificant:                          false,
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
	}
)

func GetByClassName(className string) (FGBuildableCircuitSwitch, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return FGBuildableCircuitSwitch{}, fmt.Errorf("failed to find FGBuildableCircuitSwitch with class name %s", className)
}

var classNameToVar = map[string]FGBuildableCircuitSwitch{
	"Build_PowerSwitch_C": PowerSwitch,
}
