package store_test

import (
	"learngowithtests/app/helper"
	"learngowithtests/app/store"
	"testing"
)

func TestFileSystemPlayerStore(t *testing.T) {
	t.Run("works with an empty file", func(t *testing.T) {
		db, cleanDb := helper.CreateTmpFile(t, "")
		defer cleanDb()

		_, err := store.NewFileSystemPlayerStore(db)
		helper.AssertErrNil(t, err)
	})

	t.Run("sorted league from a reader", func(t *testing.T) {
		db, cleanDb := helper.CreateTmpFile(t, `[
			{"Name": "Adam", "Score": 25},
			{"Name": "Karina", "Score": 90}
		]`)
		defer cleanDb()

		playerStore, err := store.NewFileSystemPlayerStore(db)
		helper.AssertErrNil(t, err)

		got := playerStore.GetLeague()

		want := store.League{
			{Name: "Karina", Score: 90},
			{Name: "Adam", Score: 25},
		}

		helper.AssertLeagueBody(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		db, cleanDb := helper.CreateTmpFile(t, `[
			{"Name": "John", "Score": 25},
			{"Name": "Kate", "Score": 22}
		]`)
		defer cleanDb()

		store, err := store.NewFileSystemPlayerStore(db)
		helper.AssertErrNil(t, err)

		got := store.GetPlayerScore("Kate")

		want := 22

		helper.AssertScore(t, got, want)
	})

	t.Run("save score to existing player", func(t *testing.T) {
		db, cleanDb := helper.CreateTmpFile(t, `[
			{"Name": "Costya", "Score": 4},	
			{"Name": "Katia", "Score": 5}
		]`)
		defer cleanDb()

		store, err := store.NewFileSystemPlayerStore(db)
		helper.AssertErrNil(t, err)

		store.PostPlayerScore("Katia")

		got := store.GetPlayerScore("Katia")
		want := 6

		helper.AssertScore(t, got, want)
	})

	t.Run("save score to new player", func(t *testing.T) {
		db, cleanDb := helper.CreateTmpFile(t, `[
			{"Name": "Arsen", "Score": 5}
		]`)
		defer cleanDb()

		store, err := store.NewFileSystemPlayerStore(db)
		helper.AssertErrNil(t, err)

		store.PostPlayerScore("Nastya")

		got := store.GetPlayerScore("Nastya")
		want := 1

		helper.AssertScore(t, got, want)
	})
}
