package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
)

type StubPlayerStore struct {
	score map[string]int
}

func (s *StubPlayerStore) getPlayerScore(name string) int {
	return s.score[name]
}

func TestGetPlayerScore(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{
			"Pepper":20,
			"Floyd":10,
		},
	}
	server := &PlayerServer{store:store}
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
}

func newGetPlayerScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func assertScore(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}