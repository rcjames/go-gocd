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
