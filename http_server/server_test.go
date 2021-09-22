package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) PostPlayerScore(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Legolas": 20,
			"Gimli":   12,
		},
		nil,
	}
	server := &PlayerServer{store}
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
			req := newGetScoreRequest(test.player)
			res := httptest.NewRecorder()

			server.ServeHTTP(res, req)

			assertResCode(t, res.Code, test.statusCode)
			assertResBody(t, res.Body.String(), test.score)
		})
	}

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("non-existing")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResCode(t, res.Code, http.StatusNotFound)
	})
}

func TestPOSTStore(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{store}

	t.Run("returns status accepted on POST", func(t *testing.T) {
		player := "Legolas"

		req := newPostScoreRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResCode(t, res.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to PostPlayerScore, want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})

}
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got score %q, want %q", got, want)
	}
}

func assertResCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want %d", got, want)
	}
}
