package save

import (
	"io/ioutil"
	"testing"

	"github.com/mattetti/filebuffer"
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
		body: filebuffer.New(data),
	}

	h, err := p.parseHeader()
	require.NoError(t, err)

	out := filebuffer.New([]byte{})
	p = &parser{
		body: out,
	}

	err = p.serializeHeader(h)
	require.NoError(t, err)

	assert.Equal(t, data, out.Buff.Bytes())
}
