package server

import (
	"encoding/json"
	"fmt"
	"learngowithtests/app/store"
	"net/http"
	"strings"
)

const (
	ContentTypeJSON = "application/json"
)

// PlayerStore is an interface that implements FileSystemPlayerStore's functions
type PlayerStore interface {
	GetPlayerScore(name string) int
	PostPlayerScore(name string)
	GetLeague() store.League
}

// PlayerServer is a struct that has PlayerStore and http.Handler fields
type PlayerServer struct {
	PlayerStore PlayerStore
	http.Handler
}

// NewPlayerServer returns a player server with the given PlayerStore
func NewPlayerServer(playerStore PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.PlayerStore = playerStore

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
	w.Header().Set("content-type", ContentTypeJSON)
	json.NewEncoder(w).Encode(p.PlayerStore.GetLeague())
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) postScore(w http.ResponseWriter, player string) {
	p.PlayerStore.PostPlayerScore(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) getScore(w http.ResponseWriter, player string) {
	score := p.PlayerStore.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, p.PlayerStore.GetPlayerScore(player))
}
