package poker

import (
	"bufio"
	"io"
	"learngowithtests/app/server"
	"strings"
)

type CLI struct {
	playerStore server.PlayerStore
	in          *bufio.Scanner
}

func NewCLI(playerStore server.PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: playerStore,
		in:          bufio.NewScanner(in),
	}
}

func (c *CLI) PlayPoker() {
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
