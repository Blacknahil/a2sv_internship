package main

import "fmt"

func main() {

	nums := []int{2, 3, 4, 5}
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println("total", total)

	hash_map := map[string]string{"a": "almaz", "b": "bereket", "c": "chirst"}

	for k, v := range hash_map {
		fmt.Printf("%s --> %s\n", k, v)
	}

	for i, c := range "nahom" {
		fmt.Println("i", i, " c=", c)
	}
}
