package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostScoreAndGetScore(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Fredry"

	server.ServeHTTP(httptest.NewRecorder(), newRequest(player, http.MethodPost))
	server.ServeHTTP(httptest.NewRecorder(), newRequest(player, http.MethodPost))
	server.ServeHTTP(httptest.NewRecorder(), newRequest(player, http.MethodPost))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newRequest(player, http.MethodGet))

	assertResCode(t, response.Code, http.StatusOK)
	assertResBody(t, response.Body.String(), "3")
}
