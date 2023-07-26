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

func TestGetMaterialModifications(t *testing.T) {
	c, server := NewMockClient(t)
	defer server.Close()

	materialModifications, err := c.GetMaterialModifications("03c8d2a131154436b6ef2d621c6f773837481aaea8c5c1bb9c0cb9b5bc64a2f1")

	require.Empty(t, err, "get material modifications should not have returned an error")
	require.Equal(t, "my hola mundo changes", materialModifications.Embedded.Modifications[0].Comment, "get unexpected comment on material modification")
}
