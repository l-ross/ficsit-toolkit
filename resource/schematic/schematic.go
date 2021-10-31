package Schematic

import (
	"fmt"
)

type FGSchematic struct {
	ClassName               string
	FullName                string
	MCost                   string
	MDescription            string
	MDisplayName            string
	MIncludeInBuilds        string
	MMenuPriority           float64
	MRelevantEvents         string
	MRelevantShopSchematics string
	MSchematicDependencies  []struct {
		Class                              string
		MGamePhase                         string
		MRequireAllSchematicsToBePurchased bool
		MSchematics                        string
	}

	MSchematicIcon      string
	MSmallSchematicIcon string
	MSubCategories      string
	MTechTier           int
	MTimeToComplete     float64
	MType               string
	MUnlocks            []struct {
		Class                         string
		MItemsToGive                  string
		MNumArmEquipmentSlotsToUnlock int
		MNumInventorySlotsToUnlock    int
		MRecipes                      string
		MResourcePairsToAddToScanner  string
		MResourcesToAddToScanner      string
		MSchematics                   string
	}
}

var (
	ACarapace0 = FGSchematic{
		ClassName:               "Research_ACarapace_0_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_ACarapace_0.Research_ACarapace_0_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AnimalParts/Desc_HogParts.Desc_HogParts_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Alien Carapace`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/AnimalParts/UI/IconDesc_HogPart_64.IconDesc_HogPart_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	ACarapace1 = FGSchematic{
		ClassName:               "Research_ACarapace_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_ACarapace_1.Research_ACarapace_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AnimalParts/Desc_HogParts.Desc_HogParts_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Structural Analysis`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Biomass_Final_64.IconDesc_Biomass_Final_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Biomass_AlienCarapace.Recipe_Biomass_AlienCarapace_C"')`,
			},
		},
	}

	ACarapace2 = FGSchematic{
		ClassName:               "Research_ACarapace_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_ACarapace_2.Research_ACarapace_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronScrew/Desc_IronScrew.Desc_IronScrew_C"',Amount=500))`,
		MDescription:            ``,
		MDisplayName:            `Rebar Gun`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/NailGun/UI/RebarGun_64.RebarGun_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_RebarGun.Recipe_RebarGun_C"')`,
			},
		},
	}

	ACarapace21 = FGSchematic{
		ClassName:               "Research_ACarapace_2_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_ACarapace_2_1.Research_ACarapace_2_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Spiked Rebars`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/SpikedRebar/UI/IconDesc_SpikedRebar_64.IconDesc_SpikedRebar_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_SpikedRebar.Recipe_SpikedRebar_C"')`,
			},
		},
	}

	ACarapace3 = FGSchematic{
		ClassName:               "Research_ACarapace_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_ACarapace_3.Research_ACarapace_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AnimalParts/Desc_HogParts.Desc_HogParts_C"',Amount=5),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=1000))`,
		MDescription:            ``,
		MDisplayName:            `Expanded Toolbelt`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/Assets/Shared/ThumbsUp_64.ThumbsUp_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                         `BP_UnlockArmEquipmentSlot_C`,
				MNumArmEquipmentSlotsToUnlock: 1,
			},
		},
	}

	AILimiter = FGSchematic{
		ClassName:               "ResourceSink_AILimiter_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier2/ResourceSink_AILimiter.ResourceSink_AILimiter_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `AI Limiter`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_1.Research_Caterium_4_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/AIlimiter/UI/IconDesc_AILimiter_256.IconDesc_AILimiter_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Electronics.SC_RSS_Electronics_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CircuitBoardHighSpeed/Desc_CircuitBoardHighSpeed.Desc_CircuitBoardHighSpeed_C"',Amount=100))`,
			},
		},
	}

	AOrganisms1 = FGSchematic{
		ClassName:               "Research_AOrganisms_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_AOrganisms_1.Research_AOrganisms_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/Desc_CrystalOscillator.Desc_CrystalOscillator_C"',Amount=5),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Stator/Desc_Stator.Desc_Stator_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/GemstoneScanner/BP_EquipmentDescriptorObjectScanner.BP_EquipmentDescriptorObjectScanner_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Object Scanner Improvements`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/InGame/-Shared/Mam_Key_64.Mam_Key_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	AOrganisms2 = FGSchematic{
		ClassName:               "Research_AOrganisms_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_AOrganisms_2.Research_AOrganisms_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AnimalParts/Desc_SpitterParts.Desc_SpitterParts_C"',Amount=5),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AnimalParts/Desc_HogParts.Desc_HogParts_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Hostile Organism Detection`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Equipment/ObjectScanner/Icons/Monsters_256.Monsters_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	AOrgans0 = FGSchematic{
		ClassName:               "Research_AOrgans_0_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_AOrgans_0.Research_AOrgans_0_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AnimalParts/Desc_SpitterParts.Desc_SpitterParts_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Alien Organs`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/AnimalParts/UI/IconDesc_SpitterPart_64.IconDesc_SpitterPart_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	AOrgans1 = FGSchematic{
		ClassName:               "Research_AOrgans_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_AOrgans_1.Research_AOrgans_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AnimalParts/Desc_SpitterParts.Desc_SpitterParts_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Organic Properties`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Biomass_Final_64.IconDesc_Biomass_Final_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Biomass_AlienOrgans.Recipe_Biomass_AlienOrgans_C"')`,
			},
		},
	}

	AOrgans2 = FGSchematic{
		ClassName:               "Research_AOrgans_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_AOrgans_2.Research_AOrgans_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AnimalParts/Desc_SpitterParts.Desc_SpitterParts_C"',Amount=5),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_Mycelia.Desc_Mycelia_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrame/Desc_ModularFrame.Desc_ModularFrame_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Medicinal Inhaler`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/Medkit/UI/Inhaler_64.Inhaler_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_MedicinalInhalerAlienOrgans.Recipe_MedicinalInhalerAlienOrgans_C"')`,
			},
		},
	}

	AOrgans3 = FGSchematic{
		ClassName:               "Research_AOrgans_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_AOrgans_3.Research_AOrgans_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AnimalParts/Desc_SpitterParts.Desc_SpitterParts_C"',Amount=5),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=2000))`,
		MDescription:            ``,
		MDisplayName:            `Inflated Pocket Dimension`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/Assets/Shared/ThumbsUp_64.ThumbsUp_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 6,
			},
		},
	}

	AdvancedAmmoPack = FGSchematic{
		ClassName:               "ResourceSink_AdvancedAmmoPack_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_AdvancedAmmoPack.ResourceSink_AdvancedAmmoPack_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Advanced Ammo Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_4_2_1.Research_Sulfur_4_2_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/CartridgeStandard/UI/Rifle_Magazine_256.Rifle_Magazine_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CartridgeStandard/Desc_CartridgeStandard.Desc_CartridgeStandard_C"',Amount=25))`,
			},
		},
	}

	AlcladSheet = FGSchematic{
		ClassName:               "ResourceSink_AlcladSheet_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier7/ResourceSink_AlcladSheet.ResourceSink_AlcladSheet_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Aluminum Alclad Sheet`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/AluminumPlate/UI/IconDesc_AluminiumSheet_256.IconDesc_AluminiumSheet_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `Texture2D /Game/FactoryGame/Resource/Parts/AluminumPlate/UI/IconDesc_AluminiumSheet_64.IconDesc_AluminiumSheet_64`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumPlate/Desc_AluminumPlate.Desc_AluminumPlate_C"',Amount=100))`,
			},
		},
	}

	AlternateAdheredIronPlate = FGSchematic{
		ClassName:               "Schematic_Alternate_AdheredIronPlate_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_AdheredIronPlate.Schematic_Alternate_AdheredIronPlate_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Adhered Iron Plate`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_AdheredIronPlate.Recipe_Alternate_AdheredIronPlate_C"')`,
			},
		},
	}

	AlternateAlcladCasing = FGSchematic{
		ClassName:               "Schematic_Alternate_AlcladCasing_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_AlcladCasing.Schematic_Alternate_AlcladCasing_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Alclad Casing`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_AlcladCasing.Recipe_Alternate_AlcladCasing_C"')`,
			},
		},
	}

	AlternateAutomatedMiner = FGSchematic{
		ClassName:               "Schematic_Alternate_AutomatedMiner_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_AutomatedMiner.Schematic_Alternate_AutomatedMiner_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Automated Miner`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-2.Schematic_5-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_AutomatedMiner.Recipe_Alternate_AutomatedMiner_C"')`,
			},
		},
	}

	AlternateBeacon1 = FGSchematic{
		ClassName:               "Schematic_Alternate_Beacon1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Beacon1.Schematic_Alternate_Beacon1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Crystal Beacon`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_1.Research_Quartz_1_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Beacon_1.Recipe_Alternate_Beacon_1_C"')`,
			},
		},
	}

	AlternateBoltedFrame = FGSchematic{
		ClassName:               "Schematic_Alternate_BoltedFrame_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_BoltedFrame.Schematic_Alternate_BoltedFrame_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Bolted Frame`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_2-1.Schematic_2-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_BoltedFrame.Recipe_Alternate_BoltedFrame_C"')`,
			},
		},
	}

	AlternateCable1 = FGSchematic{
		ClassName:               "Schematic_Alternate_Cable1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Cable1.Schematic_Alternate_Cable1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Insulated Cable`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Cable_1.Recipe_Alternate_Cable_1_C"')`,
			},
		},
	}

	AlternateCable2 = FGSchematic{
		ClassName:               "Schematic_Alternate_Cable2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Cable2.Schematic_Alternate_Cable2_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Quickwire Cable`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_1.Research_Caterium_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Cable_2.Recipe_Alternate_Cable_2_C"')`,
			},
		},
	}

	AlternateCircuitBoard1 = FGSchematic{
		ClassName:               "Schematic_Alternate_CircuitBoard1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_CircuitBoard1.Schematic_Alternate_CircuitBoard1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Silicon Circuit Board`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_0.Research_Quartz_0_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_CircuitBoard_1.Recipe_Alternate_CircuitBoard_1_C"')`,
			},
		},
	}

	AlternateCircuitBoard2 = FGSchematic{
		ClassName:               "Schematic_Alternate_CircuitBoard2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_CircuitBoard2.Schematic_Alternate_CircuitBoard2_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Caterium Circuit Board`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_1.Research_Caterium_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_CircuitBoard_2.Recipe_Alternate_CircuitBoard_2_C"')`,
			},
		},
	}

	AlternateClassicBattery = FGSchematic{
		ClassName:               "Schematic_Alternate_ClassicBattery_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_ClassicBattery.Schematic_Alternate_ClassicBattery_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Classic Battery`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-4.Schematic_7-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_ClassicBattery.Recipe_Alternate_ClassicBattery_C"')`,
			},
		},
	}

	AlternateCoal1 = FGSchematic{
		ClassName:               "Schematic_Alternate_Coal1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Coal1.Schematic_Alternate_Coal1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Charcoal`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           3,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Coal_1.Recipe_Alternate_Coal_1_C"')`,
			},
		},
	}

	AlternateCoal2 = FGSchematic{
		ClassName:               "Schematic_Alternate_Coal2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Coal2.Schematic_Alternate_Coal2_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Biocoal`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           3,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Coal_2.Recipe_Alternate_Coal_2_C"')`,
			},
		},
	}

	AlternateCoatedCable = FGSchematic{
		ClassName:               "Schematic_Alternate_CoatedCable_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_CoatedCable.Schematic_Alternate_CoatedCable_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Coated Cable`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_CoatedCable.Recipe_Alternate_CoatedCable_C"')`,
			},
		},
	}

	AlternateCoatedIronCanister = FGSchematic{
		ClassName:               "Schematic_Alternate_CoatedIronCanister_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_CoatedIronCanister.Schematic_Alternate_CoatedIronCanister_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Coated Iron Canister`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-4.Schematic_5-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_CoatedIronCanister.Recipe_Alternate_CoatedIronCanister_C"')`,
			},
		},
	}

	AlternateCoatedIronPlate = FGSchematic{
		ClassName:               "Schematic_Alternate_CoatedIronPlate_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_CoatedIronPlate.Schematic_Alternate_CoatedIronPlate_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Coated Iron Plate`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_CoatedIronPlate.Recipe_Alternate_CoatedIronPlate_C"')`,
			},
		},
	}

	AlternateCokeSteelIngot = FGSchematic{
		ClassName:               "Schematic_Alternate_CokeSteelIngot_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_CokeSteelIngot.Schematic_Alternate_CokeSteelIngot_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Coke Steel Ingot`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_CokeSteelIngot.Recipe_Alternate_CokeSteelIngot_C"')`,
			},
		},
	}

	AlternateComputer1 = FGSchematic{
		ClassName:               "Schematic_Alternate_Computer1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Computer1.Schematic_Alternate_Computer1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Caterium Computer`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-2.Schematic_5-2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_1.Research_Caterium_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Computer_1.Recipe_Alternate_Computer_1_C"')`,
			},
		},
	}

	AlternateComputer2 = FGSchematic{
		ClassName:               "Schematic_Alternate_Computer2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Computer2.Schematic_Alternate_Computer2_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Crystal Computer`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-2.Schematic_5-2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_1.Research_Quartz_1_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Computer_2.Recipe_Alternate_Computer_2_C"')`,
			},
		},
	}

	AlternateConcrete = FGSchematic{
		ClassName:               "Schematic_Alternate_Concrete_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Concrete.Schematic_Alternate_Concrete_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Fine Concrete`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_0.Research_Quartz_0_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Concrete.Recipe_Alternate_Concrete_C"')`,
			},
		},
	}

	AlternateCoolingDevice = FGSchematic{
		ClassName:               "Schematic_Alternate_CoolingDevice_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_CoolingDevice.Schematic_Alternate_CoolingDevice_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Cooling Device`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-2.Schematic_8-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_CoolingDevice.Recipe_Alternate_CoolingDevice_C"')`,
			},
		},
	}

	AlternateCopperAlloyIngot = FGSchematic{
		ClassName:               "Schematic_Alternate_CopperAlloyIngot_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_CopperAlloyIngot.Schematic_Alternate_CopperAlloyIngot_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Copper Alloy Ingot`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               0,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_CopperAlloyIngot.Recipe_Alternate_CopperAlloyIngot_C"')`,
			},
		},
	}

	AlternateCopperRotor = FGSchematic{
		ClassName:               "Schematic_Alternate_CopperRotor_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_CopperRotor.Schematic_Alternate_CopperRotor_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Copper Rotor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_2-1.Schematic_2-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_CopperRotor.Recipe_Alternate_CopperRotor_C"')`,
			},
		},
	}

	AlternateCrystalOscillator = FGSchematic{
		ClassName:               "Schematic_Alternate_CrystalOscillator_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_CrystalOscillator.Schematic_Alternate_CrystalOscillator_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Insulated Crystal Oscillator`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_2.Research_Quartz_2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_3.Research_Caterium_3_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_CrystalOscillator.Recipe_Alternate_CrystalOscillator_C"')`,
			},
		},
	}

	AlternateDilutedFuel = FGSchematic{
		ClassName:               "Schematic_Alternate_DilutedFuel_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_DilutedFuel.Schematic_Alternate_DilutedFuel_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Diluted Fuel`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_DilutedFuel.Recipe_Alternate_DilutedFuel_C"')`,
			},
		},
	}

	AlternateDilutedPackagedFuel = FGSchematic{
		ClassName:               "Schematic_Alternate_DilutedPackagedFuel_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_DilutedPackagedFuel.Schematic_Alternate_DilutedPackagedFuel_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Diluted Packaged Fuel`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_DilutedPackagedFuel.Recipe_Alternate_DilutedPackagedFuel_C"')`,
			},
		},
	}

	AlternateElectricMotor = FGSchematic{
		ClassName:               "Schematic_Alternate_ElectricMotor_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_ElectricMotor.Schematic_Alternate_ElectricMotor_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Electric Motor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-4.Schematic_7-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_ElectricMotor.Recipe_Alternate_ElectricMotor_C"')`,
			},
		},
	}

	AlternateElectroAluminumScrap = FGSchematic{
		ClassName:               "Schematic_Alternate_ElectroAluminumScrap_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_ElectroAluminumScrap.Schematic_Alternate_ElectroAluminumScrap_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Electrode - Aluminum Scrap`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_ElectroAluminumScrap.Recipe_Alternate_ElectroAluminumScrap_C"')`,
			},
		},
	}

	AlternateElectrodeCircuitBoard = FGSchematic{
		ClassName:               "Schematic_Alternate_ElectrodeCircuitBoard_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_ElectrodeCircuitBoard.Schematic_Alternate_ElectrodeCircuitBoard_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Electrode Circuit Board`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_ElectrodeCircuitBoard.Recipe_Alternate_ElectrodeCircuitBoard_C"')`,
			},
		},
	}

	AlternateElectromagneticControlRod1 = FGSchematic{
		ClassName:               "Schematic_Alternate_ElectromagneticControlRod1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_ElectromagneticControlRod1.Schematic_Alternate_ElectromagneticControlRod1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Electromagnetic Connection Rod`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-1.Schematic_8-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_1.Research_Caterium_4_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           7,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_ElectromagneticControlRod_1.Recipe_Alternate_ElectromagneticControlRod_1_C"')`,
			},
		},
	}

	AlternateEnrichedCoal = FGSchematic{
		ClassName:               "Schematic_Alternate_EnrichedCoal_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_EnrichedCoal.Schematic_Alternate_EnrichedCoal_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Compacted Coal`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_0.Research_Sulfur_0_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_EnrichedCoal.Recipe_Alternate_EnrichedCoal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Coal_1.Recipe_Alternate_Coal_1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Coal_2.Recipe_Alternate_Coal_2_C"')`,
			},
		},
	}

	AlternateFertileUranium = FGSchematic{
		ClassName:               "Schematic_Alternate_FertileUranium_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_FertileUranium.Schematic_Alternate_FertileUranium_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Fertile Uranium`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-5.Schematic_8-5_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_FertileUranium.Recipe_Alternate_FertileUranium_C"')`,
			},
		},
	}

	AlternateFlexibleFramework = FGSchematic{
		ClassName:               "Schematic_Alternate_FlexibleFramework_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_FlexibleFramework.Schematic_Alternate_FlexibleFramework_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Flexible Framework `,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_FlexibleFramework.Recipe_Alternate_FlexibleFramework_C"')`,
			},
		},
	}

	AlternateFusedWire = FGSchematic{
		ClassName:               "Schematic_Alternate_FusedWire_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_FusedWire.Schematic_Alternate_FusedWire_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Fused Wire`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_0.Research_Caterium_0_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_FusedWire.Recipe_Alternate_FusedWire_C"')`,
			},
		},
	}

	AlternateGunpowder1 = FGSchematic{
		ClassName:               "Schematic_Alternate_Gunpowder1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Gunpowder1.Schematic_Alternate_Gunpowder1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Fine Black Powder`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_EnrichedCoal.Schematic_Alternate_EnrichedCoal_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Gunpowder_1.Recipe_Alternate_Gunpowder_1_C"')`,
			},
		},
	}

	AlternateHeatFusedFrame = FGSchematic{
		ClassName:               "Schematic_Alternate_HeatFusedFrame_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_HeatFusedFrame.Schematic_Alternate_HeatFusedFrame_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Heat-Fused Frame`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-5.Schematic_8-5_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_HeatFusedFrame.Recipe_Alternate_HeatFusedFrame_C"')`,
			},
		},
	}

	AlternateHeatSink1 = FGSchematic{
		ClassName:               "Schematic_Alternate_HeatSink1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_HeatSink1.Schematic_Alternate_HeatSink1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Heat Exchanger`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-2.Schematic_8-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           7,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_HeatSink_1.Recipe_Alternate_HeatSink_1_C"')`,
			},
		},
	}

	AlternateHeavyFlexibleFrame = FGSchematic{
		ClassName:               "Schematic_Alternate_HeavyFlexibleFrame_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_HeavyFlexibleFrame.Schematic_Alternate_HeavyFlexibleFrame_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Heavy Flexible Frame`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-2.Schematic_5-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_HeavyFlexibleFrame.Recipe_Alternate_HeavyFlexibleFrame_C"')`,
			},
		},
	}

	AlternateHeavyModularFrame = FGSchematic{
		ClassName:               "Schematic_Alternate_HeavyModularFrame_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_HeavyModularFrame.Schematic_Alternate_HeavyModularFrame_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Heavy Encased Frame`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-2.Schematic_5-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_ModularFrameHeavy.Recipe_Alternate_ModularFrameHeavy_C"')`,
			},
		},
	}

	AlternateHeavyOilResidue = FGSchematic{
		ClassName:               "Schematic_Alternate_HeavyOilResidue_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_HeavyOilResidue.Schematic_Alternate_HeavyOilResidue_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Heavy Oil Residue`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_HeavyOilResidue.Recipe_Alternate_HeavyOilResidue_C"')`,
			},
		},
	}

	AlternateHighSpeedConnector = FGSchematic{
		ClassName:               "Schematic_Alternate_HighSpeedConnector_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_HighSpeedConnector.Schematic_Alternate_HighSpeedConnector_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Silicon High-Speed Connector`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_5.Research_Caterium_5_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_0.Research_Quartz_0_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_HighSpeedConnector.Recipe_Alternate_HighSpeedConnector_C"')`,
			},
		},
	}

	AlternateHighSpeedWiring = FGSchematic{
		ClassName:               "Schematic_Alternate_HighSpeedWiring_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_HighSpeedWiring.Schematic_Alternate_HighSpeedWiring_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: High-Speed Wiring`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_4-1.Schematic_4-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_1.Research_Caterium_4_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_HighSpeedWiring.Recipe_Alternate_HighSpeedWiring_C"')`,
			},
		},
	}

	AlternateIngotIron = FGSchematic{
		ClassName:               "Schematic_Alternate_IngotIron_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_IngotIron.Schematic_Alternate_IngotIron_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Iron Alloy Ingot`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_IngotIron.Recipe_Alternate_IngotIron_C"')`,
			},
		},
	}

	AlternateIngotSteel1 = FGSchematic{
		ClassName:               "Schematic_Alternate_IngotSteel1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_IngotSteel1.Schematic_Alternate_IngotSteel1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Solid Steel Ingot`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_IngotSteel_1.Recipe_Alternate_IngotSteel_1_C"')`,
			},
		},
	}

	AlternateIngotSteel2 = FGSchematic{
		ClassName:               "Schematic_Alternate_IngotSteel2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_IngotSteel2.Schematic_Alternate_IngotSteel2_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Compacted Steel Ingot`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_EnrichedCoal.Schematic_Alternate_EnrichedCoal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_IngotSteel_2.Recipe_Alternate_IngotSteel_2_C"')`,
			},
		},
	}

	AlternateInstantPlutoniumCell = FGSchematic{
		ClassName:               "Schematic_Alternate_InstantPlutoniumCell_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_InstantPlutoniumCell.Schematic_Alternate_InstantPlutoniumCell_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Instant Plutonium Cell`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-5.Schematic_8-5_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_InstantPlutoniumCell.Recipe_Alternate_InstantPlutoniumCell_C"')`,
			},
		},
	}

	AlternateInstantScrap = FGSchematic{
		ClassName:               "Schematic_Alternate_InstantScrap_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_InstantScrap.Schematic_Alternate_InstantScrap_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Instant Scrap`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-4.Schematic_7-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_InstantScrap.Recipe_Alternate_InstantScrap_C"')`,
			},
		},
	}

	AlternateInventorySlots1 = FGSchematic{
		ClassName:               "Schematic_Alternate_InventorySlots1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Upgrades/Schematic_Alternate_InventorySlots1.Schematic_Alternate_InventorySlots1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Inflated Pocket Dimension`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               2,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 6,
			},
		},
	}

	AlternateInventorySlots2 = FGSchematic{
		ClassName:               "Schematic_Alternate_InventorySlots2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Upgrades/Schematic_Alternate_InventorySlots2.Schematic_Alternate_InventorySlots2_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Inflated Pocket Dimension`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-2.Schematic_5-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 6,
			},
		},
	}

	AlternateModularFrame = FGSchematic{
		ClassName:               "Schematic_Alternate_ModularFrame_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_ModularFrame.Schematic_Alternate_ModularFrame_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Steeled Frame`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_ModularFrame.Recipe_Alternate_ModularFrame_C"')`,
			},
		},
	}

	AlternateMotor1 = FGSchematic{
		ClassName:               "Schematic_Alternate_Motor1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Motor1.Schematic_Alternate_Motor1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Rigour Motor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_4-1.Schematic_4-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_1.Research_Quartz_1_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Motor_1.Recipe_Alternate_Motor_1_C"')`,
			},
		},
	}

	AlternateNobelisk1 = FGSchematic{
		ClassName:               "Schematic_Alternate_Nobelisk1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Nobelisk1.Schematic_Alternate_Nobelisk1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Seismic Nobelisk`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_3_2_1.Research_Sulfur_3_2_1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_1.Research_Quartz_1_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Nobelisk_1.Recipe_Alternate_Nobelisk_1_C"')`,
			},
		},
	}

	AlternateNuclearFuelRod1 = FGSchematic{
		ClassName:               "Schematic_Alternate_NuclearFuelRod1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_NuclearFuelRod1.Schematic_Alternate_NuclearFuelRod1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Uranium Fuel Unit`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-1.Schematic_8-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_1.Research_Quartz_1_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           7,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_NuclearFuelRod_1.Recipe_Alternate_NuclearFuelRod_1_C"')`,
			},
		},
	}

	AlternateOCSupercomputer = FGSchematic{
		ClassName:               "Schematic_Alternate_OCSupercomputer_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_OCSupercomputer.Schematic_Alternate_OCSupercomputer_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: OC Supercomputer`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-4.Schematic_7-4_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-2.Schematic_8-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_OCSupercomputer.Recipe_Alternate_OCSupercomputer_C"')`,
			},
		},
	}

	AlternatePlastic1 = FGSchematic{
		ClassName:               "Schematic_Alternate_Plastic1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Plastic1.Schematic_Alternate_Plastic1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Recycled Plastic`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           5,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Plastic_1.Recipe_Alternate_Plastic_1_C"')`,
			},
		},
	}

	AlternatePlasticSmartPlating = FGSchematic{
		ClassName:               "Schematic_Alternate_PlasticSmartPlating_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_PlasticSmartPlating.Schematic_Alternate_PlasticSmartPlating_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Plastic Smart Plating`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_PlasticSmartPlating.Recipe_Alternate_PlasticSmartPlating_C"')`,
			},
		},
	}

	AlternatePlutoniumFuelUnit = FGSchematic{
		ClassName:               "Schematic_Alternate_PlutoniumFuelUnit_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_PlutoniumFuelUnit.Schematic_Alternate_PlutoniumFuelUnit_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Plutonium Fuel Unit`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-5.Schematic_8-5_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_PlutoniumFuelUnit.Recipe_Alternate_PlutoniumFuelUnit_C"')`,
			},
		},
	}

	AlternatePolyesterFabric = FGSchematic{
		ClassName:               "Schematic_Alternate_PolyesterFabric_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_PolyesterFabric.Schematic_Alternate_PolyesterFabric_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Polyester Fabric`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Mycelia_RS/Research_Mycelia_2.Research_Mycelia_2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_PolyesterFabric.Recipe_Alternate_PolyesterFabric_C"')`,
			},
		},
	}

	AlternatePolymerResin = FGSchematic{
		ClassName:               "Schematic_Alternate_PolymerResin_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_PolymerResin.Schematic_Alternate_PolymerResin_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Polymer Resin`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_PolymerResin.Recipe_Alternate_PolymerResin_C"')`,
			},
		},
	}

	AlternatePureAluminumIngot = FGSchematic{
		ClassName:               "Schematic_Alternate_PureAluminumIngot_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_PureAluminumIngot.Schematic_Alternate_PureAluminumIngot_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Pure Aluminum Ingot`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_PureAluminumIngot.Recipe_PureAluminumIngot_C"')`,
			},
		},
	}

	AlternatePureCateriumIngot = FGSchematic{
		ClassName:               "Schematic_Alternate_PureCateriumIngot_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_PureCateriumIngot.Schematic_Alternate_PureCateriumIngot_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Pure Caterium Ingot`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_0.Research_Caterium_0_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_PureCateriumIngot.Recipe_Alternate_PureCateriumIngot_C"')`,
			},
		},
	}

	AlternatePureCopperIngot = FGSchematic{
		ClassName:               "Schematic_Alternate_PureCopperIngot_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_PureCopperIngot.Schematic_Alternate_PureCopperIngot_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Pure Copper Ingot`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_PureCopperIngot.Recipe_Alternate_PureCopperIngot_C"')`,
			},
		},
	}

	AlternatePureIronIngot = FGSchematic{
		ClassName:               "Schematic_Alternate_PureIronIngot_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_PureIronIngot.Schematic_Alternate_PureIronIngot_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Pure Iron Ingot`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_PureIronIngot.Recipe_Alternate_PureIronIngot_C"')`,
			},
		},
	}

	AlternatePureQuartzCrystal = FGSchematic{
		ClassName:               "Schematic_Alternate_PureQuartzCrystal_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_PureQuartzCrystal.Schematic_Alternate_PureQuartzCrystal_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Pure Quartz Crystal `,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_1.Research_Quartz_1_1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_PureQuartzCrystal.Recipe_Alternate_PureQuartzCrystal_C"')`,
			},
		},
	}

	AlternateQuickwire = FGSchematic{
		ClassName:               "Schematic_Alternate_Quickwire_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Quickwire.Schematic_Alternate_Quickwire_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Fused Quickwire`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_1.Research_Caterium_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           3,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Quickwire.Recipe_Alternate_Quickwire_C"')`,
			},
		},
	}

	AlternateRadioControlSystem = FGSchematic{
		ClassName:               "Schematic_Alternate_RadioControlSystem_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_RadioControlSystem.Schematic_Alternate_RadioControlSystem_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Radio Control System`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_RadioControlSystem.Recipe_Alternate_RadioControlSystem_C"')`,
			},
		},
	}

	AlternateRadioControlUnit1 = FGSchematic{
		ClassName:               "Schematic_Alternate_RadioControlUnit1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_RadioControlUnit1.Schematic_Alternate_RadioControlUnit1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Radio Connection Unit`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-2.Schematic_8-2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_1.Research_Caterium_4_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           7,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_RadioControlUnit_1.Recipe_Alternate_RadioControlUnit_1_C"')`,
			},
		},
	}

	AlternateRecycledRubber = FGSchematic{
		ClassName:               "Schematic_Alternate_RecycledRubber_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_RecycledRubber.Schematic_Alternate_RecycledRubber_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Recycled Rubber`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_RecycledRubber.Recipe_Alternate_RecycledRubber_C"')`,
			},
		},
	}

	AlternateReinforcedIronPlate1 = FGSchematic{
		ClassName:               "Schematic_Alternate_ReinforcedIronPlate1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_ReinforcedIronPlate1.Schematic_Alternate_ReinforcedIronPlate1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Bolted Iron Plate`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               2,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_ReinforcedIronPlate_1.Recipe_Alternate_ReinforcedIronPlate_1_C"')`,
			},
		},
	}

	AlternateReinforcedIronPlate2 = FGSchematic{
		ClassName:               "Schematic_Alternate_ReinforcedIronPlate2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_ReinforcedIronPlate2.Schematic_Alternate_ReinforcedIronPlate2_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Stitched Iron Plate`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               2,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_ReinforcedIronPlate_2.Recipe_Alternate_ReinforcedIronPlate_2_C"')`,
			},
		},
	}

	AlternateReinforcedSteelPlate = FGSchematic{
		ClassName:               "Schematic_Alternate_ReinforcedSteelPlate_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_ReinforcedSteelPlate.Schematic_Alternate_ReinforcedSteelPlate_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Encased Industrial Pipe`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_4-1.Schematic_4-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_EncasedIndustrialBeam.Recipe_Alternate_EncasedIndustrialBeam_C"')`,
			},
		},
	}

	AlternateRotor = FGSchematic{
		ClassName:               "Schematic_Alternate_Rotor_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Rotor.Schematic_Alternate_Rotor_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Steel Rotor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Rotor.Recipe_Alternate_Rotor_C"')`,
			},
		},
	}

	AlternateRubberConcrete = FGSchematic{
		ClassName:               "Schematic_Alternate_RubberConcrete_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_RubberConcrete.Schematic_Alternate_RubberConcrete_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Rubber Concrete`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_RubberConcrete.Recipe_Alternate_RubberConcrete_C"')`,
			},
		},
	}

	AlternateScrew = FGSchematic{
		ClassName:               "Schematic_Alternate_Screw_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Screw.Schematic_Alternate_Screw_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Cast Screw`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Screw.Recipe_Alternate_Screw_C"')`,
			},
		},
	}

	AlternateScrew2 = FGSchematic{
		ClassName:               "Schematic_Alternate_Screw2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Screw2.Schematic_Alternate_Screw2_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Steel Screw`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Screw_2.Recipe_Alternate_Screw_2_C"')`,
			},
		},
	}

	AlternateSilica = FGSchematic{
		ClassName:               "Schematic_Alternate_Silica_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Silica.Schematic_Alternate_Silica_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Cheap Silica`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_2.Research_Quartz_1_2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Silica.Recipe_Alternate_Silica_C"')`,
			},
		},
	}

	AlternateSloppyAlumina = FGSchematic{
		ClassName:               "Schematic_Alternate_SloppyAlumina_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_SloppyAlumina.Schematic_Alternate_SloppyAlumina_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Sloppy Alumina`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_SloppyAlumina.Recipe_Alternate_SloppyAlumina_C"')`,
			},
		},
	}

	AlternateStator = FGSchematic{
		ClassName:               "Schematic_Alternate_Stator_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Stator.Schematic_Alternate_Stator_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Quickwire Stator`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_4-1.Schematic_4-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_1.Research_Caterium_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           4,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Stator.Recipe_Alternate_Stator_C"')`,
			},
		},
	}

	AlternateSteamedCopperSheet = FGSchematic{
		ClassName:               "Schematic_Alternate_SteamedCopperSheet_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_SteamedCopperSheet.Schematic_Alternate_SteamedCopperSheet_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Steamed Copper Sheet `,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_SteamedCopperSheet.Recipe_Alternate_SteamedCopperSheet_C"')`,
			},
		},
	}

	AlternateSteelCanister = FGSchematic{
		ClassName:               "Schematic_Alternate_SteelCanister_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_SteelCanister.Schematic_Alternate_SteelCanister_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Steel Canister`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-4.Schematic_5-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_SteelCanister.Recipe_Alternate_SteelCanister_C"')`,
			},
		},
	}

	AlternateSteelCoatedPlate = FGSchematic{
		ClassName:               "Schematic_Alternate_SteelCoatedPlate_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_SteelCoatedPlate.Schematic_Alternate_SteelCoatedPlate_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Steel Coated Plate `,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_SteelCoatedPlate.Recipe_Alternate_SteelCoatedPlate_C"')`,
			},
		},
	}

	AlternateSteelRod = FGSchematic{
		ClassName:               "Schematic_Alternate_SteelRod_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_SteelRod.Schematic_Alternate_SteelRod_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Steel Rod`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_SteelRod.Recipe_Alternate_SteelRod_C"')`,
			},
		},
	}

	AlternateSuperStateComputer = FGSchematic{
		ClassName:               "Schematic_Alternate_SuperStateComputer_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_SuperStateComputer.Schematic_Alternate_SuperStateComputer_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Super-State Computer`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-1.Schematic_8-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-4.Schematic_7-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_SuperStateComputer.Recipe_Alternate_SuperStateComputer_C"')`,
			},
		},
	}

	AlternateTurboBlendFuel = FGSchematic{
		ClassName:               "Schematic_Alternate_TurboBlendFuel_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_TurboBlendFuel.Schematic_Alternate_TurboBlendFuel_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Turbo Blend Fuel`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_0.Research_Sulfur_0_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_TurboBlendFuel.Recipe_Alternate_TurboBlendFuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_PackagedTurboFuel.Recipe_PackagedTurboFuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_UnpackageTurboFuel.Recipe_UnpackageTurboFuel_C"')`,
			},
		},
	}

	AlternateTurboFuel = FGSchematic{
		ClassName:               "Schematic_Alternate_TurboFuel_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_TurboFuel.Schematic_Alternate_TurboFuel_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Turbofuel`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_EnrichedCoal.Schematic_Alternate_EnrichedCoal_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           6,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Turbofuel.Recipe_Alternate_Turbofuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_PackagedTurboFuel.Recipe_PackagedTurboFuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_UnpackageTurboFuel.Recipe_UnpackageTurboFuel_C"')`,
			},
		},
	}

	AlternateTurboHeavyFuel = FGSchematic{
		ClassName:               "Schematic_Alternate_TurboHeavyFuel_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_TurboHeavyFuel.Schematic_Alternate_TurboHeavyFuel_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Turbo Heavy Fuel`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_EnrichedCoal.Schematic_Alternate_EnrichedCoal_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_TurboHeavyFuel.Recipe_Alternate_TurboHeavyFuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_PackagedTurboFuel.Recipe_PackagedTurboFuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_UnpackageTurboFuel.Recipe_UnpackageTurboFuel_C"')`,
			},
		},
	}

	AlternateTurboMotor1 = FGSchematic{
		ClassName:               "Schematic_Alternate_TurboMotor1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_TurboMotor1.Schematic_Alternate_TurboMotor1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Turbo Electric Motor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-4.Schematic_8-4_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-1.Schematic_8-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           7,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_TurboMotor_1.Recipe_Alternate_TurboMotor_1_C"')`,
			},
		},
	}

	AlternateTurboPressureMotor = FGSchematic{
		ClassName:               "Schematic_Alternate_TurboPressureMotor_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update4/Schematic_Alternate_TurboPressureMotor.Schematic_Alternate_TurboPressureMotor_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Turbo Pressure Motor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-5.Schematic_8-5_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update4/Recipe_Alternate_TurboPressureMotor.Recipe_Alternate_TurboPressureMotor_C"')`,
			},
		},
	}

	AlternateUraniumCell1 = FGSchematic{
		ClassName:               "Schematic_Alternate_UraniumCell1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_UraniumCell1.Schematic_Alternate_UraniumCell1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Infused Uranium Cell`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-1.Schematic_8-1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_1.Research_Caterium_1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_0.Research_Quartz_0_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_0.Research_Sulfur_0_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           7,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_UraniumCell_1.Recipe_Alternate_UraniumCell_1_C"')`,
			},
		},
	}

	AlternateWetConcrete = FGSchematic{
		ClassName:               "Schematic_Alternate_WetConcrete_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/New_Update3/Schematic_Alternate_WetConcrete.Schematic_Alternate_WetConcrete_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Wet Concrete`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/New_Update3/Recipe_Alternate_WetConcrete.Recipe_Alternate_WetConcrete_C"')`,
			},
		},
	}

	AlternateWire1 = FGSchematic{
		ClassName:               "Schematic_Alternate_Wire1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Wire1.Schematic_Alternate_Wire1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Iron Wire`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Wire_1.Recipe_Alternate_Wire_1_C"')`,
			},
		},
	}

	AlternateWire2 = FGSchematic{
		ClassName:               "Schematic_Alternate_Wire2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Alternate/Parts/Schematic_Alternate_Wire2.Schematic_Alternate_Wire2_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Alternate: Caterium Wire`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_0.Research_Caterium_0_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_MAM.SchematicIcon_MAM"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           3,
		MTimeToComplete:     0.000000,
		MType:               `EST_Alternate`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/AlternateRecipes/Parts/Recipe_Alternate_Wire_2.Recipe_Alternate_Wire_2_C"')`,
			},
		},
	}

	AmmoPack = FGSchematic{
		ClassName:               "ResourceSink_AmmoPack_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_AmmoPack.ResourceSink_AmmoPack_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Ammo Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_ACarapace_2_1.Research_ACarapace_2_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/SpikedRebar/UI/IconDesc_SpikedRebar_256.IconDesc_SpikedRebar_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SpikedRebar/Desc_SpikedRebar.Desc_SpikedRebar_C"',Amount=25))`,
			},
		},
	}

	Battery = FGSchematic{
		ClassName:               "ResourceSink_Battery_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier7/ResourceSink_Battery.ResourceSink_Battery_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Battery`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-2.Schematic_7-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Battery/UI/IconDesc_Battery_256.IconDesc_Battery_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Electronics.SC_RSS_Electronics_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Battery/Desc_Battery.Desc_Battery_C"',Amount=100))`,
			},
		},
	}

	Beacons = FGSchematic{
		ClassName:               "ResourceSink_Beacons_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_Beacons.ResourceSink_Beacons_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Beacons`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_1-3.Schematic_1-3_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/Beacon/UI/IconDesc_Beacon_256.IconDesc_Beacon_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Beacon/BP_EquipmentDescriptorBeacon.BP_EquipmentDescriptorBeacon_C"',Amount=10))`,
			},
		},
	}

	Biofuel = FGSchematic{
		ClassName:               "ResourceSink_Biofuel_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier2/ResourceSink_Biofuel.ResourceSink_Biofuel_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Biofuel`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_2-2.Schematic_2-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/SolidBiofuel/UI/IconDesc_SolidBiofuel_256.IconDesc_SolidBiofuel_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Biomass.SC_RSS_Biomass_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/BioFuel/Desc_Biofuel.Desc_Biofuel_C"',Amount=200))`,
			},
		},
	}

	Biomass = FGSchematic{
		ClassName:               "ResourceSink_Biomass_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier1/ResourceSink_Biomass.ResourceSink_Biomass_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Biomass`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Biomass_Final_256.IconDesc_Biomass_Final_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Biomass.SC_RSS_Biomass_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_GenericBiomass.Desc_GenericBiomass_C"',Amount=200))`,
			},
		},
	}

	BlackPowder = FGSchematic{
		ClassName:               "ResourceSink_BlackPowder_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier3/ResourceSink_BlackPowder.ResourceSink_BlackPowder_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Black Powder`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_1.Research_Sulfur_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GunPowder/UI/IconDesc_Gunpowder_256.IconDesc_Gunpowder_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Minerals.SC_RSS_Minerals_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GunPowder/Desc_Gunpowder.Desc_Gunpowder_C"',Amount=100))`,
			},
		},
	}

	Cable = FGSchematic{
		ClassName:               "ResourceSink_Cable_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier1/ResourceSink_Cable.ResourceSink_Cable_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Cable`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Cable/UI/IconDesc_Cables_256.IconDesc_Cables_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Electronics.SC_RSS_Electronics_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=100))`,
			},
		},
	}

	Caterium0 = FGSchematic{
		ClassName:               "Research_Caterium_0_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_0.Research_Caterium_0_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreGold/Desc_OreGold.Desc_OreGold_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Caterium`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/RawResources/Nodes/UI/CateriumOre_64.CateriumOre_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreGold/Desc_OreGold.Desc_OreGold_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreGold/Desc_OreGold.Desc_OreGold_C"')`,
			},
		},
	}

	Caterium1 = FGSchematic{
		ClassName:               "Research_Caterium_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_1.Research_Caterium_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreGold/Desc_OreGold.Desc_OreGold_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Caterium Ingots`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GoldIngot/UI/IconDesc_CateriumIngot_64.IconDesc_CateriumIngot_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Smelter/Recipe_IngotCaterium.Recipe_IngotCaterium_C"')`,
			},
		},
	}

	Caterium2 = FGSchematic{
		ClassName:               "Research_Caterium_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_2.Research_Caterium_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GoldIngot/Desc_GoldIngot.Desc_GoldIngot_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Quickwire`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/UI/IconDesc_Quickwire_64.IconDesc_Quickwire_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Quickwire.Recipe_Quickwire_C"')`,
			},
		},
	}

	Caterium21 = FGSchematic{
		ClassName:               "Research_Caterium_2_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_2_1.Research_Caterium_2_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/Desc_HighSpeedWire.Desc_HighSpeedWire_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Zipline`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Equipment/Zipline/UI/IconDesc_Zipline_64.IconDesc_Zipline_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_ZipLine.Recipe_ZipLine_C"')`,
			},
		},
	}

	Caterium3 = FGSchematic{
		ClassName:               "Research_Caterium_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_3.Research_Caterium_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/Desc_HighSpeedWire.Desc_HighSpeedWire_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Caterium Electronics`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/InGame/-Shared/Mam_Key_64.Mam_Key_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Caterium31 = FGSchematic{
		ClassName:               "Research_Caterium_3_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_3_1.Research_Caterium_3_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GoldIngot/Desc_GoldIngot.Desc_GoldIngot_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Inflated Pocket Dimension`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/Assets/Shared/ThumbsUp_64.ThumbsUp_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         120.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 6,
			},
		},
	}

	Caterium41 = FGSchematic{
		ClassName:               "Research_Caterium_4_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_1.Research_Caterium_4_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/Desc_HighSpeedWire.Desc_HighSpeedWire_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CopperSheet/Desc_CopperSheet.Desc_CopperSheet_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `AI Limiter`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/AIlimiter/UI/IconDesc_AILimiter_64.IconDesc_AILimiter_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_AILimiter.Recipe_AILimiter_C"')`,
			},
		},
	}

	Caterium411 = FGSchematic{
		ClassName:               "Research_Caterium_4_1_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_1_1.Research_Caterium_4_1_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CircuitBoardHighSpeed/Desc_CircuitBoardHighSpeed.Desc_CircuitBoardHighSpeed_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Smart Splitter`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/CA_SplitterSmart/UI/SmartSplitter_256.SmartSplitter_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorAttachmentSplitterSmart.Recipe_ConveyorAttachmentSplitterSmart_C"')`,
			},
		},
	}

	Caterium412 = FGSchematic{
		ClassName:               "Research_Caterium_4_1_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_1_2.Research_Caterium_4_1_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlate/Desc_SteelPlate.Desc_SteelPlate_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CircuitBoardHighSpeed/Desc_CircuitBoardHighSpeed.Desc_CircuitBoardHighSpeed_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Power Switch`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/PowerSwitch/UI/IconDesc_PowerSwitch_256.IconDesc_PowerSwitch_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerSwitch.Recipe_PowerSwitch_C"')`,
			},
		},
	}

	Caterium42 = FGSchematic{
		ClassName:               "Research_Caterium_4_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_2.Research_Caterium_4_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/Desc_HighSpeedWire.Desc_HighSpeedWire_C"',Amount=300))`,
		MDescription:            ``,
		MDisplayName:            `Power Poles Mk.2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/PowerPoleMk2/UI/PowerPoleMk2_256.PowerPoleMk2_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerPoleMk2.Recipe_PowerPoleMk2_C"')`,
			},
		},
	}

	Caterium43 = FGSchematic{
		ClassName:               "Research_Caterium_4_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_3.Research_Caterium_4_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/Desc_HighSpeedWire.Desc_HighSpeedWire_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrame/Desc_ModularFrame.Desc_ModularFrame_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Blade Runners`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Equipment/JumpingStilts/UI/IconDesc_SprintingStilts_64.IconDesc_SprintingStilts_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_BladeRunners.Recipe_BladeRunners_C"')`,
			},
		},
	}

	Caterium5 = FGSchematic{
		ClassName:               "Research_Caterium_5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_5.Research_Caterium_5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/Desc_HighSpeedWire.Desc_HighSpeedWire_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Plastic/Desc_Plastic.Desc_Plastic_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `High-Speed Connector`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/HighSpeedConnector/UI/IconDesc_HighSpeedConnector_64.IconDesc_HighSpeedConnector_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_HighSpeedConnector.Recipe_HighSpeedConnector_C"')`,
			},
		},
	}

	Caterium61 = FGSchematic{
		ClassName:               "Research_Caterium_6_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_6_1.Research_Caterium_6_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CircuitBoardHighSpeed/Desc_CircuitBoardHighSpeed.Desc_CircuitBoardHighSpeed_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedConnector/Desc_HighSpeedConnector.Desc_HighSpeedConnector_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Computer/Desc_Computer.Desc_Computer_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Supercomputer`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/ComputerSuper/UI/IconDesc_SuperComputer_64.IconDesc_SuperComputer_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_ComputerSuper.Recipe_ComputerSuper_C"')`,
			},
		},
	}

	Caterium62 = FGSchematic{
		ClassName:               "Research_Caterium_6_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_6_2.Research_Caterium_6_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedConnector/Desc_HighSpeedConnector.Desc_HighSpeedConnector_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Power Poles Mk.3`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/PowerPoleMk3/UI/PowerPoleMk3_256.PowerPoleMk3_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         360.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerPoleMk3.Recipe_PowerPoleMk3_C"')`,
			},
		},
	}

	Caterium63 = FGSchematic{
		ClassName:               "Research_Caterium_6_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_6_3.Research_Caterium_6_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedConnector/Desc_HighSpeedConnector.Desc_HighSpeedConnector_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Inflated Pocket Dimension`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/Assets/Shared/ThumbsUp_64.ThumbsUp_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         180.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Caterium71 = FGSchematic{
		ClassName:               "Research_Caterium_7_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_7_1.Research_Caterium_7_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ComputerSuper/Desc_ComputerSuper.Desc_ComputerSuper_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Programmable Splitter`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/CA_SplitterProgrammable/UI/ProgrammableSplitter_256.ProgrammableSplitter_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         480.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorAttachmentSplitterProgrammable.Recipe_ConveyorAttachmentSplitterProgrammable_C"')`,
			},
		},
	}

	Caterium72 = FGSchematic{
		ClassName:               "Research_Caterium_7_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_7_2.Research_Caterium_7_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ComputerSuper/Desc_ComputerSuper.Desc_ComputerSuper_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rubber/Desc_Rubber.Desc_Rubber_C"',Amount=300))`,
		MDescription:            ``,
		MDisplayName:            `Geothermal Generator`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/GeneratorGeoThermal/UI/GeoThermalPowerGenerator_256.GeoThermalPowerGenerator_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         480.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_GeneratorGeoThermal.Recipe_GeneratorGeoThermal_C"')`,
			},
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Geyser/Desc_Geyser.Desc_Geyser_C"',ResourceNodeType=Geyser))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Geyser/Desc_Geyser.Desc_Geyser_C"')`,
			},
		},
	}

	CeilingLight = FGSchematic{
		ClassName:               "ResourceSink_CeilingLight_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CeilingLight.ResourceSink_CeilingLight_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            `Ceiling Lights are the perfect choice for indoor illumination. Simply attach them to the ceiling, hook up power, and watch your conveyor belts shine bright even in the darkest of factories.`,
		MDisplayName:            `Indoor Lighting`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StreetLight.ResourceSink_StreetLight_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_LightTower.ResourceSink_LightTower_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPoles.ResourceSink_WallPowerPoles_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPolesMK2.ResourceSink_WallPowerPolesMK2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPolesMK3.ResourceSink_WallPowerPolesMK3_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/CeilingLight/UI/IconDesc_CeilingLight_256.IconDesc_CeilingLight_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_CeilingLight.Recipe_CeilingLight_C"')`,
			},
		},
	}

	CircuitBoard = FGSchematic{
		ClassName:               "ResourceSink_CircuitBoard_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier5/ResourceSink_CircuitBoard.ResourceSink_CircuitBoard_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Circuit Board`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/CircuitBoard/UI/IconDesc_CircuitBoard_64.IconDesc_CircuitBoard_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Electronics.SC_RSS_Electronics_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CircuitBoard/Desc_CircuitBoard.Desc_CircuitBoard_C"',Amount=200))`,
			},
		},
	}

	CoffeeCup = FGSchematic{
		ClassName:               "ResourceSink_CoffeeCup_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_CoffeeCup.ResourceSink_CoffeeCup_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            `Worried about your superfluous body needs interrupting your work? We're proud to provide our pioneers with the means to stay efficient! Get your Coffee Cup! (Coffee not included).`,
		MDisplayName:            `FICSIT™ Coffee Cup`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Equipment/Cup/UI/IconDesc_CoffeeCup_256.IconDesc_CoffeeCup_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Cup/BP_EquipmentDescriptorCup.BP_EquipmentDescriptorCup_C"',Amount=1))`,
			},
		},
	}

	ColoredAmmoPack = FGSchematic{
		ClassName:               "ResourceSink_ColoredAmmoPack_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_ColoredAmmoPack.ResourceSink_ColoredAmmoPack_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Color Ammo Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/FlowerPetals_RS/Research_FlowerPetals_3.Research_FlowerPetals_3_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/ColorCartridge/UI/IconDesc_ColorCartridge_256.IconDesc_ColorCartridge_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ColorCartridge/Desc_ColorCartridge.Desc_ColorCartridge_C"',Amount=50))`,
			},
		},
	}

	Computer = FGSchematic{
		ClassName:               "ResourceSink_Computer_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier6/ResourceSink_Computer.ResourceSink_Computer_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=6))`,
		MDescription:            ``,
		MDisplayName:            `Computer`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-2.Schematic_5-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Computer/UI/IconDesc_Computer_256.IconDesc_Computer_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Communications.SC_RSS_Communications_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Computer/Desc_Computer.Desc_Computer_C"',Amount=50))`,
			},
		},
	}

	Concrete = FGSchematic{
		ClassName:               "ResourceSink_Concrete_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier1/ResourceSink_Concrete.ResourceSink_Concrete_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Concrete`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Cement/UI/IconDesc_Concrete_256.IconDesc_Concrete_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Minerals.SC_RSS_Minerals_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=100))`,
			},
		},
	}

	ConveryWallsMetal = FGSchematic{
		ClassName:               "ResourceSink_ConveryWalls_Metal_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveryWalls_Metal.ResourceSink_ConveryWalls_Metal_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=7))`,
		MDescription:            ``,
		MDisplayName:            `Metal Conveyor Walls`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           1.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveryWalls_Normal.ResourceSink_ConveryWalls_Normal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveyorWallMount.ResourceSink_ConveyorWallMount_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Wall/UI/Wall_Conveyor_x3_Grey_256.Wall_Conveyor_x3_Grey_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Conveyor_8x4_01_Steel.Recipe_Wall_Conveyor_8x4_01_Steel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Conveyor_8x4_02_Steel.Recipe_Wall_Conveyor_8x4_02_Steel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Conveyor_8x4_03_Steel.Recipe_Wall_Conveyor_8x4_03_Steel_C"')`,
			},
		},
	}

	ConveryWallsNormal = FGSchematic{
		ClassName:               "ResourceSink_ConveryWalls_Normal_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveryWalls_Normal.ResourceSink_ConveryWalls_Normal_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Conveyor Walls`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveryWalls_Metal.ResourceSink_ConveryWalls_Metal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveyorWallMount.ResourceSink_ConveyorWallMount_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Wall/UI/Wall_Conveyor_x3_Orange_256.Wall_Conveyor_x3_Orange_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Conveyor_8x4_01.Recipe_Wall_Conveyor_8x4_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Conveyor_8x4_02.Recipe_Wall_Conveyor_8x4_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Conveyor_8x4_03.Recipe_Wall_Conveyor_8x4_03_C"')`,
			},
		},
	}

	ConveyorWallMount = FGSchematic{
		ClassName:               "ResourceSink_ConveyorWallMount_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveyorWallMount.ResourceSink_ConveyorWallMount_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Conveyor Wall Mount`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveryWalls_Normal.ResourceSink_ConveryWalls_Normal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveryWalls_Metal.ResourceSink_ConveryWalls_Metal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_HyperTubeWallAttachements.ResourceSink_HyperTubeWallAttachements_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_PipelineWallAttachments.ResourceSink_PipelineWallAttachments_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPoles.ResourceSink_WallPowerPoles_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/ConveyorPoleWall/UI/ConveyorPole_Wall_256.ConveyorPole_Wall_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Buildable/Factory/ConveyorPoleWall/UI/ConveyorPole_Wall_256.ConveyorPole_Wall_256`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorPoleWall.Recipe_ConveyorPoleWall_C"')`,
			},
		},
	}

	CopperSheet = FGSchematic{
		ClassName:               "ResourceSink_CopperSheet_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier2/ResourceSink_CopperSheet.ResourceSink_CopperSheet_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Copper Sheet`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_2-1.Schematic_2-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/CopperSheet/UI/IconDesc_CopperSheet_256.IconDesc_CopperSheet_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `Texture2D /Game/FactoryGame/Resource/Parts/CopperSheet/UI/IconDesc_CopperSheet_64.IconDesc_CopperSheet_64`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CopperSheet/Desc_CopperSheet.Desc_CopperSheet_C"',Amount=100))`,
			},
		},
	}

	CrystalOscillator = FGSchematic{
		ClassName:               "ResourceSink_CrystalOscillator_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier5/ResourceSink_CrystalOscillator.ResourceSink_CrystalOscillator_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=4))`,
		MDescription:            ``,
		MDisplayName:            `Crystal Oscillator`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_2.Research_Quartz_2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/UI/IconDesc_CrystalOscillator_256.IconDesc_CrystalOscillator_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Communications.SC_RSS_Communications_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/Desc_CrystalOscillator.Desc_CrystalOscillator_C"',Amount=100))`,
			},
		},
	}

	CurvedFoundationPack = FGSchematic{
		ClassName:               "ResourceSink_CurvedFoundationPack_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CurvedFoundationPack.ResourceSink_CurvedFoundationPack_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=8))`,
		MDescription:            ``,
		MDisplayName:            `Quarter Pipes Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           4.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FactoryCart.ResourceSink_FactoryCart_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoundationExpansionPack.ResourceSink_FoundationExpansionPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedRampPack.ResourceSink_InvertedRampPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DiagonalRamps.ResourceSink_DiagonalRamps_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedCornerRamps.ResourceSink_InvertedCornerRamps_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=512.000000,Y=512.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Ramp/UI/QuarterPipe_01_512.QuarterPipe_01_512"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_QuarterPipe.Recipe_QuarterPipe_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_QuarterPipe_02.Recipe_QuarterPipe_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_QuarterPipeCorner_01.Recipe_QuarterPipeCorner_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_QuarterPipeCorner_02.Recipe_QuarterPipeCorner_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_QuarterPipeCorner_03.Recipe_QuarterPipeCorner_03_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_QuarterPipeCorner_04.Recipe_QuarterPipeCorner_04_C"')`,
			},
		},
	}

	CyberWagon = FGSchematic{
		ClassName:               "ResourceSink_CyberWagon_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CyberWagon.ResourceSink_CyberWagon_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=20))`,
		MDescription:            ``,
		MDisplayName:            `Cyber Wagon`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           1.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FactoryCart.ResourceSink_FactoryCart_C"')`,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CyberWagon_Unlock.ResourceSink_CyberWagon_Unlock_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Vehicle/Cyberwagon/UI/Cyberwagon_256.Cyberwagon_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Vehicle/Recipe_CyberWagon.Recipe_CyberWagon_C"')`,
			},
		},
	}

	CyberWagonUnlock = FGSchematic{
		ClassName:               "ResourceSink_CyberWagon_Unlock_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CyberWagon_Unlock.ResourceSink_CyberWagon_Unlock_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Cyber Wagon available in the AWESOME Shop`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Vehicle/Cyberwagon/UI/Cyberwagon_256.Cyberwagon_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(None)`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Custom`,
		MUnlocks:                nil,
	}

	DiagonalRamps = FGSchematic{
		ClassName:               "ResourceSink_DiagonalRamps_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DiagonalRamps.ResourceSink_DiagonalRamps_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=6))`,
		MDescription:            ``,
		MDisplayName:            `Corner Ramp Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           2.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedCornerRamps.ResourceSink_InvertedCornerRamps_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoundationExpansionPack.ResourceSink_FoundationExpansionPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedRampPack.ResourceSink_InvertedRampPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Ladders.ResourceSink_Ladders_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Stairs.ResourceSink_Stairs_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=512.000000,Y=512.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Ramp/UI/TXUI_Foundation_Corner_8x2_512.TXUI_Foundation_Corner_8x2_512"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_Diagonal_8x1_02.Recipe_Ramp_Diagonal_8x1_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_Diagonal_8x2_02.Recipe_Ramp_Diagonal_8x2_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_Diagonal_8x4_02.Recipe_Ramp_Diagonal_8x4_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_Diagonal_8x1_01.Recipe_Ramp_Diagonal_8x1_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_Diagonal_8x2_01.Recipe_Ramp_Diagonal_8x2_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_Diagonal_8x4_01.Recipe_Ramp_Diagonal_8x4_01_C"')`,
			},
		},
	}

	DoorWallsMetal = FGSchematic{
		ClassName:               "ResourceSink_DoorWalls_Metal_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DoorWalls_Metal.ResourceSink_DoorWalls_Metal_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=4))`,
		MDescription:            `Space dilation is still in the R&D phase at FICSIT research facilities. It's not yet recommended to try walking through walls. Purchase a Door today for safe factory traversal!`,
		MDisplayName:            `Metal Door Walls`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           3.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DoorWalls_Normal.ResourceSink_DoorWalls_Normal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WindowedWalls.ResourceSink_WindowedWalls_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Wall/UI/Wall_Door_Center_Grey_256.Wall_Door_Center_Grey_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Door_8x4_01_Steel.Recipe_Wall_Door_8x4_01_Steel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Door_8x4_02_Steel.Recipe_Wall_Door_8x4_02_Steel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Door_8x4_03_Steel.Recipe_Wall_Door_8x4_03_Steel_C"')`,
			},
		},
	}

	DoorWallsNormal = FGSchematic{
		ClassName:               "ResourceSink_DoorWalls_Normal_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DoorWalls_Normal.ResourceSink_DoorWalls_Normal_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            `Space dilation is still in the R&D phase at FICSIT research facilities. It's not yet recommended to try walking through walls. Purchase a Door today for safe factory traversal!`,
		MDisplayName:            `Door Walls`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           2.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DoorWalls_Metal.ResourceSink_DoorWalls_Metal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WindowedWalls.ResourceSink_WindowedWalls_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Wall/UI/Wall_Door_Center_Orange_256.Wall_Door_Center_Orange_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Door_8x4_01.Recipe_Wall_Door_8x4_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Door_8x4_02.Recipe_Wall_Door_8x4_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Door_8x4_03.Recipe_Wall_Door_8x4_03_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Gate_8x4_01.Recipe_Wall_Gate_8x4_01_C"')`,
			},
		},
	}

	EmptyCanister = FGSchematic{
		ClassName:               "ResourceSink_EmptyCanister_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier5/ResourceSink_EmptyCanister.ResourceSink_EmptyCanister_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Empty Canister`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/FluidCanister/UI/IconDesc_EmptyCannister_256.IconDesc_EmptyCannister_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_OilProducts.SC_RSS_OilProducts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/FluidCanister/Desc_FluidCanister.Desc_FluidCanister_C"',Amount=100))`,
			},
		},
	}

	EncasedIndustrialBeam = FGSchematic{
		ClassName:               "ResourceSink_EncasedIndustrialBeam_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier4/ResourceSink_EncasedIndustrialBeam.ResourceSink_EncasedIndustrialBeam_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Encased Industrial Beam`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_4-1.Schematic_4-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/SteelPlateReinforced/UI/IconDesc_EncasedSteelBeam_256.IconDesc_EncasedSteelBeam_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `Texture2D /Game/FactoryGame/Resource/Parts/SteelPlateReinforced/UI/IconDesc_EncasedSteelBeam_64.IconDesc_EncasedSteelBeam_64`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlateReinforced/Desc_SteelPlateReinforced.Desc_SteelPlateReinforced_C"',Amount=100))`,
			},
		},
	}

	Explosives = FGSchematic{
		ClassName:               "ResourceSink_Explosives_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_Explosives.ResourceSink_Explosives_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Explosives`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_3_2_1.Research_Sulfur_3_2_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/NobeliskExplosive/UI/IconDesc_Explosive_256.IconDesc_Explosive_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/NobeliskExplosive/Desc_NobeliskExplosive.Desc_NobeliskExplosive_C"',Amount=10))`,
			},
		},
	}

	Fabric = FGSchematic{
		ClassName:               "ResourceSink_Fabric_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier2/ResourceSink_Fabric.ResourceSink_Fabric_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Fabric`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Mycelia_RS/Research_Mycelia_2.Research_Mycelia_2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Fabric_256.IconDesc_Fabric_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Biomass.SC_RSS_Biomass_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_Fabric.Desc_Fabric_C"',Amount=100))`,
			},
		},
	}

	FactoryCart = FGSchematic{
		ClassName:               "ResourceSink_FactoryCart_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FactoryCart.ResourceSink_FactoryCart_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `FICSIT Factory Cart™`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CyberWagon.ResourceSink_CyberWagon_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CurvedFoundationPack.ResourceSink_CurvedFoundationPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoundationExpansionPack.ResourceSink_FoundationExpansionPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DiagonalRamps.ResourceSink_DiagonalRamps_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Walkways.ResourceSink_Walkways_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Equipment/GolfCart/UI/IconDesc_GolfCart_256.IconDesc_GolfCart_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Vehicle/Recipe_FactoryCart.Recipe_FactoryCart_C"')`,
			},
		},
	}

	FactoryRailing = FGSchematic{
		ClassName:               "ResourceSink_FactoryRailing_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FactoryRailing.ResourceSink_FactoryRailing_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Factory Railing`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           3.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Walkways.ResourceSink_Walkways_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Ladders.ResourceSink_Ladders_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Fence/UI/Fences_256.Fences_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Fence/Recipe_Fence_01.Recipe_Fence_01_C"')`,
			},
		},
	}

	FlowerPetals1 = FGSchematic{
		ClassName:               "Research_FlowerPetals_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/FlowerPetals_RS/Research_FlowerPetals_1.Research_FlowerPetals_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_FlowerPetals.Desc_FlowerPetals_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Flower Petals`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GenericBiomass/UI/FlowerPetals_Final_64.FlowerPetals_Final_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	FlowerPetals2 = FGSchematic{
		ClassName:               "Research_FlowerPetals_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/FlowerPetals_RS/Research_FlowerPetals_2.Research_FlowerPetals_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronScrew/Desc_IronScrew.Desc_IronScrew_C"',Amount=250),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=250))`,
		MDescription:            ``,
		MDisplayName:            `Color Gun`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/ColorGun/UI/ColorGun_64.ColorGun_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_ColorGun.Recipe_ColorGun_C"')`,
			},
		},
	}

	FlowerPetals3 = FGSchematic{
		ClassName:               "Research_FlowerPetals_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/FlowerPetals_RS/Research_FlowerPetals_3.Research_FlowerPetals_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_FlowerPetals.Desc_FlowerPetals_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Color Cartridges`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/ColorCartridge/UI/IconDesc_ColorCartridge_64.IconDesc_ColorCartridge_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_ColorCartridge.Recipe_ColorCartridge_C"')`,
			},
		},
	}

	FoudationPillar = FGSchematic{
		ClassName:               "ResourceSink_FoudationPillar_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoudationPillar.ResourceSink_FoudationPillar_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Foundation Pillar`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           6.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Pillars/UI/Pillar_Bottom_256.Pillar_Bottom_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_PillarTop.Recipe_PillarTop_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_PillarMiddle.Recipe_PillarMiddle_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_PillarBase.Recipe_PillarBase_C"')`,
			},
		},
	}

	FoundationExpansionPack = FGSchematic{
		ClassName:               "ResourceSink_FoundationExpansionPack_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoundationExpansionPack.ResourceSink_FoundationExpansionPack_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            `Never again will you build a Ramp one step at a time. Take over the world efficiently, increase your productivity and enjoy building in diagonal with the new Double Ramps!`,
		MDisplayName:            `Double Ramp Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedRampPack.ResourceSink_InvertedRampPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DiagonalRamps.ResourceSink_DiagonalRamps_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedCornerRamps.ResourceSink_InvertedCornerRamps_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Ladders.ResourceSink_Ladders_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Stairs.ResourceSink_Stairs_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Ramp/UI/Ramp_8x8x4_256.Ramp_8x8x4_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampDouble_8x1.Recipe_RampDouble_8x1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampDouble.Recipe_RampDouble_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_8x8x8.Recipe_Ramp_8x8x8_C"')`,
			},
		},
	}

	FrameworkFoundations = FGSchematic{
		ClassName:               "ResourceSink_FrameworkFoundations_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            `Meet your decoration needs while keeping your productivity in check with this exclusive pack of foundation variants! You will feel even better working unlimited hours!`,
		MDisplayName:            `Foundations Expansion Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           5.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoudationPillar.ResourceSink_FoudationPillar_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Walkways.ResourceSink_Walkways_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FactoryRailing.ResourceSink_FactoryRailing_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Ladders.ResourceSink_Ladders_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Stairs.ResourceSink_Stairs_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Foundation/UI/FrameworkFoundation_256.FrameworkFoundation_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_Foundation_Frame_01.Recipe_Foundation_Frame_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_FoundationGlass_01.Recipe_FoundationGlass_01_C"')`,
			},
		},
	}

	GasFilters = FGSchematic{
		ClassName:               "ResourceSink_GasFilters_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_GasFilters.ResourceSink_GasFilters_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Gas Filters`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_6-4.Schematic_6-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Filter/UI/IconDesc_GasMaskFilter_256.IconDesc_GasMaskFilter_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Filter/Desc_Filter.Desc_Filter_C"',Amount=25))`,
			},
		},
	}

	GoldenCup = FGSchematic{
		ClassName:               "ResourceSink_GoldenCup_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_GoldenCup.ResourceSink_GoldenCup_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            `This limited edition cup is only awarded to those Pioneers who have completed the last available stage of Project Assembly in its 'early design stage'.`,
		MDisplayName:            `'Employee of the Planet' Cup`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:      `BP_GamePhaseReachedDependency_C`,
				MGamePhase: `EGP_FoodCourt`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Equipment/GoldenCup/UI/IconDesc_CupGold_256.IconDesc_CupGold_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/GoldenCup/BP_EquipmentDescriptorCupGold.BP_EquipmentDescriptorCupGold_C"',Amount=1))`,
			},
		},
	}

	HardDrive0 = FGSchematic{
		ClassName:               "Research_HardDrive_0_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/HardDrive_RS/Research_HardDrive_0.Research_HardDrive_0_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/CrashSites/Desc_HardDrive.Desc_HardDrive_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Hard Drive`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Environment/CrashSites/UI/HardDrive_64.HardDrive_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               0,
		MTimeToComplete:         600.000000,
		MType:                   `EST_HardDrive`,
		MUnlocks:                nil,
	}

	HealthPack = FGSchematic{
		ClassName:               "ResourceSink_HealthPack_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_HealthPack.ResourceSink_HealthPack_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Health Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/AlienOrganisms_RS/Research_AOrgans_2.Research_AOrgans_2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Nutrients_RS/Research_Nutrients_4.Research_Nutrients_4_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Mycelia_RS/Research_Mycelia_5.Research_Mycelia_5_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/Medkit/UI/Inhaler_256.Inhaler_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Medkit/Desc_Medkit.Desc_Medkit_C"',Amount=5))`,
			},
		},
	}

	HeatSink = FGSchematic{
		ClassName:               "ResourceSink_HeatSink_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier7/ResourceSink_HeatSink.ResourceSink_HeatSink_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Heat Sink`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-2.Schematic_7-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/HeatSink/UI/IconDesc_Heatsink_256.IconDesc_Heatsink_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_IndustrialParts.SC_RSS_IndustrialParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumPlateReinforced/Desc_AluminumPlateReinforced.Desc_AluminumPlateReinforced_C"',Amount=100))`,
			},
		},
	}

	HeavyModularFrame = FGSchematic{
		ClassName:               "ResourceSink_HeavyModularFrame_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier6/ResourceSink_HeavyModularFrame.ResourceSink_HeavyModularFrame_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=6))`,
		MDescription:            ``,
		MDisplayName:            `Heavy Modular Frame`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-2.Schematic_5-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/UI/IconDesc_ModularFrameHeavy_256.IconDesc_ModularFrameHeavy_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `Texture2D /Game/FactoryGame/Resource/Parts/ModularFrameHeavy/UI/IconDesc_ModularFrameHeavy_64.IconDesc_ModularFrameHeavy_64`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=50))`,
			},
		},
	}

	HighSpeedConnector = FGSchematic{
		ClassName:               "ResourceSink_HighSpeedConnector_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier5/ResourceSink_HighSpeedConnector.ResourceSink_HighSpeedConnector_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=4))`,
		MDescription:            ``,
		MDisplayName:            `High-Speed Connector`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_5.Research_Caterium_5_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/HighSpeedConnector/UI/IconDesc_HighSpeedConnector_256.IconDesc_HighSpeedConnector_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Electronics.SC_RSS_Electronics_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedConnector/Desc_HighSpeedConnector.Desc_HighSpeedConnector_C"',Amount=100))`,
			},
		},
	}

	HyperTubeWallAttachements = FGSchematic{
		ClassName:               "ResourceSink_HyperTubeWallAttachements_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_HyperTubeWallAttachements.ResourceSink_HyperTubeWallAttachements_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            `Never again will a wall stop you from swooping through your factory. With the new Next-Gen Hypertube Wall Hole included in this pack, speed has no limits!`,
		MDisplayName:            `Hyper Tube Wall Attachments`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           2.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_PipelineWallAttachments.ResourceSink_PipelineWallAttachments_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveyorWallMount.ResourceSink_ConveyorWallMount_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPoles.ResourceSink_WallPowerPoles_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/HyperTubeWallSupport/UI/IconDesc_HyperTube_WallHole_256.IconDesc_HyperTube_WallHole_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Buildable/Factory/HyperTubeWallSupport/UI/IconDesc_HyperTube_WallHole_256.IconDesc_HyperTube_WallHole_256`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_HyperTubeWallSupport.Recipe_HyperTubeWallSupport_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_HyperTubeWallHole.Recipe_HyperTubeWallHole_C"')`,
			},
		},
	}

	InvertedCornerRamps = FGSchematic{
		ClassName:               "ResourceSink_InvertedCornerRamps_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedCornerRamps.ResourceSink_InvertedCornerRamps_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Inverted Corner Ramp Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           3.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DiagonalRamps.ResourceSink_DiagonalRamps_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoundationExpansionPack.ResourceSink_FoundationExpansionPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedRampPack.ResourceSink_InvertedRampPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Ladders.ResourceSink_Ladders_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Stairs.ResourceSink_Stairs_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Ramp/UI/IconDesc_C_Inv_Ramp_8x4_01_256.IconDesc_C_Inv_Ramp_8x4_01_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampInverted_8x1_Corner_01.Recipe_RampInverted_8x1_Corner_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampInverted_8x2_Corner_01.Recipe_RampInverted_8x2_Corner_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampInverted_8x4_Corner_01.Recipe_RampInverted_8x4_Corner_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampInverted_8x1_Corner_02.Recipe_RampInverted_8x1_Corner_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampInverted_8x2_Corner_02.Recipe_RampInverted_8x2_Corner_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampInverted_8x4_Corner_02.Recipe_RampInverted_8x4_Corner_02_C"')`,
			},
		},
	}

	InvertedRampPack = FGSchematic{
		ClassName:               "ResourceSink_InvertedRampPack_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedRampPack.ResourceSink_InvertedRampPack_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Inverted Ramp Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           1.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoundationExpansionPack.ResourceSink_FoundationExpansionPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DiagonalRamps.ResourceSink_DiagonalRamps_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_InvertedCornerRamps.ResourceSink_InvertedCornerRamps_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Ladders.ResourceSink_Ladders_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Stairs.ResourceSink_Stairs_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Ramp/UI/Ramp_8x4_Inverted_256.Ramp_8x4_Inverted_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampInverted_8x1.Recipe_RampInverted_8x1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_RampInverted_8x2_01.Recipe_RampInverted_8x2_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_8x4_Inverted_01.Recipe_Ramp_8x4_Inverted_01_C"')`,
			},
		},
	}

	Ladders = FGSchematic{
		ClassName:               "ResourceSink_Ladders_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Ladders.ResourceSink_Ladders_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=4))`,
		MDescription:            ``,
		MDisplayName:            `Factory Ladder`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Stairs.ResourceSink_Stairs_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Walkways.ResourceSink_Walkways_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoundationExpansionPack.ResourceSink_FoundationExpansionPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Ladder/UI/IconDesc_Ladder_256.IconDesc_Ladder_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ladder/Recipe_Ladder.Recipe_Ladder_C"')`,
			},
		},
	}

	LightControlPanel = FGSchematic{
		ClassName:               "ResourceSink_LightControlPanel_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_LightControlPanel.ResourceSink_LightControlPanel_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Lights Control Panel`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StreetLight.ResourceSink_StreetLight_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CeilingLight.ResourceSink_CeilingLight_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_LightTower.ResourceSink_LightTower_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/LightsControlPanel/UI/IconDesc_LightsControlPanel_256.IconDesc_LightsControlPanel_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_LightsControlPanel.Recipe_LightsControlPanel_C"')`,
			},
		},
	}

	LightTower = FGSchematic{
		ClassName:               "ResourceSink_LightTower_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_LightTower.ResourceSink_LightTower_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=5))`,
		MDescription:            `Specifically designed to illuminate outdoor areas, the Flood Lights are the perfect choice for lighting up construction sites or outposts. Their height makes them tower over even the largest of buildings. And for additional convenience the Wall version can be placed on any wall, for indoor settings or if more precise light placement is required.`,
		MDisplayName:            `Flood Lights`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StreetLight.ResourceSink_StreetLight_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CeilingLight.ResourceSink_CeilingLight_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoundationExpansionPack.ResourceSink_FoundationExpansionPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DiagonalRamps.ResourceSink_DiagonalRamps_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/Floodlight/UI/IconDesc_FloodlightWall_256.IconDesc_FloodlightWall_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_FloodlightPole.Recipe_FloodlightPole_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_FloodlightWall.Recipe_FloodlightWall_C"')`,
			},
		},
	}

	ModularFrame = FGSchematic{
		ClassName:               "ResourceSink_ModularFrame_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier2/ResourceSink_ModularFrame.ResourceSink_ModularFrame_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=4))`,
		MDescription:            ``,
		MDisplayName:            `Modular Frame`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_2-1.Schematic_2-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/ModularFrame/UI/IconDesc_ModularFrame_256.IconDesc_ModularFrame_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `Texture2D /Game/FactoryGame/Resource/Parts/ModularFrame/UI/IconDesc_ModularFrame_64.IconDesc_ModularFrame_64`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrame/Desc_ModularFrame.Desc_ModularFrame_C"',Amount=50))`,
			},
		},
	}

	Motor = FGSchematic{
		ClassName:               "ResourceSink_Motor_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier5/ResourceSink_Motor.ResourceSink_Motor_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Motor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_4-1.Schematic_4-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Motor/UI/IconDesc_Engine_256.IconDesc_Engine_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_IndustrialParts.SC_RSS_IndustrialParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=50))`,
			},
		},
	}

	Mycelia1 = FGSchematic{
		ClassName:               "Research_Mycelia_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Mycelia_RS/Research_Mycelia_1.Research_Mycelia_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_Mycelia.Desc_Mycelia_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Mycelia`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Mycelia_64.IconDesc_Mycelia_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Biomass_Mycelia.Recipe_Biomass_Mycelia_C"')`,
			},
		},
	}

	Mycelia2 = FGSchematic{
		ClassName:               "Research_Mycelia_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Mycelia_RS/Research_Mycelia_2.Research_Mycelia_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_Mycelia.Desc_Mycelia_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_GenericBiomass.Desc_GenericBiomass_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Fabric`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GenericBiomass/UI/IconDesc_Fabric_64.IconDesc_Fabric_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Fabric.Recipe_Fabric_C"')`,
			},
		},
	}

	Mycelia3 = FGSchematic{
		ClassName:               "Research_Mycelia_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Mycelia_RS/Research_Mycelia_3.Research_Mycelia_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_Fabric.Desc_Fabric_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Parachute`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/Beacon/UI/Parachute_64.Parachute_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_Parachute.Recipe_Parachute_C"')`,
			},
		},
	}

	Mycelia4 = FGSchematic{
		ClassName:               "Research_Mycelia_4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Mycelia_RS/Research_Mycelia_4.Research_Mycelia_4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/DesertShroom/Desc_Shroom.Desc_Shroom_C"',Amount=1),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Berry/Desc_Berry.Desc_Berry_C"',Amount=2),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Nut/Desc_Nut.Desc_Nut_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Medical Properties`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/InGame/-Shared/Mam_Key_64.Mam_Key_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Mycelia5 = FGSchematic{
		ClassName:               "Research_Mycelia_5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Mycelia_RS/Research_Mycelia_5.Research_Mycelia_5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_Mycelia.Desc_Mycelia_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=25))`,
		MDescription:            ``,
		MDisplayName:            `Medicinal Inhaler`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/Medkit/UI/Inhaler_64.Inhaler_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_MedicinalInhaler.Recipe_MedicinalInhaler_C"')`,
			},
		},
	}

	Nutrients0 = FGSchematic{
		ClassName:               "Research_Nutrients_0_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Nutrients_RS/Research_Nutrients_0.Research_Nutrients_0_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Berry/Desc_Berry.Desc_Berry_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Paleberry`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Environment/Berry/UI/Berry_64.Berry_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Nutrients1 = FGSchematic{
		ClassName:               "Research_Nutrients_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Nutrients_RS/Research_Nutrients_1.Research_Nutrients_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Nut/Desc_Nut.Desc_Nut_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Beryl Nut`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Environment/Nut/UI/Nut_64_new.Nut_64_new"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Nutrients2 = FGSchematic{
		ClassName:               "Research_Nutrients_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Nutrients_RS/Research_Nutrients_2.Research_Nutrients_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/DesertShroom/Desc_Shroom.Desc_Shroom_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Bacon Agaric`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Environment/DesertShroom/UI/Mushroom_64.Mushroom_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Nutrients3 = FGSchematic{
		ClassName:               "Research_Nutrients_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Nutrients_RS/Research_Nutrients_3.Research_Nutrients_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Stator/Desc_Stator.Desc_Stator_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=500))`,
		MDescription:            ``,
		MDisplayName:            `Nutritional Mixture`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/InGame/-Shared/Mam_Key_64.Mam_Key_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Nutrients4 = FGSchematic{
		ClassName:               "Research_Nutrients_4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Nutrients_RS/Research_Nutrients_4.Research_Nutrients_4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/DesertShroom/Desc_Shroom.Desc_Shroom_C"',Amount=2),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Berry/Desc_Berry.Desc_Berry_C"',Amount=4),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Nut/Desc_Nut.Desc_Nut_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Nutritional Inhaler`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/Medkit/UI/Inhaler_64.Inhaler_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_NutritionalInhaler.Recipe_NutritionalInhaler_C"')`,
			},
		},
	}

	PackagedBiofuel = FGSchematic{
		ClassName:               "ResourceSink_PackagedBiofuel_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier6/ResourceSink_PackagedBiofuel.ResourceSink_PackagedBiofuel_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Packaged Liquid Biofuel`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-4.Schematic_5-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Turbofuel/UI/IconDesc_LiquidBiofuel_256.IconDesc_LiquidBiofuel_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Biomass.SC_RSS_Biomass_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/BioFuel/Desc_PackagedBiofuel.Desc_PackagedBiofuel_C"',Amount=100))`,
			},
		},
	}

	PackagedFuel = FGSchematic{
		ClassName:               "ResourceSink_PackagedFuel_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier6/ResourceSink_PackagedFuel.ResourceSink_PackagedFuel_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Packaged Fuel`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-4.Schematic_5-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Fuel/UI/IconDesc_Fuel_256.IconDesc_Fuel_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_OilProducts.SC_RSS_OilProducts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Fuel/Desc_Fuel.Desc_Fuel_C"',Amount=100))`,
			},
		},
	}

	Parachutes = FGSchematic{
		ClassName:               "ResourceSink_Parachutes_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_Parachutes.ResourceSink_Parachutes_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Parachutes`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Mycelia_RS/Research_Mycelia_3.Research_Mycelia_3_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/Beacon/UI/Parachute_256.Parachute_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Beacon/Desc_Parachute.Desc_Parachute_C"',Amount=10))`,
			},
		},
	}

	PetroleumCoke = FGSchematic{
		ClassName:               "ResourceSink_PetroleumCoke_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier5/ResourceSink_PetroleumCoke.ResourceSink_PetroleumCoke_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Petroleum Coke`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/PetroleumCoke/UI/IconDesc_PetroleumCoke_256.IconDesc_PetroleumCoke_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_OilProducts.SC_RSS_OilProducts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/PetroleumCoke/Desc_PetroleumCoke.Desc_PetroleumCoke_C"',Amount=200))`,
			},
		},
	}

	PipelineWallAttachments = FGSchematic{
		ClassName:               "ResourceSink_PipelineWallAttachments_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_PipelineWallAttachments.ResourceSink_PipelineWallAttachments_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Pipeline Wall Attachments`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           1.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_HyperTubeWallAttachements.ResourceSink_HyperTubeWallAttachements_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveyorWallMount.ResourceSink_ConveyorWallMount_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPoles.ResourceSink_WallPowerPoles_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/PipelineSupportWallHole/UI/PipeSupportWallHole_256.PipeSupportWallHole_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Buildable/Factory/PipelineSupportWallHole/UI/PipeSupportWallHole_256.PipeSupportWallHole_256`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipeSupportWall.Recipe_PipeSupportWall_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipeSupportWallHole.Recipe_PipeSupportWallHole_C"')`,
			},
		},
	}

	Plastic = FGSchematic{
		ClassName:               "ResourceSink_Plastic_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier5/ResourceSink_Plastic.ResourceSink_Plastic_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Plastic`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Plastic/UI/IconDesc_Plastic_256.IconDesc_Plastic_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_OilProducts.SC_RSS_OilProducts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Plastic/Desc_Plastic.Desc_Plastic_C"',Amount=100))`,
			},
		},
	}

	Plate = FGSchematic{
		ClassName:               "ResourceSink_Plate_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier1/ResourceSink_Plate.ResourceSink_Plate_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Iron Plate`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/IronPlate/UI/IconDesc_IronPlates_256.IconDesc_IronPlates_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Resource/Parts/IronPlate/UI/IconDesc_IronPlates_64.IconDesc_IronPlates_64`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=100))`,
			},
		},
	}

	PolymerResin = FGSchematic{
		ClassName:               "ResourceSink_PolymerResin_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier6/ResourceSink_PolymerResin.ResourceSink_PolymerResin_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Polymer Resin`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/PolymerResin/UI/IconDesc_PolymerResin_256.IconDesc_PolymerResin_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_OilProducts.SC_RSS_OilProducts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/PolymerResin/Desc_PolymerResin.Desc_PolymerResin_C"',Amount=200))`,
			},
		},
	}

	PowerSlugs1 = FGSchematic{
		ClassName:               "Research_PowerSlugs_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/PowerSlugs_RS/Research_PowerSlugs_1.Research_PowerSlugs_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Crystal/Desc_Crystal.Desc_Crystal_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Green Power Slugs`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugGreen_64.PowerSlugGreen_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_PowerCrystalShard_1.Recipe_PowerCrystalShard_1_C"')`,
			},
		},
	}

	PowerSlugs2 = FGSchematic{
		ClassName:               "Research_PowerSlugs_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/PowerSlugs_RS/Research_PowerSlugs_2.Research_PowerSlugs_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Overclock Production`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/Assets/Shared/Overclock_Icon.Overclock_Icon"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class: `BP_UnlockBuildOverclock_C`,
			},
		},
	}

	PowerSlugs3 = FGSchematic{
		ClassName:               "Research_PowerSlugs_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/PowerSlugs_RS/Research_PowerSlugs_3.Research_PowerSlugs_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Slug Scanning`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugGreen_64.PowerSlugGreen_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         120.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	PowerSlugs4 = FGSchematic{
		ClassName:               "Research_PowerSlugs_4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/PowerSlugs_RS/Research_PowerSlugs_4.Research_PowerSlugs_4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Crystal/Desc_Crystal_mk2.Desc_Crystal_mk2_C"',Amount=1),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Yellow Power Shards`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugYellow_64.PowerSlugYellow_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_PowerCrystalShard_2.Recipe_PowerCrystalShard_2_C"')`,
			},
		},
	}

	PowerSlugs5 = FGSchematic{
		ClassName:               "Research_PowerSlugs_5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/PowerSlugs_RS/Research_PowerSlugs_5.Research_PowerSlugs_5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Environment/Crystal/Desc_Crystal_mk3.Desc_Crystal_mk3_C"',Amount=1),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrame/Desc_ModularFrame.Desc_ModularFrame_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Purple Power Shards`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Environment/Crystal/UI/PowerSlugPurple_64.PowerSlugPurple_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_PowerCrystalShard_3.Recipe_PowerCrystalShard_3_C"')`,
			},
		},
	}

	Quartz0 = FGSchematic{
		ClassName:               "Research_Quartz_0_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_0.Research_Quartz_0_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Quartz`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/QuartzCrystal/UI/IconDesc_QuartzCrystal_64.IconDesc_QuartzCrystal_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C"')`,
			},
		},
	}

	Quartz11 = FGSchematic{
		ClassName:               "Research_Quartz_1_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_1.Research_Quartz_1_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C"',Amount=20))`,
		MDescription:            ``,
		MDisplayName:            `Quartz Crystals`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/QuartzCrystal/UI/IconDesc_QuartzResource_64.IconDesc_QuartzResource_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_QuartzCrystal.Recipe_QuartzCrystal_C"')`,
			},
		},
	}

	Quartz12 = FGSchematic{
		ClassName:               "Research_Quartz_1_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_2.Research_Quartz_1_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C"',Amount=20))`,
		MDescription:            ``,
		MDisplayName:            `Silica`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Silica/UI/IconDesc_Silica_64.IconDesc_Silica_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Silica.Recipe_Silica_C"')`,
			},
		},
	}

	Quartz2 = FGSchematic{
		ClassName:               "Research_Quartz_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_2.Research_Quartz_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/QuartzCrystal/Desc_QuartzCrystal.Desc_QuartzCrystal_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Crystal Oscillator`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/UI/IconDesc_CrystalOscillator_64.IconDesc_CrystalOscillator_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_CrystalOscillator.Recipe_CrystalOscillator_C"')`,
			},
		},
	}

	Quartz3 = FGSchematic{
		ClassName:               "Research_Quartz_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_3.Research_Quartz_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/Desc_CrystalOscillator.Desc_CrystalOscillator_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Signal Technologies`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/InGame/-Shared/Mam_Key_64.Mam_Key_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Quartz31 = FGSchematic{
		ClassName:               "Research_Quartz_3_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_3_1.Research_Quartz_3_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/Desc_CrystalOscillator.Desc_CrystalOscillator_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrame/Desc_ModularFrame.Desc_ModularFrame_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `The Explorer`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Vehicle/Explorer/UI/Explorer_256.Explorer_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Vehicle/Recipe_Explorer.Recipe_Explorer_C"')`,
			},
		},
	}

	Quartz32 = FGSchematic{
		ClassName:               "Research_Quartz_3_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_3_2.Research_Quartz_3_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/Desc_CrystalOscillator.Desc_CrystalOscillator_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Beacon/BP_EquipmentDescriptorBeacon.BP_EquipmentDescriptorBeacon_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Frequency Mapping`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/Assets/Shared/Map_Icon.Map_Icon"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class: `BP_UnlockMap_C`,
			},
		},
	}

	Quartz33 = FGSchematic{
		ClassName:               "Research_Quartz_3_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_3_3.Research_Quartz_3_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/Desc_CrystalOscillator.Desc_CrystalOscillator_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumPlate/Desc_AluminumPlate.Desc_AluminumPlate_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Radio Control Unit`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/RadioControlUnit/UI/IconDesc_RadioControlUnit_64.IconDesc_RadioControlUnit_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_RadioControlUnit.Recipe_RadioControlUnit_C"')`,
			},
		},
	}

	Quartz4 = FGSchematic{
		ClassName:               "Research_Quartz_4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_4.Research_Quartz_4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/Desc_CrystalOscillator.Desc_CrystalOscillator_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Beacon/BP_EquipmentDescriptorBeacon.BP_EquipmentDescriptorBeacon_C"',Amount=15))`,
		MDescription:            ``,
		MDisplayName:            `Radar Technology`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/RadarTower/UI/RadarTower_256.RadarTower_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_RadarTower.Recipe_RadarTower_C"')`,
			},
		},
	}

	Quartz41 = FGSchematic{
		ClassName:               "Research_Quartz_4_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_4_1.Research_Quartz_4_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CrystalOscillator/Desc_CrystalOscillator.Desc_CrystalOscillator_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Beacon/BP_EquipmentDescriptorBeacon.BP_EquipmentDescriptorBeacon_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Radio Signal Scanning`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Equipment/ObjectScanner/Icons/CrashSite_64.CrashSite_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         300.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Quickwire = FGSchematic{
		ClassName:               "ResourceSink_Quickwire_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier2/ResourceSink_Quickwire.ResourceSink_Quickwire_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Quickwire`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_2.Research_Caterium_2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/UI/IconDesc_Quickwire_256.IconDesc_Quickwire_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Electronics.SC_RSS_Electronics_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/Desc_HighSpeedWire.Desc_HighSpeedWire_C"',Amount=500))`,
			},
		},
	}

	RadiationFilters = FGSchematic{
		ClassName:               "ResourceSink_RadiationFilters_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Equipment/ResourceSink_RadiationFilters.ResourceSink_RadiationFilters_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Radiation Filters`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-3.Schematic_7-3_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/IodineInfusedFilter/UI/IconDesc_HazmatFilter_256.IconDesc_HazmatFilter_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/SchematicCategories/SC_Walls.SC_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IodineInfusedFilter/Desc_HazmatFilter.Desc_HazmatFilter_C"',Amount=10))`,
			},
		},
	}

	RadioControlUnit = FGSchematic{
		ClassName:               "ResourceSink_RadioControlUnit_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier7/ResourceSink_RadioControlUnit.ResourceSink_RadioControlUnit_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=7))`,
		MDescription:            ``,
		MDisplayName:            `Radio Control Unit`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_3_3.Research_Quartz_3_3_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/RadioControlUnit/UI/IconDesc_RadioControlUnit_256.IconDesc_RadioControlUnit_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Communications.SC_RSS_Communications_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameLightweight/Desc_ModularFrameLightweight.Desc_ModularFrameLightweight_C"',Amount=50))`,
			},
		},
	}

	ReinforcedIronPlate = FGSchematic{
		ClassName:               "ResourceSink_ReinforcedIronPlate_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier2/ResourceSink_ReinforcedIronPlate.ResourceSink_ReinforcedIronPlate_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Reinforced Iron Plate`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/UI/IconDesc_ReinforcedIronPlates_256.IconDesc_ReinforcedIronPlates_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Resource/Parts/IronPlateReinforced/UI/IconDesc_ReinforcedIronPlates_64.IconDesc_ReinforcedIronPlates_64`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=100))`,
			},
		},
	}

	Rod = FGSchematic{
		ClassName:               "ResourceSink_Rod_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier1/ResourceSink_Rod.ResourceSink_Rod_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Iron Rod`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/IronRod/UI/IconDesc_IronRods_256.IconDesc_IronRods_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Resource/Parts/IronRod/UI/IconDesc_IronRods_64.IconDesc_IronRods_64`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=100))`,
			},
		},
	}

	Rotor = FGSchematic{
		ClassName:               "ResourceSink_Rotor_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier2/ResourceSink_Rotor.ResourceSink_Rotor_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Rotor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_2-1.Schematic_2-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Rotor/UI/IconDesc_Rotor_256.IconDesc_Rotor_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_IndustrialParts.SC_RSS_IndustrialParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=100))`,
			},
		},
	}

	Rubber = FGSchematic{
		ClassName:               "ResourceSink_Rubber_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier5/ResourceSink_Rubber.ResourceSink_Rubber_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Rubber`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Rubber/UI/IconDesc_Rubber_256.IconDesc_Rubber_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_OilProducts.SC_RSS_OilProducts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rubber/Desc_Rubber.Desc_Rubber_C"',Amount=100))`,
			},
		},
	}

	SaveCompatibility = FGSchematic{
		ClassName:               "Schematic_SaveCompatibility_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Schematic_SaveCompatibility.Schematic_SaveCompatibility_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Save Compatibility`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=32.000000,Y=32.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               0,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Custom`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Fence/Recipe_Fence_01.Recipe_Fence_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Stairs/Recipe_Stair_1b.Recipe_Stair_1b_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Window_8x4_03.Recipe_Wall_Window_8x4_03_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Window_8x4_03_Steel.Recipe_Wall_Window_8x4_03_Steel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Conveyor_8x4_04.Recipe_Wall_Conveyor_8x4_04_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Conveyor_8x4_04_Steel.Recipe_Wall_Conveyor_8x4_04_Steel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_JumpPad.Recipe_JumpPad_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_JumpPadTilted.Recipe_JumpPadTilted_C"')`,
			},
		},
	}

	Schematic1_1 = FGSchematic{
		ClassName:               "Schematic_1-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_1-1.Schematic_1-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Base Building`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Structure.SchematicIcon_Structure"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               1,
		MTimeToComplete:         120.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_LookoutTower.Recipe_LookoutTower_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_Foundation_8x1_01.Recipe_Foundation_8x1_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_Foundation_8x2_01.Recipe_Foundation_8x2_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Foundations/Recipe_Foundation_8x4_01.Recipe_Foundation_8x4_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_8x1_01.Recipe_Ramp_8x1_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_8x2_01.Recipe_Ramp_8x2_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Ramps/Recipe_Ramp_8x4_01.Recipe_Ramp_8x4_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_8x4_01.Recipe_Wall_8x4_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_8x4_02.Recipe_Wall_8x4_02_C"')`,
			},
		},
	}

	Schematic1_2 = FGSchematic{
		ClassName:               "Schematic_1-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_1-2.Schematic_1-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=150),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=150),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=300))`,
		MDescription:            ``,
		MDisplayName:            `Logistics`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Logistics.SchematicIcon_Logistics"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               1,
		MTimeToComplete:         240.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorAttachmentSplitter.Recipe_ConveyorAttachmentSplitter_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorAttachmentMerger.Recipe_ConveyorAttachmentMerger_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorLiftMk1.Recipe_ConveyorLiftMk1_C"')`,
			},
			{
				Class: `BP_UnlockBuildEfficiency_C`,
			},
		},
	}

	Schematic1_3 = FGSchematic{
		ClassName:               "Schematic_1-3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_1-3.Schematic_1-3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=300),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronScrew/Desc_IronScrew.Desc_IronScrew_C"',Amount=300),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Field Research`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Equipment.SchematicIcon_Equipment"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               1,
		MTimeToComplete:         180.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_Mam.Recipe_Mam_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_ObjectScanner.Recipe_ObjectScanner_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_Beacon.Recipe_Beacon_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_StoragePlayer.Recipe_StoragePlayer_C"')`,
			},
			{
				Class:                         `BP_UnlockArmEquipmentSlot_C`,
				MNumArmEquipmentSlotsToUnlock: 1,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Schematic2_1 = FGSchematic{
		ClassName:               "Schematic_2-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_2-1.Schematic_2-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronScrew/Desc_IronScrew.Desc_IronScrew_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=300))`,
		MDescription:            ``,
		MDisplayName:            `Part Assembly`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               2,
		MTimeToComplete:         360.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_AssemblerMk1.Recipe_AssemblerMk1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_CopperSheet.Recipe_CopperSheet_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_Rotor.Recipe_Rotor_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_ModularFrame.Recipe_ModularFrame_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/SpaceElevatorParts/Recipe_SpaceElevatorPart_1.Recipe_SpaceElevatorPart_1_C"')`,
			},
		},
	}

	Schematic2_2 = FGSchematic{
		ClassName:               "Schematic_2-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_2-2.Schematic_2-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronScrew/Desc_IronScrew.Desc_IronScrew_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Obstacle Clearing`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Equipment.SchematicIcon_Equipment"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               2,
		MTimeToComplete:         180.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_Chainsaw.Recipe_Chainsaw_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Biofuel.Recipe_Biofuel_C"')`,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Schematic2_3 = FGSchematic{
		ClassName:               "Schematic_2-3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_2-3.Schematic_2-3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=300),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=150))`,
		MDescription:            ``,
		MDisplayName:            `Jump Pads`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Structure.SchematicIcon_Structure"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               2,
		MTimeToComplete:         240.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_JumpPadAdjustable.Recipe_JumpPadAdjustable_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_UJellyLandingPad.Recipe_UJellyLandingPad_C"')`,
			},
		},
	}

	Schematic2_5 = FGSchematic{
		ClassName:               "Schematic_2-5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_2-5.Schematic_2-5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=400),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Resource Sink Bonus Program`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               2,
		MTimeToComplete:         300.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ResourceSink.Recipe_ResourceSink_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ResourceSinkShop.Recipe_ResourceSinkShop_C"')`,
			},
		},
	}

	Schematic3_1 = FGSchematic{
		ClassName:               "Schematic_3-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_3-1.Schematic_3-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=150),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=300))`,
		MDescription:            ``,
		MDisplayName:            `Coal Power`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         480.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_GeneratorCoal.Recipe_GeneratorCoal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_WaterPump.Recipe_WaterPump_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_Pipeline.Recipe_Pipeline_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipeSupport.Recipe_PipeSupport_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipelineJunction_Cross.Recipe_PipelineJunction_Cross_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipelinePump.Recipe_PipelinePump_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipeStorageTank.Recipe_PipeStorageTank_C"')`,
			},
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Coal/Desc_Coal.Desc_Coal_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Coal/Desc_Coal.Desc_Coal_C"')`,
			},
		},
	}

	Schematic3_2 = FGSchematic{
		ClassName:               "Schematic_3-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_3-2.Schematic_3-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=300),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=300))`,
		MDescription:            ``,
		MDisplayName:            `Logistics Mk.2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Logistics.SchematicIcon_Logistics"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               2,
		MTimeToComplete:         360.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorBeltMk2.Recipe_ConveyorBeltMk2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorPoleStackable.Recipe_ConveyorPoleStackable_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorLiftMk2.Recipe_ConveyorLiftMk2_C"')`,
			},
		},
	}

	Schematic3_3 = FGSchematic{
		ClassName:               "Schematic_3-3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_3-3.Schematic_3-3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrame/Desc_ModularFrame.Desc_ModularFrame_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=400))`,
		MDescription:            ``,
		MDisplayName:            `Vehicular Transport`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Vehicle.SchematicIcon_Vehicle"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         240.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Vehicle/Recipe_Tractor.Recipe_Tractor_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_TruckStation.Recipe_TruckStation_C"')`,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Schematic3_4 = FGSchematic{
		ClassName:               "Schematic_3-4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrame/Desc_ModularFrame.Desc_ModularFrame_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=150),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=300),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=1000))`,
		MDescription:            ``,
		MDisplayName:            `Basic Steel Production`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         480.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_SmelterMk1.Recipe_SmelterMk1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Smelter/Recipe_IngotSteel.Recipe_IngotSteel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_SteelBeam.Recipe_SteelBeam_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_SteelPipe.Recipe_SteelPipe_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/SpaceElevatorParts/Recipe_SpaceElevatorPart_2.Recipe_SpaceElevatorPart_2_C"')`,
			},
		},
	}

	Schematic4_1 = FGSchematic{
		ClassName:               "Schematic_4-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_4-1.Schematic_4-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=1500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=300))`,
		MDescription:            ``,
		MDisplayName:            `Advanced Steel Production`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               4,
		MTimeToComplete:         600.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_MinerMk2.Recipe_MinerMk2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_EncasedIndustrialBeam.Recipe_EncasedIndustrialBeam_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_Stator.Recipe_Stator_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_Motor.Recipe_Motor_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/SpaceElevatorParts/Recipe_SpaceElevatorPart_3.Recipe_SpaceElevatorPart_3_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_ModularFrameHeavy.Recipe_ModularFrameHeavy_C"')`,
			},
		},
	}

	Schematic4_2 = FGSchematic{
		ClassName:               "Schematic_4-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_4-2.Schematic_4-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rotor/Desc_Rotor.Desc_Rotor_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=1500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Improved Melee Combat`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Equipment.SchematicIcon_Equipment"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               4,
		MTimeToComplete:         180.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_XenoBasher.Recipe_XenoBasher_C"')`,
			},
			{
				Class:                         `BP_UnlockArmEquipmentSlot_C`,
				MNumArmEquipmentSlotsToUnlock: 1,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Schematic4_4 = FGSchematic{
		ClassName:               "Schematic_4-4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_4-4.Schematic_4-4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CopperSheet/Desc_CopperSheet.Desc_CopperSheet_C"',Amount=300),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=300),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlateReinforced/Desc_SteelPlateReinforced.Desc_SteelPlateReinforced_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Hyper Tubes`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Structure.SchematicIcon_Structure"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               4,
		MTimeToComplete:         600.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipe_PipeHyperStart.Recipe_PipeHyperStart_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipeHyper.Recipe_PipeHyper_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipeHyperSupport.Recipe_PipeHyperSupport_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_HyperPoleStackable.Recipe_HyperPoleStackable_C"')`,
			},
		},
	}

	Schematic5_1 = FGSchematic{
		ClassName:               "Schematic_5-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_5-1.Schematic_5-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlateReinforced/Desc_SteelPlateReinforced.Desc_SteelPlateReinforced_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CopperSheet/Desc_CopperSheet.Desc_CopperSheet_C"',Amount=500))`,
		MDescription:            ``,
		MDisplayName:            `Oil Processing`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               5,
		MTimeToComplete:         720.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_OilPump.Recipe_OilPump_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_OilRefinery.Recipe_OilRefinery_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_Valve.Recipe_Valve_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_Plastic.Recipe_Plastic_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_Rubber.Recipe_Rubber_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_LiquidFuel.Recipe_LiquidFuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_PetroleumCoke.Recipe_PetroleumCoke_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_CircuitBoard.Recipe_CircuitBoard_C"')`,
			},
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/CrudeOil/Desc_LiquidOil.Desc_LiquidOil_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/CrudeOil/Desc_LiquidOil.Desc_LiquidOil_C"')`,
			},
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-1-1.Schematic_5-1-1_C"')`,
			},
		},
	}

	Schematic5_1_1 = FGSchematic{
		ClassName:               "Schematic_5-1-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_5-1-1.Schematic_5-1-1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Oil Processing 2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               5,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Custom`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_ResidualPlastic.Recipe_ResidualPlastic_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_ResidualRubber.Recipe_ResidualRubber_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_ResidualFuel.Recipe_ResidualFuel_C"')`,
			},
		},
	}

	Schematic5_2 = FGSchematic{
		ClassName:               "Schematic_5-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_5-2.Schematic_5-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Plastic/Desc_Plastic.Desc_Plastic_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rubber/Desc_Rubber.Desc_Rubber_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=1000))`,
		MDescription:            ``,
		MDisplayName:            `Industrial Manufacturing`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               5,
		MTimeToComplete:         720.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ManufacturerMk1.Recipe_ManufacturerMk1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Vehicle/Recipe_Truck.Recipe_Truck_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_Computer.Recipe_Computer_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/SpaceElevatorParts/Recipe_SpaceElevatorPart_4.Recipe_SpaceElevatorPart_4_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/SpaceElevatorParts/Recipe_SpaceElevatorPart_5.Recipe_SpaceElevatorPart_5_C"')`,
			},
		},
	}

	Schematic5_3 = FGSchematic{
		ClassName:               "Schematic_5-3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_5-3.Schematic_5-3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlate/Desc_SteelPlate.Desc_SteelPlate_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=500))`,
		MDescription:            ``,
		MDisplayName:            `Logistics Mk.3`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Logistics.SchematicIcon_Logistics"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               4,
		MTimeToComplete:         300.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerStorageMk1.Recipe_PowerStorageMk1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_StorageContainerMk2.Recipe_StorageContainerMk2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorBeltMk3.Recipe_ConveyorBeltMk3_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorLiftMk3.Recipe_ConveyorLiftMk3_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipeSupportStackable.Recipe_PipeSupportStackable_C"')`,
			},
		},
	}

	Schematic5_4 = FGSchematic{
		ClassName:               "Schematic_5-4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_5-4.Schematic_5-4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=25),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Plastic/Desc_Plastic.Desc_Plastic_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=3000))`,
		MDescription:            ``,
		MDisplayName:            `Alternative Fluid Transport`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Logistics.SchematicIcon_Logistics"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               5,
		MTimeToComplete:         480.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_Packager.Recipe_Packager_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_FluidCanister.Recipe_FluidCanister_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_PackagedWater.Recipe_PackagedWater_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_PackagedCrudeOil.Recipe_PackagedCrudeOil_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_Fuel.Recipe_Fuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_PackagedOilResidue.Recipe_PackagedOilResidue_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_PackagedBiofuel.Recipe_PackagedBiofuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_LiquidBiofuel.Recipe_LiquidBiofuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_IndustrialTank.Recipe_IndustrialTank_C"')`,
			},
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_5-4-1.Schematic_5-4-1_C"')`,
			},
		},
	}

	Schematic5_4_1 = FGSchematic{
		ClassName:               "Schematic_5-4-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_5-4-1.Schematic_5-4-1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            ` Alternative Fluid Transport 2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               5,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Custom`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_UnpackageWater.Recipe_UnpackageWater_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_UnpackageOil.Recipe_UnpackageOil_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_UnpackageFuel.Recipe_UnpackageFuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_UnpackageOilResidue.Recipe_UnpackageOilResidue_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_UnpackageBioFuel.Recipe_UnpackageBioFuel_C"')`,
			},
		},
	}

	Schematic6_1 = FGSchematic{
		ClassName:               "Schematic_6-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_6-1.Schematic_6-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Computer/Desc_Computer.Desc_Computer_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlateReinforced/Desc_SteelPlateReinforced.Desc_SteelPlateReinforced_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rubber/Desc_Rubber.Desc_Rubber_C"',Amount=400))`,
		MDescription:            ``,
		MDisplayName:            `Expanded Power Infrastructure`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               6,
		MTimeToComplete:         900.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_GeneratorFuel.Recipe_GeneratorFuel_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorBeltMk4.Recipe_ConveyorBeltMk4_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorLiftMk4.Recipe_ConveyorLiftMk4_C"')`,
			},
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreGold/Desc_OreGold.Desc_OreGold_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreGold/Desc_OreGold.Desc_OreGold_C"')`,
			},
		},
	}

	Schematic6_2 = FGSchematic{
		ClassName:               "Schematic_6-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_6-2.Schematic_6-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Plastic/Desc_Plastic.Desc_Plastic_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rubber/Desc_Rubber.Desc_Rubber_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Fuel/Desc_Fuel.Desc_Fuel_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Jetpack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Equipment.SchematicIcon_Equipment"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               6,
		MTimeToComplete:         300.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_JetPack.Recipe_JetPack_C"')`,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Schematic6_3 = FGSchematic{
		ClassName:               "Schematic_6-3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_6-3.Schematic_6-3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Computer/Desc_Computer.Desc_Computer_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlate/Desc_SteelPlate.Desc_SteelPlate_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=600))`,
		MDescription:            ``,
		MDisplayName:            `Monorail Train Technology`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Vehicle.SchematicIcon_Vehicle"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               6,
		MTimeToComplete:         900.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Vehicle/Train/Recipe_Locomotive.Recipe_Locomotive_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Vehicle/Train/Recipe_FreightWagon.Recipe_FreightWagon_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_RailroadTrack.Recipe_RailroadTrack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Buildable/Factory/Train/Station/Recipe_TrainStation.Recipe_TrainStation_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Buildable/Factory/Train/Station/Recipe_TrainDockingStation.Recipe_TrainDockingStation_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Buildable/Factory/Train/Station/Recipe_TrainDockingStationLiquid.Recipe_TrainDockingStationLiquid_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Buildable/Factory/Train/Station/Recipe_TrainPlatformEmpty.Recipe_TrainPlatformEmpty_C"')`,
			},
		},
	}

	Schematic6_4 = FGSchematic{
		ClassName:               "Schematic_6-4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_6-4.Schematic_6-4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rubber/Desc_Rubber.Desc_Rubber_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Plastic/Desc_Plastic.Desc_Plastic_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GenericBiomass/Desc_Fabric.Desc_Fabric_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Gas Mask`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Equipment.SchematicIcon_Equipment"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               5,
		MTimeToComplete:         300.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_Gasmask.Recipe_Gasmask_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_FilterGasMask.Recipe_FilterGasMask_C"')`,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Schematic6_5 = FGSchematic{
		ClassName:               "Schematic_6-5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_6-5.Schematic_6-5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CopperSheet/Desc_CopperSheet.Desc_CopperSheet_C"',Amount=1000),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Plastic/Desc_Plastic.Desc_Plastic_C"',Amount=400),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rubber/Desc_Rubber.Desc_Rubber_C"',Amount=400),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Pipeline Engineering Mk.2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Logistics.SchematicIcon_Logistics"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               6,
		MTimeToComplete:         600.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipelineMK2.Recipe_PipelineMK2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PipelinePumpMK2.Recipe_PipelinePumpMK2_C"')`,
			},
		},
	}

	Schematic7_1 = FGSchematic{
		ClassName:               "Schematic_7-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_7-1.Schematic_7-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Computer/Desc_Computer.Desc_Computer_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rubber/Desc_Rubber.Desc_Rubber_C"',Amount=500))`,
		MDescription:            ``,
		MDisplayName:            `Bauxite Refinement`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               7,
		MTimeToComplete:         600.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_AluminaSolution.Recipe_AluminaSolution_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Packager/Recipe_PackagedAlumina.Recipe_PackagedAlumina_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_AluminumScrap.Recipe_AluminumScrap_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Smelter/Recipe_IngotAluminum.Recipe_IngotAluminum_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_AluminumSheet.Recipe_AluminumSheet_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_AluminumCasing.Recipe_AluminumCasing_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_RadioControlUnit.Recipe_RadioControlUnit_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_Blender.Recipe_Blender_C"')`,
			},
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreBauxite/Desc_OreBauxite.Desc_OreBauxite_C"'),(ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreBauxite/Desc_OreBauxite.Desc_OreBauxite_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/RawQuartz/Desc_RawQuartz.Desc_RawQuartz_C"')`,
			},
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-1-1.Schematic_7-1-1_C"')`,
			},
		},
	}

	Schematic7_1_1 = FGSchematic{
		ClassName:               "Schematic_7-1-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_7-1-1.Schematic_7-1-1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Bauxite Refinement 2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               7,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Custom`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Packager/Recipe_UnpackageAlumina.Recipe_UnpackageAlumina_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_QuartzCrystal.Recipe_QuartzCrystal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_CrystalOscillator.Recipe_CrystalOscillator_C"')`,
			},
		},
	}

	Schematic7_2 = FGSchematic{
		ClassName:               "Schematic_7-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_7-2.Schematic_7-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumPlate/Desc_AluminumPlate.Desc_AluminumPlate_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlateReinforced/Desc_SteelPlateReinforced.Desc_SteelPlateReinforced_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C"',Amount=300))`,
		MDescription:            ``,
		MDisplayName:            `Logistics Mk.5`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Logistics.SchematicIcon_Logistics"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               7,
		MTimeToComplete:         60.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorBeltMk5.Recipe_ConveyorBeltMk5_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorLiftMk5.Recipe_ConveyorLiftMk5_C"')`,
			},
		},
	}

	Schematic7_3 = FGSchematic{
		ClassName:               "Schematic_7-3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_7-3.Schematic_7-3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumCasing/Desc_AluminumCasing.Desc_AluminumCasing_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/HighSpeedWire/Desc_HighSpeedWire.Desc_HighSpeedWire_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Filter/Desc_Filter.Desc_Filter_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Hazmat Suit`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Equipment.SchematicIcon_Equipment"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               7,
		MTimeToComplete:         300.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_HazmatSuit.Recipe_HazmatSuit_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_FilterHazmat.Recipe_FilterHazmat_C"')`,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Schematic7_4 = FGSchematic{
		ClassName:               "Schematic_7-4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_7-4.Schematic_7-4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameLightweight/Desc_ModularFrameLightweight.Desc_ModularFrameLightweight_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumPlate/Desc_AluminumPlate.Desc_AluminumPlate_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumCasing/Desc_AluminumCasing.Desc_AluminumCasing_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=300))`,
		MDescription:            ``,
		MDisplayName:            `Aeronautical Engineering`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Vehicle.SchematicIcon_Vehicle"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               7,
		MTimeToComplete:         900.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Sulfur/Desc_Sulfur.Desc_Sulfur_C"'))`,
				MResourcesToAddToScanner:     ``,
			},
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/OilRefinery/Recipe_SulfuricAcid.Recipe_SulfuricAcid_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Packager/Recipe_PackagedSulfuricAcid.Recipe_PackagedSulfuricAcid_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_Battery.Recipe_Battery_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_ComputerSuper.Recipe_ComputerSuper_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_DroneTransport.Recipe_DroneTransport_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_DroneStation.Recipe_DroneStation_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/SpaceElevatorParts/Recipe_SpaceElevatorPart_7.Recipe_SpaceElevatorPart_7_C"')`,
			},
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-4-1.Schematic_7-4-1_C"')`,
			},
		},
	}

	Schematic7_4_1 = FGSchematic{
		ClassName:               "Schematic_7-4-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_7-4-1.Schematic_7-4-1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Aeronautical Engineering 2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               7,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Custom`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Packager/Recipe_UnpackageSulfuricAcid.Recipe_UnpackageSulfuricAcid_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_AILimiter.Recipe_AILimiter_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_HighSpeedConnector.Recipe_HighSpeedConnector_C"')`,
			},
		},
	}

	Schematic8_1 = FGSchematic{
		ClassName:               "Schematic_8-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_8-1.Schematic_8-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ComputerSuper/Desc_ComputerSuper.Desc_ComputerSuper_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=1000),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=2000))`,
		MDescription:            ``,
		MDisplayName:            `Nuclear Power`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               8,
		MTimeToComplete:         900.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreUranium/Desc_OreUranium.Desc_OreUranium_C"'))`,
				MResourcesToAddToScanner:     ``,
			},
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_UraniumCell.Recipe_UraniumCell_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_ElectromagneticControlRod.Recipe_ElectromagneticControlRod_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_NuclearFuelRod.Recipe_NuclearFuelRod_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_GeneratorNuclear.Recipe_GeneratorNuclear_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/SpaceElevatorParts/Recipe_SpaceElevatorPart_6.Recipe_SpaceElevatorPart_6_C"')`,
			},
		},
	}

	Schematic8_2 = FGSchematic{
		ClassName:               "Schematic_8-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_8-2.Schematic_8-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameLightweight/Desc_ModularFrameLightweight.Desc_ModularFrameLightweight_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumCasing/Desc_AluminumCasing.Desc_AluminumCasing_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumPlate/Desc_AluminumPlate.Desc_AluminumPlate_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=3000))`,
		MDescription:            ``,
		MDisplayName:            `Advanced Aluminum Production`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               8,
		MTimeToComplete:         900.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_FrackingSmasher.Recipe_FrackingSmasher_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_FrackingExtractor.Recipe_FrackingExtractor_C"')`,
			},
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Water/Desc_Water.Desc_Water_C"',ResourceNodeType=FrackingCore),(ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/CrudeOil/Desc_LiquidOil.Desc_LiquidOil_C"',ResourceNodeType=FrackingCore),(ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/NitrogenGas/Desc_NitrogenGas.Desc_NitrogenGas_C"',ResourceNodeType=FrackingCore))`,
				MResourcesToAddToScanner:     ``,
			},
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_GasTank.Recipe_GasTank_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Packager/Recipe_PackagedNitrogen.Recipe_PackagedNitrogen_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_HeatSink.Recipe_HeatSink_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Blender/Recipe_CoolingSystem.Recipe_CoolingSystem_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Blender/Recipe_FusedModularFrame.Recipe_FusedModularFrame_C"')`,
			},
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-2-1.Schematic_8-2-1_C"')`,
			},
		},
	}

	Schematic8_2_1 = FGSchematic{
		ClassName:               "Schematic_8-2-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_8-2-1.Schematic_8-2-1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Advanced Aluminum Production 2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               8,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Custom`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Packager/Recipe_UnpackageNitrogen.Recipe_UnpackageNitrogen_C"')`,
			},
		},
	}

	Schematic8_3 = FGSchematic{
		ClassName:               "Schematic_8-3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_8-3.Schematic_8-3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Motor/Desc_Motor.Desc_Motor_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Computer/Desc_Computer.Desc_Computer_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/AluminumPlate/Desc_AluminumPlate.Desc_AluminumPlate_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Hover Pack`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Equipment.SchematicIcon_Equipment"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               7,
		MTimeToComplete:         300.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_Hoverpack.Recipe_Hoverpack_C"')`,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Schematic8_4 = FGSchematic{
		ClassName:               "Schematic_8-4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_8-4.Schematic_8-4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameFused/Desc_ModularFrameFused.Desc_ModularFrameFused_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ComputerSuper/Desc_ComputerSuper.Desc_ComputerSuper_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=1000))`,
		MDescription:            ``,
		MDisplayName:            `Leading-edge Production`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               8,
		MTimeToComplete:         300.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_MotorTurbo.Recipe_MotorTurbo_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_MinerMk3.Recipe_MinerMk3_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/SpaceElevatorParts/Recipe_SpaceElevatorPart_8.Recipe_SpaceElevatorPart_8_C"')`,
			},
		},
	}

	Schematic8_5 = FGSchematic{
		ClassName:               "Schematic_8-5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_8-5.Schematic_8-5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ElectromagneticControlRod/Desc_ElectromagneticControlRod.Desc_ElectromagneticControlRod_C"',Amount=400),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CoolingSystem/Desc_CoolingSystem.Desc_CoolingSystem_C"',Amount=400),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameFused/Desc_ModularFrameFused.Desc_ModularFrameFused_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/MotorLightweight/Desc_MotorLightweight.Desc_MotorLightweight_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Particle Enrichment`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               8,
		MTimeToComplete:         1200.000000,
		MType:                   `EST_Milestone`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Blender/Recipe_NitricAcid.Recipe_NitricAcid_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Packager/Recipe_PackagedNitricAcid.Recipe_PackagedNitricAcid_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Blender/Recipe_NonFissileUranium.Recipe_NonFissileUranium_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/HadronCollider/Recipe_Plutonium.Recipe_Plutonium_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_PlutoniumCell.Recipe_PlutoniumCell_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Manufacturer/Recipe_PlutoniumFuelRod.Recipe_PlutoniumFuelRod_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_HadronCollider.Recipe_HadronCollider_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_CopperDust.Recipe_CopperDust_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_PressureConversionCube.Recipe_PressureConversionCube_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/SpaceElevatorParts/Recipe_SpaceElevatorPart_9.Recipe_SpaceElevatorPart_9_C"')`,
			},
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_8-5-1.Schematic_8-5-1_C"')`,
			},
		},
	}

	Schematic8_5_1 = FGSchematic{
		ClassName:               "Schematic_8-5-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Progression/Schematic_8-5-1.Schematic_8-5-1_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Particle Enrichment 2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Factory.SchematicIcon_Factory"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               8,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Custom`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Packager/Recipe_UnpackageNitricAcid.Recipe_UnpackageNitricAcid_C"')`,
			},
		},
	}

	Screw = FGSchematic{
		ClassName:               "ResourceSink_Screw_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier1/ResourceSink_Screw.ResourceSink_Screw_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Screw`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/IronScrew/UI/IconDesc_IronScrews_256.IconDesc_IronScrews_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Resource/Parts/IronScrew/UI/IconDesc_IronScrews_64.IconDesc_IronScrews_64`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronScrew/Desc_IronScrew.Desc_IronScrew_C"',Amount=500))`,
			},
		},
	}

	Silica = FGSchematic{
		ClassName:               "ResourceSink_Silica_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier7/ResourceSink_Silica.ResourceSink_Silica_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Silica`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Quartz_RS/Research_Quartz_1_2.Research_Quartz_1_2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Silica/UI/IconDesc_Silica_256.IconDesc_Silica_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Minerals.SC_RSS_Minerals_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Silica/Desc_Silica.Desc_Silica_C"',Amount=100))`,
			},
		},
	}

	Stairs = FGSchematic{
		ClassName:               "ResourceSink_Stairs_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Stairs.ResourceSink_Stairs_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=2))`,
		MDescription:            ``,
		MDisplayName:            `Stairs`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           1.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Ladders.ResourceSink_Ladders_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Walkways.ResourceSink_Walkways_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FoundationExpansionPack.ResourceSink_FoundationExpansionPack_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Stair/UI/StairLeft_256.StairLeft_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Stairs/Recipe_Stairs_Left_01.Recipe_Stairs_Left_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Stairs/Recipe_Stairs_Right_01.Recipe_Stairs_Right_01_C"')`,
			},
		},
	}

	StartingRecipes = FGSchematic{
		ClassName:               "Schematic_StartingRecipes_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Schematic_StartingRecipes.Schematic_StartingRecipes_C`,
		MCost:                   ``,
		MDescription:            ``,
		MDisplayName:            `Starting Blueprints`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=32.000000,Y=32.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               0,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Custom`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Smelter/Recipe_IngotIron.Recipe_IngotIron_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_IronPlate.Recipe_IronPlate_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_IronRod.Recipe_IronRod_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_TradingPost.Recipe_TradingPost_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_OreIron.Recipe_OreIron_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_OreCopper.Recipe_OreCopper_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_OreBauxite.Recipe_OreBauxite_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_OreCaterium.Recipe_OreCaterium_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_OreUranium.Recipe_OreUranium_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_CrudeOil.Recipe_CrudeOil_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_Sulfur.Recipe_Sulfur_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_Limestone.Recipe_Limestone_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_Coal.Recipe_Coal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/RawResources/Recipe_RawQuartz.Recipe_RawQuartz_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_XenoZapper.Recipe_XenoZapper_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_WorkBench.Recipe_WorkBench_C"')`,
			},
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreIron/Desc_OreIron.Desc_OreIron_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreIron/Desc_OreIron.Desc_OreIron_C"')`,
			},
		},
	}

	Stator = FGSchematic{
		ClassName:               "ResourceSink_Stator_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier4/ResourceSink_Stator.ResourceSink_Stator_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=3))`,
		MDescription:            ``,
		MDisplayName:            `Stator`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_4-1.Schematic_4-1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Stator/UI/IconDesc_Stator_256.IconDesc_Stator_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_IndustrialParts.SC_RSS_IndustrialParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Stator/Desc_Stator.Desc_Stator_C"',Amount=100))`,
			},
		},
	}

	StatueBronzePioneer = FGSchematic{
		ClassName:               "ResourceSink_StatueBronzePioneer_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueBronzePioneer.ResourceSink_StatueBronzePioneer_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=25))`,
		MDescription:            ``,
		MDisplayName:            `Adequate Pioneering`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueSilverPioneer.ResourceSink_StatueSilverPioneer_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueGoldPioneer.ResourceSink_StatueGoldPioneer_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Decor/Statues/UI/Award_Statue_Character_Bronze_256.Award_Statue_Character_Bronze_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Pioneering.SC_RSS_Pioneering_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Decoration/Desc_CharacterRunStatue.Desc_CharacterRunStatue_C"',Amount=1))`,
			},
		},
	}

	StatueGoldPioneer = FGSchematic{
		ClassName:               "ResourceSink_StatueGoldPioneer_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueGoldPioneer.ResourceSink_StatueGoldPioneer_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=150))`,
		MDescription:            ``,
		MDisplayName:            `Satisfactory Pioneering`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           2.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueBronzePioneer.ResourceSink_StatueBronzePioneer_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueSilverPioneer.ResourceSink_StatueSilverPioneer_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Decor/Statues/UI/Award_Statue_Character_Gold_256.Award_Statue_Character_Gold_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Pioneering.SC_RSS_Pioneering_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Decoration/Desc_CharacterSpin_Statue.Desc_CharacterSpin_Statue_C"',Amount=1))`,
			},
		},
	}

	StatueGoldenNut = FGSchematic{
		ClassName:               "ResourceSink_StatueGoldenNut_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueGoldenNut.ResourceSink_StatueGoldenNut_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1000))`,
		MDescription:            `It's a Golden Nut. It's shiny and its price is totally justified. Don't overthink it, you need one! Purchase a Golden Nut today! (Warning: Ingestion is not recommended).`,
		MDisplayName:            `Golden Nut`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           6.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueGoldPioneer.ResourceSink_StatueGoldPioneer_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueLizardDoggo.ResourceSink_StatueLizardDoggo_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Decor/Statues/UI/Award_Statue_Nut_256.Award_Statue_Nut_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Massage-2ABb.SC_RSS_Massage-2ABb_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Decoration/Desc_GoldenNut_Statue.Desc_GoldenNut_Statue_C"',Amount=1))`,
			},
		},
	}

	StatueHoggo = FGSchematic{
		ClassName:               "ResourceSink_Statue_Hoggo_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Statue_Hoggo.ResourceSink_Statue_Hoggo_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Silver Hog`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           3.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueLizardDoggo.ResourceSink_StatueLizardDoggo_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueSpaceGiraffe.ResourceSink_StatueSpaceGiraffe_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Decor/Statues/UI/Award_Statue_Hog_256.Award_Statue_Hog_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Massage-2ABb.SC_RSS_Massage-2ABb_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Decoration/Desc_Hog_Statue.Desc_Hog_Statue_C"',Amount=1))`,
			},
		},
	}

	StatueLizardDoggo = FGSchematic{
		ClassName:               "ResourceSink_StatueLizardDoggo_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueLizardDoggo.ResourceSink_StatueLizardDoggo_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=100))`,
		MDescription:            `You love Doggos but are not of the responsible breed? Worry not! Purchase the Golden Lizard Doggo today and enjoy cuteness without worrying about it starving!`,
		MDisplayName:            `Lizard Doggo`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           4.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Statue_Hoggo.ResourceSink_Statue_Hoggo_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueSpaceGiraffe.ResourceSink_StatueSpaceGiraffe_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Decor/Statues/UI/Award_Statue_LizardDoggo_256.Award_Statue_LizardDoggo_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Massage-2ABb.SC_RSS_Massage-2ABb_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Decoration/Desc_DoggoStatue.Desc_DoggoStatue_C"',Amount=1))`,
			},
		},
	}

	StatueSilverPioneer = FGSchematic{
		ClassName:               "ResourceSink_StatueSilverPioneer_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueSilverPioneer.ResourceSink_StatueSilverPioneer_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=50))`,
		MDescription:            `Are you up for the challenge? Show your worth with this recognition efforts and proudly showcase it in your factory if you're really sure you deserve it.`,
		MDisplayName:            `Pretty Good Pioneering`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           1.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueBronzePioneer.ResourceSink_StatueBronzePioneer_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueGoldPioneer.ResourceSink_StatueGoldPioneer_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Decor/Statues/UI/Award_Statue_Character_Silver_256.Award_Statue_Character_Silver_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Pioneering.SC_RSS_Pioneering_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Decoration/Desc_CharacterClap_Statue.Desc_CharacterClap_Statue_C"',Amount=1))`,
			},
		},
	}

	StatueSpaceGiraffe = FGSchematic{
		ClassName:               "ResourceSink_StatueSpaceGiraffe_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueSpaceGiraffe.ResourceSink_StatueSpaceGiraffe_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Confusing Creature`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           5.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Statue_Hoggo.ResourceSink_Statue_Hoggo_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StatueLizardDoggo.ResourceSink_StatueLizardDoggo_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Decor/Statues/UI/Award_Statue_SpaceGiraffe_256.Award_Statue_SpaceGiraffe_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Massage-2ABb.SC_RSS_Massage-2ABb_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Decoration/Desc_SpaceGiraffeStatue.Desc_SpaceGiraffeStatue_C"',Amount=1))`,
			},
		},
	}

	SteelBeam = FGSchematic{
		ClassName:               "ResourceSink_SteelBeam_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier3/ResourceSink_SteelBeam.ResourceSink_SteelBeam_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Steel Beam`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/SteelPlate/UI/IconDesc_SteelBeam_256.IconDesc_SteelBeam_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `Texture2D /Game/FactoryGame/Resource/Parts/SteelPlate/UI/IconDesc_SteelBeam_64.IconDesc_SteelBeam_64`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlate/Desc_SteelPlate.Desc_SteelPlate_C"',Amount=100))`,
			},
		},
	}

	SteelPipe = FGSchematic{
		ClassName:               "ResourceSink_SteelPipe_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier3/ResourceSink_SteelPipe.ResourceSink_SteelPipe_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Steel Pipe`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_3-4.Schematic_3-4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/SteelPipe/UI/IconDesc_SteelPipe_256.IconDesc_SteelPipe_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `Texture2D /Game/FactoryGame/Resource/Parts/SteelPipe/UI/IconDesc_SteelPipe_64.IconDesc_SteelPipe_64`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_StandardParts.SC_RSS_StandardParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=100))`,
			},
		},
	}

	StreetLight = FGSchematic{
		ClassName:               "ResourceSink_StreetLight_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_StreetLight.ResourceSink_StreetLight_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            `No more night-time related walkway accidents involving Trucks and poor lighting conditions! Light up those pathways and roads using the FICSIT Street Lights!`,
		MDisplayName:            `Street Light`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_LightTower.ResourceSink_LightTower_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_CeilingLight.ResourceSink_CeilingLight_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Walkways.ResourceSink_Walkways_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Stairs.ResourceSink_Stairs_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/StreetLight/UI/IconDesc_StreetLight_256.IconDesc_StreetLight_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_StreetLight.Recipe_StreetLight_C"')`,
			},
		},
	}

	Sulfur0 = FGSchematic{
		ClassName:               "Research_Sulfur_0_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_0.Research_Sulfur_0_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Sulfur/Desc_Sulfur.Desc_Sulfur_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Sulfur`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/RawResources/Sulfur/UI/Sulfur_64.Sulfur_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Sulfur/Desc_Sulfur.Desc_Sulfur_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Sulfur/Desc_Sulfur.Desc_Sulfur_C"')`,
			},
		},
	}

	Sulfur1 = FGSchematic{
		ClassName:               "Research_Sulfur_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_1.Research_Sulfur_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Sulfur/Desc_Sulfur.Desc_Sulfur_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Coal/Desc_Coal.Desc_Coal_C"',Amount=25))`,
		MDescription:            ``,
		MDisplayName:            `Black Powder`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/GunPowder/UI/IconDesc_Gunpowder_64.IconDesc_Gunpowder_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_Gunpowder.Recipe_Gunpowder_C"')`,
			},
		},
	}

	Sulfur2 = FGSchematic{
		ClassName:               "Research_Sulfur_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_2.Research_Sulfur_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GunPowder/Desc_Gunpowder.Desc_Gunpowder_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Volatile Applications`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/InGame/-Shared/Mam_Key_64.Mam_Key_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Sulfur3 = FGSchematic{
		ClassName:               "Research_Sulfur_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_3.Research_Sulfur_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlate/Desc_SteelPlate.Desc_SteelPlate_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/GemstoneScanner/BP_EquipmentDescriptorObjectScanner.BP_EquipmentDescriptorObjectScanner_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Detonator Prototype`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/NobeliskDetonator/UI/Detonator_64.Detonator_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         180.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Sulfur31 = FGSchematic{
		ClassName:               "Research_Sulfur_3_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_3_1.Research_Sulfur_3_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlateReinforced/Desc_SteelPlateReinforced.Desc_SteelPlateReinforced_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/GemstoneScanner/BP_EquipmentDescriptorObjectScanner.BP_EquipmentDescriptorObjectScanner_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Nobelisk Detonator`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/NobeliskDetonator/UI/Detonator_64.Detonator_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         180.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_NobeliskDetonator.Recipe_NobeliskDetonator_C"')`,
			},
		},
	}

	Sulfur32 = FGSchematic{
		ClassName:               "Research_Sulfur_3_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_3_2.Research_Sulfur_3_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GunPowder/Desc_Gunpowder.Desc_Gunpowder_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Prototype Explosives`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/NobeliskExplosive/UI/IconDesc_Explosive_64.IconDesc_Explosive_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         180.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Sulfur321 = FGSchematic{
		ClassName:               "Research_Sulfur_3_2_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_3_2_1.Research_Sulfur_3_2_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GunPowder/Desc_Gunpowder.Desc_Gunpowder_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Nobelisk Explosives`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/NobeliskExplosive/UI/IconDesc_Explosive_64.IconDesc_Explosive_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_Nobelisk.Recipe_Nobelisk_C"')`,
			},
		},
	}

	Sulfur4 = FGSchematic{
		ClassName:               "Research_Sulfur_4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_4.Research_Sulfur_4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CircuitBoard/Desc_CircuitBoard.Desc_CircuitBoard_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronScrew/Desc_IronScrew.Desc_IronScrew_C"',Amount=250))`,
		MDescription:            ``,
		MDisplayName:            `Rifle Prototype`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/Rifle/UI/RifleMK1_64.RifleMK1_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         180.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Sulfur41 = FGSchematic{
		ClassName:               "Research_Sulfur_4_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_4_1.Research_Sulfur_4_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/CircuitBoard/Desc_CircuitBoard.Desc_CircuitBoard_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ModularFrameHeavy/Desc_ModularFrameHeavy.Desc_ModularFrameHeavy_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `The Rifle`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Equipment/Rifle/UI/RifleMK1_64.RifleMK1_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         180.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_SpaceRifleMk1.Recipe_SpaceRifleMk1_C"')`,
			},
		},
	}

	Sulfur42 = FGSchematic{
		ClassName:               "Research_Sulfur_4_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_4_2.Research_Sulfur_4_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GunPowder/Desc_Gunpowder.Desc_Gunpowder_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Equipment/Beacon/BP_EquipmentDescriptorBeacon.BP_EquipmentDescriptorBeacon_C"',Amount=5))`,
		MDescription:            ``,
		MDisplayName:            `Prototype Cartridges`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/CartridgeStandard/UI/Rifle_Magazine_64.Rifle_Magazine_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         180.000000,
		MType:                   `EST_MAM`,
		MUnlocks:                nil,
	}

	Sulfur421 = FGSchematic{
		ClassName:               "Research_Sulfur_4_2_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_4_2_1.Research_Sulfur_4_2_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GunPowder/Desc_Gunpowder.Desc_Gunpowder_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPipe/Desc_SteelPipe.Desc_SteelPipe_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Rubber/Desc_Rubber.Desc_Rubber_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Rifle Cartridges`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/CartridgeStandard/UI/Rifle_Magazine_64.Rifle_Magazine_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_Cartridge.Recipe_Cartridge_C"')`,
			},
		},
	}

	Sulfur5 = FGSchematic{
		ClassName:               "Research_Sulfur_5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_5.Research_Sulfur_5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GunPowder/Desc_Gunpowder.Desc_Gunpowder_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlateReinforced/Desc_SteelPlateReinforced.Desc_SteelPlateReinforced_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Expanded Toolbelt`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/Assets/Shared/ThumbsUp_64.ThumbsUp_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         180.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                         `BP_UnlockArmEquipmentSlot_C`,
				MNumArmEquipmentSlotsToUnlock: 1,
			},
		},
	}

	Sulfur6 = FGSchematic{
		ClassName:               "Research_Sulfur_6_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/Sulfur_RS/Research_Sulfur_6.Research_Sulfur_6_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/GunPowder/Desc_Gunpowder.Desc_Gunpowder_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/SteelPlate/Desc_SteelPlate.Desc_SteelPlate_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `Inflated Pocket Dimension`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Interface/UI/Assets/Shared/ThumbsUp_64.ThumbsUp_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         180.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 6,
			},
		},
	}

	SuperComputer = FGSchematic{
		ClassName:               "ResourceSink_SuperComputer_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier7/ResourceSink_SuperComputer.ResourceSink_SuperComputer_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=8))`,
		MDescription:            ``,
		MDisplayName:            `Supercomputer`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_6_1.Research_Caterium_6_1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/ComputerSuper/UI/IconDesc_SuperComputer_256.IconDesc_SuperComputer_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Communications.SC_RSS_Communications_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ComputerSuper/Desc_ComputerSuper.Desc_ComputerSuper_C"',Amount=50))`,
			},
		},
	}

	TurboMotor = FGSchematic{
		ClassName:               "ResourceSink_TurboMotor_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier7/ResourceSink_TurboMotor.ResourceSink_TurboMotor_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=8))`,
		MDescription:            ``,
		MDisplayName:            `Turbo Motor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Progression/Schematic_7-2.Schematic_7-2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/TurboMotor/UI/IconDesc_TurboMotor_256.IconDesc_TurboMotor_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_IndustrialParts.SC_RSS_IndustrialParts_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/MotorLightweight/Desc_MotorLightweight.Desc_MotorLightweight_C"',Amount=50))`,
			},
		},
	}

	Tutorial1 = FGSchematic{
		ClassName:               "Schematic_Tutorial1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial1.Schematic_Tutorial1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `HUB Upgrade 1`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Hub_1.SchematicIcon_Hub_1"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               0,
		MTimeToComplete:         0.000000,
		MType:                   `EST_Tutorial`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_Workshop.Recipe_Workshop_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_PortableMiner.Recipe_PortableMiner_C"')`,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Tutorial15 = FGSchematic{
		ClassName:               "Schematic_Tutorial1_5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial1_5.Schematic_Tutorial1_5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=20),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `HUB Upgrade 2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial1.Schematic_Tutorial1_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Hub_1.SchematicIcon_Hub_1"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Tutorial`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_SmelterBasicMk1.Recipe_SmelterBasicMk1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerLine.Recipe_PowerLine_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Smelter/Recipe_IngotCopper.Recipe_IngotCopper_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Wire.Recipe_Wire_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Cable.Recipe_Cable_C"')`,
			},
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreCopper/Desc_OreCopper.Desc_OreCopper_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/OreCopper/Desc_OreCopper.Desc_OreCopper_C"')`,
			},
		},
	}

	Tutorial2 = FGSchematic{
		ClassName:               "Schematic_Tutorial2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial2.Schematic_Tutorial2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=20),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=20),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=20))`,
		MDescription:            ``,
		MDisplayName:            `HUB Upgrade 3`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial1_5.Schematic_Tutorial1_5_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Hub_2.SchematicIcon_Hub_2"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Tutorial`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConstructorMk1.Recipe_ConstructorMk1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerPoleMk1.Recipe_PowerPoleMk1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Concrete.Recipe_Concrete_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Screw.Recipe_Screw_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Assembler/Recipe_IronPlateReinforced.Recipe_IronPlateReinforced_C"')`,
			},
			{
				Class:                        `BP_UnlockScannableResource_C`,
				MResourcePairsToAddToScanner: `((ResourceDescriptor=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Stone/Desc_Stone.Desc_Stone_C"'))`,
				MResourcesToAddToScanner:     `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/RawResources/Stone/Desc_Stone.Desc_Stone_C"')`,
			},
		},
	}

	Tutorial3 = FGSchematic{
		ClassName:               "Schematic_Tutorial3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial3.Schematic_Tutorial3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=75),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=20),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `HUB Upgrade 4`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial2.Schematic_Tutorial2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Hub_3.SchematicIcon_Hub_3"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Tutorial`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorPole.Recipe_ConveyorPole_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_ConveyorBeltMk1.Recipe_ConveyorBeltMk1_C"')`,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Tutorial4 = FGSchematic{
		ClassName:               "Schematic_Tutorial4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial4.Schematic_Tutorial4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=75),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cable/Desc_Cable.Desc_Cable_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=20))`,
		MDescription:            ``,
		MDisplayName:            `HUB Upgrade 5`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial3.Schematic_Tutorial3_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Hub_4.SchematicIcon_Hub_4"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Tutorial`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_MinerMk1.Recipe_MinerMk1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_StorageContainerMk1.Recipe_StorageContainerMk1_C"')`,
			},
			{
				Class:                      `BP_UnlockInventorySlot_C`,
				MNumInventorySlotsToUnlock: 3,
			},
		},
	}

	Tutorial5 = FGSchematic{
		ClassName:               "Schematic_Tutorial5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial5.Schematic_Tutorial5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronRod/Desc_IronRod.Desc_IronRod_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Cement/Desc_Cement.Desc_Cement_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `HUB Upgrade 6`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Tutorial/Schematic_Tutorial4.Schematic_Tutorial4_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/TradingPost/UI/SchematicIcons/SchematicIcon_Hub_5.SchematicIcon_Hub_5"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      ``,
		MTechTier:           0,
		MTimeToComplete:     0.000000,
		MType:               `EST_Tutorial`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/TowTruck/Recipe_SpaceElevator.Recipe_SpaceElevator_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_GeneratorBiomass.Recipe_GeneratorBiomass_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Biomass_Wood.Recipe_Biomass_Wood_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Constructor/Recipe_Biomass_Leaves.Recipe_Biomass_Leaves_C"')`,
			},
		},
	}

	Walkways = FGSchematic{
		ClassName:               "ResourceSink_Walkways_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Walkways.ResourceSink_Walkways_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=6))`,
		MDescription:            `Ever wanted to admire your factories from a safe and convenient place? Build some Walkways and get inspired by your own labour to do even better!`,
		MDisplayName:            `Walkways`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           2.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FactoryRailing.ResourceSink_FactoryRailing_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_Ladders.ResourceSink_Ladders_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Walkway/UI/WalkwayStraight_256.WalkwayStraight_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walkways/Recipe_Walkway_Straight.Recipe_Walkway_Straight_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walkways/Recipe_Walkway_Ramp.Recipe_Walkway_Ramp_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walkways/Recipe_Walkway_Turn.Recipe_Walkway_Turn_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walkways/Recipe_Walkway_T.Recipe_Walkway_T_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walkways/Recipe_Walkway_Cross.Recipe_Walkway_Cross_C"')`,
			},
		},
	}

	WallPowerPoles = FGSchematic{
		ClassName:               "ResourceSink_WallPowerPoles_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPoles.ResourceSink_WallPowerPoles_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=4))`,
		MDescription:            `Tired of bulky power poles? Save space and organize an efficient power grid in your factory with the all new and revolutionary Power Outlets for walls and ceilings!`,
		MDisplayName:            `Wall Power Outlets Mk.1`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           3.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPolesMK2.ResourceSink_WallPowerPolesMK2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPolesMK3.ResourceSink_WallPowerPolesMK3_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveyorWallMount.ResourceSink_ConveyorWallMount_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_PipelineWallAttachments.ResourceSink_PipelineWallAttachments_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_HyperTubeWallAttachements.ResourceSink_HyperTubeWallAttachements_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/PowerPoleWall/UI/PowerPoleWall_MK1_256.PowerPoleWall_MK1_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerPoleWall.Recipe_PowerPoleWall_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerPoleWallDouble.Recipe_PowerPoleWallDouble_C"')`,
			},
		},
	}

	WallPowerPolesMK2 = FGSchematic{
		ClassName:               "ResourceSink_WallPowerPolesMK2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPolesMK2.ResourceSink_WallPowerPolesMK2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=5))`,
		MDescription:            `Tired of bulky power poles? Save space and organize an efficient power grid in your factory with the all new and revolutionary Power Outlets for walls and ceilings!`,
		MDisplayName:            `Wall Power Outlets Mk.2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           4.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPoles.ResourceSink_WallPowerPoles_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPolesMK3.ResourceSink_WallPowerPolesMK3_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveyorWallMount.ResourceSink_ConveyorWallMount_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_PipelineWallAttachments.ResourceSink_PipelineWallAttachments_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_HyperTubeWallAttachements.ResourceSink_HyperTubeWallAttachements_C"')`,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_4_2.Research_Caterium_4_2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/PowerPoleWall/UI/PowerPoleWall_MK2_256.PowerPoleWall_MK2_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerPoleWallMk2.Recipe_PowerPoleWallMk2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerPoleWallDoubleMk2.Recipe_PowerPoleWallDoubleMk2_C"')`,
			},
		},
	}

	WallPowerPolesMK3 = FGSchematic{
		ClassName:               "ResourceSink_WallPowerPolesMK3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPolesMK3.ResourceSink_WallPowerPolesMK3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=6))`,
		MDescription:            `Tired of bulky power poles? Save space and organize an efficient power grid in your factory with the all new and revolutionary Power Outlets for walls and ceilings!`,
		MDisplayName:            `Wall Power Outlets Mk.3`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           5.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPoles.ResourceSink_WallPowerPoles_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WallPowerPolesMK2.ResourceSink_WallPowerPolesMK2_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_ConveyorWallMount.ResourceSink_ConveyorWallMount_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_PipelineWallAttachments.ResourceSink_PipelineWallAttachments_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_HyperTubeWallAttachements.ResourceSink_HyperTubeWallAttachements_C"')`,
		MSchematicDependencies: []struct {
			Class                              string
			MGamePhase                         string
			MRequireAllSchematicsToBePurchased bool
			MSchematics                        string
		}{
			{
				Class:                              `BP_SchematicPurchasedDependency_C`,
				MRequireAllSchematicsToBePurchased: true,
				MSchematics:                        `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/Research/Caterium_RS/Research_Caterium_6_2.Research_Caterium_6_2_C"')`,
			},
		},
		MSchematicIcon:      `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Factory/PowerPoleWall/UI/PowerPoleWall_MK3_256.PowerPoleWall_MK3_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon: `None`,
		MSubCategories:      `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:           1,
		MTimeToComplete:     0.000000,
		MType:               `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerPoleWallMk3.Recipe_PowerPoleWallMk3_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Recipe_PowerPoleWallDoubleMk3.Recipe_PowerPoleWallDoubleMk3_C"')`,
			},
		},
	}

	WindowedWalls = FGSchematic{
		ClassName:               "ResourceSink_WindowedWalls_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/ResourceSink_WindowedWalls.ResourceSink_WindowedWalls_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=9))`,
		MDescription:            ``,
		MDisplayName:            `Windowed Walls`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           4.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DoorWalls_Normal.ResourceSink_DoorWalls_Normal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_DoorWalls_Metal.ResourceSink_DoorWalls_Metal_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSink/ResourceSink_FrameworkFoundations.ResourceSink_FrameworkFoundations_C"')`,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=512.000000,Y=512.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Buildable/Building/Wall/UI/Wall_Window_8x4_01_512.Wall_Window_8x4_01_512"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Walls.SC_RSS_Walls_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Window_8x4_01.Recipe_Wall_Window_8x4_01_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Window_8x4_02.Recipe_Wall_Window_8x4_02_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Window_8x4_03.Recipe_Wall_Window_8x4_03_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Buildings/Walls/Recipe_Wall_Window_8x4_04.Recipe_Wall_Window_8x4_04_C"')`,
			},
		},
	}

	Wire = FGSchematic{
		ClassName:               "ResourceSink_Wire_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/ResourceSink/Parts/Tier1/ResourceSink_Wire.ResourceSink_Wire_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/ResourceSinkCoupon/Desc_ResourceSinkCoupon.Desc_ResourceSinkCoupon_C"',Amount=1))`,
		MDescription:            ``,
		MDisplayName:            `Wire`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         ``,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Resource/Parts/Wire/UI/IconDesc_Wire_256.IconDesc_Wire_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Schematics/ResourceSinkShopCategories/SC_RSS_Electronics.SC_RSS_Electronics_C"')`,
		MTechTier:               1,
		MTimeToComplete:         0.000000,
		MType:                   `EST_ResourceSink`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:        `BP_UnlockGiveItem_C`,
				MItemsToGive: `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Wire/Desc_Wire.Desc_Wire_C"',Amount=500))`,
			},
		},
	}

	XMas1 = FGSchematic{
		ClassName:               "Research_XMas_1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_1.Research_XMas_1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_Gift.Desc_Gift_C"',Amount=100))`,
		MDescription:            ``,
		MDisplayName:            `FICSMAS Tree Base`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Gift_64.IconDesc_Gift_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Gift_64.IconDesc_Gift_64`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/TreeDecor/Recipe_XMassTree.Recipe_XMassTree_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_XmasBranch.Recipe_XmasBranch_C"')`,
			},
		},
	}

	XMas1_1 = FGSchematic{
		ClassName:               "Research_XMas_1-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_1-1.Research_XMas_1-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBranch.Desc_XmasBranch_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `Candy Cane Basher`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_CaneEquipment_64.IconDesc_CaneEquipment_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Equipment/Recipe_CandyCaneBasher.Recipe_CandyCaneBasher_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_CandyCane.Recipe_CandyCane_C"')`,
			},
		},
	}

	XMas1_2 = FGSchematic{
		ClassName:               "Research_XMas_1-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_1-2.Research_XMas_1-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_CandyCane.Desc_CandyCane_C"',Amount=10))`,
		MDescription:            ``,
		MDisplayName:            `Candy Cane Decor`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Buildings/UI/IconDesc_CandyCaneDecor_256.IconDesc_CandyCaneDecor_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/CandyCaneDecor/Recipe_CandyCaneDecor.Recipe_CandyCaneDecor_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_XmasBow.Recipe_XmasBow_C"')`,
			},
		},
	}

	XMas2 = FGSchematic{
		ClassName:               "Research_XMas_2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_2.Research_XMas_2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBall1.Desc_XmasBall1_C"',Amount=1),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_CandyCane.Desc_CandyCane_C"',Amount=20),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBow.Desc_XmasBow_C"',Amount=30))`,
		MDescription:            ``,
		MDisplayName:            `Giant FICSMAS Tree: Upgrade 1`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasBall_Red_64.IconDesc_XmasBall_Red_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/TreeDecor/Schematic_XMassTree_T1.Schematic_XMassTree_T1_C"')`,
			},
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_XmasBall1.Recipe_XmasBall1_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_XmasBall2.Recipe_XmasBall2_C"')`,
			},
		},
	}

	XMas2_1 = FGSchematic{
		ClassName:               "Research_XMas_2-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_2-1.Research_XMas_2-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_CandyCane.Desc_CandyCane_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBow.Desc_XmasBow_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `A friend`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Buildings/UI/IconDesc_Snowman_256.IconDesc_Snowman_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Events/Christmas/Buildings/UI/IconDesc_Snowman_256.IconDesc_Snowman_256`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/SnowmanDecor/Recipe_Snowman.Recipe_Snowman_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_Snow.Recipe_Snow_C"')`,
			},
		},
	}

	XMas2_2 = FGSchematic{
		ClassName:               "Research_XMas_2-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_2-2.Research_XMas_2-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBall1.Desc_XmasBall1_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBall2.Desc_XmasBall2_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBranch.Desc_XmasBranch_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `FICSMAS Gift Tree`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_ChristmasTree_256.IconDesc_ChristmasTree_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Holiday/Recipe_TreeGiftProducer.Recipe_TreeGiftProducer_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_XmasBall3.Recipe_XmasBall3_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_XmasBall4.Recipe_XmasBall4_C"')`,
			},
		},
	}

	XMas3 = FGSchematic{
		ClassName:               "Research_XMas_3_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_3.Research_XMas_3_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBallCluster.Desc_XmasBallCluster_C"',Amount=1),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBall3.Desc_XmasBall3_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBall4.Desc_XmasBall4_C"',Amount=200))`,
		MDescription:            ``,
		MDisplayName:            `Giant FICSMAS Tree: Upgrade 2`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Balls_64.IconDesc_Balls_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Balls_64.IconDesc_Balls_64`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/TreeDecor/Schematic_XMassTree_T2.Schematic_XMassTree_T2_C"')`,
			},
		},
	}

	XMas3_1 = FGSchematic{
		ClassName:               "Research_XMas_3-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_3-1.Research_XMas_3-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBall1.Desc_XmasBall1_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBall2.Desc_XmasBall2_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBall3.Desc_XmasBall3_C"',Amount=50),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBall4.Desc_XmasBall4_C"',Amount=50))`,
		MDescription:            ``,
		MDisplayName:            `FICSMAS Lights`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_XmasLights_256.IconDesc_XmasLights_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/PowerLineLights/Recipe_xmassLights.Recipe_xmassLights_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_XmasBallCluster.Recipe_XmasBallCluster_C"')`,
			},
		},
	}

	XMas3_2 = FGSchematic{
		ClassName:               "Research_XMas_3-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_3-2.Research_XMas_3-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBallCluster.Desc_XmasBallCluster_C"',Amount=10),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_Snow.Desc_Snow_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBranch.Desc_XmasBranch_C"',Amount=500))`,
		MDescription:            ``,
		MDisplayName:            `It's snowing!`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=256.000000,Y=256.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Buildings/UI/IconDesc_SnowDispenser_256.IconDesc_SnowDispenser_256"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Events/Christmas/Buildings/UI/IconDesc_SnowDispenser_256.IconDesc_SnowDispenser_256`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/SnowDispenser/Recipe_SnowDispenser.Recipe_SnowDispenser_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_XmasWreath.Recipe_XmasWreath_C"')`,
			},
		},
	}

	XMas4 = FGSchematic{
		ClassName:               "Research_XMas_4_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_4.Research_XMas_4_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasStar.Desc_XmasStar_C"',Amount=1),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasWreath.Desc_XmasWreath_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBallCluster.Desc_XmasBallCluster_C"',Amount=400))`,
		MDescription:            ``,
		MDisplayName:            `Giant FICSMAS Tree: Upgrade 3`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Star_64.IconDesc_Star_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Star_64.IconDesc_Star_64`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/TreeDecor/Schematic_XMassTree_T3.Schematic_XMassTree_T3_C"')`,
			},
		},
	}

	XMas4_1 = FGSchematic{
		ClassName:               "Research_XMas_4-1_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_4-1.Research_XMas_4-1_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasWreath.Desc_XmasWreath_C"',Amount=100),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBallCluster.Desc_XmasBallCluster_C"',Amount=200),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBow.Desc_XmasBow_C"',Amount=500))`,
		MDescription:            ``,
		MDisplayName:            `FICSMAS Wreath`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Wreath_64.IconDesc_Wreath_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Wreath_64.IconDesc_Wreath_64`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/WreathDecor/Recipe_WreathDecor.Recipe_WreathDecor_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Recipe_XmasStar.Recipe_XmasStar_C"')`,
			},
		},
	}

	XMas4_2 = FGSchematic{
		ClassName:               "Research_XMas_4-2_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_4-2.Research_XMas_4-2_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_Snow.Desc_Snow_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_CandyCane.Desc_CandyCane_C"',Amount=500),(ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasBow.Desc_XmasBow_C"',Amount=500))`,
		MDescription:            ``,
		MDisplayName:            `Snowfight!`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_SnowballMittens_64.IconDesc_SnowballMittens_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `None`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:    `BP_UnlockRecipe_C`,
				MRecipes: `(BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Holiday/Recipe_SnowballWeapon.Recipe_SnowballWeapon_C"',BlueprintGeneratedClass'"/Game/FactoryGame/Recipes/Holiday/Recipe_Snowball.Recipe_Snowball_C"')`,
			},
		},
	}

	XMas5 = FGSchematic{
		ClassName:               "Research_XMas_5_C",
		FullName:                `BlueprintGeneratedClass /Game/FactoryGame/Schematics/Research/XMas_RS/Research_XMas_5.Research_XMas_5_C`,
		MCost:                   `((ItemClass=BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Parts/Desc_XmasStar.Desc_XmasStar_C"',Amount=500))`,
		MDescription:            ``,
		MDisplayName:            `Giant FICSMAS Tree: Upgrade 4`,
		MIncludeInBuilds:        `IIB_PublicBuilds`,
		MMenuPriority:           0.000000,
		MRelevantEvents:         `(EV_Christmas)`,
		MRelevantShopSchematics: ``,
		MSchematicDependencies:  nil,
		MSchematicIcon:          `(ImageSize=(X=64.000000,Y=64.000000),Margin=(),TintColor=(SpecifiedColor=(R=1.000000,G=1.000000,B=1.000000,A=1.000000)),ResourceObject=Texture2D'"/Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Star_64.IconDesc_Star_64"',UVRegion=(Min=(X=0.000000,Y=0.000000),Max=(X=0.000000,Y=0.000000),bIsValid=0),DrawAs=Image)`,
		MSmallSchematicIcon:     `Texture2D /Game/FactoryGame/Events/Christmas/Parts/UI/IconDesc_Star_64.IconDesc_Star_64`,
		MSubCategories:          ``,
		MTechTier:               3,
		MTimeToComplete:         3.000000,
		MType:                   `EST_MAM`,
		MUnlocks: []struct {
			Class                         string
			MItemsToGive                  string
			MNumArmEquipmentSlotsToUnlock int
			MNumInventorySlotsToUnlock    int
			MRecipes                      string
			MResourcePairsToAddToScanner  string
			MResourcesToAddToScanner      string
			MSchematics                   string
		}{
			{
				Class:       `BP_UnlockSchematic_C`,
				MSchematics: `(BlueprintGeneratedClass'"/Game/FactoryGame/Events/Christmas/Buildings/TreeDecor/Schematic_XMassTree_T4.Schematic_XMassTree_T4_C"')`,
			},
		},
	}
)

