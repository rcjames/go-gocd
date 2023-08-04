package gocd

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
