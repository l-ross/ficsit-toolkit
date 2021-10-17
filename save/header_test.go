package save

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeader(t *testing.T) {
	// Parse a header then serialize it back and confirm it
	// matches the original.

	data, err := ioutil.ReadFile("testdata/header.dat")
	require.NoError(t, err)

	p := &Parser{
		body: bytes.NewReader(data),
	}

	h1, err := p.ParseHeader()
	require.NoError(t, err)

	out := &bytes.Buffer{}
	s := &Save{
		Header: h1,
		w:      out,
	}

	err = s.serializeHeader()
	require.NoError(t, err)
	assert.Equal(t, data, out.Bytes())
}
