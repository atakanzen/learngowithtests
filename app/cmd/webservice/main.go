package main

import (
	"fmt"
	"learngowithtests/app/server"
	"learngowithtests/app/store"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(fmt.Sprintf("./db/%s", dbFileName), os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("could not open %s, %v", dbFileName, err)
	}

	playerStore, err := store.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("could not create FS player store, %v", err)
	}

	playerServer := server.NewPlayerServer(playerStore)

	if err := http.ListenAndServe(":5000", playerServer); err != nil {
		log.Fatalf("could not listen on port :5000, %v", err)
	}
}
