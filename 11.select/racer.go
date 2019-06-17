package main

import (
	"net/http"
	"time"
	"errors"
)

const tenSecond  = time.Second * 10

var ErrTimeOut = errors.New("time out")

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecond)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", ErrTimeOut
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		http.Get(url)
		ch <- true
	}()

	return ch
}
