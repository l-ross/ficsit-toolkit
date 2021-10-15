package save

import (
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
		{
			// Note that some lengths defined in the data are incorrect.
			// Version Info:
			// - SaveVersion = 25
			// - BuildVersion = 152331
			name:     "InventoryItem",
			testData: "testdata/prop_struct_inventoryitem.dat",
			assertValue: func(t *testing.T, s *StructPropertyValue) {
				s1, err := s.GetArbitraryStruct()
				require.NoError(t, err)
				require.Len(t, s1.Properties, 1)
				s2, err := s1.Properties[0].GetStructValue()
				require.NoError(t, err)
				v, err := s2.GetInventoryItemStruct()
				require.NoError(t, err)
				assert.Equal(t, "/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C", v.ItemName)
				assert.Equal(t, int32(4), v.NumItems)
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
				body: data,
			}
			objData, err := p.parseObjectData(0, int32(len(data)))
			require.NoError(t, err)
			require.NotNil(t, objData)
			assert.Zero(t, p.buf.Len(), "we should have consumed the entire reader")

			require.Len(t, objData.Properties, 1, "we should have 1 property")
			s, err := objData.Properties[0].GetStructValue()
			require.NoError(t, err)
			tt.assertValue(t, s)
		})
	}
}
