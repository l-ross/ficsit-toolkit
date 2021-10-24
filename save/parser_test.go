package save

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/ViRb3/slicewriteseek"

	"github.com/stretchr/testify/require"
)

// Simple E2E test.
func TestParser_Parse(t *testing.T) {
	f, err := os.Open("../exp/Example.sav")
	require.NoError(t, err)
	defer f.Close()

	s, err := Parse(f)
	require.NoError(t, err)
	require.NotNil(t, s)

	fOut, err := os.OpenFile("out.dump", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	require.NoError(t, err)

	err = Serialize(s, fOut)
	require.NoError(t, err)

	fOut.Close()

	data, err := ioutil.ReadFile("out.dump")
	require.NoError(t, err)

	p := &parser{
		body: &slicewriteseek.SliceWriteSeeker{
			Buffer: data,
		},
	}
	s2 := &Save{}
	err = p.parseBody(s2)
	require.NoError(t, err)
}
