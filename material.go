package gocd

import "fmt"

// GetAllMaterials gets a list of Materials configured across all pipelines via
// the "[Get all materials]" API.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	materials := c.GetAllMaterials()
//	for _, m := range materials {
//		fmt.Println(m.Type)
//	}
//
// [Get all materials]: https://api.gocd.org/current/#get-all-materials
func (c *GoCDClient) GetAllMaterials() ([]Material, error) {
	var allMaterials GetAllMaterialsResponse

	_, err := c.getRequest("go/api/config/materials", "", &allMaterials)
	if err != nil {
		return nil, err
	}

	return allMaterials.Embedded.Materials, err
}

// GetMaterialModificationsWithOffset fetches Material modifications with a
// paginated offset based on the provided fingerprint using the "[Get material
// modifications]" API.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	materials := c.GetAllMaterials()
//	for _, m := range materials {
//		changes := c.GetMaterialModificationsWithOffset(m.Fingerprint, 0)
//		fmt.Printf("%s\n", changes.Embedded.Materials[0].Comment)
//	}
//
// [Get material modifications]: https://api.gocd.org/current/#get-material-modifications
func (c *GoCDClient) GetMaterialModificationsWithOffset(fingerprint string, offset int) (GetMaterialModificationsResponse, error) {
	var allModifications GetMaterialModificationsResponse
	requestPath := fmt.Sprintf("go/api/materials/%s/modifications", fingerprint)

	// Assuming there is an offset
	if offset > -1 {
		requestPath = fmt.Sprintf("%s/%d", requestPath, offset)
	}

	_, err := c.getRequest(requestPath, "", &allModifications)
	if err != nil {
		return allModifications, err
	}

	return allModifications, err
}

// GetMaterialModifications fetch the first page of Material modifiations based
// on the provided fingerprint. See GetMaterialModificationsWithOffset, for
// further information, but without the offset parameter.
func (c *GoCDClient) GetMaterialModifications(fingerprint string) (GetMaterialModificationsResponse, error) {
	return c.GetMaterialModificationsWithOffset(fingerprint, -1)
}
