package save

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
	"testing"

	"github.com/l-ross/ficsit-toolkit/save/data"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_Parse(t *testing.T) {
	// Simple E2E test. Parse a save file and then serialize it.

	t.Parallel()

	f, err := os.Open("testdata/example.sav")
	require.NoError(t, err)
	defer f.Close()

	s, err := Parse(f)
	require.NoError(t, err)
	require.NotNil(t, s)

	var b bytes.Buffer
	err = s.Serialize(&b)
	require.NoError(t, err)
}

func TestIdentical(t *testing.T) {
	// Verify that when we serialize the decompressed body of a save file it is
	// identical to the original.

	t.Parallel()

	// Get the decompressed body of a save file.
	f, err := os.Open("testdata/example.sav")
	require.NoError(t, err)
	defer f.Close()

	body, err := DumpBody(f)
	require.NoError(t, err)

	// Store the hash of the original body.
	originalHash := fmt.Sprintf("%x", md5.Sum(body))

	// Parse the body.
	p := &parser{
		Data: data.NewFromBytes(body),
	}

	save := &Save{
		Components:       make(map[string]*Component),
		Entities:         make(map[string]*Entity),
		CollectedObjects: make([]*ObjectReference, 0),

		entityOrder:              make([]string, 0),
		componentOrder:           make([]string, 0),
		objects:                  make([]object, 0),
		componentsByInstanceName: make(map[string]*Component),
	}

	err = p.parseBody(save)
	require.NoError(t, err)

	// Serialize the body
	d2 := data.New()
	s := &serializer{
		Data: d2,
	}

	err = s.serializeBody(save)
	require.NoError(t, err)

	// Verify the original and serialized bodies match.
	newHash := fmt.Sprintf("%x", md5.Sum(d2.Bytes()))
	assert.Equal(t, originalHash, newHash)
}
