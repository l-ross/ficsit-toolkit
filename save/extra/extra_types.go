package extra

import (
	"fmt"

	"github.com/l-ross/ficsit-toolkit/save/data"
)

const (
	CircuitSubsystemExtraType Type = "CircuitSubsystem"
	ConveyorBeltExtraType     Type = "ConveyorBelt"
	GameModeExtraType         Type = "GameMode"
	GameStateExtraType        Type = "GameState"
	PlayerStateExtraType      Type = "PlayerState"
	PowerLineExtraType        Type = "PowerLine"
	TrainExtraType            Type = "Train"
	UnknownExtraType          Type = "Unknown"
	VehicleExtraType          Type = "Vehicle"
)

//
// CircuitSubsystem
//

type CircuitSubsystem struct {
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
		Value: &CircuitSubsystem{},
	}
}

func (e *Extra) GetCircuitSubsystem() (*CircuitSubsystem, error) {
	if v, ok := e.Value.(*CircuitSubsystem); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *CircuitSubsystem) parse(d *data.Data) error {
	count, err := d.ReadInt32()
	if err != nil {
		return err
	}

	e.Circuits = make([]*Circuit, count)

	for i := int32(0); i < count; i++ {
		c := &Circuit{}

		c.ID, err = d.ReadInt32()
		if err != nil {
			return err
		}

		c.LevelName, err = d.ReadString()
		if err != nil {
			return err
		}

		c.PathName, err = d.ReadString()
		if err != nil {
			return err
		}

		e.Circuits[i] = c
	}

	return nil
}

