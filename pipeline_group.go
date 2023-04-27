package gocd

import "fmt"

type GetAllPipelineGroupsResponse struct {
	Links    Links `json:"_links,omitempty"`
	Embedded struct {
		Groups []PipelineGroup `json:"groups"`
	} `json:"_embedded"`
}

type PipelineGroup struct {
	Links         Links                      `json:"_links,omitempty"`
	Name          string                     `json:"name"`
	Authorization PipelineGroupAuthorization `json:"authorization,omitempty"`
	Pipelines     []struct {
		Links Links  `json:"_links,omitempty"`
		Name  string `json:"name"`
	} `json:"pipelines"`
}

type PipelineGroupAuthorization struct {
	View    PipelineGroupAuthorizationsRule `json:"view,omitempty"`
	Operate PipelineGroupAuthorizationsRule `json:"operate,omitempty"`
	Admins  PipelineGroupAuthorizationsRule `json:"admins,omitempty"`
}

type PipelineGroupAuthorizationsRule struct {
	Users []string `json:"users,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

func (c *GoCDClient) GetAllPipelineGroups() ([]PipelineGroup, error) {
	var pipelineGroupsResponse GetAllPipelineGroupsResponse

	_, err := c.getRequest("go/api/admin/pipeline_groups", "", &pipelineGroupsResponse)
	if err != nil {
		return nil, err
	}

	return pipelineGroupsResponse.Embedded.Groups, nil
}

func (c *GoCDClient) GetPipelineGroup(pipelineGroupName string) (PipelineGroup, string, error) {
	var pipelineGroup PipelineGroup
	requestPath := fmt.Sprintf("go/api/admin/pipeline_groups/%s", pipelineGroupName)

	etag, err := c.getRequest(requestPath, "", &pipelineGroup)
	if err != nil {
		return pipelineGroup, "", err
	}

	return pipelineGroup, etag, nil
}

func (c *GoCDClient) CreatePipelineGroup(pipelineGroup PipelineGroup) (PipelineGroup, string, error) {
	var pipelineGroupResponse PipelineGroup

	etag, err := c.postRequest("go/api/admin/pipeline_groups", "", pipelineGroup, &pipelineGroupResponse)
	if err != nil {
		return pipelineGroupResponse, "", err
	}

	return pipelineGroupResponse, etag, nil
}

func (c *GoCDClient) UpdatePipelineGroup(pipelineGroupName, etag string, pipelineGroup PipelineGroup) (PipelineGroup, string, error) {
	var pipelineGroupResponse PipelineGroup
	requestPath := fmt.Sprintf("go/api/admin/pipeline_groups/%s", pipelineGroupName)

	etag, err := c.putRequest(requestPath, "", etag, pipelineGroup, &pipelineGroupResponse)
	if err != nil {
		return pipelineGroupResponse, "", err
	}

	return pipelineGroupResponse, etag, nil
}

func (c *GoCDClient) DeletePipelineGroup(pipelineGroupName string) (string, error) {
	var message DeleteMessage
	requestPath := fmt.Sprintf("go/api/admin/pipeline_groups/%s", pipelineGroupName)

	err := c.deleteRequest(requestPath, "", &message)
	if err != nil {
		return "", err
	}

	return message.Message, nil
}
