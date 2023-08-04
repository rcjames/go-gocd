package gocd

// A GetAllPipelineGroupsResponse object is used for handling the response from
// the "[Get all pipeline groups]" endpoint which contains and _embedded block
// with the PipelineGroups in.
//
// [Get all pipeline groups]: https://api.gocd.org/current/#get-all-pipeline-groups
type GetAllPipelineGroupsResponse struct {
	// Links are metadata about requests returned by GoCD. This does not need
	// to be provided.
	Links Links `json:"_links,omitempty"`
	// Links are metadata about requests returned by GoCD. This does not need
	// to be provided.
	Embedded struct {
		Groups []PipelineGroup `json:"groups"`
	} `json:"_embedded"`
}

// A PipelineGroup is a configuration for a [Pipeline Group].
//
// [Pipeline Group]: https://api.gocd.org/current/#the-pipeline-group-object
type PipelineGroup struct {
	// Links are metadata about requests returned by GoCD. This does not need
	// to be provided.
	Links         Links                      `json:"_links,omitempty"`
	Name          string                     `json:"name"`
	Authorization PipelineGroupAuthorization `json:"authorization,omitempty"`
	Pipelines     []struct {
		// Links are metadata about requests returned by GoCD. This does not need
		// to be provided.
		Links Links  `json:"_links,omitempty"`
		Name  string `json:"name"`
	} `json:"pipelines"`
}

// A PipelineGroupAuthorization contains the [Pipeline group authorization
// settings] for a PipelineGroup.
//
// [Pipeline group authorization settings]: https://api.gocd.org/current/#the-pipeline-group-authorization-configuration
type PipelineGroupAuthorization struct {
	View    PipelineGroupAuthorizationsRule `json:"view,omitempty"`
	Operate PipelineGroupAuthorizationsRule `json:"operate,omitempty"`
	Admins  PipelineGroupAuthorizationsRule `json:"admins,omitempty"`
}

// A PipelineGroupAuthorizationRule contains lists of Users and Roles for use
// in the PipelineGroupAuthorization object.
type PipelineGroupAuthorizationsRule struct {
	Users []string `json:"users,omitempty"`
	Roles []string `json:"roles,omitempty"`
}
