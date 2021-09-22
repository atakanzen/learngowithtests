package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	PostPlayerScore(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.getScore(w, player)
	case http.MethodPost:
		p.postScore(w, player)
	}
}

func (p *PlayerServer) postScore(w http.ResponseWriter, player string) {
	p.store.PostPlayerScore(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) getScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
