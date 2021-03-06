// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableWalkway

import (
	"fmt"
)

type FGBuildableWalkway struct {
	Name                               string
	ClassName                          string
	MAllowColoring                     bool
	MAttachmentPoints                  string
	MBuildEffectSpeed                  float64
	MCreateClearanceMeshRepresentation bool
	MDescription                       string
	MDisableSnapOn                     string
	MDisplayName                       string
	MElevation                         float64
	MForceNetUpdateOnRegisterPlayer    bool
	MHideOnBuildEffectStart            bool
	MHighlightVector                   string
	MInteractingPlayers                string
	MIsUseable                         bool
	MShouldModifyWorldGrid             bool
	MShouldShowAttachmentPointVisuals  bool
	MShouldShowHighlight               bool
	MSize                              float64
	MSkipBuildEffect                   bool
	MToggleDormancyOnInteraction       bool
	MaxRenderDistance                  float64
}

var (
	CatwalkCorner = FGBuildableWalkway{
		Name:                               "CatwalkCorner",
		ClassName:                          "Build_CatwalkCorner_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Front=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                      `Catwalk Corner`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	CatwalkCross = FGBuildableWalkway{
		Name:                               "CatwalkCross",
		ClassName:                          "Build_CatwalkCross_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Top=True,Bottom=True)`,
		MDisplayName:                      `Catwalk Crossing`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	CatwalkRamp = FGBuildableWalkway{
		Name:                               "CatwalkRamp",
		ClassName:                          "Build_CatwalkRamp_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Right=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                      `Catwalk Ramp`,
		MElevation:                        200.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	CatwalkStairs = FGBuildableWalkway{
		Name:                               "CatwalkStairs",
		ClassName:                          "Build_CatwalkStairs_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Right=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                      `Catwalk Stairs`,
		MElevation:                        400.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	CatwalkStraight = FGBuildableWalkway{
		Name:                               "CatwalkStraight",
		ClassName:                          "Build_CatwalkStraight_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Right=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                      `Catwalk Straight`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	CatwalkT = FGBuildableWalkway{
		Name:                               "CatwalkT",
		ClassName:                          "Build_CatwalkT_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Front=True,Top=True,Bottom=True)`,
		MDisplayName:                      `Catwalk T-Crossing`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	WalkwayCross = FGBuildableWalkway{
		Name:                               "WalkwayCross",
		ClassName:                          "Build_WalkwayCross_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Top=True,Bottom=True)`,
		MDisplayName:                      `Walkway Crossing`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	WalkwayRamp = FGBuildableWalkway{
		Name:                               "WalkwayRamp",
		ClassName:                          "Build_WalkwayRamp_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Right=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                      `Walkway Ramp`,
		MElevation:                        200.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	WalkwayStraight = FGBuildableWalkway{
		Name:                               "WalkwayStraight",
		ClassName:                          "Build_WalkwayStraight_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Right=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                      `Walkway Straight`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	WalkwayT = FGBuildableWalkway{
		Name:                               "WalkwayT",
		ClassName:                          "Build_WalkwayT_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Front=True,Top=True,Bottom=True)`,
		MDisplayName:                      `Walkway T-Crossing`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	WalkwayTrun = FGBuildableWalkway{
		Name:                               "WalkwayTrun",
		ClassName:                          "Build_WalkwayTrun_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                    `(Front=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                      `Walkway Turn`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSize:                             400.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}
)

func GetByClassName(className string) (FGBuildableWalkway, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableWalkway{}, fmt.Errorf("failed to find FGBuildableWalkway with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableWalkway{
	"Build_CatwalkCorner_C":   CatwalkCorner,
	"Build_CatwalkCross_C":    CatwalkCross,
	"Build_CatwalkRamp_C":     CatwalkRamp,
	"Build_CatwalkStairs_C":   CatwalkStairs,
	"Build_CatwalkStraight_C": CatwalkStraight,
	"Build_CatwalkT_C":        CatwalkT,
	"Build_WalkwayCross_C":    WalkwayCross,
	"Build_WalkwayRamp_C":     WalkwayRamp,
	"Build_WalkwayStraight_C": WalkwayStraight,
	"Build_WalkwayT_C":        WalkwayT,
	"Build_WalkwayTrun_C":     WalkwayTrun,
}
