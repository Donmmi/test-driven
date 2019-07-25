package poker_test

import (
	"testing"
	"strings"
	"donmmi/test-driven/http"
)

func Test_CLI(t *testing.T) {
	in := strings.NewReader(`pepper wins`)
	playerStore := &poker.StubPlayerStore{}
	cli := poker.NewCLI(playerStore, in)
	cli.PlayPoker()

	poker.AssertWinner(t, playerStore, "pepper")
}
