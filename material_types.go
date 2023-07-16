package gocd

type Material struct {
	Type        string              `json:"type,omitempty"`
	Fingerprint string              `json:"fingerprint,omitempty"`
	Attributes  *MaterialAttributes `json:"attributes,omitempty"`
}

type MaterialAttributes struct {
	AutoUpdate          bool                      `json:"auto_update,omitempty"`
	Branch              string                    `json:"branch,omitempty"`
	CheckExternals      bool                      `json:"check_externals,omitempty"`
	Destination         string                    `json:"destination,omitempty"`
	Domain              string                    `json:"domain,omitempty"`
	EncryptedPassword   string                    `json:"encrypted_password,omitempty"`
	Filter              *MaterialAttributesFilter `json:"filter,omitempty"`
	IgnoreForScheduling bool                      `json:"ignore_for_scheduling,omitempty"`
	InvertFilter        bool                      `json:"invert_filter,omitempty"`
	Name                string                    `json:"name,omitempty"`
	Password            string                    `json:"password,omitempty"`
	Pipeline            string                    `json:"pipeline,omitempty"`
	Port                string                    `json:"port,omitempty"`
	ProjectPath         string                    `json:"project_path,omitempty"`
	Ref                 string                    `json:"ref,omitempty"`
	ShallowClone        bool                      `json:"shallow_clone,omitempty"`
	Stage               string                    `json:"stage,omitempty"`
	SubmoduleFolder     string                    `json:"submodule_folder,omitempty"`
	Url                 string                    `json:"url,omitempty"`
	UseTickets          bool                      `json:"use_tickets,omitempty"`
	Username            string                    `json:"username,omitempty"`
	View                string                    `json:"view,omitempty"`
}

type MaterialAttributesFilter struct {
	Ignore []string `json:"ignore,omitempty"`
}

type GetAllMaterialsResponse struct {
	Links    Links `json:"_links,omitempty"`
	Embedded struct {
		Materials []Material `json:"materials"`
	} `json:"_embedded"`
}
