package gocd

import "fmt"

// GetAllArtifactStores gets a list of all available artifact stores via the
// "[Get all artifact stores]" API.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	artifactStores, _ := c.GetAllArtifactStores()
//	for _, a := range artifactStores {
//	  fmt.Println(a.Id)
//	}
//
// [Get all artifact stores]: https://api.gocd.org/current/#get-all-artifact-stores
func (c *GoCDClient) GetAllArtifactStores() ([]ArtifactStore, error) {
	var artifactStores GetAllArtifactStoresResponse

	_, err := c.getRequest("go/api/admin/artifact_stores", "", &artifactStores)
	if err != nil {
		return nil, err
	}

	return artifactStores.Embedded.ArtifactStores, err
}

// GetArtifactStore fetches the artifact store and ETAG for a given ID, by
// calling the "[Get an artifact store]" API.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	artifactStore, etag, _ := c.GetArtifactStore("docker")
//	fmt.Printf("PluginId: %s, ETAG: %s", artifactStore.PluginId, etag)
//
// [Get an artifact store]: https://api.gocd.org/current/#get-an-artifact-store
func (c *GoCDClient) GetArtifactStore(id string) (ArtifactStore, string, error) {
	var artifactStore ArtifactStore
	requestPath := fmt.Sprintf("go/api/admin/artifact_stores/%s", id)

	etag, err := c.getRequest(requestPath, "", &artifactStore)
	if err != nil {
		return artifactStore, "", err
	}

	return artifactStore, etag, err

}

// CreateArtifactStore creates an artifact store by calling the "[Create an artifact store]"
// API, returning the created ArtifactStore object and the ETAG.
//
// Example usage:
//
//	 c := gocd.New(hostname, username, password)
//	 artifactStore := ArtifactStore{
//		  Id:       "docker",
//		  PluginId: "cd.go.artifact.docker.registry",
//	 }
//
//	 var properties = make(map[string]string)
//	 properties["RegistryURL"] = "https://your_docker_registry_url"
//	 properties["Username"] = "admin"
//	 properties["Password"] = "badger"
//	 properties["RegistryType"] = "other"
//	 for k, v := range properties {
//	   p := ConfigurationProperty{
//	     Key:   k,
//	     Value: v,
//	   }
//	   artifactStore.AddProperty(p)
//	 }
//
//	 artifactStoreResponse, etag, _ := c.CreateArtifactStore(artifactStore)
//	 fmt.Printf("Plugin Id: %s, ETAG: %s", artifactStoreResponse.PluginId, etag)
//
// [Create an artifact store]: https://api.gocd.org/current/#create-an-artifact-store
func (c *GoCDClient) CreateArtifactStore(as ArtifactStore) (ArtifactStore, string, error) {
	var artifactStoreResponse ArtifactStore

	etag, err := c.postRequest("go/api/admin/artifact_stores", "", as, &artifactStoreResponse)
	if err != nil {
		return artifactStoreResponse, "", err
	}

	return artifactStoreResponse, etag, nil
}

// UpdateArtifaceStore updates the configuration for an existing artifact store
// by calling the "[Update an artifact store]" API and returning the new
// configuration and teh ETAG.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	artifactStoreResponse, etag, _ := c.GetArtifactStore("docker")
//	artifactStoreReponse.Properties[0].Value = "updated"
//	updatedArtifactStoreResponse, _, _ := c.UpdateArtifactStore("docker", etag, artifactStoreResponse)
//
// [Update an artifact store]: https://api.gocd.org/current/#update-an-artifact-store
func (c *GoCDClient) UpdateArtifactStore(id, etag string, as ArtifactStore) (ArtifactStore, string, error) {
	var artifactStore ArtifactStore
	requestPath := fmt.Sprintf("go/api/admin/artifact_stores/%s", id)

	etag, err := c.putRequest(requestPath, "", etag, as, &artifactStore)
	if err != nil {
		return artifactStore, "", err
	}

	return artifactStore, etag, err
}

// DeleteArtifactStore deletes the artifact store for the provided Id using the
// "[Delete an artifact store]" API, returning a message on the success of the
// deletion.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	msg, _ := c.DeleteArtifactStore("docker")
//	fmt.Printf("Delete messages: %s", msg)
//
// [Delete an artifact store]: https://api.gocd.org/current/#delete-an-artifact-store
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
