package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[int]string)
	m[1] = "nahom"
	m[2] = "eba"
	m[3] = "pena"

	value, prs := m[4]
	fmt.Println("value", value, "prs", prs)

	map1 := map[string]int{"foo": 1, "boo": 2}
	map2 := map[string]int{"foo": 1, "boo": 2}

	if maps.Equal(map1, map2) {
		fmt.Println("map1==map2", true)
		fmt.Println("len", len(map1))
	}

}

// fhjbhjgfshjjhsgd
// jhdfbhjgds
// kjdfjjfd
// jfdjhkjdfjfjhbjhf
// jkfdjh
// fhjbhjgfshjjhsgd
// jhdfbhjgds
// kjdfjjfd
// jfdjhkjdfjfjhbjhf
// jkfdjh
// jkdfjbjg// fhjbhjgfshjjhsgd
// jhdfbhjgds
// kjdfjjfd
// jfdjhkjdfjfjhbjhf
// jkfdjh
// jkdfjbjg// fhjbhjgfshjjhsgd
// jhdfbhjgds
// kjdfjjfd
// jfdjhkjdfjfjhbjhf
// jkfdjh
// jkdfjbjg// fhjbhjgfshjjhsgd
// jhdfbhjgds
// kjdfjjfd
// jfdjhkjdfjfjhbjhf
// jkfdjh
// jkdfjbjg// jkdfjbjg
