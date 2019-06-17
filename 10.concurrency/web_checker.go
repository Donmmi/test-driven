package main

import (
	"net/http"
	"fmt"
)

type Checker func(url string) bool

type CheckResult struct {
	url string
	status bool
}

func WebChecker(checker Checker, urls []string) map[string]bool {
	res := make(map[string]bool)
	ch := make(chan CheckResult)
	for _, url := range urls {
		go func(url string) {
			ch <- CheckResult{url:url, status:checker(url)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		checkResult := <-ch
		res[checkResult.url] = checkResult.status
	}

	return res
}

func myChecker(url string) bool {
	_, err := http.Get(url)
	if err != nil {
		return false
	}
	return true
}

func main() {
	urls := []string{
		"http://www.google.com",
		"http:/www.sina.com",
		"http://www.baidu.com",
	}
	fmt.Println(WebChecker(myChecker, urls))
}
