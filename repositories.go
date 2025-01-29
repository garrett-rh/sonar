package main

import (
	"encoding/json"
	"log"
)

type Repositories struct {
	Repositories []string `json:"repositories"`
}

func (r *RegistryClient) GetRepositories() Repositories {
	var repositories Repositories

	resp := r.Get("/v2/_catalog")

	err := json.NewDecoder(resp.Body).Decode(&repositories)
	if err != nil {
		log.Fatal(err)
	}

	return repositories

}
