package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Course struct {
	name string
	mark float64
}

func calculate_grade(hashMap *map[string]float64) float64 {
	total := 0.0
	n := 0

	for _, value := range *hashMap {
		total += value
		n += 1
	}

	return total / float64(n)
}

func check_input(input string, hashMap *map[string]float64) (bool, Course) {

	// trim any extra spaces or newlines
	input = strings.TrimSpace(input)

	//split the input into parts
	parts := strings.Fields(input)

	if len(parts) != 2 {
		return false, Course{}
	}

	subject := parts[0]
	_, exists := (*hashMap)[subject]
	if exists {
		fmt.Printf("\033[31m Already saved the result for the course \033[0m \n")
		return false, Course{}
	}
	mark, err := strconv.ParseFloat(parts[1], 64)
	if err != nil || mark < 0 || mark > 100 {
		return false, Course{}
	}

	return true, Course{name: subject, mark: mark}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("\033[31m Usage: go run grade_calculator number(int) \033[0m ")
	} else {

		subject_number, err := strconv.ParseInt(os.Args[1], 10, 64)

		if err != nil {
			fmt.Printf("\033[31m Error: %s \033[0m \n", err)
			return
		}

		hashMap := make(map[string]float64)
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter subject name and mark (e.g., Math 85) \n ")

		for i := int64(0); i < subject_number; i++ {
			// enter the subject name and mark we get
			fmt.Printf("input %d: ", i+1)
			input, _ := reader.ReadString('\n')

			fmt.Println("you have entered ", input)
			// check user input

			status, course := check_input(input, &hashMap)

			if !status {
				// if not valid prompt the user to add again
				// provide an exit mechanism
				fmt.Println("\033[31m Invalid input. Please enter in the format: <subject> <mark> \033[0m")
				i--
				continue

			}

			hashMap[course.name] = float64(course.mark)
			// if valid add
		}

		// fmt.Println("successfully added ")

		ans := calculate_grade(&hashMap)
		fmt.Printf("\033[34m ANSWER=%f \033[0m\n", ans)

	}
}
