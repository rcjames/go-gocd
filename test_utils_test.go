package gocd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

// https://medium.com/zus-health/mocking-outbound-http-requests-in-go-youre-probably-doing-it-wrong-60373a38d2aa
func NewMockClient(t *testing.T) (GoCDClient, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			regex := regexp.MustCompile(`go/api/(.*)`)
			matches := regex.FindStringSubmatch(r.URL.Path)
			filename := fmt.Sprintf("test-data/%s.json", matches[1])

			s, err := ioutil.ReadFile(filename)
			if err != nil {
				return
			}

			w.Write(s)
		} else if r.Method == "POST" {
			w.Write([]byte("{}"))
		}
	}))

	c := New(server.URL, "", "")

	return c, server
}
