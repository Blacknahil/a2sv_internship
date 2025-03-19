package main

import (
	"fmt"
)

func main() {

	var arr [5]int
	fmt.Println(arr)

	arr[2] = 1
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[2])

	// length of the arrayconst
	fmt.Println("len=", len(arr))

	//one line array intialization
	b := [4]int{0, 1, 34, 43}
	fmt.Println("b", b)

	b_count := [...]int{848, 949, 9349, 939, 949}
	fmt.Println("b_count", b_count)

	c := [...]int{8, 5, 5: 400, 500}
	fmt.Println("c", c)

	// two dimensional arrays

	var twoD [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD", twoD)

	var twoD2 = [3][3]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println("twoD 2=", twoD2)

}

// fjhfjfsbhjshjshjsd
// jhsfdhgs
// kjdfjghjsdfg
// fbjgsj
// kjdfjghjsdfg
// kdbjksdjb/ fbjgsj
// jhkdfj// fjhfjfsbhjshjshjsd
// jhsfdhgs
// kjdfjghjsdfg
// fbjgsj
// jhkdfj// fjhfjfsbhjshjshjsd
// jhsfdhgs
// kjdfjghjsdfg
// fbjgsj
// jhkdfj// fjhfjfsbhjshjshjsd
// jhsfdhgs
// kjdfjghjsdfg
// fbjgsj
// jhkdfj// fjhfjfsbhjshjshjsd
// jhsfdhgs
// kjdfjghjsdfg
// fbjgsj
// jhkdfj// fjhfjfsbhjshjshjsd
// jhsfdhgs
// kjdfjghjsdfg
// fbjgsj
// jhkdfj// fjhfjfsbhjshjshjsd
// jhsfdhgs
// kjdfjghjsdfg
// fbjgsj
// jhkdfj// jhkdfj
