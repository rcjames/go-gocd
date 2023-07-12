package gocd

type GetAllArtifactStoresResponse struct {
	Links    Links `json:"_links,omitempty"`
	Embedded struct {
		ArtifactStores []ArtifactStore `json:"artifact_stores"`
	} `json:"_embedded"`
}

type ArtifactStore struct {
	Links      Links                   `json:"_links,omitempty"`
	Id         string                  `json:"id,omitempty"`
	PluginId   string                  `json:"plugin_id,omitempty"`
	Properties []ConfigurationProperty `json:"properties,omitempty"`
}
