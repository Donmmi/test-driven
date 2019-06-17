package main

import (
	"io"
	"os"
	"fmt"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (s *ConfigurableSleeper) Sleep() {
	time.Sleep(s.duration)
}

const countNum  = 3

func count(w io.Writer, sleeper Sleeper) {
	for i := countNum; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}

	fmt.Fprintln(w, "Go!")
}

func main() {
	sleeper := &ConfigurableSleeper{time.Second}
	count(os.Stdout, sleeper)
}
