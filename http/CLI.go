package poker

import (
	"bufio"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in *bufio.Scanner
}

func NewCLI(store PlayerStore, in *bufio.Scanner) *CLI {
	return &CLI{store, in}
}

func (c *CLI) PlayPoker() {
	winner := exactWinner(c.readLine())
	c.playerStore.Record(winner)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}

func exactWinner(line string) string {
	return strings.Replace(line, " wins", "", 1)
}