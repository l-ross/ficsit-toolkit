package save

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntity_parse(t *testing.T) {
	// Parse an entity, then serialize it back and confirm it
	// matches the original.
	t.Parallel()

	// Parse entity
	p1 := createTestParserFromFile(t, "testdata/entity.dat")
	e := &Entity{}
	err := e.parse(p1)
	require.NoError(t, err)
	assertAllBufferRead(t, p1)

	// Verify entity
	assert.Equal(t,
		"Persistent_Level:PersistentLevel.FGPipeNetwork_2147302283",
		e.InstanceName,
	)

	// Serialize entity
	p2 := createTestParserInMemory()
	err = e.serialize(p2)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p1, p2)
}

func TestEntity_parseData(t *testing.T) {
	// Parse an entity's data, then serialize it back and confirm
	// it matches the original.
	t.Parallel()

	// Parse entity data
	p1 := createTestParserFromFile(t, "testdata/entity_data.dat")
	e := &Entity{}
	err := e.parseData(p1)
	require.NoError(t, err)
	assertAllBufferRead(t, p1)

	// Verify entity data
	assert.Len(t, e.Properties, 3)

	// Serialize entity data
	p2 := createTestParserInMemory()
	err = e.serializeData(p2)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p1, p2)
}
