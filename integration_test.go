//go:build integration
// +build integration

package gocd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePipelineGroup(t *testing.T) {
	c := New("http://localhost:8153", "", "")

	pipelineGroup := PipelineGroup{
		Name: "first",
	}

	pipelineGroupResponse, _, err := c.CreatePipelineGroup(pipelineGroup)
	require.Empty(t, err, "create pipeline group should not have thrown an error")
	require.Equal(t, pipelineGroup.Name, pipelineGroupResponse.Name, "unexpected pipeline group name on creation")

	pipelineGroups, err := c.GetAllPipelineGroups()
	require.Empty(t, err, "get all pipeline groups should not have thrown an error")
	require.Equal(t, 1, len(pipelineGroups), "unexpected number of pipeline groups")

	// TODO - Test UpdatePipelineGroup (trouble with pipeline changing (etags???))

	msg, err := c.DeletePipelineGroup(pipelineGroup.Name)
	require.Empty(t, err, "delete pipeline group should not have thrown an error")

	want := fmt.Sprintf("Pipeline group with name '%s' was deleted successfully!", pipelineGroup.Name)
	require.Equal(t, want, msg, "unexpected message when deleting pipeline group")

	pipelineGroups, err = c.GetAllPipelineGroups()
	require.Empty(t, err, "get all pipeline groups should not have thrown an error")
	require.Equal(t, 0, len(pipelineGroups), "unexpected number of pipeline groups after delete")
}
