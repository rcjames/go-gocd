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

type GoCDClient struct {
	Host     string
	Client   *http.Client
	Username string
	Password string
}

type HttpError struct {
	StatusCode int
	Err        error
}

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

func (h *HttpError) Error() string {
	return fmt.Sprintf("status %d: error %v", h.StatusCode, h.Err)
}
