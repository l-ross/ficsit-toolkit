package save

import (
	"io/ioutil"
	"testing"

	"github.com/ViRb3/slicewriteseek"

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
			testData: "testdata/prop_array_interface.dat",
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
			testData: "testdata/prop_array_struct.dat",
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
			testData: "testdata/prop_bool.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetBoolValue()
				require.NoError(t, err)
				assert.True(t, v)
			},
		},
		{
			name:     "ByteProperty",
			testData: "testdata/prop_byte.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetByteValue()
				require.NoError(t, err)
				assert.Equal(t, []byte{0xAA}, v)
			},
		},
		{
			name:     "DoubleProperty",
			testData: "testdata/prop_double.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetDoubleValue()
				require.NoError(t, err)
				assert.Equal(t, float64(1), v)
			},
		},
		{
			name:     "FloatProperty",
			testData: "testdata/prop_float.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetFloatValue()
				require.NoError(t, err)
				assert.Equal(t, float32(1), v)
			},
		},
		{
			name:     "Int8Property",
			testData: "testdata/prop_int8.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetInt8Value()
				require.NoError(t, err)
				assert.Equal(t, int8(1), v)
			},
		},
		{
			name:     "Int64Property",
			testData: "testdata/prop_int64.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetInt64Value()
				require.NoError(t, err)
				assert.Equal(t, int64(1), v)
			},
		},
		{
			name:     "InterfaceProperty",
			testData: "testdata/prop_interface.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetInterfaceValue()
				require.NoError(t, err)
				assert.Equal(t, "TheLevelName", v.LevelName)
				assert.Equal(t, "ThePathName", v.PathName)
			},
		},
		{
			name:     "IntProperty",
			testData: "testdata/prop_int.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetIntValue()
				require.NoError(t, err)
				assert.Equal(t, int32(4), v)
			},
		},
		{
			name:     "MapProperty_IntKey_IntValue",
			testData: "testdata/prop_map_int_int.dat",
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
			testData: "testdata/prop_name.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetNameValue()
				require.NoError(t, err)
				assert.Equal(t, "AStringValue", v)
			},
		},
		{
			name:     "ObjectProperty",
			testData: "testdata/prop_object.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetObjectValue()
				require.NoError(t, err)
				assert.Equal(t, "TheLevelName", v.LevelName)
				assert.Equal(t, "ThePathName", v.PathName)
			},
		},
		{
			name:     "StringProperty",
			testData: "testdata/prop_string.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetStringValue()
				require.NoError(t, err)
				assert.Equal(t, "AStringValue", v)
			},
		},
		{
			name:     "StructProperty",
			testData: "testdata/prop_struct.dat",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetStructValue()
				require.NoError(t, err)
				a, err := v.GetArbitraryStruct()
				require.NoError(t, err)
				require.Len(t, a.Properties, 3)
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

			props, err := p.parseProperties()
			require.NoError(t, err)
			assert.Equal(t, p.body.Index, int64(len(data)), "we should have consumed the entire reader")

			require.Len(t, props, 1, "we should have 1 property")
			prop := props[0]
			tt.assertValue(t, prop)

			out := slicewriteseek.New()
			p = &parser{
				body: out,
			}

			err = p.serializeProperties(props)
			require.NoError(t, err)

			assert.Equal(t, data, out.Buffer)

			err = ioutil.WriteFile("out.dump", out.Buffer, 0644)
			require.NoError(t, err)
		})
	}
}
