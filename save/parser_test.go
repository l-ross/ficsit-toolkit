package save

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// Simple E2E test.
func TestParser_Parse(t *testing.T) {
	f, err := os.Open("../exp/Example.sav")
	require.NoError(t, err)
	defer f.Close()

	p, err := NewParser(f)
	require.NoError(t, err)

	s, err := p.Parse()
	require.NoError(t, err)
	require.NotNil(t, s)

	out, err := json.MarshalIndent(s, "", "  ")
	require.NoError(t, err)

	err = ioutil.WriteFile("out.json", out, 0644)
	require.NoError(t, err)
}
