package main

import (
	"testing"
	"reflect"
	"fmt"
)

type People struct {
	Name string
	Profile Profile
}

type Profile struct {
	Country string
	Age int
}

func TestWalk(t *testing.T) {
	var got []string
	fn := func(input string) {
		got = append(got, input)
	}

	cases := []struct{
		Name string
		Value interface{}
		Expected []string
	} {
		{
			"struct with one string",
			struct {
				Name string
			}{
				"Chris",
			},
			[]string{"Chris"},
		},
		{
			"struct with non string",
			struct {
				Name string
				Age int
			}{
				"Chris",
				29,
			},
			[]string{"Chris"},
		},
		{
			"struct with multi string",
			struct {
				Name string
				Country string
			}{
				"Chris",
				"USA",
			},
			[]string{"Chris", "USA"},
		},
		{
			"pointer of struct with multi string",
			&struct {
				Name string
				Country string
				Age int
			}{
				"Chris","" +
					"USA",
				29,
			},
			[]string{"Chris", "USA"},
		},
		{
			"struct contain a struct",
			People{
				"Chris",
				Profile{
					"USA",
					29,
				},
			},
			[]string{"Chris", "USA"},
		},
		{
			"slice of struct",
			[]People{
				{
					"Chris",
					Profile{
						"USA",
						29,
					},
				},
				{
					"Tom",
					Profile{
						"China",
						30,
					},
				},
			},
			[]string{"Chris", "USA", "Tom", "China"},
		},
		{
			"array of struct",
			[2]People{
				{
					"Chris",
					Profile{
						"USA",
						29,
					},
				},
				{
					"Tom",
					Profile{
						"China",
						30,
					},
				},
			},
			[]string{"Chris", "USA", "Tom", "China"},
		},
	}

	for _, c := range cases {
		got = make([]string, 0)

		t.Run(fmt.Sprintf("%s", c.Name), func(t *testing.T) {
			Walk(c.Value, fn)

			if !reflect.DeepEqual(got, c.Expected) {
				t.Errorf("got:[%v], expected:[%v]", got, c.Expected)
			}
		})
	}

	t.Run("with map", func(t *testing.T) {
		got = make([]string, 0)
		x := map[string]interface{}{
			"Name":"Chris",
			"Age":29,
			"Country":"USA",
			"Hobbit":"swimming",
		}
		Walk(x, fn)

		for i := 0; i < len(got); i++ {
			assertContainValue(t, x, got[i])
		}
	})
}

func assertContainValue(t *testing.T, m map[string]interface{}, want string) bool {
	t.Helper()
	for _, v := range m {
		if want == v {
			return true
		}
	}
	return false
}