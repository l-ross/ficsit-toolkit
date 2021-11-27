// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableRailroadSignal

import (
	"fmt"
)

type FGBuildableRailroadSignal struct {
	Name                               string
	ClassName                          string
	MAllowColoring                     bool
	MAspect                            string
	MAttachmentPoints                  string
	MBlockValidation                   string
	MBuildEffectSpeed                  float64
	MCreateClearanceMeshRepresentation bool
	MDescription                       string
	MDisplayName                       string
	MForceNetUpdateOnRegisterPlayer    bool
	MGuardedConnections                string
	MHideOnBuildEffectStart            bool
	MHighlightVector                   string
	MInteractingPlayers                string
	MIsBiDirectional                   bool
	MIsPathSignal                      bool
	MIsUseable                         bool
	MObservedConnections               string
	MOnAspectChangedDelegate           string
	MOnBlockValidationChangedDelegate  string
	MPreviousAspect                    string
	MShouldModifyWorldGrid             bool
	MShouldShowAttachmentPointVisuals  bool
	MShouldShowHighlight               bool
	MSignificanceRange                 float64
	MSkipBuildEffect                   bool
	MToggleDormancyOnInteraction       bool
	MVisualState                       int
	MaxRenderDistance                  float64
}

var (
	RailroadBlockSignal = FGBuildableRailroadSignal{
		Name:                               "RailroadBlockSignal",
		ClassName:                          "Build_RailroadBlockSignal_C",
		MAllowColoring:                     true,
		MAspect:                            `RSA_None`,
		MAttachmentPoints:                  ``,
		MBlockValidation:                   `RBV_Unvalidated`,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Train Signals are used to direct the movement of Trains to avoid collisions and bottlenecks.

Block Signals can be placed on Railways to create 'Blocks' between each other. When a Train is occupying such a Block, other Trains will be unable to enter it.

Caution: Signals are directional! Trains are unable to move against this direction, so be sure to set up Signals in both directions for bi-directional Railways.`,
		MDisplayName:                      `Block Signal`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MGuardedConnections:               ``,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsBiDirectional:                  false,
		MIsPathSignal:                     false,
		MIsUseable:                        true,
		MObservedConnections:              ``,
		MOnAspectChangedDelegate:          `()`,
		MOnBlockValidationChangedDelegate: `()`,
		MPreviousAspect:                   `RSA_None`,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSignificanceRange:                75000.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MVisualState:                      0,
		MaxRenderDistance:                 -1.000000,
	}

	RailroadPathSignal = FGBuildableRailroadSignal{
		Name:                               "RailroadPathSignal",
		ClassName:                          "Build_RailroadPathSignal_C",
		MAllowColoring:                     true,
		MAspect:                            `RSA_None`,
		MAttachmentPoints:                  ``,
		MBlockValidation:                   `RBV_Unvalidated`,
		MBuildEffectSpeed:                  0.000000,
		MCreateClearanceMeshRepresentation: true,
		MDescription: `Train Signals are used to direct the movement of Trains to avoid collisions and bottlenecks.

Path Signals are advanced signals that are especially useful for bi-directional Railways and complex intersections. They function similarly to Block Signals but rather than occupying the entire Block, Trains can reserve a specific path through it and will only enter the Block if their path allows them to fully pass through it.

Caution: Signals are directional! Trains are unable to move against this direction, so be sure to set up Signals in both directions for bi-directional Railways.`,
		MDisplayName:                      `Path Signal`,
		MForceNetUpdateOnRegisterPlayer:   false,
		MGuardedConnections:               ``,
		MHideOnBuildEffectStart:           false,
		MHighlightVector:                  `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:               ``,
		MIsBiDirectional:                  false,
		MIsPathSignal:                     true,
		MIsUseable:                        true,
		MObservedConnections:              ``,
		MOnAspectChangedDelegate:          `()`,
		MOnBlockValidationChangedDelegate: `()`,
		MShouldModifyWorldGrid:            true,
		MShouldShowAttachmentPointVisuals: false,
		MShouldShowHighlight:              false,
		MSignificanceRange:                75000.000000,
		MSkipBuildEffect:                  false,
		MToggleDormancyOnInteraction:      false,
		MVisualState:                      0,
		MaxRenderDistance:                 -1.000000,
	}
)

func GetByClassName(className string) (FGBuildableRailroadSignal, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableRailroadSignal{}, fmt.Errorf("failed to find FGBuildableRailroadSignal with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableRailroadSignal{
	"Build_RailroadBlockSignal_C": RailroadBlockSignal,
	"Build_RailroadPathSignal_C":  RailroadPathSignal,
}
