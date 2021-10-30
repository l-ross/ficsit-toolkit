package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenStruct(t *testing.T) {
	t.Parallel()

	ty, v, err := genStruct("(B=1,G=2,R=3,A=4)", false)
	require.NoError(t, err)

	fmt.Println(ty)
	fmt.Println()
	fmt.Println(v)
}
