package main

import (
	"encoding/json"
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

func TestGETPlayerScore(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Legolas": 20,
			"Gimli":   12,
		},
		nil,
	}
	server := NewPlayerServer(store)
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

			server.ServeHTTP(res, req)

			assertResCode(t, res.Code, test.statusCode)
			assertResBody(t, res.Body.String(), test.score)
		})
	}

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newScoreRequest("non-existing", http.MethodGet)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResCode(t, res.Code, http.StatusNotFound)
	})
}

func TestPOSTPlayerScore(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
	}
	server := NewPlayerServer(store)

	t.Run("returns status accepted on POST", func(t *testing.T) {
		player := "Legolas"

		req := newScoreRequest(player, http.MethodPost)
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

func TestGetLeague(t *testing.T) {
	store := &StubPlayerStore{}
	server := NewPlayerServer(store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/league", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		var got []Player
		err := json.NewDecoder(res.Body).Decode(&got)

		if err != nil {
			t.Fatalf("unable to parse response %q, into slice of Player '%v'", res.Body, err)
		}

		assertResCode(t, res.Code, http.StatusOK)
	})
}

func newScoreRequest(name, method string) *http.Request {
	req, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", name), nil)
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
