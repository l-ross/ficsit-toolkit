package save

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestHeader(t *testing.T) {
	// Parse a header then serialize it back and confirm it
	// matches the original data.
	t.Parallel()

	// Parse header
	p1 := createTestParserFromFile(t, "testdata/header.dat")
	h, err := p1.parseHeader()
	require.NoError(t, err)
	assertAllBufferRead(t, p1)

	// Verify header
	assert.Equal(t, "Luke", h.SessionName)
	assert.Equal(t, int32(152331), h.BuildVersion)

	// Serialize header
	p2 := createTestParserInMemory()
	err = p2.serializeHeader(h)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p1, p2)
}
