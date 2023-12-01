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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Printf("== Day 1 ==\n    P1 -> ")
	process(1)
	fmt.Printf("    P2 -> ")
	process(2)
	fmt.Printf("\n")
}

func read_line_by_line(filename string) ([]string, error) {
	result := make([]string, 0)

	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Printf("Err: %s\n", err)
		return result, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Err: %s\n", err)
		return result, err
	}

	return result, nil
}

func process(mode int) {
	// mode = 1 = part 1, mode = 2 = part 2
	// since both of them are very similar with some minor changes, i decided to do it like this (´･_･`)
	line_helper := func(data string) string {
		rules := strings.NewReplacer(
			// weird little hack to deal with overlapping numbers such as "oneight"
			// not proud of this shit
			"one", "1e",
			"two", "2o",
			"three", "3e",
			"four", "4r",
			"five", "5e",
			"six", "6x",
			"seven", "7n",
			"eight", "8t",
			"nine", "9e",
		)

		data = rules.Replace(data)
		data = rules.Replace(data)
		return data
	}

	input, err := read_line_by_line("./input.txt")
	if err != nil {
		return
	}

	nums := make([]int, 0)

	for _, line := range input {
		if mode == 2 {
			line = line_helper(line)
		}

		pairs := [2]string{}
		encounter := 0
		for i := 0; i < (len(line)); i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				if encounter == 0 {
					pairs[0] = string(line[i])
					encounter++
				} else {
					pairs[1] = string(line[i])
					encounter++
				}
				if encounter == 1 {
					pairs[1] = pairs[0]
				}
			}
		}
		n, err := strconv.Atoi(pairs[0]+pairs[1])
		if err != nil {
			log.Printf("Err: %s\n", err)
			return
		}
		nums = append(nums, n)
	}

	ans := 0

	for _, item := range nums {
		ans += item
	}

	fmt.Println(ans)
}