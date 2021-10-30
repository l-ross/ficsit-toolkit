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
	p := createTestParser(t, "testdata/header.dat")
	h, err := p.parseHeader()
	require.NoError(t, err)
	assertAllBufferRead(t, p)

	// Verify header
	assert.Equal(t, "Luke", h.SessionName)
	assert.Equal(t, int32(152331), h.BuildVersion)

	// Serialize header
	s := createTestSerializer()
	err = s.serializeHeader(h)
	require.NoError(t, err)

	// Verify serialization is correct
	assertBuffersEqual(t, p, s)
}
