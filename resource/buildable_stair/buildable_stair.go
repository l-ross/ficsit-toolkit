package BuildableStair

import (
	"fmt"
)

type FGBuildableStair struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MDescription                            string
	MDisableAttachmentSnapOn                string
	MDisableSnapOn                          string
	MDisplayName                            string
	MElevation                              float64
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
	MSize                                   float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
}

var (
	Stair1b = FGBuildableStair{
		ClassName:                               "Build_Stair_1b_C",
		MAllowCleranceSeparationEvenIfStackedOn: true,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `The Stairs allow you to move to the upper floors of your buildings.
Height: 4 m`,
		MDisableAttachmentSnapOn:        `()`,
		MDisableSnapOn:                  `()`,
		MDisplayName:                    `Stairs`,
		MElevation:                      0.000000,
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
		MSize:                           800.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	StairsLeft01 = FGBuildableStair{
		ClassName:                               "Build_Stairs_Left_01_C",
		MAllowCleranceSeparationEvenIfStackedOn: true,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations.
Makes it easier to get to other floors of your structures.`,
		MDisableAttachmentSnapOn:        `()`,
		MDisableSnapOn:                  `()`,
		MDisplayName:                    `Stairs Left`,
		MElevation:                      0.000000,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         200.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSize:                           800.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	StairsRight01 = FGBuildableStair{
		ClassName:                               "Build_Stairs_Right_01_C",
		MAllowCleranceSeparationEvenIfStackedOn: true,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations.
Makes it easier to get to other floors of your structures.`,
		MDisableAttachmentSnapOn:        `()`,
		MDisableSnapOn:                  `()`,
		MDisplayName:                    `Stairs Right`,
		MElevation:                      0.000000,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         200.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSize:                           800.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableStair, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableStair with name %s", className)
}

var classNameToVar = map[string]*FGBuildableStair{
	"Build_Stair_1b_C":        &Stair1b,
	"Build_Stairs_Left_01_C":  &StairsLeft01,
	"Build_Stairs_Right_01_C": &StairsRight01,
}