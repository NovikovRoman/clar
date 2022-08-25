package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_toSnake(t *testing.T) {
	data := map[string]string{
		"DdsfsAS23Sd_sad": "ddsfs_a_s23_sd_sad",
		"toSnake":         "to_snake",
		"HeLlo World":     "he_llo _world",
	}

	for s, r := range data {
		require.Equal(t, toSnake(s), r)
	}
}
