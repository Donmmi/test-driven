package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	score  map[string]int
	calls  []string
	league []Player
}

func (s *StubPlayerStore) getPlayerScore(name string) int {
	return s.score[name]
}

func (s *StubPlayerStore) record(name string) {
	s.calls = append(s.calls, name)
}

func (s *StubPlayerStore) getLeague() []Player {
	return s.league
}

func TestGetPlayerScore(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(store)
	t.Run("get Pepper's score", func(t *testing.T) {
		request := newGetPlayerScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"
		assertScore(t, got, want)
	})

	t.Run("get Floyd's score", func(t *testing.T) {
		request := newGetPlayerScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"
		assertScore(t, got, want)
	})

	t.Run("get non existing player's score", func(t *testing.T) {
		request := newGetPlayerScoreRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound
		assertStatusCode(t, got, want)
	})
}

func TestRecord(t *testing.T) {
	store := &StubPlayerStore{}
	server := NewPlayerServer(store)
	t.Run("record Pepper's score", func(t *testing.T) {
		player := "Pepper"
		request := newRecordWinnerRequest(player)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted
		assertStatusCode(t, got, want)
		assertWinner(t, store, player)

		contentTypeGot := response.Header().Get("content-type")
		contentTypeWant := ContentTypeJson
		if contentTypeGot != contentTypeWant {
			t.Errorf("got:[%s], expected:[%s]", contentTypeGot, contentTypeWant)
		}
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Pepper", 20},
		{"Floyd", 10},
	}
	store := &StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		wantedLeague,
	}
	server := NewPlayerServer(store)

	request := newGetLeagueRequest()
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	got := response.Code
	want := http.StatusOK
	assertStatusCode(t, got, want)

	gotLeague := getLeagueFromResponse(t, response.Body)
	assertLeague(t, gotLeague, wantedLeague)
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	var players []Player
	err := json.NewDecoder(body).Decode(&players)
	if err != nil {
		t.Fatal(err)
	}
	return players
}

func newGetPlayerScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newRecordWinnerRequest(winner string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", winner), nil)
	return req
}

func newGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func assertScore(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%d], expected:[%d]", got, want)
	}
}

func assertWinner(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.calls) != 1 {
		t.Errorf("got:[%d], expected:[%d]", len(store.calls), 1)
	}

	if store.calls[0] != winner {
		t.Errorf("got:[%s], expected:[%s]", store.calls[0], winner)
	}
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:[%v], expected:[%v]", got, want)
	}
}
