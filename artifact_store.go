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

func (c *GoCDClient) CreateArtifactStore(as ArtifactStore) (ArtifactStore, string, error) {
	var artifactStoreResponse ArtifactStore

	etag, err := c.postRequest("go/api/admin/artifact_stores", "", as, &artifactStoreResponse)
	if err != nil {
		return artifactStoreResponse, "", err
	}

	return artifactStoreResponse, etag, nil
}

func (c *GoCDClient) UpdateArtifactStore(id, etag string, as ArtifactStore) (ArtifactStore, string, error) {
	var artifactStore ArtifactStore
	requestPath := fmt.Sprintf("go/api/admin/artifact_stores/%s", id)

	etag, err := c.putRequest(requestPath, "", etag, as, &artifactStore)
	if err != nil {
		return artifactStore, "", err
	}

	return artifactStore, etag, err
}

func (c *GoCDClient) DeleteArtifactStore(id string) (string, error) {
	var message DeleteMessage
	requestPath := fmt.Sprintf("go/api/admin/artifact_stores/%s", id)

	err := c.deleteRequest(requestPath, "", &message)
	if err != nil {
		return "", err
	}

	return message.Message, nil
}

func (as *ArtifactStore) AddProperty(cp ConfigurationProperty) {
	as.Properties = append(as.Properties, cp)
}

func (as *ArtifactStore) GetPropertyValue(key string) string {
	for _, p := range as.Properties {
		if p.Key == key {
			return p.Value
		}
	}

	return ""
}