func (e *CircuitSubsystem) serialize(d *data.Data) error {
	err := d.WriteInt32(int32(len(e.Circuits)))
	if err != nil {
		return err
	}

	for _, c := range e.Circuits {
		err = d.WriteInt32(c.ID)
		if err != nil {
			return err
		}

		err = d.WriteString(c.LevelName)
		if err != nil {
			return err
		}

		err = d.WriteString(c.PathName)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// ConveyorBelt
//

type ConveyorBelt struct {
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
		Value: &ConveyorBelt{},
	}
}

func (e *Extra) GetConveyorBelt() (*ConveyorBelt, error) {
	if v, ok := e.Value.(*ConveyorBelt); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *ConveyorBelt) parse(d *data.Data) error {
	itemCount, err := d.ReadInt32()
	if err != nil {
		return err
	}

	e.Items = make([]*ConveyorBeltItem, itemCount)

	for i := int32(0); i < itemCount; i++ {
		item := &ConveyorBeltItem{}

		// UNKNOWN_DATA
		_, err = d.ReadInt32()
		if err != nil {
			return err
		}

		item.ResourceName, err = d.ReadString()
		if err != nil {
			return err
		}

		item.LevelName, err = d.ReadString()
		if err != nil {
			return err
		}

		item.PathName, err = d.ReadString()
		if err != nil {
			return err
		}

		item.Position, err = d.ReadFloat32()
		if err != nil {
			return err
		}

		e.Items[i] = item
	}

	return nil
}

func (e *ConveyorBelt) serialize(d *data.Data) error {
	err := d.WriteInt32(int32(len(e.Items)))
	if err != nil {
		return err
	}

	for _, i := range e.Items {
		// UNKNOWN_DATA
		err = d.WriteInt32(0)
		if err != nil {
			return err
		}

		err = d.WriteString(i.ResourceName)
		if err != nil {
			return err
		}

		err = d.WriteString(i.LevelName)
		if err != nil {
			return err
		}

		err = d.WriteString(i.PathName)
		if err != nil {
			return err
		}

		err = d.WriteFloat32(i.Position)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// GameMode
//

type GameMode struct {
	Objects []*Reference
}

func newGameMode(_ int32) *Extra {
	return &Extra{
		Type:  GameModeExtraType,
		Value: &GameMode{},
	}
}

func (e *Extra) GetGameMode() (*GameMode, error) {
	if v, ok := e.Value.(*GameMode); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *GameMode) parse(d *data.Data) error {
	count, err := d.ReadInt32()
	if err != nil {
		return err
	}

	for i := int32(0); i < count; i++ {
		o := &Reference{}

		o.LevelName, err = d.ReadString()
		if err != nil {
			return err
		}

		o.PathName, err = d.ReadString()
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *GameMode) serialize(d *data.Data) error {
	err := d.WriteInt32(int32(len(e.Objects)))
	if err != nil {
		return err
	}

	for _, o := range e.Objects {
		err = d.WriteString(o.LevelName)
		if err != nil {
			return err
		}

		err = d.WriteString(o.PathName)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// GameState
//

type GameState struct {
	Objects []*Reference
}

func newGameState(_ int32) *Extra {
	return &Extra{
		Type:  GameStateExtraType,
		Value: &GameState{},
	}
}

func (e *Extra) GetGameState() (*GameState, error) {
	if v, ok := e.Value.(*GameState); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *GameState) parse(d *data.Data) error {
	count, err := d.ReadInt32()
	if err != nil {
		return err
	}

	for i := int32(0); i < count; i++ {
		o := &Reference{}

		o.LevelName, err = d.ReadString()
		if err != nil {
			return err
		}

		o.PathName, err = d.ReadString()
		if err != nil {
			return err
		}

		e.Objects = append(e.Objects, o)
	}

	return nil
}

func (e *GameState) serialize(d *data.Data) error {
	err := d.WriteInt32(int32(len(e.Objects)))
	if err != nil {
		return err
	}

	for _, o := range e.Objects {
		err = d.WriteString(o.LevelName)
		if err != nil {
			return err
		}

		err = d.WriteString(o.PathName)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// PlayerState
//

type PlayerState struct {
	Data []byte

	len int32
}

func newPlayerState(l int32) *Extra {
	return &Extra{
		Type: PlayerStateExtraType,
		Value: &PlayerState{
			len: l,
		},
	}
}

func (e *Extra) GetPlayerState() (*PlayerState, error) {
	if v, ok := e.Value.(*PlayerState); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *PlayerState) parse(d *data.Data) error {
	var err error
	e.Data, err = d.ReadBytes(e.len)
	if err != nil {
		return err
	}

	return nil
}

func (e *PlayerState) serialize(d *data.Data) error {
	return d.WriteBytes(e.Data)
}

//
// PowerLine
//

type PowerLine struct {
	SourceLevelName string
	SourcePathName  string
	TargetLevelName string
	TargetPathName  string
}

func newPowerLine(_ int32) *Extra {
	return &Extra{
		Type:  PowerLineExtraType,
		Value: &PowerLine{},
	}
}

func (e *Extra) GetPowerLine() (*PowerLine, error) {
	if v, ok := e.Value.(*PowerLine); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *PowerLine) parse(d *data.Data) error {
	var err error
	e.SourceLevelName, err = d.ReadString()
	if err != nil {
		return err
	}

	e.SourcePathName, err = d.ReadString()
	if err != nil {
		return err
	}

	e.TargetLevelName, err = d.ReadString()
	if err != nil {
		return err
	}

	e.TargetPathName, err = d.ReadString()
	if err != nil {
		return err
	}

	return nil
}

func (e *PowerLine) serialize(d *data.Data) error {
	err := d.WriteString(e.SourceLevelName)
	if err != nil {
		return err
	}

	err = d.WriteString(e.SourcePathName)
	if err != nil {
		return err
	}

	err = d.WriteString(e.TargetLevelName)
	if err != nil {
		return err
	}

	err = d.WriteString(e.TargetPathName)
	if err != nil {
		return err
	}

	return nil
}

//
// Train
//

type Train struct {
	PreviousLevelName string
	PreviousPathName  string
	NextLevelName     string
	NextPathName      string
}

func newTrainExtra(_ int32) *Extra {
	return &Extra{
		Type:  TrainExtraType,
		Value: &Train{},
	}
}

func (e *Extra) GetTrain() (*Train, error) {
	if v, ok := e.Value.(*Train); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *Train) parse(d *data.Data) error {
	// UNKNOWN_DATA
	_, err := d.ReadInt32()
	if err != nil {
		return err
	}

	e.PreviousLevelName, err = d.ReadString()
	if err != nil {
		return err
	}

	e.PreviousPathName, err = d.ReadString()
	if err != nil {
		return err
	}

	e.NextLevelName, err = d.ReadString()
	if err != nil {
		return err
	}

	e.NextPathName, err = d.ReadString()
	if err != nil {
		return err
	}

	return nil
}

func (e *Train) serialize(d *data.Data) error {
	err := d.WriteInt32(0)
	if err != nil {
		return err
	}

	err = d.WriteString(e.PreviousLevelName)
	if err != nil {
		return err
	}

	err = d.WriteString(e.PreviousPathName)
	if err != nil {
		return err
	}

	err = d.WriteString(e.NextLevelName)
	if err != nil {
		return err
	}

	err = d.WriteString(e.NextPathName)
	if err != nil {
		return err
	}

	return nil
}

//
// Unknown
//

// UnknownExtra is the default extra that will be returned if we encounter extra data in a
// save.Entity with a TypePath that we don't have explicit handling for.
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

func (e *Extra) GetUnknown() (*UnknownExtra, error) {
	if v, ok := e.Value.(*UnknownExtra); ok {
		return v, nil
	}

	return nil, fmt.Errorf("wrong extra type: %s", e.Type)
}

func (e *UnknownExtra) parse(d *data.Data) error {
	var err error
	e.Data, err = d.ReadBytes(e.len)
	if err != nil {
		return err
	}

	return nil
}

func (e *UnknownExtra) serialize(d *data.Data) error {
	return d.WriteBytes(e.Data)
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

func (e *Vehicle) parse(d *data.Data) error {
	objCount, err := d.ReadInt32()
	if err != nil {
		return err
	}

	for x := int32(0); x < objCount; x++ {
		name, err := d.ReadString()
		if err != nil {
			return err
		}

		data, err := d.ReadBytes(53)
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

func (e *Vehicle) serialize(d *data.Data) error {
	err := d.WriteInt32(int32(len(e.Data)))
	if err != nil {
		return err
	}

	for _, data := range e.Data {
		err := d.WriteString(data.Name)
		if err != nil {
			return err
		}

		err = d.WriteBytes(data.Data)
		if err != nil {
			return err
		}
	}

	return nil
}
