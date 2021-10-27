package poker

import (
	"bufio"
	"io"
	"learngowithtests/app/server"
	"strings"
	"time"
)

type BlindAlerter interface {
	ScheduleAlertAt(time.Duration, int)
}

type CLI struct {
	playerStore  server.PlayerStore
	in           *bufio.Scanner
	blindAlerter BlindAlerter
}

func NewCLI(playerStore server.PlayerStore, in io.Reader, blindAlerter BlindAlerter) *CLI {
	return &CLI{
		playerStore:  playerStore,
		in:           bufio.NewScanner(in),
		blindAlerter: blindAlerter,
	}
}

func (c *CLI) PlayPoker() {
	c.blindAlerter.ScheduleAlertAt(5*time.Second, 100)
	userInput := c.readLine()
	c.playerStore.PostPlayerScore(extractWinner(userInput))
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
