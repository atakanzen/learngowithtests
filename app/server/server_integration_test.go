package server_test

import (
	"fmt"
	"learngowithtests/app/helper"
	"learngowithtests/app/server"
	"learngowithtests/app/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testScoreCount = 3

func TestPostScoreAndGetScore(t *testing.T) {
	db, cleanDb := helper.CreateTmpFile(t, "")
	defer cleanDb()

	playerStore, err := store.NewFileSystemPlayerStore(db)
	helper.AssertErrNil(t, err)

	playerServer := server.NewPlayerServer(playerStore)
	player := "Fredry"

	for i := 0; i < testScoreCount; i++ {
		playerServer.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newScoreRequest(player, http.MethodGet))

		helper.AssertResCode(t, response.Code, http.StatusOK)
		helper.AssertResBody(t, response.Body.String(), fmt.Sprint(testScoreCount))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newLeagueRequest())

		helper.AssertResCode(t, response.Code, http.StatusOK)
		got := helper.GetLeagueBody(t, response.Body)
		want := store.League{
			{Name: "Fredry", Score: testScoreCount},
		}

		helper.AssertLeagueBody(t, got, want)
	})
}
