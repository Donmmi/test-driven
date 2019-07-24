package poker_test

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"donmmi/test-driven/http"
)

func TestRecordAndRetrievingThem(t *testing.T) {
	//store := NewInMemoryPlayerStore()
	tmpFile, clean := poker.CreateTmpFile(t, ``)
	defer clean()
	store, err := poker.NewFileSystemPlayerStore(tmpFile)
	poker.AssertNoError(t, err)

	server := poker.NewPlayerServer(store)

	// Record 3 times
	player := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), poker.NewRecordWinnerRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewRecordWinnerRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewRecordWinnerRequest(player))

	t.Run("get player Score", func(t *testing.T) {
		// get results
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetPlayerScoreRequest(player))
		poker.AssertStatusCode(t, response.Code, http.StatusOK)

		got := response.Body.String()
		want := "3"
		if got != want {
			t.Errorf("got:[%s], expected:[%s]", got, want)
		}
	})

	t.Run("get League", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetLeagueRequest())

		poker.AssertStatusCode(t, response.Code, http.StatusOK)

		gotLeague := poker.GetLeagueFromResponse(t, response.Body)
		wantedLeague := []poker.Player{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, gotLeague, wantedLeague)
	})
}
