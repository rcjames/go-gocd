package gocd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GoCDClient struct {
	Host     string
	Client   *http.Client
	Username string
	Password string
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

func (c *GoCDClient) getRequest(path, apiVersion string, response interface{}) error {
	req, err := c.setupRequest(path, "GET", apiVersion, nil)
	if err != nil {
		return err
	}

	err = c.makeRequest(req, response)
	if err != nil {
		return err
	}

	return nil
}

func (c *GoCDClient) postRequest(path, apiVersion string, data, response interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	dataBody := bytes.NewReader(b)

	req, err := c.setupRequest(path, "POST", apiVersion, dataBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	err = c.makeRequest(req, &response)
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

func (c *GoCDClient) makeRequest(req *http.Request, response interface{}) error {
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return err // TODO - Create this error
	}

	body, err := ioutil.ReadAll(resp.Body) // TODO - Handle error
	if err != nil {
		fmt.Printf("Error reading boxy for %s", req.URL)
		return err
	}

	if err := json.Unmarshal(body, &response); err != nil { // Parse []byte to the go struct pointer
		fmt.Printf("Can not unmarshal JSON: %s\n", err)
		return err
	}

	return nil

}
