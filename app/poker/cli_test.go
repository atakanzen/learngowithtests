package poker_test

import (
	"fmt"
	"learngowithtests/app/helper"
	"learngowithtests/app/mock"
	"learngowithtests/app/poker"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	var dummySpyAlerter = &mock.SpyBlindAlerter{}

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

	t.Run("it schedules blind bet alert", func(t *testing.T) {
		in := strings.NewReader("Muslum wins\n")
		playerStore := &mock.StubPlayerStore{}
		blindAlerter := &mock.SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.Alerts) != 1 {
			t.Fatal("expected a blind alert but didn't get one")
		}
	})
}
