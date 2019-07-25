package poker

import (
	"bufio"
	"strings"
	"io"
)

type CLI struct {
	playerStore PlayerStore
	in *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{store, bufio.NewScanner(in)}
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