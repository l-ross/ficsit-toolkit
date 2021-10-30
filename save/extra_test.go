package save

import (
	"testing"

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
			testData: "testdata/extra/circuit_subsystem.dat",
			typePath: "/Game/FactoryGame/-Shared/Blueprint/BP_CircuitSubsystem.BP_CircuitSubsystem_C",
			assertValue: func(t *testing.T, e *Extra) {
				c, err := e.getCircuitSubsystem()
				require.NoError(t, err)
				assert.Len(t, c.Circuits, 1)
			},
		},
		{
			name:     "ConveyorBelt",
			testData: "testdata/extra/conveyor_belt.dat",
			typePath: "/Game/FactoryGame/Buildable/Factory/ConveyorBeltMk1/Build_ConveyorBeltMk1.Build_ConveyorBeltMk1_C",
			assertValue: func(t *testing.T, e *Extra) {
				c, err := e.getConveyorBelt()
				require.NoError(t, err)
				assert.Len(t, c.Items, 4)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Parse extra
			p := createTestParser(t, tt.testData)
			e := getExtra(tt.typePath)(int32(len(p.body.Buffer)))
			err := e.Value.parse(p)
			require.NoError(t, err)
			assertAllBufferRead(t, p)

			// Verify extra
			tt.assertValue(t, e)

			// Serialize extra
			s := createTestSerializer()
			err = e.Value.serialize(s)
			require.NoError(t, err)

			// Verify serialization is correct
			assertBuffersEqual(t, p, s)
		})
	}
}
