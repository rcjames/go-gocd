package gocd

// A Pipeline is a configuration for a [Pipeline].
//
// [Pipeline]: https://api.gocd.org/current/#the-pipeline-config-object
type Pipeline struct {
	// Links are metadata about requests returned by GoCD. This does not need
	// to be provided.
	Links                Links                         `json:"_links,omitempty"`
	Group                string                        `json:"group,omitempty"`
	LabelTemplate        string                        `json:"label_template,omitempty"`
	LockBehavior         string                        `json:"lock_behavior,omitempty"`
	Name                 string                        `json:"name,omitempty"`
	Template             string                        `json:"template,omitempty"`
	Origin               *PipelineConfigRepo           `json:"origin,omitempty"`
	Parameters           []PipelineParameter           `json:"parameters,omitempty"`
	EnvironmentVariables []PipelineEnvironmentVariable `json:"environment_variables,omitempty"`
	Materials            []Material                    `json:"materials,omitempty"`
	Stages               []PipelineStage               `json:"stages,omitempty"`
	TrackingTool         *PipelineTrackingTool         `json:"tracking_tool,omitempty"`
	Timer                *PipelineTimer                `json:"timer,omitempty"`
}

// A PipelineConfigRepo maps to a "[Config repo origin]" object.
//
// [Config repo origin]: https://api.gocd.org/current/#the-generic-tracking-tool-object
type PipelineConfigRepo struct {
	// Links are metadata about requests returned by GoCD. This does not need
	// to be provided.
	Links *Links `json:"_links,omitempty"`
	Type  string `json:"type,omitempty"`
	Id    string `json:"id,omitempty"`
}

// A PipelineParameter maps to a "[Pipeline parameter]" object.
//
// [Pipeline parameter]: https://api.gocd.org/current/#the-pipeline-parameter-object
type PipelineParameter struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// A PipelineEnvironmentVariable maps to a "[Environment variable]" object.
//
// [Environment variable]: https://api.gocd.org/current/#the-environment-variable-object
type PipelineEnvironmentVariable struct {
	Name           string `json:"name,omitempty"`
	Value          string `json:"value,omitempty"`
	EncryptedValue string `json:"encrypted_value,omitempty"`
	Secure         bool   `json:"secure,omitempty"`
}

// A PipelineStage maps to the "[stage]" object.
//
// [stage]: https://api.gocd.org/current/#the-stage-object
type PipelineStage struct {
	Name                  string                        `json:"name,omitempty"`
	FetchMaterial         bool                          `json:"fetch_material,omitempty"`
	CleanWorkingDirectory bool                          `json:"clean_working_directory,omitempty"`
	NeverCleanupArtifacts bool                          `json:"never_cleanup_artifacts,omitempty"`
	Approval              *PipelineStageApproval        `json:"approval,omitempty"`
	EnvironmentVariables  []PipelineEnvironmentVariable `json:"environment_variables,omitempty"`
	Jobs                  []PipelineJob                 `json:"jobs,omitempty"`
}

// A PipelineStageApproval maps to the "[approval]" object.
//
// [approval]: https://api.gocd.org/current/#the-approval-object
type PipelineStageApproval struct {
	Type               string                              `json:"type,omitempty"`
	AllowOnlyOnSuccess bool                                `json:"allow_only_on_success,omitempty"`
	Authorization      *PipelineStageApprovalAuthorization `json:"authorization,omitempty"`
}

