package main

import (
	"learngowithtests/mocking"
	"os"
	"time"
)

func main() {
	// log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(di.GreetHandler)))
	sleeper := mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepCb: time.Sleep}
	mocking.Countdown(os.Stdout, &sleeper)
}
