package main

import (
	"testing"
	"reflect"
	"time"
)

func SpyChecker(url string) bool {
	if url == "test" {
		return true
	} else {
		return false
	}
}

func SpySlowerChecker(_ string) bool {
	time.Sleep(time.Millisecond * 20)
	return true
}

func TestWebChecker(t *testing.T) {
	urls := []string{
		"test",
		"test2",
		"test3",
		"test4",
	}

	got := WebChecker(SpyChecker, urls)
	want := map[string]bool{
		"test":true,
		"test2":false,
		"test3":false,
		"test4":false,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:[%v], expected:[%v]", got, want)
	}
}

func BenchmarkSlowerChecker(b *testing.B) {
	urls := make([]string, 0)
	for i := 0; i < 100; i++ {
		urls = append(urls, "test")
	}

	for i := 0; i < b.N; i++ {
		WebChecker(SpySlowerChecker, urls)
	}
}