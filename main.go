package main

import (
	di "learngowithtests/dependency_injection"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(di.GreetHandler)))
}
