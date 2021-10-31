package BuildableAttachmentMerger

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableAttachmentMerger struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MCurrentInputIndex                      int
	MCurrentInventoryIndex                  int
	MDescription                            string
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MMaterialNameToInstanceManager          string
	MMaxPotential                           float64
	MMaxPotentialIncreasePerCrystal         float64
	MMinPotential                           float64
	MMinimumProducingTime                   float64
	MMinimumStoppedTime                     float64
	MNumCyclesForProductivity               int
	MOnHasPowerChanged                      string
	MOnHasProductionChanged                 string
	MOnHasStandbyChanged                    string
	MOnProductionStatusChanged              string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	ConveyorAttachmentMerger = FGBuildableAttachmentMerger{
		ClassName:                               "Build_ConveyorAttachmentMerger_C",
		MAddToSignificanceManager:               false,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MCurrentInputIndex:                      -1,
		MCurrentInventoryIndex:                  0,
		MDescription:                            `Merges up to three conveyor belts into one.`,
		MDisplayName:                            `Conveyor Merger`,
		MEffectUpdateInterval:                   0.000000,
		MExcludeFromMaterialInstancing:          ``,
		MFluidStackSizeDefault:                  resource.Fluid,
		MFluidStackSizeMultiplier:               1,
		MFogPlaneTransforms:                     `((Translation=(X=0.000000,Y=185.000000,Z=50.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-0.707112,W=0.707101),Translation=(X=184.999939,Y=0.000286,Z=50.000000)),(Rotation=(X=0.000000,Y=-0.000000,Z=1.000000,W=0.000004),Translation=(X=0.000129,Y=-185.000000,Z=50.000000)),(Rotation=(X=0.000000,Y=-0.000000,Z=0.707098,W=0.707116),Translation=(X=-184.999939,Y=-0.000425,Z=50.000000)))`,
		MForceNetUpdateOnRegisterPlayer:         false,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MIsUseable:                              false,
		MMaterialNameToInstanceManager:          `()`,
		MMaxPotential:                           1.000000,
		MMaxPotentialIncreasePerCrystal:         0.500000,
		MMinPotential:                           0.010000,
		MMinimumProducingTime:                   -1.000000,
		MMinimumStoppedTime:                     -1.000000,
		MNumCyclesForProductivity:               20,
		MOnHasPowerChanged:                      `()`,
		MOnHasProductionChanged:                 `()`,
		MOnHasStandbyChanged:                    `()`,
		MOnProductionStatusChanged:              `()`,
		MPowerConsumption:                       0.000000,
		MPowerConsumptionExponent:               1.600000,
		MShouldModifyWorldGrid:                  true,
		MShouldShowHighlight:                    false,
		MSignificanceRange:                      18000.000000,
		MSkipBuildEffect:                        false,
		MToggleDormancyOnInteraction:            false,
		MaxRenderDistance:                       -1.000000,
		OnReplicationDetailActorCreatedEvent:    `()`,
	}
)

func GetByClassName(className string) (*FGBuildableAttachmentMerger, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableAttachmentMerger with name %s", className)
}

var classNameToVar = map[string]*FGBuildableAttachmentMerger{
	"Build_ConveyorAttachmentMerger_C": &ConveyorAttachmentMerger,
}