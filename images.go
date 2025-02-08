package sonar

import (
	"encoding/json"
	"fmt"
	"log"
)

type Images struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func (r *RegistryClient) GetTags(name string) Images {
	var images Images

	resp := r.Get(fmt.Sprintf("/%s/tags/list", name))

	err := json.NewDecoder(resp.Body).Decode(&images)
	if err != nil {
		log.Fatal(err)
	}

	return images
}
