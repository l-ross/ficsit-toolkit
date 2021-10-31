package BuildableWalkway

import (
	"fmt"
)

type FGBuildableWalkway struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MDescription                            string
	MDisableSnapOn                          string
	MDisplayName                            string
	MElevation                              float64
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
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
	WalkwayCross = FGBuildableWalkway{
		ClassName:                               "Build_WalkwayCross_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                  `(Top=True,Bottom=True)`,
		MDisplayName:                    `Walkway Crossing`,
		MElevation:                      0.000000,
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
		MSize:                           400.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	WalkwayRamp = FGBuildableWalkway{
		ClassName:                               "Build_WalkwayRamp_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                  `(Right=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                    `Walkway Ramp`,
		MElevation:                      200.000000,
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
		MSize:                           400.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	WalkwayStraight = FGBuildableWalkway{
		ClassName:                               "Build_WalkwayStraight_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                  `(Right=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                    `Walkway Straight`,
		MElevation:                      0.000000,
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
		MSize:                           400.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	WalkwayT = FGBuildableWalkway{
		ClassName:                               "Build_WalkwayT_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                  `(Front=True,Top=True,Bottom=True)`,
		MDisplayName:                    `Walkway T-Crossing`,
		MElevation:                      0.000000,
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
		MSize:                           400.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}

	WalkwayTrun = FGBuildableWalkway{
		ClassName:                               "Build_WalkwayTrun_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription: `Snaps to foundations and other walkways.
Specifically made for humans to walk on.`,
		MDisableSnapOn:                  `(Front=True,Left=True,Top=True,Bottom=True)`,
		MDisplayName:                    `Walkway Turn`,
		MElevation:                      0.000000,
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
		MSize:                           400.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableWalkway, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableWalkway with name %s", className)
}

var classNameToVar = map[string]*FGBuildableWalkway{
	"Build_WalkwayCross_C":    &WalkwayCross,
	"Build_WalkwayRamp_C":     &WalkwayRamp,
	"Build_WalkwayStraight_C": &WalkwayStraight,
	"Build_WalkwayT_C":        &WalkwayT,
	"Build_WalkwayTrun_C":     &WalkwayTrun,
}
