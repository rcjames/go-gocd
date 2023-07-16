package gocd

func (c *GoCDClient) GetAllMaterials() ([]Material, error) {
	var allMaterials GetAllMaterialsResponse

	_, err := c.getRequest("go/api/config/materials", "", &allMaterials)
	if err != nil {
		return nil, err
	}

	return allMaterials.Embedded.Materials, err
}
