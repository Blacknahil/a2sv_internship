package main

import (
	"fmt"
	"slices"
	// "slices"
)

func main() {

	var s []string
	fmt.Println("uninit", s, s == nil, len(s) == 0)

	var t = make([]string, 3)
	fmt.Println("t=", t, "len", len(t), "cap", cap(t))

	t[0] = "a"
	t[1] = "b"
	t[2] = "c"
	t = append(t, "d")

	fmt.Println("t", t)

	r := make([]string, len(t))
	copy(r, t)

	fmt.Println("copied", r)

	l := r[1:4]
	fmt.Println("sliced", l)

	t1 := []string{"a", "s", "d"}
	t2 := []string{"a", "s", "d"}
	if slices.Equal(t1, t2) {
		fmt.Println("t1==t2")
	}

	// a two dimensional data with different inner slices
	twod := make([][]int, 3)

	for i := range 3 {
		innerL := i + 1
		twod[i] = make([]int, innerL)

		for j := range innerL {
			twod[i][j] = i + j
		}

	}

	fmt.Println("two D slices", twod)

}
