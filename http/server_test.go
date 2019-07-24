package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"donmmi/test-driven/http"
)

func TestGetPlayerScore(t *testing.T) {
	store := &poker.StubPlayerStore{
		Score:map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		Calls:nil,
		League:nil,
	}
	server := poker.NewPlayerServer(store)
	t.Run("get Pepper's Score", func(t *testing.T) {
		request := poker.NewGetPlayerScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"
		poker.AssertScore(t, got, want)
	})

	t.Run("get Floyd's Score", func(t *testing.T) {
		request := poker.NewGetPlayerScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"
		poker.AssertScore(t, got, want)
	})

	t.Run("get non existing player's Score", func(t *testing.T) {
		request := poker.NewGetPlayerScoreRequest("Apollo")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound
		poker.AssertStatusCode(t, got, want)
	})
}

func TestRecord(t *testing.T) {
	store := &poker.StubPlayerStore{}
	server := poker.NewPlayerServer(store)
	t.Run("Record Pepper's Score", func(t *testing.T) {
		player := "Pepper"
		request := poker.NewRecordWinnerRequest(player)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusAccepted
		poker.AssertStatusCode(t, got, want)
		poker.AssertWinner(t, store, player)

		contentTypeGot := response.Header().Get("content-type")
		contentTypeWant := poker.ContentTypeJson
		if contentTypeGot != contentTypeWant {
			t.Errorf("got:[%s], expected:[%s]", contentTypeGot, contentTypeWant)
		}
	})
}

func TestLeague(t *testing.T) {
	wantedLeague := []poker.Player{
		{"Pepper", 20},
		{"Floyd", 10},
	}
	store := &poker.StubPlayerStore{
		Score:map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		Calls:nil,
		League:wantedLeague,
	}
	server := poker.NewPlayerServer(store)

	request := poker.NewGetLeagueRequest()
	response := httptest.NewRecorder()

	server.ServeHTTP(response, request)

	got := response.Code
	want := http.StatusOK
	poker.AssertStatusCode(t, got, want)

	gotLeague := poker.GetLeagueFromResponse(t, response.Body)
	poker.AssertLeague(t, gotLeague, wantedLeague)
}
