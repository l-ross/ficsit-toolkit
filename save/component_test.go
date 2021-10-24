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
	p1 := createTestParserFromFile(t, "testdata/component.dat")
	c := &Component{}
	err := c.parse(p1)
	require.NoError(t, err)
	assertAllBufferRead(t, p1)

	// Verify component
	assert.Equal(t,
		"Persistent_Level:PersistentLevel.Build_Workshop_C_2146449192.FGFactoryLegs",
		c.InstanceName,
	)

	// Serialize component
	p2 := createTestParserInMemory()
	err = c.serialize(p2)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p1, p2)
}

func TestComponent_parseData(t *testing.T) {
	// Parse a component's data, then serialize it back and confirm
	// it matches the original.

	// Parse component data
	p1 := createTestParserFromFile(t, "testdata/component_data.dat")
	c := &Component{}
	err := c.parseData(p1)
	require.NoError(t, err)
	assertAllBufferRead(t, p1)

	// Verify component data
	require.Len(t, c.Properties, 1)
	assert.Equal(t, "mCachedFeetOffset", c.Properties[0].Name)

	// Serialize component data
	p2 := createTestParserInMemory()
	err = c.serializeData(p2)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p1, p2)
}
