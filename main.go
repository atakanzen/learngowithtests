package main

import (
	"learngowithtests/mocking"
	"os"
)

func main() {
	// log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(di.GreetHandler)))
	mocking.Countdown(os.Stdout)
}
