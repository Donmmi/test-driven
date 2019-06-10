package main

func Sum(a []int) int {
	s := 0
	for _, e := range a {
		s += e
	}
	return s
}

func SumAll(slices ...[]int) []int {
	r := make([]int, 0)
	for _, slice := range slices {
		r = append(r, Sum(slice))
	}
	return r
}

func SumTail(slices ...[]int) []int {
	r := make([]int, 0)
	for _, slice := range slices {
		if len(slice) > 0 {
			r = append(r, Sum(slice[1:]))
		} else {
			r = append(r, 0)
		}
	}
	return r
}

func main() {

}
