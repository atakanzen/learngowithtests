package server_test

import (
	"learngowithtests/app/helper"
	"learngowithtests/app/server"
	"learngowithtests/app/store"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostScoreAndGetScore(t *testing.T) {
	db, cleanDb := helper.CreateTmpFile(t, "")
	defer cleanDb()

	playerStore, err := store.NewFileSystemPlayerStore(db)
	helper.AssertErrNil(t, err)

	playerServer := server.NewPlayerServer(playerStore)
	player := "Fredry"

	// TODO with for loop atr the amount of specified n
	playerServer.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	playerServer.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))
	playerServer.ServeHTTP(httptest.NewRecorder(), newScoreRequest(player, http.MethodPost))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newScoreRequest(player, http.MethodGet))

		helper.AssertResCode(t, response.Code, http.StatusOK)
		helper.AssertResBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, newLeagueRequest())

		helper.AssertResCode(t, response.Code, http.StatusOK)
		got := helper.GetLeagueBody(t, response.Body)
		want := store.League{
			{Name: "Fredry", Score: 3},
		}

		helper.AssertLeagueBody(t, got, want)
	})
}
