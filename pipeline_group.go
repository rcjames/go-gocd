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

	err := c.getRequest("go/api/admin/pipeline_groups", "", &pipelineGroupsResponse)
	if err != nil {
		fmt.Printf("failed to get all pipelines: %s\n", err)
		return nil, err
	}

	return pipelineGroupsResponse.Embedded.Groups, nil
}

func (c *GoCDClient) GetPipelineGroup(pipelineGroupName string) (PipelineGroup, error) {
	var pipelineGroup PipelineGroup
	requestPath := fmt.Sprintf("go/api/admin/pipeline_groups/%s", pipelineGroupName)

	err := c.getRequest(requestPath, "", &pipelineGroup)
	if err != nil {
		fmt.Printf("failed to get pipeline group %s: %s", pipelineGroupName, err)
		return pipelineGroup, err
	}

	return pipelineGroup, nil
}

func (c *GoCDClient) CreatePipelineGroup(pipelineGroup PipelineGroup) (PipelineGroup, error) {
	var pipelineGroupResponse PipelineGroup

	err := c.postRequest("go/api/admin/pipeline_groups", "", pipelineGroup, &pipelineGroupResponse)
	if err != nil {
		fmt.Printf("failed to create pipeline group %s: %s", pipelineGroup.Name, err)
		return pipelineGroupResponse, err
	}

	return pipelineGroupResponse, nil
}
