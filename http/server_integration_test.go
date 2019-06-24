package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestRecordAndRetrievingThem(t *testing.T) {
	//store := NewInMemoryPlayerStore()
	tmpFile, clean := createTmpFile(t, ``)
	defer clean()
	store, err := NewFileSystemPlayerStore(tmpFile)
	assertNoError(t, err)

	server := NewPlayerServer(store)

	// record 3 times
	player := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinnerRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinnerRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinnerRequest(player))

	t.Run("get player score", func(t *testing.T) {
		// get results
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetPlayerScoreRequest(player))
		assertStatusCode(t, response.Code, http.StatusOK)

		got := response.Body.String()
		want := "3"
		if got != want {
			t.Errorf("got:[%s], expected:[%s]", got, want)
		}
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLeagueRequest())

		assertStatusCode(t, response.Code, http.StatusOK)

		gotLeague := getLeagueFromResponse(t, response.Body)
		wantedLeague := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, gotLeague, wantedLeague)
	})
}
