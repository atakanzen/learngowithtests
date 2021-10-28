package poker_test

import (
	"fmt"
	"learngowithtests/app/helper"
	"learngowithtests/app/mock"
	"learngowithtests/app/poker"
	"strings"
	"testing"
	"time"
)

func TestCLI(t *testing.T) {
	var dummySpyAlerter = &mock.SpyBlindAlerter{}

	t.Run("it schedules prints for blind bets", func(t *testing.T) {
		in := strings.NewReader("Gece wins\n")
		playerStore := &mock.StubPlayerStore{}
		blindAlerter := &mock.SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []mock.ScheduledAlert{
			{At: 0 * time.Second, BetAmount: 100},
			{At: 10 * time.Minute, BetAmount: 200},
			{At: 20 * time.Minute, BetAmount: 300},
			{At: 30 * time.Minute, BetAmount: 400},
			{At: 40 * time.Minute, BetAmount: 500},
			{At: 50 * time.Minute, BetAmount: 600},
			{At: 60 * time.Minute, BetAmount: 800},
			{At: 70 * time.Minute, BetAmount: 1000},
			{At: 80 * time.Minute, BetAmount: 2000},
			{At: 90 * time.Minute, BetAmount: 4000},
			{At: 100 * time.Minute, BetAmount: 8000},
		}

		for i, c := range cases {
			t.Run(fmt.Sprint(c), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled, %v", i, blindAlerter.Alerts)
				}

				alert := blindAlerter.Alerts[i]
				helper.AssertScheduledAlert(t, alert, c)
			})
		}

	})

	t.Run("record a win for Howl", func(t *testing.T) {
		player := "Howl"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", player))

		playerStore := &mock.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		helper.AssertPlayerWin(t, playerStore, player)
	})

	t.Run("record a win for Sophie", func(t *testing.T) {
		player := "Sophie"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", player))

		playerStore := &mock.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		helper.AssertPlayerWin(t, playerStore, player)
	})

}
