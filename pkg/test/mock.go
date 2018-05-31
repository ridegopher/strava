package mock

import (
	"github.com/strava/go.strava"
	"net/http"
)

var mockDirectory = "testdata"

func NewMockClient(token, data string) *Client {
	c := strava.NewClient(token)
	c.httpClient = &http.Client{
		Transport: &MockTransport{
			token:     token,
			directory: cassetteDirectory,
			data:      data},
	}

	return c
}

type MockTransport struct {
	http.Transport
	token     string
	directory string
	data      string
}
