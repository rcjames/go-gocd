package gocd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetVersion(t *testing.T) {
	want := "16.6.0"
	c, server := NewMockClient(t)
	defer server.Close()

	version, _ := c.GetVersion()
	got := version.Version

	require.Equal(t, want, got, "Version was unexpected")
}
