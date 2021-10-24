package save

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Simple E2E test.
func TestParser_Parse(t *testing.T) {
	debug = true

	f, err := os.Open("../exp/Example.sav")
	require.NoError(t, err)
	defer f.Close()

	s, err := Parse(f)
	require.NoError(t, err)
	require.NotNil(t, s)
}
