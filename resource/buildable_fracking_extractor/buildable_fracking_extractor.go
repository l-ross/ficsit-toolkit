package BuildableFrackingExtractor

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableFrackingExtractor struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MAllowedResourceForms                   string
	MAllowedResources                       string
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MDescription                            string
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MExtractCycleTime                       float64
	MExtractStartupTime                     float64
	MExtractStartupTimer                    float64
	MExtractorTypeName                      string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MItemsPerCycle                          int
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
	MOnlyAllowCertainResources              bool
	MPipeOutputConnections                  string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MReplicatedFlowRate                     float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	FrackingExtractor = FGBuildableFrackingExtractor{
		ClassName:                               "Build_FrackingExtractor_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAllowedResourceForms:                   `(RF_LIQUID,RF_GAS)`,
		MAllowedResources:                       `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/CrudeOil/Desc_LiquidOil.Desc_LiquidOil_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/NitrogenGas/Desc_NitrogenGas.Desc_NitrogenGas_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Water/Desc_Water.Desc_Water_C"')`,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MDescription: `Normal extraction rate: 60m³ fluid per minute.
Head Lift: 10 meters.
(Allows fluids to be transported 10 meters upwards.)

Can be placed on the activated sub-nodes of a Resource Well to collect the pressurized resources. Does not require Power.`,
		MDisplayName:                         `Resource Well Extractor`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MExtractCycleTime:                    1.000000,
		MExtractStartupTime:                  -10.000000,
		MExtractStartupTimer:                 10.000000,
		MExtractorTypeName:                   `None`,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MItemsPerCycle:                       1000,
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
		MOnlyAllowCertainResources:           true,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MReplicatedFlowRate:                  0.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableFrackingExtractor, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableFrackingExtractor with name %s", className)
}

var classNameToVar = map[string]*FGBuildableFrackingExtractor{
	"Build_FrackingExtractor_C": &FrackingExtractor,
}