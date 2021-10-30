package save

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ViRb3/slicewriteseek"
	"github.com/stretchr/testify/require"
)

func createTestParser(t *testing.T, filepath string) *Parser {
	data, err := ioutil.ReadFile(filepath)
	require.NoError(t, err)

	p := &Parser{
		body: &slicewriteseek.SliceWriteSeeker{
			Buffer: data,
		},
	}

	return p
}

func createTestSerializer() *Serializer {
	return &Serializer{
		body: slicewriteseek.New(),
	}
}

func assertBuffersEqual(t *testing.T, p *Parser, s *Serializer) {
	assert.Equal(t, p.body.Buffer, s.body.Buffer)
}

func assertAllBufferRead(t *testing.T, p *Parser) {
	assert.Equal(t,
		p.body.Index,
		int64(len(p.body.Buffer)),
		"we should have consumed the entire reader",
	)
}
