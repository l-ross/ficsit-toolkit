package BuildablePowerStorage

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildablePowerStorage struct {
	ClassName                               string
	MActivationEventID                      int
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBatteryStatus                          string
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MCurrentGameTimeSinceStateChange        float64
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
	MIndicatorLevel                         int
	MIndicatorLevelMax                      int
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
	MPowerInputCapacity                     float64
	MPowerStore                             float64
	MPowerStoreCapacity                     float64
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MToggleDormancyOnInteraction            bool
	M_PreviousBatteryStatus                 string
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	PowerStorageMk1 = FGBuildablePowerStorage{
		ClassName:                               "Build_PowerStorageMk1_C",
		MActivationEventID:                      0,
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBatteryStatus:                          `BS_Idle`,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     false,
		MCurrentGameTimeSinceStateChange:        0.000000,
		MDescription: `Storage Capacity: 100 MWh (100 MW for 1 hour)
Max Charge Rate: 100 MW
Max Discharge Rate: Unlimited 

Can be connected to a Power Grid to store excess power production. The stored power can be used later in cases of high consumption.`,
		MDisplayName:                         `Power Storage`,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  ``,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MIndicatorLevel:                      0,
		MIndicatorLevelMax:                   5,
		MInteractingPlayers:                  ``,
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
		MPowerInputCapacity:                  100.000000,
		MPowerStore:                          0.000000,
		MPowerStoreCapacity:                  100.000000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   14000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		M_PreviousBatteryStatus:              `BS_Idle`,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildablePowerStorage, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildablePowerStorage with name %s", className)
}

var classNameToVar = map[string]*FGBuildablePowerStorage{
	"Build_PowerStorageMk1_C": &PowerStorageMk1,
}