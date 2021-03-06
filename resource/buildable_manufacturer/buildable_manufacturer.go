// Code generated by ../../gen/docs_json. DO NOT EDIT.

package BuildableManufacturer

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableManufacturer struct {
	Name                                 string
	ClassName                            string
	BIsPendingToKillVfx                  bool
	CurrentPackagingMode                 string
	CurrentPotentialConvert              string
	IsPowered                            bool
	MAddToSignificanceManager            bool
	MAllowColoring                       bool
	MAttachmentPoints                    string
	MBuildEffectSpeed                    float64
	MCachedSkeletalMeshes                string
	MCanChangePotential                  bool
	MColor                               string
	MCreateClearanceMeshRepresentation   bool
	MCurrentColorVFX                     string
	MCurrentColor_VFX                    string
	MCurrentRecipeChanged                string
	MCurrentRecipeCheck                  string
	MDescription                         string
	MDisplayName                         string
	MDoesHaveShutdownAnimation           bool
	MEffectUpdateInterval                float64
	MFactoryInputConnections             string
	MFactoryOutputConnections            string
	MFluidStackSizeDefault               resource.StackSize
	MFluidStackSizeMultiplier            int
	MForceNetUpdateOnRegisterPlayer      bool
	MHideOnBuildEffectStart              bool
	MHighlightVector                     string
	MInteractingPlayers                  string
	MIsPendingToKillVFX                  bool
	MIsRadioActive                       bool
	MIsUseable                           bool
	MManufacturingSpeed                  float64
	MMaxPotential                        float64
	MMaxPotentialIncreasePerCrystal      float64
	MMinPotential                        float64
	MMinimumProducingTime                float64
	MMinimumStoppedTime                  float64
	MNumCyclesForProductivity            int
	MOnHasPowerChanged                   string
	MOnHasProductionChanged              string
	MOnHasStandbyChanged                 string
	MPipeInputConnections                string
	MPipeOutputConnections               string
	MPowerConsumption                    float64
	MPowerConsumptionExponent            float64
	MPreviousRecipeCheck                 string
	MProductionEffectsRunning            bool
	MShouldModifyWorldGrid               bool
	MShouldShowAttachmentPointVisuals    bool
	MShouldShowHighlight                 bool
	MSignificanceRange                   float64
	MSkipBuildEffect                     bool
	MSocketStoppedAkComponents           string
	MStoppedAkComponents                 string
	MStoppedProducingAnimationSounds     bool
	MToggleDormancyOnInteraction         bool
	M_NotifyNameREferences               string
	MaxRenderDistance                    float64
	OnReplicationDetailActorCreatedEvent string
}

