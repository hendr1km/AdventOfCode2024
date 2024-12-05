package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("data1.txt")
	defer file.Close()

	var column1 []int
	var column2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		columns := strings.Fields(scanner.Text())

		c1, _ := strconv.Atoi(columns[0])
		c2, _ := strconv.Atoi(columns[1])

		column1 = append(column1, c1)
		column2 = append(column2, c2)
	}

	sort.Ints(column1)
	sort.Ints(column2)

	part1(column1, column2)
	part2(column1, column2)
}

func part1(column1, column2 []int) {
	sum := 0
	for i := 0; i < len(column1); i++ {
		diff := column1[i] - column2[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	fmt.Println(sum)
}

func part2(column1, column2 []int) {

	m2 := make(map[int]int)

	for _, num := range column2 {
		m2[num]++
	}

	sim := 0
	for _, num := range column1 {
		sim += num * m2[num]
	}

	fmt.Println(sim)

}
