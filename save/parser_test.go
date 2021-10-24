package save

import (
	"encoding/binary"
	"fmt"
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

	fOut, err := os.OpenFile("dump.dat", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	require.NoError(t, err)
	defer fOut.Close()

	err = Serialize(s, fOut)
	require.NoError(t, err)

	//for _, e := range s.Entities {
	//	out := slicewriteseek.New()
	//	p := &parser{
	//		body: out,
	//	}
	//
	//	err = e.serializeData(p)
	//	require.NoError(t, err)
	//
	//	assert.Equal(t, e.objectData, out.Buffer)
	//
	//	if t.Failed() {
	//		fmt.Println(e.objectDataPos)
	//		fmt.Println(e.TypePath)
	//
	//		err = ioutil.WriteFile("entity_in.dump", e.objectData, 0644)
	//		require.NoError(t, err)
	//
	//		err = ioutil.WriteFile("entity_out.dump", out.Buffer, 0644)
	//		require.NoError(t, err)
	//
	//		t.FailNow()
	//	}
	//}

	//for _, c := range s.Components {
	//	out := slicewriteseek.New()
	//	p := &parser{
	//		body: out,
	//	}
	//
	//	err = c.serializeData(p)
	//	require.NoError(t, err)
	//
	//	assert.Equal(t, c.objectData, out.Buffer)
	//
	//	if t.Failed() {
	//		fmt.Println(c.objectDataPos)
	//		fmt.Println(c.TypePath)
	//
	//		err = ioutil.WriteFile("component_in.dump", c.objectData, 0644)
	//		require.NoError(t, err)
	//
	//		err = ioutil.WriteFile("component_out.dump", out.Buffer, 0644)
	//		require.NoError(t, err)
	//
	//		t.FailNow()
	//	}
	//}
}

func TestInt32(t *testing.T) {
	d := make([]byte, 4)
	binary.LittleEndian.PutUint32(d, 10000)
	fmt.Printf("%v\n", d)
}
