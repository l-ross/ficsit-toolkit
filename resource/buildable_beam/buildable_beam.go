// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableBeam

import (
	"fmt"
)

type FGBuildableBeam struct {
	Name                               string
	ClassName                          string
	MAllowColoring                     bool
	MAttachmentPoints                  string
	MBuildEffectSpeed                  float64
	MCreateClearanceMeshRepresentation bool
	MDefaultLength                     float64
	MDescription                       string
	MDisplayName                       string
	MForceNetUpdateOnRegisterPlayer    bool
	MHideOnBuildEffectStart            bool
	MHighlightVector                   string
	MInteractingPlayers                string
	MIsUseable                         bool
	MLength                            float64
	MMaxLength                         float64
	MShouldModifyWorldGrid             bool
	MShouldShowAttachmentPointVisuals  bool
	MShouldShowHighlight               bool
	MSize                              float64
	MSkipBuildEffect                   bool
	MToggleDormancyOnInteraction       bool
	MaxRenderDistance                  float64
}

var (
	Beam = FGBuildableBeam{
		Name:                               "Beam",
		ClassName:                          "Build_Beam_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDefaultLength:                     400.000000,
		MDescription: `Snaps to other Beams and various other structural buildings.
Beams support multiple build modes for different use cases.`,
		MDisplayName:                      `Metal Beam`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MLength:                           0.000000,
		MMaxLength:                        4000.000000,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             100.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	BeamPainted = FGBuildableBeam{
		Name:                               "BeamPainted",
		ClassName:                          "Build_Beam_Painted_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDefaultLength:                     400.000000,
		MDescription: `Snaps to other Beams and various other structural buildings.
Beams support multiple build modes for different use cases.`,
		MDisplayName:                      `Painted Beam`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MLength:                           0.000000,
		MMaxLength:                        4000.000000,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             100.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}
)

func GetByClassName(className string) (FGBuildableBeam, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableBeam{}, fmt.Errorf("failed to find FGBuildableBeam with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableBeam{
	"Build_Beam_C":         Beam,
	"Build_Beam_Painted_C": BeamPainted,
}
