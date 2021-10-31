package BuildableSplitterSmart

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableSplitterSmart struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MCurrentInventoryIndex                  int
	MCurrentOutputIndex                     int
	MDescription                            string
	MDisplayName                            string
	MDistributionTable                      string
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
	MItemToLastOutputMap                    string
	MLastItem                               string
	MLastOutputIndex                        int
	MMaterialNameToInstanceManager          string
	MMaxNumSortRules                        int
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
	OnSortRulesChangedDelegate              string
}

var (
	ConveyorAttachmentSplitterProgrammable = FGBuildableSplitterSmart{
		ClassName:                               "Build_ConveyorAttachmentSplitterProgrammable_C",
		MAddToSignificanceManager:               false,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MCurrentInventoryIndex:                  0,
		MCurrentOutputIndex:                     -1,
		MDescription: `Splits conveyor belts in three. 
You can set rules for each output to decide exactly where each part should go.`,
		MDisplayName:                         `Programmable Splitter`,
		MDistributionTable:                   ``,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Translation=(X=0.000000,Y=185.000000,Z=50.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-0.707112,W=0.707101),Translation=(X=184.999939,Y=0.000286,Z=50.000000)),(Rotation=(X=0.000000,Y=-0.000000,Z=1.000000,W=0.000004),Translation=(X=0.000129,Y=-185.000000,Z=50.000000)),(Rotation=(X=0.000000,Y=-0.000000,Z=0.707098,W=0.707116),Translation=(X=-184.999939,Y=-0.000425,Z=50.000000)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MItemToLastOutputMap:                 `()`,
		MLastItem:                            `(ItemState=())`,
		MLastOutputIndex:                     0,
		MMaterialNameToInstanceManager:       `()`,
		MMaxNumSortRules:                     64,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                -1.000000,
		MMinimumStoppedTime:                  -1.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
		OnSortRulesChangedDelegate:           `()`,
	}

	ConveyorAttachmentSplitterSmart = FGBuildableSplitterSmart{
		ClassName:                               "Build_ConveyorAttachmentSplitterSmart_C",
		MAddToSignificanceManager:               false,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MCurrentInventoryIndex:                  0,
		MCurrentOutputIndex:                     -1,
		MDescription: `Splits conveyor belts in three.
You can set a rule for each output to decide exactly what part should go there.`,
		MDisplayName:                         `Smart Splitter`,
		MDistributionTable:                   ``,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Translation=(X=0.000000,Y=185.000000,Z=50.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-0.707112,W=0.707101),Translation=(X=184.999939,Y=0.000286,Z=50.000000)),(Rotation=(X=0.000000,Y=-0.000000,Z=1.000000,W=0.000004),Translation=(X=0.000129,Y=-185.000000,Z=50.000000)),(Rotation=(X=0.000000,Y=-0.000000,Z=0.707098,W=0.707116),Translation=(X=-184.999939,Y=-0.000425,Z=50.000000)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MItemToLastOutputMap:                 `()`,
		MLastItem:                            `(ItemState=())`,
		MLastOutputIndex:                     0,
		MMaterialNameToInstanceManager:       `()`,
		MMaxNumSortRules:                     0,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                -1.000000,
		MMinimumStoppedTime:                  -1.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MOnProductionStatusChanged:           `()`,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
		OnSortRulesChangedDelegate:           `()`,
	}
)

func GetByClassName(className string) (*FGBuildableSplitterSmart, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableSplitterSmart with name %s", className)
}

var classNameToVar = map[string]*FGBuildableSplitterSmart{
	"Build_ConveyorAttachmentSplitterProgrammable_C": &ConveyorAttachmentSplitterProgrammable,
	"Build_ConveyorAttachmentSplitterSmart_C":        &ConveyorAttachmentSplitterSmart,
}