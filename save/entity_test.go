package save

import (
	"io/ioutil"
	"testing"

	"github.com/mattetti/filebuffer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntity(t *testing.T) {
	// Parse an entity, then serialize it back and confirm it
	// matches the original.
	t.Parallel()

	data, err := ioutil.ReadFile("testdata/entity.dat")
	require.NoError(t, err)

	p := &Parser{
		body: filebuffer.New(data),
	}

	e, err := p.parseEntity()
	require.NoError(t, err)
	assert.Equal(t,
		"Persistent_Level:PersistentLevel.FGPipeNetwork_2147302283",
		e.InstanceName,
	)

	out := filebuffer.New([]byte{})
	s := &Save{
		body: out,
	}

	err = s.serializeEntity(e)
	require.NoError(t, err)

	assert.Equal(t, data, out.Buff.Bytes())
}
