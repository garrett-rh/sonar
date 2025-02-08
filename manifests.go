package sonar

import (
	"encoding/json"
	"fmt"
	"log"
)

type Manifest struct {
	SchemaVersion int    `json:"schemaVersion"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
	Architecture  string `json:"architecture"`
	FsLayers      []struct {
		BlobSum string `json:"blobSum"`
	} `json:"fsLayers"`
}

func (r *RegistryClient) GetManifest(image string, tag string) Manifest {
	var manifest Manifest

	resp := r.Get(fmt.Sprintf("/%s/manifests/%s", image, tag))

	err := json.NewDecoder(resp.Body).Decode(&manifest)
	if err != nil {
		log.Fatal(err)
	}

	return manifest
}
