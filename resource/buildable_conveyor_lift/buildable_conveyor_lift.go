package BuildableConveyorLift

import (
	"fmt"
)

type FGBuildableConveyorLift struct {
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
	MIsReversed                             bool
	MIsUseable                              bool
	MItemMeshMap                            string
	MItems                                  string
	MMaterialNameToInstanceManager          string
	MMeshHeight                             float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSkipBuildEffect                        bool
	MSpeed                                  float64
	MToggleDormancyOnInteraction            bool
	MTopTransform                           string
	MaxRenderDistance                       float64
}

var (
	ConveyorLiftMk1 = FGBuildableConveyorLift{
		ClassName:                               "Build_ConveyorLiftMk1_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 60 resources per minute. Used to move resources between floors.`,
		MDisplayName:                            `Conveyor Lift Mk.1`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsReversed:                             false,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshHeight:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  120.000000,
		MToggleDormancyOnInteraction:            false,
		MTopTransform:                           `(Rotation=(X=0.000000,Y=0.000000,Z=0.000000,W=1.000000),Translation=(X=0.000000,Y=0.000000,Z=0.000000),Scale3D=(X=1.000000,Y=1.000000,Z=1.000000))`,
		MaxRenderDistance:                       -1.000000,
	}

	ConveyorLiftMk2 = FGBuildableConveyorLift{
		ClassName:                               "Build_ConveyorLiftMk2_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 120 resources per minute. Used to move resources between floors.`,
		MDisplayName:                            `Conveyor Lift Mk.2`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsReversed:                             false,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshHeight:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  240.000000,
		MToggleDormancyOnInteraction:            false,
		MTopTransform:                           `(Rotation=(X=0.000000,Y=0.000000,Z=0.000000,W=1.000000),Translation=(X=0.000000,Y=0.000000,Z=0.000000),Scale3D=(X=1.000000,Y=1.000000,Z=1.000000))`,
		MaxRenderDistance:                       -1.000000,
	}

	ConveyorLiftMk3 = FGBuildableConveyorLift{
		ClassName:                               "Build_ConveyorLiftMk3_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 270 resources per minute. Used to move resources between floors.`,
		MDisplayName:                            `Conveyor Lift Mk.3`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsReversed:                             false,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshHeight:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  540.000000,
		MToggleDormancyOnInteraction:            false,
		MTopTransform:                           `(Rotation=(X=0.000000,Y=0.000000,Z=0.000000,W=1.000000),Translation=(X=0.000000,Y=0.000000,Z=0.000000),Scale3D=(X=1.000000,Y=1.000000,Z=1.000000))`,
		MaxRenderDistance:                       -1.000000,
	}

	ConveyorLiftMk4 = FGBuildableConveyorLift{
		ClassName:                               "Build_ConveyorLiftMk4_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 480 resources per minute. Used to move resources between floors.`,
		MDisplayName:                            `Conveyor Lift Mk.4`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsReversed:                             false,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshHeight:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  960.000000,
		MToggleDormancyOnInteraction:            false,
		MTopTransform:                           `(Rotation=(X=0.000000,Y=0.000000,Z=0.000000,W=1.000000),Translation=(X=0.000000,Y=0.000000,Z=0.000000),Scale3D=(X=1.000000,Y=1.000000,Z=1.000000))`,
		MaxRenderDistance:                       -1.000000,
	}

	ConveyorLiftMk5 = FGBuildableConveyorLift{
		ClassName:                               "Build_ConveyorLiftMk5_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MDescription:                            `Transports up to 780 resources per minute. Used to move resources between floors.`,
		MDisplayName:                            `Conveyor Lift Mk.5`,
		MExcludeFromMaterialInstancing:          ``,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsReversed:                             false,
		MIsUseable:                              false,
		MItemMeshMap:                            `()`,
		MItems:                                  `()`,
		MMaterialNameToInstanceManager:          `()`,
		MMeshHeight:                             200.000000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSkipBuildEffect:                        false,
		MSpeed:                                  1560.000000,
		MToggleDormancyOnInteraction:            false,
		MTopTransform:                           `(Rotation=(X=0.000000,Y=0.000000,Z=0.000000,W=1.000000),Translation=(X=0.000000,Y=0.000000,Z=0.000000),Scale3D=(X=1.000000,Y=1.000000,Z=1.000000))`,
		MaxRenderDistance:                       -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableConveyorLift, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableConveyorLift with name %s", className)
}

var classNameToVar = map[string]*FGBuildableConveyorLift{
	"Build_ConveyorLiftMk1_C": &ConveyorLiftMk1,
	"Build_ConveyorLiftMk2_C": &ConveyorLiftMk2,
	"Build_ConveyorLiftMk3_C": &ConveyorLiftMk3,
	"Build_ConveyorLiftMk4_C": &ConveyorLiftMk4,
	"Build_ConveyorLiftMk5_C": &ConveyorLiftMk5,
}
