package extra

import (
	"testing"

	"github.com/l-ross/ficsit-toolkit/save/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtra(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		testData    string
		typePath    string
		assertValue func(t *testing.T, e *Extra)
	}{
		{
			name:     "CircuitSubsystem",
			testData: "testdata/circuit_subsystem.dat",
			typePath: "/Game/FactoryGame/-Shared/Blueprint/BP_CircuitSubsystem.BP_CircuitSubsystem_C",
			assertValue: func(t *testing.T, e *Extra) {
				c, err := e.GetCircuitSubsystem()
				require.NoError(t, err)
				require.Len(t, c.Circuits, 1)
				assert.Equal(t, "Persistent_Level:PersistentLevel.CircuitSubsystem.FGPowerCircuit_2147247253", c.Circuits[0].PathName)
			},
		},
		{
			name:     "ConveyorBelt",
			testData: "testdata/conveyor_belt.dat",
			typePath: "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk1/Build_ConveyorBeltMk1.Build_ConveyorBeltMk1_C",
			assertValue: func(t *testing.T, e *Extra) {
				c, err := e.GetConveyorBelt()
				require.NoError(t, err)
				require.Len(t, c.Items, 4)
				assert.Equal(t, "/Game/FactoryGame/Resource/Parts/IronPlate/Desc_IronPlate.Desc_IronPlate_C", c.Items[0].ResourceName)
			},
		},
		{
			name:     "GameState",
			testData: "testdata/game_state.dat",
			typePath: "/Game/FactoryGame/-Shared/Blueprint/BP_GameState.BP_GameState_C",
			assertValue: func(t *testing.T, e *Extra) {
				g, err := e.GetGameState()
				require.NoError(t, err)
				require.Len(t, g.Objects, 1)
				assert.Equal(t, "Persistent_Level:PersistentLevel.BP_PlayerState_C_2147479729", g.Objects[0].PathName)
			},
		},
		{
			name:     "PlayerState",
			testData: "testdata/player_state.dat",
			typePath: "/Game/FactoryGame/Character/Player/BP_PlayerState.BP_PlayerState_C",
			assertValue: func(t *testing.T, e *Extra) {
				_, err := e.GetPlayerState()
				require.NoError(t, err)
			},
		},
		{
			name:     "PowerLine",
			testData: "testdata/power_line.dat",
			typePath: "/Game/FactoryGame/Buildable/Factory/PowerLine/Build_PowerLine.Build_PowerLine_C",
			assertValue: func(t *testing.T, e *Extra) {
				p, err := e.GetPowerLine()
				require.NoError(t, err)
				assert.Equal(t, "Persistent_Level:PersistentLevel.Build_PowerPoleMk1_C_2147403233.PowerConnection", p.SourcePathName)
			},
		},
		{
			name:     "Train",
			testData: "testdata/train.dat",
			typePath: "/Game/FactoryGame/Buildable/Vehicle/Train/Wagon/BP_FreightWagon.BP_FreightWagon_C",
			assertValue: func(t *testing.T, e *Extra) {
				tr, err := e.GetTrain()
				require.NoError(t, err)
				assert.Equal(t, "Persistent_Level:PersistentLevel.BP_Locomotive_C_2147390965", tr.PreviousPathName)
			},
		},
		{
			name:     "Vehicle",
			testData: "testdata/vehicle.dat",
			typePath: "/Game/FactoryGame/Buildable/Vehicle/Tractor/BP_Tractor.BP_Tractor_C",
			assertValue: func(t *testing.T, e *Extra) {
				v, err := e.GetVehicle()
				require.NoError(t, err)
				require.Len(t, v.Data, 1)
				assert.Equal(t, "Root", v.Data[0].Name)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Parse extra
			d := data.TestData(t, tt.testData)
			e, err := Parse(tt.typePath, int32(d.Len()), d)
			require.NoError(t, err)
			data.AssertAllBufferRead(t, d)

			// Verify extra
			tt.assertValue(t, e)

			// Serialize extra
			d2 := data.New()
			err = e.Value.serialize(d2)
			require.NoError(t, err)

			// Verify serialization is correct
			assert.Equal(t, d.Bytes(), d2.Bytes())
		})
	}
}
