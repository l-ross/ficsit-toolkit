package save

type ExtraType string

// TODO: Train
// TODO: Vehicle
const (
	CircuitSubsystemExtraType ExtraType = "CircuitSubsystem"
	ConveyorBeltExtraType     ExtraType = "ConveyorBelt"
	GameModeExtraType         ExtraType = "GameMode"
	GameStateExtraType        ExtraType = "GameState"
	PowerLineExtraType        ExtraType = "PowerLine"
	UnknownExtraType          ExtraType = "Unknown"
)

type Extra struct {
	Type  ExtraType
	Value ExtraValue
}

type ExtraValue interface {
	parse(p *parser) error
}

func hasExtra(c string) func() *Extra {
	switch c {
	case "/Game/FactoryGame/-Shared/Blueprint/BP_CircuitSubsystem.BP_CircuitSubsystem_C":
		return newCircuitSubsystem
	case "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk1/Build_ConveyorBeltMk1.Build_ConveyorBeltMk1_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk2/Build_ConveyorBeltMk1.Build_ConveyorBeltMk2_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk3/Build_ConveyorBeltMk3.Build_ConveyorBeltMk3_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk4/Build_ConveyorBeltMk4.Build_ConveyorBeltMk4_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk5/Build_ConveyorBeltMk5.Build_ConveyorBeltMk5_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk1/Build_ConveyorLiftMk1.Build_ConveyorLiftMk1_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk2/Build_ConveyorLiftMk2.Build_ConveyorLiftMk2_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk3/Build_ConveyorLiftMk3.Build_ConveyorLiftMk3_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk4/Build_ConveyorLiftMk1.Build_ConveyorLiftMk1_C",
		"/Game/FactoryGame/Buildable/Factory/ConveyorLiftMk5/Build_ConveyorLiftMk5.Build_ConveyorLiftMk5_C":
		return newConveyorBelt
	case "/Game/FactoryGame/-Shared/Blueprint/BP_GameMode.BP_GameMode_C":
		return newGameMode
	case "/Game/FactoryGame/-Shared/Blueprint/BP_GameState.BP_GameState_C":
		return newGameState
	case "/Game/FactoryGame/Buildable/Factory/PowerLine/Build_PowerLine.Build_PowerLine_C",
		"/Game/FactoryGame/Events/Christmas/Buildings/PowerLineLights/Build_XmassLightsLine.Build_XmassLightsLine_C":
		return newPowerLine
	}
	return nil
}

//
// CircuitSubsystem

type CircuitSubsystemExtra struct {
	Circuits []*Circuit
}

type Circuit struct {
	ID        int32
	LevelName string
	PathName  string
}

func newCircuitSubsystem() *Extra {
	return &Extra{
		Type:  CircuitSubsystemExtraType,
		Value: &CircuitSubsystemExtra{},
	}
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

func newConveyorBelt() *Extra {
	return &Extra{
		Type:  ConveyorBeltExtraType,
		Value: &ConveyorBeltExtra{},
	}
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

//
// GameMode
//

type GameModeExtra struct {
	Objects []*ObjectReference
}

func newGameMode() *Extra {
	return &Extra{
		Type:  GameModeExtraType,
		Value: &GameModeExtra{},
	}
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

//
// GameState
//

type GameStateExtra struct {
	Objects []*ObjectReference
}

func newGameState() *Extra {
	return &Extra{
		Type:  GameStateExtraType,
		Value: &GameStateExtra{},
	}
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
	}

	return nil
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

func newPowerLine() *Extra {
	return &Extra{
		Type:  PowerLineExtraType,
		Value: &PowerLineExtra{},
	}
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

//
// UnknownExtra
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

func (e *UnknownExtra) parse(p *parser) error {
	var err error
	e.Data, err = p.readBytes(e.len)
	if err != nil {
		return err
	}

	return nil
}
