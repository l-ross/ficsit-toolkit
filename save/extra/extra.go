// Package extra provides handling for the extra data that may be appended on to a save.Entity
//
// The type of the Extra depends on the TypePath of the save.Entity.
package extra

import (
	"github.com/l-ross/ficsit-toolkit/save/data"
)

type Extra struct {
	Type Type

	// The Value of the Extra.
	//
	// Accessing the value can be achieved by calling the appropriate Get method
	// on the Extra based on its Type.
	Value Value
}

type Value interface {
	parse(d *data.Data) error

	serialize(d *data.Data) error
}

type Type string

// Parse an Extra
func Parse(typePath string, l int32, d *data.Data) (*Extra, error) {
	e := getExtra(typePath)(l)

	err := e.Value.parse(d)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// Serialize an Extra
func (e *Extra) Serialize(d *data.Data) error {
	return e.Value.serialize(d)
}

func getExtra(c string) func(l int32) *Extra {
	switch c {
	case "/Game/FactoryGame/-Shared/Blueprint/BP_CircuitSubsystem.BP_CircuitSubsystem_C":
		return newCircuitSubsystem
	case "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk1/Build_ConveyorBeltMk1.Build_ConveyorBeltMk1_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk2/Build_ConveyorBeltMk2.Build_ConveyorBeltMk2_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk3/Build_ConveyorBeltMk3.Build_ConveyorBeltMk3_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk4/Build_ConveyorBeltMk4.Build_ConveyorBeltMk4_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk5/Build_ConveyorBeltMk5.Build_ConveyorBeltMk5_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk1/Build_ConveyorLiftMk1.Build_ConveyorLiftMk1_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk2/Build_ConveyorLiftMk2.Build_ConveyorLiftMk2_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk3/Build_ConveyorLiftMk3.Build_ConveyorLiftMk3_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk4/Build_ConveyorLiftMk4.Build_ConveyorLiftMk4_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk5/Build_ConveyorLiftMk5.Build_ConveyorLiftMk5_C":
		return newConveyorBelt
	case "/Game/FactoryGame/-Shared/Blueprint/BP_GameMode.BP_GameMode_C":
		return newGameMode
	case "/Game/FactoryGame/-Shared/Blueprint/BP_GameState.BP_GameState_C":
		return newGameState
	case "/Game/FactoryGame/Character/Player/BP_PlayerState.BP_PlayerState_C":
		return newPlayerState
	case "/Game/FactoryGame/Buildable/Factory/PowerLine/Build_PowerLine.Build_PowerLine_C",
		"/Game/FactoryGame/Events/Christmas/Buildings/PowerLineLights/Build_XmassLightsLine.Build_XmassLightsLine_C":
		return newPowerLine
	case "/Game/FactoryGame/Buildable/Vehicle/Train/Wagon/BP_FreightWagon.BP_FreightWagon_C",
		"/Game/FactoryGame/Buildable/Vehicle/Train/Locomotive/BP_Locomotive.BP_Locomotive_C":
		return newTrainExtra
	case "/Game/FactoryGame/Buildable/Vehicle/Tractor/BP_Tractor.BP_Tractor_C",
		"/Game/FactoryGame/Buildable/Vehicle/Truck/BP_Truck.BP_Truck_C",
		"/Game/FactoryGame/Buildable/Vehicle/Explorer/BP_Explorer.BP_Explorer_C",
		"/Game/FactoryGame/Buildable/Vehicle/Cyberwagon/Testa_BP_WB.Testa_BP_WB_C",
		"/Game/FactoryGame/Buildable/Vehicle/Golfcart/BP_Golfcart.BP_Golfcart_C",
		"/Game/FactoryGame/Buildable/Factory/DroneStation/BP_DroneTransport.BP_DroneTransport_C":
		return newVehicle
	}
	return newUnknownExtra
}

type Reference struct {
	LevelName string `json:"levelName"`
	PathName  string `json:"pathName"`
}
