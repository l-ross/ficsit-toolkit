package save

import (
	"io/ioutil"
	"testing"

	"github.com/ViRb3/slicewriteseek"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponent_parse(t *testing.T) {
	// Parse a component, then serialize it back and confirm it
	// matches the original.
	t.Parallel()

	data, err := ioutil.ReadFile("testdata/component.dat")
	require.NoError(t, err)

	p := &parser{
		body: &slicewriteseek.SliceWriteSeeker{
			Buffer: data,
		},
	}

	c := &Component{}
	err = c.parse(p)
	require.NoError(t, err)
	assert.Equal(t,
		"Persistent_Level:PersistentLevel.Build_Workshop_C_2146449192.FGFactoryLegs",
		c.InstanceName,
	)

	out := slicewriteseek.New()
	p = &parser{
		body: out,
	}

	err = c.serialize(p)
	require.NoError(t, err)
	assert.Equal(t, data, out.Buffer)
}

func TestComponent_parseData(t *testing.T) {
	// Parse a component's data, then serialize it back and confirm
	// it matches the original.

	data, err := ioutil.ReadFile("testdata/component_data.dat")
	require.NoError(t, err)

	p := &parser{
		body: &slicewriteseek.SliceWriteSeeker{
			Buffer: data,
		},
	}

	c := &Component{}
	err = c.parseData(p)

	out := slicewriteseek.New()
	p = &parser{
		body: out,
	}

	err = c.serializeData(p)
	require.NoError(t, err)
	assert.Equal(t, data, out.Buffer)

	err = ioutil.WriteFile("out.dump", out.Buffer, 0644)
}
