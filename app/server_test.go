package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) PostPlayerScore(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeagueTable() []Player {
	return s.league
}

func TestGETPlayerScore(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Legolas": 20,
			"Gimli":   12,
		},
		nil,
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

	t.Run("it returns league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Sam", 1},
			{"Neil", 2},
			{"Bill", 3},
		}

		store := &StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(store)

		req := newLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := getLeagueBody(t, res.Body)

		assertResCode(t, res.Code, http.StatusOK)
		assertLeagueBody(t, got, wantedLeague)
		assertContentType(t, res.Header().Get("content-type"), contentTypeJSON)
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

func getLeagueBody(t testing.TB, body *bytes.Buffer) []Player {
	t.Helper()

	league, err := NewLeague(body)

	if err != nil {
		t.Error(err)
	}

	return league
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

func assertLeagueBody(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertContentType(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response did not have content-type of \"application/json\", got %q", got)
	}
}
