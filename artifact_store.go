package gocd

import "fmt"

func (c *GoCDClient) GetAllArtifactStores() ([]ArtifactStore, error) {
	var artifactStores GetAllArtifactStoresResponse

	_, err := c.getRequest("go/api/admin/artifact_stores", "", &artifactStores)
	if err != nil {
		return nil, err
	}

	return artifactStores.Embedded.ArtifactStores, err
}

func (c *GoCDClient) GetArtifactStore(id string) (ArtifactStore, string, error) {
	var artifactStore ArtifactStore
	requestPath := fmt.Sprintf("go/api/admin/artifact_stores/%s", id)

	etag, err := c.getRequest(requestPath, "", &artifactStore)
	if err != nil {
		return artifactStore, "", err
	}

	return artifactStore, etag, err

}

/*
func (c *GoCDClient) CreateArtifactStore(as ArtifactStore) (ArtifactStore, string, error) {
}

func (c *GoCDClient) UpdateArtifactStore(id, etag string, as ArtifactStore) (ArtifactStore, string, error) {

}

func (c *GoCDClient) DeleteArtifactStore(id string) (string, error) {

}
*/
