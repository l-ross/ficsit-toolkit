package satisfactorysave

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestStructs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		testData    string
		assertValue func(t *testing.T, s *StructPropertyValue)
	}{
		{
			name:     "Box",
			testData: "testdata/struct_box.dat",
			assertValue: func(t *testing.T, s *StructPropertyValue) {
				v, err := s.GetBoxStruct()
				require.NoError(t, err)
				require.Len(t, v.Min, 3)
				require.Len(t, v.Max, 3)
				assert.Equal(t, float32(-269069.38), v.Min[0])
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			data, err := ioutil.ReadFile(tt.testData)
			require.NoError(t, err)

			p := &Parser{
				body: bytes.NewReader(data),
			}
			objData, err := p.parseObjectData()
			require.NoError(t, err)
			require.NotNil(t, objData)
			assert.Zero(t, p.body.Len(), "we should have consumed the entire reader")

			require.Len(t, objData.Properties, 1, "we should have 1 property")
			s, err := objData.Properties[0].GetStructValue()
			require.NoError(t, err)
			tt.assertValue(t, s)
		})
	}
}
