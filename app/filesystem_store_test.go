package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		db, cleanDb := createTmpFile(t, `[
			{"Name": "Adam", "Score": 25},
			{"Name": "Karina", "Score": 90}
		]`)
		defer cleanDb()

		store := NewFileSystemPlayerStore(db)

		got := store.GetLeague()

		want := []Player{
			{"Adam", 25},
			{"Karina", 90},
		}

		assertLeagueBody(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		db, cleanDb := createTmpFile(t, `[
			{"Name": "John", "Score": 25},
			{"Name": "Kate", "Score": 22}
		]`)
		defer cleanDb()

		store := NewFileSystemPlayerStore(db)

		got := store.GetPlayerScore("Kate")

		want := 22

		assertScore(t, got, want)
	})

	t.Run("save score to existing player", func(t *testing.T) {
		db, cleanDb := createTmpFile(t, `[
			{"Name": "Costya", "Score": 4},	
			{"Name": "Katia", "Score": 5}
		]`)
		defer cleanDb()

		store := NewFileSystemPlayerStore(db)

		store.PostPlayerScore("Katia")

		got := store.GetPlayerScore("Katia")
		want := 6

		assertScore(t, got, want)
	})

	t.Run("save score to new player", func(t *testing.T) {
		db, cleanDb := createTmpFile(t, `[
			{"Name": "Arsen", "Score": 5}
		]`)
		defer cleanDb()

		store := NewFileSystemPlayerStore(db)

		store.PostPlayerScore("Nastya")

		got := store.GetPlayerScore("Nastya")
		want := 1

		assertScore(t, got, want)
	})
}

func assertScore(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got score %d, want %d", got, want)
	}
}

func createTmpFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create tmp file, %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
