package BuildableFactorySimpleProducer

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableFactorySimpleProducer struct {
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
	MEventType                              string
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
	MTimeToProduceItem                      float64
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	TreeGiftProducer = FGBuildableFactorySimpleProducer{
		ClassName:                               "Build_TreeGiftProducer_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MDescription: `It comes bearing gifts.
Produces 15 Gifts per minute.`,
		MDisplayName:                         `FICSMAS Gift Tree`,
		MEffectUpdateInterval:                0.000000,
		MEventType:                           `EV_Christmas`,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           false,
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
		MTimeToProduceItem:                   4.000000,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableFactorySimpleProducer, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableFactorySimpleProducer with name %s", className)
}

var classNameToVar = map[string]*FGBuildableFactorySimpleProducer{
	"Build_TreeGiftProducer_C": &TreeGiftProducer,
}