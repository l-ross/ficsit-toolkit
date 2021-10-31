package BuildableWall

import (
	"fmt"
)

type FGBuildableWall struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHeight                                 float64
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MMaterialNameToInstanceManager          string
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MWidth                                  float64
	MaxRenderDistance                       float64
}

var (
	Fence01 = FGBuildableWall{
		ClassName:                               "Build_Fence_01_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `The Fence can be built on the edges of floors to prevent you from falling off.`,
		MDisplayName:                            `Fence`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHeight:                                 250.000000,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MMaterialNameToInstanceManager:          `()`,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MWidth:                                  800.000000,
		MaxRenderDistance:                       -1.000000,
	}

	Wall8x401 = FGBuildableWall{
		ClassName:                               "Build_Wall_8x4_01_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
Useful for building multi-floor structures.`,
		MDisplayName:                    `Wall 8m x 4m`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	Wall8x402 = FGBuildableWall{
		ClassName:                               "Build_Wall_8x4_02_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
Useful for building multi-floor structures.`,
		MDisplayName:                    `Wall 8m x 4m`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallConveyor8x401 = FGBuildableWall{
		ClassName:                               "Build_Wall_Conveyor_8x4_01_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Walls connect to other Walls and Floors. Use these to make buildings with several floors.
Has 3 conveyor belt connections.
Height: 4 m`,
		MDisplayName:                    `Wall Conveyor x3`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallConveyor8x401Steel = FGBuildableWall{
		ClassName:                               "Build_Wall_Conveyor_8x4_01_Steel_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Walls connect to other Walls and Floors. Use these to make buildings with several floors.
Has 3 conveyor belt connections.
Height: 4 m`,
		MDisplayName:                    `Wall Conveyor x3`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallConveyor8x402 = FGBuildableWall{
		ClassName:                               "Build_Wall_Conveyor_8x4_02_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Walls connect to other Walls and Floors. Use these to make buildings with several floors.
Has 2 conveyor belt connections.
Height: 4 m`,
		MDisplayName:                    `Wall Conveyor x2`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallConveyor8x402Steel = FGBuildableWall{
		ClassName:                               "Build_Wall_Conveyor_8x4_02_Steel_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Walls connect to other Walls and Floors. Use these to make buildings with several floors.
Has 2 conveyor belt connections.
Height: 4 m`,
		MDisplayName:                    `Wall Conveyor x2`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallConveyor8x403 = FGBuildableWall{
		ClassName:                               "Build_Wall_Conveyor_8x4_03_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Walls connect to other Walls and Floors. Use these to make buildings with several floors.
Has 1 conveyor belt connection.
Height: 4 m`,
		MDisplayName:                    `Wall Conveyor x1`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallConveyor8x403Steel = FGBuildableWall{
		ClassName:                               "Build_Wall_Conveyor_8x4_03_Steel_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Walls connect to other Walls and Floors. Use these to make buildings with several floors.
Has 1 conveyor belt connection.
Height: 4 m`,
		MDisplayName:                    `Wall Conveyor x1`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallConveyor8x404 = FGBuildableWall{
		ClassName:                               "Build_Wall_Conveyor_8x4_04_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Walls connect to other Walls and Floors. Use these to make buildings with several floors.
Has 1 conveyor belt connection perpendicular to the wall.
Height: 4 m`,
		MDisplayName:                    `Wall Conveyor Perpendicular`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallConveyor8x404Steel = FGBuildableWall{
		ClassName:                               "Build_Wall_Conveyor_8x4_04_Steel_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Walls connect to other Walls and Floors. Use these to make buildings with several floors.
Has 1 conveyor belt connection perpendicular to the wall.
Height: 4 m`,
		MDisplayName:                    `Wall Conveyor Perpendicular`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallDoor8x401 = FGBuildableWall{
		ClassName:                               "Build_Wall_Door_8x4_01_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The door allows Pioneers to pass through the wall.`,
		MDisplayName:                    `Center Door Wall`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallDoor8x401Steel = FGBuildableWall{
		ClassName:                               "Build_Wall_Door_8x4_01_Steel_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The door allows Pioneers to pass through the wall.`,
		MDisplayName:                    `Center Door Wall`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallDoor8x402 = FGBuildableWall{
		ClassName:                               "Build_Wall_Door_8x4_02_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The door allows Pioneers to pass through the wall.`,
		MDisplayName:                    `Left Door Wall`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallDoor8x402Steel = FGBuildableWall{
		ClassName:                               "Build_Wall_Door_8x4_02_Steel_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The door allows Pioneers to pass through the wall.`,
		MDisplayName:                    `Left Door Wall`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallDoor8x403 = FGBuildableWall{
		ClassName:                               "Build_Wall_Door_8x4_03_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The door allows Pioneers to pass through the wall.`,
		MDisplayName:                    `Right Door Wall`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallDoor8x403Steel = FGBuildableWall{
		ClassName:                               "Build_Wall_Door_8x4_03_Steel_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The door allows Pioneers to pass through the wall.`,
		MDisplayName:                    `Right Door Wall`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallGate8x401 = FGBuildableWall{
		ClassName:                               "Build_Wall_Gate_8x4_01_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The gate allows Pioneers to pass through the wall.`,
		MDisplayName:                    `Gate Wall`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallWindow8x401 = FGBuildableWall{
		ClassName:                               "Build_Wall_Window_8x4_01_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The windows allow Pioneers to see through the wall.`,
		MDisplayName:                    `Single Window`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallWindow8x402 = FGBuildableWall{
		ClassName:                               "Build_Wall_Window_8x4_02_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The windows allow Pioneers to see through the wall.`,
		MDisplayName:                    `Frame Window`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallWindow8x403 = FGBuildableWall{
		ClassName:                               "Build_Wall_Window_8x4_03_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The windows allow Pioneers to see through the wall.`,
		MDisplayName:                    `Panel Window`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}

	WallWindow8x404 = FGBuildableWall{
		ClassName:                               "Build_Wall_Window_8x4_04_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walls.
The windows allow Pioneers to see through the wall.`,
		MDisplayName:                    `Reinforced Window`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         400.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MWidth:                          800.000000,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableWall, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableWall with name %s", className)
}

var classNameToVar = map[string]*FGBuildableWall{
	"Build_Fence_01_C":                   &Fence01,
	"Build_Wall_8x4_01_C":                &Wall8x401,
	"Build_Wall_8x4_02_C":                &Wall8x402,
	"Build_Wall_Conveyor_8x4_01_C":       &WallConveyor8x401,
	"Build_Wall_Conveyor_8x4_01_Steel_C": &WallConveyor8x401Steel,
	"Build_Wall_Conveyor_8x4_02_C":       &WallConveyor8x402,
	"Build_Wall_Conveyor_8x4_02_Steel_C": &WallConveyor8x402Steel,
	"Build_Wall_Conveyor_8x4_03_C":       &WallConveyor8x403,
	"Build_Wall_Conveyor_8x4_03_Steel_C": &WallConveyor8x403Steel,
	"Build_Wall_Conveyor_8x4_04_C":       &WallConveyor8x404,
	"Build_Wall_Conveyor_8x4_04_Steel_C": &WallConveyor8x404Steel,
	"Build_Wall_Door_8x4_01_C":           &WallDoor8x401,
	"Build_Wall_Door_8x4_01_Steel_C":     &WallDoor8x401Steel,
	"Build_Wall_Door_8x4_02_C":           &WallDoor8x402,
	"Build_Wall_Door_8x4_02_Steel_C":     &WallDoor8x402Steel,
	"Build_Wall_Door_8x4_03_C":           &WallDoor8x403,
	"Build_Wall_Door_8x4_03_Steel_C":     &WallDoor8x403Steel,
	"Build_Wall_Gate_8x4_01_C":           &WallGate8x401,
	"Build_Wall_Window_8x4_01_C":         &WallWindow8x401,
	"Build_Wall_Window_8x4_02_C":         &WallWindow8x402,
	"Build_Wall_Window_8x4_03_C":         &WallWindow8x403,
	"Build_Wall_Window_8x4_04_C":         &WallWindow8x404,
}