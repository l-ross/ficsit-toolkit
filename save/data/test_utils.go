package data

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestData(t *testing.T, filePath string) *Data {
	f, err := os.Open(filePath)
	require.NoError(t, err)
	t.Cleanup(func() {
		f.Close()
	})

	d, err := NewFromReader(f)
	require.NoError(t, err)

	return d
}

func AssertAllBufferRead(t *testing.T, d *Data) {
	assert.Equal(t,
		d.Index(),
		int64(len(d.Bytes())),
		"we should have consumed the entire reader",
	)
}
