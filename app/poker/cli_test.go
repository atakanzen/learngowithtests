package poker_test

import (
	"fmt"
	"learngowithtests/app/helper"
	"learngowithtests/app/poker"
	"learngowithtests/app/store/mock"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record a win for Howl", func(t *testing.T) {
		player := "Howl"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", player))

		playerStore := &mock.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		helper.AssertPlayerWin(t, playerStore, player)
	})

	t.Run("record a win for Sophie", func(t *testing.T) {
		player := "Sophie"
		in := strings.NewReader(fmt.Sprintf("%s wins\n", player))

		playerStore := &mock.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		helper.AssertPlayerWin(t, playerStore, player)
	})
}
