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

	id_sum := 0

	for scanner.Scan() {
		id, colors := game_parser(scanner.Text())

		fmt.Printf("%d: ", id)
		for k, v := range colors {
			fmt.Printf("%s: %d, ", k, v)
		}

		// This is for part 1
		//   12 red
		//   13 green
		//   14 blue
		if is_game_possible(colors, 12, 13, 14) {
			id_sum += id
		}

		fmt.Println()
	}

	fmt.Println(id_sum)
}

func game_parser(line string) (int, map[string]int) {
	colors := make(map[string]int)
	split_colors := strings.Split(line, ": ")
	id, _ := strconv.Atoi(strings.Split(split_colors[0], " ")[1])

	for _, hand := range strings.Split(split_colors[1], "; ") {
		for _, color := range strings.Split(hand, ", ") {
			num_color := strings.Split(color, " ")

			num, _ := strconv.Atoi(num_color[0])

			// Want to get the max, rather than the sum cause a color can be
			//   in multiple hands.
			colors[num_color[1]] = max_int(num, colors[num_color[1]])
		}
	}

	return id, colors
}

func is_game_possible(colors map[string]int, red int, green int, blue int) bool {
	for k, v := range colors {
		if k == "red" && v > red ||
			k == "blue" && v > blue ||
			k == "green" && v > green {
			return false
		}
	}

	return true
}

func max_int(left int, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}
