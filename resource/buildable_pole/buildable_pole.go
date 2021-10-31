package BuildablePole

import (
	"fmt"
)

type FGBuildablePole struct {
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCanStack                               bool
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
	MStackHeight                            float64
	MToggleDormancyOnInteraction            bool
	MUseStaticHeight                        bool
	MaxRenderDistance                       float64
}

var (
	ConveyorPole = FGBuildablePole{
		ClassName:                               "Build_ConveyorPole_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCanStack:                               false,
		MDescription: `Can be used as a connection for conveyor belts. The height of the pole can be adjusted.
Useful to route conveyor belts in a more controlled manner and over long distances.`,
		MDisplayName:                    `Conveyor Pole`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHeight:                         100.000000,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      false,
		MMaterialNameToInstanceManager:  `()`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSkipBuildEffect:                false,
		MStackHeight:                    200.000000,
		MToggleDormancyOnInteraction:    false,
		MUseStaticHeight:                false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildablePole, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildablePole with name %s", className)
}

var classNameToVar = map[string]*FGBuildablePole{
	"Build_ConveyorPole_C": &ConveyorPole,
}