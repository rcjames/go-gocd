package gocd

import "fmt"

func (c *GoCDClient) GetAllMaterials() ([]Material, error) {
	var allMaterials GetAllMaterialsResponse

	_, err := c.getRequest("go/api/config/materials", "", &allMaterials)
	if err != nil {
		return nil, err
	}

	return allMaterials.Embedded.Materials, err
}

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

func (c *GoCDClient) GetMaterialModifications(fingerprint string) (GetMaterialModificationsResponse, error) {
	return c.GetMaterialModificationsWithOffset(fingerprint, -1)
}
