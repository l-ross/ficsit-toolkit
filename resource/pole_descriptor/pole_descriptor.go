// Code generated by ../../gen/docs_json. DO NOT EDIT.

package PoleDescriptor

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGPoleDescriptor struct {
	Name                    string
	ClassName               string
	MAbbreviatedDisplayName string
	MBuildMenuPriority      float64
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
	MSmallIcon              string
	MStackSize              resource.StackSize
	MSubCategories          string
}

var (
	ConveyorPole = FGPoleDescriptor{
		Name:                    "ConveyorPole",
		ClassName:               "Desc_ConveyorPole_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      30.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Factory/ConveyorPole/UI/IconDesc_ConveyorPole_512.IconDesc_ConveyorPole_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Factory/ConveyorPole/UI/IconDesc_ConveyorPole_256.IconDesc_ConveyorPole_256`,
		MStackSize:              resource.Medium,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_ConverPole.SC_ConverPole_C"')`,
	}

	ConveyorPoleStackable = FGPoleDescriptor{
		Name:                    "ConveyorPoleStackable",
		ClassName:               "Desc_ConveyorPoleStackable_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      30.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Factory/ConveyorPoleMulti/UI/ConveyorPoleMulti_512.ConveyorPoleMulti_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Factory/ConveyorPoleMulti/UI/ConveyorPoleMulti_256.ConveyorPoleMulti_256`,
		MStackSize:              resource.Medium,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_ConverPole.SC_ConverPole_C"')`,
	}

	HyperPoleStackable = FGPoleDescriptor{
		Name:                    "HyperPoleStackable",
		ClassName:               "Desc_HyperPoleStackable_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      10.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Factory/HyperTubeWallSupport/UI/IconDesc_HyperTubeStackable_512.IconDesc_HyperTubeStackable_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Factory/HyperTubeWallSupport/UI/IconDesc_HyperTubeStackable_256.IconDesc_HyperTubeStackable_256`,
		MStackSize:              resource.Medium,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_HyperTubes.SC_HyperTubes_C"')`,
	}

	PipeHyperSupport = FGPoleDescriptor{
		Name:                    "PipeHyperSupport",
		ClassName:               "Desc_PipeHyperSupport_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      3.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Factory/PipeHyperSupport/UI/HyperTubePole_512.HyperTubePole_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Factory/PipeHyperSupport/UI/HyperTubePole_256.HyperTubePole_256`,
		MStackSize:              resource.Medium,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_HyperTubes.SC_HyperTubes_C"')`,
	}

	PipeSupportStackable = FGPoleDescriptor{
		Name:                    "PipeSupportStackable",
		ClassName:               "Desc_PipeSupportStackable_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      1.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Factory/PipePole/UI/PipePole_Stackable_512.PipePole_Stackable_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Factory/PipePole/UI/PipePole_Stackable_256.PipePole_Stackable_256`,
		MStackSize:              resource.Medium,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_PipeSupport.SC_PipeSupport_C"')`,
	}

	PipelineSupport = FGPoleDescriptor{
		Name:                    "PipelineSupport",
		ClassName:               "Desc_PipelineSupport_C",
		MAbbreviatedDisplayName: ``,
		MBuildMenuPriority:      1.000000,
		MCanBeDiscarded:         true,
		MDescription:            ``,
		MDisplayName:            ``,
		MEnergyValue:            0.000000,
		MFluidColor:             `(B=0,G=0,R=0,A=0)`,
		MForm:                   resource.Invalid,
		MGasColor:               `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:      `Texture2D /Game/FactoryGame/Buildable/Factory/PipePole/UI/PipePole_512.PipePole_512`,
		MRadioactiveDecay:       0.000000,
		MRememberPickUp:         false,
		MSmallIcon:              `Texture2D /Game/FactoryGame/Buildable/Factory/PipePole/UI/PipePole_256.PipePole_256`,
		MStackSize:              resource.Medium,
		MSubCategories:          `(BlueprintGeneratedClass'"/Game/FactoryGame/Interface/UI/InGame/BuildMenu/BuildCategories/Sub_Transport/SC_PipeSupport.SC_PipeSupport_C"')`,
	}
)

func GetByClassName(className string) (FGPoleDescriptor, error) {
	if v, ok := ClassNameToDescriptor[className]; ok {
		return v, nil
	}

	return FGPoleDescriptor{}, fmt.Errorf("failed to find FGPoleDescriptor with class name %s", className)
}

var ClassNameToDescriptor = map[string]FGPoleDescriptor{
	"Desc_ConveyorPole_C":          ConveyorPole,
	"Desc_ConveyorPoleStackable_C": ConveyorPoleStackable,
	"Desc_HyperPoleStackable_C":    HyperPoleStackable,
	"Desc_PipeHyperSupport_C":      PipeHyperSupport,
	"Desc_PipeSupportStackable_C":  PipeSupportStackable,
	"Desc_PipelineSupport_C":       PipelineSupport,
}
