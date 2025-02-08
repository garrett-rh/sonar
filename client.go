package sonar

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type RegistryClient struct {
	client *http.Client
	uri    string
}

func NewRegistryClient(uri string, timeout time.Duration) *RegistryClient {
	return &RegistryClient{
		client: &http.Client{
			Timeout: timeout,
		}, uri: uri,
	}
}

func (r *RegistryClient) Check() int {
	resp := r.Get("")

	return resp.StatusCode
}

func (r *RegistryClient) Get(endpoint string) *http.Response {
	request, err := http.NewRequest("GET", fmt.Sprintf("%s%s", r.uri, endpoint), nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := r.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