var (
	AssemblerMk1 = FGBuildableManufacturer{
		Name:                               "AssemblerMk1",
		ClassName:                          "Build_AssemblerMk1_C",
		IsPowered:                          false,
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                true,
		MCreateClearanceMeshRepresentation: true,
		MCurrentRecipeChanged:              `()`,
		MDescription: `Crafts two parts into another part.

Can be automated by feeding parts into it with a conveyor belt connected to the input. The produced parts can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Assembler`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    15.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	Blender = FGBuildableManufacturer{
		Name:                               "Blender",
		ClassName:                          "Build_Blender_C",
		IsPowered:                          false,
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                true,
		MColor:                             ``,
		MCreateClearanceMeshRepresentation: true,
		MCurrentColorVFX:                   `(R=0.000000,G=0.000000,B=0.000000,A=0.000000)`,
		MCurrentRecipeChanged:              `()`,
		MDescription: `The Blender is capable of blending fluids and combining them with solid parts in various processes.
Head Lift: 10 meters.
(Allows fluids to be transported 10 meters upwards).

Contains both Conveyor Belt and Pipe inputs and outputs.`,
		MDisplayName:                         `Blender`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsRadioActive:                       false,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    75.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		M_NotifyNameREferences:               `("Arm_04_ClawBase","Arm_02_SFXSocket")`,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	ConstructorMk1 = FGBuildableManufacturer{
		Name:                               "ConstructorMk1",
		ClassName:                          "Build_ConstructorMk1_C",
		CurrentPotentialConvert:            `((1, 1.000000),(2, 1.200000),(0, 0.650000))`,
		IsPowered:                          false,
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                true,
		MCreateClearanceMeshRepresentation: true,
		MCurrentRecipeChanged:              `()`,
		MCurrentRecipeCheck:                ``,
		MDescription: `Crafts one part into another part.

Can be automated by feeding parts into it with a conveyor belt connected to the input. The produced parts can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Constructor`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                0.000000,
		MMinimumStoppedTime:                  0.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    4.000000,
		MPowerConsumptionExponent:            1.600000,
		MPreviousRecipeCheck:                 ``,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   8000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	FoundryMk1 = FGBuildableManufacturer{
		Name:                               "FoundryMk1",
		ClassName:                          "Build_FoundryMk1_C",
		BIsPendingToKillVfx:                false,
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                true,
		MCreateClearanceMeshRepresentation: true,
		MCurrentRecipeChanged:              `()`,
		MDescription: `Smelts two resources into alloy ingots.

Can be automated by feeding ore into it with a conveyor belt connected to the inputs. The produced ingots can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Foundry`,
		MDoesHaveShutdownAnimation:           true,
		MEffectUpdateInterval:                0.000000,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    16.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   17000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	ManufacturerMk1 = FGBuildableManufacturer{
		Name:                               "ManufacturerMk1",
		ClassName:                          "Build_ManufacturerMk1_C",
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                true,
		MCreateClearanceMeshRepresentation: true,
		MCurrentRecipeChanged:              `()`,
		MDescription: `Crafts three or four parts into another part.

Can be automated by feeding parts into it with a conveyor belt connected to the input. The produced parts can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Manufacturer`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    55.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MSocketStoppedAkComponents:           ``,
		MStoppedAkComponents:                 ``,
		MStoppedProducingAnimationSounds:     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	OilRefinery = FGBuildableManufacturer{
		Name:                               "OilRefinery",
		ClassName:                          "Build_OilRefinery_C",
		IsPowered:                          false,
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                true,
		MCreateClearanceMeshRepresentation: true,
		MCurrentRecipeChanged:              `()`,
		MDescription: `Refines fluid and/or solid parts into other parts.
Head Lift: 10 meters.
(Allows fluids to be transported 10 meters upwards.)

Contains both a Conveyor Belt and Pipe input and output, to allow for the automation of various recipes.`,
		MDisplayName:                         `Refinery`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    30.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   17000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	Packager = FGBuildableManufacturer{
		Name:                               "Packager",
		ClassName:                          "Build_Packager_C",
		CurrentPackagingMode:               ``,
		IsPowered:                          false,
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                true,
		MCreateClearanceMeshRepresentation: true,
		MCurrentColor_VFX:                  `(R=1.000000,G=1.000000,B=1.000000,A=0.000000)`,
		MCurrentRecipeChanged:              `()`,
		MDescription: `Used for the packaging and unpacking of fluids.
Head Lift: 10 meters.
(Allows fluids to be transported 10 meters upwards.)

Contains both a Conveyor Belt and Pipe input and output, to allow for the automation of various recipes.`,
		MDisplayName:                         `Packager`,
		MDoesHaveShutdownAnimation:           false,
		MEffectUpdateInterval:                0.000000,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    10.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}

	SmelterMk1 = FGBuildableManufacturer{
		Name:                               "SmelterMk1",
		ClassName:                          "Build_SmelterMk1_C",
		MAddToSignificanceManager:          true,
		MAllowColoring:                     true,
		MAttachmentPoints:                  ``,
		MBuildEffectSpeed:                  0.000000,
		MCachedSkeletalMeshes:              ``,
		MCanChangePotential:                true,
		MCreateClearanceMeshRepresentation: true,
		MCurrentRecipeChanged:              `()`,
		MDescription: `Smelts ore into ingots.

Can be automated by feeding ore into it with a conveyor belt connected to the input. The produced ingots can be automatically extracted by connecting a conveyor belt to the output.`,
		MDisplayName:                         `Smelter`,
		MDoesHaveShutdownAnimation:           true,
		MEffectUpdateInterval:                0.000000,
		MFactoryInputConnections:             ``,
		MFactoryOutputConnections:            ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsPendingToKillVFX:                  false,
		MIsUseable:                           true,
		MManufacturingSpeed:                  1.000000,
		MMaxPotential:                        1.000000,
		MMaxPotentialIncreasePerCrystal:      0.500000,
		MMinPotential:                        0.010000,
		MMinimumProducingTime:                2.000000,
		MMinimumStoppedTime:                  5.000000,
		MNumCyclesForProductivity:            20,
		MOnHasPowerChanged:                   `()`,
		MOnHasProductionChanged:              `()`,
		MOnHasStandbyChanged:                 `()`,
		MPipeInputConnections:                ``,
		MPipeOutputConnections:               ``,
		MPowerConsumption:                    4.000000,
		MPowerConsumptionExponent:            1.600000,
		MProductionEffectsRunning:            false,
		MShouldModifyWorldGrid:               true,
		MShouldShowAttachmentPointVisuals:    false,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   15000.000000,
		MSkipBuildEffect:                     false,
		MToggleDormancyOnInteraction:         false,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (FGBuildableManufacturer, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGBuildableManufacturer{}, fmt.Errorf("failed to find FGBuildableManufacturer with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGBuildableManufacturer{
	"Build_AssemblerMk1_C":    AssemblerMk1,
	"Build_Blender_C":         Blender,
	"Build_ConstructorMk1_C":  ConstructorMk1,
	"Build_FoundryMk1_C":      FoundryMk1,
	"Build_ManufacturerMk1_C": ManufacturerMk1,
	"Build_OilRefinery_C":     OilRefinery,
	"Build_Packager_C":        Packager,
	"Build_SmelterMk1_C":      SmelterMk1,
}
