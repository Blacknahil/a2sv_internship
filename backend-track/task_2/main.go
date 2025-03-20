package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func get_input(reader *bufio.Reader) string {

	fmt.Printf("Input: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func palidrome_checker(input string) bool {
	i, j := 0, len(input)-1

	for i < j {
		if input[i] != input[j] {
			return false
		}
		i++
		j--
	}
	return true

}

func word_frequency_count(input string) *map[string]int {

	dic := make(map[string]int)

	// parts := strings.Fields(input)
	length := len(input)
	var cur string

	for i := 0; i < length; i++ {
		if unicode.IsLetter(rune(input[i])) {
			cur += string(unicode.ToLower(rune(input[i])))
		} else {
			if cur != "" {
				dic[cur] += 1
			}

			cur = ""
		}
	}

	if cur != "" {
		dic[cur] += 1
	}

	return &dic
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What do you want to do ?")
	fmt.Println("Enter 'p' to check word palidrome")
	fmt.Println("Enter 'f' to check word frequency ")
	fmt.Printf("choose: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "p" || input == "P" {
		fmt.Println("checking palidrome")
		string_input := get_input(reader)
		isPali := palidrome_checker(string_input)
		if isPali {
			fmt.Printf("word = %s is palidrome\n", string_input)
		} else {
			fmt.Printf("word = %s is not palidrome\n", string_input)
		}

	} else if input == "f" || input == "F" {
		fmt.Println("checking word frequency")
		string_input := get_input(reader)
		word_map := word_frequency_count(string_input)
		fmt.Println("frequency", *word_map)

	} else {
		fmt.Println("\033[31m Invalid input \033[0m")
	}

}

// kdjfbjdsjksdkj
//jkdfbjkgkdfbjdf jkdfdjkfjkfsdfjkfjksdgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjkdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf dfjkfjksdjkskjdkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
//jkdfbjkgkdfbjdf jkdfjgjksdjbfjksdbfjsjdhjsdjhjksdjk
// dhjbhjsjhsjhd
