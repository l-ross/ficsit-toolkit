// Code generated by ../../gen/docs_json. DO NOT EDIT.

package Buildable

import (
	"fmt"
)

type FGBuildable struct {
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
	MOccupiedText                      string
	MShouldModifyWorldGrid             bool
	MShouldShowAttachmentPointVisuals  bool
	MShouldShowHighlight               bool
	MSkipBuildEffect                   bool
	MToggleDormancyOnInteraction       bool
	MaxRenderDistance                  float64
	Tier                               int
}

var (
	ConveyorPoleWall = FGBuildable{
		Name:                               "ConveyorPoleWall",
		ClassName:                          "Build_ConveyorPoleWall_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Can be attached to walls and is used as a connection for conveyor belts.
Useful to route conveyor belts in a more controlled manner and over long distances.`,
		MDisplayName:                      `Conveyor Wall Mount`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	HyperTubeWallHole = FGBuildable{
		Name:                               "HyperTubeWallHole",
		ClassName:                          "Build_HyperTubeWallHole_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription:                       `Can be attached to walls to allow Hypertubes to pass through them.`,
		MDisplayName:                       `Hypertube Wall Hole`,
		MForceNetUpdateOnRegisterPlayer:    false,
		MHideOnBuildEffectStart:            false,
		MHighlightVector:                   `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                ``,
		MIsUseable:                         false,
		MShouldModifyWorldGrid:             true,
		MShouldShowAttachmentPointVisuals:  false,
		MShouldShowHighlight:               false,
		MSkipBuildEffect:                   false,
		MToggleDormancyOnInteraction:       false,
		MaxRenderDistance:                  -1.000000,
	}

	HyperTubeWallSupport = FGBuildable{
		Name:                               "HyperTubeWallSupport",
		ClassName:                          "Build_HyperTubeWallSupport_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Can be attached to walls. 
Supports for Hypertubes to allow for longer distances.`,
		MDisplayName:                      `Hypertube Wall Support`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	PipelineSupportWall = FGBuildable{
		Name:                               "PipelineSupportWall",
		ClassName:                          "Build_PipelineSupportWall_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Can be attached to walls.
Used to connect Pipelines over longer distances.`,
		MDisplayName:                      `Pipeline Wall Support`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        false,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	PipelineSupportWallHole = FGBuildable{
		Name:                               "PipelineSupportWallHole",
		ClassName:                          "Build_PipelineSupportWallHole_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription:                       `Can be attached to walls, allowing Pipelines to pass through them.`,
		MDisplayName:                       `Pipeline Wall Hole`,
		MForceNetUpdateOnRegisterPlayer:    false,
		MHideOnBuildEffectStart:            false,
		MHighlightVector:                   `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                ``,
		MIsUseable:                         false,
		MShouldModifyWorldGrid:             true,
		MShouldShowAttachmentPointVisuals:  false,
		MShouldShowHighlight:               false,
		MSkipBuildEffect:                   false,
		MToggleDormancyOnInteraction:       false,
		MaxRenderDistance:                  -1.000000,
	}

	WorkBench = FGBuildable{
		Name:                               "WorkBench",
		ClassName:                          "Build_WorkBench_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Allows you to manually craft a large range of different parts. 
These parts can then be used in construction of different factory buildings, vehicles and equipment.`,
		MDisplayName:                      `Craft Bench`,
		MForceNetUpdateOnRegisterPlayer:   true,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsUseable:                        true,
		MOccupiedText:                     `Craftbench is occupied!`,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MaxRenderDistance:                 -1.000000,
	}

	Workshop = FGBuildable{
		Name:                               "Workshop",
		ClassName:                          "Build_Workshop_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription:                       `Used to manually craft equipment.`,
		MDisplayName:                       `Equipment Workshop`,
		MForceNetUpdateOnRegisterPlayer:    false,
		MHideOnBuildEffectStart:            false,
		MHighlightVector:                   `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                ``,
		MIsUseable:                         true,
		MOccupiedText:                      `Equipment Workshop is occupied!`,
		MShouldModifyWorldGrid:             true,
		MShouldShowAttachmentPointVisuals:  false,
		MShouldShowHighlight:               false,
		MSkipBuildEffect:                   false,
		MToggleDormancyOnInteraction:       false,
		MaxRenderDistance:                  -1.000000,
	}

	WreathDecor = FGBuildable{
		Name:                               "WreathDecor",
		ClassName:                          "Build_WreathDecor_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription:                       `Can be attached to walls. Pretty.`,
		MDisplayName:                       `FICSMAS Wreath`,
		MForceNetUpdateOnRegisterPlayer:    false,
		MHideOnBuildEffectStart:            false,
		MHighlightVector:                   `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                ``,
		MIsUseable:                         false,
		MShouldModifyWorldGrid:             true,
		MShouldShowAttachmentPointVisuals:  false,
		MShouldShowHighlight:               false,
		MSkipBuildEffect:                   false,
		MToggleDormancyOnInteraction:       false,
		MaxRenderDistance:                  -1.000000,
	}

	XmassTree = FGBuildable{
		Name:                               "XmassTree",
		ClassName:                          "Build_XmassTree_C",
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription:                       `This special delivery gigantic FICSMAS Tree is decorated by progressing the FICSMAS Holiday Event MAM Tree.`,
		MDisplayName:                       `Giant FICSMAS Tree`,
		MForceNetUpdateOnRegisterPlayer:    false,
		MHideOnBuildEffectStart:            false,
		MHighlightVector:                   `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                ``,
		MIsUseable:                         false,
		MShouldModifyWorldGrid:             true,
		MShouldShowAttachmentPointVisuals:  false,
		MShouldShowHighlight:               false,
		MSkipBuildEffect:                   false,
		MToggleDormancyOnInteraction:       false,
		MaxRenderDistance:                  -1.000000,
		Tier:                               0,
	}
)

func GetByClassName(className string) (FGBuildable, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildable{}, fmt.Errorf("failed to find FGBuildable with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildable{
	"Build_ConveyorPoleWall_C":        ConveyorPoleWall,
	"Build_HyperTubeWallHole_C":       HyperTubeWallHole,
	"Build_HyperTubeWallSupport_C":    HyperTubeWallSupport,
	"Build_PipelineSupportWall_C":     PipelineSupportWall,
	"Build_PipelineSupportWallHole_C": PipelineSupportWallHole,
	"Build_WorkBench_C":               WorkBench,
	"Build_Workshop_C":                Workshop,
	"Build_WreathDecor_C":             WreathDecor,
	"Build_XmassTree_C":               XmassTree,
}
