//go:build unit
// +build unit

package gocd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPipeline(t *testing.T) {
	c, server := NewMockClient(t)
	defer server.Close()

	pipeline, _, _ := c.GetPipeline("new_pipeline")

	require.Equal(t, "new_pipeline", pipeline.Name, "Unexpected pipeline name")
	require.Equal(t, "gocd", pipeline.Origin.Type, "Unexpected origin type")
	require.Equal(t, "git@github.com:sample_repo/example.git", pipeline.Materials[0].Attributes.Url, "Unexpected material git repo")
	require.Equal(t, "success", pipeline.Stages[0].Approval.Type, "Unexpected approval type for first stage")
	require.Equal(t, "passed", pipeline.Stages[0].Jobs[0].Tasks[0].Attributes.RunIf[0], "Unexpected first taks run_if parameter")
	require.Equal(t, "gocd/gocd-server", pipeline.Stages[0].Jobs[0].Artifacts[0].Configuration[0].Value, "Unexpected artifact configuration")
	require.Equal(t, "j2", pipeline.Stages[1].Jobs[0].Name, "Unexpected name for the first job of the second stage")
}
