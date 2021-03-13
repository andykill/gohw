package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	err := Copy("/dev/null", "/dev/null", 0, 0)
	require.NoError(t, err)
	err = Copy("/dev/null", "/dev/null", 10, 0)
	require.Error(t, err, ErrOffsetExceedsFileSize)
	err = Copy("/tmp/file_not_found", "/dev/null", 0, 0)
	require.Error(t, err, ErrUnsupportedFile)
}
