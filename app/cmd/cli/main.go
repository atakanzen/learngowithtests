package main

import (
	"fmt"
	"learngowithtests/app/poker"
	"learngowithtests/app/store"
	"log"
	"os"
)

const dbFileName = "game.cli.db.json"

func main() {

	fmt.Println("Let's play the game!")
	fmt.Println("Type '{Name} wins' to score a win for the given player")

	playerStore, close, err := store.NewFileSystemPlayerStoreFromFile(fmt.Sprintf("../../db/%s", dbFileName))
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := poker.NewCLI(playerStore, os.Stdin, poker.BlindAlerterFunc(poker.StdOutAlerter))
	game.PlayPoker()
}
