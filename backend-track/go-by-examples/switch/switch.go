package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, "as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Wow it is a weekend")
	default:
		fmt.Println("well work hard! Weekday it is!")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("Its after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		case string:
			fmt.Println("I am string")
		case float32:
			fmt.Println("I am float 64")
		default:
			fmt.Printf("I am unique of type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI("nahom")
	whatAmI(84834)
	whatAmI(388949834.9398848)
	whatAmI(9832899834892898)

}

// kjsdjfkjkds
// kjdfjklk
// kjsdjfkjkds
// kjdsfjksd
// jdkfjkd
// jdkfjkdjkdfjkkjdfkj

// jdkfjkdjkdfjkkjdfkj
// kjfkjd
// kjsdjfkjkds
// kjdsfjksd
// bfiuauireuijdkfjkd

// jdkfjkdjkdfjkkjdfkj
// kjfkjd
// kjsdjfkjkds
// kjdsfjksd
// jdkfjkd
// fjdjsjdsj/ jdkfjkdjkdfjkkjdfkj
// kjfkjd
// kjdfjklk // kjdfjklk // kjdfjklk
