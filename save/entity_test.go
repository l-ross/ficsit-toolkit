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
	p := createTestParser(t, "testdata/entity.dat")
	e := &Entity{}
	err := e.parse(p)
	require.NoError(t, err)
	assertAllBufferRead(t, p)

	// Verify entity
	assert.Equal(t,
		"Persistent_Level:PersistentLevel.FGPipeNetwork_2147302283",
		e.InstanceName,
	)

	// Serialize entity
	s := createTestSerializer()
	err = e.serialize(s)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p, s)
}

func TestEntity_parseData(t *testing.T) {
	// Parse an entity's data, then serialize it back and confirm
	// it matches the original.
	t.Parallel()

	// Parse entity data
	p := createTestParser(t, "testdata/entity_data.dat")
	e := &Entity{}
	err := e.parseData(p)
	require.NoError(t, err)
	assertAllBufferRead(t, p)

	// Verify entity data
	assert.Len(t, e.Properties, 3)

	// Serialize entity data
	s := createTestSerializer()
	err = e.serializeData(s)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p, s)
}
