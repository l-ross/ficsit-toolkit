package property

import (
	"testing"

	"github.com/l-ross/ficsit-toolkit/save/data"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStructs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		structType  StructType
		testData    string
		assertValue func(t *testing.T, s *StructPropertyValue)
	}{
		{
			name:       "Box",
			structType: BoxStructType,
			testData:   "testdata/struct_type/box.dat",
			assertValue: func(t *testing.T, s *StructPropertyValue) {
				v, err := s.GetBoxStruct()
				require.NoError(t, err)
				require.Len(t, v.Min, 3)
				require.Len(t, v.Max, 3)
				assert.Equal(t, float32(-269069.38), v.Min[0])
			},
		},
		{
			name:       "Color",
			structType: ColorStructType,
			testData:   "testdata/struct_type/color.dat",
			assertValue: func(t *testing.T, s *StructPropertyValue) {
				v, err := s.GetColorStruct()
				require.NoError(t, err)
				assert.Equal(t, byte(0xFF), v.G)
			},
		},
		{
			name:       "InventoryItem",
			structType: InventoryItemStructType,
			testData:   "testdata/struct_type/inventory_item.dat",
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
			d := data.TestData(t, tt.testData)
			props, err := ParseProperties(d)
			require.NoError(t, err)
			data.AssertAllBufferRead(t, d)

			// Verify struct
			require.Len(t, props, 1, "we should have 1 property")
			structValue, err := props[0].GetStructValue()
			require.NoError(t, err)
			tt.assertValue(t, structValue)

			// Serialize struct
			d2 := data.New()
			err = SerializeProperties(props, d2)
			require.NoError(t, err)

			// Verify serialization is correct
			assert.Equal(t, d.Bytes(), d2.Bytes())
		})
	}
}
