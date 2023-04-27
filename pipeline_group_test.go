//go:build unit
// +build unit

package gocd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllPipelineGroups(t *testing.T) {
	c, server := NewMockClient(t)
	defer server.Close()

	pipelineGroups, _ := c.GetAllPipelineGroups()

	require.Equal(t, "first", pipelineGroups[0].Name, "Unexpected pipeline group name")
	require.Equal(t, "up42", pipelineGroups[0].Pipelines[0].Name, "Unexpected pipeline name")
	require.Equal(t, "https://ci.example.com/go/api/admin/pipeline_groups/first", pipelineGroups[0].Links.Self, "Unexpected pipeline group link ref")
	require.Equal(t, "https://ci.example.com/go/api/admin/pipelines/up42", pipelineGroups[0].Pipelines[0].Links.Self, "Unexpected pipeline link ref")
}

func TestGetPipelineGroup(t *testing.T) {
	c, server := NewMockClient(t)
	defer server.Close()

	pipelineGroup, _, _ := c.GetPipelineGroup("first")

	require.Equal(t, "first", pipelineGroup.Name, "Unexpected pipeline group name")
	require.Equal(t, "up42", pipelineGroup.Pipelines[0].Name, "Unexpected pipeline name")
	require.Equal(t, "https://ci.example.com/go/api/admin/pipelines/up42", pipelineGroup.Pipelines[0].Links.Self, "Unexpected pipeline link ref")
}
