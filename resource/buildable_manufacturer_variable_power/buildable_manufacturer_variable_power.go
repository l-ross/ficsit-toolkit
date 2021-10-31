package BuildableManufacturerVariablePower

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableManufacturerVariablePower struct {
	ClassName                               string
	IsPowered                               bool
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MCurrentProducingSeekTime               float64
	MCurrentRecipeChanged                   string
	MDescription                            string
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MEstimatedMaximumPowerConsumption       float64
	MEstimatedMininumPowerConsumption       float64
	MExcludeFromMaterialInstancing          string
	MFactoryInputConnections                string
	MFactoryOutputConnections               string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MGameTimeAtProducing                    float64
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MLightningTimer                         string
	MManufacturingSpeed                     float64
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
	MPipeInputConnections                   string
	MPipeOutputConnections                  string
	MPowerConsumption                       float64
	MPowerConsumptionExponent               float64
	MSequenceDuration                       float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MStartVector_VFX_Large_End              string
	MStartVector_VFX_Large_Start            string
	MStartVector_VFX_Medium_End             string
	MStartVector_VFX_Medium_Start           string
	MStartVector_VFX_Small_End              string
	MStartVector_VFX_Small_Start            string
	MToggleDormancyOnInteraction            bool
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	HadronCollider = FGBuildableManufacturerVariablePower{
		ClassName:                               "Build_HadronCollider_C",
		IsPowered:                               false,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCurrentProducingSeekTime:               0.000000,
		MCurrentRecipeChanged:                   `()`,
		MDescription: `The FICSIT Particle Accelerator uses electromagnetic fields to propel particles to very high speeds and energies. The specific design allows for a variety of processes, such as matter generation and conversion.

Warning: Power usage is extremely high, unstable, and varies per recipe.`,
		MDisplayName:                         `Particle Accelerator`,
		MEffectUpdateInterval:                0.000000,
		MEstimatedMaximumPowerConsumption:    1500.000000,
		MEstimatedMininumPowerConsumption:    250.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MGameTimeAtProducing:                 0.000000,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MLightningTimer:                      `()`,
		MManufacturingSpeed:                  1.000000,
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
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    0.000000,
		MPowerConsumptionExponent:            1.600000,
		MSequenceDuration:                    60.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   20000.000000,
		MSkipBuildEffect:                     false,
		MStartVector_VFX_Large_End:           `((X=0.000000,Y=0.000000,Z=424.000000),(X=0.000000,Y=373.000000,Z=208.000000),(X=0.000000,Y=-378.000000,Z=211.000000))`,
		MStartVector_VFX_Large_Start:         `((X=0.000000,Y=0.000000,Z=-430.000000),(X=0.000000,Y=-372.000000,Z=-218.000000))`,
		MStartVector_VFX_Medium_End:          `((X=0.000000,Y=-212.000000,Z=380.000000),(X=0.000000,Y=224.000000,Z=379.000000))`,
		MStartVector_VFX_Medium_Start:        `((X=0.000000,Y=-219.000000,Z=-381.000000),(X=0.000000,Y=-432.000000,Z=-1.100000))`,
		MStartVector_VFX_Small_End:           `((X=0.000000,Y=0.000000,Z=400.000000),(X=0.000000,Y=-382.000000,Z=194.000000),(X=0.000000,Y=365.000000,Z=196.000000))`,
		MStartVector_VFX_Small_Start:         `((X=0.000000,Y=0.000000,Z=-420.000000),(X=0.000000,Y=365.000000,Z=-221.000000))`,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableManufacturerVariablePower, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableManufacturerVariablePower with name %s", className)
}

var classNameToVar = map[string]*FGBuildableManufacturerVariablePower{
	"Build_HadronCollider_C": &HadronCollider,
}
