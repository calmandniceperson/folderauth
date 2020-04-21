// Author(s): Michael Koeppl

package folderauth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	Hostname   string
	Username   string
	Token      string
	HTTPClient *http.Client
}

func hostnameHasCorrectPrefix(hostname string) bool {
	return strings.HasPrefix(hostname, "http://") || strings.HasPrefix(hostname, "https://")
}

// NewClient creates a new Client and returns it. If the hostname is invalid,
// an error will be returned.
func NewClient(hostname, username, token string, httpClient *http.Client) (*Client, error) {
	if !hostnameHasCorrectPrefix(hostname) {
		return nil, fmt.Errorf("Invalid hostname %s", hostname)
	}

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{hostname, username, token, httpClient}, nil
}

func attachAuthHeader(req *http.Request, username, token string) {
	encodedAuthString := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, token)))
	req.Header.Add("Authorization", "Basic "+encodedAuthString)
}

// performRequest creates a request, attaches the basic authentication header
// to it and sends the request off.
func (c *Client) performRequest(method, route string, body io.Reader, contentType string) (*http.Response, error) {
	url := fmt.Sprintf("%s/folder-auth/%s", strings.TrimSuffix(c.Hostname, "/"), strings.TrimPrefix(route, "/"))

	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, body)
		req.Header.Add("Content-Type", contentType)
	}
	if err != nil {
		return nil, err
	}
	attachAuthHeader(req, c.Username, c.Token)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		d, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return resp, errors.New(string(d))
	}

	return resp, nil
}
