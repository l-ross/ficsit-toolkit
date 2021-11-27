// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableDoor

import (
	"fmt"
)

type FGBuildableDoor struct {
	Name                               string
	ClassName                          string
	BigOverlapList                     string
	BlendExp                           float64
	EasingFunction                     string
	IsDoorOpen                         bool
	MAllowColoring                     bool
	MAngledVariants                    string
	MAngularDepth                      float64
	MAnimationRate                     float64
	MAttachmentPoints                  string
	MBuildEffectSpeed                  float64
	MCanBeLocked                       bool
	MCreateClearanceMeshRepresentation bool
	MDescription                       string
	MDisplayName                       string
	MElevation                         float64
	MForceNetUpdateOnRegisterPlayer    bool
	MHeight                            float64
	MHideOnBuildEffectStart            bool
	MHighlightVector                   string
	MInteractingPlayers                string
	MIsUseable                         bool
	MMovementRate                      float64
	MShouldModifyWorldGrid             bool
	MShouldShowAttachmentPointVisuals  bool
	MShouldShowHighlight               bool
	MSkipBuildEffect                   bool
	MToggleDormancyOnInteraction       bool
	MWallType                          string
	MWidth                             float64
	MaxRenderDistance                  float64
	Steps                              int
}

var (
	GateAutomated8x4 = FGBuildableDoor{
		Name:                               "GateAutomated8x4",
		ClassName:                          "Build_Gate_Automated_8x4_C",
		BigOverlapList:                     ``,
		BlendExp:                           10.000000,
		EasingFunction:                     `Linear`,
		MAllowColoring:                     true,
		MAngledVariants:                    ``,
		MAngularDepth:                      0.000000,
		MAnimationRate:                     0.016670,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCanBeLocked:                       false,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Automatically opens when living beings or vehicles approach it.
Gate settings can be configured.
Snaps to foundations and other walls.

Size: 8m x 4m`,
		MDisplayName:                      `Automated Gate`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHeight:                           400.000000,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MMovementRate:                     125.000000,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MWallType:                         `BWT_Normal`,
		MWidth:                            800.000000,
		MaxRenderDistance:                 -1.000000,
		Steps:                             2,
	}

	WallConcreteCDoor8x4 = FGBuildableDoor{
		Name:                               "WallConcreteCDoor8x4",
		ClassName:                          "Build_Wall_Concrete_CDoor_8x4_C",
		BlendExp:                           5.000000,
		EasingFunction:                     `ExpoInOut`,
		IsDoorOpen:                         false,
		MAllowColoring:                     true,
		MAngledVariants:                    ``,
		MAngularDepth:                      0.000000,
		MAnimationRate:                     0.016667,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCanBeLocked:                       true,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `The door allows Pioneers to pass through the wall.
Door settings can be configured.
Snaps to Foundations and other Walls.

Size: 8m x 4m`,
		MDisplayName:                      `Center Door Wall`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHeight:                           400.000000,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MMovementRate:                     200.000000,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MWallType:                         `BWT_Normal`,
		MWidth:                            800.000000,
		MaxRenderDistance:                 -1.000000,
		Steps:                             2,
	}

	WallConcreteSDoor8x4 = FGBuildableDoor{
		Name:                               "WallConcreteSDoor8x4",
		ClassName:                          "Build_Wall_Concrete_SDoor_8x4_C",
		BlendExp:                           5.000000,
		EasingFunction:                     `ExpoInOut`,
		IsDoorOpen:                         false,
		MAllowColoring:                     true,
		MAngledVariants:                    ``,
		MAngularDepth:                      0.000000,
		MAnimationRate:                     0.016667,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCanBeLocked:                       true,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `The door allows Pioneers to pass through the wall.
Door settings can be configured.
Snaps to Foundations and other Walls.

Size: 8m x 4m`,
		MDisplayName:                      `Side Door Wall`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHeight:                           400.000000,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MMovementRate:                     200.000000,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MWallType:                         `BWT_Normal`,
		MWidth:                            800.000000,
		MaxRenderDistance:                 -1.000000,
		Steps:                             2,
	}

	WallDoor8x401 = FGBuildableDoor{
		Name:                               "WallDoor8x401",
		ClassName:                          "Build_Wall_Door_8x4_01_C",
		BlendExp:                           5.000000,
		EasingFunction:                     `ExpoInOut`,
		IsDoorOpen:                         false,
		MAllowColoring:                     true,
		MAngledVariants:                    ``,
		MAngularDepth:                      0.000000,
		MAnimationRate:                     0.016667,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCanBeLocked:                       true,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `The door allows Pioneers to pass through the wall.
Door settings can be configured.
Snaps to Foundations and other Walls.

Size: 8m x 4m`,
		MDisplayName:                      `Center Door Wall`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHeight:                           400.000000,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MMovementRate:                     200.000000,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MWallType:                         `BWT_Normal`,
		MWidth:                            800.000000,
		MaxRenderDistance:                 -1.000000,
		Steps:                             2,
	}

	WallDoor8x401Steel = FGBuildableDoor{
		Name:                               "WallDoor8x401Steel",
		ClassName:                          "Build_Wall_Door_8x4_01_Steel_C",
		BlendExp:                           5.000000,
		EasingFunction:                     `ExpoInOut`,
		IsDoorOpen:                         false,
		MAllowColoring:                     true,
		MAngledVariants:                    ``,
		MAngularDepth:                      0.000000,
		MAnimationRate:                     0.016667,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCanBeLocked:                       true,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `The door allows Pioneers to pass through the wall.
Door settings can be configured.
Snaps to Foundations and other Walls.

Size: 8m x 4m`,
		MDisplayName:                      `Center Door Wall`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHeight:                           400.000000,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MMovementRate:                     200.000000,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MWallType:                         `BWT_Normal`,
		MWidth:                            800.000000,
		MaxRenderDistance:                 -1.000000,
		Steps:                             2,
	}

	WallDoor8x403 = FGBuildableDoor{
		Name:                               "WallDoor8x403",
		ClassName:                          "Build_Wall_Door_8x4_03_C",
		BlendExp:                           5.000000,
		EasingFunction:                     `ExpoInOut`,
		IsDoorOpen:                         false,
		MAllowColoring:                     true,
		MAngledVariants:                    ``,
		MAngularDepth:                      0.000000,
		MAnimationRate:                     0.016667,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCanBeLocked:                       true,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `The door allows Pioneers to pass through the wall.
Door settings can be configured.
Snaps to Foundations and other Walls.

Size: 8m x 4m`,
		MDisplayName:                      `Side Door Wall`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHeight:                           400.000000,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MMovementRate:                     200.000000,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MWallType:                         `BWT_Normal`,
		MWidth:                            800.000000,
		MaxRenderDistance:                 -1.000000,
		Steps:                             2,
	}

	WallDoor8x403Steel = FGBuildableDoor{
		Name:                               "WallDoor8x403Steel",
		ClassName:                          "Build_Wall_Door_8x4_03_Steel_C",
		BlendExp:                           5.000000,
		EasingFunction:                     `ExpoInOut`,
		IsDoorOpen:                         false,
		MAllowColoring:                     true,
		MAngledVariants:                    ``,
		MAngularDepth:                      0.000000,
		MAnimationRate:                     0.016667,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCanBeLocked:                       true,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `The door allows Pioneers to pass through the wall.
Door settings can be configured.
Snaps to Foundations and other Walls.

Size: 8m x 4m`,
		MDisplayName:                      `Side Door Wall`,
		MElevation:                        0.000000,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHeight:                           400.000000,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MMovementRate:                     200.000000,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MWallType:                         `BWT_Normal`,
		MWidth:                            800.000000,
		MaxRenderDistance:                 -1.000000,
		Steps:                             2,
	}
)

func GetByClassName(className string) (FGBuildableDoor, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableDoor{}, fmt.Errorf("failed to find FGBuildableDoor with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableDoor{
	"Build_Gate_Automated_8x4_C":      GateAutomated8x4,
	"Build_Wall_Concrete_CDoor_8x4_C": WallConcreteCDoor8x4,
	"Build_Wall_Concrete_SDoor_8x4_C": WallConcreteSDoor8x4,
	"Build_Wall_Door_8x4_01_C":        WallDoor8x401,
	"Build_Wall_Door_8x4_01_Steel_C":  WallDoor8x401Steel,
	"Build_Wall_Door_8x4_03_C":        WallDoor8x403,
	"Build_Wall_Door_8x4_03_Steel_C":  WallDoor8x403Steel,
}
