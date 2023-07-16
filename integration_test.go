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

	pipelineGroupResponse, etag, err := c.CreatePipelineGroup(pipelineGroup)
	require.Empty(t, err, "create pipeline group should not have thrown an error")
	require.Equal(t, pipelineGroup.Name, pipelineGroupResponse.Name, "unexpected pipeline group name on creation")

	pipelineGroupName = "second"
	pipelineGroupResponse.Name = pipelineGroupName
	pipelineGroupResponse2, _, err := c.UpdatePipelineGroup(pipelineGroup.Name, etag, pipelineGroupResponse)
	require.Empty(t, err, "update pipeline group should not have thrown an error")
	require.Equal(t, pipelineGroupName, pipelineGroupResponse2.Name, "unexpected pipeline group name on update")

	pipelineGroups, err := c.GetAllPipelineGroups()
	require.Empty(t, err, "get all pipeline groups should not have thrown an error")
	require.Equal(t, 1, len(pipelineGroups), "unexpected number of pipeline groups")

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

	msg, err = c.DeletePipelineGroup(pipelineGroupName)
	require.Empty(t, err, "delete pipeline group should not have thrown an error")
	want = fmt.Sprintf("Pipeline group with name '%s' was deleted successfully!", pipelineGroupName)
	require.Equal(t, want, msg, "unexpected message when deleting pipeline group")

	pipelineGroups, err = c.GetAllPipelineGroups()
	require.Empty(t, err, "get all pipeline groups should not have thrown an error")
	require.Equal(t, 0, len(pipelineGroups), "unexpected number of pipeline groups after delete")
}

func TestCreateArtifactStore(t *testing.T) {
	c := New("http://localhost:8153", "", "")

	artifactStoreId := "docker"
	artifactStore := ArtifactStore{
		Id:       artifactStoreId,
		PluginId: "cd.go.artifact.docker.registry",
	}

	var properties = make(map[string]string)
	properties["RegistryURL"] = "https://your_docker_registry_url"
	properties["Username"] = "admin"
	properties["Password"] = "badger"
	properties["RegistryType"] = "other"
	for k, v := range properties {
		p := ConfigurationProperty{
			Key:   k,
			Value: v,
		}
		artifactStore.AddProperty(p)
	}

	artifactStoreResponse, etag, err := c.CreateArtifactStore(artifactStore)
	require.Empty(t, err, "create artifact store should not have thrown an error")
	require.Equal(t, artifactStoreId, artifactStoreResponse.Id, "unexpected artifact store id")

	artifactStores, err := c.GetAllArtifactStores()
	require.Empty(t, err, "get all artifact stores should not have thrown an error")
	require.Equal(t, 1, len(artifactStores), "unexpected number of artifact stores")

	var usernameKey int
	for i, p := range artifactStoreResponse.Properties {
		if p.Key == "Username" {
			usernameKey = i
		}
	}
	artifactStoreResponse.Properties[usernameKey].Value = "root"
	artifactStoreResponse2, _, err := c.UpdateArtifactStore(artifactStoreId, etag, artifactStoreResponse)
	require.Empty(t, err, "update artifact store should not have thrown an error")
	require.Equal(t, "root", artifactStoreResponse2.GetPropertyValue("Username"), "unexpected registry password")

	msg, err := c.DeleteArtifactStore(artifactStoreId)
	require.Empty(t, err, "delete artifact store should not have thrown an error")
	want := fmt.Sprintf("Artifact store with id '%s' was deleted successfully!", artifactStoreId)
	require.Equal(t, want, msg, "unexpected message when deleting artifact store")
}
