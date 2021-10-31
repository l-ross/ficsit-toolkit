package BuildableDroneStation

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/resource"
)

type FGBuildableDroneStation struct {
	ClassName                               string
	MAddToSignificanceManager               bool
	MAllowCleranceSeparationEvenIfStackedOn bool
	MBatteryClasses                         string
	MBatteryStorageSizeX                    int
	MBatteryStorageSizeY                    int
	MBlockSharingMaterialInstanceMapping    bool
	MBuildEffectSpeed                       float64
	MCachedSkeletalMeshes                   string
	MCanChangePotential                     bool
	MDescription                            string
	MDisplayName                            string
	MDroneDockingLocationLocal              string
	MDroneDockingQueue                      string
	MDroneDockingStartLocationLocal         string
	MDroneQueueRadius                       float64
	MDroneQueueSeparationRadius             float64
	MDroneQueueVerticalSeparation           float64
	MEffectUpdateInterval                   float64
	MExcludeFromMaterialInstancing          string
	MFluidStackSizeDefault                  resource.StackSize
	MFluidStackSizeMultiplier               int
	MFogPlaneTransforms                     string
	MForceNetUpdateOnRegisterPlayer         bool
	MHideOnBuildEffectStart                 bool
	MHighlightVector                        string
	MInteractingPlayers                     string
	MIsUseable                              bool
	MItemTransferringStage                  string
	MMapText                                string
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
	MShouldModifyWorldGrid                  bool
	MShouldShowHighlight                    bool
	MSignificanceRange                      float64
	MSkipBuildEffect                        bool
	MSocketStoppedAkComponents              string
	MStackTransferSize                      float64
	MStationHasDronesInQueue                bool
	MStoppedAkComponents                    string
	MStoppedProducingAnimationSounds        bool
	MStorageSizeX                           int
	MStorageSizeY                           int
	MToggleDormancyOnInteraction            bool
	MTransferProgress                       float64
	MTransferSpeed                          float64
	MTripInformationSampleCount             int
	MTripPowerCost                          float64
	MTripPowerPerMeterCost                  float64
	M_DockingStates                         string
	M_OffsetTime                            float64
	MaxRenderDistance                       float64
	OnReplicationDetailActorCreatedEvent    string
}

var (
	DroneStation = FGBuildableDroneStation{
		ClassName:                               "Build_DroneStation_C",
		MAddToSignificanceManager:               true,
		MAllowCleranceSeparationEvenIfStackedOn: false,
		MBatteryClasses:                         `(BlueprintGeneratedClass'"/Game/FactoryGame/Resource/Parts/Battery/Desc_Battery.Desc_Battery_C"')`,
		MBatteryStorageSizeX:                    1,
		MBatteryStorageSizeY:                    1,
		MBlockSharingMaterialInstanceMapping:    false,
		MBuildEffectSpeed:                       0.000000,
		MCachedSkeletalMeshes:                   ``,
		MCanChangePotential:                     true,
		MDescription: `Drone Ports can have one other Port assigned as their transport destination.
Each Drone Port can contain a single Drone, which transports available input back and forth between its home and destination Ports.

The Drone Port interface provides delivery details and allows management of Port connections.`,
		MDisplayName:                         `Drone Port`,
		MDroneDockingLocationLocal:           `(X=0.000000,Y=0.000000,Z=570.000000)`,
		MDroneDockingQueue:                   ``,
		MDroneDockingStartLocationLocal:      `(X=-800.000000,Y=0.000000,Z=860.000000)`,
		MDroneQueueRadius:                    2000.000000,
		MDroneQueueSeparationRadius:          0.000000,
		MDroneQueueVerticalSeparation:        800.000000,
		MEffectUpdateInterval:                0.000000,
		MExcludeFromMaterialInstancing:       ``,
		MFluidStackSizeDefault:               resource.Fluid,
		MFluidStackSizeMultiplier:            1,
		MFogPlaneTransforms:                  `((Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000005),Translation=(X=-200.000000,Y=-900.000000,Z=140.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000005),Translation=(X=600.000000,Y=-900.000000,Z=140.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000005),Translation=(X=200.000000,Y=-900.000000,Z=140.000000)),(Rotation=(X=-0.000000,Y=0.000000,Z=-1.000000,W=0.000005),Translation=(X=-600.000000,Y=-900.000000,Z=140.000000)),(Rotation=(X=0.000000,Y=-0.000000,Z=0.000010,W=1.000000),Translation=(X=0.000007,Y=840.000000,Z=140.000000)))`,
		MForceNetUpdateOnRegisterPlayer:      false,
		MHideOnBuildEffectStart:              false,
		MHighlightVector:                     `(X=0.000000,Y=0.000000,Z=0.000000)`,
		MInteractingPlayers:                  ``,
		MIsUseable:                           true,
		MItemTransferringStage:               `ITS_NONE`,
		MMapText:                             `Drone Port`,
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
		MPowerConsumption:                    100.000000,
		MPowerConsumptionExponent:            1.600000,
		MShouldModifyWorldGrid:               true,
		MShouldShowHighlight:                 false,
		MSignificanceRange:                   18000.000000,
		MSkipBuildEffect:                     false,
		MSocketStoppedAkComponents:           ``,
		MStackTransferSize:                   1.000000,
		MStationHasDronesInQueue:             false,
		MStoppedAkComponents:                 ``,
		MStoppedProducingAnimationSounds:     false,
		MStorageSizeX:                        3,
		MStorageSizeY:                        6,
		MToggleDormancyOnInteraction:         false,
		MTransferProgress:                    0.000000,
		MTransferSpeed:                       0.000000,
		MTripInformationSampleCount:          1,
		MTripPowerCost:                       24000.000000,
		MTripPowerPerMeterCost:               6.000000,
		M_DockingStates:                      `DS_UNDOCKED`,
		M_OffsetTime:                         0.000000,
		MaxRenderDistance:                    -1.000000,
		OnReplicationDetailActorCreatedEvent: `()`,
	}
)

func GetByClassName(className string) (*FGBuildableDroneStation, error) {
	if v, ok := classNameToVar[className]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("failed to find FGBuildableDroneStation with name %s", className)
}

var classNameToVar = map[string]*FGBuildableDroneStation{
	"Build_DroneStation_C": &DroneStation,
}