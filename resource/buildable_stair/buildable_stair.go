// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableStair

import (
	"fmt"
)

type FGBuildableStair struct {
	Name                               string
	ClassName                          string
	MAllowColoring                     bool
	MAttachmentPoints                  string
	MBuildEffectSpeed                  float64
	MCreateClearanceMeshRepresentation bool
	MDescription                       string
	MDisplayName                       string
	MForceNetUpdateOnRegisterPlayer    bool
	MHeight                            float64
	MHideOnBuildEffectStart            bool
	MHighlightVector                   string
	MInteractingPlayers                string
	MIsUseable                         bool
	MShouldModifyWorldGrid             bool
	MShouldShowAttachmentPointVisuals  bool
	MShouldShowHighlight               bool
	MSize                              float64
	MSkipBuildEffect                   bool
	MStairDirection                    string
	MToggleDormancyOnInteraction       bool
	MaxRenderDistance                  float64
}

var (
	StairsLeft01 = FGBuildableStair{
		Name:                               "StairsLeft01",
		ClassName:                          "Build_Stairs_Left_01_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations.
Makes it easier to get to other floors of your structures.`,
		MDisplayName:                      `Stairs Left`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHeight:                           200.000000,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             800.000000,
		MSkipBuildEffect:                  false,
		MStairDirection:                   `EBSD_Left`,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	StairsRight01 = FGBuildableStair{
		Name:                               "StairsRight01",
		ClassName:                          "Build_Stairs_Right_01_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations.
Makes it easier to get to other floors of your structures.`,
		MDisplayName:                      `Stairs Right`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHeight:                           200.000000,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             800.000000,
		MSkipBuildEffect:                  false,
		MStairDirection:                   `EBSD_Right`,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}
)

func GetByClassName(className string) (FGBuildableStair, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableStair{}, fmt.Errorf("failed to find FGBuildableStair with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableStair{
	"Build_Stairs_Left_01_C":  StairsLeft01,
	"Build_Stairs_Right_01_C": StairsRight01,
}
