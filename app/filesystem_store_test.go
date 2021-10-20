package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("works with an empty file", func(t *testing.T) {
		db, cleanDb := createTmpFile(t, "")
		defer cleanDb()

		_, err := NewFileSystemPlayerStore(db)
		assertErrNil(t, err)
	})

	t.Run("sorted league from a reader", func(t *testing.T) {
		db, cleanDb := createTmpFile(t, `[
			{"Name": "Adam", "Score": 25},
			{"Name": "Karina", "Score": 90}
		]`)
		defer cleanDb()

		store, err := NewFileSystemPlayerStore(db)
		assertErrNil(t, err)

		got := store.GetLeague()

		want := []Player{
			{"Karina", 90},
			{"Adam", 25},
		}

		assertLeagueBody(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		db, cleanDb := createTmpFile(t, `[
			{"Name": "John", "Score": 25},
			{"Name": "Kate", "Score": 22}
		]`)
		defer cleanDb()

		store, err := NewFileSystemPlayerStore(db)
		assertErrNil(t, err)

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

		store, err := NewFileSystemPlayerStore(db)
		assertErrNil(t, err)

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

		store, err := NewFileSystemPlayerStore(db)
		assertErrNil(t, err)

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

func assertErrNil(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("did not expect error, but got one %q,", got)
	}
}

func createTmpFile(t testing.TB, initialData string) (*os.File, func()) {
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
