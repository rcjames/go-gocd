package gocd

import "fmt"

func (c *GoCDClient) GetPipeline(name string) (Pipeline, string, error) {
	var pipeline Pipeline

	requestPath := fmt.Sprintf("go/api/admin/pipelines/%s", name)
	etag, err := c.getRequest(requestPath, "", &pipeline)
	if err != nil {
		return pipeline, "", err
	}

	return pipeline, etag, err
}
