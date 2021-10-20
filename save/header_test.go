package save

import (
	"io/ioutil"
	"testing"

	"github.com/ViRb3/slicewriteseek"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeader(t *testing.T) {
	// Parse a header then serialize it back and confirm it
	// matches the original data.
	t.Parallel()

	data, err := ioutil.ReadFile("testdata/header.dat")
	require.NoError(t, err)

	p := &parser{
		body: &slicewriteseek.SliceWriteSeeker{
			Buffer: data,
		},
	}

	h, err := p.parseHeader()
	require.NoError(t, err)

	out := slicewriteseek.New()
	p = &parser{
		body: out,
	}

	err = p.serializeHeader(h)
	require.NoError(t, err)

	assert.Equal(t, data, out.Buffer)
}
