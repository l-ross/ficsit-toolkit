package save

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ViRb3/slicewriteseek"
	"github.com/stretchr/testify/require"
)

func createTestParser(t *testing.T, filepath string) *parser {
	data, err := ioutil.ReadFile(filepath)
	require.NoError(t, err)

	p := &parser{
		body: &slicewriteseek.SliceWriteSeeker{
			Buffer: data,
		},
	}

	return p
}

func createTestSerializer() *serializer {
	return &serializer{
		body: slicewriteseek.New(),
	}
}

func assertBuffersEqual(t *testing.T, p *parser, s *serializer) {
	assert.Equal(t, p.body.Buffer, s.body.Buffer)
}

func assertAllBufferRead(t *testing.T, p *parser) {
	assert.Equal(t,
		p.body.Index,
		int64(len(p.body.Buffer)),
		"we should have consumed the entire reader",
	)
}
