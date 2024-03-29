package gocd

// A GetAllArtifactStoreResponse object is used for handling the response from
// the "[Get all artifact stores]" endpoint which contains an _embedded block
// with the ArtifactStores in.
//
// [Get all artifact stores]: https://api.gocd.org/current/#get-all-artifact-stores
type GetAllArtifactStoresResponse struct {
	// Links are metadata about requests returned by GoCD. This does not need
	// to be provided.
	Links Links `json:"_links,omitempty"`
	// Embedded is the object containing the response objects
	Embedded struct {
		ArtifactStores []ArtifactStore `json:"artifact_stores"`
	} `json:"_embedded"`
}

// A ArtifactStore is a configuration for an [Artifact Store]. See [Plugin guide]
// for further information.
//
// [Artifact Store]: https://api.gocd.org/current/#the-artifact-store-object
// [Plugin guide]: https://docs.gocd.org/current/extension_points/plugin_user_guide.html
type ArtifactStore struct {
	// Links are metadata about requests returned by GoCD. This does not need
	// to be provided.
	Links      Links                   `json:"_links,omitempty"`
	Id         string                  `json:"id,omitempty"`
	PluginId   string                  `json:"plugin_id,omitempty"`
	Properties []ConfigurationProperty `json:"properties,omitempty"`
}
