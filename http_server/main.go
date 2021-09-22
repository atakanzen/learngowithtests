package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct {
}

func (i InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i InMemoryPlayerStore) PostPlayerScore(name string) {

}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
