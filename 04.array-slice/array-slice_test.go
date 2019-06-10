package main

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {
	t.Run("sum slice", func(t *testing.T) {
		got := Sum([]int{1,2,3,4})
		want := 10

		assertSum(t, got, want)
	})

	t.Run("sum all", func(t *testing.T) {
		got := SumAll([]int{1,2,3}, []int{4,5,6})
		want := []int{6, 15}

		checkSlice(t, got, want)
	})

	t.Run("sum tail", func(t *testing.T) {
		got := SumTail([]int{1,2,3}, []int{4,5,6})
		want := []int{5, 11}

		checkSlice(t, got, want)
	})

	t.Run("sum tail safe", func(t *testing.T) {
		got := SumTail([]int{}, []int{4,5,6})
		want := []int{0, 11}

		checkSlice(t, got, want)
	})
}

func assertSum(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got:[%d], expected:[%d]", got, want)
	}
}

func checkSlice(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:[%v], expected:[%v]", got, want)
	}
}
