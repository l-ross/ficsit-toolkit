// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableLadder

import (
	"fmt"
)

type FGBuildableLadder struct {
	Name                               string
	ClassName                          string
	MAllowColoring                     bool
	MAttachmentPoints                  string
	MBuildEffectSpeed                  float64
	MCreateClearanceMeshRepresentation bool
	MDescription                       string
	MDisplayName                       string
	MForceNetUpdateOnRegisterPlayer    bool
	MHideOnBuildEffectStart            bool
	MHighlightVector                   string
	MInteractingPlayers                string
	MIsUseable                         bool
	MLadderMeshes                      string
	MMaxSegmentCount                   int
	MMeshHeight                        float64
	MNumSegments                       int
	MShouldModifyWorldGrid             bool
	MShouldShowAttachmentPointVisuals  bool
	MShouldShowHighlight               bool
	MSkipBuildEffect                   bool
	MToggleDormancyOnInteraction       bool
	MWidth                             float64
	MaxRenderDistance                  float64
}

var (
	Ladder = FGBuildableLadder{
		Name:                               "Ladder",
		ClassName:                          "Build_Ladder_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription:                       `A ladder with a default height of 2 meters, which can be extended while building. Snaps to walls and foundations.`,
		MDisplayName:                       `Ladder`,
		MForceNetUpdateOnRegisterPlayer:    false,
		MHideOnBuildEffectStart:            false,
		MHighlightVector:                   `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                ``,
		MIsUseable:                         false,
		MLadderMeshes:                      ``,
		MMaxSegmentCount:                   10,
		MMeshHeight:                        200.000000,
		MNumSegments:                       0,
		MShouldModifyWorldGrid:             true,
		MShouldShowAttachmentPointVisuals:  false,
		MShouldShowHighlight:               false,
		MSkipBuildEffect:                   false,
		MToggleDormancyOnInteraction:       false,
		MWidth:                             80.000000,
		MaxRenderDistance:                  -1.000000,
	}
)

func GetByClassName(className string) (FGBuildableLadder, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableLadder{}, fmt.Errorf("failed to find FGBuildableLadder with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableLadder{
	"Build_Ladder_C": Ladder,
}
