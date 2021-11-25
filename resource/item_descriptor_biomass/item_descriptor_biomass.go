// Code generated by ../../gen/docs_json. DO NOT EDIT.

package ItemDescriptorBiomass

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGItemDescriptorBiomass struct {
	Name                    string
	ClassName               string
	MAbbreviatedDisplayName string
	MCanBeDiscarded         bool
	MDescription            string
	MDisplayName            string
	MEnergyValue            float64
	MFluidColor             string
	MForm                   resource.Form
	MGasColor               string
	MPersistentBigIcon      string
	MRadioactiveDecay       float64
	MRememberPickUp         bool
	MResourceSinkPoints     int
	MSmallIcon              string
	MStackSize              resource.StackSize
}

var (
	Biofuel = FGItemDescriptorBiomass{
		Name:                    "Biofuel",
		ClassName:               "Desc_Biofuel_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `The most energy-efficient form of solid biomass. Can be used as fuel for the Chainsaw.`,
		MDisplayName:            `Solid Biofuel`,
		MEnergyValue:            450.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/SolidBiofuel/UI/IconDesc_SolidBiofuel_256.IconDesc_SolidBiofuel_256`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     48,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/SolidBiofuel/UI/IconDesc_SolidBiofuel_64.IconDesc_SolidBiofuel_64`,
		MStackSize:              resource.Big,
	}

	Fabric = FGItemDescriptorBiomass{
		Name:                    "Fabric",
		ClassName:               "Desc_Fabric_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for equipment crafting.
Flexible but durable fabric.`,
		MDisplayName:        `Fabric`,
		MEnergyValue:        15.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Fabric_256.IconDesc_Fabric_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 140,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Fabric_64.IconDesc_Fabric_64`,
		MStackSize:          resource.Medium,
	}

	FlowerPetals = FGItemDescriptorBiomass{
		Name:                    "FlowerPetals",
		ClassName:               "Desc_FlowerPetals_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
Colorful native flower petals.`,
		MDisplayName:        `Flower Petals`,
		MEnergyValue:        100.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/FlowerPetals_Final_256.FlowerPetals_Final_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 10,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/FlowerPetals_Final_64.FlowerPetals_Final_64`,
		MStackSize:          resource.Huge,
	}

	GenericBiomass = FGItemDescriptorBiomass{
		Name:                    "GenericBiomass",
		ClassName:               "Desc_GenericBiomass_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Primarily used as fuel.
Biomass burners and vehicles can use it for power.
Biomass is much more energy-efficient than raw biological matter.`,
		MDisplayName:        `Biomass`,
		MEnergyValue:        180.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Biomass_Final_256.IconDesc_Biomass_Final_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 12,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Biomass_Final_64.IconDesc_Biomass_Final_64`,
		MStackSize:          resource.Big,
	}

	HogParts = FGItemDescriptorBiomass{
		Name:                    "HogParts",
		ClassName:               "Desc_HogParts_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Thick and sturdy natural armor plates from alien creatures.`,
		MDisplayName:            `Alien Carapace`,
		MEnergyValue:            250.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/AnimalParts/UI/IconDesc_HogPart_256.IconDesc_HogPart_256`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         true,
		MResourceSinkPoints:     0,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/AnimalParts/UI/IconDesc_HogPart_64.IconDesc_HogPart_64`,
		MStackSize:              resource.Small,
	}

	Leaves = FGItemDescriptorBiomass{
		Name:                    "Leaves",
		ClassName:               "Desc_Leaves_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Primarily used as fuel.
Biomass Burners and vehicles can use it for power.`,
		MDisplayName:        `Leaves`,
		MEnergyValue:        15.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/IconDesc_Leaves_256.IconDesc_Leaves_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 3,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Leaves_64.IconDesc_Leaves_64`,
		MStackSize:          resource.Huge,
	}

	LiquidBiofuel = FGItemDescriptorBiomass{
		Name:                    "LiquidBiofuel",
		ClassName:               "Desc_LiquidBiofuel_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Liquid Biofuel can be used to generate power or packaged to be used as fuel for Vehicles.`,
		MDisplayName:            `Liquid Biofuel`,
		MEnergyValue:            0.750000,
		MFluidColor:             `(B=44,G=83,R=59,A=255)`,
		MForm:                   resource.Liquid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/BioFuel/UI/IconDesc_LiquidBiofuel_Pipe_512.IconDesc_LiquidBiofuel_Pipe_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     261,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/BioFuel/UI/IconDesc_LiquidBiofuel_Pipe_256.IconDesc_LiquidBiofuel_Pipe_256`,
		MStackSize:              resource.Fluid,
	}

	Mycelia = FGItemDescriptorBiomass{
		Name:                    "Mycelia",
		ClassName:               "Desc_Mycelia_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Used for crafting.
Biomass Burners and vehicles can use it for power.`,
		MDisplayName:        `Mycelia`,
		MEnergyValue:        20.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Mycelia_256.IconDesc_Mycelia_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 10,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Mycelia_64.IconDesc_Mycelia_64`,
		MStackSize:          resource.Big,
	}

	PackagedAlumina = FGItemDescriptorBiomass{
		Name:                    "PackagedAlumina",
		ClassName:               "Desc_PackagedAlumina_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Alumina Solution, packaged for alternative transport.`,
		MDisplayName:            `Packaged Alumina Solution`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=42,G=83,R=58,A=255)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/Alumina/UI/IconDesc_PackagedAluminaSolution_256.IconDesc_PackagedAluminaSolution_256`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     160,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/Alumina/UI/IconDesc_PackagedAluminaSolution_64.IconDesc_PackagedAluminaSolution_64`,
		MStackSize:              resource.Medium,
	}

	PackagedBiofuel = FGItemDescriptorBiomass{
		Name:                    "PackagedBiofuel",
		ClassName:               "Desc_PackagedBiofuel_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Liquid Biofuel, packaged for alternative transport. Can be used as fuel for Vehicles.`,
		MDisplayName:            `Packaged Liquid Biofuel`,
		MEnergyValue:            750.000000,
		MFluidColor:             `(B=42,G=83,R=58,A=255)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/Turbofuel/UI/IconDesc_LiquidBiofuel_256.IconDesc_LiquidBiofuel_256`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     370,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/Turbofuel/UI/IconDesc_LiquidBiofuel_64.IconDesc_LiquidBiofuel_64`,
		MStackSize:              resource.Medium,
	}

	PackagedNitricAcid = FGItemDescriptorBiomass{
		Name:                    "PackagedNitricAcid",
		ClassName:               "Desc_PackagedNitricAcid_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Nitric Acid, packaged for alternative transport.`,
		MDisplayName:            `Packaged Nitric Acid`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=42,G=83,R=58,A=255)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/NitricAcid/UI/IconDesc_PackagedNitricAcid_256.IconDesc_PackagedNitricAcid_256`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     412,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/NitricAcid/UI/IconDesc_PackagedNitricAcid_64.IconDesc_PackagedNitricAcid_64`,
		MStackSize:              resource.Medium,
	}

	PackagedNitrogenGas = FGItemDescriptorBiomass{
		Name:                    "PackagedNitrogenGas",
		ClassName:               "Desc_PackagedNitrogenGas_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Nitrogen Gas, packaged for alternative transport.`,
		MDisplayName:            `Packaged Nitrogen Gas`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=42,G=83,R=58,A=255)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/PackagedNitrogen/UI/IconDesc_PackagedNitrogen_256.IconDesc_PackagedNitrogen_256`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     312,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/PackagedNitrogen/UI/IconDesc_PackagedNitrogen_64.IconDesc_PackagedNitrogen_64`,
		MStackSize:              resource.Medium,
	}

	PackagedSulfuricAcid = FGItemDescriptorBiomass{
		Name:                    "PackagedSulfuricAcid",
		ClassName:               "Desc_PackagedSulfuricAcid_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Sulfuric Acid, packaged for alternative transport.`,
		MDisplayName:            `Packaged Sulfuric Acid`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=42,G=83,R=58,A=255)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/SulfuricAcid/UI/IconDesc_PckagedSulphuricAcid_256.IconDesc_PckagedSulphuricAcid_256`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     152,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/SulfuricAcid/UI/IconDesc_PckagedSulphuricAcid_64.IconDesc_PckagedSulphuricAcid_64`,
		MStackSize:              resource.Medium,
	}

	SpitterParts = FGItemDescriptorBiomass{
		Name:                    "SpitterParts",
		ClassName:               "Desc_SpitterParts_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription:            `Organs from alien creatures.`,
		MDisplayName:            `Alien Organs`,
		MEnergyValue:            250.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/AnimalParts/UI/IconDesc_SpitterPart_256.IconDesc_SpitterPart_256`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         true,
		MResourceSinkPoints:     0,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/AnimalParts/UI/IconDesc_SpitterPart_64.IconDesc_SpitterPart_64`,
		MStackSize:              resource.Small,
	}

	Wood = FGItemDescriptorBiomass{
		Name:                    "Wood",
		ClassName:               "Desc_Wood_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MDescription: `Primarily used as fuel.
Biomass Burners and vehicles can use it for power.`,
		MDisplayName:        `Wood`,
		MEnergyValue:        100.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Wood_256.IconDesc_Wood_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 30,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Wood_64.IconDesc_Wood_64`,
		MStackSize:          resource.Big,
	}
)

func GetByClassName(className string) (FGItemDescriptorBiomass, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGItemDescriptorBiomass{}, fmt.Errorf("failed to find FGItemDescriptorBiomass with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGItemDescriptorBiomass{
	"Desc_Biofuel_C":              Biofuel,
	"Desc_Fabric_C":               Fabric,
	"Desc_FlowerPetals_C":         FlowerPetals,
	"Desc_GenericBiomass_C":       GenericBiomass,
	"Desc_HogParts_C":             HogParts,
	"Desc_Leaves_C":               Leaves,
	"Desc_LiquidBiofuel_C":        LiquidBiofuel,
	"Desc_Mycelia_C":              Mycelia,
	"Desc_PackagedAlumina_C":      PackagedAlumina,
	"Desc_PackagedBiofuel_C":      PackagedBiofuel,
	"Desc_PackagedNitricAcid_C":   PackagedNitricAcid,
	"Desc_PackagedNitrogenGas_C":  PackagedNitrogenGas,
	"Desc_PackagedSulfuricAcid_C": PackagedSulfuricAcid,
	"Desc_SpitterParts_C":         SpitterParts,
	"Desc_Wood_C":                 Wood,
}
