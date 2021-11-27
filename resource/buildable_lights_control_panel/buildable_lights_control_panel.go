// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableLightsControlPanel

import (
	"fmt"
)

type FGBuildableLightsControlPanel struct {
	Name                               string
	ClassName                          string
	MAllowColoring                     bool
	MAttachmentPoints                  string
	MBuildEffectSpeed                  float64
	MConnections                       string
	MControlledBuildables              string
	MCreateClearanceMeshRepresentation bool
	MDescription                       string
	MDisplayName                       string
	MForceNetUpdateOnRegisterPlayer    bool
	MHideOnBuildEffectStart            bool
	MHighlightVector                   string
	MInteractingPlayers                string
	MIsBridgeConnected                 bool
	MIsEnabled                         bool
	MIsUseable                         bool
	MLightControlData                  string
	MOnControlledBuildablesChanged     string
	MShouldModifyWorldGrid             bool
	MShouldShowAttachmentPointVisuals  bool
	MShouldShowHighlight               bool
	MSkipBuildEffect                   bool
	MToggleDormancyOnInteraction       bool
	MaxRenderDistance                  float64
	OnLightControlPanelStateChanged    string
}

var (
	LightsControlPanel = FGBuildableLightsControlPanel{
		Name:                               "LightsControlPanel",
		ClassName:                          "Build_LightsControlPanel_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MConnections:                       ``,
		MControlledBuildables:              ``,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Useful for sectioning and modifying many lights at once.

Controls all Lights connected to the Power Grid attached to the 'Light Power Connector'.
(Other Control Panels and Power Switches interrupt the connection.)`,
		MDisplayName:                      `Lights Control Panel`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsBridgeConnected:                false,
		MIsEnabled:                        false,
		MIsUseable:                        true,
		MLightControlData:                 `(Intensity=1.000000)`,
		MOnControlledBuildablesChanged:    `()`,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
		OnLightControlPanelStateChanged:   `()`,
	}
)

func GetByClassName(className string) (FGBuildableLightsControlPanel, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableLightsControlPanel{}, fmt.Errorf("failed to find FGBuildableLightsControlPanel with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableLightsControlPanel{
	"Build_LightsControlPanel_C": LightsControlPanel,
}
