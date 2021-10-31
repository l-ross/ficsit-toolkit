package BuildableStorage

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableStorage struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
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
	MInventorySizeX                         int
	MInventorySizeY                         int
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
	MStackingHeight                         float64
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	StorageContainerMk1 = FGBuildableStorage{
		ClassName:                               "Build_StorageContainerMk1_C",
		MAddToSignificanceManager:               false,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDescription: `Contains 24 slots for storing large amounts of items.
Has an input and output for conveyor belts.`,
		MDisplayName:                         `Storage Container`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000012),Translation=(X=0.003135,Y=-472.392761,Z=175.000000)),(Rotation=(X=0.000000,Y=-0.000000,Z=0.000018,W=1.000000),Translation=(X=-0.001938,Y=475.408813,Z=175.000000)))`,
		MForceNetUpdateOnRegisterPlayer:      true,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MInventorySizeX:                      6,
		MInventorySizeY:                      4,
		MIsUseable:                           true,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
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
		MStackingHeight:                      400.000000,
		MToggleDormancyOnInteraction:         true,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	StorageContainerMk2 = FGBuildableStorage{
		ClassName:                               "Build_StorageContainerMk2_C",
		MAddToSignificanceManager:               false,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDescription: `Contains 48 slots for storing large amounts of items.
Has two inputs and outputs for conveyor belts.`,
		MDisplayName:                         `Industrial Storage Container`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      true,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MInventorySizeX:                      8,
		MInventorySizeY:                      6,
		MIsUseable:                           true,
		MMaterialNameToInstanceManager:       `()`,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
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
		MStackingHeight:                      800.000000,
		MToggleDormancyOnInteraction:         true,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	StoragePlayer = FGBuildableStorage{
		ClassName:                               "Build_StoragePlayer_C",
		MAddToSignificanceManager:               false,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDescription:                            `Contains 25 slots for storing large amounts of items.`,
		MDisplayName:                            `Personal Storage Box`,
		MEffectUpdateInterval:                   0.000000,
		MExcludeFromMaterialInstancing:          ``,
		MFluidStackSizeDefault:                  resource.Fluid,
		MFluidStackSizeMultiplier:               1,
		MFogPlaneTransforms:                     ``,
		MForceNetUpdateOnRegisterPlayer:         true,
		MHideOnBuildEffectStart:                 false,
		MHighlightVector:                        `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                     ``,
		MInventorySizeX:                         5,
		MInventorySizeY:                         5,
		MIsUseable:                              true,
		MMaterialNameToInstanceManager:          `()`,
		MMaxPotential:                           1.000000,
		MMaxPotentialIncreasePerCrystal:         0.500000,
		MMinPotential:                           0.010000,
		MMinimumProducingTime:                   2.000000,
		MMinimumStoppedTime:                     5.000000,
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
		MStackingHeight:                         0.000000,
		MToggleDormancyOnInteraction:            true,
		MaxRenderDistance:                       -1.000000,
		OnReplicationDetailActorCreatedEvent:    `()`,
	}
)

func GetByClassName(className string) (*FGBuildableStorage, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableStorage with name %s", className)
}

var classNameToVar = map[string]*FGBuildableStorage{
	"Build_StorageContainerMk1_C": &StorageContainerMk1,
	"Build_StorageContainerMk2_C": &StorageContainerMk2,
	"Build_StoragePlayer_C":       &StoragePlayer,
}