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

func (c *GoCDClient) CreatePipeline(pipeline Pipeline) (Pipeline, string, error) {
	var pipelineResponse Pipeline
	pipelineRequest := PipelineCreateRequest{
		Group:    pipeline.Group,
		Pipeline: pipeline,
	}

	etag, err := c.postRequest("go/api/admin/pipelines", "", pipelineRequest, &pipelineResponse)
	if err != nil {
		return pipelineResponse, "", err
	}

	return pipelineResponse, etag, nil
}

// TODO - Write test
func (c *GoCDClient) UpdatePipeline(pipelineName, etag string, pipeline Pipeline) (Pipeline, string, error) {
	var pipelineResponse Pipeline
	requestPath := fmt.Sprintf("go/api/admin/pipelines/%s", pipelineName)

	etag, err := c.putRequest(requestPath, "", etag, pipeline, &pipelineResponse)
	if err != nil {
		return pipelineResponse, "", err
	}

	return pipelineResponse, etag, nil
}

func (c *GoCDClient) DeletePipeline(name string) (string, error) {
	var message DeleteMessage
	requestPath := fmt.Sprintf("go/api/admin/pipelines/%s", name)

	err := c.deleteRequest(requestPath, "", &message)
	if err != nil {
		return "", err
	}

	return message.Message, nil
}

func (j *PipelineJob) AddTask(task PipelineTask) {
	j.Tasks = append(j.Tasks, task)
}

func (s *PipelineStage) AddJob(job PipelineJob) {
	s.Jobs = append(s.Jobs, job)
}

func (p *Pipeline) AddStage(stage PipelineStage) {
	p.Stages = append(p.Stages, stage)
}

func (p *Pipeline) AddMaterial(material Material) {
	p.Materials = append(p.Materials, material)
}
