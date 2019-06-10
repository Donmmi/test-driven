package main

import "strings"

func Repeat(c string, times int) string {
	repeat := ""
	for i := 0; i < times; i++ {
		repeat += c
	}
	return repeat
}

func RepeatStandLib(c string, times int) string {
	return strings.Repeat(c, times)
}

func main() {

}
