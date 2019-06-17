package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("get faster url", func(t *testing.T) {
		fastServer := NewServer(0)
		slowServer := NewServer(time.Millisecond * 20)

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		got, err := Racer(fastUrl, slowUrl)
		want := fastUrl

		assertNoError(t, err)
		assertUrl(t, got, want)
	})

	t.Run("if more than 10s, return an error", func(t *testing.T) {
		fastServer := NewServer(time.Millisecond * 30)
		slowServer := NewServer(time.Millisecond * 50)

		fastUrl := fastServer.URL
		slowUrl := slowServer.URL

		_, err := ConfigurableRacer(fastUrl, slowUrl, time.Millisecond * 20)
		assertError(t, err, ErrTimeOut)
	})
}

func NewServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}

func assertError(t *testing.T, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an err")
	}

	if err.Error() != want.Error() {
		t.Errorf("got:[%s], expected:[%s]", err.Error(), ErrTimeOut.Error())
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("unexpected error")
	}
}

func assertUrl(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%s], expected:[%s]", got, want)
	}
}