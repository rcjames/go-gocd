package gocd

import "fmt"

// GetAllPipelineGroups gets a list of all PipelineGroups via the "[Get all
// pipeline groups]" endpoint.
//
// Example usage:
//
//	 c := gocd.New(hostname, username, password)
//	 pipelineGroups, _ := c.GetAllPipelineGroups
//	 for _, pg := range pipelineGroups {
//		fmt.Println(pg.Name)
//	 }
//
// [Get all pipeline groups]: https://api.gocd.org/current/#get-all-pipeline-groups
func (c *GoCDClient) GetAllPipelineGroups() ([]PipelineGroup, error) {
	var pipelineGroupsResponse GetAllPipelineGroupsResponse

	_, err := c.getRequest("go/api/admin/pipeline_groups", "", &pipelineGroupsResponse)
	if err != nil {
		return nil, err
	}

	return pipelineGroupsResponse.Embedded.Groups, nil
}

// GetPipelineGroup returns the configuration for a specific pipeline group via
// the "[Get a pipeline group]" endpoint and returns the pipeline group and the
// ETAG.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	pg, _, _ := c.GetPipelineGroup("group1")
//	for _, p := range pg.Pipelines {
//		fmt.Println(p.Name)
//	}
//
// [Get a pipeline group]: https://api.gocd.org/current/#get-a-pipeline-group
func (c *GoCDClient) GetPipelineGroup(pipelineGroupName string) (PipelineGroup, string, error) {
	var pipelineGroup PipelineGroup
	requestPath := fmt.Sprintf("go/api/admin/pipeline_groups/%s", pipelineGroupName)

	etag, err := c.getRequest(requestPath, "", &pipelineGroup)
	if err != nil {
		return pipelineGroup, "", err
	}

	return pipelineGroup, etag, nil
}

// CreatePipelineGroup creates a PipelineGroup via the "[Create a pipeline group]"
// endpoint and returns the created pipeline group and ETAG.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	c.CreatePipeline(PipelineGroup{
//		Name: "group1",
//	})
//
// [Create a pipeline group]: https://api.gocd.org/current/#create-a-pipeline-group
func (c *GoCDClient) CreatePipelineGroup(pipelineGroup PipelineGroup) (PipelineGroup, string, error) {
	var pipelineGroupResponse PipelineGroup

	etag, err := c.postRequest("go/api/admin/pipeline_groups", "", pipelineGroup, &pipelineGroupResponse)
	if err != nil {
		return pipelineGroupResponse, "", err
	}

	return pipelineGroupResponse, etag, nil
}

// UpdatePipelineGroup updates the provided pipeline group name to match the
// PipelineGroup config via the "[Update a pipeline group]" endpoint. ETAG must
// be up to date with the current config for GoCD to allow an update.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//
//	pipelineGroupName := "group1"
//	pg, etag, _ := gocd.GetPipelineGroup(pipelineGroupName)
//	pg.Name = "group2"
//	c.UpdatePipelineGroup(pipelineGroupName, etag, pg)
//
// [Update a pipeline group]: https://api.gocd.org/current/#update-a-pipeline-group
func (c *GoCDClient) UpdatePipelineGroup(pipelineGroupName, etag string, pipelineGroup PipelineGroup) (PipelineGroup, string, error) {
	var pipelineGroupResponse PipelineGroup
	requestPath := fmt.Sprintf("go/api/admin/pipeline_groups/%s", pipelineGroupName)

	etag, err := c.putRequest(requestPath, "", etag, pipelineGroup, &pipelineGroupResponse)
	if err != nil {
		return pipelineGroupResponse, "", err
	}

	return pipelineGroupResponse, etag, nil
}

// DeletePipelineGroup deletes the pipeline group using the "[Delete a pipeline
// group]" endpoint and returns the DeleteMessage.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	msg, _ := c.DeletePipelineGroup("group1")
//	fmt.Println(msg)
//
// [Delete a pipeline group]: https://api.gocd.org/current/#delete-a-pipeline-group
func (c *GoCDClient) DeletePipelineGroup(pipelineGroupName string) (string, error) {
	var message DeleteMessage
	requestPath := fmt.Sprintf("go/api/admin/pipeline_groups/%s", pipelineGroupName)

	err := c.deleteRequest(requestPath, "", &message)
	if err != nil {
		return "", err
	}

	return message.Message, nil
}
