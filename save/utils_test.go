package save

import (
	"io/ioutil"
	"testing"

	"github.com/l-ross/ficsit-toolkit/save/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTestParser(t *testing.T, filepath string) *parser {
	b, err := ioutil.ReadFile(filepath)
	require.NoError(t, err)

	p := &parser{
		Data: data.NewFromBytes(b),
	}

	return p
}

func createTestSerializer() *serializer {
	return &serializer{
		Data: data.New(),
	}
}

func assertBuffersEqual(t *testing.T, p *parser, s *serializer) {
	assert.Equal(t, p.Bytes(), s.Bytes())
}

func assertAllBufferRead(t *testing.T, p *parser) {
	assert.Equal(t,
		p.Index(),
		int64(len(p.Bytes())),
		"we should have consumed the entire reader",
	)
}
