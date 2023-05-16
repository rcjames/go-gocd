//go:build integration
// +build integration

package gocd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePipeline(t *testing.T) {
	c := New("http://localhost:8153", "", "")

	pipelineGroupName := "first"

	pipelineGroup := PipelineGroup{
		Name: pipelineGroupName,
	}

	pipelineGroupResponse, _, err := c.CreatePipelineGroup(pipelineGroup)
	require.Empty(t, err, "create pipeline group should not have thrown an error")
	require.Equal(t, pipelineGroup.Name, pipelineGroupResponse.Name, "unexpected pipeline group name on creation")

	pipelineGroups, err := c.GetAllPipelineGroups()
	require.Empty(t, err, "get all pipeline groups should not have thrown an error")
	require.Equal(t, 1, len(pipelineGroups), "unexpected number of pipeline groups")

	// TODO - Test UpdatePipelineGroup (trouble with pipeline changing (etags???))
	pipelineName := "pipeline"

	pipelineTask := PipelineTask{
		Type: "exec",
		Attributes: &PipelineTaskAttributes{
			Command:   "echo",
			Arguments: []string{"1"},
		},
	}

	pipelineJob := PipelineJob{
		Name: "job1",
	}
	pipelineJob.AddTask(pipelineTask)

	pipelineStage := PipelineStage{
		Name: "stage1",
	}
	pipelineStage.AddJob(pipelineJob)

	materialBranch := "main"
	pipelineMaterial := PipelineMaterial{
		Type: "git",
		Attributes: &PipelineMaterialAttributes{
			Name:         "go-gocd",
			Url:          "https://github.com/rcjames/go-gocd",
			Branch:       materialBranch,
			ShallowClone: true,
		},
	}

	pipeline := Pipeline{
		Name:  pipelineName,
		Group: pipelineGroupName,
	}
	pipeline.AddStage(pipelineStage)
	pipeline.AddMaterial(pipelineMaterial)

	pipelineResponse, pipelineEtag, err := c.CreatePipeline(pipeline)
	require.Empty(t, err, "create pipeline should not have thrown an error")
	require.Equal(t, pipelineName, pipelineResponse.Name, "unexpected pipeline name")
	require.Equal(t, materialBranch, pipelineResponse.Materials[0].Attributes.Branch, "unexpected meterial branch")

	updatedPipelineStageName := "stageA"
	pipelineResponse.Stages[0].Name = updatedPipelineStageName
	pipelineResponse2, _, err := c.UpdatePipeline(pipelineResponse.Name, pipelineEtag, pipelineResponse)
	require.Empty(t, err, "update pipeline should not have thrown an error")
	require.Equal(t, updatedPipelineStageName, pipelineResponse2.Stages[0].Name, "updated pipeline stage name is not expected")

	msg, err := c.DeletePipeline(pipelineName)
	require.Empty(t, err, "delete pipeline should not have thown an error")
	want := fmt.Sprintf("Pipeline with name '%s' was deleted successfully!", pipelineName)
	require.Equal(t, want, msg, "unexpected message when deleting pipeline")

	msg, err = c.DeletePipelineGroup(pipelineGroup.Name)
	require.Empty(t, err, "delete pipeline group should not have thrown an error")

	want = fmt.Sprintf("Pipeline group with name '%s' was deleted successfully!", pipelineGroup.Name)
	require.Equal(t, want, msg, "unexpected message when deleting pipeline group")

	pipelineGroups, err = c.GetAllPipelineGroups()
	require.Empty(t, err, "get all pipeline groups should not have thrown an error")
	require.Equal(t, 0, len(pipelineGroups), "unexpected number of pipeline groups after delete")
}
