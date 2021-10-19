package main

import (
	"fmt"
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

	store := NewFileSystemPlayerStore(db)
	server := NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port :5000, %v", err)
	}
}
