package poker

import (
	"net/http"
	"testing"
	"reflect"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type StubPlayerStore struct {
	Score  map[string]int
	Calls  []string
	League []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.Score[name]
}

func (s *StubPlayerStore) Record(name string) {
	s.Calls = append(s.Calls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.League
}

func GetLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	league, err := getLeague(body)
	if err != nil {
		t.Error(err)
	}
	return league
}

func NewGetPlayerScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func NewRecordWinnerRequest(winner string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", winner), nil)
	return req
}

func NewGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func AssertScore(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}

func AssertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%d], expected:[%d]", got, want)
	}
}

func AssertWinner(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.Calls) != 1 {
		t.Errorf("got:[%d], expected:[%d]", len(store.Calls), 1)
	}

	if store.Calls[0] != winner {
		t.Errorf("got:[%s], expected:[%s]", store.Calls[0], winner)
	}
}

func AssertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:[%v], expected:[%v]", got, want)
	}
}

func AssertPlayerScore(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%d], expected:[%d]", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("unexpected err:", err)
	}
}

func CreateTmpFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()
	tmpFile, err := ioutil.TempFile("", "db")
	if err != nil {
		panic(err)
	}
	_, err = tmpFile.Write([]byte(initialData))
	if err != nil {
		panic(err)
	}

	_, err = tmpFile.Seek(0 ,0)
	if err != nil {
		panic(err)
	}

	clean := func() {
		err = tmpFile.Close()
		if err != nil {
			panic(err)
		}
		err = os.Remove(tmpFile.Name())
		if err != nil {
			panic(err)
		}
	}
	return tmpFile, clean
}