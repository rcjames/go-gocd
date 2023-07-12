package gocd

func (c *GoCDClient) GetAllArtifactStores() ([]ArtifactStore, error) {
	var artifactStores GetAllArtifactStoresResponse

	_, err := c.getRequest("go/api/admin/artifact_stores", "", &artifactStores)
	if err != nil {
		return nil, err
	}

	return artifactStores.Embedded.ArtifactStores, err
}
