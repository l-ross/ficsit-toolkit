package save

import (
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
			testData: "testdata/struct/box.dat",
			assertValue: func(t *testing.T, s *StructPropertyValue) {
				v, err := s.GetBoxStruct()
				require.NoError(t, err)
				require.Len(t, v.Min, 3)
				require.Len(t, v.Max, 3)
				assert.Equal(t, float32(-269069.38), v.Min[0])
			},
		},
		{
			name:     "InventoryItem",
			testData: "testdata/struct/inventory_item.dat",
			assertValue: func(t *testing.T, s *StructPropertyValue) {
				s1, err := s.GetArbitraryStruct()
				require.NoError(t, err)
				require.Len(t, s1.Properties, 1)
				s2, err := s1.Properties[0].GetStructValue()
				require.NoError(t, err)
				v, err := s2.GetInventoryItemStruct()
				require.NoError(t, err)
				assert.Equal(t, "/Game/FactoryGame/Resource/Parts/IronPlateReinforced/Desc_IronPlateReinforced.Desc_IronPlateReinforced_C", v.ItemName)
				assert.Equal(t, int32(24), v.NumItems)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Parse struct
			p1 := createTestParserFromFile(t, tt.testData)
			props, err := p1.parseProperties()
			require.NoError(t, err)
			assertAllBufferRead(t, p1)

			// Verify struct
			require.Len(t, props, 1, "we should have 1 property")
			s, err := props[0].GetStructValue()
			require.NoError(t, err)
			tt.assertValue(t, s)

			// Serialize struct
			p2 := createTestParserInMemory()
			err = p2.serializeProperties(props)
			require.NoError(t, err)

			// Verify serialization is correct
			assertBuffersEqual(t, p1, p2)
		})
	}
}
