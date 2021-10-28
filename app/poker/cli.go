package poker

import (
	"bufio"
	"io"
	"learngowithtests/app/server"
	"strings"
	"time"
)

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
	c.scheduleBlindAlerts()
	userInput := c.readLine()
	c.playerStore.PostPlayerScore(extractWinner(userInput))
}

func (c *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		c.blindAlerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
