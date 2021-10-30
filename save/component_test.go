package save

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponent_parse(t *testing.T) {
	// Parse a component, then serialize it back and confirm it
	// matches the original.
	t.Parallel()

	// Parse component
	p := createTestParser(t, "testdata/component.dat")
	c := &Component{}
	err := c.parse(p)
	require.NoError(t, err)
	assertAllBufferRead(t, p)

	// Verify component
	assert.Equal(t,
		"Persistent_Level:PersistentLevel.Build_Workshop_C_2146449192.FGFactoryLegs",
		c.InstanceName,
	)

	// Serialize component
	s := createTestSerializer()
	err = c.serialize(s)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p, s)
}

func TestComponent_parseData(t *testing.T) {
	// Parse a component's data, then serialize it back and confirm
	// it matches the original.

	// Parse component data
	p := createTestParser(t, "testdata/component_data.dat")
	c := &Component{}
	err := c.parseData(p)
	require.NoError(t, err)
	assertAllBufferRead(t, p)

	// Verify component data
	require.Len(t, c.Properties, 1)
	assert.Equal(t, "mCachedFeetOffset", c.Properties[0].Name)

	// Serialize component data
	s := createTestSerializer()
	err = c.serializeData(s)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p, s)
}
