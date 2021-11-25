// Code generated by ../../gen/docs_json. DO NOT EDIT.

package Buildable

import (
	"fmt"
)

type FGBuildable struct {
	Name                                    string
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MMaterialNameToInstanceManager          string
	MOccupiedText                           string
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	Tier                                    int
}

var (
	ConveyorPoleWall = FGBuildable{
		Name:                                    "ConveyorPoleWall",
		ClassName:                               "Build_ConveyorPoleWall_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Can be attached to walls and is used as a connection for conveyor belts.
Useful to route conveyor belts in a more controlled manner and over long distances.`,
		MDisplayName:                    `Conveyor Wall Mount`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	HyperTubeWallHole = FGBuildable{
		Name:                                    "HyperTubeWallHole",
		ClassName:                               "Build_HyperTubeWallHole_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Can be attached to walls to allow Hyper Tubes to pass through them.`,
		MDisplayName:                            `Hyper Tube Wall Hole`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     `((Rotation=(X=0.000000,Y=-0.000000,Z=0.707107,W=0.707106),Translation=(X=-28.918512,Y=0.000427,Z=100.000000),Scale3D=(X=0.726597,Y=0.726597,Z=0.726597)),(Rotation=(X=-0.000000,Y=0.000000,Z=-0.707108,W=0.707105),Translation=(X=29.806301,Y=-0.000015,Z=100.000000),Scale3D=(X=0.726597,Y=0.726597,Z=0.726597)))`,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MMaterialNameToInstanceManager:          `()`,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
	}

	HyperTubeWallSupport = FGBuildable{
		Name:                                    "HyperTubeWallSupport",
		ClassName:                               "Build_HyperTubeWallSupport_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Can be attached to walls. 
Supports for Hyper Tubes to allow for longer distances.`,
		MDisplayName:                    `Hyper Tube Wall Support`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PipelineSupportWall = FGBuildable{
		Name:                                    "PipelineSupportWall",
		ClassName:                               "Build_PipelineSupportWall_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Can be attached to walls.
Used to connect Pipelines over longer distances.`,
		MDisplayName:                    `Pipeline Wall Support`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	PipelineSupportWallHole = FGBuildable{
		Name:                                    "PipelineSupportWallHole",
		ClassName:                               "Build_PipelineSupportWallHole_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Can be attached to walls, allowing Pipelines to pass through them.`,
		MDisplayName:                            `Pipeline Wall Hole`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MMaterialNameToInstanceManager:          `()`,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
	}

	WorkBench = FGBuildable{
		Name:                                    "WorkBench",
		ClassName:                               "Build_WorkBench_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Allows you to manually craft a large range of different parts. 
These parts can then be used in construction of different factory buildings, vehicles and equipment.`,
		MDisplayName:                    `Craft Bench`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: true,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      true,
		MMaterialNameToInstanceManager:  `()`,
		MOccupiedText:                   `Craftbench is occupied!`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	Workshop = FGBuildable{
		Name:                                    "Workshop",
		ClassName:                               "Build_Workshop_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Used to manually craft equipment.`,
		MDisplayName:                            `Equipment Workshop`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              true,
		MMaterialNameToInstanceManager:          `()`,
		MOccupiedText:                           `Equipment Workshop is occupied!`,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
	}

	WreathDecor = FGBuildable{
		Name:                                    "WreathDecor",
		ClassName:                               "Build_WreathDecor_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Can be attached to walls. Pretty.`,
		MDisplayName:                            `FICSMAS Wreath`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MMaterialNameToInstanceManager:          `()`,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
	}

	XmassTree = FGBuildable{
		Name:                                    "XmassTree",
		ClassName:                               "Build_XmassTree_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    true,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `This special delivery gigantic FICSMAS Tree is decorated by progressing the FICSMAS Holiday Event MAM Tree.`,
		MDisplayName:                            `Giant FICSMAS Tree`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MMaterialNameToInstanceManager:          `()`,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
		Tier:                                    0,
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
