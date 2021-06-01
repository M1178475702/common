package main

import "strconv"


func main() {
	a := 1
	b := 2

	r := max(a, b)
	println(`max: `+strconv.Itoa(r))
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	var c int = -10e6
	return b
}