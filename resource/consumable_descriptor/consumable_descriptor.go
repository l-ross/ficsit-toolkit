package ConsumableDescriptor

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGConsumableDescriptor struct {
	ClassName               string
	MAbbreviatedDisplayName string
	MCanBeDiscarded         bool
	MCustomHandsMeshScale   float64
	MCustomLocation         string
	MCustomRotation         string
	MDescription            string
	MDisplayName            string
	MEnergyValue            float64
	MFluidColor             string
	MForm                   resource.Form
	MGasColor               string
	MHealthGain             float64
	MPersistentBigIcon      string
	MRadioactiveDecay       float64
	MRememberPickUp         bool
	MResourceSinkPoints     int
	MSmallIcon              string
	MStackSize              resource.StackSize
}

var (
	Berry = FGConsumableDescriptor{
		ClassName:               "Desc_Berry_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCustomHandsMeshScale:   1.000000,
		MCustomLocation:         `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MCustomRotation:         `(Pitch=0.000000,Yaw=-90.000000,Roll=0.000000)`,
		MDescription: `Slot: Hands
Consumable

Can be eaten to restore one health segment.`,
		MDisplayName:        `Paleberry`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MHealthGain:         10.000000,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Environment/Berry/UI/Berry_256.Berry_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Environment/Berry/UI/Berry_64.Berry_64`,
		MStackSize:          resource.Small,
	}

	EquipmentDescriptorBeacon = FGConsumableDescriptor{
		ClassName:               "BP_EquipmentDescriptorBeacon_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCustomHandsMeshScale:   1.000000,
		MCustomLocation:         `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MCustomRotation:         `(Pitch=0.000000,Yaw=0.000000,Roll=0.000000)`,
		MDescription: `Slot: Hands
Consumable

Used to mark areas of interest. Displayed on your compass with the color and name you set for it.`,
		MDisplayName:        `Beacon`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/Beacon/UI/IconDesc_Beacon_256.IconDesc_Beacon_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 320,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/Beacon/UI/IconDesc_Beacon_64.IconDesc_Beacon_64`,
		MStackSize:          resource.Medium,
	}

	Medkit = FGConsumableDescriptor{
		ClassName:               "Desc_Medkit_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCustomHandsMeshScale:   0.200000,
		MCustomLocation:         `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MCustomRotation:         `(Pitch=0.000000,Yaw=0.000000,Roll=0.000000)`,
		MDescription: `Slot: Hands
Consumable

Can be inhaled to fully restore health.`,
		MDisplayName:        `Medicinal Inhaler`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MHealthGain:         100.000000,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/Medkit/UI/Inhaler_256.Inhaler_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 67,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/Medkit/UI/Inhaler_64.Inhaler_64`,
		MStackSize:          resource.Small,
	}

	Nut = FGConsumableDescriptor{
		ClassName:               "Desc_Nut_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCustomHandsMeshScale:   1.000000,
		MCustomLocation:         `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MCustomRotation:         `(Pitch=0.000000,Yaw=-90.000000,Roll=0.000000)`,
		MDescription: `Slot: Hands
Consumable

Can be eaten to restore half a health segment.`,
		MDisplayName:        `Beryl Nut`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MHealthGain:         5.000000,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Environment/Nut/UI/Nut_256_New.Nut_256_New`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Environment/Nut/UI/Nut_64_new.Nut_64_new`,
		MStackSize:          resource.Medium,
	}

	Parachute = FGConsumableDescriptor{
		ClassName:               "Desc_Parachute_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCustomHandsMeshScale:   1.000000,
		MCustomLocation:         `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MCustomRotation:         `(Pitch=0.000000,Yaw=0.000000,Roll=0.000000)`,
		MDescription: `Slot: Body
Consumable

Slows down your fall when activated in mid-air.`,
		MDisplayName:        `Parachute`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Equipment/Beacon/UI/Parachute_256.Parachute_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     false,
		MResourceSinkPoints: 608,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Equipment/Beacon/UI/Parachute_64.Parachute_64`,
		MStackSize:          resource.Small,
	}

	Shroom = FGConsumableDescriptor{
		ClassName:               "Desc_Shroom_C",
		MAbbreviatedDisplayName: ``,
		MCanBeDiscarded:         true,
		MCustomHandsMeshScale:   1.000000,
		MCustomLocation:         `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MCustomRotation:         `(Pitch=0.000000,Yaw=-90.000000,Roll=0.000000)`,
		MDescription: `Slot: Hands
Consumable

Can be eaten to restore two health segments.`,
		MDisplayName:        `Bacon Agaric`,
		MEnergyValue:        0.000000,
		MFluidColor:         `(B=0,G=0,R=0,A=0)`,
		MForm:               resource.Solid,
		MGasColor:           `(B=0,G=0,R=0,A=0)`,
		MHealthGain:         20.000000,
		MPersistentBigIcon:  `Texture2D /Game/FactoryGame/Resource/Environment/DesertShroom/UI/Mushroom_256.Mushroom_256`,
		MRadioactiveDecay:   0.000000,
		MRememberPickUp:     true,
		MResourceSinkPoints: 0,
		MSmallIcon:          `Texture2D /Game/FactoryGame/Resource/Environment/DesertShroom/UI/Mushroom_64.Mushroom_64`,
		MStackSize:          resource.Small,
	}
)

func GetByClassName(className string) (*FGConsumableDescriptor, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGConsumableDescriptor with name %s", className)
}

var classNameToVar = map[string]*FGConsumableDescriptor{
	"Desc_Berry_C":                   &Berry,
	"BP_EquipmentDescriptorBeacon_C": &EquipmentDescriptorBeacon,
	"Desc_Medkit_C":                  &Medkit,
	"Desc_Nut_C":                     &Nut,
	"Desc_Parachute_C":               &Parachute,
	"Desc_Shroom_C":                  &Shroom,
}