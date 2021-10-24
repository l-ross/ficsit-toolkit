package save

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ViRb3/slicewriteseek"
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

			data, err := ioutil.ReadFile(tt.testData)
			require.NoError(t, err)

			p := &parser{
				body: &slicewriteseek.SliceWriteSeeker{
					Buffer: data,
				},
			}

			e := getExtra(tt.typePath)(int32(len(data)))
			err = e.Value.parse(p)
			require.NoError(t, err)
			assert.Equal(t, p.body.Index, int64(len(data)), "we should have consumed the entire reader")

			tt.assertValue(t, e)

			out := slicewriteseek.New()
			p = &parser{
				body: out,
			}

			err = e.Value.serialize(p)
			require.NoError(t, err)

			assert.Equal(t, data, out.Buffer)
		})
	}
}
