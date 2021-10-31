package BuildableGeneratorFuel

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableGeneratorFuel struct {
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
		MByproductAmount           string
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
	MRTPCInterval                        float64
	MRequiresSupplementalResource        bool
	MShouldModifyWorldGrid               bool
	MShouldShowHighlight                 bool
	MSignificanceRange                   float64
	MSkipBuildEffect                     bool
	MSupplementalLoadAmount              int
	MSupplementalToPowerRatio            float64
	MToggleDormancyOnInteraction         bool
	M_CurrentPotential                   int
	M_SFXSockets                         string
	MaxRenderDistance                    float64
	OnReplicationDetailActorCreatedEvent string
}

var (
	GeneratorBiomass = FGBuildableGeneratorFuel{
		ClassName:                               "Build_GeneratorBiomass_C",
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
		MDefaultFuelClasses:                     `(/Script/FactoryGame.FGItemDescriptorBiomass)`,
		MDescription: `Burns various forms of biomass to generate electricity for the power grid.
Has no input and must therefore be fed biomass manually.

Resource consumption will automatically be lowered to meet power demands.`,
		MDisplayName:                    `Biomass Burner`,
		MEffectUpdateInterval:           1.000000,
		MExcludeFromMaterialInstancing:  ``,
		MFluidStackSizeDefault:          resource.Fluid,
		MFluidStackSizeMultiplier:       1,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MFuel: []struct {
			MByproduct                 string
			MByproductAmount           string
			MFuelClass                 string
			MSupplementalResourceClass string
		}{
			{
				MByproduct:                 ``,
				MByproductAmount:           ``,
				MFuelClass:                 `FGItemDescriptorBiomass`,
				MSupplementalResourceClass: ``,
			},
		},
		MFuelClasses:                         ``,
		MFuelLoadAmount:                      1,
		MFuelResourceForm:                    resource.Solid,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsFullBlast:                         false,
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
		MPowerProduction:                     30.000000,
		MPowerProductionExponent:             1.300000,
		MRequiresSupplementalResource:        false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   12000.000000,
		MSkipBuildEffect:                     false,
		MSupplementalLoadAmount:              0,
		MSupplementalToPowerRatio:            0.000000,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	GeneratorCoal = FGBuildableGeneratorFuel{
		ClassName:                               "Build_GeneratorCoal_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MAvailableFuelClasses:                   ``,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedInputConnections:                 ``,
		MCachedPipeInputConnections:             ``,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MDefaultFuelClasses:                     `(/Game/FactoryGame/Resource/RawResources/Coal/Desc_Coal.Desc_Coal_C,/Game/FactoryGame/Resource/Parts/CompactedCoal/Desc_CompactedCoal.Desc_CompactedCoal_C,/Game/FactoryGame/Resource/Parts/PetroleumCoke/Desc_PetroleumCoke.Desc_PetroleumCoke_C)`,
		MDescription: `Burns Coal to boil Water, the produced steam rotates turbines to generate electricity for the power grid.
Has a Conveyor Belt and Pipe input, so both the Coal and Water supply can be automated.

Caution: Always generates at the set clockspeed. Shuts down if fuel requirements are not met.`,
		MDisplayName:                    `Coal Generator`,
		MEffectUpdateInterval:           0.000000,
		MExcludeFromMaterialInstancing:  ``,
		MFluidStackSizeDefault:          resource.Fluid,
		MFluidStackSizeMultiplier:       1,
		MFogPlaneTransforms:             `((Translation=(X=199.904510,Y=1143.047119,Z=162.721100)))`,
		MForceNetUpdateOnRegisterPlayer: false,
		MFuel: []struct {
			MByproduct                 string
			MByproductAmount           string
			MFuelClass                 string
			MSupplementalResourceClass string
		}{
			{
				MByproduct:                 ``,
				MByproductAmount:           ``,
				MFuelClass:                 `Desc_Coal_C`,
				MSupplementalResourceClass: `Desc_Water_C`,
			},
			{
				MByproduct:                 ``,
				MByproductAmount:           ``,
				MFuelClass:                 `Desc_CompactedCoal_C`,
				MSupplementalResourceClass: `Desc_Water_C`,
			},
			{
				MByproduct:                 ``,
				MByproductAmount:           ``,
				MFuelClass:                 `Desc_PetroleumCoke_C`,
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
		MPowerProduction:                     75.000000,
		MPowerProductionExponent:             1.300000,
		MRequiresSupplementalResource:        true,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   20000.000000,
		MSkipBuildEffect:                     false,
		MSupplementalLoadAmount:              1000,
		MSupplementalToPowerRatio:            10.000000,
		MToggleDormancyOnInteraction:         false,
		M_CurrentPotential:                   1,
		M_SFXSockets:                         `("AudioSocketTurbine","CoalGeneratorPotential")`,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	GeneratorFuel = FGBuildableGeneratorFuel{
		ClassName:                               "Build_GeneratorFuel_C",
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
		MDefaultFuelClasses:                     `(/Game/FactoryGame/Resource/Parts/Fuel/Desc_LiquidFuel.Desc_LiquidFuel_C,/Game/FactoryGame/Resource/Parts/Turbofuel/Desc_LiquidTurboFuel.Desc_LiquidTurboFuel_C,/Game/FactoryGame/Resource/Parts/BioFuel/Desc_LiquidBiofuel.Desc_LiquidBiofuel_C)`,
		MDescription: `Consumes Fuel to generate electricity for the power grid.
Has a Pipe input so the Fuel supply can be automated.

Caution: Always generates at the set clockspeed. Shuts down if fuel requirements are not met.`,
		MDisplayName:                    `Fuel Generator`,
		MEffectUpdateInterval:           2.000000,
		MExcludeFromMaterialInstancing:  ``,
		MFluidStackSizeDefault:          resource.Fluid,
		MFluidStackSizeMultiplier:       1,
		MFogPlaneTransforms:             ``,
		MForceNetUpdateOnRegisterPlayer: false,
		MFuel: []struct {
			MByproduct                 string
			MByproductAmount           string
			MFuelClass                 string
			MSupplementalResourceClass string
		}{
			{
				MByproduct:                 ``,
				MByproductAmount:           ``,
				MFuelClass:                 `Desc_LiquidFuel_C`,
				MSupplementalResourceClass: ``,
			},
			{
				MByproduct:                 ``,
				MByproductAmount:           ``,
				MFuelClass:                 `Desc_LiquidTurboFuel_C`,
				MSupplementalResourceClass: ``,
			},
			{
				MByproduct:                 ``,
				MByproductAmount:           ``,
				MFuelClass:                 `Desc_LiquidBiofuel_C`,
				MSupplementalResourceClass: ``,
			},
		},
		MFuelClasses:                         ``,
		MFuelLoadAmount:                      1000,
		MFuelResourceForm:                    resource.Liquid,
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
		MPowerProduction:                     150.000000,
		MPowerProductionExponent:             1.300000,
		MRTPCInterval:                        0.000000,
		MRequiresSupplementalResource:        false,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MSupplementalLoadAmount:              0,
		MSupplementalToPowerRatio:            0.000000,
		MToggleDormancyOnInteraction:         false,
		M_CurrentPotential:                   0,
		M_SFXSockets:                         `("AudioSocket_Exhaust","AudioSocket_Root")`,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableGeneratorFuel, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableGeneratorFuel with name %s", className)
}

var classNameToVar = map[string]*FGBuildableGeneratorFuel{
	"Build_GeneratorBiomass_C": &GeneratorBiomass,
	"Build_GeneratorCoal_C":    &GeneratorCoal,
	"Build_GeneratorFuel_C":    &GeneratorFuel,
}