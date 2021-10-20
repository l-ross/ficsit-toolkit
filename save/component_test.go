package save

import (
	"io/ioutil"
	"testing"

	"github.com/mattetti/filebuffer"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComponent(t *testing.T) {
	// Parse a component, then serialize it back and confirm it
	// matches the original.
	t.Parallel()

	data, err := ioutil.ReadFile("testdata/component.dat")
	require.NoError(t, err)

	p := &parser{
		body: filebuffer.New(data),
	}

	c, err := p.parseComponent()
	require.NoError(t, err)
	assert.Equal(t,
		"Persistent_Level:PersistentLevel.Build_Workshop_C_2146449192.FGFactoryLegs",
		c.InstanceName,
	)

	out := filebuffer.New([]byte{})
	p = &parser{
		body: out,
	}

	err = p.serializeComponent(c)
	require.NoError(t, err)

	assert.Equal(t, data, out.Buff.Bytes())
}
