package server_test

import (
	"fmt"
	"learngowithtests/app/helper"
	"learngowithtests/app/mock"
	"learngowithtests/app/server"
	"learngowithtests/app/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayerScore(t *testing.T) {
	playerStore := &mock.StubPlayerStore{
		Scores: map[string]int{
			"Legolas": 20,
			"Gimli":   12,
		},
		WinCalls: nil,
		League:   nil,
	}
	playerServer := server.NewPlayerServer(playerStore)
	cases := []struct {
		name       string
		player     string
		score      string
		statusCode int
	}{
		{"returns Legolas' score", "Legolas", "20", 200},
		{"returns Gimli's score", "Gimli", "12", 200},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req := newScoreRequest(test.player, http.MethodGet)
			res := httptest.NewRecorder()

			playerServer.ServeHTTP(res, req)

			helper.AssertResCode(t, res.Code, test.statusCode)
			helper.AssertResBody(t, res.Body.String(), test.score)
		})
	}

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newScoreRequest("non-existing", http.MethodGet)
		res := httptest.NewRecorder()

		playerServer.ServeHTTP(res, req)

		helper.AssertResCode(t, res.Code, http.StatusNotFound)
	})
}

func TestPOSTPlayerScore(t *testing.T) {
	playerStore := &mock.StubPlayerStore{
		Scores:   map[string]int{},
		WinCalls: nil,
		League:   nil,
	}
	playerServer := server.NewPlayerServer(playerStore)

	t.Run("returns status accepted on POST", func(t *testing.T) {
		player := "Legolas"

		req := newScoreRequest(player, http.MethodPost)
		res := httptest.NewRecorder()

		playerServer.ServeHTTP(res, req)

		helper.AssertResCode(t, res.Code, http.StatusAccepted)
		helper.AssertPlayerWin(t, playerStore, player)
	})
}

func TestGetLeague(t *testing.T) {

	t.Run("it returns league table as JSON", func(t *testing.T) {
		wantedLeague := store.League{
			{Name: "Sam", Score: 1},
			{Name: "Neil", Score: 2},
			{Name: "Bill", Score: 3},
		}

		playerStore := &mock.StubPlayerStore{
			Scores:   nil,
			WinCalls: nil,
			League:   wantedLeague,
		}
		playerServer := server.NewPlayerServer(playerStore)

		req := newLeagueRequest()
		res := httptest.NewRecorder()

		playerServer.ServeHTTP(res, req)

		got := helper.GetLeagueBody(t, res.Body)

		helper.AssertResCode(t, res.Code, http.StatusOK)
		helper.AssertLeagueBody(t, got, wantedLeague)
		helper.AssertContentType(t, res.Header().Get("content-type"), server.ContentTypeJSON)
	})
}

func newScoreRequest(name, method string) *http.Request {
	req, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
