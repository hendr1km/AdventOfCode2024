package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	pattern1 := `mul\((\d{1,3}),(\d{1,3})\)`
	pattern2 := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`

	res1 := Process(pattern1)
	res2 := Process2(pattern2)

	fmt.Println(res1)
	fmt.Println(res2)
}

func Process(pattern string) int {
	file, _ := os.Open("data3.txt")
	defer file.Close()

	re := regexp.MustCompile(pattern)

	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()

		matches := re.FindAllStringSubmatch(input, -1)

		for _, match := range matches {

			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			prod := num1 * num2
			res += prod
		}
	}

	return res
}

func Process2(pattern string) int {

	file, _ := os.Open("data3.txt")
	defer file.Close()
	re := regexp.MustCompile(pattern)

	res := 0

	scanner := bufio.NewScanner(file)

	open := true
	for scanner.Scan() {
		input := scanner.Text()
		matches := re.FindAllStringSubmatch(input, -1)

		for _, match := range matches {

			if match[0] == "don't()" {
				open = false
				continue
			} else if match[0] == "do()" {
				open = true
				continue
			} else if open == false {
				continue
			}

			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])

			prod := num1 * num2
			res += prod

		}
	}

	return res
}
