package main

import (
	"strings"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		db := strings.NewReader(`[
			{"Name": "Adam", "Score": 25},
			{"Name": "Karina", "Score": 90}
		]`)

		store := FileSystemPlayerStore{db}

		got := store.GetLeague()

		want := []Player{
			{"Adam", 25},
			{"Karina", 90},
		}

		assertLeagueBody(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		db := strings.NewReader(`[
			{"Name": "John", "Score": 25},
			{"Name": "Kate", "Score": 22}
		]`)

		store := FileSystemPlayerStore{db}

		got := store.GetPlayerScore("Kate")

		want := 22

		assertScore(t, got, want)
	})
}

func assertScore(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got score %d, want %d", got, want)
	}
}
