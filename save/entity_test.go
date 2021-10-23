package save

import (
	"io/ioutil"
	"testing"

	"github.com/ViRb3/slicewriteseek"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntity_parse(t *testing.T) {
	// Parse an entity, then serialize it back and confirm it
	// matches the original.
	t.Parallel()

	data, err := ioutil.ReadFile("testdata/entity.dat")
	require.NoError(t, err)

	p := &parser{
		body: &slicewriteseek.SliceWriteSeeker{
			Buffer: data,
		},
	}

	e := &Entity{}
	err = e.parse(p)
	require.NoError(t, err)
	assert.Equal(t,
		"Persistent_Level:PersistentLevel.FGPipeNetwork_2147302283",
		e.InstanceName,
	)

	out := slicewriteseek.New()
	p = &parser{
		body: out,
	}

	err = e.serialize(p)
	require.NoError(t, err)

	assert.Equal(t, data, out.Buffer)
}

func TestEntity_parseData(t *testing.T) {
	// Parse an entity's data, then serialize it back and confirm
	// it matches the original.
	t.Parallel()

	data, err := ioutil.ReadFile("testdata/entity_data.dat")
	require.NoError(t, err)

	p := &parser{
		body: &slicewriteseek.SliceWriteSeeker{
			Buffer: data,
		},
	}

	e := &Entity{}
	err = e.parseData(p)
	require.NoError(t, err)
	assert.Len(t, e.Properties, 3)

	out := slicewriteseek.New()
	p = &parser{
		body: out,
	}

	err = e.serializeData(p)
	require.NoError(t, err)

	assert.Equal(t, data, out.Buffer)
}
