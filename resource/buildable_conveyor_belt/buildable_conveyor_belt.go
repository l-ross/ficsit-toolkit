package BuildableConveyorBelt

import (
	"fmt"
)

type FGBuildableConveyorBelt struct {
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
	MItemMeshMap                            string
	MItems                                  string
	MMaterialNameToInstanceManager          string
	MMeshLength                             float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MSpeed                                  float64
	MSplineData                             string
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
}

var (
	ConveyorBeltMk1 = FGBuildableConveyorBelt{
		ClassName:                               "Build_ConveyorBeltMk1_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 60 resources per minute. Used to move resources between buildings.`,
		MDisplayName:                            `Conveyor Belt Mk.1`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 true,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshLength:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  120.000000,
		MSplineData:                             ``,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
	}

	ConveyorBeltMk2 = FGBuildableConveyorBelt{
		ClassName:                               "Build_ConveyorBeltMk2_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 120 resources per minute. Used to move resources between buildings.`,
		MDisplayName:                            `Conveyor Belt Mk.2`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 true,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshLength:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  240.000000,
		MSplineData:                             ``,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
	}

	ConveyorBeltMk3 = FGBuildableConveyorBelt{
		ClassName:                               "Build_ConveyorBeltMk3_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 270 resources per minute. Used to move resources between buildings.`,
		MDisplayName:                            `Conveyor Belt Mk.3`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 true,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshLength:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  540.000000,
		MSplineData:                             ``,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
	}

	ConveyorBeltMk4 = FGBuildableConveyorBelt{
		ClassName:                               "Build_ConveyorBeltMk4_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 480 resources per minute. Used to move resources between buildings.`,
		MDisplayName:                            `Conveyor Belt Mk.4`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 true,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshLength:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  960.000000,
		MSplineData:                             ``,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
	}

	ConveyorBeltMk5 = FGBuildableConveyorBelt{
		ClassName:                               "Build_ConveyorBeltMk5_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 780 resources per minute. Used to move resources between buildings.`,
		MDisplayName:                            `Conveyor Belt Mk.5`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 true,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshLength:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  1560.000000,
		MSplineData:                             ``,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableConveyorBelt, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableConveyorBelt with name %s", className)
}

var classNameToVar = map[string]*FGBuildableConveyorBelt{
	"Build_ConveyorBeltMk1_C": &ConveyorBeltMk1,
	"Build_ConveyorBeltMk2_C": &ConveyorBeltMk2,
	"Build_ConveyorBeltMk3_C": &ConveyorBeltMk3,
	"Build_ConveyorBeltMk4_C": &ConveyorBeltMk4,
	"Build_ConveyorBeltMk5_C": &ConveyorBeltMk5,
}