// A PipelineStageApprovalAuthorization maps to the "[authorization]" object.
//
// [authorization]: https://api.gocd.org/current/#the-authorization-object
type PipelineStageApprovalAuthorization struct {
	Users []string `json:"users,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

// A PipelineJob maps to the "[job]" object.
//
// [job]: https://api.gocd.org/current/#the-job-object
type PipelineJob struct {
	Name                 string                        `json:"name,omitempty"`
	RunInstanceCount     string                        `json:"run_instance_count,omitempty"`
	Timeout              int                           `json:"timeout,omitempty"`
	EnvironmentVariables []PipelineEnvironmentVariable `json:"environment_variables,omitempty"`
	Resources            []string                      `json:"resources,omitempty"`
	Tasks                []PipelineTask                `json:"tasks,omitempty"`
	Tabs                 []PipelineTab                 `json:"tabs,omitempty"`
	Artifacts            []PipelineArtifact            `json:"artifacts,omitempty"`
	ElasticProfileId     string                        `json:"elastic_profile_id,omitempty"`
}

// A PipelineTask maps to the "[task]" object.
//
// [task]: https://api.gocd.org/current/#the-task-object
type PipelineTask struct {
	Type       string                  `json:"type,omitempty"`
	Attributes *PipelineTaskAttributes `json:"attributes,omitempty"`
}

// A PipelineTaskAttributes maps to the various types of [task attributes]. See
// the documentation for which attributes need configuring for which task types.
//
// [task attributes]: https://api.gocd.org/current/#the-exec-task-attributes
type PipelineTaskAttributes struct {
	Arguments           []string                              `json:"arguments,omitempty"`
	ArtifactId          string                                `json:"artifact_id,omitempty"`
	ArtifactOrigin      string                                `json:"artifact_origin,omitempty"`
	Command             string                                `json:"command,omitempty"`
	Configuration       []PipelineConfiguration               `json:"configuration,omitempty"`
	Destination         string                                `json:"destination,omitempty"`
	IsSourceAFile       bool                                  `json:"is_source_a_file,omitempty"`
	Job                 string                                `json:"job,omitempty"`
	OnCancel            *PipelineCancelTask                   `json:"on_cancel,omitempty"`
	Pipeline            string                                `json:"pipeline,omitempty"`
	PluginConfiguration *PipelineAttributePluginConfiguration `json:"plugin_configuration,omitempty"`
	RunIf               []string                              `json:"run_if,omitempty"`
	Source              string                                `json:"source,omitempty"`
	Stage               string                                `json:"stage,omitempty"`
	WorkingDirectory    string                                `json:"working_directory,omitempty"`
}

// Duplicate of PipelineTask which is required to avoid recursive struct for
// OnCancel.
type PipelineCancelTask struct {
	Type       string                        `json:"type,omitempty"`
	Attributes *PipelineCancelTaskAttributes `json:"attributes,omitempty"`
}

// Duplicate of PipelineTaskAttributes which is required to avoid recursive
// struct for OnCancel.
type PipelineCancelTaskAttributes struct {
	Arguments           []string                              `json:"arguments,omitempty"`
	ArtifactId          string                                `json:"artifact_id,omitempty"`
	ArtifactOrigin      string                                `json:"artifact_origin,omitempty"`
	Command             string                                `json:"command,omitempty"`
	Configuration       []PipelineConfiguration               `json:"configuration,omitempty"`
	Destination         string                                `json:"destination,omitempty"`
	IsSourceAFile       bool                                  `json:"is_source_a_file,omitempty"`
	Job                 string                                `json:"job,omitempty"`
	Pipeline            string                                `json:"pipeline,omitempty"`
	PluginConfiguration *PipelineAttributePluginConfiguration `json:"plugin_configuration,omitempty"`
	RunIf               []string                              `json:"run_if,omitempty"`
	Source              string                                `json:"source,omitempty"`
	Stage               string                                `json:"stage,omitempty"`
	WorkingDirectory    string                                `json:"working_directory,omitempty"`
}

// A PluginAttributePluginConfiguration is used to parse the object required
// by the [pluggable task] object.
//
// [pluggable task]: https://api.gocd.org/current/#the-pluggable-task-object
type PipelineAttributePluginConfiguration struct {
	Id      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
}

// A PipelineTab maps to the [tab] object.
//
// [tab]: https://api.gocd.org/current/#the-tab-object
type PipelineTab struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

// A PipelineArtifact maps to the [pipeline config artifact] object.
//
// [pipeline config artifact]: https://api.gocd.org/current/#the-pipeline-config-artifact-object
type PipelineArtifact struct {
	Type          string                  `json:"type,omitempty"`
	Source        string                  `json:"source,omitempty"`
	Destination   string                  `json:"destination,omitempty"`
	Id            string                  `json:"id,omitempty,omitempty"`
	StoreId       string                  `json:"store_id,omitempty"`
	Configuration []PipelineConfiguration `json:"configuration,omitempty"`
}

// A PipelineConfiguration maps to the configuration field for the [pipeline
// config artifact] object.
//
// [pipeline config artifact]: https://api.gocd.org/current/#the-pipeline-config-artifact-object
type PipelineConfiguration struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// A PipelineTrackingTool maps to the [tracking tool] object.
//
// [tracking tool]: https://api.gocd.org/current/#the-tracking-tool-object
type PipelineTrackingTool struct {
	Type       string                          `json:"type,omitempty"`
	Attributes *PipelineTrackingToolAttributes `json:"attributes,omitempty"`
}

// A PipelineTrackingToolAttributes is used to handle the attributes object of
// the [tracking tool] object.
//
// [tracking tool]: https://api.gocd.org/current/#the-tracking-tool-object
type PipelineTrackingToolAttributes struct {
	UrlPattern string `json:"url_pattern,omitempty"`
	Regex      string `json:"regex,omitempty"`
}

// A PipelineTimer maps to the [timer] object.
//
// [timer]: https://api.gocd.org/current/#the-timer-object
type PipelineTimer struct {
	Spec          string `json:"spec,omitempty"`
	OnlyOnChanges bool   `json:"only_on_changes,omitempty"`
}

// A PipelineCreateRequest is used to wrap the CreatePipeline request object
// to be used by the [Create a pipeline] endpoint.
//
// [Create a pipeline]: https://api.gocd.org/current/#create-a-pipeline
type PipelineCreateRequest struct {
	Group    string   `json:"group,omitempty"`
	Pipeline Pipeline `json:"pipeline,omitempty"`
}
