package satisfactorysave

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestPropertyValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		testData    string
		assertValue func(t *testing.T, p *Property)
	}{
		{
			name:     "ArrayProperty_InterfaceValueType",
			testData: "testdata/array_interface.prop",
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
			name:     "BoolProperty",
			testData: "testdata/bool.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetBoolValue()
				require.NoError(t, err)
				assert.True(t, v)
			},
		},
		{
			name:     "ByteProperty",
			testData: "testdata/byte.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetByteValue()
				require.NoError(t, err)
				assert.Equal(t, []byte{0xAA}, v)
			},
		},
		{
			name:     "DoubleProperty",
			testData: "testdata/double.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetDoubleValue()
				require.NoError(t, err)
				assert.Equal(t, float64(1), v)
			},
		},
		{
			name:     "FloatProperty",
			testData: "testdata/float.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetFloatValue()
				require.NoError(t, err)
				assert.Equal(t, float32(1), v)
			},
		},
		{
			name:     "Int8Property",
			testData: "testdata/int8.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetInt8Value()
				require.NoError(t, err)
				assert.Equal(t, int8(1), v)
			},
		},
		{
			name:     "Int64Property",
			testData: "testdata/int64.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetInt64Value()
				require.NoError(t, err)
				assert.Equal(t, int64(1), v)
			},
		},
		{
			name:     "IntProperty",
			testData: "testdata/int.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetIntValue()
				require.NoError(t, err)
				assert.Equal(t, int32(4), v)
			},
		},
		{
			name:     "InterfaceProperty",
			testData: "testdata/interface.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetInterfaceValue()
				require.NoError(t, err)
				assert.Equal(t, "TheLevelName", v.LevelName)
				assert.Equal(t, "ThePathName", v.PathName)
			},
		},
		{
			name:     "NameProperty",
			testData: "testdata/name.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetNameValue()
				require.NoError(t, err)
				assert.Equal(t, "AStringValue", v)
			},
		},
		{
			name:     "ObjectProperty",
			testData: "testdata/object.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetObjectValue()
				require.NoError(t, err)
				assert.Equal(t, "TheLevelName", v.LevelName)
				assert.Equal(t, "ThePathName", v.PathName)
			},
		},
		{
			name:     "StringProperty",
			testData: "testdata/string.prop",
			assertValue: func(t *testing.T, p *Property) {
				v, err := p.GetStringValue()
				require.NoError(t, err)
				assert.Equal(t, "AStringValue", v)
			},
		},
		{
			name:     "StructProperty",
			testData: "testdata/struct.prop",
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

			p := &Parser{
				body: bytes.NewReader(data),
			}
			objData, err := p.parseObjectData()
			require.NoError(t, err)
			require.NotNil(t, objData)
			require.Len(t, objData.Properties, 1, "we should have 1 property")
			assert.Zero(t, p.body.Len(), "we should have consumed the entire reader")

			tt.assertValue(t, objData.Properties[0])
		})
	}
}
