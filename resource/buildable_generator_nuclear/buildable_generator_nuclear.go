package BuildableGeneratorNuclear

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableGeneratorNuclear struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MAvailableFuelClasses                   string
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedInputConnections                 string
	MCachedLoadPercentage                   float64
	MCachedPipeInputConnections             string
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MCurrentGeneratorNuclearWarning         string
	MDefaultFuelClasses                     string
	MDescription                            string
	MDisplayName                            string
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MFuel                                   []struct {
		MByproduct                 string
		MByproductAmount           int
		MFuelClass                 string
		MSupplementalResourceClass string
	}

	MFuelClasses                         string
	MFuelLoadAmount                      int
	MFuelResourceForm                    resource.Form
	MHideOnBuildEffectStart              bool
	MHighlightVector                     string
	MInteractingPlayers                  string
	MIsFullBlast                         bool
	MIsUseable                           bool
	MLoadPercentage                      float64
	MMaterialNameToInstanceManager       string
	MMaxPotential                        float64
	MMaxPotentialIncreasePerCrystal      float64
	MMinPotential                        float64
	MMinimumProducingTime                float64
	MMinimumStoppedTime                  float64
	MNumCyclesForProductivity            int
	MOnHasPowerChanged                   string
	MOnHasProductionChanged              string
	MOnHasStandbyChanged                 string
	MOnProductionStatusChanged           string
	MPowerConsumption                    float64
	MPowerConsumptionExponent            float64
	MPowerProduction                     float64
	MPowerProductionExponent             float64
	MRequiresSupplementalResource        bool
	MShouldModifyWorldGrid               bool
	MShouldShowHighlight                 bool
	MSignificanceRange                   float64
	MSkipBuildEffect                     bool
	MSupplementalLoadAmount              int
	MSupplementalToPowerRatio            float64
	MToggleDormancyOnInteraction         bool
	MWasteLeftFromCurrentFuel            int
	MaxRenderDistance                    float64
	OnReplicationDetailActorCreatedEvent string
}

var (
	GeneratorNuclear = FGBuildableGeneratorNuclear{
		ClassName:                               "Build_GeneratorNuclear_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAvailableFuelClasses:                   ``,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedInputConnections:                 ``,
		MCachedLoadPercentage:                   0.000000,
		MCachedPipeInputConnections:             ``,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MCurrentGeneratorNuclearWarning:         `GNW_None`,
		MDefaultFuelClasses:                     `(/Game/FactoryGame/Resource/Parts/NuclearFuelRod/Desc_NuclearFuelRod.Desc_NuclearFuelRod_C,/Game/FactoryGame/Resource/Parts/PlutoniumFuelRods/Desc_PlutoniumFuelRod.Desc_PlutoniumFuelRod_C)`,
		MDescription: `Consumes Nuclear Fuel Rods and Water to produce electricity for the power grid.

Produces Nuclear Waste, which is extracted from the conveyor belt output.

Caution: Always generates at the set clockspeed. Shuts down if fuel requirements are not met.`,
		MDisplayName:                    `Nuclear Power Plant`,
		MEffectUpdateInterval:           2.000000,
		MExcludeFromMaterialInstancing:  ``,
		MFluidStackSizeDefault:          resource.Fluid,
		MFluidStackSizeMultiplier:       1,
		MFogPlaneTransforms:             `((Translation=(X=-200.000000,Y=2088.737549,Z=181.784302)),(Translation=(X=200.000000,Y=2088.737549,Z=181.784302)))`,
		MForceNetUpdateOnRegisterPlayer: false,
		MFuel: []struct {
			MByproduct                 string
			MByproductAmount           int
			MFuelClass                 string
			MSupplementalResourceClass string
		}{
			{
				MByproduct:                 `Desc_NuclearWaste_C`,
				MByproductAmount:           50,
				MFuelClass:                 `Desc_NuclearFuelRod_C`,
				MSupplementalResourceClass: `Desc_Water_C`,
			},
			{
				MByproduct:                 `Desc_PlutoniumWaste_C`,
				MByproductAmount:           10,
				MFuelClass:                 `Desc_PlutoniumFuelRod_C`,
				MSupplementalResourceClass: `Desc_Water_C`,
			},
		},
		MFuelClasses:                         ``,
		MFuelLoadAmount:                      1,
		MFuelResourceForm:                    resource.Solid,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsFullBlast:                         true,
		MIsUseable:                           true,
		MLoadPercentage:                      0.000000,
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
		MPowerProduction:                     2500.000000,
		MPowerProductionExponent:             1.321928,
		MRequiresSupplementalResource:        true,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   8000.000000,
		MSkipBuildEffect:                     false,
		MSupplementalLoadAmount:              10000,
		MSupplementalToPowerRatio:            2.000000,
		MToggleDormancyOnInteraction:         false,
		MWasteLeftFromCurrentFuel:            0,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableGeneratorNuclear, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableGeneratorNuclear with name %s", className)
}

var classNameToVar = map[string]*FGBuildableGeneratorNuclear{
	"Build_GeneratorNuclear_C": &GeneratorNuclear,
}
