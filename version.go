package gocd

type Version struct {
	Version     string `json:"version"`
	BuildNumber string `json:"build_number"`
	GitSha      string `json:"git_sha"`
	FullVersion string `json:"full_version"`
	CommitUrl   string `json:"commit_url"`
}

func (c *GoCDClient) GetVersion() (Version, error) {
	var version Version

	_, err := c.getRequest("go/api/version", "", &version)
	if err != nil {
		return version, err
	}
	return version, nil
}
