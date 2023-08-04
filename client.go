package gocd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GoCDClient contains the connection information and HTTP client to allow
// requests to be made.
type GoCDClient struct {
	Host     string
	Client   *http.Client
	Username string
	Password string
}

// HttpError wraps up the status code and error object returned by GoCD.
type HttpError struct {
	StatusCode int
	Err        error
}

// New creates a new GoCDClient which can be used for making request. Currently
// only basic authentication has been implemented.
//
// Example usage:
//
//	c := gocd.New(hostname, username, password)
func New(host, username, password string) GoCDClient {
	client := &http.Client{}
	gocdClient := GoCDClient{
		Host:     host,
		Client:   client,
		Username: username,
		Password: password,
	}
	return gocdClient
}

// getRequest is for internal use only. It makes a GET request to the specified
// path, setting the optional API version and using a provided interface to
// return the response. The ETAG and any errors are returned from the funtion.
func (c *GoCDClient) getRequest(path, apiVersion string, response interface{}) (string, error) {
	req, err := c.setupRequest(path, "GET", apiVersion, nil)
	if err != nil {
		return "", err
	}

	etag, err := c.makeRequest(req, response)
	if err != nil {
		return "", err
	}

	return etag, nil
}

// postRequest is for internal use only. It makes a POST request to the specified
// path, setting the optional API version. The provided interfaces contain the
// object to be created and then to return the response object. The ETAG and any
// errors are returned from the function.
func (c *GoCDClient) postRequest(path, apiVersion string, data, response interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	dataBody := bytes.NewReader(b)

	req, err := c.setupRequest(path, "POST", apiVersion, dataBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	etag, err := c.makeRequest(req, &response)
	if err != nil {
		return "", err
	}

	return etag, nil
}

// putRequest is for internal use only. It operates in a similar way to postRequest
// but requires an ETAG which is set as a header.
func (c *GoCDClient) putRequest(path, apiVersion, etag string, data, response interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	dataBody := bytes.NewReader(b)

	req, err := c.setupRequest(path, "PUT", apiVersion, dataBody)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	if etag != "" {
		req.Header.Set("If-Match", etag)
	}

	etag, err = c.makeRequest(req, &response)
	if err != nil {
		return "", err
	}

	return etag, nil
}

// deleteRequest is for internal use only. It makes a DELETE request to the
// specified path, setting the optional API version and using the provided
// interface to return the response.
func (c *GoCDClient) deleteRequest(path, apiVersion string, response interface{}) error {
	req, err := c.setupRequest(path, "DELETE", apiVersion, nil)
	if err != nil {
		return err
	}

	_, err = c.makeRequest(req, &response)
	if err != nil {
		return err
	}

	return nil
}

// setupRequest is for internal use only. This is a utility function used to
// setup any type of request.
func (c *GoCDClient) setupRequest(path, method, apiVersion string, data *bytes.Reader) (*http.Request, error) {
	fullUrl := fmt.Sprintf("%s/%s", c.Host, path)

	var req *http.Request
	var err error

	if data != nil {
		req, err = http.NewRequest(method, fullUrl, data)
	} else {
		req, err = http.NewRequest(method, fullUrl, nil)
	}
	if err != nil {
		return nil, err
	}

	// Use the latest API if one is not provided
	if apiVersion == "" {
		apiVersion = "application/vnd.go.cd+json"
	}
	req.Header.Set("Accept", apiVersion)

	if len(c.Username) > 0 {
		req.SetBasicAuth(c.Username, c.Password)
	}

	return req, nil
}

// makeRequest is for internal use only. This is a utility function used to
// make the web request for any type of request.
func (c *GoCDClient) makeRequest(req *http.Request, response interface{}) (string, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("status code %d response code. response body: %s", resp.StatusCode, string(body)))
	}

	if err := json.Unmarshal(body, &response); err != nil { // Parse []byte to the go struct pointer
		return "", err
	}

	etag := strings.ReplaceAll(resp.Header.Get("ETag"), string('"'), "")

	return etag, nil
}

// HttpError Error is used to handle the custom error type returned by GoCD.
func (h *HttpError) Error() string {
	return fmt.Sprintf("status %d: error %v", h.StatusCode, h.Err)
}
