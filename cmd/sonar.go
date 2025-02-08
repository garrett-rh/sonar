package main

import (
	"encoding/json"
	"fmt"
	"github.com/garrett-rh/sonar"
	"log"
	"time"
)

var version = "v2"

func main() {
	timeout := 5 * time.Second
	c := sonar.NewRegistryClient(fmt.Sprintf("http://localhost/%s", version), timeout)

	log.Println(c.Check())

	repository := c.GetRepositories()
	r, err := json.MarshalIndent(repository, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(r))

	tags := c.GetTags(repository.Repositories[0])
	t, err := json.MarshalIndent(tags, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(t))

	manifests := c.GetManifest(repository.Repositories[0], tags.Tags[0])
	m, err := json.MarshalIndent(manifests, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(m))

}
