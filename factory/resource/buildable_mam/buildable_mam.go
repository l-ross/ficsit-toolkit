// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableMAM

import (
	"fmt"
)

type FGBuildableMAM struct {
	Name                                    string
	ClassName                               string
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCurrentResearchState                   string
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
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
}

var (
	Mam = FGBuildableMAM{
		Name:                                    "Mam",
		ClassName:                               "Build_Mam_C",
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCurrentResearchState:                   `ERS_NotResearching`,
		MDescription: `The Molecular Analysis Machine is used to analyse new and exotic materials found on alien planets.
R&D will assist Pioneers through the MAM to turn any valuable data into usable research options and new technologies.`,
		MDisplayName:                    `MAM`,
		MExcludeFromMaterialInstancing:  ``,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MHideOnBuildEffectStart:         false,
		MHighlightVector:                `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:             ``,
		MIsUseable:                      true,
		MMaterialNameToInstanceManager:  `()`,
		MOccupiedText:                   `The Molecular Analysis Machine is occupied!`,
		MShouldModifyWorldGrid:          true,
		MShouldShowHighlight:            false,
		MSignificanceRange:              5000.000000,
		MSkipBuildEffect:                false,
		MToggleDormancyOnInteraction:    false,
		MaxRenderDistance:               -1.000000,
	}
)

func GetByClassName(className string) (*FGBuildableMAM, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableMAM with class name %s", className)
}

var classNameToVar = map[string]*FGBuildableMAM{
	"Build_Mam_C": &Mam,
}
