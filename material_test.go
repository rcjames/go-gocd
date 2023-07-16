//go:build unit
// +build unit

package gocd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllMaterials(t *testing.T) {
	c, server := NewMockClient(t)
	defer server.Close()

	materials, _ := c.GetAllMaterials()

	require.Equal(t, "git", materials[0].Type, "unexpected type on git material")
	require.Equal(t, "master", materials[0].Attributes.Branch, "unexpected branch on git material")
	require.Equal(t, "up42_stage", materials[3].Attributes.Stage, "unexpected stage on dependency material")
}
