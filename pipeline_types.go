package gocd

type Pipeline struct {
	Links                Links                         `json:"_links,omitempty"`
	Group                string                        `json:"group,omitempty"`
	LabelTemplate        string                        `json:"label_template,omitempty"`
	LockBehavior         string                        `json:"lock_behavior,omitempty"`
	Name                 string                        `json:"name,omitempty"`
	Template             string                        `json:"template,omitempty"`
	Origin               *PipelineConfigRepo           `json:"origin,omitempty"`
	Parameters           []PipelineParameter           `json:"parameters,omitempty`
	EnvironmentVariables []PipelineEnvironmentVariable `json:"environment_variables,omitempty"`
	Materials            []PipelineMaterial            `json:"materials,omitempty"`
	Stages               []PipelineStage               `json:"stages,omitempty"`
	TrackingTool         *PipelineTrackingTool         `json:"tracking_tool,omitempty"`
	Timer                *PipelineTimer                `json:"timer,omitempty"`
}

type PipelineConfigRepo struct {
	Links *Links `json:"_links,omitempty"`
	Type  string `json:"type,omitempty"`
	Id    string `json:"id,omitempty"`
}

type PipelineParameter struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// TODO - Validation on Value OR EncryptedValue
type PipelineEnvironmentVariable struct {
	Name           string `json:"name,omitempty"`
	Value          string `json:"value,omitempty"`
	EncryptedValue string `json:"encrypted_value,omitempty"`
	Secure         bool   `json:"secure,omitempty"`
}

type PipelineMaterial struct {
	Type       string                      `json:"type,omitempty"`
	Attributes *PipelineMaterialAttributes `json:"attributes,omitempty"`
}

type PipelineMaterialAttributes struct {
	AutoUpdate          bool                              `json:"auto_update,omitempty"`
	Branch              string                            `json:"branch,omitempty"`
	CheckExternals      bool                              `json:"check_externals,omitempty"`
	Destination         string                            `json:"destination,omitempty"`
	Domain              string                            `json:"domain,omitempty"`
	EncryptedPassword   string                            `json:"encrypted_password,omitempty"`
	Filter              *PipelineMaterialAttributesFilter `json:"filter,omitempty"`
	IgnoreForScheduling bool                              `json:"ignore_for_scheduling,omitempty"`
	InvertFilter        bool                              `json:"invert_filter,omitempty"`
	Name                string                            `json:"name,omitempty"`
	Password            string                            `json:"password,omitempty"`
	Pipeline            string                            `json:"pipeline,omitempty"`
	Port                string                            `json:"port,omitempty"`
	ProjectPath         string                            `json:"project_path,omitempty"`
	Ref                 string                            `json:"ref,omitempty"`
	ShallowClone        bool                              `json:"shallow_clone,omitempty"`
	Stage               string                            `json:"stage,omitempty"`
	SubmoduleFolder     string                            `json:"submodule_folder,omitempty"`
	Url                 string                            `json:"url,omitempty"`
	UseTickets          bool                              `json:"use_tickets,omitempty"`
	Username            string                            `json:"username,omitempty"`
	View                string                            `json:"view,omitempty"`
}

type PipelineMaterialAttributesFilter struct {
	Ignore []string `json:"ignore,omitempty"`
}

type PipelineStage struct {
	Name                  string                        `json:"name,omitempty"`
	FetchMaterial         bool                          `json:"fetch_material,omitempty"`
	CleanWorkingDirectory bool                          `json:"clean_working_directory,omitempty"`
	NeverCleanupArtifacts bool                          `json:"never_cleanup_artifacts,omitempty"`
	Approval              *PipelineStageApproval        `json:"approval,omitempty"`
	EnvironmentVariables  []PipelineEnvironmentVariable `json:"environment_variables,omitempty"`
	Jobs                  []PipelineJob                 `json:"jobs,omitempty"`
}

type PipelineStageApproval struct {
	Type               string                              `json:"type,omitempty"`
	AllowOnlyOnSuccess bool                                `json:"allow_only_on_success,omitempty"`
	Authorization      *PipelineStageApprovalAuthorization `json:"authorization,omitempty"`
}

type PipelineStageApprovalAuthorization struct {
	Users []string `json:"users,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type PipelineJob struct {
	Name                 string                        `json:"name,omitempty"`
	RunInstanceCount     string                        `json:"run_instance_count,omitempty"`
	Timeout              int                           `json:"timeout,omitempty"`
	EnvironmentVariables []PipelineEnvironmentVariable `json:"environment_variables,omitempty"`
	Resources            []string                      `json:"resources,omitempty"`
	Tasks                []PipelineTask                `json:"tasks,omitempty"`
	Tabs                 []PipelineTab                 `json:"tabs,omitempty"`
	Artifacts            []PipelineArtifact            `json:"artifacts,omitempty"`
	ElasticProfileId     string                        `json:elastic_profile_id,omitempty"`
}

type PipelineTask struct {
	Type       string                  `json:"type,omitempty"`
	Attributes *PipelineTaskAttributes `json:"attributes,omitempty"`
}

// TODO - Validation on arguments passed and types
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

// Duplicate of PipelineTask
// Required to avoid recursive struct for OnCancel
type PipelineCancelTask struct {
	Type       string                        `json:"type,omitempty"`
	Attributes *PipelineCancelTaskAttributes `json:"attributes,omitempty"`
}

// Duplicate of PipelineTaskAttributes
// Required to avoid recursive struct for OnCancel
// TODO - Validation on arguments passed and types
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

type PipelineAttributePluginConfiguration struct {
	Id      string `json:"id,omitempty"`
	Version string `json:"version,omitempty"`
}

type PipelineTab struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

type PipelineArtifact struct {
	Type          string                  `json:"type,omitempty"`
	Source        string                  `json:"source,omitempty"`
	Destination   string                  `json:"destination,omitempty"`
	Id            string                  `json:"id,omitempty,omitempty"`
	StoreId       string                  `json:"store_id,omitempty"`
	Configuration []PipelineConfiguration `json:"configuration,omitempty"`
}

type PipelineConfiguration struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type PipelineTrackingTool struct {
	Type       string                          `json:"type,omitempty"`
	Attributes *PipelineTrackingToolAttributes `json:"attributes,omitempty"`
}

type PipelineTrackingToolAttributes struct {
	UrlPattern string `json:"url_pattern,omitempty"`
	Regex      string `json:"regex,omitempty"`
}

type PipelineTimer struct {
	Spec          string `json:"spec,omitempty"`
	OnlyOnChanges bool   `json:"only_on_changes,omitempty"`
}

type PipelineCreateRequest struct {
	Group    string   `json:"group,omitempty"`
	Pipeline Pipeline `json:"pipeline,omitempty"`
}
