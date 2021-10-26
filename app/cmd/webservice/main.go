package main

import (
	"fmt"
	"learngowithtests/app/server"
	"learngowithtests/app/store"
	"log"
	"net/http"
)

const dbFileName = "game.web.db.json"

func main() {
	playerStore, close, err := store.NewFileSystemPlayerStoreFromFile(fmt.Sprintf("../../db/%s", dbFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	playerServer := server.NewPlayerServer(playerStore)

	if err := http.ListenAndServe(":5000", playerServer); err != nil {
		log.Fatalf("could not listen on port :5000, %v", err)
	}
}
