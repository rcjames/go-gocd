//go:build unit
// +build unit

package gocd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllArtifactStores(t *testing.T) {
	c, server := NewMockClient(t)
	defer server.Close()

	artifactStores, _ := c.GetAllArtifactStores()

	require.Equal(t, "hub.docker", artifactStores[0].Id, "Unexpected artifact store Id")
	require.Equal(t, "admin", artifactStores[0].Properties[1].Value, "Unexpected value in artifact store property")
}
