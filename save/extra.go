package save

import (
	"fmt"
)

type ExtraType string

const (
	CircuitSubsystemExtraType ExtraType = "CircuitSubsystem"
	ConveyorBeltExtraType     ExtraType = "ConveyorBelt"
	GameModeExtraType         ExtraType = "GameMode"
	GameStateExtraType        ExtraType = "GameState"
	PlayerStateExtraType      ExtraType = "PlayerState"
	PowerLineExtraType        ExtraType = "PowerLine"
	TrainExtraType            ExtraType = "Train"
	UnknownExtraType          ExtraType = "Unknown"
	VehicleExtraType          ExtraType = "Vehicle"
)

type Extra struct {
	Type  ExtraType
	Value ExtraValue
}

type ExtraValue interface {
	parse(p *parser) error

	serialize(s *serializer) error
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

//
// CircuitSubsystem
//

type CircuitSubsystemExtra struct {
	Circuits []*Circuit
}

type Circuit struct {
	ID        int32
	LevelName string
	PathName  string
}

func newCircuitSubsystem(_ int32) *Extra {
	return &Extra{
		Type:  CircuitSubsystemExtraType,
		Value: &CircuitSubsystemExtra{},
	}
}

func (e *Extra) GetCircuitSubsystem() (*CircuitSubsystemExtra, error) {
	if v, ok := e.Value.(*CircuitSubsystemExtra); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *CircuitSubsystemExtra) parse(p *parser) error {
	count, err := p.readInt32()
	if err != nil {
		return err
	}

	e.Circuits = make([]*Circuit, count)

	for i := int32(0); i < count; i++ {
		c := &Circuit{}

		c.ID, err = p.readInt32()
		if err != nil {
			return err
		}

		c.LevelName, err = p.readString()
		if err != nil {
			return err
		}

		c.PathName, err = p.readString()
		if err != nil {
			return err
		}

		e.Circuits[i] = c
	}

	return nil
}

func (e *CircuitSubsystemExtra) serialize(s *serializer) error {
	err := s.writeInt32(int32(len(e.Circuits)))
	if err != nil {
		return err
	}

	for _, c := range e.Circuits {
		err = s.writeInt32(c.ID)
		if err != nil {
			return err
		}

		err = s.writeString(c.LevelName)
		if err != nil {
			return err
		}

		err = s.writeString(c.PathName)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// ConveyorBelt
//

type ConveyorBeltExtra struct {
	Items []*ConveyorBeltItem
}

type ConveyorBeltItem struct {
	ResourceName string
	LevelName    string
	PathName     string
	Position     float32
}

func newConveyorBelt(_ int32) *Extra {
	return &Extra{
		Type:  ConveyorBeltExtraType,
		Value: &ConveyorBeltExtra{},
	}
}

func (e *Extra) GetConveyorBelt() (*ConveyorBeltExtra, error) {
	if v, ok := e.Value.(*ConveyorBeltExtra); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *ConveyorBeltExtra) parse(p *parser) error {
	itemCount, err := p.readInt32()
	if err != nil {
		return err
	}

	e.Items = make([]*ConveyorBeltItem, itemCount)

	for i := int32(0); i < itemCount; i++ {
		item := &ConveyorBeltItem{}

		// UNKNOWN_DATA
		_, err = p.readInt32()
		if err != nil {
			return err
		}

		item.ResourceName, err = p.readString()
		if err != nil {
			return err
		}

		item.LevelName, err = p.readString()
		if err != nil {
			return err
		}

		item.PathName, err = p.readString()
		if err != nil {
			return err
		}

		item.Position, err = p.readFloat32()
		if err != nil {
			return err
		}

		e.Items[i] = item
	}

	return nil
}

func (e *ConveyorBeltExtra) serialize(s *serializer) error {
	err := s.writeInt32(int32(len(e.Items)))
	if err != nil {
		return err
	}

	for _, i := range e.Items {
		// UNKNOWN_DATA
		err = s.writeInt32(0)
		if err != nil {
			return err
		}

		err = s.writeString(i.ResourceName)
		if err != nil {
			return err
		}

		err = s.writeString(i.LevelName)
		if err != nil {
			return err
		}

		err = s.writeString(i.PathName)
		if err != nil {
			return err
		}

		err = s.writeFloat32(i.Position)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// GameMode
//

type GameModeExtra struct {
	Objects []*ObjectReference
}

func newGameMode(_ int32) *Extra {
	return &Extra{
		Type:  GameModeExtraType,
		Value: &GameModeExtra{},
	}
}

func (e *Extra) GetGameModeExtra() (*GameModeExtra, error) {
	if v, ok := e.Value.(*GameModeExtra); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *GameModeExtra) parse(p *parser) error {
	count, err := p.readInt32()
	if err != nil {
		return err
	}

	for i := int32(0); i < count; i++ {
		o := &ObjectReference{}

		o.LevelName, err = p.readString()
		if err != nil {
			return err
		}

		o.PathName, err = p.readString()
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *GameModeExtra) serialize(s *serializer) error {
	err := s.writeInt32(int32(len(e.Objects)))
	if err != nil {
		return err
	}

	for _, o := range e.Objects {
		err = s.writeString(o.LevelName)
		if err != nil {
			return err
		}

		err = s.writeString(o.PathName)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// GameState
//

type GameStateExtra struct {
	Objects []*ObjectReference
}

func newGameState(_ int32) *Extra {
	return &Extra{
		Type:  GameStateExtraType,
		Value: &GameStateExtra{},
	}
}

func (e *Extra) GetGameStateExtra() (*GameStateExtra, error) {
	if v, ok := e.Value.(*GameStateExtra); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *GameStateExtra) parse(p *parser) error {
	count, err := p.readInt32()
	if err != nil {
		return err
	}

	for i := int32(0); i < count; i++ {
		o := &ObjectReference{}

		o.LevelName, err = p.readString()
		if err != nil {
			return err
		}

		o.PathName, err = p.readString()
		if err != nil {
			return err
		}

		e.Objects = append(e.Objects, o)
	}

	return nil
}

func (e *GameStateExtra) serialize(s *serializer) error {
	err := s.writeInt32(int32(len(e.Objects)))
	if err != nil {
		return err
	}

	for _, o := range e.Objects {
		err = s.writeString(o.LevelName)
		if err != nil {
			return err
		}

		err = s.writeString(o.PathName)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// PlayerState
//

type PlayerStateExtra struct {
	Data []byte

	len int32
}

func newPlayerState(l int32) *Extra {
	return &Extra{
		Type: PlayerStateExtraType,
		Value: &PlayerStateExtra{
			len: l,
		},
	}
}

func (e *Extra) GetPlayerStateExtra() (*PlayerStateExtra, error) {
	if v, ok := e.Value.(*PlayerStateExtra); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *PlayerStateExtra) parse(p *parser) error {
	var err error
	e.Data, err = p.readBytes(e.len)
	if err != nil {
		return err
	}

	return nil
}

func (e *PlayerStateExtra) serialize(s *serializer) error {
	return s.writeBytes(e.Data)
}

//
// PowerLine
//

type PowerLineExtra struct {
	SourceLevelName string
	SourcePathName  string
	TargetLevelName string
	TargetPathName  string
}

func newPowerLine(_ int32) *Extra {
	return &Extra{
		Type:  PowerLineExtraType,
		Value: &PowerLineExtra{},
	}
}

func (e *Extra) GetPowerLineExtra() (*PowerLineExtra, error) {
	if v, ok := e.Value.(*PowerLineExtra); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *PowerLineExtra) parse(p *parser) error {
	var err error
	e.SourceLevelName, err = p.readString()
	if err != nil {
		return err
	}

	e.SourcePathName, err = p.readString()
	if err != nil {
		return err
	}

	e.TargetLevelName, err = p.readString()
	if err != nil {
		return err
	}

	e.TargetPathName, err = p.readString()
	if err != nil {
		return err
	}

	return nil
}

func (e *PowerLineExtra) serialize(s *serializer) error {
	err := s.writeString(e.SourceLevelName)
	if err != nil {
		return err
	}

	err = s.writeString(e.SourcePathName)
	if err != nil {
		return err
	}

	err = s.writeString(e.TargetLevelName)
	if err != nil {
		return err
	}

	err = s.writeString(e.TargetPathName)
	if err != nil {
		return err
	}

	return nil
}

//
// Train
//

type TrainExtra struct {
	PreviousLevelName string
	PreviousPathName  string
	NextLevelName     string
	NextPathName      string
}

func newTrainExtra(_ int32) *Extra {
	return &Extra{
		Type:  TrainExtraType,
		Value: &TrainExtra{},
	}
}

func (e *Extra) GetTrainExtra() (*TrainExtra, error) {
	if v, ok := e.Value.(*TrainExtra); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *TrainExtra) parse(p *parser) error {
	// UNKNOWN_DATA
	_, err := p.readInt32()
	if err != nil {
		return err
	}

	e.PreviousLevelName, err = p.readString()
	if err != nil {
		return err
	}

	e.PreviousPathName, err = p.readString()
	if err != nil {
		return err
	}

	e.NextLevelName, err = p.readString()
	if err != nil {
		return err
	}

	e.NextPathName, err = p.readString()
	if err != nil {
		return err
	}

	return nil
}

func (e *TrainExtra) serialize(s *serializer) error {
	err := s.writeInt32(0)
	if err != nil {
		return err
	}

	err = s.writeString(e.PreviousLevelName)
	if err != nil {
		return err
	}

	err = s.writeString(e.PreviousPathName)
	if err != nil {
		return err
	}

	err = s.writeString(e.NextLevelName)
	if err != nil {
		return err
	}

	err = s.writeString(e.NextPathName)
	if err != nil {
		return err
	}

	return nil
}

//
// Unknown
//

type UnknownExtra struct {
	Data []byte

	len int32
}

func newUnknownExtra(l int32) *Extra {
	return &Extra{
		Type: UnknownExtraType,
		Value: &UnknownExtra{
			len: l,
		},
	}
}

func (e *Extra) GetUnknownExtra() (*UnknownExtra, error) {
	if v, ok := e.Value.(*UnknownExtra); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *UnknownExtra) parse(p *parser) error {
	var err error
	e.Data, err = p.readBytes(e.len)
	if err != nil {
		return err
	}

	return nil
}

func (e *UnknownExtra) serialize(s *serializer) error {
	return s.writeBytes(e.Data)
}

//
// Vehicle
//

type Vehicle struct {
	Data []VehicleData
}

type VehicleData struct {
	Name string
	Data []byte
}

func newVehicle(_ int32) *Extra {
	return &Extra{
		Type:  VehicleExtraType,
		Value: &Vehicle{},
	}
}

func (e *Extra) GetVehicle() (*Vehicle, error) {
	if v, ok := e.Value.(*Vehicle); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *Vehicle) parse(p *parser) error {
	objCount, err := p.readInt32()
	if err != nil {
		return err
	}

	for x := int32(0); x < objCount; x++ {
		name, err := p.readString()
		if err != nil {
			return err
		}

		data, err := p.readBytes(53)
		if err != nil {
			return err
		}

		e.Data = append(e.Data, VehicleData{
			Name: name,
			Data: data,
		})
	}

	return nil
}

func (e *Vehicle) serialize(s *serializer) error {
	err := s.writeInt32(int32(len(e.Data)))
	if err != nil {
		return err
	}

	for _, d := range e.Data {
		err := s.writeString(d.Name)
		if err != nil {
			return err
		}

		err = s.writeBytes(d.Data)
		if err != nil {
			return err
		}
	}

	return nil
}
