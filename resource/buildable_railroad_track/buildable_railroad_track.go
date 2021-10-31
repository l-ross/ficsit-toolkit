package BuildableRailroadTrack

import (
	"fmt"
)

type FGBuildableRailroadTrack struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MConnections                            string
	MDescription                            string
	MDisplayName                            string
	MExcludeFromMaterialInstancing          string
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsOwnedByPlatform                      bool
	MIsUseable                              bool
	MMaterialNameToInstanceManager          string
	MMeshLength                             float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
}

var (
	RailroadTrack = FGBuildableRailroadTrack{
		ClassName:                               "Build_RailroadTrack_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MConnections:                            `(FGRailroadTrackConnectionComponent'"/Game/FactoryGame/Buildable/Factory/Train/Track/Build_RailroadTrack.Default__Build_RailroadTrack_C:TrackConnection0"',FGRailroadTrackConnectionComponent'"/Game/FactoryGame/Buildable/Factory/Train/Track/Build_RailroadTrack.Default__Build_RailroadTrack_C:TrackConnection1"')`,
		MDescription: `Used to transport trains in a reliable and fast manner.
Has a wide turn angle so make sure to plan it out properly.`,
		MDisplayName:                    `Railway`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsOwnedByPlatform:              false,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MMeshLength:                     1200.000000,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableRailroadTrack, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableRailroadTrack with name %s", className)
}

var classNameToVar = map[string]*FGBuildableRailroadTrack{
	"Build_RailroadTrack_C": &RailroadTrack,
}
