package gocd

// A Material is a configuration for an input to a pipeline which causes the
// pipeline to trigger ([see docs]). Different MaterialAttributes are required
// based on the Material type. See the [material attributes] section of the API
// documentation.
//
// [see docs]: https://docs.gocd.org/current/introduction/concepts_in_go.html#materials
// [material attributes]: https://api.gocd.org/current/#the-material-object
type Material struct {
	// Type indicates the type of material to be configured. See API documentation
	// for further information.
	Type string `json:"type,omitempty"`
	// The fingerprint of the material
	Fingerprint string `json:"fingerprint,omitempty"`
	// The attributes for the material. See MaterialAttributes.
	Attributes *MaterialAttributes `json:"attributes,omitempty"`
}

// MaterialAttributes contains all the different types of configuration for the
// different material types. See [API documentation] for further detail on which
// fields are required for which types.
//
// [API documentation]: https://api.gocd.org/current/#the-material-object
type MaterialAttributes struct {
	// AutoUpdate - Whether to poll for new changes or not.
	AutoUpdate bool `json:"auto_update,omitempty"`
	// Branch - The git/mercurial branch to build.
	Branch string `json:"branch,omitempty"`
	// CheckExternals - Whether the changes o the externals will trigger the pipeline automatically or not.
	CheckExternals bool `json:"check_externals,omitempty"`
	// Destination - The directory (relative to the pipeline directory) in which source code will be checked out.
	Destination string `json:"destination,omitempty"`
	// Domain - The domain name for TFS authentication credentials.
	Domain string `json:"domain,omitempty"`
	// EncryptedPassword - The encrypted password for source providers.
	EncryptedPassword string `json:"encrypted_password,omitempty"`
	// Filter - The filter specifies files in changesets that should not trigger a pipeline automatically.
	Filter *MaterialAttributesFilter `json:"filter,omitempty"`
	// IgnoreForScheduling - Whether the pipeline should be triggered when there are changes in this material.
	IgnoreForScheduling bool `json:"ignore_for_scheduling,omitempty"`
	// InvertFilter - Invert filter to enable whitelist.
	InvertFilter bool `json:"invert_filter,omitempty"`
	// Name - The name of this material.
	Name string `json:"name,omitempty"`
	// Password - The password for the source provider.
	Password string `json:"password,omitempty"`
	// Pipeline - The name of a pipeline that this pipeline depends on.
	Pipeline string `json:"pipeline,omitempty"`
	// Port - Perforce server connection to use ([transport:]host:port).
	Port string `json:"port,omitempty"`
	// ProjectPath - The project path within the TFS collection.
	ProjectPath string `json:"project_path,omitempty"`
	// Ref - The unique package repository id.
	Ref string `json:"ref,omitempty"`
	// ShallowClone - Clone with -n (–depth) option if set to true.
	ShallowClone bool `json:"shallow_clone,omitempty"`
	// Stage - The name of a stage which will trigger this pipeline once it is successful.
	Stage string `json:"stage,omitempty"`
	// SubmoduleFolder - The git submodule in the git repository.
	SubmoduleFolder string `json:"submodule_folder,omitempty"`
	// Url - The URL of the source repository.
	Url string `json:"url,omitempty"`
	// UseTickets - Whether to work with the Perforce tickets or not.
	UseTickets bool `json:"use_tickets,omitempty"`
	// Username - The username to be used.
	Username string `json:"username,omitempty"`
	// View - The Perforce view.
	View string `json:"view,omitempty"`
}

// MaterialAttributesFilter contains a list of file paths as strings for the
// Materials Filter attribute. See [API documentation] for more information.
//
// [API documentation]: https://api.gocd.org/current/#the-filter-object
type MaterialAttributesFilter struct {
	// Ignore is a list of file paths to ignore or match depending on InvertFilter.
	Ignore []string `json:"ignore,omitempty"`
}

// MaterialModification contains a change made to a Material from the [Get
// material modifications] API endpoint.
//
// [Get material modifications]: https://api.gocd.org/current/#get-material-modifications
type MaterialModification struct {
	// The EmailAddress of the user making the change
	EmailAddress string `json:"email_address,omitempty"`
	// The Id for the Material change
	Id int `json:"id"`
	// The ModificationTime of the change as a unix timestamp
	ModificationTime int `json:"modification_time"`
	// The GoCD UserName for the user who made the change
	UserName string `json:"user_name,omitempty"`
	// The Comment left for the change.
	Comment string `json:"comment,omitempty"`
	// The Revision hash for the change
	Revision string `json:"revision"`
}

// A GetAllMaterialsResponse object is used for handling the response from the
// "[Get all material]" endpoint which contains an _embedded block with the
// Materials in.
//
// [Get all materials]: https://api.gocd.org/current/#get-all-materials
type GetAllMaterialsResponse struct {
	// Links are metadata about requests returned by GoCD. This does not need
	// to be provided.
	Links Links `json:"_links,omitempty"`
	// Embedded is the object containing the response objects
	Embedded struct {
		// Materials is the list of Materials
		Materials []Material `json:"materials"`
	} `json:"_embedded"`
}

// A GetMaterialModificationsResponse object is used for handling the response
// from the "[Get material modifications]" endpoint which contains an _embedded
// block with the MeaterialModifications and Pagination information in.
//
// [Get material modifications]: https://api.gocd.org/current/#get-material-modifications
type GetMaterialModificationsResponse struct {
	// Links are metadata about requests returned by GoCD. This does not need
	// to be provided.
	Links Links `json:"_links,omitempty"`
	// Embedded is the object containing the response objects
	Embedded struct {
		// Modifications is the list of MaterialModifications
		Modifications []MaterialModification `json:"modifications,omitempty"`
		// Pagination contains the pagination information for the response
		Pagination Pagination `json:"pagination,omitempty"`
	} `json:"_embedded"`
}
