package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostScoreAndGetScore(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Fredry"

	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	server.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newScoreRequest(player, http.MethodGet))

		assertResCode(t, response.Code, http.StatusOK)
		assertResBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())

		assertResCode(t, response.Code, http.StatusOK)
		got := getLeagueBody(t, response.Body)
		want := []Player{
			{"Fredry", 3},
		}

		assertLeagueBody(t, got, want)
	})
}