func GetByClassName(className string) (*FGSchematic, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGSchematic with name %s", className)
}

var classNameToVar = map[string]*FGSchematic{
	"Research_ACarapace_0_C":                           &ACarapace0,
	"Research_ACarapace_1_C":                           &ACarapace1,
	"Research_ACarapace_2_C":                           &ACarapace2,
	"Research_ACarapace_2_1_C":                         &ACarapace21,
	"Research_ACarapace_3_C":                           &ACarapace3,
	"ResourceSink_AILimiter_C":                         &AILimiter,
	"Research_AOrganisms_1_C":                          &AOrganisms1,
	"Research_AOrganisms_2_C":                          &AOrganisms2,
	"Research_AOrgans_0_C":                             &AOrgans0,
	"Research_AOrgans_1_C":                             &AOrgans1,
	"Research_AOrgans_2_C":                             &AOrgans2,
	"Research_AOrgans_3_C":                             &AOrgans3,
	"ResourceSink_AdvancedAmmoPack_C":                  &AdvancedAmmoPack,
	"ResourceSink_AlcladSheet_C":                       &AlcladSheet,
	"Schematic_Alternate_AdheredIronPlate_C":           &AlternateAdheredIronPlate,
	"Schematic_Alternate_AlcladCasing_C":               &AlternateAlcladCasing,
	"Schematic_Alternate_AutomatedMiner_C":             &AlternateAutomatedMiner,
	"Schematic_Alternate_Beacon1_C":                    &AlternateBeacon1,
	"Schematic_Alternate_BoltedFrame_C":                &AlternateBoltedFrame,
	"Schematic_Alternate_Cable1_C":                     &AlternateCable1,
	"Schematic_Alternate_Cable2_C":                     &AlternateCable2,
	"Schematic_Alternate_CircuitBoard1_C":              &AlternateCircuitBoard1,
	"Schematic_Alternate_CircuitBoard2_C":              &AlternateCircuitBoard2,
	"Schematic_Alternate_ClassicBattery_C":             &AlternateClassicBattery,
	"Schematic_Alternate_Coal1_C":                      &AlternateCoal1,
	"Schematic_Alternate_Coal2_C":                      &AlternateCoal2,
	"Schematic_Alternate_CoatedCable_C":                &AlternateCoatedCable,
	"Schematic_Alternate_CoatedIronCanister_C":         &AlternateCoatedIronCanister,
	"Schematic_Alternate_CoatedIronPlate_C":            &AlternateCoatedIronPlate,
	"Schematic_Alternate_CokeSteelIngot_C":             &AlternateCokeSteelIngot,
	"Schematic_Alternate_Computer1_C":                  &AlternateComputer1,
	"Schematic_Alternate_Computer2_C":                  &AlternateComputer2,
	"Schematic_Alternate_Concrete_C":                   &AlternateConcrete,
	"Schematic_Alternate_CoolingDevice_C":              &AlternateCoolingDevice,
	"Schematic_Alternate_CopperAlloyIngot_C":           &AlternateCopperAlloyIngot,
	"Schematic_Alternate_CopperRotor_C":                &AlternateCopperRotor,
	"Schematic_Alternate_CrystalOscillator_C":          &AlternateCrystalOscillator,
	"Schematic_Alternate_DilutedFuel_C":                &AlternateDilutedFuel,
	"Schematic_Alternate_DilutedPackagedFuel_C":        &AlternateDilutedPackagedFuel,
	"Schematic_Alternate_ElectricMotor_C":              &AlternateElectricMotor,
	"Schematic_Alternate_ElectroAluminumScrap_C":       &AlternateElectroAluminumScrap,
	"Schematic_Alternate_ElectrodeCircuitBoard_C":      &AlternateElectrodeCircuitBoard,
	"Schematic_Alternate_ElectromagneticControlRod1_C": &AlternateElectromagneticControlRod1,
	"Schematic_Alternate_EnrichedCoal_C":               &AlternateEnrichedCoal,
	"Schematic_Alternate_FertileUranium_C":             &AlternateFertileUranium,
	"Schematic_Alternate_FlexibleFramework_C":          &AlternateFlexibleFramework,
	"Schematic_Alternate_FusedWire_C":                  &AlternateFusedWire,
	"Schematic_Alternate_Gunpowder1_C":                 &AlternateGunpowder1,
	"Schematic_Alternate_HeatFusedFrame_C":             &AlternateHeatFusedFrame,
	"Schematic_Alternate_HeatSink1_C":                  &AlternateHeatSink1,
	"Schematic_Alternate_HeavyFlexibleFrame_C":         &AlternateHeavyFlexibleFrame,
	"Schematic_Alternate_HeavyModularFrame_C":          &AlternateHeavyModularFrame,
	"Schematic_Alternate_HeavyOilResidue_C":            &AlternateHeavyOilResidue,
	"Schematic_Alternate_HighSpeedConnector_C":         &AlternateHighSpeedConnector,
	"Schematic_Alternate_HighSpeedWiring_C":            &AlternateHighSpeedWiring,
	"Schematic_Alternate_IngotIron_C":                  &AlternateIngotIron,
	"Schematic_Alternate_IngotSteel1_C":                &AlternateIngotSteel1,
	"Schematic_Alternate_IngotSteel2_C":                &AlternateIngotSteel2,
	"Schematic_Alternate_InstantPlutoniumCell_C":       &AlternateInstantPlutoniumCell,
	"Schematic_Alternate_InstantScrap_C":               &AlternateInstantScrap,
	"Schematic_Alternate_InventorySlots1_C":            &AlternateInventorySlots1,
	"Schematic_Alternate_InventorySlots2_C":            &AlternateInventorySlots2,
	"Schematic_Alternate_ModularFrame_C":               &AlternateModularFrame,
	"Schematic_Alternate_Motor1_C":                     &AlternateMotor1,
	"Schematic_Alternate_Nobelisk1_C":                  &AlternateNobelisk1,
	"Schematic_Alternate_NuclearFuelRod1_C":            &AlternateNuclearFuelRod1,
	"Schematic_Alternate_OCSupercomputer_C":            &AlternateOCSupercomputer,
	"Schematic_Alternate_Plastic1_C":                   &AlternatePlastic1,
	"Schematic_Alternate_PlasticSmartPlating_C":        &AlternatePlasticSmartPlating,
	"Schematic_Alternate_PlutoniumFuelUnit_C":          &AlternatePlutoniumFuelUnit,
	"Schematic_Alternate_PolyesterFabric_C":            &AlternatePolyesterFabric,
	"Schematic_Alternate_PolymerResin_C":               &AlternatePolymerResin,
	"Schematic_Alternate_PureAluminumIngot_C":          &AlternatePureAluminumIngot,
	"Schematic_Alternate_PureCateriumIngot_C":          &AlternatePureCateriumIngot,
	"Schematic_Alternate_PureCopperIngot_C":            &AlternatePureCopperIngot,
	"Schematic_Alternate_PureIronIngot_C":              &AlternatePureIronIngot,
	"Schematic_Alternate_PureQuartzCrystal_C":          &AlternatePureQuartzCrystal,
	"Schematic_Alternate_Quickwire_C":                  &AlternateQuickwire,
	"Schematic_Alternate_RadioControlSystem_C":         &AlternateRadioControlSystem,
	"Schematic_Alternate_RadioControlUnit1_C":          &AlternateRadioControlUnit1,
	"Schematic_Alternate_RecycledRubber_C":             &AlternateRecycledRubber,
	"Schematic_Alternate_ReinforcedIronPlate1_C":       &AlternateReinforcedIronPlate1,
	"Schematic_Alternate_ReinforcedIronPlate2_C":       &AlternateReinforcedIronPlate2,
	"Schematic_Alternate_ReinforcedSteelPlate_C":       &AlternateReinforcedSteelPlate,
	"Schematic_Alternate_Rotor_C":                      &AlternateRotor,
	"Schematic_Alternate_RubberConcrete_C":             &AlternateRubberConcrete,
	"Schematic_Alternate_Screw_C":                      &AlternateScrew,
	"Schematic_Alternate_Screw2_C":                     &AlternateScrew2,
	"Schematic_Alternate_Silica_C":                     &AlternateSilica,
	"Schematic_Alternate_SloppyAlumina_C":              &AlternateSloppyAlumina,
	"Schematic_Alternate_Stator_C":                     &AlternateStator,
	"Schematic_Alternate_SteamedCopperSheet_C":         &AlternateSteamedCopperSheet,
	"Schematic_Alternate_SteelCanister_C":              &AlternateSteelCanister,
	"Schematic_Alternate_SteelCoatedPlate_C":           &AlternateSteelCoatedPlate,
	"Schematic_Alternate_SteelRod_C":                   &AlternateSteelRod,
	"Schematic_Alternate_SuperStateComputer_C":         &AlternateSuperStateComputer,
	"Schematic_Alternate_TurboBlendFuel_C":             &AlternateTurboBlendFuel,
	"Schematic_Alternate_TurboFuel_C":                  &AlternateTurboFuel,
	"Schematic_Alternate_TurboHeavyFuel_C":             &AlternateTurboHeavyFuel,
	"Schematic_Alternate_TurboMotor1_C":                &AlternateTurboMotor1,
	"Schematic_Alternate_TurboPressureMotor_C":         &AlternateTurboPressureMotor,
	"Schematic_Alternate_UraniumCell1_C":               &AlternateUraniumCell1,
	"Schematic_Alternate_WetConcrete_C":                &AlternateWetConcrete,
	"Schematic_Alternate_Wire1_C":                      &AlternateWire1,
	"Schematic_Alternate_Wire2_C":                      &AlternateWire2,
	"ResourceSink_AmmoPack_C":                          &AmmoPack,
	"ResourceSink_Battery_C":                           &Battery,
	"ResourceSink_Beacons_C":                           &Beacons,
	"ResourceSink_Biofuel_C":                           &Biofuel,
	"ResourceSink_Biomass_C":                           &Biomass,
	"ResourceSink_BlackPowder_C":                       &BlackPowder,
	"ResourceSink_Cable_C":                             &Cable,
	"Research_Caterium_0_C":                            &Caterium0,
	"Research_Caterium_1_C":                            &Caterium1,
	"Research_Caterium_2_C":                            &Caterium2,
	"Research_Caterium_2_1_C":                          &Caterium21,
	"Research_Caterium_3_C":                            &Caterium3,
	"Research_Caterium_3_1_C":                          &Caterium31,
	"Research_Caterium_4_1_C":                          &Caterium41,
	"Research_Caterium_4_1_1_C":                        &Caterium411,
	"Research_Caterium_4_1_2_C":                        &Caterium412,
	"Research_Caterium_4_2_C":                          &Caterium42,
	"Research_Caterium_4_3_C":                          &Caterium43,
	"Research_Caterium_5_C":                            &Caterium5,
	"Research_Caterium_6_1_C":                          &Caterium61,
	"Research_Caterium_6_2_C":                          &Caterium62,
	"Research_Caterium_6_3_C":                          &Caterium63,
	"Research_Caterium_7_1_C":                          &Caterium71,
	"Research_Caterium_7_2_C":                          &Caterium72,
	"ResourceSink_CeilingLight_C":                      &CeilingLight,
	"ResourceSink_CircuitBoard_C":                      &CircuitBoard,
	"ResourceSink_CoffeeCup_C":                         &CoffeeCup,
	"ResourceSink_ColoredAmmoPack_C":                   &ColoredAmmoPack,
	"ResourceSink_Computer_C":                          &Computer,
	"ResourceSink_Concrete_C":                          &Concrete,
	"ResourceSink_ConveryWalls_Metal_C":                &ConveryWallsMetal,
	"ResourceSink_ConveryWalls_Normal_C":               &ConveryWallsNormal,
	"ResourceSink_ConveyorWallMount_C":                 &ConveyorWallMount,
	"ResourceSink_CopperSheet_C":                       &CopperSheet,
	"ResourceSink_CrystalOscillator_C":                 &CrystalOscillator,
	"ResourceSink_CurvedFoundationPack_C":              &CurvedFoundationPack,
	"ResourceSink_CyberWagon_C":                        &CyberWagon,
	"ResourceSink_CyberWagon_Unlock_C":                 &CyberWagonUnlock,
	"ResourceSink_DiagonalRamps_C":                     &DiagonalRamps,
	"ResourceSink_DoorWalls_Metal_C":                   &DoorWallsMetal,
	"ResourceSink_DoorWalls_Normal_C":                  &DoorWallsNormal,
	"ResourceSink_EmptyCanister_C":                     &EmptyCanister,
	"ResourceSink_EncasedIndustrialBeam_C":             &EncasedIndustrialBeam,
	"ResourceSink_Explosives_C":                        &Explosives,
	"ResourceSink_Fabric_C":                            &Fabric,
	"ResourceSink_FactoryCart_C":                       &FactoryCart,
	"ResourceSink_FactoryRailing_C":                    &FactoryRailing,
	"Research_FlowerPetals_1_C":                        &FlowerPetals1,
	"Research_FlowerPetals_2_C":                        &FlowerPetals2,
	"Research_FlowerPetals_3_C":                        &FlowerPetals3,
	"ResourceSink_FoudationPillar_C":                   &FoudationPillar,
	"ResourceSink_FoundationExpansionPack_C":           &FoundationExpansionPack,
	"ResourceSink_FrameworkFoundations_C":              &FrameworkFoundations,
	"ResourceSink_GasFilters_C":                        &GasFilters,
	"ResourceSink_GoldenCup_C":                         &GoldenCup,
	"Research_HardDrive_0_C":                           &HardDrive0,
	"ResourceSink_HealthPack_C":                        &HealthPack,
	"ResourceSink_HeatSink_C":                          &HeatSink,
	"ResourceSink_HeavyModularFrame_C":                 &HeavyModularFrame,
	"ResourceSink_HighSpeedConnector_C":                &HighSpeedConnector,
	"ResourceSink_HyperTubeWallAttachements_C":         &HyperTubeWallAttachements,
	"ResourceSink_InvertedCornerRamps_C":               &InvertedCornerRamps,
	"ResourceSink_InvertedRampPack_C":                  &InvertedRampPack,
	"ResourceSink_Ladders_C":                           &Ladders,
	"ResourceSink_LightControlPanel_C":                 &LightControlPanel,
	"ResourceSink_LightTower_C":                        &LightTower,
	"ResourceSink_ModularFrame_C":                      &ModularFrame,
	"ResourceSink_Motor_C":                             &Motor,
	"Research_Mycelia_1_C":                             &Mycelia1,
	"Research_Mycelia_2_C":                             &Mycelia2,
	"Research_Mycelia_3_C":                             &Mycelia3,
	"Research_Mycelia_4_C":                             &Mycelia4,
	"Research_Mycelia_5_C":                             &Mycelia5,
	"Research_Nutrients_0_C":                           &Nutrients0,
	"Research_Nutrients_1_C":                           &Nutrients1,
	"Research_Nutrients_2_C":                           &Nutrients2,
	"Research_Nutrients_3_C":                           &Nutrients3,
	"Research_Nutrients_4_C":                           &Nutrients4,
	"ResourceSink_PackagedBiofuel_C":                   &PackagedBiofuel,
	"ResourceSink_PackagedFuel_C":                      &PackagedFuel,
	"ResourceSink_Parachutes_C":                        &Parachutes,
	"ResourceSink_PetroleumCoke_C":                     &PetroleumCoke,
	"ResourceSink_PipelineWallAttachments_C":           &PipelineWallAttachments,
	"ResourceSink_Plastic_C":                           &Plastic,
	"ResourceSink_Plate_C":                             &Plate,
	"ResourceSink_PolymerResin_C":                      &PolymerResin,
	"Research_PowerSlugs_1_C":                          &PowerSlugs1,
	"Research_PowerSlugs_2_C":                          &PowerSlugs2,
	"Research_PowerSlugs_3_C":                          &PowerSlugs3,
	"Research_PowerSlugs_4_C":                          &PowerSlugs4,
	"Research_PowerSlugs_5_C":                          &PowerSlugs5,
	"Research_Quartz_0_C":                              &Quartz0,
	"Research_Quartz_1_1_C":                            &Quartz11,
	"Research_Quartz_1_2_C":                            &Quartz12,
	"Research_Quartz_2_C":                              &Quartz2,
	"Research_Quartz_3_C":                              &Quartz3,
	"Research_Quartz_3_1_C":                            &Quartz31,
	"Research_Quartz_3_2_C":                            &Quartz32,
	"Research_Quartz_3_3_C":                            &Quartz33,
	"Research_Quartz_4_C":                              &Quartz4,
	"Research_Quartz_4_1_C":                            &Quartz41,
	"ResourceSink_Quickwire_C":                         &Quickwire,
	"ResourceSink_RadiationFilters_C":                  &RadiationFilters,
	"ResourceSink_RadioControlUnit_C":                  &RadioControlUnit,
	"ResourceSink_ReinforcedIronPlate_C":               &ReinforcedIronPlate,
	"ResourceSink_Rod_C":                               &Rod,
	"ResourceSink_Rotor_C":                             &Rotor,
	"ResourceSink_Rubber_C":                            &Rubber,
	"Schematic_SaveCompatibility_C":                    &SaveCompatibility,
	"Schematic_1-1_C":                                  &Schematic1_1,
	"Schematic_1-2_C":                                  &Schematic1_2,
	"Schematic_1-3_C":                                  &Schematic1_3,
	"Schematic_2-1_C":                                  &Schematic2_1,
	"Schematic_2-2_C":                                  &Schematic2_2,
	"Schematic_2-3_C":                                  &Schematic2_3,
	"Schematic_2-5_C":                                  &Schematic2_5,
	"Schematic_3-1_C":                                  &Schematic3_1,
	"Schematic_3-2_C":                                  &Schematic3_2,
	"Schematic_3-3_C":                                  &Schematic3_3,
	"Schematic_3-4_C":                                  &Schematic3_4,
	"Schematic_4-1_C":                                  &Schematic4_1,
	"Schematic_4-2_C":                                  &Schematic4_2,
	"Schematic_4-4_C":                                  &Schematic4_4,
	"Schematic_5-1_C":                                  &Schematic5_1,
	"Schematic_5-1-1_C":                                &Schematic5_1_1,
	"Schematic_5-2_C":                                  &Schematic5_2,
	"Schematic_5-3_C":                                  &Schematic5_3,
	"Schematic_5-4_C":                                  &Schematic5_4,
	"Schematic_5-4-1_C":                                &Schematic5_4_1,
	"Schematic_6-1_C":                                  &Schematic6_1,
	"Schematic_6-2_C":                                  &Schematic6_2,
	"Schematic_6-3_C":                                  &Schematic6_3,
	"Schematic_6-4_C":                                  &Schematic6_4,
	"Schematic_6-5_C":                                  &Schematic6_5,
	"Schematic_7-1_C":                                  &Schematic7_1,
	"Schematic_7-1-1_C":                                &Schematic7_1_1,
	"Schematic_7-2_C":                                  &Schematic7_2,
	"Schematic_7-3_C":                                  &Schematic7_3,
	"Schematic_7-4_C":                                  &Schematic7_4,
	"Schematic_7-4-1_C":                                &Schematic7_4_1,
	"Schematic_8-1_C":                                  &Schematic8_1,
	"Schematic_8-2_C":                                  &Schematic8_2,
	"Schematic_8-2-1_C":                                &Schematic8_2_1,
	"Schematic_8-3_C":                                  &Schematic8_3,
	"Schematic_8-4_C":                                  &Schematic8_4,
	"Schematic_8-5_C":                                  &Schematic8_5,
	"Schematic_8-5-1_C":                                &Schematic8_5_1,
	"ResourceSink_Screw_C":                             &Screw,
	"ResourceSink_Silica_C":                            &Silica,
	"ResourceSink_Stairs_C":                            &Stairs,
	"Schematic_StartingRecipes_C":                      &StartingRecipes,
	"ResourceSink_Stator_C":                            &Stator,
	"ResourceSink_StatueBronzePioneer_C":               &StatueBronzePioneer,
	"ResourceSink_StatueGoldPioneer_C":                 &StatueGoldPioneer,
	"ResourceSink_StatueGoldenNut_C":                   &StatueGoldenNut,
	"ResourceSink_Statue_Hoggo_C":                      &StatueHoggo,
	"ResourceSink_StatueLizardDoggo_C":                 &StatueLizardDoggo,
	"ResourceSink_StatueSilverPioneer_C":               &StatueSilverPioneer,
	"ResourceSink_StatueSpaceGiraffe_C":                &StatueSpaceGiraffe,
	"ResourceSink_SteelBeam_C":                         &SteelBeam,
	"ResourceSink_SteelPipe_C":                         &SteelPipe,
	"ResourceSink_StreetLight_C":                       &StreetLight,
	"Research_Sulfur_0_C":                              &Sulfur0,
	"Research_Sulfur_1_C":                              &Sulfur1,
	"Research_Sulfur_2_C":                              &Sulfur2,
	"Research_Sulfur_3_C":                              &Sulfur3,
	"Research_Sulfur_3_1_C":                            &Sulfur31,
	"Research_Sulfur_3_2_C":                            &Sulfur32,
	"Research_Sulfur_3_2_1_C":                          &Sulfur321,
	"Research_Sulfur_4_C":                              &Sulfur4,
	"Research_Sulfur_4_1_C":                            &Sulfur41,
	"Research_Sulfur_4_2_C":                            &Sulfur42,
	"Research_Sulfur_4_2_1_C":                          &Sulfur421,
	"Research_Sulfur_5_C":                              &Sulfur5,
	"Research_Sulfur_6_C":                              &Sulfur6,
	"ResourceSink_SuperComputer_C":                     &SuperComputer,
	"ResourceSink_TurboMotor_C":                        &TurboMotor,
	"Schematic_Tutorial1_C":                            &Tutorial1,
	"Schematic_Tutorial1_5_C":                          &Tutorial15,
	"Schematic_Tutorial2_C":                            &Tutorial2,
	"Schematic_Tutorial3_C":                            &Tutorial3,
	"Schematic_Tutorial4_C":                            &Tutorial4,
	"Schematic_Tutorial5_C":                            &Tutorial5,
	"ResourceSink_Walkways_C":                          &Walkways,
	"ResourceSink_WallPowerPoles_C":                    &WallPowerPoles,
	"ResourceSink_WallPowerPolesMK2_C":                 &WallPowerPolesMK2,
	"ResourceSink_WallPowerPolesMK3_C":                 &WallPowerPolesMK3,
	"ResourceSink_WindowedWalls_C":                     &WindowedWalls,
	"ResourceSink_Wire_C":                              &Wire,
	"Research_XMas_1_C":                                &XMas1,
	"Research_XMas_1-1_C":                              &XMas1_1,
	"Research_XMas_1-2_C":                              &XMas1_2,
	"Research_XMas_2_C":                                &XMas2,
	"Research_XMas_2-1_C":                              &XMas2_1,
	"Research_XMas_2-2_C":                              &XMas2_2,
	"Research_XMas_3_C":                                &XMas3,
	"Research_XMas_3-1_C":                              &XMas3_1,
	"Research_XMas_3-2_C":                              &XMas3_2,
	"Research_XMas_4_C":                                &XMas4,
	"Research_XMas_4-1_C":                              &XMas4_1,
	"Research_XMas_4-2_C":                              &XMas4_2,
	"Research_XMas_5_C":                                &XMas5,
}