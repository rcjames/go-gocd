package gocd

// A Version maps to a [version] object.
//
// [version]: https://api.gocd.org/current/#the-version-object
type Version struct {
	Version     string `json:"version"`
	BuildNumber string `json:"build_number"`
	GitSha      string `json:"git_sha"`
	FullVersion string `json:"full_version"`
	CommitUrl   string `json:"commit_url"`
}

// GetVersion feth the version information from the "[Get version]" endpoint.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
//	version, _ := c.GetVersion()
//	fmt.Println(version.FullVersion)
//
// [Get version]: https://api.gocd.org/current/#get-version
func (c *GoCDClient) GetVersion() (Version, error) {
	var version Version

	_, err := c.getRequest("go/api/version", "", &version)
	if err != nil {
		return version, err
	}
	return version, nil
}
