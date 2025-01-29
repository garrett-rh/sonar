package main

import (
	"log"
	"time"
)

func main() {
	timeout := 5 * time.Second
	c := NewRegistryClient("http://localhost", timeout)
	log.Println(c.GetRepositories().Repositories)
}
