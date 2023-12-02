package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	m := make(map[string]string)
	m["1"] = "1"
	m["2"] = "2"
	m["3"] = "3"
	m["4"] = "4"
	m["5"] = "5"
	m["6"] = "6"
	m["7"] = "7"
	m["8"] = "8"
	m["9"] = "9"

	// Worded numbers
	m["one"] = "1"
	m["two"] = "2"
	m["three"] = "3"
	m["four"] = "4"
	m["five"] = "5"
	m["six"] = "6"
	m["seven"] = "7"
	m["eight"] = "8"
	m["nine"] = "9"

	container := "one two three four five six seven eight nine"

	total := 0

	for scanner.Scan() {
		str_num := ""
		text := scanner.Text()
		i := 0

		// Parse to only show numbers
		for i < len(text) {
			str_char := string(text[i])

			n := 1
			word_num := str_char
			for len(text) > i+n && strings.Contains(container, word_num+string(text[i+n])) {
				word_num += string(text[i+n])
				n++
			}

			if m[str_char] != "" {
				str_num += m[str_char]
				i++
			} else if m[word_num] != "" {
				str_num += m[word_num]
				i += n - 1
			} else {
				i++
			}
		}

		convert_num, _ := strconv.Atoi(string(str_num[0]) + string(str_num[len(str_num)-1]))
		total += convert_num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}
