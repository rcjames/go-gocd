package gocd

import "fmt"

// GetPipeline fetches the configuration for the pipeline matching the provided
// name using the "[Get pipeline config]" endpoint, returning the pipeline config
// and the ETAG.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	pipeline, etag, _ := c.GetPipeline("pipeline1")
//	fmt.Printf("Pipeline: %s has the etag %s", pipeline.Name, etag)
//
// [Get pipeline config]: https://api.gocd.org/current/#get-pipeline-config
func (c *GoCDClient) GetPipeline(name string) (Pipeline, string, error) {
	var pipeline Pipeline

	requestPath := fmt.Sprintf("go/api/admin/pipelines/%s", name)
	etag, err := c.getRequest(requestPath, "", &pipeline)
	if err != nil {
		return pipeline, "", err
	}

	return pipeline, etag, err
}

// CreatePipeline creates a Pipeline with the provided configuration using the
// "[Create a pipeline]" endpoint, returning the pipeline configuration and the
// ETAG.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	pipelineName := "pipeline1"
//
//	pipelineTask := PipelineTask{
//		Type: "exec",
//		Attributes: &PipelineTaskAttributes{
//			Command:   "echo",
//			Arguments: []string{"1"},
//		},
//	}
//
//	pipelineJob := PipelineJob{
//		Name: "job1",
//	}
//	pipelineJob.AddTask(pipelineTask)
//
//	pipelineStage := PipelineStage{
//		Name: "stage1",
//	}
//	pipelineStage.AddJob(pipelineJob)
//
//	materialBranch := "main"
//	material := Material{
//		Type: "git",
//		Attributes: &MaterialAttributes{
//			Name:         "go-gocd",
//			Url:          "https://github.com/rcjames/go-gocd",
//			Branch:       materialBranch,
//			ShallowClone: true,
//		},
//	}
//
//	pipeline := Pipeline{
//		Name:  pipelineName,
//		Group: pipelineGroupName,
//	}
//	pipeline.AddStage(pipelineStage)
//	pipeline.AddMaterial(material)
//
//	pipelineResponse, etag, _ := c.CreatePipeline(pipeline)
//	fmt.Printf("Created pipeline %s with etag %s", pipelineResponse.Name, etag)
//
// [Create a pipeline]: https://api.gocd.org/current/#create-a-pipeline
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

// UpdatePipeline updates the configuration for an existing Pipeline using the
// "[Edit pipeline config]"" endpoint, returning the new configuration and the
// ETAG.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	pipelineName := "pipeline1"
//	p, etag, _ := gocd.GetPipeline(pipelineName)
//	p.Name = "pipeline2"
//	c.UpdatePipeline(pipelineName, etag, p)
//
// [Edit pipeline config]: https://api.gocd.org/current/#edit-pipeline-config
func (c *GoCDClient) UpdatePipeline(pipelineName, etag string, pipeline Pipeline) (Pipeline, string, error) {
	var pipelineResponse Pipeline
	requestPath := fmt.Sprintf("go/api/admin/pipelines/%s", pipelineName)

	etag, err := c.putRequest(requestPath, "", etag, pipeline, &pipelineResponse)
	if err != nil {
		return pipelineResponse, "", err
	}

	return pipelineResponse, etag, nil
}

// DeletePipeline deletes the pipeline using the "[Delete a pipeline]" endpoint
// and returns the DeleteMessage.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	msg, _ := c.DeletePipeline("pipeline1")
//	fmt.Println(msg)
//
// [Delete a pipeline]: https://api.gocd.org/current/#delete-a-pipeline
func (c *GoCDClient) DeletePipeline(name string) (string, error) {
	var message DeleteMessage
	requestPath := fmt.Sprintf("go/api/admin/pipelines/%s", name)

	err := c.deleteRequest(requestPath, "", &message)
	if err != nil {
		return "", err
	}

	return message.Message, nil
}

// AddTask is a utility function which adds a Task to a Job. See CreatePipeline
// for example usage.
func (j *PipelineJob) AddTask(task PipelineTask) {
	j.Tasks = append(j.Tasks, task)
}

// AddJob is a utility function which adds a Job to a Stage. See CreatePipeline
// for example usage.
func (s *PipelineStage) AddJob(job PipelineJob) {
	s.Jobs = append(s.Jobs, job)
}

// AddStage is a utility function which adds a Stage to a Pipeline. See
// CreatePipeline for example usage.
func (p *Pipeline) AddStage(stage PipelineStage) {
	p.Stages = append(p.Stages, stage)
}

// AddMaterial is a utility function which adds a Material to a Pipeline. See
// CreatePipeline for example usage.
func (p *Pipeline) AddMaterial(material Material) {
	p.Materials = append(p.Materials, material)
}
