package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	contentTypeJSON = "application/json"
)

type Player struct {
	Name  string
	Score int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	PostPlayerScore(name string)
	GetLeagueTable() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	p.Handler = router

	return p
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodGet:
		p.getScore(w, player)
	case http.MethodPost:
		p.postScore(w, player)
	}
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", contentTypeJSON)
	json.NewEncoder(w).Encode(p.store.GetLeagueTable())
	w.WriteHeader(http.StatusOK)
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
