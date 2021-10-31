package ResourceDescriptor

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGResourceDescriptor struct {
	ClassName               string
	MAbbreviatedDisplayName string
	MCanBeDiscarded         bool
	MCollectSpeedMultiplier float64
	MDecalSize              float64
	MDescription            string
	MDisplayName            string
	MEnergyValue            float64
	MFluidColor             string
	MForm                   resource.Form
	MGasColor               string
	MManualMiningAudioName  string
	MPersistentBigIcon      string
	MPingColor              string
	MRadioactiveDecay       float64
	MRememberPickUp         bool
	MResourceSinkPoints     int
	MSmallIcon              string
	MStackSize              resource.StackSize
}

var (
	Coal = FGResourceDescriptor{
		ClassName:               "Desc_Coal_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription:            `Mainly used as fuel for vehicles & coal generators and for steel production.`,
		MDisplayName:            `Coal`,
		MEnergyValue:            300.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName:  `Metal`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/RawResources/Coal/UI/CoalOre_256.CoalOre_256`,
		MPingColor:              `(R=0.048172,G=0.048172,B=0.048172,A=1.000000)`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         true,
		MResourceSinkPoints:     3,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/RawResources/Coal/UI/CoalOre_64.CoalOre_64`,
		MStackSize:              resource.Medium,
	}

	LiquidOil = FGResourceDescriptor{
		ClassName:               "Desc_LiquidOil_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription:            `Crude Oil is refined into all kinds of Oil-based resources, like Fuel and Plastic.`,
		MDisplayName:            `Crude Oil`,
		MEnergyValue:            0.320000,
		MFluidColor:             `(B=25,G=0,R=25,A=255)`,
		MForm:                   resource.Liquid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName:  `Metal`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/RawResources/CrudeOil/UI/LiquidOil_Pipe_512.LiquidOil_Pipe_512`,
		MPingColor:              `(R=0.000000,G=0.000000,B=0.000000,A=1.000000)`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     30,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/RawResources/CrudeOil/UI/LiquidOil_Pipe_256.LiquidOil_Pipe_256`,
		MStackSize:              resource.Fluid,
	}

	NitrogenGas = FGResourceDescriptor{
		ClassName:               "Desc_NitrogenGas_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription:            `Nitrogen can be used in a variety of ways, such as metallurgy, cooling, and Nitric Acid production. On Massage-2(AB)b, it can be extracted from underground gas wells.`,
		MDisplayName:            `Nitrogen Gas`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=89,G=89,R=89,A=255)`,
		MForm:                   resource.Gas,
		MGasColor:               `(B=255,G=255,R=255,A=255)`,
		MManualMiningAudioName:  `Metal`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/Alumina/UI/LiquidAlumina_Pipe_512.LiquidAlumina_Pipe_512`,
		MPingColor:              `(R=0.000000,G=0.000000,B=0.000000,A=1.000000)`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     10,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/Alumina/UI/LiquidAlumina_Pipe_256.LiquidAlumina_Pipe_256`,
		MStackSize:              resource.Fluid,
	}

	OreBauxite = FGResourceDescriptor{
		ClassName:               "Desc_OreBauxite_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription:            `Bauxite is used to produce Alumina, which can be further refined into the Aluminum Scrap required to produce Aluminum Ingots.`,
		MDisplayName:            `Bauxite`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName:  `Metal`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/RawResources/Nodes/UI/Bauxite_256.Bauxite_256`,
		MPingColor:              `(R=1.000000,G=0.000000,B=0.000000,A=1.000000)`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         true,
		MResourceSinkPoints:     8,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/RawResources/Nodes/UI/Bauxite_64.Bauxite_64`,
		MStackSize:              resource.Medium,
	}

	OreCopper = FGResourceDescriptor{
		ClassName:               "Desc_OreCopper_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription: `Used for crafting.
Basic resource mainly used for electricity.`,
		MDisplayName:           `Copper Ore`,
		MEnergyValue:           0.000000,
		MFluidColor:            `(B=0,G=0,R=0,A=0)`,
		MForm:                  resource.Solid,
		MGasColor:              `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName: `Metal`,
		MPersistentBigIcon:     `Texture2D /Game/FactoryGame/Resource/RawResources/Nodes/UI/IconDesc_copper_new_256.IconDesc_copper_new_256`,
		MPingColor:             `(R=0.050000,G=0.500000,B=0.150000,A=1.000000)`,
		MRadioactiveDecay:      0.000000,
		MRememberPickUp:        true,
		MResourceSinkPoints:    3,
		MSmallIcon:             `Texture2D /Game/FactoryGame/Resource/RawResources/Nodes/UI/IconDesc_copper_new_64.IconDesc_copper_new_64`,
		MStackSize:             resource.Medium,
	}

	OreGold = FGResourceDescriptor{
		ClassName:               "Desc_OreGold_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription:            `Caterium Ore is smelted into Caterium Ingots. Caterium Ingots are mostly used for advanced electronics.`,
		MDisplayName:            `Caterium Ore`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName:  `Metal`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/RawResources/Nodes/UI/CateriumOre_256.CateriumOre_256`,
		MPingColor:              `(R=1.000000,G=0.622325,B=0.000000,A=1.000000)`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         true,
		MResourceSinkPoints:     7,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/RawResources/Nodes/UI/CateriumOre_64.CateriumOre_64`,
		MStackSize:              resource.Medium,
	}

	OreIron = FGResourceDescriptor{
		ClassName:               "Desc_OreIron_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription: `Used for crafting.
The most essential basic resource.`,
		MDisplayName:           `Iron Ore`,
		MEnergyValue:           0.000000,
		MFluidColor:            `(B=0,G=0,R=0,A=0)`,
		MForm:                  resource.Solid,
		MGasColor:              `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName: `Metal`,
		MPersistentBigIcon:     `Texture2D /Game/FactoryGame/Resource/RawResources/Nodes/UI/IconDesc_iron_new_256.IconDesc_iron_new_256`,
		MPingColor:             `(R=0.460000,G=0.291200,B=0.242434,A=1.000000)`,
		MRadioactiveDecay:      0.000000,
		MRememberPickUp:        true,
		MResourceSinkPoints:    1,
		MSmallIcon:             `Texture2D /Game/FactoryGame/Resource/RawResources/Nodes/UI/IconDesc_iron_new_64.IconDesc_iron_new_64`,
		MStackSize:             resource.Medium,
	}

	OreUranium = FGResourceDescriptor{
		ClassName:               "Desc_OreUranium_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription: `Uranium is a radioactive element. 
Used to produce Encased Uranium Cells for Uranium Fuel Rods.

Caution: Moderately Radioactive.`,
		MDisplayName:           `Uranium`,
		MEnergyValue:           0.000000,
		MFluidColor:            `(B=0,G=0,R=0,A=0)`,
		MForm:                  resource.Solid,
		MGasColor:              `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName: `Metal`,
		MPersistentBigIcon:     `Texture2D /Game/FactoryGame/Resource/RawResources/OreUranium/UI/UraniumOre_256.UraniumOre_256`,
		MPingColor:             `(R=0.000000,G=1.000000,B=0.040609,A=1.000000)`,
		MRadioactiveDecay:      15.000000,
		MRememberPickUp:        true,
		MResourceSinkPoints:    35,
		MSmallIcon:             `Texture2D /Game/FactoryGame/Resource/RawResources/OreUranium/UI/UraniumOre_64.UraniumOre_64`,
		MStackSize:             resource.Medium,
	}

	RawQuartz = FGResourceDescriptor{
		ClassName:               "Desc_RawQuartz_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription:            `Raw Quartz is processed into Quartz Crystals which are used for radar and quantum technology.`,
		MDisplayName:            `Raw Quartz`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName:  `Metal`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/Parts/QuartzCrystal/UI/IconDesc_QuartzCrystal_256.IconDesc_QuartzCrystal_256`,
		MPingColor:              `(R=0.983632,G=1.000000,B=0.765000,A=1.000000)`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         true,
		MResourceSinkPoints:     15,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/Parts/QuartzCrystal/UI/IconDesc_QuartzCrystal_64.IconDesc_QuartzCrystal_64`,
		MStackSize:              resource.Medium,
	}

	Stone = FGResourceDescriptor{
		ClassName:               "Desc_Stone_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription: `Used for crafting.
Basic resource mainly used for stable foundations.`,
		MDisplayName:           `Limestone`,
		MEnergyValue:           0.000000,
		MFluidColor:            `(B=0,G=0,R=0,A=0)`,
		MForm:                  resource.Solid,
		MGasColor:              `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName: `Metal`,
		MPersistentBigIcon:     `Texture2D /Game/FactoryGame/Resource/RawResources/Stone/UI/Stone_256.Stone_256`,
		MPingColor:             `(R=0.241195,G=0.318507,B=0.415000,A=1.000000)`,
		MRadioactiveDecay:      0.000000,
		MRememberPickUp:        true,
		MResourceSinkPoints:    2,
		MSmallIcon:             `Texture2D /Game/FactoryGame/Resource/RawResources/Stone/UI/Stone_64.Stone_64`,
		MStackSize:             resource.Medium,
	}

	Sulfur = FGResourceDescriptor{
		ClassName:               "Desc_Sulfur_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription:            `Sulfur is primarily used for Black Powder.`,
		MDisplayName:            `Sulfur`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Solid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName:  `Metal`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/RawResources/Sulfur/UI/Sulfur_256.Sulfur_256`,
		MPingColor:              `(R=1.000000,G=0.956156,B=0.260000,A=1.000000)`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         true,
		MResourceSinkPoints:     11,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/RawResources/Sulfur/UI/Sulfur_64.Sulfur_64`,
		MStackSize:              resource.Medium,
	}

	Water = FGResourceDescriptor{
		ClassName:               "Desc_Water_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCollectSpeedMultiplier: 1.000000,
		MDecalSize:              200.000000,
		MDescription:            `It's water.`,
		MDisplayName:            `Water`,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=212,G=176,R=122,A=255)`,
		MForm:                   resource.Liquid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MManualMiningAudioName:  `Metal`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Resource/RawResources/Water/UI/LiquidWater_Pipe_512.LiquidWater_Pipe_512`,
		MPingColor:              `(R=0.000000,G=0.000000,B=0.000000,A=0.000000)`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MResourceSinkPoints:     5,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Resource/RawResources/Water/UI/LiquidWater_Pipe_256.LiquidWater_Pipe_256`,
		MStackSize:              resource.Fluid,
	}
)

func GetByClassName(className string) (*FGResourceDescriptor, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGResourceDescriptor with name %s", className)
}

var classNameToVar = map[string]*FGResourceDescriptor{
	"Desc_Coal_C":        &Coal,
	"Desc_LiquidOil_C":   &LiquidOil,
	"Desc_NitrogenGas_C": &NitrogenGas,
	"Desc_OreBauxite_C":  &OreBauxite,
	"Desc_OreCopper_C":   &OreCopper,
	"Desc_OreGold_C":     &OreGold,
	"Desc_OreIron_C":     &OreIron,
	"Desc_OreUranium_C":  &OreUranium,
	"Desc_RawQuartz_C":   &RawQuartz,
	"Desc_Stone_C":       &Stone,
	"Desc_Sulfur_C":      &Sulfur,
	"Desc_Water_C":       &Water,
}