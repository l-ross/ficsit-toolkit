package save

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestProperties(t *testing.T) {
	t.Parallel()

	// Parse properties, verify the properties value then serialise
	// back and verify it matches the original.

	tests := []struct {
		name        string
		testData    string
		assertValue func(t *testing.T, p *Property)
	}{
		{
			name:     "ArrayProperty_InterfaceValueType",
			testData: "testdata/property_type/array_interface.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetArrayValue()
				require.NoError(t, err)
				assert.Equal(t, &ArrayPropertyValue{
					ValueType: InterfacePropertyType,
					Values: []PropertyValue{
						&InterfacePropertyValue{
							LevelName: "LevelName1",
							PathName:  "PathName1",
						},
						&InterfacePropertyValue{
							LevelName: "LevelName2",
							PathName:  "PathName2",
						},
					},
				}, v)
			},
		},
		{
			name:     "ArrayProperty_StructProperty",
			testData: "testdata/property_type/array_struct.dat",
			assertValue: func(t *testing.T, p *Property) {
				s1, err := p.GetStructValue()
				require.NoError(t, err)
				a, err := s1.GetArbitraryStruct()
				require.NoError(t, err)
				require.Len(t, a.Properties, 1)
				arr, err := a.Properties[0].GetArrayValue()
				require.NoError(t, err)
				require.Len(t, arr.Values, 1)

			},
		},
		{
			name:     "BoolProperty",
			testData: "testdata/property_type/bool.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetBoolValue()
				require.NoError(t, err)
				assert.True(t, v)
			},
		},
		{
			name:     "ByteProperty",
			testData: "testdata/property_type/byte.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetByteValue()
				require.NoError(t, err)
				assert.Equal(t, []byte{0xAA}, v)
			},
		},
		{
			name:     "DoubleProperty",
			testData: "testdata/property_type/double.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetDoubleValue()
				require.NoError(t, err)
				assert.Equal(t, float64(1), v)
			},
		},
		{
			name:     "FloatProperty",
			testData: "testdata/property_type/float.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetFloatValue()
				require.NoError(t, err)
				assert.Equal(t, float32(1), v)
			},
		},
		{
			name:     "Int8Property",
			testData: "testdata/property_type/int8.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetInt8Value()
				require.NoError(t, err)
				assert.Equal(t, int8(1), v)
			},
		},
		{
			name:     "Int64Property",
			testData: "testdata/property_type/int64.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetInt64Value()
				require.NoError(t, err)
				assert.Equal(t, int64(1), v)
			},
		},
		{
			name:     "InterfaceProperty",
			testData: "testdata/property_type/interface.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetInterfaceValue()
				require.NoError(t, err)
				assert.Equal(t, "TheLevelName", v.LevelName)
				assert.Equal(t, "ThePathName", v.PathName)
			},
		},
		{
			name:     "IntProperty",
			testData: "testdata/property_type/int.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetIntValue()
				require.NoError(t, err)
				assert.Equal(t, int32(4), v)
			},
		},
		{
			name:     "MapProperty_IntKey_IntValue",
			testData: "testdata/property_type/map_int_int.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetMapPropertyValue()
				require.NoError(t, err)
				assert.Equal(t, IntPropertyType, v.KeyType)
				assert.Equal(t, IntPropertyType, v.ValueType)
				assert.Len(t, v.Values, 20)
			},
		},
		{
			name:     "NameProperty",
			testData: "testdata/property_type/name.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetNameValue()
				require.NoError(t, err)
				assert.Equal(t, "AStringValue", v)
			},
		},
		{
			name:     "ObjectProperty",
			testData: "testdata/property_type/object.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetObjectValue()
				require.NoError(t, err)
				assert.Equal(t, "TheLevelName", v.LevelName)
				assert.Equal(t, "ThePathName", v.PathName)
			},
		},
		{
			name:     "StringProperty",
			testData: "testdata/property_type/string.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetStringValue()
				require.NoError(t, err)
				assert.Equal(t, "AStringValue", v)
			},
		},
		{
			name:     "StringProperty_UTF16",
			testData: "testdata/property_type/string_utf16.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetStringValue()
				require.NoError(t, err)
				assert.Equal(t, "Uranium Rods In ↑\r\nRadioactive: ☢", v)
			},
		},
		{
			name:     "StructProperty",
			testData: "testdata/property_type/struct.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetStructValue()
				require.NoError(t, err)
				a, err := v.GetArbitraryStruct()
				require.NoError(t, err)
				require.Len(t, a.Properties, 3)
			},
		},
		{
			name:     "TextProperty_1",
			testData: "testdata/property_type/text_1.dat",
			assertValue: func(t *testing.T, p *Property) {
				txt, err := p.GetTextValue()
				require.NoError(t, err)

				noneTxt, err := txt.GetNoneText()
				require.NoError(t, err)
				assert.Equal(t, "Skoevde C", noneTxt.String)
			},
		},
		{
			name:     "TextProperty_2",
			testData: "testdata/property_type/text_2.dat",
			assertValue: func(t *testing.T, p *Property) {
				txt, err := p.GetTextValue()
				require.NoError(t, err)

				baseTxt, err := txt.GetBaseText()
				require.NoError(t, err)
				assert.Equal(t, "Custom Name", baseTxt.Value)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Parse property
			p := createTestParser(t, tt.testData)
			props, err := p.parseProperties()
			require.NoError(t, err)
			assertAllBufferRead(t, p)

			// Verify property
			require.Len(t, props, 1, "we should have 1 property")
			prop := props[0]
			tt.assertValue(t, prop)

			// Serialize property
			s := createTestSerializer()
			err = serializeProperties(props, s.Data)
			require.NoError(t, err)

			// Verify serialization is correct
			assertBuffersEqual(t, p, s)
		})
	}
}
