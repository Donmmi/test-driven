package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestRecordAndRetrievingThem(t *testing.T) {
	store := &InMemoryPlayerStore{map[string]int{}}
	server := &PlayerServer{store:store}

	// record 3 times
	player := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinnerRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinnerRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newRecordWinnerRequest(player))

	// get results
	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetPlayerScoreRequest(player))
	assertStatusCode(t, response.Code, http.StatusOK)

	got := response.Body.String()
	want := "3"
	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}
